@echo off

SET BIN=nano-cli
SET INC_DIR=%~dp0nng\include
SET BUILD_DIR=%~dp0nng\build
SET LIB_PATH=%BUILD_DIR%\Debug
SET NNG_LIB=nng.lib
SET NNG_DLL=nng.dll

echo "Removing object and binary files..."
del *.obj
del *.dll
del *.exe

IF EXIST %BUILD_DIR% (
  echo "Removing build directory %BUILD_DIR%..."
  rmdir /Q /S %BUILD_DIR%
)
echo "Createing build directory %BUILD_DIR%..."
mkdir %BUILD_DIR%

echo "Changing into build dir %BUILD_DIR%..."
pushd %BUILD_DIR%

echo "Cmake init command..."
cmake -DBUILD_SHARED_LIBS=ON ..

echo "Cmake build command..."
cmake --build .

REM Back to the src directory
popd

echo "Copy dll to src directory..."
copy %LIB_PATH%\%NNG_DLL% .\%NNG_DLL%

echo "Building %BIN%.exe..."
cl.exe -I%INC_DIR% /DNNG_SHARED_LIB  /EHsc /Fe"%BIN%" serv.cpp /link %LIB_PATH%\%NNG_LIB% 
