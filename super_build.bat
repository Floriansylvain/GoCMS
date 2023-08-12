@echo off

setlocal

set "GOARCH=amd64"
set "OUTPUT_DIR=./bin"

set "GO_FILE_PATH=./main/main.go"
set "PROGRAM_NAME=GohCMS"

REM Compile for Windows
set "GOOS=windows"
go build -o "%OUTPUT_DIR%\%PROGRAM_NAME%_windows.exe" "%GO_FILE_PATH%"
if %errorlevel% neq 0 (
    echo Compilation for Windows failed.
    exit /b %errorlevel%
)

REM Compile for Linux
set "GOOS=linux"
go build -o "%OUTPUT_DIR%\%PROGRAM_NAME%_linux" "%GO_FILE_PATH%"
if %errorlevel% neq 0 (
    echo Compilation for Linux failed.
    exit /b %errorlevel%
)

REM Compile for macOS
set "GOOS=darwin"
go build -o "%OUTPUT_DIR%\%PROGRAM_NAME%_darwin" "%GO_FILE_PATH%"
if %errorlevel% neq 0 (
    echo Compilation for macOS failed.
    exit /b %errorlevel%
)

echo Compilation successful for all platforms.

endlocal
