// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package postgresql

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The ``.Extension`` resource creates and manages an extension on a PostgreSQL
// server.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-postgresql/blob/master/website/docs/r/extension.html.markdown.
type Extension struct {
	pulumi.CustomResourceState

	// Which database to create the extension on. Defaults to provider database.
	Database pulumi.StringOutput `pulumi:"database"`
	// The name of the extension.
	Name pulumi.StringOutput `pulumi:"name"`
	// Sets the schema of an extension.
	Schema pulumi.StringOutput `pulumi:"schema"`
	// Sets the version number of the extension.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewExtension registers a new resource with the given unique name, arguments, and options.
func NewExtension(ctx *pulumi.Context,
	name string, args *ExtensionArgs, opts ...pulumi.ResourceOption) (*Extension, error) {
	if args == nil {
		args = &ExtensionArgs{}
	}
	var resource Extension
	err := ctx.RegisterResource("postgresql:index/extension:Extension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExtension gets an existing Extension resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExtensionState, opts ...pulumi.ResourceOption) (*Extension, error) {
	var resource Extension
	err := ctx.ReadResource("postgresql:index/extension:Extension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Extension resources.
type extensionState struct {
	// Which database to create the extension on. Defaults to provider database.
	Database *string `pulumi:"database"`
	// The name of the extension.
	Name *string `pulumi:"name"`
	// Sets the schema of an extension.
	Schema *string `pulumi:"schema"`
	// Sets the version number of the extension.
	Version *string `pulumi:"version"`
}

type ExtensionState struct {
	// Which database to create the extension on. Defaults to provider database.
	Database pulumi.StringPtrInput
	// The name of the extension.
	Name pulumi.StringPtrInput
	// Sets the schema of an extension.
	Schema pulumi.StringPtrInput
	// Sets the version number of the extension.
	Version pulumi.StringPtrInput
}

func (ExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionState)(nil)).Elem()
}

type extensionArgs struct {
	// Which database to create the extension on. Defaults to provider database.
	Database *string `pulumi:"database"`
	// The name of the extension.
	Name *string `pulumi:"name"`
	// Sets the schema of an extension.
	Schema *string `pulumi:"schema"`
	// Sets the version number of the extension.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Extension resource.
type ExtensionArgs struct {
	// Which database to create the extension on. Defaults to provider database.
	Database pulumi.StringPtrInput
	// The name of the extension.
	Name pulumi.StringPtrInput
	// Sets the schema of an extension.
	Schema pulumi.StringPtrInput
	// Sets the version number of the extension.
	Version pulumi.StringPtrInput
}

func (ExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*extensionArgs)(nil)).Elem()
}

