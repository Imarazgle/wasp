// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:revive
package fairauction

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableFinalizeAuctionParams struct {
	proxy wasmtypes.Proxy
}

// NFT identifies the auction
func (s ImmutableFinalizeAuctionParams) Nft() wasmtypes.ScImmutableNftID {
	return wasmtypes.NewScImmutableNftID(s.proxy.Root(ParamNft))
}

type MutableFinalizeAuctionParams struct {
	proxy wasmtypes.Proxy
}

// NFT identifies the auction
func (s MutableFinalizeAuctionParams) Nft() wasmtypes.ScMutableNftID {
	return wasmtypes.NewScMutableNftID(s.proxy.Root(ParamNft))
}

type ImmutablePlaceBidParams struct {
	proxy wasmtypes.Proxy
}

// NFT identifies the auction
func (s ImmutablePlaceBidParams) Nft() wasmtypes.ScImmutableNftID {
	return wasmtypes.NewScImmutableNftID(s.proxy.Root(ParamNft))
}

type MutablePlaceBidParams struct {
	proxy wasmtypes.Proxy
}

// NFT identifies the auction
func (s MutablePlaceBidParams) Nft() wasmtypes.ScMutableNftID {
	return wasmtypes.NewScMutableNftID(s.proxy.Root(ParamNft))
}

type ImmutableSetOwnerMarginParams struct {
	proxy wasmtypes.Proxy
}

// new SC owner margin in promilles
func (s ImmutableSetOwnerMarginParams) OwnerMargin() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ParamOwnerMargin))
}

type MutableSetOwnerMarginParams struct {
	proxy wasmtypes.Proxy
}

// new SC owner margin in promilles
func (s MutableSetOwnerMarginParams) OwnerMargin() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ParamOwnerMargin))
}

type ImmutableStartAuctionParams struct {
	proxy wasmtypes.Proxy
}

// description of the NFTs being auctioned
func (s ImmutableStartAuctionParams) Description() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamDescription))
}

// duration of auction in minutes
func (s ImmutableStartAuctionParams) Duration() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamDuration))
}

// minimum required amount for any bid
func (s ImmutableStartAuctionParams) MinimumBid() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ParamMinimumBid))
}

type MutableStartAuctionParams struct {
	proxy wasmtypes.Proxy
}

// description of the NFTs being auctioned
func (s MutableStartAuctionParams) Description() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamDescription))
}

// duration of auction in minutes
func (s MutableStartAuctionParams) Duration() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamDuration))
}

// minimum required amount for any bid
func (s MutableStartAuctionParams) MinimumBid() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ParamMinimumBid))
}

type ImmutableGetAuctionInfoParams struct {
	proxy wasmtypes.Proxy
}

// NFT identifies the auction
func (s ImmutableGetAuctionInfoParams) Nft() wasmtypes.ScImmutableNftID {
	return wasmtypes.NewScImmutableNftID(s.proxy.Root(ParamNft))
}

type MutableGetAuctionInfoParams struct {
	proxy wasmtypes.Proxy
}

// NFT identifies the auction
func (s MutableGetAuctionInfoParams) Nft() wasmtypes.ScMutableNftID {
	return wasmtypes.NewScMutableNftID(s.proxy.Root(ParamNft))
}
