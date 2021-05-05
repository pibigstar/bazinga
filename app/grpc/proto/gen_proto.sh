#!/usr/bin/env bash

TARGET="./"
PROTO=""

if [ -n "$1" ]; then
    PROTO=$1
fi

for file in `find . -name '*.proto' -print`;
do
  if [ -n "$PROTO" ] && [[ "$file" != *"$PROTO"* ]]; then
      continue
  fi
	echo $file
	protoc --go_out=plugins=grpc,paths=import:$TARGET $file
done