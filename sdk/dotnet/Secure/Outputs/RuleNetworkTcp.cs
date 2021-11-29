// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sysdig.Secure.Outputs
{

    [OutputType]
    public sealed class RuleNetworkTcp
    {
        public readonly bool? Matching;
        public readonly ImmutableArray<int> Ports;

        [OutputConstructor]
        private RuleNetworkTcp(
            bool? matching,

            ImmutableArray<int> ports)
        {
            Matching = matching;
            Ports = ports;
        }
    }
}