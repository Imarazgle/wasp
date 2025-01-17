// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

import * as wasmlib from '../index';
import * as sc from './index';

export class ControlAddressesCall {
    func:    wasmlib.ScView;
    results: sc.ImmutableControlAddressesResults = new sc.ImmutableControlAddressesResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewCallContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewControlAddresses);
    }
}

export class GetBlockInfoCall {
    func:    wasmlib.ScView;
    params:  sc.MutableGetBlockInfoParams = new sc.MutableGetBlockInfoParams(wasmlib.ScView.nilProxy);
    results: sc.ImmutableGetBlockInfoResults = new sc.ImmutableGetBlockInfoResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewCallContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewGetBlockInfo);
    }
}

export class GetEventsForBlockCall {
    func:    wasmlib.ScView;
    params:  sc.MutableGetEventsForBlockParams = new sc.MutableGetEventsForBlockParams(wasmlib.ScView.nilProxy);
    results: sc.ImmutableGetEventsForBlockResults = new sc.ImmutableGetEventsForBlockResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewCallContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewGetEventsForBlock);
    }
}

export class GetEventsForContractCall {
    func:    wasmlib.ScView;
    params:  sc.MutableGetEventsForContractParams = new sc.MutableGetEventsForContractParams(wasmlib.ScView.nilProxy);
    results: sc.ImmutableGetEventsForContractResults = new sc.ImmutableGetEventsForContractResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewCallContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewGetEventsForContract);
    }
}

export class GetEventsForRequestCall {
    func:    wasmlib.ScView;
    params:  sc.MutableGetEventsForRequestParams = new sc.MutableGetEventsForRequestParams(wasmlib.ScView.nilProxy);
    results: sc.ImmutableGetEventsForRequestResults = new sc.ImmutableGetEventsForRequestResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewCallContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewGetEventsForRequest);
    }
}

export class GetRequestIDsForBlockCall {
    func:    wasmlib.ScView;
    params:  sc.MutableGetRequestIDsForBlockParams = new sc.MutableGetRequestIDsForBlockParams(wasmlib.ScView.nilProxy);
    results: sc.ImmutableGetRequestIDsForBlockResults = new sc.ImmutableGetRequestIDsForBlockResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewCallContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewGetRequestIDsForBlock);
    }
}

export class GetRequestReceiptCall {
    func:    wasmlib.ScView;
    params:  sc.MutableGetRequestReceiptParams = new sc.MutableGetRequestReceiptParams(wasmlib.ScView.nilProxy);
    results: sc.ImmutableGetRequestReceiptResults = new sc.ImmutableGetRequestReceiptResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewCallContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewGetRequestReceipt);
    }
}

export class GetRequestReceiptsForBlockCall {
    func:    wasmlib.ScView;
    params:  sc.MutableGetRequestReceiptsForBlockParams = new sc.MutableGetRequestReceiptsForBlockParams(wasmlib.ScView.nilProxy);
    results: sc.ImmutableGetRequestReceiptsForBlockResults = new sc.ImmutableGetRequestReceiptsForBlockResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewCallContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewGetRequestReceiptsForBlock);
    }
}

export class IsRequestProcessedCall {
    func:    wasmlib.ScView;
    params:  sc.MutableIsRequestProcessedParams = new sc.MutableIsRequestProcessedParams(wasmlib.ScView.nilProxy);
    results: sc.ImmutableIsRequestProcessedResults = new sc.ImmutableIsRequestProcessedResults(wasmlib.ScView.nilProxy);

    public constructor(ctx: wasmlib.ScViewCallContext) {
        this.func = new wasmlib.ScView(ctx, sc.HScName, sc.HViewIsRequestProcessed);
    }
}

export class ScFuncs {
    static controlAddresses(ctx: wasmlib.ScViewCallContext): ControlAddressesCall {
        const f = new ControlAddressesCall(ctx);
        f.results = new sc.ImmutableControlAddressesResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    static getBlockInfo(ctx: wasmlib.ScViewCallContext): GetBlockInfoCall {
        const f = new GetBlockInfoCall(ctx);
        f.params = new sc.MutableGetBlockInfoParams(wasmlib.newCallParamsProxy(f.func));
        f.results = new sc.ImmutableGetBlockInfoResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    static getEventsForBlock(ctx: wasmlib.ScViewCallContext): GetEventsForBlockCall {
        const f = new GetEventsForBlockCall(ctx);
        f.params = new sc.MutableGetEventsForBlockParams(wasmlib.newCallParamsProxy(f.func));
        f.results = new sc.ImmutableGetEventsForBlockResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    static getEventsForContract(ctx: wasmlib.ScViewCallContext): GetEventsForContractCall {
        const f = new GetEventsForContractCall(ctx);
        f.params = new sc.MutableGetEventsForContractParams(wasmlib.newCallParamsProxy(f.func));
        f.results = new sc.ImmutableGetEventsForContractResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    static getEventsForRequest(ctx: wasmlib.ScViewCallContext): GetEventsForRequestCall {
        const f = new GetEventsForRequestCall(ctx);
        f.params = new sc.MutableGetEventsForRequestParams(wasmlib.newCallParamsProxy(f.func));
        f.results = new sc.ImmutableGetEventsForRequestResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    static getRequestIDsForBlock(ctx: wasmlib.ScViewCallContext): GetRequestIDsForBlockCall {
        const f = new GetRequestIDsForBlockCall(ctx);
        f.params = new sc.MutableGetRequestIDsForBlockParams(wasmlib.newCallParamsProxy(f.func));
        f.results = new sc.ImmutableGetRequestIDsForBlockResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    static getRequestReceipt(ctx: wasmlib.ScViewCallContext): GetRequestReceiptCall {
        const f = new GetRequestReceiptCall(ctx);
        f.params = new sc.MutableGetRequestReceiptParams(wasmlib.newCallParamsProxy(f.func));
        f.results = new sc.ImmutableGetRequestReceiptResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    static getRequestReceiptsForBlock(ctx: wasmlib.ScViewCallContext): GetRequestReceiptsForBlockCall {
        const f = new GetRequestReceiptsForBlockCall(ctx);
        f.params = new sc.MutableGetRequestReceiptsForBlockParams(wasmlib.newCallParamsProxy(f.func));
        f.results = new sc.ImmutableGetRequestReceiptsForBlockResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }

    static isRequestProcessed(ctx: wasmlib.ScViewCallContext): IsRequestProcessedCall {
        const f = new IsRequestProcessedCall(ctx);
        f.params = new sc.MutableIsRequestProcessedParams(wasmlib.newCallParamsProxy(f.func));
        f.results = new sc.ImmutableIsRequestProcessedResults(wasmlib.newCallResultsProxy(f.func));
        return f;
    }
}
