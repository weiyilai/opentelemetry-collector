// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pmetric

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.opentelemetry.io/collector/pdata/internal"
	otlpmetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/metrics/v1"
	"go.opentelemetry.io/collector/pdata/internal/json"
)

func TestGauge_MoveTo(t *testing.T) {
	ms := generateTestGauge()
	dest := NewGauge()
	ms.MoveTo(dest)
	assert.Equal(t, NewGauge(), ms)
	assert.Equal(t, generateTestGauge(), dest)
	dest.MoveTo(dest)
	assert.Equal(t, generateTestGauge(), dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.MoveTo(newGauge(&otlpmetrics.Gauge{}, &sharedState)) })
	assert.Panics(t, func() { newGauge(&otlpmetrics.Gauge{}, &sharedState).MoveTo(dest) })
}

func TestGauge_CopyTo(t *testing.T) {
	ms := NewGauge()
	orig := NewGauge()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestGauge()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newGauge(&otlpmetrics.Gauge{}, &sharedState)) })
}

func TestGauge_MarshalAndUnmarshalJSON(t *testing.T) {
	stream := json.BorrowStream(nil)
	defer json.ReturnStream(stream)
	src := generateTestGauge()
	src.marshalJSONStream(stream)
	require.NoError(t, stream.Error())

	// Append an unknown field at the start to ensure unknown fields are skipped
	// and the unmarshal logic continues.
	buf := stream.Buffer()
	assert.EqualValues(t, '{', buf[0])
	iter := json.BorrowIterator(append([]byte(`{"unknown": "string",`), buf[1:]...))
	defer json.ReturnIterator(iter)
	dest := NewGauge()
	dest.unmarshalJSONIter(iter)
	require.NoError(t, iter.Error())

	assert.Equal(t, src, dest)
}

func TestGauge_DataPoints(t *testing.T) {
	ms := NewGauge()
	assert.Equal(t, NewNumberDataPointSlice(), ms.DataPoints())
	fillTestNumberDataPointSlice(ms.DataPoints())
	assert.Equal(t, generateTestNumberDataPointSlice(), ms.DataPoints())
}

func generateTestGauge() Gauge {
	tv := NewGauge()
	fillTestGauge(tv)
	return tv
}

func fillTestGauge(tv Gauge) {
	fillTestNumberDataPointSlice(tv.DataPoints())
}
