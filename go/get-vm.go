package main

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/property"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

const (
	URL      = "https://127.0.0.1/sdk"
	UserName = "user"
	Password = "pass"
	Insecure = true
)

func exit(err error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	os.Exit(1)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	u, err := url.Parse(URL)
	if err != nil {
		exit(err)
	}

	u.User = url.UserPassword(UserName, Password)

	c, err := govmomi.NewClient(ctx, u, Insecure)
	if err != nil {
		exit(err)
	}
	f := find.NewFinder(c.Client, true)

	dc, err := f.DefaultDatacenter(ctx)
	if err != nil {
		exit(err)
	}

	f.SetDatacenter(dc)

	vms, err := f.VirtualMachineList(ctx, "*")
	if err != nil {
		exit(err)
	}

	var refs []types.ManagedObjectReference
	for _, vm := range vms {
		refs = append(refs, vm.Reference())
	}

	pc := property.DefaultCollector(c.Client)
	var vmt []mo.VirtualMachine
	err = pc.Retrieve(ctx, refs, []string{"name"}, &vmt)
	if err != nil {
		exit(err)
	}

	for _, vm := range vmt {
		fmt.Println(vm.Name)
	}
}
