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
	"bytes"
	"encoding/json"
	"fmt"
	"net"
  "strings"
	"strconv"
	"gopkg.in/check.v1"
)

func (s *S) TestBillableItem_marshall(c *check.C) {

	o := &BillableItem{
		ClientId:      String("tom"),
		Description:   String("walker.wheesy.com billed"),
		Hours:         String("10"),
		Amount:        String("100"),
		InvoiceAction: String("noinvoice"),
	}

	want := `{
	"clientid" :  "tom",
	"description" : "walker.wheesy.com billed",
	"hours" : "10",
	"amount" : "100",
	"invoiceaction" : "noinvoice"
	}`

	j, err := json.Marshal(o)
	c.Assert(err, check.IsNil)

	w := new(bytes.Buffer)
	err = json.Compact(w, []byte(want))
	c.Assert(err, check.IsNil)

	c.Assert(w.String(), check.Equals, string(j))

}

func (s *S) TestBillableService_Create_specifiedOrder_error(c *check.C) {
	s.mux.HandleFunc("/bi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"email": "joe@simpsons.com"}`)
	})

	user, _, err := s.client.Billables.Create(map[string]string{"clientemail": "joe@simpsons.com"})
	c.Assert(err, check.NotNil)
	c.Assert(user, check.IsNil)
}

func (s *S) TestBillableService_Create_specifiedOrder(c *check.C) {
	s.mux.HandleFunc("/includes/api.php", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"clientid": "raj@megam.io"}`)
	})

	user, _, err := s.client.Billables.Create(map[string]string{"clientid": "joe@simpsons.com",
		"username": "testing", "password": "awesome"})
	c.Assert(err, check.IsNil)
	want := &BillableItem{ClientId: String("raj@megam.io")}
	c.Assert(user, check.DeepEquals, want)
}

func (s *S) TestBillableService_Create_invalidOrder(c *check.C) {
	_, _, err := s.client.Billables.Create(map[string]string{"clientid": "%"})
	c.Assert(err, check.NotNil)
}

func (s *S) TestBillableService_Create(c *check.C) {
	addr := strings.Join([]string{"192.168.0.133", strconv.Itoa(80)}, ":")
	_, err := net.Dial("tcp", addr)
	c.Assert(err, check.IsNil)
	//	if err == nil {
	//		c.Skip("WHMCS isn't running. You can't rest it live.")
	//	}
	//	defer conn.Close()
	client := NewClient(nil, "http://192.168.0.133/whmcs/")
	a := map[string]string{
		"username":      "Megam",
		"password":      GetMD5Hash("team4megam"),
		"clientid":      "67",
		"description":   "testing billableitems",
		"hours":         "1",
		"amount":        "0.3",
		"invoiceaction": "nextcron",
	}
	fmt.Println(client)
	_, _, err = client.Billables.Create(a)
		c.Assert(err, check.IsNil)

}
