// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Deprecated: postgresql.DefaultPrivileg has been deprecated in favor of postgresql.DefaultPrivileges
type DefaultPrivileg struct {
	pulumi.CustomResourceState

	// The database to grant default privileges for this role
	Database pulumi.StringOutput `pulumi:"database"`
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type)
	ObjectType pulumi.StringOutput `pulumi:"objectType"`
	// Target role for which to alter default privileges.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// The list of privileges to apply as default privileges
	Privileges pulumi.StringArrayOutput `pulumi:"privileges"`
	// The name of the role to which grant default privileges on
	Role pulumi.StringOutput `pulumi:"role"`
	// The database schema to set default privileges for this role
	Schema pulumi.StringOutput `pulumi:"schema"`
}

// NewDefaultPrivileg registers a new resource with the given unique name, arguments, and options.
func NewDefaultPrivileg(ctx *pulumi.Context,
	name string, args *DefaultPrivilegArgs, opts ...pulumi.ResourceOption) (*DefaultPrivileg, error) {
	if args == nil || args.Database == nil {
		return nil, errors.New("missing required argument 'Database'")
	}
	if args == nil || args.ObjectType == nil {
		return nil, errors.New("missing required argument 'ObjectType'")
	}
	if args == nil || args.Owner == nil {
		return nil, errors.New("missing required argument 'Owner'")
	}
	if args == nil || args.Privileges == nil {
		return nil, errors.New("missing required argument 'Privileges'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.Schema == nil {
		return nil, errors.New("missing required argument 'Schema'")
	}
	if args == nil {
		args = &DefaultPrivilegArgs{}
	}
	var resource DefaultPrivileg
	err := ctx.RegisterResource("postgresql:index/defaultPrivileg:DefaultPrivileg", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultPrivileg gets an existing DefaultPrivileg resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultPrivileg(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultPrivilegState, opts ...pulumi.ResourceOption) (*DefaultPrivileg, error) {
	var resource DefaultPrivileg
	err := ctx.ReadResource("postgresql:index/defaultPrivileg:DefaultPrivileg", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultPrivileg resources.
type defaultPrivilegState struct {
	// The database to grant default privileges for this role
	Database *string `pulumi:"database"`
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type)
	ObjectType *string `pulumi:"objectType"`
	// Target role for which to alter default privileges.
	Owner *string `pulumi:"owner"`
	// The list of privileges to apply as default privileges
	Privileges []string `pulumi:"privileges"`
	// The name of the role to which grant default privileges on
	Role *string `pulumi:"role"`
	// The database schema to set default privileges for this role
	Schema *string `pulumi:"schema"`
}

type DefaultPrivilegState struct {
	// The database to grant default privileges for this role
	Database pulumi.StringPtrInput
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type)
	ObjectType pulumi.StringPtrInput
	// Target role for which to alter default privileges.
	Owner pulumi.StringPtrInput
	// The list of privileges to apply as default privileges
	Privileges pulumi.StringArrayInput
	// The name of the role to which grant default privileges on
	Role pulumi.StringPtrInput
	// The database schema to set default privileges for this role
	Schema pulumi.StringPtrInput
}

func (DefaultPrivilegState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultPrivilegState)(nil)).Elem()
}

type defaultPrivilegArgs struct {
	// The database to grant default privileges for this role
	Database string `pulumi:"database"`
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type)
	ObjectType string `pulumi:"objectType"`
	// Target role for which to alter default privileges.
	Owner string `pulumi:"owner"`
	// The list of privileges to apply as default privileges
	Privileges []string `pulumi:"privileges"`
	// The name of the role to which grant default privileges on
	Role string `pulumi:"role"`
	// The database schema to set default privileges for this role
	Schema string `pulumi:"schema"`
}

// The set of arguments for constructing a DefaultPrivileg resource.
type DefaultPrivilegArgs struct {
	// The database to grant default privileges for this role
	Database pulumi.StringInput
	// The PostgreSQL object type to set the default privileges on (one of: table, sequence, function, type)
	ObjectType pulumi.StringInput
	// Target role for which to alter default privileges.
	Owner pulumi.StringInput
	// The list of privileges to apply as default privileges
	Privileges pulumi.StringArrayInput
	// The name of the role to which grant default privileges on
	Role pulumi.StringInput
	// The database schema to set default privileges for this role
	Schema pulumi.StringInput
}

func (DefaultPrivilegArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultPrivilegArgs)(nil)).Elem()
}
