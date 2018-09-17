// Copyright 2018, Goomba project Authors. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package orderedset_test

import (
	"testing"

	"github.com/goombaio/orderedset"
)

func BenchmarkOrderedSet_Add(b *testing.B) {
	s := orderedset.NewOrderedSet()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(i ^ 2)
	}
}

func BenchmarkOrderedSet_Remove(b *testing.B) {
	s := orderedset.NewOrderedSet()
	for i := 0; i < b.N; i++ {
		s.Add(i ^ 2)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Remove(i ^ 2)
	}
}

var resultBenchmarkOrderedSetContains bool

func BenchmarkOrderedSet_Contains(b *testing.B) {
	s := orderedset.NewOrderedSet()
	for i := 0; i < b.N; i++ {
		s.Add(i ^ 2)
	}

	var contains bool

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		contains = s.Contains(i ^ 2)
	}

	resultBenchmarkOrderedSetContains = contains
}

var resultBenchmarkOrderedSetEmpty bool

func BenchmarkOrderedSet_Empty(b *testing.B) {
	s := orderedset.NewOrderedSet()
	for i := 0; i < b.N; i++ {
		s.Add(i ^ 2)
	}

	var empty bool

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		empty = s.Empty()
	}

	resultBenchmarkOrderedSetEmpty = empty
}

var resultBenchmarkOrderedSetSize int

func BenchmarkOrderedSet_Size(b *testing.B) {
	s := orderedset.NewOrderedSet()
	for i := 0; i < b.N; i++ {
		s.Add(i ^ 2)
	}

	var size int

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		size = s.Size()
	}

	resultBenchmarkOrderedSetSize = size
}

var resultBenchmarkOrderedSetValues []interface{}

func BenchmarkOrderedSet_Values(b *testing.B) {
	s := orderedset.NewOrderedSet()
	for i := 0; i < b.N; i++ {
		s.Add(i ^ 2)
	}

	var values []interface{}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		values = s.Values()
	}

	resultBenchmarkOrderedSetValues = values
}

var resultBenchmarkOrderedSetString string

func BenchmarkOrderedSet_String(b *testing.B) {
	s := orderedset.NewOrderedSet()
	for i := 0; i < b.N; i++ {
		s.Add(i ^ 2)
	}

	var str string

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str = s.String()
	}

	resultBenchmarkOrderedSetString = str
}
