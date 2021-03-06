# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables

__all__ = [
    'SchemaPolicy',
    'ProviderClientcert',
]

@pulumi.output_type
class SchemaPolicy(dict):
    def __init__(__self__, *,
                 create: Optional[bool] = None,
                 create_with_grant: Optional[bool] = None,
                 role: Optional[str] = None,
                 usage: Optional[bool] = None,
                 usage_with_grant: Optional[bool] = None):
        """
        :param bool create: Should the specified ROLE have CREATE privileges to the specified SCHEMA.
        :param bool create_with_grant: Should the specified ROLE have CREATE privileges to the specified SCHEMA and the ability to GRANT the CREATE privilege to other ROLEs.
        :param str role: The ROLE who is receiving the policy.  If this value is empty or not specified it implies the policy is referring to the [`PUBLIC` role](https://www.postgresql.org/docs/current/static/sql-grant.html).
        :param bool usage: Should the specified ROLE have USAGE privileges to the specified SCHEMA.
        :param bool usage_with_grant: Should the specified ROLE have USAGE privileges to the specified SCHEMA and the ability to GRANT the USAGE privilege to other ROLEs.
        """
        if create is not None:
            pulumi.set(__self__, "create", create)
        if create_with_grant is not None:
            pulumi.set(__self__, "create_with_grant", create_with_grant)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if usage is not None:
            pulumi.set(__self__, "usage", usage)
        if usage_with_grant is not None:
            pulumi.set(__self__, "usage_with_grant", usage_with_grant)

    @property
    @pulumi.getter
    def create(self) -> Optional[bool]:
        """
        Should the specified ROLE have CREATE privileges to the specified SCHEMA.
        """
        return pulumi.get(self, "create")

    @property
    @pulumi.getter(name="createWithGrant")
    def create_with_grant(self) -> Optional[bool]:
        """
        Should the specified ROLE have CREATE privileges to the specified SCHEMA and the ability to GRANT the CREATE privilege to other ROLEs.
        """
        return pulumi.get(self, "create_with_grant")

    @property
    @pulumi.getter
    def role(self) -> Optional[str]:
        """
        The ROLE who is receiving the policy.  If this value is empty or not specified it implies the policy is referring to the [`PUBLIC` role](https://www.postgresql.org/docs/current/static/sql-grant.html).
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def usage(self) -> Optional[bool]:
        """
        Should the specified ROLE have USAGE privileges to the specified SCHEMA.
        """
        return pulumi.get(self, "usage")

    @property
    @pulumi.getter(name="usageWithGrant")
    def usage_with_grant(self) -> Optional[bool]:
        """
        Should the specified ROLE have USAGE privileges to the specified SCHEMA and the ability to GRANT the USAGE privilege to other ROLEs.
        """
        return pulumi.get(self, "usage_with_grant")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProviderClientcert(dict):
    def __init__(__self__, *,
                 cert: str,
                 key: str):
        pulumi.set(__self__, "cert", cert)
        pulumi.set(__self__, "key", key)

    @property
    @pulumi.getter
    def cert(self) -> str:
        return pulumi.get(self, "cert")

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


