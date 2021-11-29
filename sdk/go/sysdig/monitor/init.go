// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitor

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-sysdig/sdk/go/sysdig"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "sysdig:Monitor/alertAnomaly:AlertAnomaly":
		r = &AlertAnomaly{}
	case "sysdig:Monitor/alertDashboard:AlertDashboard":
		r = &AlertDashboard{}
	case "sysdig:Monitor/alertDowntime:AlertDowntime":
		r = &AlertDowntime{}
	case "sysdig:Monitor/alertEvent:AlertEvent":
		r = &AlertEvent{}
	case "sysdig:Monitor/alertGroupOutlier:AlertGroupOutlier":
		r = &AlertGroupOutlier{}
	case "sysdig:Monitor/alertMetric:AlertMetric":
		r = &AlertMetric{}
	case "sysdig:Monitor/alertPromql:AlertPromql":
		r = &AlertPromql{}
	case "sysdig:Monitor/notificationChannelEmail:NotificationChannelEmail":
		r = &NotificationChannelEmail{}
	case "sysdig:Monitor/notificationChannelOpsgenie:NotificationChannelOpsgenie":
		r = &NotificationChannelOpsgenie{}
	case "sysdig:Monitor/notificationChannelPagerduty:NotificationChannelPagerduty":
		r = &NotificationChannelPagerduty{}
	case "sysdig:Monitor/notificationChannelSlack:NotificationChannelSlack":
		r = &NotificationChannelSlack{}
	case "sysdig:Monitor/notificationChannelSns:NotificationChannelSns":
		r = &NotificationChannelSns{}
	case "sysdig:Monitor/notificationChannelVictorops:NotificationChannelVictorops":
		r = &NotificationChannelVictorops{}
	case "sysdig:Monitor/notificationChannelWebhook:NotificationChannelWebhook":
		r = &NotificationChannelWebhook{}
	case "sysdig:Monitor/team:Team":
		r = &Team{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := sysdig.PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"sysdig",
		"Monitor/alertAnomaly",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sysdig",
		"Monitor/alertDashboard",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sysdig",
		"Monitor/alertDowntime",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sysdig",
		"Monitor/alertEvent",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sysdig",
		"Monitor/alertGroupOutlier",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sysdig",
		"Monitor/alertMetric",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sysdig",
		"Monitor/alertPromql",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sysdig",
		"Monitor/notificationChannelEmail",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sysdig",
		"Monitor/notificationChannelOpsgenie",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sysdig",
		"Monitor/notificationChannelPagerduty",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sysdig",
		"Monitor/notificationChannelSlack",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sysdig",
		"Monitor/notificationChannelSns",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sysdig",
		"Monitor/notificationChannelVictorops",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sysdig",
		"Monitor/notificationChannelWebhook",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"sysdig",
		"Monitor/team",
		&module{version},
	)
}
