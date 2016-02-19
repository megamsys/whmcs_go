package whmcs

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

type S struct {
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the GitHub client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server

	url     string
}

var _ = check.Suite(&S{})


func (s *S) SetUpTest(c *check.C) {
	// test server
	s.mux = http.NewServeMux()

	s.server = httptest.NewServer(s.mux)

	// whmcs client configured to use test server
	s.client = NewClient(nil, s.server.URL)

	//the defaultBaseURL
	s.url = "https://www.megam.io/billing/"
}

func (s *S) TearDownTest(c *check.C) {
	s.server.Close()
}

/*func (s *S) SetUpTest(c *check.C) {
	server, err := otesting.NewServer("127.0.0.1:5555")
	c.Assert(err, check.IsNil)
	s.server = server
	s.p, err = newFakeOneProvisioner(s.server.URL())
	c.Assert(err, check.IsNil)
}

func (s *S) TestTearDownTest(c *check.C) {
	s.server.Stop()
}*/
