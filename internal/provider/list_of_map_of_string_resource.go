package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &ListOfMapOfStringResource{}
	_ resource.ResourceWithImportState = &ListOfMapOfStringResource{}
)

func NewListOfMapOfStringResource() resource.Resource {
	return &ListOfMapOfStringResource{}
}

type ListOfMapOfStringResource struct{}

func (r *ListOfMapOfStringResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_list_of_map_of_string"
}

func (r *ListOfMapOfStringResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"list_of_map_of_string": schema.ListAttribute{
				Required: true,
				ElementType: types.MapType{
					ElemType: types.StringType,
				},
			},
		},
	}
}

func (r ListOfMapOfStringResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ListOfMapOfStringResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Example testing logic based on acceptance testing configuration
	if data.ListOfMapOfString.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("list_of_map_of_string"),
			"Unexpected Value",
			"Expected ListOfMapOfString to be known, got null",
		)

		return
	}

	var listOfTerraformMap []types.Map

	resp.Diagnostics.Append(data.ListOfMapOfString.ElementsAs(ctx, &listOfTerraformMap, false)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if len(listOfTerraformMap) != 1 {
		resp.Diagnostics.AddAttributeError(
			path.Root("list_of_map_of_string"),
			"Unexpected Value",
			fmt.Sprintf("Expected ListOfMapOfString to have 1 list element, got %d", len(listOfTerraformMap)),
		)

		return
	}

	for _, terraformMap := range listOfTerraformMap {
		if terraformMap.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("list_of_map_of_string"),
				"Unexpected Value",
				"Expected map under list to be known, got null",
			)

			return
		}

		var mapOfTerraformString map[string]types.String

		resp.Diagnostics.Append(terraformMap.ElementsAs(ctx, &mapOfTerraformString, false)...)

		if resp.Diagnostics.HasError() {
			return
		}

		if len(mapOfTerraformString) != 1 {
			resp.Diagnostics.AddAttributeError(
				path.Root("list_of_map_of_string"),
				"Unexpected Value",
				fmt.Sprintf("Expected ListOfMapOfString list element to have 1 map element, got %d", len(mapOfTerraformString)),
			)

			return
		}

		for key, terraformString := range mapOfTerraformString {
			if key != "testkey" {
				resp.Diagnostics.AddAttributeError(
					path.Root("list_of_map_of_string"),
					"Unexpected Value",
					fmt.Sprintf("Expected map key to be \"testkey\", got %q", key),
				)

				return
			}

			if terraformString.IsNull() {
				resp.Diagnostics.AddAttributeError(
					path.Root("list_of_map_of_string"),
					"Unexpected Value",
					"Expected string under map under list to be known, got null",
				)

				return
			}

			if terraformString.ValueString() != "testvalue" {
				resp.Diagnostics.AddAttributeError(
					path.Root("list_of_map_of_string"),
					"Unexpected Value",
					fmt.Sprintf("Expected map value to be \"testvalue\", got %q", terraformString.ValueString()),
				)

				return
			}
		}
	}

	data.ID = types.StringValue("testing")

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r ListOfMapOfStringResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ListOfMapOfStringResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r ListOfMapOfStringResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ListOfMapOfStringResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r ListOfMapOfStringResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ListOfMapOfStringResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
}

func (r ListOfMapOfStringResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

type ListOfMapOfStringResourceModel struct {
	ID                types.String `tfsdk:"id"`
	ListOfMapOfString types.List   `tfsdk:"list_of_map_of_string"`
}
