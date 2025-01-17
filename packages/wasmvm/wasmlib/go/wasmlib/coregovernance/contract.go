// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

package coregovernance

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"

type AddAllowedStateControllerAddressCall struct {
	Func   *wasmlib.ScFunc
	Params MutableAddAllowedStateControllerAddressParams
}

type AddCandidateNodeCall struct {
	Func   *wasmlib.ScFunc
	Params MutableAddCandidateNodeParams
}

type ChangeAccessNodesCall struct {
	Func   *wasmlib.ScFunc
	Params MutableChangeAccessNodesParams
}

type ClaimChainOwnershipCall struct {
	Func *wasmlib.ScFunc
}

type DelegateChainOwnershipCall struct {
	Func   *wasmlib.ScFunc
	Params MutableDelegateChainOwnershipParams
}

type RemoveAllowedStateControllerAddressCall struct {
	Func   *wasmlib.ScFunc
	Params MutableRemoveAllowedStateControllerAddressParams
}

type RevokeAccessNodeCall struct {
	Func   *wasmlib.ScFunc
	Params MutableRevokeAccessNodeParams
}

type RotateStateControllerCall struct {
	Func   *wasmlib.ScFunc
	Params MutableRotateStateControllerParams
}

type SetChainInfoCall struct {
	Func   *wasmlib.ScFunc
	Params MutableSetChainInfoParams
}

type SetFeePolicyCall struct {
	Func   *wasmlib.ScFunc
	Params MutableSetFeePolicyParams
}

type GetAllowedStateControllerAddressesCall struct {
	Func    *wasmlib.ScView
	Results ImmutableGetAllowedStateControllerAddressesResults
}

type GetChainInfoCall struct {
	Func    *wasmlib.ScView
	Results ImmutableGetChainInfoResults
}

type GetChainNodesCall struct {
	Func    *wasmlib.ScView
	Results ImmutableGetChainNodesResults
}

type GetChainOwnerCall struct {
	Func    *wasmlib.ScView
	Results ImmutableGetChainOwnerResults
}

type GetFeePolicyCall struct {
	Func    *wasmlib.ScView
	Results ImmutableGetFeePolicyResults
}

type GetMaxBlobSizeCall struct {
	Func    *wasmlib.ScView
	Results ImmutableGetMaxBlobSizeResults
}

type Funcs struct{}

var ScFuncs Funcs

func (sc Funcs) AddAllowedStateControllerAddress(ctx wasmlib.ScFuncCallContext) *AddAllowedStateControllerAddressCall {
	f := &AddAllowedStateControllerAddressCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncAddAllowedStateControllerAddress)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

// access nodes
func (sc Funcs) AddCandidateNode(ctx wasmlib.ScFuncCallContext) *AddCandidateNodeCall {
	f := &AddCandidateNodeCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncAddCandidateNode)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

func (sc Funcs) ChangeAccessNodes(ctx wasmlib.ScFuncCallContext) *ChangeAccessNodesCall {
	f := &ChangeAccessNodesCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncChangeAccessNodes)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

// chain owner
func (sc Funcs) ClaimChainOwnership(ctx wasmlib.ScFuncCallContext) *ClaimChainOwnershipCall {
	return &ClaimChainOwnershipCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncClaimChainOwnership)}
}

func (sc Funcs) DelegateChainOwnership(ctx wasmlib.ScFuncCallContext) *DelegateChainOwnershipCall {
	f := &DelegateChainOwnershipCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncDelegateChainOwnership)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

func (sc Funcs) RemoveAllowedStateControllerAddress(ctx wasmlib.ScFuncCallContext) *RemoveAllowedStateControllerAddressCall {
	f := &RemoveAllowedStateControllerAddressCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncRemoveAllowedStateControllerAddress)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

func (sc Funcs) RevokeAccessNode(ctx wasmlib.ScFuncCallContext) *RevokeAccessNodeCall {
	f := &RevokeAccessNodeCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncRevokeAccessNode)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

// state controller
func (sc Funcs) RotateStateController(ctx wasmlib.ScFuncCallContext) *RotateStateControllerCall {
	f := &RotateStateControllerCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncRotateStateController)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

// chain info
func (sc Funcs) SetChainInfo(ctx wasmlib.ScFuncCallContext) *SetChainInfoCall {
	f := &SetChainInfoCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncSetChainInfo)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

// fees
func (sc Funcs) SetFeePolicy(ctx wasmlib.ScFuncCallContext) *SetFeePolicyCall {
	f := &SetFeePolicyCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncSetFeePolicy)}
	f.Params.Proxy = wasmlib.NewCallParamsProxy(&f.Func.ScView)
	return f
}

// state controller
func (sc Funcs) GetAllowedStateControllerAddresses(ctx wasmlib.ScViewCallContext) *GetAllowedStateControllerAddressesCall {
	f := &GetAllowedStateControllerAddressesCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetAllowedStateControllerAddresses)}
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.Proxy)
	return f
}

// chain info
func (sc Funcs) GetChainInfo(ctx wasmlib.ScViewCallContext) *GetChainInfoCall {
	f := &GetChainInfoCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetChainInfo)}
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.Proxy)
	return f
}

// access nodes
func (sc Funcs) GetChainNodes(ctx wasmlib.ScViewCallContext) *GetChainNodesCall {
	f := &GetChainNodesCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetChainNodes)}
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.Proxy)
	return f
}

// chain owner
func (sc Funcs) GetChainOwner(ctx wasmlib.ScViewCallContext) *GetChainOwnerCall {
	f := &GetChainOwnerCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetChainOwner)}
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.Proxy)
	return f
}

// fees
func (sc Funcs) GetFeePolicy(ctx wasmlib.ScViewCallContext) *GetFeePolicyCall {
	f := &GetFeePolicyCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetFeePolicy)}
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.Proxy)
	return f
}

func (sc Funcs) GetMaxBlobSize(ctx wasmlib.ScViewCallContext) *GetMaxBlobSizeCall {
	f := &GetMaxBlobSizeCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetMaxBlobSize)}
	wasmlib.NewCallResultsProxy(f.Func, &f.Results.Proxy)
	return f
}

var exportMap = wasmlib.ScExportMap{
	Names: []string{
		FuncAddAllowedStateControllerAddress,
		FuncAddCandidateNode,
		FuncChangeAccessNodes,
		FuncClaimChainOwnership,
		FuncDelegateChainOwnership,
		FuncRemoveAllowedStateControllerAddress,
		FuncRevokeAccessNode,
		FuncRotateStateController,
		FuncSetChainInfo,
		FuncSetFeePolicy,
		ViewGetAllowedStateControllerAddresses,
		ViewGetChainInfo,
		ViewGetChainNodes,
		ViewGetChainOwner,
		ViewGetFeePolicy,
		ViewGetMaxBlobSize,
	},
	Funcs: []wasmlib.ScFuncContextFunction{
		wasmlib.FuncError,
		wasmlib.FuncError,
		wasmlib.FuncError,
		wasmlib.FuncError,
		wasmlib.FuncError,
		wasmlib.FuncError,
		wasmlib.FuncError,
		wasmlib.FuncError,
		wasmlib.FuncError,
		wasmlib.FuncError,
	},
	Views: []wasmlib.ScViewContextFunction{
		wasmlib.ViewError,
		wasmlib.ViewError,
		wasmlib.ViewError,
		wasmlib.ViewError,
		wasmlib.ViewError,
		wasmlib.ViewError,
	},
}

func OnDispatch(index int32) {
	if index == -1 {
		exportMap.Export()
		return
	}

	panic("Calling core contract?")
}
