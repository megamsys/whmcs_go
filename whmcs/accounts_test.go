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
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"

	"gopkg.in/check.v1"
)

func (s *S) TestAccount_marshall(c *check.C) {
	u := &Account{
		FirstName:   String("tom"),
		LastName:    String("walker"),
		Email:       String("joe@simpson.com"),
		Address1:    String("1660 parklawm ave"),
		City:        String("edina"),
		State:       String("mn"),
		PostCode:    String("49289"),
		Country:     String("USA"),
		PhoneNumber: String("612-908-789"),
		Password:    String("bill4me"),
		Status:      String("Active"),
	}

	want := `{
	"firstname" :  "tom",
	"lastname" : "walker",
	"email" : "joe@simpson.com",
	"address1" : "1660 parklawm ave",
	"city" : "edina",
	"state" : "mn",
	"postcode" : "49289",
	"country" : "USA",
	"phonenumber" : "612-908-789",
	"password2" : "bill4me",
	"status": "Active"
	}`

	j, err := json.Marshal(u)
	c.Assert(err, check.IsNil)

	w := new(bytes.Buffer)
	err = json.Compact(w, []byte(want))
	c.Assert(err, check.IsNil)

	c.Assert(w.String(), check.Equals, string(j))

}

func (s *S) TestAccountsService_Get_specifiedUser_error(c *check.C) {
	s.mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"email": "joe@simpsons.com"}`)
	})

	user, _, err := s.client.Accounts.Get(map[string]string{"clientemail": "joe@simpsons.com"})
	c.Assert(err, check.NotNil)
	c.Assert(user, check.IsNil)
}

func (s *S) TestAccountsService_Get_specifiedUser(c *check.C) {
	s.mux.HandleFunc("/includes/api.php", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"email": "raj@megam.io"}`)
	})

	user, _, err := s.client.Accounts.Get(map[string]string{"clientemail": "joe@simpsons.com",
		"username": "testing", "password": "awesome"})
	c.Assert(err, check.IsNil)
	want := &Account{Email: String("raj@megam.io")}
	c.Assert(user, check.DeepEquals, want)
}

func (s *S) TestAccountsService_Get_invalidUser(c *check.C) {
	_, _, err := s.client.Accounts.Get(map[string]string{"clientemail": "%"})
	c.Assert(err, check.NotNil)
}

func (s *S) TestAccount(c *check.C) {
	addr := strings.Join([]string{"103.56.92.20", strconv.Itoa(80)}, ":")
	_, err := net.Dial("tcp", addr)
	c.Assert(err, check.IsNil)
	//	if err == nil {
	//		c.Skip("WHMCS isn't running. You can't rest it live.")
	//	}
	//	defer conn.Close()

	client := NewClient(nil, "http://103.56.92.20/whmcs/")
	a := map[string]string{
		"username":    "megamsys",
		"password":    GetMD5Hash("megam"),
		"accesskey":   "team4megam",
		"firstname":   "Jonathan",
		"lastname":    "Philipos",
		"email":       "jp@det.io",
		"address1":    "Panara",
		"city":        "Western sydney",
		"state":       "Sydney",
		"postcode":    "00001",
		"country":     "AU",
		"phonenumber": "981999000",
		"password2":   "temp4det",
	}
	_, _, err = client.Accounts.Create(a)
	c.Assert(err, check.IsNil)
}

func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}
