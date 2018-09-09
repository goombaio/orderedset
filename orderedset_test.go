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
	"fmt"
	"testing"

	"github.com/goombaio/orderedset"
)

type customType struct {
	foo string
}

func TestOrderedSet_Add(t *testing.T) {
	s := orderedset.NewOrderedSet()
	s.Add("e", "f", "g", "c", "d", "x", "b", "a")
	s.Add("b") //overwrite
	structValue := customType{"svalue"}
	s.Add(structValue)
	s.Add(&structValue)
	s.Add(true)

	actualOutput := s.Values()
	expectedOutput := []interface{}{"e", "f", "g", "c", "d", "x", "b", "a", structValue, &structValue, true}
	if !sameElements(actualOutput, expectedOutput) {
		t.Errorf("Got %v expected %v", actualOutput, expectedOutput)
	}
}

func TestOrderedSet_Remove(t *testing.T) {
	s := orderedset.NewOrderedSet()
	s.Add("e", "f", "g", "c", "d", "x", "b", "a")
	s.Add("b") //overwrite
	structValue := customType{"svalue"}
	s.Add(structValue)
	s.Add(&structValue)
	s.Add(true)

	s.Remove("f", "g", &structValue, true)

	actualOutput := s.Values()
	expectedOutput := []interface{}{"e", "c", "d", "x", "b", "a", structValue}
	if !sameElements(actualOutput, expectedOutput) {
		t.Errorf("Got %v expected %v", actualOutput, expectedOutput)
	}

	// already removed
	s.Remove("f", "g", &structValue, true)
	actualOutput = s.Values()
	expectedOutput = []interface{}{"e", "c", "d", "x", "b", "a", structValue}
	if !sameElements(actualOutput, expectedOutput) {
		t.Errorf("Got %v expected %v", actualOutput, expectedOutput)
	}
}

func TestOrderedSet_Contains(t *testing.T) {
	s := orderedset.NewOrderedSet()
	s.Add("e", "f", "g", "c", "d", "x", "b", "a")
	s.Add("b") //overwrite
	structValue := customType{"svalue"}
	s.Add(structValue)
	s.Add(&structValue)
	s.Add(true)

	table := []struct {
		input          []interface{}
		expectedOutput bool
	}{
		{[]interface{}{"c", "d", &structValue}, true},
		{[]interface{}{"c", "d", nil}, false},
		{[]interface{}{true}, true},
		{[]interface{}{structValue}, true},
	}

	for _, test := range table {
		actualOutput := s.Contains(test.input...)
		if actualOutput != test.expectedOutput {
			t.Errorf("Got %v expected %v", actualOutput, test.expectedOutput)
		}
	}
}

func TestOrderedSet_Empty(t *testing.T) {
	s := orderedset.NewOrderedSet()
	if empty := s.Empty(); !empty {
		t.Errorf("Got %v expected %v", empty, true)
	}
	s.Add("e", "f", "g", "c", "d", "x", "b", "a")
	if empty := s.Empty(); empty {
		t.Errorf("Got %v expected %v", empty, false)
	}
	s.Remove("e", "f", "g", "c", "d", "x", "b", "a")
	if empty := s.Empty(); !empty {
		t.Errorf("Got %v expected %v", empty, true)
	}
}

func TestOrderedSet_Values(t *testing.T) {
	s := orderedset.NewOrderedSet()
	s.Add("e", "f", "g", "c", "d", "x", "b", "a")
	s.Add("b") //overwrite
	structValue := customType{"svalue"}
	s.Add(structValue)
	s.Add(&structValue)
	s.Add(true)

	actualOutput := s.Values()
	expectedOutput := []interface{}{"e", "f", "g", "c", "d", "x", "b", "a", structValue, &structValue, true}
	if !sameElements(actualOutput, expectedOutput) {
		t.Errorf("Got %v expected %v", actualOutput, expectedOutput)
	}
}

func TestOrderedSet_Size(t *testing.T) {
	s := orderedset.NewOrderedSet()
	if size := s.Size(); size != 0 {
		t.Errorf("Got %v expected %v", size, 0)
	}
	s.Add("e", "f", "g", "c", "d", "x", "b", "a")
	s.Add("e", "f", "g", "c", "d", "x", "b", "a", "z") // overwrite
	if size := s.Size(); size != 9 {
		t.Errorf("Got %v expected %v", size, 9)
	}
	s.Remove("e", "f", "g", "c", "d", "x", "b", "a", "z")
	if size := s.Size(); size != 0 {
		t.Errorf("Got %v expected %v", size, 0)
	}
}

func TestOrderedSet_Stringer(t *testing.T) {
	s := orderedset.NewOrderedSet()
	s.Add("foo", "bar")
	expected := "[foo bar]"
	result := fmt.Sprintf("%s", s)
	if expected != result {
		t.Fatalf("OrderedSet_Stringer expected to be %q but got %q", expected, result)
	}
}

func sameElements(a []interface{}, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for _, av := range a {
		found := false
		for _, bv := range b {
			if av == bv {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
