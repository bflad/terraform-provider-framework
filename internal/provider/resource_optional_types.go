package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type optionalTypesResourceType struct{}

func (t optionalTypesResourceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:     types.StringType,
				Computed: true,
			},
			"optional_types_bool": {
				Type:     types.BoolType,
				Optional: true,
			},
			"optional_types_float64": {
				Type:     types.Float64Type,
				Optional: true,
			},
			"optional_types_int64": {
				Type:     types.Int64Type,
				Optional: true,
			},
			"optional_types_string": {
				Type:     types.StringType,
				Optional: true,
			},
		},
	}, nil
}

func (t optionalTypesResourceType) NewResource(ctx context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return optionalTypesResource{}, nil
}

type optionalTypesResourceTypeData struct {
	ID                   string        `tfsdk:"id"`
	OptionalTypesBool    types.Bool    `tfsdk:"optional_types_bool"`
	OptionalTypesFloat64 types.Float64 `tfsdk:"optional_types_float64"`
	OptionalTypesInt64   types.Int64   `tfsdk:"optional_types_int64"`
	OptionalTypesString  types.String  `tfsdk:"optional_types_string"`
}

type optionalTypesResource struct{}

func (r optionalTypesResource) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	data := optionalTypesResourceTypeData{
		ID: "testing",
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r optionalTypesResource) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	// Intentionally blank for testing.
}

func (r optionalTypesResource) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	// Intentionally blank for testing.
}

func (r optionalTypesResource) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	// Intentionally blank for testing.
}

func (r optionalTypesResource) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStateNotImplemented(ctx, "intentionally not implementated", resp)
}
