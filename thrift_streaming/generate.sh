#!/bin/bash

# Copyright 2023 CloudWeGo Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# Regenerate kitex_gen* directories when there's any related change to codegen (both kitex&thriftgo)

export PS4='[thrift_streaming][$(date "+%F %T")] '

cd `dirname $0`
ROOT=`pwd`

set -e
set -x

# Old binaries: kitex <= v0.8.0 && thriftgo <= v0.3.4
OLD=$ROOT/binaries/github-old

# New binaries: kitex >= v0.8.1 && thriftgo >= v0.3.5
NEW=$ROOT/binaries/github-new

# kitex >= v0.8.1 && thriftgo <= v0.3.4
NEW_THRIFTGO_OLD_KITEX=$ROOT/binaries/github-new-thriftgo-old-kitex

module='-module github.com/cloudwego/kitex-tests'
idl=idl/api.thrift

SAVE_PATH=$PATH

kitex_cmd() {
  kitex --no-dependency-check $@
}

# the env should be set in the main run.sh
if [ -z "$KITEX_LATEST_VERSION" ]; then
  export KITEX_LATEST_VERSION=`go list -m github.com/cloudwego/kitex@develop | cut -d" " -f2`
fi

# generate with old kitex and thriftgo WITHOUT thrift streaming support
function generate_old() {
    echo -e "\ngenerate_old\n"
    dir=$OLD
    export PATH=$OLD:$SAVE_PATH

    mkdir -p $dir
    if [ ! -f "$dir/kitex" ]; then
        GOBIN=$dir go install github.com/cloudwego/kitex/tool/cmd/kitex@v0.8.0
    fi
    if [ ! -f "$dir/thriftgo" ]; then
        GOBIN=$dir go install github.com/cloudwego/thriftgo@v0.3.4
    fi
    if [ ! -f "$dir/kitex" -o ! -f "$dir/thriftgo" ]; then
        echo "[old] Unable to install kitex or thriftgo to $dir, please check before continue."
        exit 1
    fi

    kitex -version

    # Thrift Old
    rm -rf kitex_gen_old
    kitex -gen-path kitex_gen_old $module $idl
}

function generate_new() {
    echo -e "\ngenerate_new\n"
    dir=$NEW
    export PATH=$dir:$SAVE_PATH

    mkdir -p $dir
    if [ -d "$LOCAL_REPO" ]; then
        SAVE_DIR=`pwd`
        cd $LOCAL_REPO/tool/cmd/kitex && go build && cp kitex $dir
        cd $SAVE_DIR
    else
        GOBIN=$dir go install github.com/cloudwego/kitex/tool/cmd/kitex@$KITEX_LATEST_VERSION
    fi
    if [ ! -f "$dir/thriftgo" ]; then
        GOBIN=$dir go install github.com/cloudwego/thriftgo@latest
    fi

    if [ ! -f "$dir/kitex" -o ! -f "$dir/thriftgo" ]; then
        echo "[new] Unable to install kitex or thriftgo to $dir, please check before continue."
        exit 1
    fi

    rm -rf kitex_gen
    kitex -version

    # Thrift
    kitex_cmd $module $idl
    kitex_cmd $module --combine-service idl/combine.thrift
    kitex_cmd $module --combine-service idl/combine_extend.thrift

    # Thrift Slim
    kitex_cmd -thrift template=slim -gen-path kitex_gen_slim $module $idl
    kitex_cmd -thrift template=slim -gen-path kitex_gen_slim $module --combine-service idl/combine.thrift
    kitex_cmd -thrift template=slim -gen-path kitex_gen_slim $module --combine-service idl/combine_extend.thrift

    # KitexPB
    kitex_cmd $module idl/api.proto

    # GRPC
    kitex_cmd $module idl/api_no_stream.proto
}

function generate_new_thriftgo_old_kitex() {
    echo -e "\ngenerate_new_thriftgo_old_kitex\n"
    dir=$NEW_THRIFTGO_OLD_KITEX
    export PATH=$dir:$SAVE_PATH

    mkdir -p $dir
    if [ ! -f "$dir/kitex" ]; then
        GOBIN=$dir go install github.com/cloudwego/kitex/tool/cmd/kitex@v0.8.0
    fi
    if [ ! -f "$dir/thriftgo" ]; then
        GOBIN=$dir go install github.com/cloudwego/thriftgo@latest
    fi
    if [ ! -f "$dir/kitex" -o ! -f "$dir/thriftgo" ]; then
        echo "[cross] Unable to install kitex or thriftgo to $dir, please check before continue."
        exit 1
    fi

    rm -rf kitex_gen_cross
    kitex -version
    # Thrift
    kitex -gen-path kitex_gen_cross $module $idl
}

if [ -d "$LOCAL_REPO" ]; then
    go mod edit -replace github.com/cloudwego/kitex=$LOCAL_REPO
    go mod edit -replace github.com/cloudwego/kitex/pkg/protocol/bthrift=${LOCAL_REPO}/pkg/protocol/bthrift
else
    go get github.com/cloudwego/kitex@$KITEX_LATEST_VERSION
fi

generate_new

generate_new_thriftgo_old_kitex

# regenerate kitex_gen_old (using kitex 0.8.0/thriftgo 0.3.4 without thrift-streaming support)
if [ ! -z "$TEST_GENERATE_OLD" ]; then
  generate_old
fi

cd exitserver
echo -e "\nbuilding exitserver @ $PWD\nPlease run ./update_go_mod.sh if any Go compatibility issue\n"
go build
mv exitserver $ROOT/binaries
