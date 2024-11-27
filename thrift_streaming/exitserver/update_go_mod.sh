#!/bin/bash
#
# Copyright 2024 CloudWeGo Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This script is used for updating go.mod
# Some packages like bytedance/sonic may not compatible with latest Go version,
# and we have to update go.mod regularly

rm -f go.mod
rm -f go.sum
go mod init github.com/cloudwego/kitex-tests/thrift_streaming/exitserver

# using legacy kitex version on purpose
go get github.com/cloudwego/kitex@v0.9.0-rc4
go get github.com/apache/thrift@v0.13.0
go get golang.org/x/net@v0.23.0 # fix security issues
go mod edit -replace github.com/apache/thrift=github.com/apache/thrift@v0.13.0

# updating tricky packages to make sure it works ...
go get github.com/bytedance/sonic@latest
go get github.com/choleraehyq/pid@latest

go mod tidy