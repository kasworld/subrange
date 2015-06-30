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

import (
	"fmt"
)

// range integer base
type i64base struct {
	v     int64
	min   int64
	max   int64
	state State_Type
}

func (si i64base) String() string {
	return fmt.Sprintf("%v(%v~%v)", si.v, si.min, si.max)
}
func newi64base(v, min, max int64) i64base {
	rtn := i64base{
		v:   v,
		min: min,
		max: max,
	}
	return rtn
}
func (si i64base) GetValue() int64 {
	return si.v
}
func (si i64base) GetState() State_Type {
	return si.state
}
func (si i64base) GetRate() float64 {
	l := si.max - si.min
	vl := si.v - si.min
	return float64(vl) / float64(l)
}
func (si i64base) GetLen() int64 {
	return si.max - si.min + 1
}
func (si i64base) GetMax() int64 {
	return si.max
}
func (si i64base) GetMin() int64 {
	return si.min
}
func (si *i64base) ClearState() {
	si.state = SI_Normal
}

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

///

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
		si.state = SI_Nan
	case si.v < si.min:
		si.v = Wrap(si.v-si.min, si.GetLen()) + si.min
		si.state = SI_WrapedUnder
	case si.v > si.max:
		si.v = Wrap(si.v-si.min, si.GetLen()) + si.min
		si.state = SI_WrapedOver
	default:
		si.state = SI_Normal
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
