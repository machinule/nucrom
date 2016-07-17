#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
MECHDIR=$DIR/mechanic
cd $MECHDIR/heat && go test
cd $MECHDIR/year && go test
cd $MECHDIR/province && go test
cd $MECHDIR/pseudorandom && go test
cd $MECHDIR/points && go test
cd $MECHDIR/leaders && go test
