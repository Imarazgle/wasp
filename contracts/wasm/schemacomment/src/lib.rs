// COPYRIGHT OF A TEST SCHEMA DEFINITION 1
// COPYRIGHT OF A TEST SCHEMA DEFINITION 2

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use schemacomment::*;
use wasmlib::*;

use crate::consts::*;
use crate::events::*;
use crate::params::*;
use crate::results::*;
use crate::state::*;
use crate::structs::*;
use crate::typedefs::*;

mod consts;
mod contract;
mod events;
mod params;
mod results;
mod state;
mod structs;
mod typedefs;

mod schemacomment;

const EXPORT_MAP: ScExportMap = ScExportMap {
    names: &[
    	FUNC_TEST_FUNC1,
    	VIEW_TEST_VIEW1,
	],
    funcs: &[
    	func_test_func1_thunk,
	],
    views: &[
    	view_test_view1_thunk,
	],
};

#[no_mangle]
fn on_call(index: i32) {
	ScExports::call(index, &EXPORT_MAP);
}

#[no_mangle]
fn on_load() {
    ScExports::export(&EXPORT_MAP);
}

pub struct TestFunc1Context {
	events:  SchemaCommentEvents,
	params: ImmutableTestFunc1Params,
	results: MutableTestFunc1Results,
	state: MutableSchemaCommentState,
}

fn func_test_func1_thunk(ctx: &ScFuncContext) {
	ctx.log("schemacomment.funcTestFunc1");
	let f = TestFunc1Context {
		events:  SchemaCommentEvents {},
		params: ImmutableTestFunc1Params { proxy: params_proxy() },
		results: MutableTestFunc1Results { proxy: results_proxy() },
		state: MutableSchemaCommentState { proxy: state_proxy() },
	};
	ctx.require(f.params.name().exists(), "missing mandatory name");
	ctx.require(f.params.value().exists(), "missing mandatory value");
	func_test_func1(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("schemacomment.funcTestFunc1 ok");
}

pub struct TestView1Context {
	params: ImmutableTestView1Params,
	results: MutableTestView1Results,
	state: ImmutableSchemaCommentState,
}

fn view_test_view1_thunk(ctx: &ScViewContext) {
	ctx.log("schemacomment.viewTestView1");
	let f = TestView1Context {
		params: ImmutableTestView1Params { proxy: params_proxy() },
		results: MutableTestView1Results { proxy: results_proxy() },
		state: ImmutableSchemaCommentState { proxy: state_proxy() },
	};
	ctx.require(f.params.name().exists(), "missing mandatory name");
	view_test_view1(ctx, &f);
	ctx.results(&f.results.proxy.kv_store);
	ctx.log("schemacomment.viewTestView1 ok");
}
