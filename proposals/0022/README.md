---
id: 0022
title: Re-imagining Boots
status: ideation
authors: Jacob Weinstock <jweinstock@equinix.com>
---

## Summary

Re-imagining boots. Boots does way too many things.
Here is a short list of just what I'm aware of.

- dhcp server
- pxe, tftp, http server
  - custom/bespoke kernel cmdline values: example: eclypsium
  - custom installations
    - custom ipxe
    - nixos
    - vmware
    - coreos
    - rancher
- syslog server
- hardware discovery ([/hardware-components](https://github.com/tinkerbell/boots/blob/70440b27cb1559770ef485596b9c3a4a253a4dfc/http.go#L82))
- phone home functionality ([/phone-home](https://github.com/tinkerbell/boots/blob/70440b27cb1559770ef485596b9c3a4a253a4dfc/http.go#L69))
- failure/event system ([/problem](https://github.com/tinkerbell/boots/blob/70440b27cb1559770ef485596b9c3a4a253a4dfc/http.go#L71))
- business rule engine for who and what should be allowed to PXE (https://github.com/tinkerbell/boots/blob/70440b27cb1559770ef485596b9c3a4a253a4dfc/http.go#L106)
  - metadata reader - talks to cacher or hegel or tink server

It it also full of many Equinix Metal only specific use-cases.
This proposal addresses the following

- redesign of existing boots functionality into smaller distinct and purposeful services
- removal or refactoring of Equinix Metal specific functionality
- add proxyDHCP support

## Goals and not Goals

Goals

- enable Tinkerbell stack to work in existing DHCP environments
- make running DHCP via boots optional
- re-architect existing boots functionality into smaller purposeful services
- enable Tinkerbell stack to work in existing DHCP environments
- split out proxyDHCP functionality

Non-Goals

- modifying architecture, code or APIs outside of the PXE phase

## Content

We can generalize the Tinkerbell machine provisioning lifecycle into 3 distinct phases.
The following describes the phases and how this proposed architecture fits into the PXE phase.

- Phase 1: **PXE**
  - **Goal**: PXE boot a machine into a selected operating system installation environment
- Phase 2: **Operating system installation environment boot**
  - **Goal**: Have the operating system installation environment ready to receive the go from tink server to install an operating system
- Phase 3: **Operating system installation**
  - **Goal**: Install the operating system and run actions defined in the tink workflow

See the architecture diagram [here](./tinkerbell-lifecycle.png).

The following are some advantages gained by this re-architecture.

- we will be able to integrate with existing DHCP servers
  - allows the use of both dynamic and static DHCP addresses
- DHCP via boots can be enable/disabled at runtime for users who need a DHCP server
- we can focus our efforts around 3 core areas that arguably differentiate the Tinkerbell stack
  - workflow building - tink server
  - installing operating systems (i.e. - running workflows) - tink-worker/[actions](https://docs.tinkerbell.org/actions/action-architecture/)
  - rules engine - determining what a machine should be given to pxe
- simpler and more maintainable code bases as they are more focused and singular in purpose
- simpler mental model for the provisioning lifecycle of a machine

The following functionality from existing boots will not exist in the PXE phase of the proposed architecture.
These are concerns of later phases so they will not be addressed.

- syslog server
- hardware discovery ([/hardware-components](https://github.com/tinkerbell/boots/blob/70440b27cb1559770ef485596b9c3a4a253a4dfc/http.go#L82))
- phone home functionality ([/phone-home](https://github.com/tinkerbell/boots/blob/70440b27cb1559770ef485596b9c3a4a253a4dfc/http.go#L69))
- failure/event system ([/problem](https://github.com/tinkerbell/boots/blob/70440b27cb1559770ef485596b9c3a4a253a4dfc/http.go#L71))

## Trade Offs

- more services add to the operational complexity/overhead to deploy and maintain
- time and effort to make the changes

## Progress

There is demo implementation available [here](https://github.com/jacobweinstock/tinkerbell-next) to show how this could work.

> Note - the demo/POC setup doesn't do anything more than boot into the operating system installation environments.
> The PXE phase.
> There will need to be more work done to get a full workflow run.

## System-context-diagram

Again, this proposal is for the `Phase: PXE`.
The following diagram illustrates where in the Tinkerbell machine provisioning lifecycle this proposal lives.

![machine provisioning lifecycle](./tinkerbell-lifecycle.png#3)

## Alternatives

One alternative is to overhaul the existing boots code base without breaking the service up.
We would still need to add proxyDHCP support and make rules more flexible.
We would still have a service that did way to many things, was a single point of failure, and was limited in its scalability and interoperability.
