// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class NotificationChannelWebhook extends pulumi.CustomResource {
    /**
     * Get an existing NotificationChannelWebhook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NotificationChannelWebhookState, opts?: pulumi.CustomResourceOptions): NotificationChannelWebhook {
        return new NotificationChannelWebhook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'sysdig:Secure/notificationChannelWebhook:NotificationChannelWebhook';

    /**
     * Returns true if the given object is an instance of NotificationChannelWebhook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NotificationChannelWebhook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NotificationChannelWebhook.__pulumiType;
    }

    public readonly enabled!: pulumi.Output<boolean>;
    public readonly name!: pulumi.Output<string>;
    public readonly notifyWhenOk!: pulumi.Output<boolean>;
    public readonly notifyWhenResolved!: pulumi.Output<boolean>;
    public readonly sendTestNotification!: pulumi.Output<boolean | undefined>;
    public readonly url!: pulumi.Output<string>;
    public /*out*/ readonly version!: pulumi.Output<number>;

    /**
     * Create a NotificationChannelWebhook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NotificationChannelWebhookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NotificationChannelWebhookArgs | NotificationChannelWebhookState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NotificationChannelWebhookState | undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notifyWhenOk"] = state ? state.notifyWhenOk : undefined;
            inputs["notifyWhenResolved"] = state ? state.notifyWhenResolved : undefined;
            inputs["sendTestNotification"] = state ? state.sendTestNotification : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as NotificationChannelWebhookArgs | undefined;
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.notifyWhenOk === undefined) && !opts.urn) {
                throw new Error("Missing required property 'notifyWhenOk'");
            }
            if ((!args || args.notifyWhenResolved === undefined) && !opts.urn) {
                throw new Error("Missing required property 'notifyWhenResolved'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notifyWhenOk"] = args ? args.notifyWhenOk : undefined;
            inputs["notifyWhenResolved"] = args ? args.notifyWhenResolved : undefined;
            inputs["sendTestNotification"] = args ? args.sendTestNotification : undefined;
            inputs["url"] = args ? args.url : undefined;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NotificationChannelWebhook.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NotificationChannelWebhook resources.
 */
export interface NotificationChannelWebhookState {
    enabled?: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    notifyWhenOk?: pulumi.Input<boolean>;
    notifyWhenResolved?: pulumi.Input<boolean>;
    sendTestNotification?: pulumi.Input<boolean>;
    url?: pulumi.Input<string>;
    version?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a NotificationChannelWebhook resource.
 */
export interface NotificationChannelWebhookArgs {
    enabled: pulumi.Input<boolean>;
    name?: pulumi.Input<string>;
    notifyWhenOk: pulumi.Input<boolean>;
    notifyWhenResolved: pulumi.Input<boolean>;
    sendTestNotification?: pulumi.Input<boolean>;
    url: pulumi.Input<string>;
}