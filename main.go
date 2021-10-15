package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bflad/terraform-provider-framework/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

//go:generate terraform fmt -recursive ./examples/
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	err := tfsdk.Serve(context.Background(), provider.New, tfsdk.ServeOpts{
		Name: "bflad/framework",
	})

	if err != nil {
		fmt.Printf("error serving provider: %s", err)
		os.Exit(1)
	}
}
