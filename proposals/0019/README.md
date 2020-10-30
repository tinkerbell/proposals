---
id: 0019
title: PBnJ.next
status: ideation
authors: Jacob Weinstock <jweinstock@equinix.com>
---

## Summary

PBnJ is in need of some updating to increase the maintainability,
robustness and extensibility of the service. Under the hood PBnJ
currently relies on `ipmitool` and `racadm`. This limits its hardware
agnostic functionality. Scaling the service is also a concern at the
moment as there is an in-memory task service for some endpoints.

This proposal is intended to provide an option for how to increase the
maintainability, robustness and extensibility of the service.  

## Goals and not Goals

Goals

* Add more ways to interact with a BMC.
* Add basic user management capabilities.
* Add a gRPC interface.
* Asynchronous operations.
* Increase maintainability through updated code base.
* Architecture that scales well.
* Serve both http and gRPC protocols.
* Support Tinkerbell centralize logging and events

Non-Goals

* Provide all BMC functionality.

## Content

PBnJ.next will have all existing functionality available through a gRPC
interface and would augment that functionality to be more robust and
support additional hardware vendors by implementing libraries like
[bmclib](https://github.com/bmc-toolbox/bmclib),
[gebn/bmc](https://github.com/gebn/bmc), and
[gofish](https://github.com/stmcginnis/gofish). As standardization across
hardware vendors is notoriously absent, an action against a BMC
would be tried using all of these libraries until one is successful.
The existing Http interface would be supported in the near term with eventual deprecation plans.

## Progress

There is current work in progress
[here](https://github.com/tinkerbell/pbnj/tree/pbnj.next).

## System-context-diagram

[example power on request flow/code architecture](./PBnJ-RequestFlow.png)

## APIs

Proposed protocol buffers are available for review
[here](https://github.com/tinkerbell/pbnj/tree/pbnj.next/api/v1).

## Alternatives

Keep the existing service as is and bolt on more functionality.
This would still require a significant rewrite to make scalable
and add new functionality.
