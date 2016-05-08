#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
rm -rf $DIR/gen/*
mkdir $DIR/gen/tmp
cp $DIR/src/*.proto $DIR/gen/tmp
cp $DIR/src/mechanic/*.proto $DIR/gen/tmp
cd $DIR/gen
protoc -I tmp/ tmp/*.proto --go_out=plugins=grpc:.
cd -
echo "Proto compiled"
