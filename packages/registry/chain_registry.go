// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0
package registry

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/samber/lo"

	"github.com/iotaledger/hive.go/core/generics/onchangemap"
	"github.com/iotaledger/hive.go/core/ioutils"
	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/wasp/packages/cryptolib"
	"github.com/iotaledger/wasp/packages/isc"
	"github.com/iotaledger/wasp/packages/util"
)

// ChainRecord represents chain the node is participating in.
type ChainRecord struct {
	id          isc.ChainID
	Active      bool
	AccessNodes []*cryptolib.PublicKey
}

func NewChainRecord(chainID isc.ChainID, active bool) *ChainRecord {
	return &ChainRecord{
		id:          chainID,
		Active:      active,
		AccessNodes: []*cryptolib.PublicKey{},
	}
}

func (r *ChainRecord) ID() *isc.ChainID {
	return &r.id
}

func (r *ChainRecord) ChainID() isc.ChainID {
	return r.id
}

func (r *ChainRecord) Clone() onchangemap.Item[string, *isc.ChainID] {
	return &ChainRecord{
		id:          r.id,
		Active:      r.Active,
		AccessNodes: util.CloneSlice(r.AccessNodes),
	}
}

func (r *ChainRecord) AddAccessNode(pubKey *cryptolib.PublicKey) error {
	if lo.ContainsBy(r.AccessNodes, func(p *cryptolib.PublicKey) bool {
		return p.Equals(pubKey)
	}) {
		return fmt.Errorf("node is already an access node")
	}
	r.AccessNodes = append(r.AccessNodes, pubKey)
	return nil
}

func (r *ChainRecord) RemoveAccessNode(pubKey *cryptolib.PublicKey) (modified bool) {
	newAccessNodes := []*cryptolib.PublicKey{}
	for _, p := range r.AccessNodes {
		if p.Equals(pubKey) {
			modified = true
		} else {
			newAccessNodes = append(newAccessNodes, p)
		}
	}
	r.AccessNodes = newAccessNodes
	return modified
}

type jsonChainRecord struct {
	ChainID     string   `json:"chainID"`
	Active      bool     `json:"active"`
	AccessNodes []string `json:"accessNodes"`
}

func (r *ChainRecord) MarshalJSON() ([]byte, error) {
	accessNodesPubKeysHex := make([]string, 0)
	for _, accessNodePubKey := range r.AccessNodes {
		accessNodesPubKeysHex = append(accessNodesPubKeysHex, cryptolib.PublicKeyToHex(accessNodePubKey))
	}

	return json.Marshal(&jsonChainRecord{
		ChainID:     r.ID().String(),
		Active:      r.Active,
		AccessNodes: accessNodesPubKeysHex,
	})
}

func (r *ChainRecord) UnmarshalJSON(bytes []byte) error {
	j := &jsonChainRecord{}
	if err := json.Unmarshal(bytes, j); err != nil {
		return err
	}

	if j.ChainID == "" {
		return errors.New("missing chainID")
	}

	_, address, err := iotago.ParseBech32(j.ChainID)
	if err != nil {
		return err
	}

	if address.Type() != iotago.AddressAlias {
		return errors.New("chainID is not an alias address")
	}

	aliasAddress, ok := address.(*iotago.AliasAddress)
	if !ok {
		return errors.New("chainID is not an alias address")
	}

	accessNodesPubKeys := make([]*cryptolib.PublicKey, len(j.AccessNodes))
	for i, accessNodePubKeyHex := range j.AccessNodes {
		accessNodePubKey, err := cryptolib.NewPublicKeyFromHex(accessNodePubKeyHex)
		if err != nil {
			return err
		}

		accessNodesPubKeys[i] = accessNodePubKey
	}

	*r = ChainRecord{
		id:          isc.ChainID(aliasAddress.AliasID()),
		Active:      j.Active,
		AccessNodes: accessNodesPubKeys,
	}

	return nil
}

type ChainRecordRegistry struct {
	storeOnChangeMap *onchangemap.OnChangeMap[string, *isc.ChainID, *ChainRecord]
}

var _ ChainRecordRegistryProvider = &ChainRecordRegistry{}

// NewChainRecordRegistry creates new instance of the chain registry implementation.
func NewChainRecordRegistry(storeCallback func(chainRecords []*ChainRecord) error) *ChainRecordRegistry {
	return &ChainRecordRegistry{
		storeOnChangeMap: onchangemap.NewOnChangeMap[string, *isc.ChainID](storeCallback),
	}
}

func (p *ChainRecordRegistry) EnableStoreOnChange() {
	p.storeOnChangeMap.CallbackEnabled(true)
}

func (p *ChainRecordRegistry) ChainRecord(chainID isc.ChainID) (*ChainRecord, error) {
	chainRecord, err := p.storeOnChangeMap.Get(&chainID)
	if err != nil {
		// chain record doesn't exist
		return nil, nil
	}

	return chainRecord, nil
}

func (p *ChainRecordRegistry) ChainRecords() ([]*ChainRecord, error) {
	return lo.Values(p.storeOnChangeMap.All()), nil
}

func (p *ChainRecordRegistry) ForEachActiveChainRecord(consumer func(*ChainRecord) bool) error {
	chainRecords, err := p.ChainRecords()
	if err != nil {
		return err
	}

	for _, chainRecord := range chainRecords {
		if !chainRecord.Active {
			continue
		}

		if !consumer(chainRecord) {
			return nil
		}
	}

	return nil
}

func (p *ChainRecordRegistry) AddChainRecord(chainRecord *ChainRecord) error {
	return p.storeOnChangeMap.Add(chainRecord)
}

func (p *ChainRecordRegistry) DeleteChainRecord(chainID isc.ChainID) error {
	return p.storeOnChangeMap.Delete(&chainID)
}

// UpdateChainRecord modifies a ChainRecord in the Registry.
func (p *ChainRecordRegistry) UpdateChainRecord(chainID isc.ChainID, callback func(*ChainRecord) bool) (*ChainRecord, error) {
	return p.storeOnChangeMap.Modify(&chainID, callback)
}

func (p *ChainRecordRegistry) ActivateChainRecord(chainID isc.ChainID) (*ChainRecord, error) {
	return p.UpdateChainRecord(chainID, func(r *ChainRecord) bool {
		if r.Active {
			// chain was already active
			return false
		}
		r.Active = true
		return true
	})
}

func (p *ChainRecordRegistry) DeactivateChainRecord(chainID isc.ChainID) (*ChainRecord, error) {
	return p.UpdateChainRecord(chainID, func(r *ChainRecord) bool {
		if !r.Active {
			// chain was already disabled
			return false
		}
		r.Active = false
		return true
	})
}

type jsonChainRecords struct {
	ChainRecords []*ChainRecord `json:"chainRecords"`
}

func LoadChainRecordsJSONFromFile(filePath string, chainRecordsRegistry *ChainRecordRegistry) error {
	tmpChainRecords := &jsonChainRecords{}
	if err := ioutils.ReadJSONFromFile(filePath, tmpChainRecords); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("unable to unmarshal json file: %w", err)
	}

	for i := range tmpChainRecords.ChainRecords {
		if err := chainRecordsRegistry.AddChainRecord(tmpChainRecords.ChainRecords[i]); err != nil {
			return fmt.Errorf("unable to add ChainRecord to registry: %w", err)
		}
	}

	return nil
}

func WriteChainRecordsJSONToFile(filePath string, chainRecords []*ChainRecord) error {
	if err := os.MkdirAll(path.Dir(filePath), 0o770); err != nil {
		return fmt.Errorf("unable to create folder \"%s\": %w", path.Dir(filePath), err)
	}

	if err := ioutils.WriteJSONToFile(filePath, &jsonChainRecords{ChainRecords: chainRecords}, 0o600); err != nil {
		return fmt.Errorf("unable to marshal json file: %w", err)
	}

	return nil
}
