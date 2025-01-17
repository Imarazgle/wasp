// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

package coreaccounts

import (
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"
)

type ImmutableFoundryCreateNewResults struct {
	Proxy wasmtypes.Proxy
}

func (s ImmutableFoundryCreateNewResults) FoundrySN() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.Proxy.Root(ResultFoundrySN))
}

type MutableFoundryCreateNewResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableFoundryCreateNewResults() MutableFoundryCreateNewResults {
	return MutableFoundryCreateNewResults{Proxy: wasmlib.NewResultsProxy()}
}

func (s MutableFoundryCreateNewResults) FoundrySN() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.Proxy.Root(ResultFoundrySN))
}

type ArrayOfImmutableNftID struct {
	Proxy wasmtypes.Proxy
}

func (a ArrayOfImmutableNftID) Length() uint32 {
	return a.Proxy.Length()
}

func (a ArrayOfImmutableNftID) GetNftID(index uint32) wasmtypes.ScImmutableNftID {
	return wasmtypes.NewScImmutableNftID(a.Proxy.Index(index))
}

type ImmutableAccountNFTsResults struct {
	Proxy wasmtypes.Proxy
}

func (s ImmutableAccountNFTsResults) NftIDs() ArrayOfImmutableNftID {
	return ArrayOfImmutableNftID{Proxy: s.Proxy.Root(ResultNftIDs)}
}

type ArrayOfMutableNftID struct {
	Proxy wasmtypes.Proxy
}

func (a ArrayOfMutableNftID) AppendNftID() wasmtypes.ScMutableNftID {
	return wasmtypes.NewScMutableNftID(a.Proxy.Append())
}

func (a ArrayOfMutableNftID) Clear() {
	a.Proxy.ClearArray()
}

func (a ArrayOfMutableNftID) Length() uint32 {
	return a.Proxy.Length()
}

func (a ArrayOfMutableNftID) GetNftID(index uint32) wasmtypes.ScMutableNftID {
	return wasmtypes.NewScMutableNftID(a.Proxy.Index(index))
}

type MutableAccountNFTsResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableAccountNFTsResults() MutableAccountNFTsResults {
	return MutableAccountNFTsResults{Proxy: wasmlib.NewResultsProxy()}
}

func (s MutableAccountNFTsResults) NftIDs() ArrayOfMutableNftID {
	return ArrayOfMutableNftID{Proxy: s.Proxy.Root(ResultNftIDs)}
}

type MapAgentIDToImmutableBytes struct {
	Proxy wasmtypes.Proxy
}

func (m MapAgentIDToImmutableBytes) GetBytes(key wasmtypes.ScAgentID) wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(m.Proxy.Key(wasmtypes.AgentIDToBytes(key)))
}

type ImmutableAccountsResults struct {
	Proxy wasmtypes.Proxy
}

func (s ImmutableAccountsResults) AllAccounts() MapAgentIDToImmutableBytes {
	return MapAgentIDToImmutableBytes(s)
}

type MapAgentIDToMutableBytes struct {
	Proxy wasmtypes.Proxy
}

func (m MapAgentIDToMutableBytes) Clear() {
	m.Proxy.ClearMap()
}

func (m MapAgentIDToMutableBytes) GetBytes(key wasmtypes.ScAgentID) wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(m.Proxy.Key(wasmtypes.AgentIDToBytes(key)))
}

type MutableAccountsResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableAccountsResults() MutableAccountsResults {
	return MutableAccountsResults{Proxy: wasmlib.NewResultsProxy()}
}

func (s MutableAccountsResults) AllAccounts() MapAgentIDToMutableBytes {
	return MapAgentIDToMutableBytes(s)
}

type MapTokenIDToImmutableBigInt struct {
	Proxy wasmtypes.Proxy
}

func (m MapTokenIDToImmutableBigInt) GetBigInt(key wasmtypes.ScTokenID) wasmtypes.ScImmutableBigInt {
	return wasmtypes.NewScImmutableBigInt(m.Proxy.Key(wasmtypes.TokenIDToBytes(key)))
}

type ImmutableBalanceResults struct {
	Proxy wasmtypes.Proxy
}

func (s ImmutableBalanceResults) Balances() MapTokenIDToImmutableBigInt {
	return MapTokenIDToImmutableBigInt(s)
}

type MapTokenIDToMutableBigInt struct {
	Proxy wasmtypes.Proxy
}

func (m MapTokenIDToMutableBigInt) Clear() {
	m.Proxy.ClearMap()
}

func (m MapTokenIDToMutableBigInt) GetBigInt(key wasmtypes.ScTokenID) wasmtypes.ScMutableBigInt {
	return wasmtypes.NewScMutableBigInt(m.Proxy.Key(wasmtypes.TokenIDToBytes(key)))
}

type MutableBalanceResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableBalanceResults() MutableBalanceResults {
	return MutableBalanceResults{Proxy: wasmlib.NewResultsProxy()}
}

func (s MutableBalanceResults) Balances() MapTokenIDToMutableBigInt {
	return MapTokenIDToMutableBigInt(s)
}

type ImmutableBalanceBaseTokenResults struct {
	Proxy wasmtypes.Proxy
}

func (s ImmutableBalanceBaseTokenResults) Balance() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.Proxy.Root(ResultBalance))
}

type MutableBalanceBaseTokenResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableBalanceBaseTokenResults() MutableBalanceBaseTokenResults {
	return MutableBalanceBaseTokenResults{Proxy: wasmlib.NewResultsProxy()}
}

func (s MutableBalanceBaseTokenResults) Balance() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.Proxy.Root(ResultBalance))
}

type ImmutableBalanceNativeTokenResults struct {
	Proxy wasmtypes.Proxy
}

func (s ImmutableBalanceNativeTokenResults) Tokens() wasmtypes.ScImmutableBigInt {
	return wasmtypes.NewScImmutableBigInt(s.Proxy.Root(ResultTokens))
}

type MutableBalanceNativeTokenResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableBalanceNativeTokenResults() MutableBalanceNativeTokenResults {
	return MutableBalanceNativeTokenResults{Proxy: wasmlib.NewResultsProxy()}
}

func (s MutableBalanceNativeTokenResults) Tokens() wasmtypes.ScMutableBigInt {
	return wasmtypes.NewScMutableBigInt(s.Proxy.Root(ResultTokens))
}

type ImmutableFoundryOutputResults struct {
	Proxy wasmtypes.Proxy
}

func (s ImmutableFoundryOutputResults) FoundryOutputBin() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.Proxy.Root(ResultFoundryOutputBin))
}

type MutableFoundryOutputResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableFoundryOutputResults() MutableFoundryOutputResults {
	return MutableFoundryOutputResults{Proxy: wasmlib.NewResultsProxy()}
}

func (s MutableFoundryOutputResults) FoundryOutputBin() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.Proxy.Root(ResultFoundryOutputBin))
}

type ImmutableGetAccountNonceResults struct {
	Proxy wasmtypes.Proxy
}

func (s ImmutableGetAccountNonceResults) AccountNonce() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.Proxy.Root(ResultAccountNonce))
}

type MutableGetAccountNonceResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableGetAccountNonceResults() MutableGetAccountNonceResults {
	return MutableGetAccountNonceResults{Proxy: wasmlib.NewResultsProxy()}
}

func (s MutableGetAccountNonceResults) AccountNonce() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.Proxy.Root(ResultAccountNonce))
}

type MapTokenIDToImmutableBytes struct {
	Proxy wasmtypes.Proxy
}

func (m MapTokenIDToImmutableBytes) GetBytes(key wasmtypes.ScTokenID) wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(m.Proxy.Key(wasmtypes.TokenIDToBytes(key)))
}

type ImmutableGetNativeTokenIDRegistryResults struct {
	Proxy wasmtypes.Proxy
}

func (s ImmutableGetNativeTokenIDRegistryResults) Mapping() MapTokenIDToImmutableBytes {
	return MapTokenIDToImmutableBytes(s)
}

type MapTokenIDToMutableBytes struct {
	Proxy wasmtypes.Proxy
}

func (m MapTokenIDToMutableBytes) Clear() {
	m.Proxy.ClearMap()
}

func (m MapTokenIDToMutableBytes) GetBytes(key wasmtypes.ScTokenID) wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(m.Proxy.Key(wasmtypes.TokenIDToBytes(key)))
}

type MutableGetNativeTokenIDRegistryResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableGetNativeTokenIDRegistryResults() MutableGetNativeTokenIDRegistryResults {
	return MutableGetNativeTokenIDRegistryResults{Proxy: wasmlib.NewResultsProxy()}
}

func (s MutableGetNativeTokenIDRegistryResults) Mapping() MapTokenIDToMutableBytes {
	return MapTokenIDToMutableBytes(s)
}

type ImmutableNftDataResults struct {
	Proxy wasmtypes.Proxy
}

func (s ImmutableNftDataResults) NftData() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.Proxy.Root(ResultNftData))
}

type MutableNftDataResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableNftDataResults() MutableNftDataResults {
	return MutableNftDataResults{Proxy: wasmlib.NewResultsProxy()}
}

func (s MutableNftDataResults) NftData() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.Proxy.Root(ResultNftData))
}

type ImmutableTotalAssetsResults struct {
	Proxy wasmtypes.Proxy
}

func (s ImmutableTotalAssetsResults) Assets() MapTokenIDToImmutableBigInt {
	return MapTokenIDToImmutableBigInt(s)
}

type MutableTotalAssetsResults struct {
	Proxy wasmtypes.Proxy
}

func NewMutableTotalAssetsResults() MutableTotalAssetsResults {
	return MutableTotalAssetsResults{Proxy: wasmlib.NewResultsProxy()}
}

func (s MutableTotalAssetsResults) Assets() MapTokenIDToMutableBigInt {
	return MapTokenIDToMutableBigInt(s)
}
