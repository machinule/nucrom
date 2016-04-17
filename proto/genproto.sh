#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
rm -rf $DIR/gen/*
mkdir $DIR/gen/tmp
cp $DIR/src/*.proto $DIR/gen/tmp
cp $DIR/src/mechanic/*.proto $DIR/gen/tmp
protoc --go_out=$DIR/gen --proto_path=$DIR/gen/tmp $DIR/gen/tmp/*.proto
echo "Proto compiled"
