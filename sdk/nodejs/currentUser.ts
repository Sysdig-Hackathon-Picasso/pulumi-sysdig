// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function currentUser(opts?: pulumi.InvokeOptions): Promise<CurrentUserResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("sysdig:index/currentUser:CurrentUser", {
    }, opts);
}

/**
 * A collection of values returned by CurrentUser.
 */
export interface CurrentUserResult {
    readonly email: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly lastName: string;
    readonly name: string;
    readonly systemRole: string;
}