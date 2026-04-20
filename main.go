// Package main provides the entry point for the Infra Terraform provider.
package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/infrahq/terraform-provider-infra/internal/provider"
)

// version is set during build via ldflags.
var version = "dev"

func main() {
	var debugMode bool

	// Default debug to false; pass -debug flag explicitly when using delve
	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		// NOTE: Using my personal fork's registry address for local testing.
		// Switch back to registry.terraform.io/infrahq/infra for upstream builds.
		Address: "registry.terraform.io/infrahq/infra",
		Debug:   debugMode,
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)
	if err != nil {
		log.Fatal(err)
	}
}
