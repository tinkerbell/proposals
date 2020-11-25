---
id: 0016
title: Propose Default workflow
status: ideation
authors: Dan <daniel.finneran@gmail.com>
---

## Summary

This additional feature would allow an end user to specify a workflow to be used against a machine that has no `hardware` profile specified within Tinkerbell. The main feature that this would enable would be for DCIM or hardware discovery where any new unknown machines are automatically prossed by a specified workflow that could perform actions such as:

 - Secure wipe
 - Forensic analysis
 - DataCentre inventory
 - Hardware capture


This could ultimately lead to the capability to build pools of hardware (grouped by a hardware profile (GPUs/SSDs/CPU type)) that we could start to provide "cloud like" functionality around.

## Goals

1. Specify a default workflow
2. If Tinkerbell doesn't "know" the hardware then assign that workflow

## Not Goals

N/A

## Content

The same functionality was implemented in `plunder` in order to build a pool of discovered hardware, which the `cluster-api` provider could allocate unused hardware when provisioning was required. 

## Concerns

Default behaviour is fine when we're looking at a discovery (READ-ONLY) workflows, a workflow that is potentially destructive could be dangerous when enabled in an unknown environment.
