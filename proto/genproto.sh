#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
rm -rf $DIR/gen/*
protoc --go_out=$DIR/gen --proto_path=$DIR/src $DIR/src/*.proto $DIR/src/mechanics/*.proto $DIR/src/systems/*.proto $DIR/src/common/*.proto
echo "Proto compiled"
