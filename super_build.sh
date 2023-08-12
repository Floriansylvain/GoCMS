#!/bin/sh

export GOARCH=amd64
OUTPUT_DIR=./bin

GO_FILE_PATH="./main/main.go"
PROGRAM_NAME=GohCMS

# Compile for Windows
GOOS=windows go build -o "$OUTPUT_DIR/${PROGRAM_NAME}_windows.exe" "$GO_FILE_PATH"
if [ $? -ne 0 ]; then
    echo "Compilation for Windows failed."
    exit $?
fi

# Compile for Linux
GOOS=linux go build -o "$OUTPUT_DIR/${PROGRAM_NAME}_linux" "$GO_FILE_PATH"
if [ $? -ne 0 ]; then
    echo "Compilation for Linux failed."
    exit $?
fi

# Compile for macOS
GOOS=darwin go build -o "$OUTPUT_DIR/${PROGRAM_NAME}_darwin" "$GO_FILE_PATH"
if [ $? -ne 0 ]; then
    echo "Compilation for macOS failed."
    exit $?
fi

echo "Compilation successful for all platforms."
