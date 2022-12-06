// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

package solotutorial

import (
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"
)

type ImmutableStoreStringParams struct {
	proxy wasmtypes.Proxy
}

func NewImmutableStoreStringParams() ImmutableStoreStringParams {
	return ImmutableStoreStringParams{proxy: wasmlib.NewParamsProxy()}
}

func (s ImmutableStoreStringParams) Str() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamStr))
}

type MutableStoreStringParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableStoreStringParams) Str() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamStr))
}
