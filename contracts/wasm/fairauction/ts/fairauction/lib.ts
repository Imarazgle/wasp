// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmlib from "wasmlib";
import * as sc from "./index";

const exportMap: wasmlib.ScExportMap = {
	names: [
		sc.FuncFinalizeAuction,
		sc.FuncInit,
		sc.FuncPlaceBid,
		sc.FuncSetOwnerMargin,
		sc.FuncStartAuction,
		sc.ViewGetAuctionInfo,
	],
	funcs: [
		funcFinalizeAuctionThunk,
		funcInitThunk,
		funcPlaceBidThunk,
		funcSetOwnerMarginThunk,
		funcStartAuctionThunk,
	],
	views: [
		viewGetAuctionInfoThunk,
	],
};

export function on_call(index: i32): void {
	wasmlib.WasmVMHost.connect();
	wasmlib.ScExports.call(index, exportMap);
}

export function on_load(): void {
	wasmlib.WasmVMHost.connect();
	wasmlib.ScExports.export(exportMap);
}

function funcFinalizeAuctionThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("fairauction.funcFinalizeAuction");
	let f = new sc.FinalizeAuctionContext();

	// only SC itself can invoke this function
	ctx.require(ctx.caller().equals(ctx.accountID()), "no permission");

	ctx.require(f.params.nft().exists(), "missing mandatory nft");
	sc.funcFinalizeAuction(ctx, f);
	ctx.log("fairauction.funcFinalizeAuction ok");
}

function funcInitThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("fairauction.funcInit");
	let f = new sc.InitContext();
	sc.funcInit(ctx, f);
	ctx.log("fairauction.funcInit ok");
}

function funcPlaceBidThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("fairauction.funcPlaceBid");
	let f = new sc.PlaceBidContext();
	ctx.require(f.params.nft().exists(), "missing mandatory nft");
	sc.funcPlaceBid(ctx, f);
	ctx.log("fairauction.funcPlaceBid ok");
}

function funcSetOwnerMarginThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("fairauction.funcSetOwnerMargin");
	let f = new sc.SetOwnerMarginContext();

	// only SC creator can set owner margin
	const access = f.state.owner();
	ctx.require(access.exists(), "access not set: owner");
	ctx.require(ctx.caller().equals(access.value()), "no permission");

	ctx.require(f.params.ownerMargin().exists(), "missing mandatory ownerMargin");
	sc.funcSetOwnerMargin(ctx, f);
	ctx.log("fairauction.funcSetOwnerMargin ok");
}

function funcStartAuctionThunk(ctx: wasmlib.ScFuncContext): void {
	ctx.log("fairauction.funcStartAuction");
	let f = new sc.StartAuctionContext();
	ctx.require(f.params.minimumBid().exists(), "missing mandatory minimumBid");
	sc.funcStartAuction(ctx, f);
	ctx.log("fairauction.funcStartAuction ok");
}

function viewGetAuctionInfoThunk(ctx: wasmlib.ScViewContext): void {
	ctx.log("fairauction.viewGetAuctionInfo");
	let f = new sc.GetAuctionInfoContext();
	const results = new wasmlib.ScDict([]);
	f.results = new sc.MutableGetAuctionInfoResults(results.asProxy());
	ctx.require(f.params.nft().exists(), "missing mandatory nft");
	sc.viewGetAuctionInfo(ctx, f);
	ctx.results(results);
	ctx.log("fairauction.viewGetAuctionInfo ok");
}
