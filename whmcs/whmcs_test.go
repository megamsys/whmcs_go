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
/*
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"gopkg.in/check.v1"
)

type values map[string]string

func testFormValues(t *testing.T, r *http.Request, values values) {
	want := url.Values{}
	for k, v := range values {
		want.Add(k, v)
	}

	r.ParseForm()
	if got := r.Form; !reflect.DeepEqual(got, want) {
		t.Errorf("Request parameters: %v, want %v", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %s, want %s", header, got, want)
	}
}

func testURLParseError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func testBody(t *testing.T, r *http.Request, want string) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Errorf("Error reading request body: %v", err)
	}
	if got := string(b); got != want {
		t.Errorf("request Body is %s, want %s", got, want)
	}
}


func (s *S) TestNewClient(c *check.C) {
	cl := NewClient(nil, s.url)
	got := cl.BaseURL.String()
	c.Assert(got, check.Equals, "https://www.megam.io/billing/")
	got = cl.ApiEndSux
	c.Assert(got, check.Equals, "includes/api.php")
}

func (s *S) TestNewWRequest_emptydata(c *check.C) {
	cl := NewClient(nil, s.url)
	req, _ := cl.NewWRequest(map[string]string{}, "testing")
	got := req.url.String()
	c.Assert(got, check.Equals, s.url+cl.ApiEndSux)
}

func (s *S) TestNewWRequest_invaliddata(c *check.C) {
	cl := NewClient(nil, s.url)
	req, err := cl.NewWRequest(map[string]string{"usernama": "testings"},"testing")
	c.Assert(err, check.IsNil)
	c.Assert(req, check.NotNil)
}

func (s *S) TestNewRequest_noAction(c *check.C) {
	cl := NewClient(nil, s.url)
	req, err := cl.NewWRequest(map[string]string{"username": "testing", "password": "testpass"}, "testing")
	c.Assert(err, check.IsNil)
	c.Assert(req, check.NotNil)
}

func (s *S) TestNewRequest_Apikey(c *check.C) {
	cl := NewClient(nil, s.url)
	req, err := cl.NewWRequest(map[string]string{"api_key": "testing"}, "testing")
	c.Assert(err, check.IsNil)
	c.Assert(req, check.NotNil)
}

func (s *S) TestDo(c *check.C) {
	type foo struct {
		A string
	}

	s.mux.HandleFunc("/includes/api.php", func(w http.ResponseWriter, r *http.Request) {
		c.Assert(r.Method, check.Equals, "POST")
		fmt.Fprint(w, `{"A":"a"}`)
	})
	req, _ := s.client.NewWRequest(map[string]string{"username": "first", "password": "pass", "action": "GetClients"}, "testing")
	body := new(foo)
	s.client.Do(*req, body)
	want := &foo{A: "a"}
	c.Assert(body, check.DeepEquals, want)
}

func (s *S) TestDo_httpError(c *check.C) {
	type foo struct {
		A string
	}
	http.NewServeMux().HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Request", 400)
	})

	req, _ := s.client.NewWRequest(map[string]string{"username": "first", "password": "pass", "action": "GetClients"}, "testing")
	body := new(foo)
	_, err := s.client.Do(*req, body)
	c.Assert(err, check.NotNil)
}

func (s *S) TestSanitizeURL(c *check.C) {
	tests := []struct {
		in, want string
	}{
		{"/?a=b", "/?a=b"},
		{"/?a=b&accesskey=secret", "/?a=b&accesskey=REDACTED"},
		{"/?a=b&client_id=id&accesskey=secret", "/?a=b&accesskey=REDACTED&client_id=id"},
	}

	for _, tt := range tests {
		inURL, _ := url.Parse(tt.in)
		want, _ := url.Parse(tt.want)

		c.Assert(sanitizeURL(inURL), check.DeepEquals, want)
	}
}

func (s *S) TestCheckResponse(c *check.C) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusBadRequest,
		Body: ioutil.NopCloser(strings.NewReader(`{"message":"m",
			"errors": [{"resource": "r", "field": "f", "code": "c"}]}`)),
	}
	err := CheckResponse(res).(*ErrorResponse)
	c.Assert(err, check.NotNil)

	want := &ErrorResponse{
		Response: res,
		Message:  "m",
		Errors:   []Error{{Resource: "r", Field: "f", Code: "c"}},
	}
	c.Assert(err, check.DeepEquals, want)
}

// ensure that we properly handle API errors that do not contain a response body
func (s *S) TestCheckResponse_noBody(c *check.C) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusBadRequest,
		Body:       ioutil.NopCloser(strings.NewReader("")),
	}
	err := CheckResponse(res).(*ErrorResponse)

	c.Assert(err, check.NotNil)

	want := &ErrorResponse{
		Response: res,
	}

	c.Assert(err, check.DeepEquals, want)
}

func (s *S) TestParseBooleanResponse_true(c *check.C) {
	result, err := parseBoolResponse(nil)
	c.Assert(err, check.IsNil)
	c.Assert(result, check.Equals, true)
}

func (s *S) TestParseBooleanResponse_false(c *check.C) {
	v := &ErrorResponse{Response: &http.Response{StatusCode: http.StatusNotFound}}
	result, err := parseBoolResponse(v)
	c.Assert(err, check.IsNil)
	c.Assert(result, check.Equals, false)
}

func (s *S) TestParseBooleanResponse_error(c *check.C) {
	v := &ErrorResponse{Response: &http.Response{StatusCode: http.StatusBadRequest}}
	result, err := parseBoolResponse(v)
	c.Assert(err, check.NotNil)
	c.Assert(result, check.Equals, false)
}

func (s *S) TestErrorResponse_Error(c *check.C) {
	res := &http.Response{Request: &http.Request{}}
	err := ErrorResponse{Message: "m", Response: res}
	c.Assert(err, check.NotNil)
	c.Assert(len(err.Error()) > 0, check.Equals, true)
}
*/
