// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sysdig
{
    [SysdigResourceType("sysdig:index/user:User")]
    public partial class User : Pulumi.CustomResource
    {
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        [Output("firstName")]
        public Output<string?> FirstName { get; private set; } = null!;

        [Output("lastName")]
        public Output<string?> LastName { get; private set; } = null!;

        [Output("systemRole")]
        public Output<string?> SystemRole { get; private set; } = null!;

        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("sysdig:index/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("sysdig:index/user:User", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : Pulumi.ResourceArgs
    {
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        [Input("firstName")]
        public Input<string>? FirstName { get; set; }

        [Input("lastName")]
        public Input<string>? LastName { get; set; }

        [Input("systemRole")]
        public Input<string>? SystemRole { get; set; }

        public UserArgs()
        {
        }
    }

    public sealed class UserState : Pulumi.ResourceArgs
    {
        [Input("email")]
        public Input<string>? Email { get; set; }

        [Input("firstName")]
        public Input<string>? FirstName { get; set; }

        [Input("lastName")]
        public Input<string>? LastName { get; set; }

        [Input("systemRole")]
        public Input<string>? SystemRole { get; set; }

        [Input("version")]
        public Input<int>? Version { get; set; }

        public UserState()
        {
        }
    }
}