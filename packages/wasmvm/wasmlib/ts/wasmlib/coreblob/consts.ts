// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";

export const ScName        = "blob";
export const ScDescription = "Blob Contract";
export const HScName       = new wasmtypes.ScHname(0xfd91bc63);

export const ParamBlobs       = "this";
export const ParamDescription = "d";
export const ParamField       = "field";
export const ParamHash        = "hash";
export const ParamProgBinary  = "p";
export const ParamVMType      = "v";

export const ResultBlobSizes = "this";
export const ResultBytes     = "bytes";
export const ResultHash      = "hash";

export const FuncStoreBlob    = "storeBlob";
export const ViewGetBlobField = "getBlobField";
export const ViewGetBlobInfo  = "getBlobInfo";
export const ViewListBlobs    = "listBlobs";

export const HFuncStoreBlob    = new wasmtypes.ScHname(0xddd4c281);
export const HViewGetBlobField = new wasmtypes.ScHname(0x1f448130);
export const HViewGetBlobInfo  = new wasmtypes.ScHname(0xfde4ab46);
export const HViewListBlobs    = new wasmtypes.ScHname(0x62ca7990);
