// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Sysdig.Monitor
{
    [SysdigResourceType("sysdig:Monitor/alertEvent:AlertEvent")]
    public partial class AlertEvent : Pulumi.CustomResource
    {
        [Output("capture")]
        public Output<Outputs.AlertEventCapture?> Capture { get; private set; } = null!;

        [Output("customNotification")]
        public Output<Outputs.AlertEventCustomNotification?> CustomNotification { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        [Output("eventCount")]
        public Output<int> EventCount { get; private set; } = null!;

        [Output("eventName")]
        public Output<string> EventName { get; private set; } = null!;

        [Output("eventRel")]
        public Output<string> EventRel { get; private set; } = null!;

        [Output("multipleAlertsBies")]
        public Output<ImmutableArray<string>> MultipleAlertsBies { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("notificationChannels")]
        public Output<ImmutableArray<int>> NotificationChannels { get; private set; } = null!;

        [Output("renotificationMinutes")]
        public Output<int?> RenotificationMinutes { get; private set; } = null!;

        [Output("scope")]
        public Output<string?> Scope { get; private set; } = null!;

        [Output("severity")]
        public Output<int?> Severity { get; private set; } = null!;

        [Output("source")]
        public Output<string> Source { get; private set; } = null!;

        [Output("team")]
        public Output<int> Team { get; private set; } = null!;

        [Output("triggerAfterMinutes")]
        public Output<int> TriggerAfterMinutes { get; private set; } = null!;

        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a AlertEvent resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AlertEvent(string name, AlertEventArgs args, CustomResourceOptions? options = null)
            : base("sysdig:Monitor/alertEvent:AlertEvent", name, args ?? new AlertEventArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AlertEvent(string name, Input<string> id, AlertEventState? state = null, CustomResourceOptions? options = null)
            : base("sysdig:Monitor/alertEvent:AlertEvent", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AlertEvent resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AlertEvent Get(string name, Input<string> id, AlertEventState? state = null, CustomResourceOptions? options = null)
        {
            return new AlertEvent(name, id, state, options);
        }
    }

    public sealed class AlertEventArgs : Pulumi.ResourceArgs
    {
        [Input("capture")]
        public Input<Inputs.AlertEventCaptureArgs>? Capture { get; set; }

        [Input("customNotification")]
        public Input<Inputs.AlertEventCustomNotificationArgs>? CustomNotification { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("eventCount", required: true)]
        public Input<int> EventCount { get; set; } = null!;

        [Input("eventName", required: true)]
        public Input<string> EventName { get; set; } = null!;

        [Input("eventRel", required: true)]
        public Input<string> EventRel { get; set; } = null!;

        [Input("multipleAlertsBies")]
        private InputList<string>? _multipleAlertsBies;
        public InputList<string> MultipleAlertsBies
        {
            get => _multipleAlertsBies ?? (_multipleAlertsBies = new InputList<string>());
            set => _multipleAlertsBies = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationChannels")]
        private InputList<int>? _notificationChannels;
        public InputList<int> NotificationChannels
        {
            get => _notificationChannels ?? (_notificationChannels = new InputList<int>());
            set => _notificationChannels = value;
        }

        [Input("renotificationMinutes")]
        public Input<int>? RenotificationMinutes { get; set; }

        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("severity")]
        public Input<int>? Severity { get; set; }

        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        [Input("triggerAfterMinutes", required: true)]
        public Input<int> TriggerAfterMinutes { get; set; } = null!;

        public AlertEventArgs()
        {
        }
    }

    public sealed class AlertEventState : Pulumi.ResourceArgs
    {
        [Input("capture")]
        public Input<Inputs.AlertEventCaptureGetArgs>? Capture { get; set; }

        [Input("customNotification")]
        public Input<Inputs.AlertEventCustomNotificationGetArgs>? CustomNotification { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("eventCount")]
        public Input<int>? EventCount { get; set; }

        [Input("eventName")]
        public Input<string>? EventName { get; set; }

        [Input("eventRel")]
        public Input<string>? EventRel { get; set; }

        [Input("multipleAlertsBies")]
        private InputList<string>? _multipleAlertsBies;
        public InputList<string> MultipleAlertsBies
        {
            get => _multipleAlertsBies ?? (_multipleAlertsBies = new InputList<string>());
            set => _multipleAlertsBies = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notificationChannels")]
        private InputList<int>? _notificationChannels;
        public InputList<int> NotificationChannels
        {
            get => _notificationChannels ?? (_notificationChannels = new InputList<int>());
            set => _notificationChannels = value;
        }

        [Input("renotificationMinutes")]
        public Input<int>? RenotificationMinutes { get; set; }

        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("severity")]
        public Input<int>? Severity { get; set; }

        [Input("source")]
        public Input<string>? Source { get; set; }

        [Input("team")]
        public Input<int>? Team { get; set; }

        [Input("triggerAfterMinutes")]
        public Input<int>? TriggerAfterMinutes { get; set; }

        [Input("version")]
        public Input<int>? Version { get; set; }

        public AlertEventState()
        {
        }
    }
}