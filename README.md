# Sysdig Pulumi Provider

This repository hosts the [Pulumi][pulumi] provider for [sysdig][sysdig]. It is a
wrapper around the [Terraform][terraform] [Sysdig provider][tr_sysdig].

Pulumi allows to manage infrastructure as code, leveraging many different possible
languages to manage the resources. It is possible to declare infrastructure as code
in the following languages:

    - Golang
    - Python
    - JavaScript
    - TypeScript
    - .NET

## Use with Pulumi

> :warning: **This provider is not currently published** in the registries of the
> various languages.

### Python

```bash
pip install --extra-index-url https://pypi.degenerazione.xyz pulumi_sysdig
```

### Golang

```bash
go get github.com/Sysdig-Hackathon-Picasso/pulumi-sysdig/sdk/go/...
```
### Node.js

> :hourglass_flowing_sand: This will be published to npm, to be able to use `npm` or
> `yarn`

### .NET

> :hourglass_flowing_sand: This will be published to nuget, to be able to use `nuget`

## Configuration

The following configuration points are available for the `foo` provider:

 - `sysdig:sysdig_monitor_insecure_tls` (boolean)
 - `sysdig:sysdig_monitor_api_token` (secret)
 - `sysdig:sysdig_monitor_url` (string)
 - `sysdig:sysdig_secure_api_token` (secret)
 - `sysdig:sysdig_secure_insecure_tls` (boolean)
 - `sysdig:sysdig_secure_url` (string)

## Develop

### Prerequisites

Ensure the following tools are installed and present in your `$PATH`:

* [`pulumictl`](https://github.com/pulumi/pulumictl#installation)
* [Go 1.16](https://golang.org/dl/) or 1.latest
* [NodeJS](https://nodejs.org/en/) 14.x.  We recommend using [nvm](https://github.com/nvm-sh/nvm) to manage NodeJS installations.
* [Yarn](https://yarnpkg.com/)
* [TypeScript](https://www.typescriptlang.org/)
* [Python](https://www.python.org/downloads/) (called as `python3`).  For recent versions of MacOS, the system-installed version is fine.
* [.NET](https://dotnet.microsoft.com/download)

### Build

To create the `.tar.gz`s for the release:

```bash
make dist
```

This will generate:

* `dist/sysdig-*.tar.gz` - pulumi plugin
* `sdk/python/bin/dist/pulumi_sysdig-*.tar.gz` - sysdig python sdk
* `nuget/Pulumi.Sysdig.*.nupkg` - .NET sdk package

[pulumi]: https://www.pulumi.com
[sysdig]: https://sysdig.com
[terraform]: https://www.terraform.io
[tr_sysdig]: https://github.com/sysdiglabs/terraform-provider-sysdig
