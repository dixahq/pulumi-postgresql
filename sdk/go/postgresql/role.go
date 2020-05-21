// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package postgresql

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The ``.Role`` resource creates and manages a role on a PostgreSQL
// server.
//
// When a ``.Role`` resource is removed, the PostgreSQL ROLE will
// automatically run a [`REASSIGN
// OWNED`](https://www.postgresql.org/docs/current/static/sql-reassign-owned.html)
// and [`DROP
// OWNED`](https://www.postgresql.org/docs/current/static/sql-drop-owned.html) to
// the `CURRENT_USER` (normally the connected user for the provider).  If the
// specified PostgreSQL ROLE owns objects in multiple PostgreSQL databases in the
// same PostgreSQL Cluster, one PostgreSQL provider per database must be created
// and all but the final ``.Role`` must specify a `skipDropRole`.
//
// > **Note:** All arguments including role name and password will be stored in the raw state as plain-text.
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
type Role struct {
	pulumi.CustomResourceState

	// Defines whether a role bypasses every
	// row-level security (RLS) policy.  Default value is `false`.
	BypassRowLevelSecurity pulumi.BoolPtrOutput `pulumi:"bypassRowLevelSecurity"`
	// If this role can log in, this specifies how
	// many concurrent connections the role can establish. `-1` (the default) means no
	// limit.
	ConnectionLimit pulumi.IntPtrOutput `pulumi:"connectionLimit"`
	// Defines a role's ability to execute `CREATE
	// DATABASE`.  Default value is `false`.
	CreateDatabase pulumi.BoolPtrOutput `pulumi:"createDatabase"`
	// Defines a role's ability to execute `CREATE ROLE`.
	// A role with this privilege can also alter and drop other roles.  Default value
	// is `false`.
	CreateRole pulumi.BoolPtrOutput `pulumi:"createRole"`
	// Deprecated: Rename PostgreSQL role resource attribute "encrypted" to "encrypted_password"
	Encrypted pulumi.StringPtrOutput `pulumi:"encrypted"`
	// Defines whether the password is stored
	// encrypted in the system catalogs.  Default value is `true`.  NOTE: this value
	// is always set (to the conservative and safe value), but may interfere with the
	// behavior of
	// [PostgreSQL's `passwordEncryption` setting](https://www.postgresql.org/docs/current/static/runtime-config-connection.html#GUC-PASSWORD-ENCRYPTION).
	EncryptedPassword pulumi.BoolPtrOutput `pulumi:"encryptedPassword"`
	// Defines whether a role "inherits" the privileges of
	// roles it is a member of.  Default value is `true`.
	Inherit pulumi.BoolPtrOutput `pulumi:"inherit"`
	// Defines whether role is allowed to log in.  Roles without
	// this attribute are useful for managing database privileges, but are not users
	// in the usual sense of the word.  Default value is `false`.
	Login pulumi.BoolPtrOutput `pulumi:"login"`
	// The name of the role. Must be unique on the PostgreSQL
	// server instance where it is configured.
	Name pulumi.StringOutput `pulumi:"name"`
	// Sets the role's password. A password is only of use
	// for roles having the `login` attribute set to true.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Defines whether a role is allowed to initiate
	// streaming replication or put the system in and out of backup mode.  Default
	// value is `false`
	Replication pulumi.BoolPtrOutput `pulumi:"replication"`
	// Defines list of roles which will be granted to this new role.
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// Alters the search path of this new role. Note that
	// due to limitations in the implementation, values cannot contain the substring
	// `", "`.
	SearchPaths pulumi.StringArrayOutput `pulumi:"searchPaths"`
	// When a PostgreSQL ROLE exists in multiple
	// databases and the ROLE is dropped, the
	// [cleanup of ownership of objects](https://www.postgresql.org/docs/current/static/role-removal.html)
	// in each of the respective databases must occur before the ROLE can be dropped
	// from the catalog.  Set this option to true when there are multiple databases
	// in a PostgreSQL cluster using the same PostgreSQL ROLE for object ownership.
	// This is the third and final step taken when removing a ROLE from a database.
	SkipDropRole pulumi.BoolPtrOutput `pulumi:"skipDropRole"`
	// When a PostgreSQL ROLE exists in multiple
	// databases and the ROLE is dropped, a
	// [`REASSIGN OWNED`](https://www.postgresql.org/docs/current/static/sql-reassign-owned.html) in
	// must be executed on each of the respective databases before the `DROP ROLE`
	// can be executed to dropped the ROLE from the catalog.  This is the first and
	// second steps taken when removing a ROLE from a database (the second step being
	// an implicit
	// [`DROP OWNED`](https://www.postgresql.org/docs/current/static/sql-drop-owned.html)).
	SkipReassignOwned pulumi.BoolPtrOutput `pulumi:"skipReassignOwned"`
	// Defines [`statementTimeout`](https://www.postgresql.org/docs/current/runtime-config-client.html#RUNTIME-CONFIG-CLIENT-STATEMENT) setting for this role which allows to abort any statement that takes more than the specified amount of time.
	StatementTimeout pulumi.IntPtrOutput `pulumi:"statementTimeout"`
	// Defines whether the role is a "superuser", and
	// therefore can override all access restrictions within the database.  Default
	// value is `false`.
	Superuser pulumi.BoolPtrOutput `pulumi:"superuser"`
	// Defines the date and time after which the role's
	// password is no longer valid.  Established connections past this `validTime`
	// will have to be manually terminated.  This value corresponds to a PostgreSQL
	// datetime. If omitted or the magic value `NULL` is used, `validUntil` will be
	// set to `infinity`.  Default is `NULL`, therefore `infinity`.
	ValidUntil pulumi.StringPtrOutput `pulumi:"validUntil"`
}

// NewRole registers a new resource with the given unique name, arguments, and options.
func NewRole(ctx *pulumi.Context,
	name string, args *RoleArgs, opts ...pulumi.ResourceOption) (*Role, error) {
	if args == nil {
		args = &RoleArgs{}
	}
	var resource Role
	err := ctx.RegisterResource("postgresql:index/role:Role", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRole gets an existing Role resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleState, opts ...pulumi.ResourceOption) (*Role, error) {
	var resource Role
	err := ctx.ReadResource("postgresql:index/role:Role", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Role resources.
type roleState struct {
	// Defines whether a role bypasses every
	// row-level security (RLS) policy.  Default value is `false`.
	BypassRowLevelSecurity *bool `pulumi:"bypassRowLevelSecurity"`
	// If this role can log in, this specifies how
	// many concurrent connections the role can establish. `-1` (the default) means no
	// limit.
	ConnectionLimit *int `pulumi:"connectionLimit"`
	// Defines a role's ability to execute `CREATE
	// DATABASE`.  Default value is `false`.
	CreateDatabase *bool `pulumi:"createDatabase"`
	// Defines a role's ability to execute `CREATE ROLE`.
	// A role with this privilege can also alter and drop other roles.  Default value
	// is `false`.
	CreateRole *bool `pulumi:"createRole"`
	// Deprecated: Rename PostgreSQL role resource attribute "encrypted" to "encrypted_password"
	Encrypted *string `pulumi:"encrypted"`
	// Defines whether the password is stored
	// encrypted in the system catalogs.  Default value is `true`.  NOTE: this value
	// is always set (to the conservative and safe value), but may interfere with the
	// behavior of
	// [PostgreSQL's `passwordEncryption` setting](https://www.postgresql.org/docs/current/static/runtime-config-connection.html#GUC-PASSWORD-ENCRYPTION).
	EncryptedPassword *bool `pulumi:"encryptedPassword"`
	// Defines whether a role "inherits" the privileges of
	// roles it is a member of.  Default value is `true`.
	Inherit *bool `pulumi:"inherit"`
	// Defines whether role is allowed to log in.  Roles without
	// this attribute are useful for managing database privileges, but are not users
	// in the usual sense of the word.  Default value is `false`.
	Login *bool `pulumi:"login"`
	// The name of the role. Must be unique on the PostgreSQL
	// server instance where it is configured.
	Name *string `pulumi:"name"`
	// Sets the role's password. A password is only of use
	// for roles having the `login` attribute set to true.
	Password *string `pulumi:"password"`
	// Defines whether a role is allowed to initiate
	// streaming replication or put the system in and out of backup mode.  Default
	// value is `false`
	Replication *bool `pulumi:"replication"`
	// Defines list of roles which will be granted to this new role.
	Roles []string `pulumi:"roles"`
	// Alters the search path of this new role. Note that
	// due to limitations in the implementation, values cannot contain the substring
	// `", "`.
	SearchPaths []string `pulumi:"searchPaths"`
	// When a PostgreSQL ROLE exists in multiple
	// databases and the ROLE is dropped, the
	// [cleanup of ownership of objects](https://www.postgresql.org/docs/current/static/role-removal.html)
	// in each of the respective databases must occur before the ROLE can be dropped
	// from the catalog.  Set this option to true when there are multiple databases
	// in a PostgreSQL cluster using the same PostgreSQL ROLE for object ownership.
	// This is the third and final step taken when removing a ROLE from a database.
	SkipDropRole *bool `pulumi:"skipDropRole"`
	// When a PostgreSQL ROLE exists in multiple
	// databases and the ROLE is dropped, a
	// [`REASSIGN OWNED`](https://www.postgresql.org/docs/current/static/sql-reassign-owned.html) in
	// must be executed on each of the respective databases before the `DROP ROLE`
	// can be executed to dropped the ROLE from the catalog.  This is the first and
	// second steps taken when removing a ROLE from a database (the second step being
	// an implicit
	// [`DROP OWNED`](https://www.postgresql.org/docs/current/static/sql-drop-owned.html)).
	SkipReassignOwned *bool `pulumi:"skipReassignOwned"`
	// Defines [`statementTimeout`](https://www.postgresql.org/docs/current/runtime-config-client.html#RUNTIME-CONFIG-CLIENT-STATEMENT) setting for this role which allows to abort any statement that takes more than the specified amount of time.
	StatementTimeout *int `pulumi:"statementTimeout"`
	// Defines whether the role is a "superuser", and
	// therefore can override all access restrictions within the database.  Default
	// value is `false`.
	Superuser *bool `pulumi:"superuser"`
	// Defines the date and time after which the role's
	// password is no longer valid.  Established connections past this `validTime`
	// will have to be manually terminated.  This value corresponds to a PostgreSQL
	// datetime. If omitted or the magic value `NULL` is used, `validUntil` will be
	// set to `infinity`.  Default is `NULL`, therefore `infinity`.
	ValidUntil *string `pulumi:"validUntil"`
}

type RoleState struct {
	// Defines whether a role bypasses every
	// row-level security (RLS) policy.  Default value is `false`.
	BypassRowLevelSecurity pulumi.BoolPtrInput
	// If this role can log in, this specifies how
	// many concurrent connections the role can establish. `-1` (the default) means no
	// limit.
	ConnectionLimit pulumi.IntPtrInput
	// Defines a role's ability to execute `CREATE
	// DATABASE`.  Default value is `false`.
	CreateDatabase pulumi.BoolPtrInput
	// Defines a role's ability to execute `CREATE ROLE`.
	// A role with this privilege can also alter and drop other roles.  Default value
	// is `false`.
	CreateRole pulumi.BoolPtrInput
	// Deprecated: Rename PostgreSQL role resource attribute "encrypted" to "encrypted_password"
	Encrypted pulumi.StringPtrInput
	// Defines whether the password is stored
	// encrypted in the system catalogs.  Default value is `true`.  NOTE: this value
	// is always set (to the conservative and safe value), but may interfere with the
	// behavior of
	// [PostgreSQL's `passwordEncryption` setting](https://www.postgresql.org/docs/current/static/runtime-config-connection.html#GUC-PASSWORD-ENCRYPTION).
	EncryptedPassword pulumi.BoolPtrInput
	// Defines whether a role "inherits" the privileges of
	// roles it is a member of.  Default value is `true`.
	Inherit pulumi.BoolPtrInput
	// Defines whether role is allowed to log in.  Roles without
	// this attribute are useful for managing database privileges, but are not users
	// in the usual sense of the word.  Default value is `false`.
	Login pulumi.BoolPtrInput
	// The name of the role. Must be unique on the PostgreSQL
	// server instance where it is configured.
	Name pulumi.StringPtrInput
	// Sets the role's password. A password is only of use
	// for roles having the `login` attribute set to true.
	Password pulumi.StringPtrInput
	// Defines whether a role is allowed to initiate
	// streaming replication or put the system in and out of backup mode.  Default
	// value is `false`
	Replication pulumi.BoolPtrInput
	// Defines list of roles which will be granted to this new role.
	Roles pulumi.StringArrayInput
	// Alters the search path of this new role. Note that
	// due to limitations in the implementation, values cannot contain the substring
	// `", "`.
	SearchPaths pulumi.StringArrayInput
	// When a PostgreSQL ROLE exists in multiple
	// databases and the ROLE is dropped, the
	// [cleanup of ownership of objects](https://www.postgresql.org/docs/current/static/role-removal.html)
	// in each of the respective databases must occur before the ROLE can be dropped
	// from the catalog.  Set this option to true when there are multiple databases
	// in a PostgreSQL cluster using the same PostgreSQL ROLE for object ownership.
	// This is the third and final step taken when removing a ROLE from a database.
	SkipDropRole pulumi.BoolPtrInput
	// When a PostgreSQL ROLE exists in multiple
	// databases and the ROLE is dropped, a
	// [`REASSIGN OWNED`](https://www.postgresql.org/docs/current/static/sql-reassign-owned.html) in
	// must be executed on each of the respective databases before the `DROP ROLE`
	// can be executed to dropped the ROLE from the catalog.  This is the first and
	// second steps taken when removing a ROLE from a database (the second step being
	// an implicit
	// [`DROP OWNED`](https://www.postgresql.org/docs/current/static/sql-drop-owned.html)).
	SkipReassignOwned pulumi.BoolPtrInput
	// Defines [`statementTimeout`](https://www.postgresql.org/docs/current/runtime-config-client.html#RUNTIME-CONFIG-CLIENT-STATEMENT) setting for this role which allows to abort any statement that takes more than the specified amount of time.
	StatementTimeout pulumi.IntPtrInput
	// Defines whether the role is a "superuser", and
	// therefore can override all access restrictions within the database.  Default
	// value is `false`.
	Superuser pulumi.BoolPtrInput
	// Defines the date and time after which the role's
	// password is no longer valid.  Established connections past this `validTime`
	// will have to be manually terminated.  This value corresponds to a PostgreSQL
	// datetime. If omitted or the magic value `NULL` is used, `validUntil` will be
	// set to `infinity`.  Default is `NULL`, therefore `infinity`.
	ValidUntil pulumi.StringPtrInput
}

func (RoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleState)(nil)).Elem()
}

type roleArgs struct {
	// Defines whether a role bypasses every
	// row-level security (RLS) policy.  Default value is `false`.
	BypassRowLevelSecurity *bool `pulumi:"bypassRowLevelSecurity"`
	// If this role can log in, this specifies how
	// many concurrent connections the role can establish. `-1` (the default) means no
	// limit.
	ConnectionLimit *int `pulumi:"connectionLimit"`
	// Defines a role's ability to execute `CREATE
	// DATABASE`.  Default value is `false`.
	CreateDatabase *bool `pulumi:"createDatabase"`
	// Defines a role's ability to execute `CREATE ROLE`.
	// A role with this privilege can also alter and drop other roles.  Default value
	// is `false`.
	CreateRole *bool `pulumi:"createRole"`
	// Deprecated: Rename PostgreSQL role resource attribute "encrypted" to "encrypted_password"
	Encrypted *string `pulumi:"encrypted"`
	// Defines whether the password is stored
	// encrypted in the system catalogs.  Default value is `true`.  NOTE: this value
	// is always set (to the conservative and safe value), but may interfere with the
	// behavior of
	// [PostgreSQL's `passwordEncryption` setting](https://www.postgresql.org/docs/current/static/runtime-config-connection.html#GUC-PASSWORD-ENCRYPTION).
	EncryptedPassword *bool `pulumi:"encryptedPassword"`
	// Defines whether a role "inherits" the privileges of
	// roles it is a member of.  Default value is `true`.
	Inherit *bool `pulumi:"inherit"`
	// Defines whether role is allowed to log in.  Roles without
	// this attribute are useful for managing database privileges, but are not users
	// in the usual sense of the word.  Default value is `false`.
	Login *bool `pulumi:"login"`
	// The name of the role. Must be unique on the PostgreSQL
	// server instance where it is configured.
	Name *string `pulumi:"name"`
	// Sets the role's password. A password is only of use
	// for roles having the `login` attribute set to true.
	Password *string `pulumi:"password"`
	// Defines whether a role is allowed to initiate
	// streaming replication or put the system in and out of backup mode.  Default
	// value is `false`
	Replication *bool `pulumi:"replication"`
	// Defines list of roles which will be granted to this new role.
	Roles []string `pulumi:"roles"`
	// Alters the search path of this new role. Note that
	// due to limitations in the implementation, values cannot contain the substring
	// `", "`.
	SearchPaths []string `pulumi:"searchPaths"`
	// When a PostgreSQL ROLE exists in multiple
	// databases and the ROLE is dropped, the
	// [cleanup of ownership of objects](https://www.postgresql.org/docs/current/static/role-removal.html)
	// in each of the respective databases must occur before the ROLE can be dropped
	// from the catalog.  Set this option to true when there are multiple databases
	// in a PostgreSQL cluster using the same PostgreSQL ROLE for object ownership.
	// This is the third and final step taken when removing a ROLE from a database.
	SkipDropRole *bool `pulumi:"skipDropRole"`
	// When a PostgreSQL ROLE exists in multiple
	// databases and the ROLE is dropped, a
	// [`REASSIGN OWNED`](https://www.postgresql.org/docs/current/static/sql-reassign-owned.html) in
	// must be executed on each of the respective databases before the `DROP ROLE`
	// can be executed to dropped the ROLE from the catalog.  This is the first and
	// second steps taken when removing a ROLE from a database (the second step being
	// an implicit
	// [`DROP OWNED`](https://www.postgresql.org/docs/current/static/sql-drop-owned.html)).
	SkipReassignOwned *bool `pulumi:"skipReassignOwned"`
	// Defines [`statementTimeout`](https://www.postgresql.org/docs/current/runtime-config-client.html#RUNTIME-CONFIG-CLIENT-STATEMENT) setting for this role which allows to abort any statement that takes more than the specified amount of time.
	StatementTimeout *int `pulumi:"statementTimeout"`
	// Defines whether the role is a "superuser", and
	// therefore can override all access restrictions within the database.  Default
	// value is `false`.
	Superuser *bool `pulumi:"superuser"`
	// Defines the date and time after which the role's
	// password is no longer valid.  Established connections past this `validTime`
	// will have to be manually terminated.  This value corresponds to a PostgreSQL
	// datetime. If omitted or the magic value `NULL` is used, `validUntil` will be
	// set to `infinity`.  Default is `NULL`, therefore `infinity`.
	ValidUntil *string `pulumi:"validUntil"`
}

// The set of arguments for constructing a Role resource.
type RoleArgs struct {
	// Defines whether a role bypasses every
	// row-level security (RLS) policy.  Default value is `false`.
	BypassRowLevelSecurity pulumi.BoolPtrInput
	// If this role can log in, this specifies how
	// many concurrent connections the role can establish. `-1` (the default) means no
	// limit.
	ConnectionLimit pulumi.IntPtrInput
	// Defines a role's ability to execute `CREATE
	// DATABASE`.  Default value is `false`.
	CreateDatabase pulumi.BoolPtrInput
	// Defines a role's ability to execute `CREATE ROLE`.
	// A role with this privilege can also alter and drop other roles.  Default value
	// is `false`.
	CreateRole pulumi.BoolPtrInput
	// Deprecated: Rename PostgreSQL role resource attribute "encrypted" to "encrypted_password"
	Encrypted pulumi.StringPtrInput
	// Defines whether the password is stored
	// encrypted in the system catalogs.  Default value is `true`.  NOTE: this value
	// is always set (to the conservative and safe value), but may interfere with the
	// behavior of
	// [PostgreSQL's `passwordEncryption` setting](https://www.postgresql.org/docs/current/static/runtime-config-connection.html#GUC-PASSWORD-ENCRYPTION).
	EncryptedPassword pulumi.BoolPtrInput
	// Defines whether a role "inherits" the privileges of
	// roles it is a member of.  Default value is `true`.
	Inherit pulumi.BoolPtrInput
	// Defines whether role is allowed to log in.  Roles without
	// this attribute are useful for managing database privileges, but are not users
	// in the usual sense of the word.  Default value is `false`.
	Login pulumi.BoolPtrInput
	// The name of the role. Must be unique on the PostgreSQL
	// server instance where it is configured.
	Name pulumi.StringPtrInput
	// Sets the role's password. A password is only of use
	// for roles having the `login` attribute set to true.
	Password pulumi.StringPtrInput
	// Defines whether a role is allowed to initiate
	// streaming replication or put the system in and out of backup mode.  Default
	// value is `false`
	Replication pulumi.BoolPtrInput
	// Defines list of roles which will be granted to this new role.
	Roles pulumi.StringArrayInput
	// Alters the search path of this new role. Note that
	// due to limitations in the implementation, values cannot contain the substring
	// `", "`.
	SearchPaths pulumi.StringArrayInput
	// When a PostgreSQL ROLE exists in multiple
	// databases and the ROLE is dropped, the
	// [cleanup of ownership of objects](https://www.postgresql.org/docs/current/static/role-removal.html)
	// in each of the respective databases must occur before the ROLE can be dropped
	// from the catalog.  Set this option to true when there are multiple databases
	// in a PostgreSQL cluster using the same PostgreSQL ROLE for object ownership.
	// This is the third and final step taken when removing a ROLE from a database.
	SkipDropRole pulumi.BoolPtrInput
	// When a PostgreSQL ROLE exists in multiple
	// databases and the ROLE is dropped, a
	// [`REASSIGN OWNED`](https://www.postgresql.org/docs/current/static/sql-reassign-owned.html) in
	// must be executed on each of the respective databases before the `DROP ROLE`
	// can be executed to dropped the ROLE from the catalog.  This is the first and
	// second steps taken when removing a ROLE from a database (the second step being
	// an implicit
	// [`DROP OWNED`](https://www.postgresql.org/docs/current/static/sql-drop-owned.html)).
	SkipReassignOwned pulumi.BoolPtrInput
	// Defines [`statementTimeout`](https://www.postgresql.org/docs/current/runtime-config-client.html#RUNTIME-CONFIG-CLIENT-STATEMENT) setting for this role which allows to abort any statement that takes more than the specified amount of time.
	StatementTimeout pulumi.IntPtrInput
	// Defines whether the role is a "superuser", and
	// therefore can override all access restrictions within the database.  Default
	// value is `false`.
	Superuser pulumi.BoolPtrInput
	// Defines the date and time after which the role's
	// password is no longer valid.  Established connections past this `validTime`
	// will have to be manually terminated.  This value corresponds to a PostgreSQL
	// datetime. If omitted or the magic value `NULL` is used, `validUntil` will be
	// set to `infinity`.  Default is `NULL`, therefore `infinity`.
	ValidUntil pulumi.StringPtrInput
}

func (RoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleArgs)(nil)).Elem()
}
