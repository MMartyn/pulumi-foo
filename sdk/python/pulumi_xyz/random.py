# coding=utf-8
# *** WARNING: this file was generated by pulumi. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RandomArgs', 'Random']

@pulumi.input_type
class RandomArgs:
    def __init__(__self__, *,
                 length: pulumi.Input[int]):
        """
        The set of arguments for constructing a Random resource.
        """
        pulumi.set(__self__, "length", length)

    @property
    @pulumi.getter
    def length(self) -> pulumi.Input[int]:
        return pulumi.get(self, "length")

    @length.setter
    def length(self, value: pulumi.Input[int]):
        pulumi.set(self, "length", value)


class Random(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 length: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a Random resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RandomArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Random resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param RandomArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RandomArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 length: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RandomArgs.__new__(RandomArgs)

            if length is None and not opts.urn:
                raise TypeError("Missing required property 'length'")
            __props__.__dict__["length"] = length
            __props__.__dict__["result"] = None
        super(Random, __self__).__init__(
            'foo:index:Random',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Random':
        """
        Get an existing Random resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RandomArgs.__new__(RandomArgs)

        __props__.__dict__["length"] = None
        __props__.__dict__["result"] = None
        return Random(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def length(self) -> pulumi.Output[int]:
        return pulumi.get(self, "length")

    @property
    @pulumi.getter
    def result(self) -> pulumi.Output[str]:
        return pulumi.get(self, "result")

