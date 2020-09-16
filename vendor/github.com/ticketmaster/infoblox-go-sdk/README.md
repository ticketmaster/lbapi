# Infoblox-go-sdk
## Models
- RecordA
- RecordCname
- RecordPTR
- RecordHost

## Clients
- RecordA (Fetch, FetchByName, FetchByIP, Modify, Delete)
- RecordCname (Fetch, FetchByName, FetchByIP, Modify, Delete)
- RecordPTR (Fetch, FetchByName, FetchByIP, Modify, Delete)
- RecordHost (Fetch, FetchByName, FetchByIP, Modify, Delete)

## Business Logic (infoblox)
### Fetch(filter string)
The `filter` can be any part of an FQDN. The more specific, the faster the query will execute. Fetch will return the A or Host record of the matched name, and its referenced PTR and CName/Aliases. Each object will also include all associated IB references.

### FetchByIP(ip string)
Performs the same operation as Fetch, but does so by IP address.

### CreateA(rec infoblox.Record)
Creates an A record, and all referenced objects (i.e., PTR and Cname records).

### ModifyA(rec infoblox.Record)
Updates an A record, and all referenced objects (i.e., PTR and Cname records). This oepration will delete PTR records if the IP address has changed, and create/delete Cname records.

### DeleteA(rec infoblox.Record)
Deletes an A record and all of its referenced objects.

## Usage
### Clients
The clients are designed to provide the most amount of flexibility to developers. Developers can instantiate a client, attach a client session to the client and then call any of the client methods. These are raw operations and rely heavily on the business logic of the appliance.

### Business Logic (infoblox)
The infoblox package includes the `Infoblox` helper struct that includes all defined clients and methods to provide some automation. Develepers can call the `infoblox.Set(client.Host)` method to instantiate the helper object. A `client.Host` parameter must be passed into `Set`; otherwise, the helper will not be able to communicate with Infoblox. See below or reference the inflox_test.go file for more details.

```
host := new(client.Host)
host.Name = InfobloxServer
host.UserName = Username
host.Password = Password
o := infoblox.Set(host)
defer o.Unset()
if err != nil {
	log.Print(err)
	t.Fail()
}

r,_: = o.Fetch("mydns")

log.Printf("%+v", r)
```

Once the helper object is created, the developer can call any of the business logic functions above.

## Credits
- Author: Carlos Villanueva, carlos.villanueva@ticketmaster.com