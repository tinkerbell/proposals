---
id: 0009
title: Automated GitHub Action Runner
status: ideation 
authors: David McKay <david@rawkode.com>
---

## Summary:

The `tink` and `portal` repositories use GitHub Actions to run their automated tests on pull-request and push to master.

A dedicated GitHub Action runner is provided for the [Tinkerbell organization](https://github.com/tinkerbell), this runner was manually created and provisioned by [Gianluca](https://github.com/gianarb); this leaves the runner vulnerable; should it suffering any problems, restoring it will be a manual process.


## Goals:

- Automate the creation and provisioning of the dedicated runner

## Content: 

- Create a new repository, called `tinkerbell/test-infra`
- Using Pulumi, with TypeScript or Go, to manage the Packet server and provisioning with the tooling required to rebuild the runner when required

## Progress:

This will be updated as the proof of concept PR is opened on the newly created [test-infra repository](https://github.com/tinkerbell/test-infra).
