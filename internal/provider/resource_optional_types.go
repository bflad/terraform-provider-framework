package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource = &optionalTypesResource{}
)

func NewOptionalTypesResource() resource.Resource {
	return &optionalTypesResource{}
}

type optionalTypesResource struct{}

type optionalTypesResourceTypeData struct {
	ID                   string        `tfsdk:"id"`
	OptionalTypesBool    types.Bool    `tfsdk:"optional_types_bool"`
	OptionalTypesFloat64 types.Float64 `tfsdk:"optional_types_float64"`
	OptionalTypesInt64   types.Int64   `tfsdk:"optional_types_int64"`
	OptionalTypesString  types.String  `tfsdk:"optional_types_string"`
}

func (r *optionalTypesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_optional_types"
}

func (r *optionalTypesResource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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

func (r optionalTypesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	data := optionalTypesResourceTypeData{
		ID: "testing",
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r optionalTypesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Intentionally blank for testing.
}

func (r optionalTypesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Intentionally blank for testing.
}

func (r optionalTypesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Intentionally blank for testing.
}
