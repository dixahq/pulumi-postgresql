// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class Database extends pulumi.CustomResource {
    /**
     * Get an existing Database resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseState, opts?: pulumi.CustomResourceOptions): Database {
        return new Database(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'postgresql:index/database:Database';

    /**
     * Returns true if the given object is an instance of Database.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Database {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Database.__pulumiType;
    }

    /**
     * If `false` then no one can connect to this
     * database. The default is `true`, allowing connections (except as restricted by
     * other mechanisms, such as `GRANT` or `REVOKE CONNECT`).
     */
    public readonly allowConnections!: pulumi.Output<boolean | undefined>;
    /**
     * How many concurrent connections can be
     * established to this database. `-1` (the default) means no limit.
     */
    public readonly connectionLimit!: pulumi.Output<number | undefined>;
    /**
     * Character set encoding to use in the new database
     */
    public readonly encoding!: pulumi.Output<string>;
    /**
     * If `true`, then this database can be cloned by any
     * user with `CREATEDB` privileges; if `false` (the default), then only
     * superusers or the owner of the database can clone it.
     */
    public readonly isTemplate!: pulumi.Output<boolean>;
    /**
     * Collation order (LC_COLLATE) to use in the new database
     */
    public readonly lcCollate!: pulumi.Output<string>;
    /**
     * Character classification (LC_CTYPE) to use in the new database
     */
    public readonly lcCtype!: pulumi.Output<string>;
    /**
     * The name of the database. Must be unique on the PostgreSQL
     * server instance where it is configured.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The role name of the user who will own the database, or
     * `DEFAULT` to use the default (namely, the user executing the command). To
     * create a database owned by another role or to change the owner of an existing
     * database, you must be a direct or indirect member of the specified role, or
     * the username in the provider is a superuser.
     */
    public readonly owner!: pulumi.Output<string>;
    /**
     * The name of the tablespace that will be
     * associated with the database, or `DEFAULT` to use the template database's
     * tablespace.  This tablespace will be the default tablespace used for objects
     * created in this database.
     */
    public readonly tablespaceName!: pulumi.Output<string>;
    /**
     * The name of the template from which to create the new database
     */
    public readonly template!: pulumi.Output<string>;

    /**
     * Create a Database resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseArgs | DatabaseState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DatabaseState | undefined;
            inputs["allowConnections"] = state ? state.allowConnections : undefined;
            inputs["connectionLimit"] = state ? state.connectionLimit : undefined;
            inputs["encoding"] = state ? state.encoding : undefined;
            inputs["isTemplate"] = state ? state.isTemplate : undefined;
            inputs["lcCollate"] = state ? state.lcCollate : undefined;
            inputs["lcCtype"] = state ? state.lcCtype : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["owner"] = state ? state.owner : undefined;
            inputs["tablespaceName"] = state ? state.tablespaceName : undefined;
            inputs["template"] = state ? state.template : undefined;
        } else {
            const args = argsOrState as DatabaseArgs | undefined;
            inputs["allowConnections"] = args ? args.allowConnections : undefined;
            inputs["connectionLimit"] = args ? args.connectionLimit : undefined;
            inputs["encoding"] = args ? args.encoding : undefined;
            inputs["isTemplate"] = args ? args.isTemplate : undefined;
            inputs["lcCollate"] = args ? args.lcCollate : undefined;
            inputs["lcCtype"] = args ? args.lcCtype : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["owner"] = args ? args.owner : undefined;
            inputs["tablespaceName"] = args ? args.tablespaceName : undefined;
            inputs["template"] = args ? args.template : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Database.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Database resources.
 */
export interface DatabaseState {
    /**
     * If `false` then no one can connect to this
     * database. The default is `true`, allowing connections (except as restricted by
     * other mechanisms, such as `GRANT` or `REVOKE CONNECT`).
     */
    readonly allowConnections?: pulumi.Input<boolean>;
    /**
     * How many concurrent connections can be
     * established to this database. `-1` (the default) means no limit.
     */
    readonly connectionLimit?: pulumi.Input<number>;
    /**
     * Character set encoding to use in the new database
     */
    readonly encoding?: pulumi.Input<string>;
    /**
     * If `true`, then this database can be cloned by any
     * user with `CREATEDB` privileges; if `false` (the default), then only
     * superusers or the owner of the database can clone it.
     */
    readonly isTemplate?: pulumi.Input<boolean>;
    /**
     * Collation order (LC_COLLATE) to use in the new database
     */
    readonly lcCollate?: pulumi.Input<string>;
    /**
     * Character classification (LC_CTYPE) to use in the new database
     */
    readonly lcCtype?: pulumi.Input<string>;
    /**
     * The name of the database. Must be unique on the PostgreSQL
     * server instance where it is configured.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The role name of the user who will own the database, or
     * `DEFAULT` to use the default (namely, the user executing the command). To
     * create a database owned by another role or to change the owner of an existing
     * database, you must be a direct or indirect member of the specified role, or
     * the username in the provider is a superuser.
     */
    readonly owner?: pulumi.Input<string>;
    /**
     * The name of the tablespace that will be
     * associated with the database, or `DEFAULT` to use the template database's
     * tablespace.  This tablespace will be the default tablespace used for objects
     * created in this database.
     */
    readonly tablespaceName?: pulumi.Input<string>;
    /**
     * The name of the template from which to create the new database
     */
    readonly template?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    /**
     * If `false` then no one can connect to this
     * database. The default is `true`, allowing connections (except as restricted by
     * other mechanisms, such as `GRANT` or `REVOKE CONNECT`).
     */
    readonly allowConnections?: pulumi.Input<boolean>;
    /**
     * How many concurrent connections can be
     * established to this database. `-1` (the default) means no limit.
     */
    readonly connectionLimit?: pulumi.Input<number>;
    /**
     * Character set encoding to use in the new database
     */
    readonly encoding?: pulumi.Input<string>;
    /**
     * If `true`, then this database can be cloned by any
     * user with `CREATEDB` privileges; if `false` (the default), then only
     * superusers or the owner of the database can clone it.
     */
    readonly isTemplate?: pulumi.Input<boolean>;
    /**
     * Collation order (LC_COLLATE) to use in the new database
     */
    readonly lcCollate?: pulumi.Input<string>;
    /**
     * Character classification (LC_CTYPE) to use in the new database
     */
    readonly lcCtype?: pulumi.Input<string>;
    /**
     * The name of the database. Must be unique on the PostgreSQL
     * server instance where it is configured.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The role name of the user who will own the database, or
     * `DEFAULT` to use the default (namely, the user executing the command). To
     * create a database owned by another role or to change the owner of an existing
     * database, you must be a direct or indirect member of the specified role, or
     * the username in the provider is a superuser.
     */
    readonly owner?: pulumi.Input<string>;
    /**
     * The name of the tablespace that will be
     * associated with the database, or `DEFAULT` to use the template database's
     * tablespace.  This tablespace will be the default tablespace used for objects
     * created in this database.
     */
    readonly tablespaceName?: pulumi.Input<string>;
    /**
     * The name of the template from which to create the new database
     */
    readonly template?: pulumi.Input<string>;
}
