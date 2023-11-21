// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Foo.Bar.Outputs
{

    [OutputType]
    public sealed class NestedObject
    {
        public readonly ImmutableArray<string> OptProp;
        public readonly string ReqProp;

        [OutputConstructor]
        private NestedObject(
            ImmutableArray<string> optProp,

            string reqProp)
        {
            OptProp = optProp;
            ReqProp = reqProp;
        }
    }
}