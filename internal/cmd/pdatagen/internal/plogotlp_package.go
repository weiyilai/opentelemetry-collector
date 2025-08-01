// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "go.opentelemetry.io/collector/internal/cmd/pdatagen/internal"
import (
	"path/filepath"
)

var plogotlp = &Package{
	info: &PackageInfo{
		name: "plogotlp",
		path: filepath.Join("plog", "plogotlp"),
		imports: []string{
			`otlpcollectorlog "go.opentelemetry.io/collector/pdata/internal/data/protogen/collector/logs/v1"`,
		},
		testImports: []string{
			`"testing"`,
			``,
			`"github.com/stretchr/testify/assert"`,
			``,
			`"go.opentelemetry.io/collector/pdata/internal"`,
		},
	},
	structs: []baseStruct{
		exportLogsPartialSuccess,
	},
}

var exportLogsPartialSuccess = &messageStruct{
	structName:     "ExportPartialSuccess",
	description:    "// ExportPartialSuccess represents the details of a partially successful export request.",
	originFullName: "otlpcollectorlog.ExportLogsPartialSuccess",
	fields: []Field{
		&PrimitiveField{
			fieldName: "RejectedLogRecords",
			protoType: ProtoTypeInt64,
		},
		&PrimitiveField{
			fieldName: "ErrorMessage",
			protoType: ProtoTypeString,
		},
	},
}
