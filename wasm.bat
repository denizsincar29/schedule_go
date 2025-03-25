@echo off
rem this script builds the go code and copies the goroot wasm_exec.js file to the output directory

echo creating output directory
set OUTDIR=dist
mkdir %OUTDIR%

echo building go code
set GOOS=js
set GOARCH=wasm
go build -o %OUTDIR%/main.wasm ./main_wasm
set GOOS=
set GOARCH=

rem set goroot to go env GOROOT
for /f "tokens=*" %%i in ('go env GOROOT') do set GOROOT=%%i
echo GOROOT: %GOROOT%

rem copy the wasm_exec.js file to the output directory
echo copying wasm_exec.js and index.html to output directory
copy "%GOROOT%\lib\wasm\wasm_exec.js" %OUTDIR%\wasm_exec.js
copy index.html %OUTDIR%\index.html
echo "Build process complete.  Files located in %OUTDIR%"