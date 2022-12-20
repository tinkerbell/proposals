---
id: 0031
title: Prevention of invalid Hardware configurations and mitigation of accidental application state overwrites
status: prediscussion
authors: Chris Doherty <chris.doherty4@gmail.com>
---

# Prevention of invalid Hardware configurations and mitigation of accidental application state overwrites

## Summary

The Tink repository defines several Custom Resource Definitions including the `Hardware` CRD.
`Hardware` contains both static data about the hardware and dynamic state updated by the Tinkerbell
stack. When provisioning Kubernetes clusters using Cluster API Provider Tinkerbell (CAPT), the 
`Hardware`  resources dynamic state is the source of truth that aids the system in determining  
whether a piece of hardware is provisioned and consequently in-use.

Operators of Tinkerbell submit additional hardware to the cluster using `kubectl`. When hardware
is submitted it is not gated by Kubernetes Admission Controllers. Consequently, it is possible
for operators to submit hardware with invalid data such as duplicate MAC address' or IPs. 
Additionally, it's possible for operators to accidentally overwrite dynamic state with a value
that indicates the hardware is not provisioned. When CAPT identifies hardware that should be
provisioned is not, it takes remedial action to bring the node online.

## Goals

This proposal seeks to prevent invalid hardware objects being submitted to a Kubernetes cluster
and mitigate the risk of human operators inadvertantly overwriting important hardware state.

## Proposal

Write a webhook that expects the user submitting a request to be a part of a configurable role.

## 