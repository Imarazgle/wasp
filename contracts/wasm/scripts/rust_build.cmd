@echo off
cd %1
if not exist schema.yaml goto :xit
echo Building %1
schema -rust %2
echo Compiling %1_main_bg.wasm
wasm-pack build rs\%1_main
:xit
cd ..
