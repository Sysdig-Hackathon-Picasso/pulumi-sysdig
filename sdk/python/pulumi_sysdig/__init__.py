# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .current_user import *
from .fargate_workload_agent import *
from .provider import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_sysdig.config as __config
    config = __config
    import pulumi_sysdig.monitor as __monitor
    monitor = __monitor
    import pulumi_sysdig.secure as __secure
    secure = __secure
else:
    config = _utilities.lazy_import('pulumi_sysdig.config')
    monitor = _utilities.lazy_import('pulumi_sysdig.monitor')
    secure = _utilities.lazy_import('pulumi_sysdig.secure')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "sysdig",
  "mod": "Monitor/alertAnomaly",
  "fqn": "pulumi_sysdig.monitor",
  "classes": {
   "sysdig:Monitor/alertAnomaly:AlertAnomaly": "AlertAnomaly"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Monitor/alertDashboard",
  "fqn": "pulumi_sysdig.monitor",
  "classes": {
   "sysdig:Monitor/alertDashboard:AlertDashboard": "AlertDashboard"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Monitor/alertDowntime",
  "fqn": "pulumi_sysdig.monitor",
  "classes": {
   "sysdig:Monitor/alertDowntime:AlertDowntime": "AlertDowntime"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Monitor/alertEvent",
  "fqn": "pulumi_sysdig.monitor",
  "classes": {
   "sysdig:Monitor/alertEvent:AlertEvent": "AlertEvent"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Monitor/alertGroupOutlier",
  "fqn": "pulumi_sysdig.monitor",
  "classes": {
   "sysdig:Monitor/alertGroupOutlier:AlertGroupOutlier": "AlertGroupOutlier"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Monitor/alertMetric",
  "fqn": "pulumi_sysdig.monitor",
  "classes": {
   "sysdig:Monitor/alertMetric:AlertMetric": "AlertMetric"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Monitor/alertPromql",
  "fqn": "pulumi_sysdig.monitor",
  "classes": {
   "sysdig:Monitor/alertPromql:AlertPromql": "AlertPromql"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Monitor/notificationChannelEmail",
  "fqn": "pulumi_sysdig.monitor",
  "classes": {
   "sysdig:Monitor/notificationChannelEmail:NotificationChannelEmail": "NotificationChannelEmail"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Monitor/notificationChannelOpsgenie",
  "fqn": "pulumi_sysdig.monitor",
  "classes": {
   "sysdig:Monitor/notificationChannelOpsgenie:NotificationChannelOpsgenie": "NotificationChannelOpsgenie"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Monitor/notificationChannelPagerduty",
  "fqn": "pulumi_sysdig.monitor",
  "classes": {
   "sysdig:Monitor/notificationChannelPagerduty:NotificationChannelPagerduty": "NotificationChannelPagerduty"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Monitor/notificationChannelSlack",
  "fqn": "pulumi_sysdig.monitor",
  "classes": {
   "sysdig:Monitor/notificationChannelSlack:NotificationChannelSlack": "NotificationChannelSlack"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Monitor/notificationChannelSns",
  "fqn": "pulumi_sysdig.monitor",
  "classes": {
   "sysdig:Monitor/notificationChannelSns:NotificationChannelSns": "NotificationChannelSns"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Monitor/notificationChannelVictorops",
  "fqn": "pulumi_sysdig.monitor",
  "classes": {
   "sysdig:Monitor/notificationChannelVictorops:NotificationChannelVictorops": "NotificationChannelVictorops"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Monitor/notificationChannelWebhook",
  "fqn": "pulumi_sysdig.monitor",
  "classes": {
   "sysdig:Monitor/notificationChannelWebhook:NotificationChannelWebhook": "NotificationChannelWebhook"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Monitor/team",
  "fqn": "pulumi_sysdig.monitor",
  "classes": {
   "sysdig:Monitor/team:Team": "Team"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Secure/benchmarkTask",
  "fqn": "pulumi_sysdig.secure",
  "classes": {
   "sysdig:Secure/benchmarkTask:BenchmarkTask": "BenchmarkTask"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Secure/cloudAccount",
  "fqn": "pulumi_sysdig.secure",
  "classes": {
   "sysdig:Secure/cloudAccount:CloudAccount": "CloudAccount"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Secure/list",
  "fqn": "pulumi_sysdig.secure",
  "classes": {
   "sysdig:Secure/list:List": "List"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Secure/macro",
  "fqn": "pulumi_sysdig.secure",
  "classes": {
   "sysdig:Secure/macro:Macro": "Macro"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Secure/notificationChannelEmail",
  "fqn": "pulumi_sysdig.secure",
  "classes": {
   "sysdig:Secure/notificationChannelEmail:NotificationChannelEmail": "NotificationChannelEmail"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Secure/notificationChannelOpsgenie",
  "fqn": "pulumi_sysdig.secure",
  "classes": {
   "sysdig:Secure/notificationChannelOpsgenie:NotificationChannelOpsgenie": "NotificationChannelOpsgenie"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Secure/notificationChannelPagerduty",
  "fqn": "pulumi_sysdig.secure",
  "classes": {
   "sysdig:Secure/notificationChannelPagerduty:NotificationChannelPagerduty": "NotificationChannelPagerduty"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Secure/notificationChannelSlack",
  "fqn": "pulumi_sysdig.secure",
  "classes": {
   "sysdig:Secure/notificationChannelSlack:NotificationChannelSlack": "NotificationChannelSlack"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Secure/notificationChannelSns",
  "fqn": "pulumi_sysdig.secure",
  "classes": {
   "sysdig:Secure/notificationChannelSns:NotificationChannelSns": "NotificationChannelSns"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Secure/notificationChannelVictorops",
  "fqn": "pulumi_sysdig.secure",
  "classes": {
   "sysdig:Secure/notificationChannelVictorops:NotificationChannelVictorops": "NotificationChannelVictorops"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Secure/notificationChannelWebhook",
  "fqn": "pulumi_sysdig.secure",
  "classes": {
   "sysdig:Secure/notificationChannelWebhook:NotificationChannelWebhook": "NotificationChannelWebhook"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Secure/policy",
  "fqn": "pulumi_sysdig.secure",
  "classes": {
   "sysdig:Secure/policy:Policy": "Policy"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Secure/ruleContainer",
  "fqn": "pulumi_sysdig.secure",
  "classes": {
   "sysdig:Secure/ruleContainer:RuleContainer": "RuleContainer"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Secure/ruleFalco",
  "fqn": "pulumi_sysdig.secure",
  "classes": {
   "sysdig:Secure/ruleFalco:RuleFalco": "RuleFalco"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Secure/ruleFilesystem",
  "fqn": "pulumi_sysdig.secure",
  "classes": {
   "sysdig:Secure/ruleFilesystem:RuleFilesystem": "RuleFilesystem"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Secure/ruleNetwork",
  "fqn": "pulumi_sysdig.secure",
  "classes": {
   "sysdig:Secure/ruleNetwork:RuleNetwork": "RuleNetwork"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Secure/ruleProcess",
  "fqn": "pulumi_sysdig.secure",
  "classes": {
   "sysdig:Secure/ruleProcess:RuleProcess": "RuleProcess"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Secure/ruleSyscall",
  "fqn": "pulumi_sysdig.secure",
  "classes": {
   "sysdig:Secure/ruleSyscall:RuleSyscall": "RuleSyscall"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Secure/team",
  "fqn": "pulumi_sysdig.secure",
  "classes": {
   "sysdig:Secure/team:Team": "Team"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Secure/vulnerabilityException",
  "fqn": "pulumi_sysdig.secure",
  "classes": {
   "sysdig:Secure/vulnerabilityException:VulnerabilityException": "VulnerabilityException"
  }
 },
 {
  "pkg": "sysdig",
  "mod": "Secure/vulnerabilityExceptionList",
  "fqn": "pulumi_sysdig.secure",
  "classes": {
   "sysdig:Secure/vulnerabilityExceptionList:VulnerabilityExceptionList": "VulnerabilityExceptionList"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "sysdig",
  "token": "pulumi:providers:sysdig",
  "fqn": "pulumi_sysdig",
  "class": "Provider"
 }
]
"""
)
