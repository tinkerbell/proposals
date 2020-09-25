---
id: 0009
title: Automated GitHub Action Runner
status: ideation
authors: David McKay <david@rawkode.com>
---

## Summary:

The `tink` and `portal` repositories use GitHub Actions to run their automated tests on pull-request and push to master.

A dedicated GitHub Action runner is provided for the [Tinkerbell organization](https://github.com/tinkerbell), this runner was manually created and provisioned by [Gianluca](https://github.com/gianarb). As this runner wasn't automated from the beginning, should anything happen to the server that requires it to be shutdown or migrated, it would prove timely and cumbersome to rebuild.

This proposal is to use a best practice approach, Infrastructure as Code, to deliver the test infrastructure needed to support the Tinkerbell project. Doing so also provides the following benefits:

- Complete transparency through open sourcing the automation to provision the runner
- Anyone can contribute and help us maintain and improve the runner
- Anyone can contribute runners using other cloud providers
- Anyone can build and operate their own runner

## Goals:

- Automate the creation and provisioning of a dedicated GitHub Actions runner

## Content: 

### Open Source

We will keep all the automation open-source, in a newly created repository: `tinkerbell/test-infra`.

### Pulumi

Using [Pulumi](https://pulumi.com) and its [Go SDK](https://www.pulumi.com/docs/intro/languages/go/), we will provide the automation that is capable of creating a server on Packet Cloud; which will be provisioned via user-data.

The structure of this code will ensure that the server creation and provisioning are handled separately, allowing for alternative cloud providers to be added/replaced at a later date; as decided by the Tinkerbell community.

#### Operating System

The runner will be provisioned on an Ubuntu 20.04 server.

#### Security

The runner will be configured with a catch-all `iptables` configuration that blocks all in-bound traffic. This is possible as the communication method of the GitHub Runner is a one-way poll.

SSH access, if required, MUST go over the Serial over SSH connection and require an approved SSH key AND root password.

You can read more about this [here](https://docs.github.com/en/actions/hosting-your-own-runners/about-self-hosted-runners#communication-between-self-hosted-runners-and-github)

All outbound traffic will also be restricted to the following domains.

- github.com
- api.github.com
- *.actions.githubusercontent.com
- <what else do we need for tinkerbell to build / run / test ?>

We **will not** enable any other IPs, including Ubuntu update servers; we will rely on the automation to deprovision and reprovision with the latest updates on a regular cadence of 7 days.

#### Fork Control

We will need to ensure that our runner *ONLY* picks up jobs after a maintainer, or trusted contributed, specifically allows it with `/ok-to-test`.

This allows for the team to inspect the pull-request for malicious code that could somehow exploit the compute resources available through the runner.

#### Secret Management

Pulumi has built-in support for secret management that we will explicitly reject from being used. Instead, we will rely on GitHub Secrets to inject any API keys required for the automation to complete. This allows the permission model for modifying / changing the provisioning of the runner to be handled at the GitHub org / repository levels.

## Progress:

This will be updated as the proof of concept PR is opened on the newly created [test-infra repository](https://github.com/tinkerbell/test-infra).
