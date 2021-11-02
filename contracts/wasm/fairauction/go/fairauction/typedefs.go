// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package fairauction

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type ArrayOfImmutableAgentID struct {
	objID int32
}

func (a ArrayOfImmutableAgentID) Length() int32 {
	return wasmlib.GetLength(a.objID)
}

func (a ArrayOfImmutableAgentID) GetAgentID(index int32) wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(a.objID, wasmlib.Key32(index))
}

type ImmutableBidderList = ArrayOfImmutableAgentID

type ArrayOfMutableAgentID struct {
	objID int32
}

func (a ArrayOfMutableAgentID) Clear() {
	wasmlib.Clear(a.objID)
}

func (a ArrayOfMutableAgentID) Length() int32 {
	return wasmlib.GetLength(a.objID)
}

func (a ArrayOfMutableAgentID) GetAgentID(index int32) wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(a.objID, wasmlib.Key32(index))
}

type MutableBidderList = ArrayOfMutableAgentID

type MapAgentIDToImmutableBid struct {
	objID int32
}

func (m MapAgentIDToImmutableBid) GetBid(key wasmlib.ScAgentID) ImmutableBid {
	return ImmutableBid{objID: m.objID, keyID: key.KeyID()}
}

type ImmutableBids = MapAgentIDToImmutableBid

type MapAgentIDToMutableBid struct {
	objID int32
}

func (m MapAgentIDToMutableBid) Clear() {
	wasmlib.Clear(m.objID)
}

func (m MapAgentIDToMutableBid) GetBid(key wasmlib.ScAgentID) MutableBid {
	return MutableBid{objID: m.objID, keyID: key.KeyID()}
}

type MutableBids = MapAgentIDToMutableBid