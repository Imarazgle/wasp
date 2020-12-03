// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0
package sandbox

import (
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/vm/vmcontext"
	"github.com/iotaledger/wasp/packages/vm/vmtypes"
)

func init() {
	vmcontext.NewSandboxView = newView
}

type sandboxView struct {
	vmctx *vmcontext.VMContext
}

func newView(vmctx *vmcontext.VMContext) vmtypes.SandboxView {
	return sandboxView{vmctx}
}

func (s sandboxView) Params() dict.Dict {
	return s.vmctx.Params()
}

func (s sandboxView) State() kv.KVStore {
	return s.vmctx.State()
}

func (s sandboxView) MyBalances() coretypes.ColoredBalances {
	return s.vmctx.GetMyBalances()
}

func (s sandboxView) Call(contractHname coretypes.Hname, entryPoint coretypes.Hname, params dict.Dict) (dict.Dict, error) {
	return s.vmctx.Call(contractHname, entryPoint, params, nil)
}

func (s sandboxView) MyContractID() coretypes.ContractID {
	return s.vmctx.CurrentContractID()
}

func (s sandboxView) Event(msg string) {
	s.vmctx.EventPublisher().Publish(msg)
}

func (s sandboxView) Eventf(format string, args ...interface{}) {
	s.vmctx.EventPublisher().Publishf(format, args...)
}