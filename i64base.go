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

import "fmt"

// range integer base
type i64base struct {
	v     int64
	min   int64
	max   int64
	state SubRangeState
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
func (si i64base) GetState() SubRangeState {
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
	si.state = SR_Normal
}
