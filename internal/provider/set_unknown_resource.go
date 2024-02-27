package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource                = &SetUnknownResource{}
	_ resource.ResourceWithImportState = &SetUnknownResource{}
)

func NewSetUnknownResource() resource.Resource {
	return &SetUnknownResource{}
}

type SetUnknownResource struct{}

func (r *SetUnknownResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_set_unknown"
}

func (r *SetUnknownResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"set_nested_attribute": schema.SetNestedAttribute{
				Computed: true,
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"name": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"permission": schema.StringAttribute{
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r SetUnknownResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SetUnknownResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.ID = types.StringValue("testing")

	// Pretend the API filled in Computed attributes
	if !data.SetNestedAttribute.IsNull() {
		var setNestedAttributeObjects []types.Object

		resp.Diagnostics.Append(data.SetNestedAttribute.ElementsAs(ctx, &setNestedAttributeObjects, false)...)

		if resp.Diagnostics.HasError() {
			return
		}

		for idx, setNestedAttributeObject := range setNestedAttributeObjects {
			if setNestedAttributeObject.IsNull() {
				continue
			}

			var object SetNestedAttributeObjectModel

			resp.Diagnostics.Append(setNestedAttributeObject.As(ctx, &object, basetypes.ObjectAsOptions{})...)

			switch object.Id.ValueString() {
			case "id-123":
				object.Name = types.StringValue("name-123")
			case "id-456":
				object.Name = types.StringValue("name-456")
			}

			updatedSetNestedAttributeObject, diags := types.ObjectValueFrom(ctx, setNestedAttributeObject.AttributeTypes(ctx), object)

			resp.Diagnostics.Append(diags...)

			if resp.Diagnostics.HasError() {
				continue
			}

			setNestedAttributeObjects[idx] = updatedSetNestedAttributeObject
		}

		updatedSetNestedAttribute, diags := types.SetValueFrom(ctx, data.SetNestedAttribute.ElementType(ctx), setNestedAttributeObjects)

		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}

		data.SetNestedAttribute = updatedSetNestedAttribute
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r SetUnknownResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SetUnknownResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r SetUnknownResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SetUnknownResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Pretend the API filled in Computed attributes
	if !data.SetNestedAttribute.IsNull() {
		var setNestedAttributeObjects []types.Object

		resp.Diagnostics.Append(data.SetNestedAttribute.ElementsAs(ctx, &setNestedAttributeObjects, false)...)

		if resp.Diagnostics.HasError() {
			return
		}

		for idx, setNestedAttributeObject := range setNestedAttributeObjects {
			if setNestedAttributeObject.IsNull() {
				continue
			}

			var object SetNestedAttributeObjectModel

			resp.Diagnostics.Append(setNestedAttributeObject.As(ctx, &object, basetypes.ObjectAsOptions{})...)

			switch object.Id.ValueString() {
			case "id-123":
				object.Name = types.StringValue("name-123")
			case "id-456":
				object.Name = types.StringValue("name-456")
			}

			updatedSetNestedAttributeObject, diags := types.ObjectValueFrom(ctx, setNestedAttributeObject.AttributeTypes(ctx), object)

			resp.Diagnostics.Append(diags...)

			if resp.Diagnostics.HasError() {
				continue
			}

			setNestedAttributeObjects[idx] = updatedSetNestedAttributeObject
		}

		updatedSetNestedAttribute, diags := types.SetValueFrom(ctx, data.SetNestedAttribute.ElementType(ctx), setNestedAttributeObjects)

		resp.Diagnostics.Append(diags...)

		if resp.Diagnostics.HasError() {
			return
		}

		data.SetNestedAttribute = updatedSetNestedAttribute
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, data)...)
}

func (r SetUnknownResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SetUnknownResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
}

func (r SetUnknownResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

type SetUnknownResourceModel struct {
	ID                 types.String `tfsdk:"id"`
	SetNestedAttribute types.Set    `tfsdk:"set_nested_attribute"`
}

type SetNestedAttributeObjectModel struct {
	Id         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	Permission types.String `tfsdk:"permission"`
}
