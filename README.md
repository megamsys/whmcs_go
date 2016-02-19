# whmcs #

whmcs is a Go client library for accessing the [WHMCS billing system][].

**Documentation:**

whmcs requires Go version 1.6 or greater.

## Usage ##

```go
import "github.com/megamsys/whmcs_go/whmcs"
```

Construct a new WHMCS client, then use the various services on the client to
access different parts of the WHMCS API.  

### authentication

WHMCS supports two ways of authentication

Pass the variables using the Key USERNAME, PASSWORD (or) ACCCESSKEY in a map.

- username, password

`(or)`

- accesskey

### To onboard a client named  "willnorris":

```go
client := whmcs.NewClient(nil, "https://www.wheesy.com/billing")

accounts, _, err := client.Accounts.Create(map[string]string{"firstname": "willnorris"}) //Please refer the api, there are more fiels to be passed

```

### To create an order for a client named  "willnorris":

```go
client := whmcs.NewClient(nil, "https://www.wheesy.com/billing")

accounts, _, err := client.Orders.Create(map[string]string{"firstname": "willnorris"})
```

### To add a billable item for client  "willnorris":

```go
client := whmcs.NewClient(nil, "https://www.wheesy.com/billing")

accounts, _, err := client.Billables.Create(map[string]string{"firstname": "willnorris"})


For complete usage of whmcs, see the full [package docs][].

[WHMCs API]: https://developer.github.com/v3/

The supported API are

* clients
* orders
* billableitem

We will keep adding more as we go along

## License ##

This library is distributed under the BSD-style license found in the [LICENSE](./LICENSE)
file.
