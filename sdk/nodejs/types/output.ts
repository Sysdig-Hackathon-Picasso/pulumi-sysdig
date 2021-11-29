// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export namespace Monitor {
    export interface AlertAnomalyCapture {
        duration: number;
        filename: string;
        filter?: string;
    }

    export interface AlertAnomalyCustomNotification {
        append?: string;
        prepend?: string;
        title: string;
    }

    export interface AlertDashboardPanel {
        autosizeText?: boolean;
        content?: string;
        description?: string;
        height: number;
        name: string;
        posX: number;
        posY: number;
        queries?: outputs.Monitor.AlertDashboardPanelQuery[];
        transparentBackground?: boolean;
        type: string;
        visibleTitle?: boolean;
        width: number;
    }

    export interface AlertDashboardPanelQuery {
        promql: string;
        unit: string;
    }

    export interface AlertDashboardScope {
        comparator?: string;
        metric: string;
        values?: string[];
        variable?: string;
    }

    export interface AlertDowntimeCapture {
        duration: number;
        filename: string;
        filter?: string;
    }

    export interface AlertDowntimeCustomNotification {
        append?: string;
        prepend?: string;
        title: string;
    }

    export interface AlertEventCapture {
        duration: number;
        filename: string;
        filter?: string;
    }

    export interface AlertEventCustomNotification {
        append?: string;
        prepend?: string;
        title: string;
    }

    export interface AlertGroupOutlierCapture {
        duration: number;
        filename: string;
        filter?: string;
    }

    export interface AlertGroupOutlierCustomNotification {
        append?: string;
        prepend?: string;
        title: string;
    }

    export interface AlertMetricCapture {
        duration: number;
        filename: string;
        filter?: string;
    }

    export interface AlertMetricCustomNotification {
        append?: string;
        prepend?: string;
        title: string;
    }

    export interface AlertPromqlCapture {
        duration: number;
        filename: string;
        filter?: string;
    }

    export interface AlertPromqlCustomNotification {
        append?: string;
        prepend?: string;
        title: string;
    }

    export interface TeamEntrypoint {
        selection?: string;
        type: string;
    }

    export interface TeamUserRole {
        email: string;
        role?: string;
    }

}

export namespace Secure {
    export interface PolicyAction {
        captures?: outputs.Secure.PolicyActionCapture[];
        container?: string;
    }

    export interface PolicyActionCapture {
        secondsAfterEvent: number;
        secondsBeforeEvent: number;
    }

    export interface RuleFalcoException {
        comps: string[];
        fields: string[];
        name: string;
        values: string;
    }

    export interface RuleFilesystemReadOnly {
        matching?: boolean;
        paths: string[];
    }

    export interface RuleFilesystemReadWrite {
        matching?: boolean;
        paths: string[];
    }

    export interface RuleNetworkTcp {
        matching?: boolean;
        ports: number[];
    }

    export interface RuleNetworkUdp {
        matching?: boolean;
        ports: number[];
    }

    export interface TeamUserRole {
        email: string;
        role?: string;
    }

}
