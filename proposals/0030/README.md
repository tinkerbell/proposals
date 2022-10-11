---
id: 0030
title: Support for network booting in environments using trunked ports without a native VLAN
status: published
authors: Jacob Weinstock <jakobweinstock@gmail.com>
---

## Summary

Currently, the Tinkerbell stack does not support network booting in environments that use trunked ports without a native vlan.
This limits users to either modify their network configurations or not use the Tinkerbell stack.
The desire is for the Tinkerbell stack to not require potentially unacceptable network configurations to support network booting in environments configured with trunked ports without a native vlan.

## Terms

**VLAN:** Virtual local area network; logical identifier for isolating a network

**Trunk/Tagged:** A port enabled for VLAN tagging

**Access/Untagged:** A port that does not tag and only accepts a single VLAN

**Native VLAN:** A special VLAN whose traffic traverses on the trunk without any VLAN tag

VLAN-enabled ports are generally categorized in one of two ways, tagged or untagged. These may also be referred to as "trunk" or "access" respectively. The purpose of a tagged or "trunked" port is to pass traffic for multiple VLAN's, whereas an untagged or "access" port accepts traffic for only a single VLAN. -- [ref](https://documentation.meraki.com/General_Administration/Tools_and_Troubleshooting/Fundamentals_of_802.1Q_VLAN_Tagging)

## Goals and not Goals

Goals:

* Enable network booting in environments using trunked ports without a native vlan.

Non Goals:

* Custom workflow actions for configuring OS level VLAN tagging.
* Enabling of VLAN tagging via a machine's BMC with Rufio.

## Content

### Components

The following is an initial estimate of the changes that will be needed in order to enable the VLAN tagging functionality.

| Code Repo | Required Changes |
| --------- | ---------------- |
| ipxedust  | add `vcreate` capability to embedded iPXE script |
| boots     | 1. add VLAN ID to DHCP opt 43.116 2. add `vlanid=` to kernel command line in the auto.ipxe script |
| dhcp      | add VLAN ID to DHCP opt 43.116 |
| tink      | add VLAN ID to hardware spec for interfaces |
| hook      | add virtual interface with VLAN tagging capability based on `vlanid=` found in `/proc/cmdline` |

### User Experience

In order to enable VLAN tagging support from the Tinkerbell stack side, a user would only need to add the VLAN ID to the hardware spec for the interface. Generally, there will still be additional steps a user would need to take. Some outside of the Tinkerbell stack.

1. Tag interface(s) on switches for which the physical machine is connected.
2. Enable and set a VLAN ID in the out-of-band software for the physical machine.
3. Create a template workflow action to add VLAN tagging to an interface for the OS being installed.

All other processes around workflows, etc would be the same as they are currently.

### Failure modes

Failures to read the VLAN ID from DHCP option 43.116 in iPXE will not block iPXE from booting. The VLAN ID will just not be set and not propagated through to Hook.

### Tradeoffs

This feature adds complexity to the stack. Testing will most likely be manual and very dependent on the environment.

## System-context-diagram

![System-context-diagram](vlan_tag_feature.png)

## APIs

Example Hardware spec with added vlan id.

```yaml
apiVersion: "tinkerbell.org/v1alpha1"
kind: Hardware
metadata:
  name: sm01
  namespace: default
spec:
  disks:
    - device: /dev/nvme0n1
  metadata:
    facility:
      facility_code: onprem
    manufacturer:
      slug: supermicro
    instance:
      userdata: ""
      hostname: "sm01"
      id: "3c:ec:ef:4c:4f:54"
      operating_system:
        distro: "ubuntu"
        os_slug: "ubuntu_20_04"
        version: "20.04"
  interfaces:
    - dhcp:
        arch: x86_64
        hostname: sm01
        ip:
          address: 172.16.10.100
          gateway: 172.16.10.1
          netmask: 255.255.255.0
        lease_time: 86400
        mac: 3c:ec:ef:4c:4f:54
        name_servers:
          - 172.16.10.1
          - 10.1.1.11
        uefi: true
        vlan_id: "12"
      netboot:
        allowPXE: true
        allowWorkflow: true
```

## Alternatives

* Don't support VLAN tagging.
* Use [DHCP option 132](https://www.iana.org/assignments/bootp-dhcp-parameters/bootp-dhcp-parameters.xhtml) instead of DHCP option 43.116.
  * It's not clear whether Option 132 (IEEE 802.1Q VLAN ID) is used only for VOIP.
  * This proposes a very vendor (Tinkerbell) specific way of passing around a VLAN ID, so I opted for the vendor options area of DHCP via option 43.
  * Here are some examples of both option 132 and option 43 being used. Note, that all examples of option 132 are for VOIP use cases.
    * https://www.sonicwall.com/support/knowledge-base/how-to-configure-dhcp-option-132-to-get-dhcp-lease-from-vlan/220309054946100/
    * https://wiki.unify.com/wiki/VLAN_ID_Discovery_over_DHCP#Example_configuration_of_an_Cisco_Switch
    * https://www.teradici.com/web-help/TER1206003/4.x-4.9/07_HowTo/Config_VLAN.htm
    * https://wiki.freepbx.org/display/PHON/DHCP+VLAN+Option+132
