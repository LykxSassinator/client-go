// Copyright 2021 TiKV Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// NOTE: The code in this file is based on code from the
// TiDB project, licensed under the Apache License v 2.0
//
// https://github.com/pingcap/tidb/tree/cc5e161ac06827589c4966674597c137cc9e809c/store/tikv/logutil/log.go
//

// Copyright 2021 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logutil

import (
	"context"
	"testing"

	"github.com/pingcap/log"
	"go.uber.org/zap"
)

// BgLogger returns the default global logger.
func BgLogger() *zap.Logger {
	return log.L()
}

// Logger gets a contextual logger from current context.
// contextual logger will output common fields from context.
func Logger(ctx context.Context) *zap.Logger {
	if ctxlogger, ok := ctx.Value(CtxLogKey).(*zap.Logger); ok {
		return ctxlogger
	}
	return log.L()
}

type ctxLogKeyType struct{}

// CtxLogKey is the key to retrieve logger from context.
// It can be assigned to another value.
var CtxLogKey interface{} = ctxLogKeyType{}

// AssertWarn panics when in testing mode, and logs a warning msg otherwise.
func AssertWarn(logger *zap.Logger, msg string, fields ...zap.Field) {
	if testing.Testing() {
		logger.Panic(msg, fields...)
	}
	logger.Warn(msg, fields...)
}
