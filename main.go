package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/bflad/terraform-provider-framework/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

//go:generate terraform fmt -recursive ./examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	err := providerserver.Serve(context.Background(), provider.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/bflad/framework",
		Debug:   debug,
	})

	if err != nil {
		fmt.Printf("error serving provider: %s", err)
		os.Exit(1)
	}
}
