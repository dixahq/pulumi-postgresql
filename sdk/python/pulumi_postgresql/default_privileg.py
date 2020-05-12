# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

warnings.warn("postgresql.DefaultPrivileg has been deprecated in favor of postgresql.DefaultPrivileges", DeprecationWarning)
class DefaultPrivileg(pulumi.CustomResource):
    database: pulumi.Output[str]
    """
    The database to grant default privileges for this role
    """
    object_type: pulumi.Output[str]
    """
    The PostgreSQL object type to set the default privileges on (one of: table, sequence)
    """
    owner: pulumi.Output[str]
    """
    Target role for which to alter default privileges.
    """
    privileges: pulumi.Output[list]
    """
    The list of privileges to apply as default privileges
    """
    role: pulumi.Output[str]
    """
    The name of the role to which grant default privileges on
    """
    schema: pulumi.Output[str]
    """
    The database schema to set default privileges for this role
    """
    warnings.warn("postgresql.DefaultPrivileg has been deprecated in favor of postgresql.DefaultPrivileges", DeprecationWarning)
    def __init__(__self__, resource_name, opts=None, database=None, object_type=None, owner=None, privileges=None, role=None, schema=None, __props__=None, __name__=None, __opts__=None):
        """

        Deprecated: postgresql.DefaultPrivileg has been deprecated in favor of postgresql.DefaultPrivileges

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database: The database to grant default privileges for this role
        :param pulumi.Input[str] object_type: The PostgreSQL object type to set the default privileges on (one of: table, sequence)
        :param pulumi.Input[str] owner: Target role for which to alter default privileges.
        :param pulumi.Input[list] privileges: The list of privileges to apply as default privileges
        :param pulumi.Input[str] role: The name of the role to which grant default privileges on
        :param pulumi.Input[str] schema: The database schema to set default privileges for this role
        """
        pulumi.log.warn("DefaultPrivileg is deprecated: postgresql.DefaultPrivileg has been deprecated in favor of postgresql.DefaultPrivileges")
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if database is None:
                raise TypeError("Missing required property 'database'")
            __props__['database'] = database
            if object_type is None:
                raise TypeError("Missing required property 'object_type'")
            __props__['object_type'] = object_type
            if owner is None:
                raise TypeError("Missing required property 'owner'")
            __props__['owner'] = owner
            if privileges is None:
                raise TypeError("Missing required property 'privileges'")
            __props__['privileges'] = privileges
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            if schema is None:
                raise TypeError("Missing required property 'schema'")
            __props__['schema'] = schema
        super(DefaultPrivileg, __self__).__init__(
            'postgresql:index/defaultPrivileg:DefaultPrivileg',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, database=None, object_type=None, owner=None, privileges=None, role=None, schema=None):
        """
        Get an existing DefaultPrivileg resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database: The database to grant default privileges for this role
        :param pulumi.Input[str] object_type: The PostgreSQL object type to set the default privileges on (one of: table, sequence)
        :param pulumi.Input[str] owner: Target role for which to alter default privileges.
        :param pulumi.Input[list] privileges: The list of privileges to apply as default privileges
        :param pulumi.Input[str] role: The name of the role to which grant default privileges on
        :param pulumi.Input[str] schema: The database schema to set default privileges for this role
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["database"] = database
        __props__["object_type"] = object_type
        __props__["owner"] = owner
        __props__["privileges"] = privileges
        __props__["role"] = role
        __props__["schema"] = schema
        return DefaultPrivileg(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

