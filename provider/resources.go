// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sysdig

import (
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/draios/terraform-provider-sysdig/sysdig"
	"github.com/pulumi/pulumi-sysdig/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "sysdig"
	// modules:
	mainMod    = "index"   // the sysdig module
	monitorMod = "Monitor" // the sysdig module
	secureMod  = "Secure"  // the sysdig module
	iamMod     = "Iam"     // the sysdig module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// boolRef returns a reference to the bool argument.
func boolRef(b bool) *bool {
	return &b
}

// stringValue gets a string value from a property map if present, else ""
func stringValue(vars resource.PropertyMap, prop resource.PropertyKey) string {
	val, ok := vars[prop]
	if ok && val.IsString() {
		return val.StringValue()
	}
	return ""
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// managedByPulumi is a default used for some managed resources, in the absence of something more meaningful.
var managedByPulumi = &tfbridge.DefaultInfo{Value: "Managed by Pulumi"}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(sysdig.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "sysdig",
		Description: "A Pulumi package for creating and managing sysdig cloud resources.",
		Keywords:    []string{"pulumi", "sysdig"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-sysdig",
		Config:      map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: makeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: makeResource(mainMod, "IamRole")}
			//
			// "aws_acm_certificate": {
			// 	Tok: makeResource(mainMod, "Certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: makeType(mainPkg, "Tags")},
			// 	},
			// },
			"sysdig_monitor_alert_anomaly":                  {Tok: makeResource(monitorMod, "AlertAnomaly")},
			"sysdig_monitor_alert_downtime":                 {Tok: makeResource(monitorMod, "AlertDowntime")},
			"sysdig_monitor_alert_event":                    {Tok: makeResource(monitorMod, "AlertEvent")},
			"sysdig_monitor_alert_group_outlier":            {Tok: makeResource(monitorMod, "GroupOutlier")},
			"sysdig_monitor_alert_metric":                   {Tok: makeResource(monitorMod, "Metric")},
			"sysdig_monitor_alert_promql":                   {Tok: makeResource(monitorMod, "Promql")},
			"sysdig_monitor_dashboard":                      {Tok: makeResource(monitorMod, "AlertDashboard")},
			"sysdig_monitor_notification_channel_email":     {Tok: makeResource(monitorMod, "NotificationChannelEmail")},
			"sysdig_monitor_notification_channel_opsgenie":  {Tok: makeResource(monitorMod, "NotificationChannelOpsgenie")},
			"sysdig_monitor_notification_channel_pagerduty": {Tok: makeResource(monitorMod, "NotificationChannelPagerduty")},
			"sysdig_monitor_notification_channel_slack":     {Tok: makeResource(monitorMod, "NotificationChannelSlack")},
			"sysdig_monitor_notification_channel_sns":       {Tok: makeResource(monitorMod, "NotificationChannelSns")},
			"sysdig_monitor_notification_channel_victorops": {Tok: makeResource(monitorMod, "NotificationChannelVictorops")},
			"sysdig_monitor_notification_channel_webhook":   {Tok: makeResource(monitorMod, "NotificationChannelWebhook")},
			"sysdig_monitor_team":                           {Tok: makeResource(monitorMod, "Team")},
			"sysdig_secure_benchmark_task":                  {Tok: makeResource(secureMod, "BenchmarkTask")},
			"sysdig_secure_cloud_account":                   {Tok: makeResource(secureMod, "CloudAccount")},
			"sysdig_secure_list":                            {Tok: makeResource(secureMod, "List")},
			"sysdig_secure_macro":                           {Tok: makeResource(secureMod, "Macro")},
			"sysdig_secure_notification_channel_email":      {Tok: makeResource(secureMod, "NotificationChannelEmail")},
			"sysdig_secure_notification_channel_opsgenie":   {Tok: makeResource(secureMod, "NotificationChannelOpsgenie")},
			"sysdig_secure_notification_channel_pagerduty":  {Tok: makeResource(secureMod, "NotificationChannelPagerduty")},
			"sysdig_secure_notification_channel_slack":      {Tok: makeResource(secureMod, "NotificationChannelSlack")},
			"sysdig_secure_notification_channel_sns":        {Tok: makeResource(secureMod, "NotificationChannelSns")},
			"sysdig_secure_notification_channel_victorops":  {Tok: makeResource(secureMod, "NotificationChannelVictorops")},
			"sysdig_secure_notification_channel_webhook":    {Tok: makeResource(secureMod, "NotificationChannelWebhook")},
			"sysdig_secure_policy":                          {Tok: makeResource(secureMod, "Policy")},
			"sysdig_secure_rule_container":                  {Tok: makeResource(secureMod, "RuleContainer")},
			"sysdig_secure_rule_falco":                      {Tok: makeResource(secureMod, "RuleFalco")},
			"sysdig_secure_rule_filesystem":                 {Tok: makeResource(secureMod, "RuleFilesystem")},
			"sysdig_secure_rule_network":                    {Tok: makeResource(secureMod, "RuleNetwork")},
			"sysdig_secure_rule_process":                    {Tok: makeResource(secureMod, "RuleProcess")},
			"sysdig_secure_rule_syscall":                    {Tok: makeResource(secureMod, "RuleSyscall")},
			"sysdig_secure_team":                            {Tok: makeResource(secureMod, "Team")},
			"sysdig_secure_vulnerability_exception":         {Tok: makeResource(secureMod, "VulnerabilityException")},
			"sysdig_secure_vulnerability_exception_list":    {Tok: makeResource(secureMod, "VulnerabilityExceptionList")},
			"sysdig_user": {Tok: makeResource(mainMod, "User")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: makeDataSource(mainMod, "getAmi")},
			"sysdig_current_user":                  {Tok: makeDataSource(mainMod, "CurrentUser")},
			"sysdig_fargate_workload_agent":        {Tok: makeDataSource(mainMod, "FargateWorkloadAgent")},
			"sysdig_secure_notification_channel":   {Tok: makeDataSource(secureMod, "NotificationChannel")},
			"sysdig_secure_trusted_cloud_identity": {Tok: makeDataSource(secureMod, "TrustedCloudIdentity")},
			"sysdig_user":                          {Tok: makeDataSource(mainMod, "User")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
