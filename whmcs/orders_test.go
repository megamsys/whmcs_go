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
	"net/http"

	"gopkg.in/check.v1"
)

func (s *S) TestOrder_marshall(c *check.C) {

	o := &Order{
		ClientId:      String("tom"),
		PID:           String("walker"),
		Domain:        String("simpson.com"),
		BillingCycle:  String("monthly"),
		DomainType:    String("regular"),
		RegPeriod:     String("3"),
		EppCode:       Int(1),
		NameServer1:   String("ns1.docu.sign.co"),
		PaymentMethod: String("paypal"),
		HostName:      String("awelightning.com"),
	}

	want := `{
	"clientid" :  "tom",
	"pid" : "walker",
	"Domain" : "simpson.com",
	"billingcycle" : "monthly",
	"domaintype" : "regular",
	"regperiod" : "3",
	"eppcode" : 1,
	"nameserver1" : "ns1.docu.sign.co",
	"paymentmethod" : "paypal",
	"hostname" : "awelightning.com"
	}`

	j, err := json.Marshal(o)
	c.Assert(err, check.IsNil)

	w := new(bytes.Buffer)
	err = json.Compact(w, []byte(want))
	c.Assert(err, check.IsNil)

	c.Assert(w.String(), check.Equals, string(j))

}

func (s *S) TestOrderService_Status_specifiedOrder_error(c *check.C) {
	s.mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"email": "joe@simpsons.com"}`)
	})

	user, _, err := s.client.Orders.Status(map[string]string{"clientemail": "joe@simpsons.com"})
	c.Assert(err, check.NotNil)
	c.Assert(user, check.IsNil)
}

func (s *S) TestOrderService_Status_specifiedOrder(c *check.C) {
	s.mux.HandleFunc("/includes/api.php", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"clientid": "raj@megam.io"}`)
	})

	user, _, err := s.client.Orders.Status(map[string]string{"clientid": "joe@simpsons.com",
		"username": "testing", "password": "awesome"})
	c.Assert(err, check.IsNil)
	want := &Order{ClientId: String("raj@megam.io")}
	c.Assert(user, check.DeepEquals, want)
}

func (s *S) TestOrderService_Status_invalidOrder(c *check.C) {
	_, _, err := s.client.Orders.Status(map[string]string{"clientid": "%"})
	c.Assert(err, check.NotNil)
}
