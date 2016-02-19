/*
** Copyright [2013-2016] [Megam Systems]
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
 */

package whmcs

import (
	"fmt"
	"time"

	"gopkg.in/check.v1"
)

func (s *S) TestStringify(c *check.C) {
	var nilPointer *string

	var tests = []struct {
		in  interface{}
		out string
	}{
		// basic types
		{"foo", `"foo"`},
		{123, `123`},
		{1.5, `1.5`},
		{false, `false`},
		{
			[]string{"a", "b"},
			`["a" "b"]`,
		},
		{
			struct {
				A []string
			}{nil},
			// nil slice is skipped
			`{}`,
		},
		{
			struct {
				A string
			}{"foo"},
			// structs not of a named type get no prefix
			`{A:"foo"}`,
		},

		// pointers
		{nilPointer, `<nil>`},
		{String("foo"), `"foo"`},
		{Int(123), `123`},
		{Bool(false), `false`},
		{
			[]*string{String("a"), String("b")},
			`["a" "b"]`,
		},

		// actual WHMCS structs
		{
			Timestamp{time.Date(2006, 01, 02, 15, 04, 05, 0, time.UTC)},
			`whmcs.Timestamp{2006-01-02 15:04:05 +0000 UTC}`,
		},
		{
			&Timestamp{time.Date(2006, 01, 02, 15, 04, 05, 0, time.UTC)},
			`whmcs.Timestamp{2006-01-02 15:04:05 +0000 UTC}`,
		},
		{
			Account{Email: String("joe@simpsons.com")},
			`whmcs.Account{Email:"joe@simpsons.com"}`,
		},
		{
			Order{ClientId: String("joe123")},
			`whmcs.Order{ClientId:"joe123"}`,
		},
	}

	for _, tt := range tests {
		s := Stringify(tt.in)
		c.Assert(s, check.Equals, tt.out)
	}
}

// Directly test the String() methods on various WHMCS types.  We don't do an
// exaustive test of all the various field types, since TestStringify() above
// takes care of that.  Rather, we just make sure that Stringify() is being
// used to build the strings, which we do by verifying that pointers are
// stringified as their underlying value.
func (s *S) TestString(c *check.C) {
	var tests = []struct {
		in  interface{}
		out string
	}{
		{Account{Email: String("n")}, `whmcs.Account{Email:"n"}`},
		{Order{ClientId: String("n")}, `whmcs.Order{ClientId:"n"}`},
		{BillableItem{ClientId: String("s")}, `whmcs.BillableItem{ClientId:"s"}`},
	}

	for _, tt := range tests {
		s := tt.in.(fmt.Stringer).String()
		if s != tt.out {
			c.Assert(s, check.Equals, tt.out)
		}
	}
}
