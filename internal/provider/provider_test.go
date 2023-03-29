package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"framework": providerserver.NewProtocol6WithError(New()),
}

func testMustReadFile(t *testing.T, path string) string {
	t.Helper()

	contents, err := os.ReadFile(path)

	if err != nil {
		t.Fatalf("unexpected error reading %s: %s", path, err)
	}

	return string(contents)
}
