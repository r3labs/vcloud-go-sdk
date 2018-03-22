# vcloud-go-sdk
A fully featured go sdk for vcloud director


## Quick start

To install:

```sh
$ go get github.com/r3labs/vcloud-go-sdk
```


## Basic Usage

#### Authentication

To create a new client:

```go
import (
    "github.com/r3labs/vcloud-go-sdk/models"
    "github.com/r3labs/vcloud-go-sdk/config"
    "github.com/r3labs/vcloud-go-sdk/client"
)

func main() {
    cfg := config.New("https://vcloud/", "27.0").WithCredentials("username", "password")
    vcloud := client.New(cfg)

    err := vcloud.Authenticate()
}
```

#### Organisations

To list all organisations:

```go
orgs, err := vcloud.Orgs.List()
```

To get an organisation by name:

```go
orgs, err := vcloud.Orgs.List()
if err != nil {
	return nil, err
}

orgID := orgs.ByName("my-org")[0].ID()

org, err :=  vcloud.Orgs.Get(orgID)
```

#### vDCs

To list all vdc's:

```go
orgs, err := vcloud.Orgs.List()
if err != nil {
	return nil, err
}

orgID := orgs.ByName("my-org")[0].ID()

org, err :=  vcloud.Orgs.Get(orgID)

vdcLinks := o.Links.ByType(models.TypesVdc)
```


To get a vdc by name:
```go
vdcLink := o.Links.ByType(models.TypesVdc).ByName("my-vdc")

vdc, err := vcloud.Vdcs.Get(vdcLink[0].ID())
```

#### Instances

To list all instances in a vdc:


```go
...

vdc, err := vcloud.Vdcs.Get(vdcLink[0].ID())

vapps := vdc.VAppRefs()
```

To create a new vApp from a vApp template:

```go
...

vdc, err := vcloud.Vdcs.Get(vdcLink[0].ID())

// retrieve network reference
nwref := vdc.NetworkRefs().Get("my-network")

// retrieve vapp template reference from catalog
catref := org.CatalogRefs().Get("my-vapp-catalog")

catalog, err := vcloud.Catalogs.Get(catref.ID())

templateref := catalog.Items().ByName("my-vapp-template")

template, err := vcloud.Catalogs.GetItem(templateref.GetID())


// create vapp request
params := &models.InstantiateVAppParams{
		Name:        "my-new-vapp",
		Description: "vapp from a template",
		AcceptEULAs: true,
		Deploy:      false,
		PowerOn:     false,
}

params.SetNetwork(nwref.Name, "bridged", nwref.Href)
params.SetSource("my-vapp-template", template.Entity.Href)

// create vapp and wait for task to complete
vapp, err := vcloud.VApps.Create(vdc.GetID(), params)

for _, task := range vapp.GetTasks() {
    err = vcloud.Tasks.Wait(&task)
}
```

To start/stop a vm:

```go
vm := vapp.Children.Vms[0]

task, err = vcloud.Vms.PowerOff(vm.GetID())
task, err = vcloud.Vms.PowerOn(vm.GetID())

err = vcloud.Tasks.Wait(task)
```

#### Networks

To list all vdc networks:

```go
networkRefs := vdc.NetworkRefs()
```

To get a network by name:

```go
networkRef := vdc.NetworkRefs().Get("my-network")

network, err := vcloud.Networks.Get(networkRef.ID())
```

#### Edge Gateways

To list all edge gateways:

```go
gatewayRecords, err := vcloud.Queries.RecordsFilter(models.QueryEdgeGateway, "vdc=="+vdc.Href, "1")
```

To get a edge gateway by name:

```go
gatewayRecords, err := vcloud.Queries.RecordsFilter(models.QueryEdgeGateway, "vdc=="+vdc.Href, "1")

id := gatewayRecords.EdgeGatewayRecords[0].ID()

gateway, err := vcloud.Gateways.Get(id)
```





## Contributing

Please read through our
[contributing guidelines](CONTRIBUTING.md).
Included are directions for opening issues, coding standards, and notes on
development.

Moreover, if your pull request contains patches or features, you must include
relevant unit tests.


## Copyright and License

Code and documentation copyright since 2017 ernest.io authors.

Code released under
[the Mozilla Public License Version 2.0](LICENSE).
