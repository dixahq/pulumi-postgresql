// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The ``postgresql..Extension`` resource creates and manages an extension on a PostgreSQL
 * server.
 *
 *
 * ## Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as postgresql from "@pulumi/postgresql";
 *
 * const myExtension = new postgresql.Extension("myExtension", {});
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-postgresql/blob/master/website/docs/r/postgresql_extension.html.markdown.
 */
export class Extension extends pulumi.CustomResource {
    /**
     * Get an existing Extension resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ExtensionState, opts?: pulumi.CustomResourceOptions): Extension {
        return new Extension(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'postgresql:index/extension:Extension';

    /**
     * Returns true if the given object is an instance of Extension.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Extension {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Extension.__pulumiType;
    }

    /**
     * Which database to create the extension on. Defaults to provider database.
     */
    public readonly database!: pulumi.Output<string>;
    /**
     * The name of the extension.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Sets the schema of an extension.
     */
    public readonly schema!: pulumi.Output<string>;
    /**
     * Sets the version number of the extension.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a Extension resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ExtensionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ExtensionArgs | ExtensionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ExtensionState | undefined;
            inputs["database"] = state ? state.database : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["schema"] = state ? state.schema : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as ExtensionArgs | undefined;
            inputs["database"] = args ? args.database : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["schema"] = args ? args.schema : undefined;
            inputs["version"] = args ? args.version : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Extension.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Extension resources.
 */
export interface ExtensionState {
    /**
     * Which database to create the extension on. Defaults to provider database.
     */
    readonly database?: pulumi.Input<string>;
    /**
     * The name of the extension.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Sets the schema of an extension.
     */
    readonly schema?: pulumi.Input<string>;
    /**
     * Sets the version number of the extension.
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Extension resource.
 */
export interface ExtensionArgs {
    /**
     * Which database to create the extension on. Defaults to provider database.
     */
    readonly database?: pulumi.Input<string>;
    /**
     * The name of the extension.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Sets the schema of an extension.
     */
    readonly schema?: pulumi.Input<string>;
    /**
     * Sets the version number of the extension.
     */
    readonly version?: pulumi.Input<string>;
}
