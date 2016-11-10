# whmcs #

whmcs is a Go client library for accessing the [WHMCS billing system][http://www.whmcs.com].

**Documentation:**

This is an enterprise addon for *MegamVertice*. Read [docs.megam.io](https://docs.megam.io) for more information.

whmcs requires Go version 1.6 or greater.

## Usage ##

```go
import "github.com/megamsys/whmcs_go/whmcs"
```

Construct a new WHMCS client, then use the various services on the client to
access different parts of the WHMCS API.  

### authentication

WHMCS supports two ways of authentication

You can either pass the variables using the Key USERNAME, PASSWORD (or) ACCCESSKEY in a map.

- username, password

`(or)`

- accesskey

### To onboard a client named  "willnorris":

```go
client := whmcs.NewClient(nil, "https://www.wheesy.com/billing")

accounts, _, err := client.Accounts.Create(map[string]string{"firstname": "willnorris"})
//Please refer the api, there are more fields to be passed

```

### To create an order for a client named  "willnorris":

```go

client := whmcs.NewClient(nil, "https://www.wheesy.com/billing")

accounts, _, err := client.Orders.Create(map[string]string{"firstname": "willnorris"})
//Please refer the api, there are more fields to be passed
```

### To add a billable item for client  "willnorris":

```go
client := whmcs.NewClient(nil, "https://www.wheesy.com/billing")

accounts, _, err := client.Billables.Create(map[string]string{"firstname": "willnorris"})
//Please refer the api, there are more fields to be passed
```

For complete usage of whmcs, see the full [package docs][].

[WHMCs API]: https://developer.github.com/v3/

The supported API are

* clients `[Add, Update, Suspend]`

* orders `[List, Get, Create]`

* billableitem `[Create]`

We will keep adding more as we go along, adding an API is very easy.

### Documentation

Refer [documentation] (http://docs.megam.io)


We are glad to help if you have questions, or request for new features..

[twitter @megamsys](http://twitter.com/megamsys) [email support@megam.io](<support@megam.io>)




# License


|                      |                                          |
|:---------------------|:-----------------------------------------|
| **Author:**          | Ranjitha (<ranjithar@megam.ion>)
| 	                   | KishorekumarNeelamegam (<nkishore@megam.io>)
| **Copyright:**       | Copyright (c) 2013-2016 Megam Systems.
| **License:**         | Apache License, Version 2.0

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
