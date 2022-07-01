// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

import * as wasmtypes from "wasmlib/wasmtypes";

export const ScName        = "donatewithfeedback";
export const ScDescription = "";
export const HScName       = new wasmtypes.ScHname(0x696d7f66);

export const ParamAmount   = "amount";
export const ParamFeedback = "feedback";
export const ParamNr       = "nr";
export const ParamOwner    = "owner";

export const ResultAmount        = "amount";
export const ResultCount         = "count";
export const ResultDonator       = "donator";
export const ResultError         = "error";
export const ResultFeedback      = "feedback";
export const ResultMaxDonation   = "maxDonation";
export const ResultTimestamp     = "timestamp";
export const ResultTotalDonation = "totalDonation";

export const StateLog           = "log";
export const StateMaxDonation   = "maxDonation";
export const StateOwner         = "owner";
export const StateTotalDonation = "totalDonation";

export const FuncDonate       = "donate";
export const FuncInit         = "init";
export const FuncWithdraw     = "withdraw";
export const ViewDonation     = "donation";
export const ViewDonationInfo = "donationInfo";

export const HFuncDonate       = new wasmtypes.ScHname(0xdc9b133a);
export const HFuncInit         = new wasmtypes.ScHname(0x1f44d644);
export const HFuncWithdraw     = new wasmtypes.ScHname(0x9dcc0f41);
export const HViewDonation     = new wasmtypes.ScHname(0xbdb245ba);
export const HViewDonationInfo = new wasmtypes.ScHname(0xc8f7c726);
