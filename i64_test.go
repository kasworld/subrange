// Copyright 2015 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package subrange

import (
	"testing"
)

func TestWi(t *testing.T) {
	wi := NewI64w(0, 0, 10)
	for i := 0; i < 20; i++ {
		wi.Add(1)
		t.Logf("%v", wi)
	}
	for i := 0; i < 20; i++ {
		wi.Add(-1)
		t.Logf("%v", wi)
	}
}

func TestSi(t *testing.T) {
	wi := NewI64s(0, 0, 10)
	for i := 0; i < 20; i++ {
		wi.Add(1)
		t.Logf("%v", wi)
	}
	for i := 0; i < 20; i++ {
		wi.Add(-1)
		t.Logf("%v", wi)
	}
}

func TestFi(t *testing.T) {
	wi := NewF64(0, 0, 10)
	for i := 0; i < 20; i++ {
		wi.Add(1)
		t.Logf("%v", wi)
	}
	for i := 0; i < 20; i++ {
		wi.Add(-1)
		t.Logf("%v", wi)
	}
}

func TestNan(t *testing.T) {
	wi := I64w{}
	t.Logf("wi %v", wi.GetState() == SI_Nan)
	wi.Normalize()
	t.Logf("wi %v", wi.GetState() == SI_Nan)
}
