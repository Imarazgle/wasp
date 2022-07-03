// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:revive
package testwasmlib

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

const (
	ScName        = "testwasmlib"
	ScDescription = "Exercise several aspects of WasmLib"
	HScName       = wasmtypes.ScHname(0x89703a45)
)

const (
	ParamAddress         = "address"
	ParamAddressBytes    = "addressBytes"
	ParamAddressString   = "addressString"
	ParamAgentBytes      = "agentBytes"
	ParamAgentID         = "agentID"
	ParamAgentString     = "agentString"
	ParamBigInt          = "bigInt"
	ParamBigIntBytes     = "bigIntBytes"
	ParamBigIntString    = "bigIntString"
	ParamBlockIndex      = "blockIndex"
	ParamBool            = "bool"
	ParamBytes           = "bytes"
	ParamChainID         = "chainID"
	ParamEthAddress      = "ethAddress"
	ParamHash            = "hash"
	ParamHashBytes       = "hashBytes"
	ParamHashString      = "hashString"
	ParamHname           = "hname"
	ParamHnameBytes      = "hnameBytes"
	ParamHnameString     = "hnameString"
	ParamIndex           = "index"
	ParamIndex0          = "index0"
	ParamIndex1          = "index1"
	ParamInt16           = "int16"
	ParamInt32           = "int32"
	ParamInt64           = "int64"
	ParamInt8            = "int8"
	ParamKey             = "key"
	ParamKeyAddr         = "keyAddr"
	ParamLhs             = "lhs"
	ParamName            = "name"
	ParamNameAddr        = "nameAddr"
	ParamNftID           = "nftID"
	ParamNftIDBytes      = "nftIDBytes"
	ParamNftIDString     = "nftIDString"
	ParamParam           = "this"
	ParamRecordIndex     = "recordIndex"
	ParamRequestID       = "requestID"
	ParamRequestIDBytes  = "requestIDBytes"
	ParamRequestIDString = "requestIDString"
	ParamRhs             = "rhs"
	ParamScAddress       = "scAddress"
	ParamScAgentID       = "scAgentID"
	ParamScBigInt        = "scBigInt"
	ParamScHash          = "scHash"
	ParamScHname         = "scHname"
	ParamScNftID         = "scNftID"
	ParamScRequestID     = "scRequestID"
	ParamScTokenID       = "scTokenID"
	ParamShift           = "shift"
	ParamString          = "string"
	ParamTokenID         = "tokenID"
	ParamTokenIDBytes    = "tokenIDBytes"
	ParamTokenIDString   = "tokenIDString"
	ParamUint16          = "uint16"
	ParamUint32          = "uint32"
	ParamUint64          = "uint64"
	ParamUint8           = "uint8"
	ParamValue           = "value"
	ParamValueAddr       = "valueAddr"
)

const (
	ResultCount     = "count"
	ResultIotas     = "iotas"
	ResultLength    = "length"
	ResultQuo       = "quo"
	ResultRandom    = "random"
	ResultRecord    = "record"
	ResultRemainder = "remainder"
	ResultRes       = "res"
	ResultValue     = "value"
	ResultValueAddr = "valueAddr"
)

const (
	StateAddressMapOfAddressArray = "addressMapOfAddressArray"
	StateAddressMapOfAddressMap   = "addressMapOfAddressMap"
	StateArrayOfAddressArray      = "arrayOfAddressArray"
	StateArrayOfAddressMap        = "arrayOfAddressMap"
	StateArrayOfStringArray       = "arrayOfStringArray"
	StateArrayOfStringMap         = "arrayOfStringMap"
	StateLatLong                  = "latLong"
	StateRandom                   = "random"
	StateStringMapOfStringArray   = "stringMapOfStringArray"
	StateStringMapOfStringMap     = "stringMapOfStringMap"
)

const (
	FuncAddressMapOfAddressArrayAppend = "addressMapOfAddressArrayAppend"
	FuncAddressMapOfAddressArrayClear  = "addressMapOfAddressArrayClear"
	FuncAddressMapOfAddressArraySet    = "addressMapOfAddressArraySet"
	FuncAddressMapOfAddressMapClear    = "addressMapOfAddressMapClear"
	FuncAddressMapOfAddressMapSet      = "addressMapOfAddressMapSet"
	FuncArrayOfAddressArrayAppend      = "arrayOfAddressArrayAppend"
	FuncArrayOfAddressArrayClear       = "arrayOfAddressArrayClear"
	FuncArrayOfAddressArraySet         = "arrayOfAddressArraySet"
	FuncArrayOfAddressMapClear         = "arrayOfAddressMapClear"
	FuncArrayOfAddressMapSet           = "arrayOfAddressMapSet"
	FuncArrayOfStringArrayAppend       = "arrayOfStringArrayAppend"
	FuncArrayOfStringArrayClear        = "arrayOfStringArrayClear"
	FuncArrayOfStringArraySet          = "arrayOfStringArraySet"
	FuncArrayOfStringMapClear          = "arrayOfStringMapClear"
	FuncArrayOfStringMapSet            = "arrayOfStringMapSet"
	FuncParamTypes                     = "paramTypes"
	FuncRandom                         = "random"
	FuncStringMapOfStringArrayAppend   = "stringMapOfStringArrayAppend"
	FuncStringMapOfStringArrayClear    = "stringMapOfStringArrayClear"
	FuncStringMapOfStringArraySet      = "stringMapOfStringArraySet"
	FuncStringMapOfStringMapClear      = "stringMapOfStringMapClear"
	FuncStringMapOfStringMapSet        = "stringMapOfStringMapSet"
	FuncTakeAllowance                  = "takeAllowance"
	FuncTakeBalance                    = "takeBalance"
	FuncTriggerEvent                   = "triggerEvent"
	ViewAddressMapOfAddressArrayLength = "addressMapOfAddressArrayLength"
	ViewAddressMapOfAddressArrayValue  = "addressMapOfAddressArrayValue"
	ViewAddressMapOfAddressMapValue    = "addressMapOfAddressMapValue"
	ViewArrayOfAddressArrayLength      = "arrayOfAddressArrayLength"
	ViewArrayOfAddressArrayValue       = "arrayOfAddressArrayValue"
	ViewArrayOfAddressMapValue         = "arrayOfAddressMapValue"
	ViewArrayOfStringArrayLength       = "arrayOfStringArrayLength"
	ViewArrayOfStringArrayValue        = "arrayOfStringArrayValue"
	ViewArrayOfStringMapValue          = "arrayOfStringMapValue"
	ViewBigIntAdd                      = "bigIntAdd"
	ViewBigIntDiv                      = "bigIntDiv"
	ViewBigIntDivMod                   = "bigIntDivMod"
	ViewBigIntMod                      = "bigIntMod"
	ViewBigIntMul                      = "bigIntMul"
	ViewBigIntShl                      = "bigIntShl"
	ViewBigIntShr                      = "bigIntShr"
	ViewBigIntSub                      = "bigIntSub"
	ViewBlockRecord                    = "blockRecord"
	ViewBlockRecords                   = "blockRecords"
	ViewCheckAddress                   = "checkAddress"
	ViewCheckAgentID                   = "checkAgentID"
	ViewCheckBigInt                    = "checkBigInt"
	ViewCheckBool                      = "checkBool"
	ViewCheckBytes                     = "checkBytes"
	ViewCheckEthAddressAndAgentID      = "checkEthAddressAndAgentID"
	ViewCheckHash                      = "checkHash"
	ViewCheckHname                     = "checkHname"
	ViewCheckIntAndUint                = "checkIntAndUint"
	ViewCheckNftID                     = "checkNftID"
	ViewCheckRequestID                 = "checkRequestID"
	ViewCheckString                    = "checkString"
	ViewCheckTokenID                   = "checkTokenID"
	ViewGetRandom                      = "getRandom"
	ViewIotaBalance                    = "iotaBalance"
	ViewStringMapOfStringArrayLength   = "stringMapOfStringArrayLength"
	ViewStringMapOfStringArrayValue    = "stringMapOfStringArrayValue"
	ViewStringMapOfStringMapValue      = "stringMapOfStringMapValue"
)

const (
	HFuncAddressMapOfAddressArrayAppend = wasmtypes.ScHname(0x328fea8f)
	HFuncAddressMapOfAddressArrayClear  = wasmtypes.ScHname(0x771b8d2d)
	HFuncAddressMapOfAddressArraySet    = wasmtypes.ScHname(0x9b51ab5e)
	HFuncAddressMapOfAddressMapClear    = wasmtypes.ScHname(0x2a5316b1)
	HFuncAddressMapOfAddressMapSet      = wasmtypes.ScHname(0x7736a024)
	HFuncArrayOfAddressArrayAppend      = wasmtypes.ScHname(0xaf6e6b46)
	HFuncArrayOfAddressArrayClear       = wasmtypes.ScHname(0x8967252b)
	HFuncArrayOfAddressArraySet         = wasmtypes.ScHname(0xc84c8afa)
	HFuncArrayOfAddressMapClear         = wasmtypes.ScHname(0x18720c38)
	HFuncArrayOfAddressMapSet           = wasmtypes.ScHname(0xd02edb28)
	HFuncArrayOfStringArrayAppend       = wasmtypes.ScHname(0x2f37355a)
	HFuncArrayOfStringArrayClear        = wasmtypes.ScHname(0xcec430af)
	HFuncArrayOfStringArraySet          = wasmtypes.ScHname(0xa74fa352)
	HFuncArrayOfStringMapClear          = wasmtypes.ScHname(0x4ed3a6d7)
	HFuncArrayOfStringMapSet            = wasmtypes.ScHname(0xbe6a8ae7)
	HFuncParamTypes                     = wasmtypes.ScHname(0x6921c4cd)
	HFuncRandom                         = wasmtypes.ScHname(0xe86c97ca)
	HFuncStringMapOfStringArrayAppend   = wasmtypes.ScHname(0x414f806d)
	HFuncStringMapOfStringArrayClear    = wasmtypes.ScHname(0xe706a05f)
	HFuncStringMapOfStringArraySet      = wasmtypes.ScHname(0x5acf713b)
	HFuncStringMapOfStringMapClear      = wasmtypes.ScHname(0xf8edb8cb)
	HFuncStringMapOfStringMapSet        = wasmtypes.ScHname(0xa28e5cbb)
	HFuncTakeAllowance                  = wasmtypes.ScHname(0x91e7bd00)
	HFuncTakeBalance                    = wasmtypes.ScHname(0x8ad1cb27)
	HFuncTriggerEvent                   = wasmtypes.ScHname(0xd5438ac6)
	HViewAddressMapOfAddressArrayLength = wasmtypes.ScHname(0x30ecea55)
	HViewAddressMapOfAddressArrayValue  = wasmtypes.ScHname(0x2775d926)
	HViewAddressMapOfAddressMapValue    = wasmtypes.ScHname(0x3e49b429)
	HViewArrayOfAddressArrayLength      = wasmtypes.ScHname(0xc22488ab)
	HViewArrayOfAddressArrayValue       = wasmtypes.ScHname(0xb67ef99e)
	HViewArrayOfAddressMapValue         = wasmtypes.ScHname(0xc9ae7aec)
	HViewArrayOfStringArrayLength       = wasmtypes.ScHname(0x3eb0d035)
	HViewArrayOfStringArrayValue        = wasmtypes.ScHname(0xf2b0f6c8)
	HViewArrayOfStringMapValue          = wasmtypes.ScHname(0x8b35543c)
	HViewBigIntAdd                      = wasmtypes.ScHname(0x2b2c5e52)
	HViewBigIntDiv                      = wasmtypes.ScHname(0x84b2dadd)
	HViewBigIntDivMod                   = wasmtypes.ScHname(0x0bcaf323)
	HViewBigIntMod                      = wasmtypes.ScHname(0x4bbefd01)
	HViewBigIntMul                      = wasmtypes.ScHname(0x976cc699)
	HViewBigIntShl                      = wasmtypes.ScHname(0x64677f67)
	HViewBigIntShr                      = wasmtypes.ScHname(0xe5621518)
	HViewBigIntSub                      = wasmtypes.ScHname(0xf3c8cd33)
	HViewBlockRecord                    = wasmtypes.ScHname(0xad13b2f8)
	HViewBlockRecords                   = wasmtypes.ScHname(0x16e249ea)
	HViewCheckAddress                   = wasmtypes.ScHname(0x64b8add5)
	HViewCheckAgentID                   = wasmtypes.ScHname(0x27ae5137)
	HViewCheckBigInt                    = wasmtypes.ScHname(0x3497622b)
	HViewCheckBool                      = wasmtypes.ScHname(0xd8a476d2)
	HViewCheckBytes                     = wasmtypes.ScHname(0xcef76f43)
	HViewCheckEthAddressAndAgentID      = wasmtypes.ScHname(0xe01ad77e)
	HViewCheckHash                      = wasmtypes.ScHname(0x6b391dec)
	HViewCheckHname                     = wasmtypes.ScHname(0x7690daec)
	HViewCheckIntAndUint                = wasmtypes.ScHname(0xfe928771)
	HViewCheckNftID                     = wasmtypes.ScHname(0xcf852738)
	HViewCheckRequestID                 = wasmtypes.ScHname(0xeb1244dc)
	HViewCheckString                    = wasmtypes.ScHname(0xb2c3fb27)
	HViewCheckTokenID                   = wasmtypes.ScHname(0xeea106dc)
	HViewGetRandom                      = wasmtypes.ScHname(0x46263045)
	HViewIotaBalance                    = wasmtypes.ScHname(0x9d3920bd)
	HViewStringMapOfStringArrayLength   = wasmtypes.ScHname(0x9f433699)
	HViewStringMapOfStringArrayValue    = wasmtypes.ScHname(0xb263e298)
	HViewStringMapOfStringMapValue      = wasmtypes.ScHname(0xb93f2535)
)
