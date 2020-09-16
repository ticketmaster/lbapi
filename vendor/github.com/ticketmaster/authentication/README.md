# Authentication

## TL&DR
Authentication/authorization library for Go based APIs. Supports Gin and Revel.

# Authentication/Authorization

## Overview

 This library enables users to authenticate either locally to a GO-based API or to an external authentication provider such as LDAP. Once authenticated, the package evaluates the user's group membership (e.g., role) and matches the aggregated roles to system-defined authorization rules (e.g., authentication.yam). Authorization rules determine what level of access the user has to an API end-point.

## Dependencies

The authentication package requires that the system implements either [Gin](https://github.com/gin-gonic/gin) or [Revel](https://revel.github.io/) for route handling. The chief differences between the two frameworks is that Gin is strictly an API framework whereas Revel is a true web framework with front-end support.

For LDAP support, the API implementing this middleware must import an authentication.yaml file with defined rules. See sample_authentication.yaml.

## How It Works

Implementation of the authentication package is quite simple. Note that the following process assumes that the routing framework is Gin. 

1. In your main.go file, create a new router object `router := gin.Default()`
2. To enable authentication, simply import the authentication package and call the `UseAuthentication` method (see below).

```go
package main
import (
	filter "github.com/ticketmaster/authentication/gin"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)
func main() {
    router := gin.Default() 
/* Default automatically creates routes for logging, etc. See gin documentation for additional options */
    err := filter.UseAuthentication(router, filter.NewAuthenticationOptions())
    ...
}
```

The `NewAuthenticationOptions` method loads all the definitions stored in the authentication.yaml file into memory for system consumption.

### Authentication/Authorization Definitions

There are two main sections to the authentication.yaml file. One section is for authentication and the other is for authorization. In the authentication section, one must instruct the system to either authenticate locally (e.g., memory) or remotely (e.g. LDAP).

#### Local Authentication

```go
authenticationClient:
  - provider: memory // memory or ldap.
    origin: testOrigin // grouping for rule association.
    users:
      - username: test // local user.
        password: testpass // plaintext password - not secure!
        name: My Name // name of user.
        email: test@test.com // email of user.
        roles: // either role or roles.
          - testRole // role/group user is a member of.
          - testRole2 // role/group user is a member of.
...
```

#### Remote Authentication

```go
authenticationClient:
  - provider: ldap // memory or ldap.
    endpoint: foo.bar.local // authentication provider.
    baseDN: DC=foo,DC=bar // ldap DN to search for user accounts.
    port: 636 // ldap port.
    useTLS: true // TLS for secure communication.
    shortDomain: foobar // instructs the package to assume techops domain.
    tlsServerName: foo.bar.local // SNI for endpoint.
...
```

#### Authorization

Next, we define the authorization rules. Although the package supports the ability to <u>explicitly allow</u> and <u>explicitly deny</u> access to routing end-points, it is best to focus on one or the other. Otherwise, we run the risk of accidentally granting access to a sensitive resource because of rule precedence.

To ensure that the API is secure, we <u>implicitly deny</u> access to all routes; thus, cutting out all access to the API right from the get go. Then we create rules that explicitly grant access to routes and HTTP methods based on role membership.  

```yaml
#^^^ authentication rules go above
authorization:
  default: deny # allow or deny. in this case, implicitly deny access to all routes
  rules:
    - ruleType: route # apply rule to route
      method: GET # HTTP method to apply the rule to
      route: 
        - /mypath # route end-point
      authorize: allow # allow or deny - since we implicitly deny, avoid adding deny rules
      role: "NotRoot" # Either a defined role for local auth or an AD Security Group
      origin: foo # grouping for rule association - must match field value in authentication section

```

A simple note on authorization rules. You can define as many rules as you wish, but be very careful with how you define the routes for each rule. It is best to group routes based on method(s), and limit the number of unique rules. For example:

- Rule 1: Allow all authenticated users in NotRoot, read access (GET)
- Rule 2: Allow all members of Operator, read and modify access (GET/PUT/PATCH)
- Rule 3: Allow all members of Administrator, read,  modify and delete access (GET/PUT/PATCH/DELETE)

As such, the methods would look like:

```yaml
# Rule 1
...
      method:
        - GET
...
# Rule 2 (Operator roles)
...
      method:
        - GET
	- PUT
	- PATCH
...
# Rule 3 (Administrator roles)
...
	  method:
        - GET
	- PUT
	- PATCH
	- DELETE
...
```

#### JSON Web Token

The last component of the authentication.yaml file is the configuration options for the JSON Web Token (JWT). For this, all we need to do is include the certificate pair for the API (used for signing/decrypting tokens) and the expiration value for the token.

```yaml
# ^^^ authentication/authorization rules
privateKey: "private.key"
publicKey: "sign.crt"
jwtExpiration: "1h"
```

To request a token, all users have to do is submit a JSON payload to `/login`, using POST. For example:

```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"username":"xyz","password":"xyz"}' \
  https://myapi/login
```

The aforementioned request will return a JSON object with a token. 

The following expands on the request and includes using the token to make a GET request to the virtualserver end-point.

```bash
# Set Token
TOKEN=$(curl -s -X POST -H 'Accept: application/json' -H 'Content-Type: application/json' --data '{"username":"{username}","password":"{password}","rememberMe":false}' https://myapi/login | jq -r '.id_token')

curl -H 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" https://myapi/mypath
```



## Credits
- Author: Mike Walker
- Contributors: Carlos Villanueva
