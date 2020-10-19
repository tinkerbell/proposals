---
id: 0017
title: Official Terraform provider for Tinkerbell
status: discussion
authors: Mateusz Gozdkek <mateusz@kinvolk.io>
---

## Summary

To support more use cases for Tinkerbell, for example to [add support of Tinkerbell platform]
to [Lokomotive], Tinkerbell should provide Terraform provider for it's API.

## Goals and not Goals

To have Terraform provider in Tinkerbell published in [Terraform registry].

## Content

Some time ago I created a PoC of Terraform provider for Tinkerbell, which currently lives [in kinvolk organization] on Github.

I think before this provider gets published to the registry, it should be moved to [Tinkerbell] organization.

This project already follows best practices from https://github.com/packethost/standards. It also uses GitHub actions for CI process.

### Requirements

In order to publish this provider to the Terraform registry, following requirements must be met:
- https://github.com/tinkerbell/terraform-provider-tinkerbell repository should be created on GitHub.
- User [invidian] should be added as a maintainer to this repository.
- We must decide, what GPG key should be used to sign Terraform provider releases, as this GPG key
  needs to be sent to Hashicorp support to be added as valid for `tinkerbell` organization.

## Alternatives

Alternatively, provider could be published using my personal account and stay as unofficial provider.

[add support of Tinkerbell platform]: https://github.com/kinvolk/lokomotive/issues/382
[Lokomotive]: https://github.com/kinvolk/lokomotive
[Terraform registry]: https://registry.terraform.io
[in Kinvolk organization]: https://github.com/kinvolk/terraform-provider-tinkerbell
[Tinkerbell]: https://github.com/tinkerbell
[invidian]: https://github.com/invidian
