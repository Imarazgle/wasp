package state

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	valuetransaction "github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/transaction"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/sctransaction"
	"github.com/iotaledger/wasp/packages/variables"
	"io"
)

// represents an interface to the mutable state of the smart contract
type VariableState interface {
	// index of the current state. State index is incremented when state transition occurs
	// index 0 means origin state
	StateIndex() uint32
	ApplyStateIndex(uint32)
	// timestamp
	Timestamp() int64
	// updates state without changing state index
	ApplyStateUpdate(stateUpd StateUpdate)
	// applies batch of state updates, state index and timestamp
	ApplyBatch(Batch) error
	// commit means saving variable state to sc db, making it persistent
	CommitToDb(address address.Address, batch Batch) error
	// return hash of the variable state. It is a root of the Merkle chain of all
	// state updates starting from the origin
	Hash() hashing.HashValue
	// the storage of variable/value pairs
	Variables() variables.Variables
	String() string
}

// State update represents update to the variable state
// it is calculated by the VM (in batches)
// State updates comes in batches, all state updates within one batch
// has same state index, state tx id and batch size. ResultBatch index is unique in batch
// ResultBatch is completed when it contains one state update for each index
type StateUpdate interface {
	// request which resulted in this state update
	RequestId() *sctransaction.RequestId
	Timestamp() int64
	WithTimestamp(int64) StateUpdate
	// the payload of variables/values
	String() string
	Variables() variables.Variables
	Write(io.Writer) error
	Read(io.Reader) error
}

// ResultBatch of state updates applicable to the variable state by applying state updates
// in a sequence defined by batch indices
type Batch interface {
	ForEach(func(uint16, StateUpdate) bool)
	StateIndex() uint32
	WithStateIndex(uint32) Batch
	StateTransactionId() valuetransaction.ID
	WithStateTransaction(valuetransaction.ID) Batch
	Timestamp() int64
	Size() uint16
	RequestIds() []*sctransaction.RequestId
	EssenceHash() *hashing.HashValue // except state transaction id
	String() string
	Write(io.Writer) error
	Read(io.Reader) error
}
