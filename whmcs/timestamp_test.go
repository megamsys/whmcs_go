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
	"encoding/json"
	"fmt"
	"time"

	"gopkg.in/check.v1"
)

const (
	emptyTimeStr         = `"0001-01-01T00:00:00Z"`
	referenceTimeStr     = `"2006-01-02T15:04:05Z"`
	referenceUnixTimeStr = `1136214245`
)

var (
	referenceTime = time.Date(2006, 01, 02, 15, 04, 05, 0, time.UTC)
	unixOrigin    = time.Unix(0, 0).In(time.UTC)
)

func (s *S) TestTimestamp_Marshal(c *check.C) {
	testCases := []struct {
		desc    string
		data    Timestamp
		want    string
		wantErr bool
		equal   bool
	}{
		{"Reference", Timestamp{referenceTime}, referenceTimeStr, false, true},
		{"Empty", Timestamp{}, emptyTimeStr, false, true},
		{"Mismatch", Timestamp{}, referenceTimeStr, false, false},
	}
	for _, tc := range testCases {
		out, err := json.Marshal(tc.data)
		c.Assert(err, check.IsNil)
		got := string(out)
		c.Assert(got == tc.want, check.Equals, tc.equal)
	}
}

func (s *S) TestTimestamp_Unmarshal(c *check.C) {
	testCases := []struct {
		desc    string
		data    string
		want    Timestamp
		wantErr bool
		equal   bool
	}{
		{"Reference", referenceTimeStr, Timestamp{referenceTime}, false, true},
		{"ReferenceUnix", `1136214245`, Timestamp{referenceTime}, false, true},
		{"Empty", emptyTimeStr, Timestamp{}, false, true},
		{"UnixStart", `0`, Timestamp{unixOrigin}, false, true},
		{"Mismatch", referenceTimeStr, Timestamp{}, false, false},
		{"MismatchUnix", `0`, Timestamp{}, false, false},
		{"Invalid", `"asdf"`, Timestamp{referenceTime}, true, false},
	}
	for _, tc := range testCases {
		var got Timestamp
		err := json.Unmarshal([]byte(tc.data), &got)
		c.Assert(err != nil, check.Equals, tc.wantErr)
		c.Assert(got.Equal(tc.want), check.Equals, tc.equal)
	}
}

func (s *S) TestTimstamp_MarshalReflexivity(c *check.C) {
	testCases := []struct {
		desc string
		data Timestamp
	}{
		{"Reference", Timestamp{referenceTime}},
		{"Empty", Timestamp{}},
	}
	for _, tc := range testCases {
		data, err := json.Marshal(tc.data)
		c.Assert(err, check.IsNil)
		var got Timestamp
		err = json.Unmarshal(data, &got)
		tb := got.Equal(tc.data)
		c.Assert(tb, check.Equals, true)
	}
}

type WrappedTimestamp struct {
	A    int
	Time Timestamp
}

func (s *S) TestWrappedTimstamp_Marshal(c *check.C) {
	testCases := []struct {
		desc    string
		data    WrappedTimestamp
		want    string
		wantErr bool
		equal   bool
	}{
		{"Reference", WrappedTimestamp{0, Timestamp{referenceTime}}, fmt.Sprintf(`{"A":0,"Time":%s}`, referenceTimeStr), false, true},
		{"Empty", WrappedTimestamp{}, fmt.Sprintf(`{"A":0,"Time":%s}`, emptyTimeStr), false, true},
		{"Mismatch", WrappedTimestamp{}, fmt.Sprintf(`{"A":0,"Time":%s}`, referenceTimeStr), false, false},
	}
	for _, tc := range testCases {
		out, err := json.Marshal(tc.data)
		c.Assert(err, check.IsNil)
		got := string(out)
		c.Assert(got == tc.want, check.Equals, tc.equal)
	}
}

func (s *S) TestWrappedTimstamp_Unmarshal(c *check.C) {
	testCases := []struct {
		desc    string
		data    string
		want    WrappedTimestamp
		wantErr bool
		equal   bool
	}{
		{"Reference", referenceTimeStr, WrappedTimestamp{0, Timestamp{referenceTime}}, false, true},
		{"ReferenceUnix", referenceUnixTimeStr, WrappedTimestamp{0, Timestamp{referenceTime}}, false, true},
		{"Empty", emptyTimeStr, WrappedTimestamp{0, Timestamp{}}, false, true},
		{"UnixStart", `0`, WrappedTimestamp{0, Timestamp{unixOrigin}}, false, true},
		{"Mismatch", referenceTimeStr, WrappedTimestamp{0, Timestamp{}}, false, false},
		{"MismatchUnix", `0`, WrappedTimestamp{0, Timestamp{}}, false, false},
		{"Invalid", `"asdf"`, WrappedTimestamp{0, Timestamp{referenceTime}}, true, false},
	}
	for _, tc := range testCases {
		var got Timestamp
		err := json.Unmarshal([]byte(tc.data), &got)
		c.Assert(err != nil, check.Equals, tc.wantErr)
		eq := got.Time.Equal(tc.want.Time.Time)
		c.Assert(eq, check.Equals, tc.equal)
	}
}

func (s *S) TestWrappedTimstamp_MarshalReflexivity(c *check.C) {
	testCases := []struct {
		desc string
		data WrappedTimestamp
	}{
		{"Reference", WrappedTimestamp{0, Timestamp{referenceTime}}},
		{"Empty", WrappedTimestamp{0, Timestamp{}}},
	}
	for _, tc := range testCases {
		bytes, err := json.Marshal(tc.data)
		c.Assert(err, check.IsNil)
		var got WrappedTimestamp
		err = json.Unmarshal(bytes, &got)
		c.Assert(got.Time.Equal(tc.data.Time), check.Equals, true)
	}
}
