#!/usr/bin/env bash

TEST_DIR=/tmp/gobl-test
SRC_DIR=$TEST_DIR/src

# start with a clean space
rm -rf $TEST_DIR

# copy the test source files
mkdir -p $SRC_DIR
cp -r a $SRC_DIR
mkdir $SRC_DIR/app
cp app.go_ $SRC_DIR/app/app.go

# build the .a files
export GOPATH=$TEST_DIR
export GOBIN=$TEST_DIR/bin
go install a/...

# pack them up
gobl pack -v a

# list the contents
FILES=`tar -tf $TEST_DIR/gobl-pkg/a.gobl`
echo contents of $TEST_DIR/gobl-pkg/a.gobl:
for f in $FILES; do
    echo "> $f"
done

# delete the original source & libs
rm -rf $SRC_DIR/a
rm -rf $TEST_DIR/pkg

# unpack the gobl file
gobl install -v $TEST_DIR/gobl-pkg/a.gobl

# build the consumer
go install app

# test the consumer
OUTPUT=`$TEST_DIR/bin/app`
if [ "$OUTPUT" != "From A: 14902" ]; then
    echo "FAIL: $OUTPUT"
    exit 1
else
    echo "SUCCESS: $OUTPUT"
fi

