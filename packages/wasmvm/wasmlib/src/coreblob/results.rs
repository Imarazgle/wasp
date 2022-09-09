// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use crate::coreblob::*;
use crate::*;

#[derive(Clone)]
pub struct ImmutableStoreBlobResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableStoreBlobResults {
    // calculated hash of blob set
    pub fn hash(&self) -> ScImmutableHash {
		ScImmutableHash::new(self.proxy.root(RESULT_HASH))
	}
}

#[derive(Clone)]
pub struct MutableStoreBlobResults {
	pub(crate) proxy: Proxy,
}

impl MutableStoreBlobResults {
    // calculated hash of blob set
    pub fn hash(&self) -> ScMutableHash {
		ScMutableHash::new(self.proxy.root(RESULT_HASH))
	}
}

#[derive(Clone)]
pub struct ImmutableGetBlobFieldResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableGetBlobFieldResults {
    // blob data
    pub fn bytes(&self) -> ScImmutableBytes {
		ScImmutableBytes::new(self.proxy.root(RESULT_BYTES))
	}
}

#[derive(Clone)]
pub struct MutableGetBlobFieldResults {
	pub(crate) proxy: Proxy,
}

impl MutableGetBlobFieldResults {
    // blob data
    pub fn bytes(&self) -> ScMutableBytes {
		ScMutableBytes::new(self.proxy.root(RESULT_BYTES))
	}
}

#[derive(Clone)]
pub struct MapStringToImmutableInt32 {
	pub(crate) proxy: Proxy,
}

impl MapStringToImmutableInt32 {
    pub fn get_int32(&self, key: &str) -> ScImmutableInt32 {
        ScImmutableInt32::new(self.proxy.key(&string_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct ImmutableGetBlobInfoResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableGetBlobInfoResults {
    // size for each named blob
    pub fn blob_sizes(&self) -> MapStringToImmutableInt32 {
		MapStringToImmutableInt32 { proxy: self.proxy.clone() }
	}
}

#[derive(Clone)]
pub struct MapStringToMutableInt32 {
	pub(crate) proxy: Proxy,
}

impl MapStringToMutableInt32 {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_int32(&self, key: &str) -> ScMutableInt32 {
        ScMutableInt32::new(self.proxy.key(&string_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct MutableGetBlobInfoResults {
	pub(crate) proxy: Proxy,
}

impl MutableGetBlobInfoResults {
    // size for each named blob
    pub fn blob_sizes(&self) -> MapStringToMutableInt32 {
		MapStringToMutableInt32 { proxy: self.proxy.clone() }
	}
}

#[derive(Clone)]
pub struct MapHashToImmutableInt32 {
	pub(crate) proxy: Proxy,
}

impl MapHashToImmutableInt32 {
    pub fn get_int32(&self, key: &ScHash) -> ScImmutableInt32 {
        ScImmutableInt32::new(self.proxy.key(&hash_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct ImmutableListBlobsResults {
	pub(crate) proxy: Proxy,
}

impl ImmutableListBlobsResults {
    // total size for each blob set
    pub fn blob_sizes(&self) -> MapHashToImmutableInt32 {
		MapHashToImmutableInt32 { proxy: self.proxy.clone() }
	}
}

#[derive(Clone)]
pub struct MapHashToMutableInt32 {
	pub(crate) proxy: Proxy,
}

impl MapHashToMutableInt32 {
    pub fn clear(&self) {
        self.proxy.clear_map();
    }

    pub fn get_int32(&self, key: &ScHash) -> ScMutableInt32 {
        ScMutableInt32::new(self.proxy.key(&hash_to_bytes(key)))
    }
}

#[derive(Clone)]
pub struct MutableListBlobsResults {
	pub(crate) proxy: Proxy,
}

impl MutableListBlobsResults {
    // total size for each blob set
    pub fn blob_sizes(&self) -> MapHashToMutableInt32 {
		MapHashToMutableInt32 { proxy: self.proxy.clone() }
	}
}
