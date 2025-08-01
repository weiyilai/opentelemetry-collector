// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pmetric

import (
	"iter"
	"sort"

	"go.opentelemetry.io/collector/pdata/internal"
	otlpmetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/metrics/v1"
	"go.opentelemetry.io/collector/pdata/internal/json"
)

// MetricSlice logically represents a slice of Metric.
//
// This is a reference type. If passed by value and callee modifies it, the
// caller will see the modification.
//
// Must use NewMetricSlice function to create new instances.
// Important: zero-initialized instance is not valid for use.
type MetricSlice struct {
	orig  *[]*otlpmetrics.Metric
	state *internal.State
}

func newMetricSlice(orig *[]*otlpmetrics.Metric, state *internal.State) MetricSlice {
	return MetricSlice{orig: orig, state: state}
}

// NewMetricSlice creates a MetricSlice with 0 elements.
// Can use "EnsureCapacity" to initialize with a given capacity.
func NewMetricSlice() MetricSlice {
	orig := []*otlpmetrics.Metric(nil)
	state := internal.StateMutable
	return newMetricSlice(&orig, &state)
}

// Len returns the number of elements in the slice.
//
// Returns "0" for a newly instance created with "NewMetricSlice()".
func (es MetricSlice) Len() int {
	return len(*es.orig)
}

// At returns the element at the given index.
//
// This function is used mostly for iterating over all the values in the slice:
//
//	for i := 0; i < es.Len(); i++ {
//	    e := es.At(i)
//	    ... // Do something with the element
//	}
func (es MetricSlice) At(i int) Metric {
	return newMetric((*es.orig)[i], es.state)
}

// All returns an iterator over index-value pairs in the slice.
//
//	for i, v := range es.All() {
//	    ... // Do something with index-value pair
//	}
func (es MetricSlice) All() iter.Seq2[int, Metric] {
	return func(yield func(int, Metric) bool) {
		for i := 0; i < es.Len(); i++ {
			if !yield(i, es.At(i)) {
				return
			}
		}
	}
}

// EnsureCapacity is an operation that ensures the slice has at least the specified capacity.
// 1. If the newCap <= cap then no change in capacity.
// 2. If the newCap > cap then the slice capacity will be expanded to equal newCap.
//
// Here is how a new MetricSlice can be initialized:
//
//	es := NewMetricSlice()
//	es.EnsureCapacity(4)
//	for i := 0; i < 4; i++ {
//	    e := es.AppendEmpty()
//	    // Here should set all the values for e.
//	}
func (es MetricSlice) EnsureCapacity(newCap int) {
	es.state.AssertMutable()
	oldCap := cap(*es.orig)
	if newCap <= oldCap {
		return
	}

	newOrig := make([]*otlpmetrics.Metric, len(*es.orig), newCap)
	copy(newOrig, *es.orig)
	*es.orig = newOrig
}

// AppendEmpty will append to the end of the slice an empty Metric.
// It returns the newly added Metric.
func (es MetricSlice) AppendEmpty() Metric {
	es.state.AssertMutable()
	*es.orig = append(*es.orig, &otlpmetrics.Metric{})
	return es.At(es.Len() - 1)
}

// MoveAndAppendTo moves all elements from the current slice and appends them to the dest.
// The current slice will be cleared.
func (es MetricSlice) MoveAndAppendTo(dest MetricSlice) {
	es.state.AssertMutable()
	dest.state.AssertMutable()
	// If they point to the same data, they are the same, nothing to do.
	if es.orig == dest.orig {
		return
	}
	if *dest.orig == nil {
		// We can simply move the entire vector and avoid any allocations.
		*dest.orig = *es.orig
	} else {
		*dest.orig = append(*dest.orig, *es.orig...)
	}
	*es.orig = nil
}

// RemoveIf calls f sequentially for each element present in the slice.
// If f returns true, the element is removed from the slice.
func (es MetricSlice) RemoveIf(f func(Metric) bool) {
	es.state.AssertMutable()
	newLen := 0
	for i := 0; i < len(*es.orig); i++ {
		if f(es.At(i)) {
			(*es.orig)[i] = nil
			continue
		}
		if newLen == i {
			// Nothing to move, element is at the right place.
			newLen++
			continue
		}
		(*es.orig)[newLen] = (*es.orig)[i]
		(*es.orig)[i] = nil
		newLen++
	}
	*es.orig = (*es.orig)[:newLen]
}

// CopyTo copies all elements from the current slice overriding the destination.
func (es MetricSlice) CopyTo(dest MetricSlice) {
	dest.state.AssertMutable()
	*dest.orig = copyOrigMetricSlice(*dest.orig, *es.orig)
}

// Sort sorts the Metric elements within MetricSlice given the
// provided less function so that two instances of MetricSlice
// can be compared.
func (es MetricSlice) Sort(less func(a, b Metric) bool) {
	es.state.AssertMutable()
	sort.SliceStable(*es.orig, func(i, j int) bool { return less(es.At(i), es.At(j)) })
}

// marshalJSONStream marshals all properties from the current struct to the destination stream.
func (ms MetricSlice) marshalJSONStream(dest *json.Stream) {
	dest.WriteArrayStart()
	if len(*ms.orig) > 0 {
		ms.At(0).marshalJSONStream(dest)
	}
	for i := 1; i < len(*ms.orig); i++ {
		dest.WriteMore()
		ms.At(i).marshalJSONStream(dest)
	}
	dest.WriteArrayEnd()
}

// unmarshalJSONIter unmarshals all properties from the current struct from the source iterator.
func (ms MetricSlice) unmarshalJSONIter(iter *json.Iterator) {
	iter.ReadArrayCB(func(iter *json.Iterator) bool {
		*ms.orig = append(*ms.orig, &otlpmetrics.Metric{})
		ms.At(ms.Len() - 1).unmarshalJSONIter(iter)
		return true
	})
}

func copyOrigMetricSlice(dest, src []*otlpmetrics.Metric) []*otlpmetrics.Metric {
	var newDest []*otlpmetrics.Metric
	if cap(dest) < len(src) {
		newDest = make([]*otlpmetrics.Metric, len(src))
		// Copy old pointers to re-use.
		copy(newDest, dest)
		// Add new pointers for missing elements from len(dest) to len(srt).
		for i := len(dest); i < len(src); i++ {
			newDest[i] = &otlpmetrics.Metric{}
		}
	} else {
		newDest = dest[:len(src)]
		// Cleanup the rest of the elements so GC can free the memory.
		// This can happen when len(src) < len(dest) < cap(dest).
		for i := len(src); i < len(dest); i++ {
			dest[i] = nil
		}
		// Add new pointers for missing elements.
		// This can happen when len(dest) < len(src) < cap(dest).
		for i := len(dest); i < len(src); i++ {
			newDest[i] = &otlpmetrics.Metric{}
		}
	}
	for i := range src {
		copyOrigMetric(newDest[i], src[i])
	}
	return newDest
}
