#!/bin/sh
# sudo apt-get install gcc-mingw-w64
 
RELATIVE_DIR=`dirname "$0"`
cd $RELATIVE_DIR
FILE_PATH=`pwd -P`
 
CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui -o  $FILE_PATH/server.exe
gio set -t 'string' $FILE_PATH/server.exe 'metadata::custom-icon' file:///$FILE_PATH/icon.png