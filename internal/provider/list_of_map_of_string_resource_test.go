package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestListOfMapOfStringResource_basic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_list_of_map_of_string" "test" {
					list_of_map_of_string = [
						{
							testkey = "testvalue"
						},
					]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_list_of_map_of_string.test", "id", "testing"),
					resource.TestCheckResourceAttr("framework_list_of_map_of_string.test", "list_of_map_of_string.#", "1"),
					resource.TestCheckResourceAttr("framework_list_of_map_of_string.test", "list_of_map_of_string.0.%", "1"),
					resource.TestCheckResourceAttr("framework_list_of_map_of_string.test", "list_of_map_of_string.0.testkey", "testvalue"),
				),
			},
		},
	})
}
