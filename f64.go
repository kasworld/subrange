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
package subrange

import (
	"fmt"
)

type F64 struct {
	v     float64
	min   float64
	max   float64
	state State_Type
}

func (si F64) String() string {
	return fmt.Sprintf("%5.2f(%5.2f~%5.2f)", si.v, si.min, si.max)
}
func NewF64(v, min, max float64) *F64 {
	rtn := &F64{
		v:   v,
		min: min,
		max: max,
	}
	rtn.Normalize()
	return rtn
}
func (si F64) GetValue() float64 {
	return si.v
}
func (si *F64) GetState() State_Type {
	return si.state
}
func (si F64) GetRate() float64 {
	l := si.max - si.min
	vl := si.v - si.min
	return float64(vl) / float64(l)
}
func (si F64) GetMax() float64 {
	return si.max
}
func (si F64) GetMin() float64 {
	return si.min
}
func (si *F64) SetMax(v float64) {
	si.max = v
	si.Normalize()
}
func (si *F64) SetMin(v float64) {
	si.min = v
	si.Normalize()
}
func (si *F64) SetValue(v float64) {
	si.v = v
	si.Normalize()
}
func (si *F64) SetRate(v float64) {
	l := si.max - si.min
	si.v = si.min + l*v
	si.Normalize()
}
func (si *F64) Normalize() {
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
func (si *F64) ClearState() {
	si.state = SI_Normal
}
func (si *F64) Add(v float64) {
	si.ClearState()
	si.v += v
	si.Normalize()
}
func (si *F64) Sub(v float64) {
	si.Add(-v)
}
