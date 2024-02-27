package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSetUnknownResource_basic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `resource "framework_set_unknown" "test" {
					set_nested_attribute = [
						{
							id = "id-123"
							permission = "permission1"
						},
						{
							id = "id-456"
							permission = "permission1"
						},
					]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_set_unknown.test", "id", "testing"),
					resource.TestCheckResourceAttr("framework_set_unknown.test", "set_nested_attribute.#", "2"),
					resource.TestCheckTypeSetElemNestedAttrs("framework_set_unknown.test", "set_nested_attribute.*", map[string]string{
						"id":         "id-123",
						"name":       "name-123",
						"permission": "permission1",
					}),
					resource.TestCheckTypeSetElemNestedAttrs("framework_set_unknown.test", "set_nested_attribute.*", map[string]string{
						"id":         "id-456",
						"name":       "name-456",
						"permission": "permission1",
					}),
				),
			},
			{
				Config: `resource "framework_set_unknown" "test" {
					set_nested_attribute = [
						{
							id = "id-456"
							permission = "permission2"
						},
					]
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("framework_set_unknown.test", "id", "testing"),
					resource.TestCheckResourceAttr("framework_set_unknown.test", "set_nested_attribute.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs("framework_set_unknown.test", "set_nested_attribute.*", map[string]string{
						"id":         "id-456",
						"name":       "name-456",
						"permission": "permission2",
					}),
				),
			},
		},
	})
}
