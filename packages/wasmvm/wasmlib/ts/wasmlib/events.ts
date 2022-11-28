// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

import {ScFuncContext} from "./context";
import {uint64ToString} from "./wasmtypes/scuint64";
import {uint32FromString} from "./wasmtypes/scuint32";

export interface IEventHandler {
    callHandler(topic: string, params: string[]): void;
}

export class EventEncoder {
    event: string;

    constructor(eventName: string) {
        this.event = eventName;
        let timestamp = new ScFuncContext().timestamp();
        // convert nanoseconds to seconds
        this.encode(uint64ToString(timestamp / 1_000_000_000n));
    }

    emit(): void {
        new ScFuncContext().event(this.event);
    }

    encode(value: string): void {
        value = value.replaceAll("~", "~~")
        value = value.replaceAll("|", "~/")
        value = value.replaceAll(" ", "~_")
        this.event += "|" + value;
    }
}

export class EventDecoder {
    msg: string[];

    constructor(msg: string[]) {
        this.msg = msg;
    }

    decode(): string {
        return this.msg.shift()!;
    }

    timestamp(): u32 {
        return uint32FromString(this.decode());
    }
}
