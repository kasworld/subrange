// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
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

func Wrap(v, l int64) int64 {
	return (v%l + l) % l
}

// wrap around
type I64w struct {
	i64base
}

func NewI64w(v, min, max int64) *I64w {
	rtn := &I64w{
		newi64base(v, min, max),
	}
	rtn.Normalize()
	return rtn
}

func (si *I64w) Normalize() {
	switch {
	case si.max <= si.min:
		si.state = SR_Nan
	case si.v < si.min:
		si.v = Wrap(si.v-si.min, si.GetLen()) + si.min
		si.state = SR_WrapedUnder
	case si.v > si.max:
		si.v = Wrap(si.v-si.min, si.GetLen()) + si.min
		si.state = SR_WrapedOver
	default:
		si.state = SR_Normal
	}
}
func (si *I64w) SetMax(v int64) {
	si.max = v
	si.Normalize()
}
func (si *I64w) SetMin(v int64) {
	si.min = v
	si.Normalize()
}
func (si *I64w) SetValue(v int64) {
	si.v = v
	si.Normalize()
}
func (si *I64w) SetRate(v float64) {
	l := si.max - si.min
	si.v = si.min + int64(float64(l)*v)
	si.Normalize()
}
func (si *I64w) Add(v int64) {
	si.ClearState()
	si.v += v
	si.Normalize()
}
