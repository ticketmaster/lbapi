# LB Management API

[![pipeline status](https://github.com/ticketmaster/lbapi/badges/master/pipeline.svg)](https://github.com/ticketmaster/lbapi/commits/master)

The LB Management API is a web service that enables users to manage Virtual IPs ([VIP](terms.md#virtual_ip)) across multiple load balancers (e.g. AVI Networks and Citrix Netscaler). The goal is to provide users an abstraction layer that follows a single pattern for creating virtual server definitions.

## Components

- Deployment Model: **Docker**
- Language: **go1.14.6**
- Database (External): **PostgreSQL**
- Current Major release: **v3.0**

### Changes from v2.0

- Performance optimizations to Netscaler resource management.
- Migration feature added to move records from Netscaler to Avi.
- Recycle feature added to house deleted records.
- Optimizations to support UI integration.
- Automatic IP, and Load Balancer assignment.
- Integration with TM Network API (part of Infoblox library).
- Create and Update records **no longer** support multiple records "[]".
- Backup feature added to post records to GIT.
- Async response to Create, Update and Delete. Status are now tracked in status table.

## High Level Architecture

### API Routes

The LB Management API exposes two primary routes: **loadbalancer** and **virtualserver**. These routes include methods that will parse a JSON payload, validate its acontent, and forward it to a load balancer for processing. These methods include:

- Fetch (HTTP_GET) - Retrieves data directly from the database.
- Modify (HTTP_PUT) - modifies the resource on both the loadbalancer and database.
- Create (HTTP_POST) - Creates the resource on both the loadbalancer and database.
- Delete (HTTP_DELETE) - Deletes the resource on both the loadbalancer and database.

> Not included above, there is an administrative route called *source*. This route is solely meant for populating the database with current resource meta data.

```mermaid
graph LR
start((Start))
lbm(API)
avi(AVI Networks Go SDK)
ns(Netscaler Go SDK)
lb(Load Balancer<br/>/api/v1/loadbalancer)
db(Postgres)
vs(Virtual Server<br/>/api/v1/virtualserver)
return((Return Payload))

start-->vs
vs-->lbm

start-->lb
lb-->lbm

lbm-->avi
avi-->lbm
lbm-->ns
ns-->lbm
db-->lbm
lbm-->db
lbm-->return
```

### Additional Features

- Infoblox Integration
  - Create, update and delete DNS records pertaining to VIPs.
  - A standard **HOST** record named prd<PRD CODE>-IP-ADDRESS.lb.netops.tmcs is created whenever a new VIP is generated.
  - Users can provide alternate names in the virtualserver payload under the `dns[]` entity and the API will automatically create them. All records created by the API are HOST records, so reverse DNS lookups **will** resolve them.
- Product Code Authorization
  - LDAP-based login group level authorization based on product code membership. In short, VIPs under a product code can only be modified by a member of the product code security group (e.g., prd1544-LimitedAccess).
- Certificate Management (**AVI VIPs only**)
  - As long as the `service_type` is set to `https`, team members can bind a certificate to their VIP. You can bind certificates to both the vip and its pool. Note that pools with mapped certificates are solely meant for certificate authentication to the backend servers.
  - Users can create/replace certificates using the virtualserver model.
  - We support both ECDSA (Preferred) and RSA certificates.
  - When a VIP is deleted, the certificate is deleted. Each VIP has its own unique instance of a certificate (even though the same certificate may be used by multiple VIPs).

## Packages

This section is divided into two main categories: internal and external packages. Internal packages refer specifically to all the logic written specifically for the API and provide its core functionality. External packages are typically written and supported by a third-party, and extend the functionality of the API.

### Internal Packages

| Package Name  | Route                | Description                                                  | Modifies Load Balancer |
| ------------- | -------------------- | ------------------------------------------------------------ | ---------------------- |
| handler       |                      | Maps functions to routes.                                    | n/a                    |
| common        |                      | Provides functions for Fetch, Modify, Create, and Delete operations. | n/a                    |
| loadbalancer  | api/v1/loadbalancer  | Provides loadbalancer specific logic.                        | no                     |
| virtualserver | api/v1/virtualserver | Provides virtualserver specific logic.                       | **yes**                |
| pool          |                      | Provides pool/backend server logic.                          | **yes**                |
| poolgroup          |                      | Provides pool group logic. (AVI only)                          | **yes**                |
| certificate   |                      | Provides certificate logic - (AVI only)    | **yes**                   |
| healthmonitor           |                      | Provides health monitor logic.         | **yes**                 |
| persistence           |                      | Provides persistence monitor logic.            | **yes**                 |
| migrate | /api/v1/migrate/virtualserver | Provides migration logic to move between Netscaler and AVI. | **yes** |
| recycle | /api/v1/recycle | Repository for deleted records. | no |
| simple | /api/v1/simple/virtualserver | Route for returning a simplified recordsets (used by the UI). | no |
| backup | /api/v1/backup/virtualserver | Posts changed records to GIT for backup. | no |
| infoblox           |                      | Provides infoblox logic.            | no                |

#### handler

The `handler` package includes functions that are directly mapped to a `gin.RouteHandler`. To avoid defining resource specific routes, the `handler` package provides functions that accept `shared.Common` as a generic interface (e.g., definition) and tests the definition for the presence of common functions (i.e., Fetch, Modify, etc.).

To dynamically add routes to a resource, add a reference to the route to the `bootstrap` package and call it from `main.go`. Each resource will invoke it's own `common.New()` object which will map the appropriate CRUD routes to it. For example:will automatically

```
func NewLoadBalancer() *LoadBalancer {
	filter := make(map[string][]string)
	filter["loadbalancerip"] = []string{}
	////////////////////////////////////////////////////////////////////////////
	o := new(LoadBalancer)
	o.Common = common.New()
	////////////////////////////////////////////////////////////////////////////
	o.Database.Table = "loadbalancers"
	o.Database.Filter = filter
	o.Database.Validate = o.validate
	////////////////////////////////////////////////////////////////////////////
	o.ModifyLb = false
	o.Route = "loadbalancer"
	////////////////////////////////////////////////////////////////////////////
	return o
}
```

The NewLoadBalancer constructor above invokes `common.New` which adds the routes and populates the filters and database settings for CRUD operations. 
#### common

All resource specific packages (ie., loadbalancer, virtualserver, etc.) implement the `common` package. As such, the `common` package includes the "common" models and functions for creating, updating, retrieving and deleting records. This reduces the code footprint and ensures that each resource operates in the same manner.

In addition to providing the appropriate methods for handlers, common contains **sdkfork**, which routes load balancer logic to the appropriate SDK for processing. Each top-level handler method (i.e., Create, Fetch, etc), maps to a similar function within the load balancer SDK. For example, common.Create maps to either avi.Create or netscaler.Create. The decision is made according based on the client provided `load_balancer_ip` field. Each top-level handler queries the list of available load balancers and creates a map containing the **cluster_ip** of the load balancer and the manufacturer. A match process takes place and the manufacturer is forwarded to the SDK fork for decision-making.

#### loadbalancer

The loadbalancer package includes logic for retrieving and modifying load balancer records. These records include meta data such as serial number, firmware version and high availability partner.

#### virtualserver

The virtual package includes logic for retrieving and modifying virtual server (VIP) records. These records include meta data such as IP,  port and service type.

##### Create

> Sample Create payload. With the exception of the nsrservicetype field, all other fields are common between Avi and Netscaler. The nsrservicetype will eventually be deprecated and replaced with servicetype.

```go
  {
    "load_balancer_ip": "10.84.24.168",
    "data": {
      "name": "foo-netops.lb.netops.tmcs",
      "service_type": "https",
      "ip": "10.28.64.40",
      "ports": [
        {
          "port": 443,
          "l4_profile": "tcp",
          "ssl_enabled": true
        },
        {
          "port": 80,
          "l4_profile": "tcp",
          "ssl_enabled": false
        }
      ],
      "dns": [
        "foo-netops.lb.netops.tmcs"
      ],
      "enabled": true,
      "certificates": [
        {
          "certificate": "\n-----BEGIN CERTIFICATE-----\nMIIEhTCCA22gAwIBAgIRAPM8fg7ELzkqqrlgI+yBR8swDQYJKoZIhvcNAQELBQAwgcsxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMRcwFQYDVQQHEw5XZXN0IEhvbGx5d29vZDEVMBMGA1UEChMMVGlja2V0bWFzdGVyMRwwGgYDVQQLExNOZXR3b3JrIEVuZ2luZWVyaW5nMSMwIQYDVQQDExpUaWNrZXRtYXN0ZXIgVGVjaE9wcyBDQSB2NDE0MDIGCSqGSIb3DQEJARYldG1uZXR3b3JrZW5naW5lZXJpbmdAdGlja2V0bWFzdGVyLmNvbTAeFw0xOTA1MTcxNjM3MTVaFw0zMTExMjgwMjA0MTVaMIG8MQswCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEUMBIGA1UEBxMLTG9zIEFuZ2VsZXMxHjAcBgNVBAoTFVRpY2tldG1hc3RlciBTZWN1cml0eTEbMBkGA1UECxMSTmV0d29yayBPcGVyYXRpb25zMR0wGwYDVQQDExRkb25vdHVzZS5uZXRvcHMudG1jczEmMCQGCSqGSIb3DQEJARYXbmV0b3BzQHRpY2tldG1hc3Rlci5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCdNbeHEAGDF6Jw1H6W0V/V4UgZh3uwSl9UMRHDuSHp0TOV79Bja/rpyQuXO6k5Ui5uAVyvydCpPr4dngbn9d1UFMO1HR3dIeMikwiuaMv3wsTCpdJ7Gteq8ww8dbTeSvCC2kBqjDwlh6alEEU+LaJxjaxJ8kY719bc2/F4y8CbiLBYX/h4Q3XvWcu6ndK9OBmAzb2B9/QTaWIoa39IAZ8mVSuY4TYKr/4iI38QMh4aRsIzyema1q081fiRa6M8jiEzAiDP2+1Z2nnWadb3MmVro63PxyJKxT60YDbxhP0/n7adWhZn/tvW7SCjzxSMzLoObXJWb7aJTUHoqOV+KijdAgMBAAGjcTBvMB8GA1UdIwQYMBaAFGDkKrpkUeEcfvpkWjVUzEzKtKl1MA4GA1UdDwEB/wQEAwIHgDAdBgNVHQ4EFgQUeqfN0DmgV5sFaoj9g/b8AljZtFEwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMA0GCSqGSIb3DQEBCwUAA4IBAQBl60FrQoSrEOGuLMBo1LBAaxxJvJgmaIpS8at9uY/8CTEXXqXgvHc6WmWzarB33s2V+p6n31S4hVi3PdU93mgMNtQ+XzvL/mtVIkOjMLWqRjR/ofKLNrrffSHj3cu4q9QV82G6zw/IGpjzd9gexwDtPWpH4F0Aspa4qDwLHZ4jieh1ImUs5evMBPkyn353sLpuI9dsojgVxqhOdD6E/t3iUnyp5NSxe/QJ/gyaVxj+3J3lLJbsRx/e2cT3HdstXrFMz2wikKWmWmDc4l+vCIDXdOtaqSG+kIcFEOc2HwgQcciQ08AXnE1rQ84aahCT33tPA+kbLrTY/iR2jYkXYHC/\n-----END CERTIFICATE-----\n",
          "key": {
            "private_key": "\n-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCdNbeHEAGDF6Jw1H6W0V/V4UgZh3uwSl9UMRHDuSHp0TOV79Bja/rpyQuXO6k5Ui5uAVyvydCpPr4dngbn9d1UFMO1HR3dIeMikwiuaMv3wsTCpdJ7Gteq8ww8dbTeSvCC2kBqjDwlh6alEEU+LaJxjaxJ8kY719bc2/F4y8CbiLBYX/h4Q3XvWcu6ndK9OBmAzb2B9/QTaWIoa39IAZ8mVSuY4TYKr/4iI38QMh4aRsIzyema1q081fiRa6M8jiEzAiDP2+1Z2nnWadb3MmVro63PxyJKxT60YDbxhP0/n7adWhZn/tvW7SCjzxSMzLoObXJWb7aJTUHoqOV+KijdAgMBAAECggEBAIEmsC91Jsjbkce/yn98Yo8DFIhklWiusMIpzb5NSV8dTpPPABOtkeSeFbeYe91rdllJJSwFUDa6JNWQDXxisAFzTZRs5yvTuxWcVXVzAp34lEyUjeY0lxkJsvO4f25fglb9lg3yRzdNycmxJDGArAM9sFzPfIONPicTSb1DJmifCNERVLtFP/hETaTpxVbGc1mTq35th4IYM9/Iu/ObJPlwDnIcgd0E5HTvpQ/Wo9m9DPyNukuuAM1YLiY+9euueX0DEVd/n6X8AfnVijx2vPXrl9gKpnf8dg7vhcrOnm5UdwNdX0YJBVFbCM4to4bGZ88XmZqp5v4wDYb+LWH1IAECgYEA0ZH6FJ2hbvyqEtWQseQ23PbCXVICFL9fJ5ftl1SbXOvUX+UZee/4jNDLbeVk7ldCDBDtJuYR1QVzdDpEAIG+T6F4IP/gJ+PI5vgbgyb0tKv+Lkod3AZFcgnApvzboIWje2i630YUMoaKDqHjlon/iWuaLt1ofGVh+jel4FEY2V0CgYEAwAoO+3RC9AQ9bRP0qKSbjJKIHkthpUh6Qx+IgEZdsn8tTSzrUe5oULW/Gq3nL/z8KKmyRuI5QnE8puMTHO9vtV2EwfzJ1KZl+0tw7Fa9rwPFYphwwZOHkZOLm4vh1/x2icGSf3c8Y/VSIgQYch81CsyrqXtLO3iiuQNZCvUQFYECgYBdKbWwoHp5alz4znPqgPdat1+kOKawLnrQkRuP4I7IehYJI0F2EZW+k7s7eXSq96Nm1cd3OWPH/QpcKuK8DvFZWQCcOuOdGAfhlX41iYXTI3p1fYFUpH0OuwMnuNSxwXbxj5czVmX4KBMLejBAZcxxfKIoH0kps7AgmchlteeECQKBgAihKxUvn0aZ3izFpcviQb8qYoWB+6xSunPDuf2Rq+o2ftGmABkZboSZ9jF7uRTV+HrXTVSUG+CZeBFDyPsW410yC6Iv+t3ccF6/gB6Os01nDPqmQQLh30iyaaaevZJYHPeJxEyIDiWrw3oV1wdh0Z9fnSMrkDDm9eD8fobYhlWBAoGBAJTs/OSoyNEAGVwCv/OpCYqVadL/3R4d7KMe+5xkOmlOWW+5wfditMf8/C4rAWU48Ft1WjgXNN/2eqCPTrSSB8J4oU0MVsB/v/bM6VHRoV1q9agVgRDAYubD9KLSIKnN6aqS0ivjMhSvHrB0vRaURgyR2L6oxK9IHKiiaO/2AYEO\n-----END PRIVATE KEY-----\n"
          }
        }
      ],
      "load_balancing_method": "roundrobin",
      "pools": [
        {
          "default_port": 80,
          "enabled": true,
          "ssl_enabled": false,
          "bindings": [
            {
              "server": {
                "ip": "1.1.1.1"
              },
              "enabled": true,
              "graceful_disable": false
            },
            {
              "server": {
                "ip": "1.1.1.2"
              },
              "enabled": true,
              "graceful_disable": false
            }
          ],
          "health_monitors": [
            {
              "type": "https",
              "response_codes": [
                "400",
                "200"
              ]
            }
          ],
          "persistence": {
            "type": "client-ip"
          }
        }
      ],
      "product_code": 1544
    }
  }
```

The preceding request will create:

1. A VIP on Avi called **foo-netops.lb.netops.tmcs**. The creation ETL process will rename to **prd1544-foo-netops.lb.netops.tmcs-abc**. The last three chars is a random id.
2. A DNS A record called prd1544-10-74-24-38.lb.netops.tmcs. Optionally, you can populate the `dns` field (see virtualserver->model.go) with []string of fqdns. The create process will test the names, and then create them in Infoblox additional Host records with PTRs.
3. A pool on Avi called **prd1544-foo-netops.lb.netops.tmcs-abc-aas**. The last three chars is a random id.
4. The pool will have two members, and both members will be enabled.
5. The VIP listens on 80 and 443 with 443 enabled for SSL, and the pool members are set to listen on port 80 via the `default_port` option.
6. The VIP will load balance requests, using `roundrobin`. Right now, we only support `leastconnection` and `roundrobin`.
7. Each port will have a `l4_profile` association. The default is `tcp` but we support `udp` as well.
8. If an IP address is not provided under $.data.ip, then the system **will automatically** assign one based on the first server binding in the first pool object. 
9. If a load balancer ip ($.load_balancer_ip) is not provided, the system **will automatically** locate a suitable load balancer to host the record. If one is not found, the system will return an error indicating that a load balancer could not be found.

> The API supports the creation of Netscaler services, but it is **HIGHLY** discouraged. A service is literally a 1:1 relationship where each service manages a single back-end server. As a result, one must create multiple services for load balancing. As a result, the API request would look like multiple pool entities with a single binding.

Both the Netscaler and AVI support 1:Many backend definitions where a single Netscaler servicegroup or Avi Pool can manage multiple backends. For simplicity purposes, we will refer to this backend definition as a pool. You can manage the state of the entire pool by toggling the `enabled` field, and quickly add members by adding new `bindings`. If `default_port` is set at the pool level, every binding will inherit the `default_port` value; otherwise, you can set the `port` value for each binding independently. Likewise, you can set the state of each binding independently of the pool.

##### Fetch

Fetching is fairly straightforward. You send a HTTP_GET to the `virtualserver` route and it will return all records in the database. If this is a new deployment, you will have to populate the database using ImportAll (see Getting Started).

The filter GET request looks like this:

`http://<host>/api/v1/virtualserver?<key>=value`

The following keys are supported:

- load_balancer_ip
- product_code
- dns
- Any top level key under the Data struct (e.g., name, port, etc.).
- limit, INTEGER, limits the number of records returned
- offset, INTEGER, used in conjunction with limit. Used to indicate the recordset page. For example, if your recordset is 10 lines long, and your limit is 5, your query would result in 2 pages. Page 1 would be records 1-5 and Page 2 would be 6-10.

We do not currently support filtering based pool or binding - it is in the works.

The value doesn't need to be exact. In fact, adding `*` anywhere within the value string will add a wildcard in its place. For example, `name=*tix*` will search for any vip that has the word tix in itemporary ro.

You can add even more granularity by adding `&` and additional keys. For example, `name=*tix*&port=80` This will return any vip that has the word tix in it and uses port 80. Also, you can have more than one declaration of the same key (e.g., `name=string&name=string2`).

##### Modify

Modifying a record requires that you have the **latest** database record ID. I emphasize latest because these IDs are ephemeral, so they will rarely ever be the same. It is recommended that you search for the record, then record the ID. You can validate that you have the right ID by running executing an HTTP_GET to `api/v1/virtualserver/<id>`. That route will return all the facts pertaining to that VIP.  Copy that return payload and paste it into an HTTP_PUT JSON request.

> It is **critical** that you copy and paste the GET response into the PUT request in its entirety. The API compares what is being submitted to what is in the database and what it is on the loadbalancer. The user submitted request will overwrite whatever configurations are on the loadbalancer and in the database!

The Modify process is fairly comprehensive and designed to manage all of the vips dependencies. If you fail to include a binding in the PUT request, but the resource on the loadbalancer had a binding defined, the API will delete that binding on the loadbalancer. Likewise, if you forget to provide a value for `enabled`, the API will treat that field as `false` and disable the resource.

The Modify operation will also modify any changes to DNS names - to include the removal of CNames that are now longer in use.

##### Delete

The most frightening of all operations. Delete will delete the VIP and all of its dependencies. This ensures that any VIP created by the API is cleaned up after removal. This process will also delete the HOST records associated with the VIP.

#### pool
The pool package includes logic for retrieving and modifying pool records. These records include meta data such pool name and port. In addition, the pool package includes logic for modifying backend server bindings.

Unlike loadbalancer and virtualserver, there is no route dedicated to pool. The reason for this decision was to ensure that all pool operations are directly called by virtualserver.


### External Packages

| Package Name  | Import Reference                          | Description                                                  | Vendor Supported |
| ------------- | ----------------------------------------- | ------------------------------------------------------------ | ---------------- |
| AVI SDK       | github.com/avinetworks/sdk/go             | Includes logic for managing AVI Networks load balancers. This includes methods for parsing data, API end-points and models. | yes              |
| Netscaler SDK | github.com/ticketmaster/nitro-go-sdk | Includes logic for management Citrix Netscaler load balancers. This includes methods for parsing data, API end-points and models. This SDK is **NOT** vendor supported; however, it does utilize API end-points provided by the vendor. | no               |
| Infoblox      | github.com/ticketmaster/infoblox-go-sdk  | Manages the creation of A/CName/PTR records.                 | no               |

###
