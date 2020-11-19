---
id: 0019
title: PBnJ.next
status: approved
authors: Jacob Weinstock <jweinstock@equinix.com>
---

## Summary

PBnJ is a Restful API service that is tasked with interacting with Base Management Controllers (BMC). Among many other things, these BMC's can power a machine on and off and set the next boot device to PXE, BIOS, disk, etc. As part of a Tinkerbell workflow execution, the first step is

> On the first boot, the Worker is PXE booted

This is where PBnJ could be called to set the next boot device to PXE and then called again to reboot the machine. PBnJ currently support BMC's that work with the `ipmitool` or `racadm` tool. The following vendors are known to work: Supermicro, Dell, ASRockRack.

Currently, there is no mechanism in Tinkerbell to hook into this functionality (unfortunately, that mechanism is not what this proposal is about). The PBnJ service is only provided as a convenience for End-Users to hook into. While PBnJ does not currently have any integration into the Tinkerbell stack, it is very likely that it will be integrated into tink-server and be used to provide additional functionality to workflows in the future.

PBnJ's two main Restful API endpoints are `/power` (on, off, cycle, etc) and `/boot` (set next boot device to pxe, bios, disk, etc). `/power` is an asynchronous endpoint that returns a task id for use with the `/task` endpoint. `/boot` is a synchronous endpoint.

> An important note about BMCs. These devices are notoriously unreliable. Also, the IPMI protocol is the original protocol of these devices but not the only one available. There are other ways to interact with BMC's (depending on your vendor and BMC software); [redfish](https://www.dmtf.org/standards/redfish) (restful API), SSH, vendor specific Web APIs, gRPC, and others.

This doc proposes the addition of all new asynchronous gRPC endpoints as a step towards deprecating the Restful endpoints. The implementation behind these gRPC endpoints will add the ability to make use of multiple interaction types (noted above) instead of just `ipmitool` and `racadm`. This will enable support for new device vendors.  

## Goals and not Goals

Goals

* Add asynchronous gRPC endpoints to mirror existing functionality
* Add more ways to interact with a BMC behind the scenes
* Add basic user management capabilities
* Increase maintainability of code base

Non-Goals

* Deprecate existing Restful API
* Expose all BMC functionality through PBnJ

## Content

PBnJ.next will have all existing functionality available through gRPC
endpoints and would augment that functionality to be more robust and
support additional hardware vendors by implementing libraries like
[bmclib](https://github.com/bmc-toolbox/bmclib),
[gebn/bmc](https://github.com/gebn/bmc), and
[gofish](https://github.com/stmcginnis/gofish). As standardization across
hardware vendors is notoriously absent, an action against a BMC
would be tried using all of these libraries until one is successful.
The existing HTTP interface would be supported in the near term with eventual deprecation plans.

## Progress

There is current work in progress
[here](https://github.com/tinkerbell/pbnj/tree/pbnj.next).

## APIs

Proposed protocol buffers are available for review
[here](https://github.com/tinkerbell/pbnj/tree/pbnj.next/api/v1).
