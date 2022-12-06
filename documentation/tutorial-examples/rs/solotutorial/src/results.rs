// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;

use crate::*;

#[derive(Clone)]
pub struct ImmutableGetStringResults {
    pub proxy: Proxy,
}

impl ImmutableGetStringResults {
    pub fn str(&self) -> ScImmutableString {
        ScImmutableString::new(self.proxy.root(RESULT_STR))
    }
}

#[derive(Clone)]
pub struct MutableGetStringResults {
    pub proxy: Proxy,
}

impl MutableGetStringResults {
    pub fn new() -> MutableGetStringResults {
        MutableGetStringResults {
            proxy: results_proxy(),
        }
    }

    pub fn str(&self) -> ScMutableString {
        ScMutableString::new(self.proxy.root(RESULT_STR))
    }
}
