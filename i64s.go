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

// ranged value with saturated arithmetic
// min, max include
package subrange

// saturated range
type I64s struct {
	i64base
}

func NewI64s(v, min, max int64) *I64s {
	rtn := &I64s{
		newi64base(v, min, max),
	}
	rtn.Normalize()
	return rtn
}
func (si *I64s) Normalize() {
	switch {
	case si.max <= si.min:
		si.state = SI_Nan
	case si.v < si.min:
		si.v = si.min
		si.state = SI_Under
	case si.v > si.max:
		si.v = si.max
		si.state = SI_Over
	default:
		si.state = SI_Normal
	}
	return
}

func (si *I64s) SetMax(v int64) {
	si.max = v
	si.Normalize()
}
func (si *I64s) SetMin(v int64) {
	si.min = v
	si.Normalize()
}
func (si *I64s) SetValue(v int64) {
	si.v = v
	si.Normalize()
}
func (si *I64s) SetRate(v float64) {
	l := si.max - si.min
	si.v = si.min + int64(float64(l)*v)
	si.Normalize()
}
func (si *I64s) Add(v int64) {
	si.ClearState()
	si.v += v
	si.Normalize()
}
