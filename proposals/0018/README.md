---
id: 0018
title: Add Instance Object Definition
status: abandoned
authors: Manny <mmendez@equinix.com>
---

## Summary

The current hardware data model does not support any OS installation details.
Installing an OS is one of the main focus areas of the Tinkerbell project at the moment.
This exists in the legacy Packet Platform under "instance" in the hardware object and a similar concept is in [metal3].

This would be used by Boots and Hegel (at least) to serve their needs.

## Goals and not Goals

We have a goal of keeping the Hardware data relatively static.
We need a locaion to place instance/provision/installation-info that changes from provision to provision.
It is not a goal to stick this info in `elastic_data`.

## Content

This proposal is about adding a storage concept for data that is not as permanent as Hardware, and that can persist across multiple Workloads.
This would mimic the EquinixMetal life-cycle of a machine, provisioned and thus occupied until the customer deletes and the machine is deprovisioned.
This life-cycle is also useful in non EM use cases as it allows a worflow to end and a future workflow to act upon the previous os intallation state.

#### Hegel

Hegel currently serves up EC2 style metadata by transforming the JSON blob in hardware.metadata to the compatible form.
The JSON blob in Hardware.metadata is not specified in order to be flexible to non EquinixMetal operators, yet Hegel's transformation logic is fixed.
So in reality this is just an undocumented JSON blob that is actually very useful to the project overall.

#### Boots

Boots is currently dynamically creating the iPXE script from data in Hardware.
We generally want to keep the Hardware data static.
But we would end up needing to modifying it to add ephemeral info, such as IPs (public/private v4 and public v6 as EquinixMetal does today), hostname, os installer to use...

## APIs

The API consists of defining a minimal and extensible Instance object as a field in Hardware, via Protobufs.
We already have a protobuf file in tree ([packet]) that defines the non-tinkerbell Instance JSON object as Protobuf.
This will be used as the starting point, moving fields that are unnecessary for Tinkerbell/Hegel/Boots into a meta field.
If interest builds we can hoist meta values into the Instance protobuf definition.

```protobuf
syntax = "proto3";

option go_package = "packet";

package github.com.tinkerbell.tink.protos.packet;

message Instance {
  string id = 1;
  string state = 2;
  string hostname = 3;

  // fork in the road! choose a path
  // path 1Left
  message CustomIPXE {
    string url = 1;
    bool always_pxe = 2;
  }
  message OperatingSystem {
    string distro = 1;
    string version = 2;
    string tag = 3; // snapshot? revision? ...?
  }
  oneof OS {
    CustomIPXE custom_ipxe = 4;
    OperatingSystem operating_system = 5;
  }
  // path 1Center
  message OperatingSystem {
    string distro = 1;
    string version = 2;
    string tag = 3; // snapshot? revision? ...?
    bytes extra = 4;
  }
  // path 1Right
  message OperatingSystem {
    string distro = 1;
    string version = 2;
    string tag = 3; // snapshot? revision? ...?
    string ipxe_url = 4;
    bool always_pxe = 5;
  }
  OperatingSystem operating_system = 4;

  message IP {
    string address = 1;
    uint32 cidr = 2;
    string gateway = 3;

    enum Family {
      UNKNOWN = 0;
      V4 = 1;
      V6 = 2;
    }
    Family family = 4;

    // fork in the road! choose a path
    // path 2Left
    enum Flags {
      NONE = 0;
      PUBLIC = 1;
      MANAGEMENT = 2;
    };
    uint32 flags = 5;  // bitwise-or of Flags

    // vs
    //path 2Right
    bool public = 5;
    bool management = 6;

  }
  repeated IP ips = 6;

  string userdata = 7;
  string crypted_root_password = 8;

  repeated string tags = 9;
  repeated string ssh_keys = 10;
}
```

### Forks in the Roads

#### Fork 1: OSes

Custom iPXE is not just another OS.
It has a very fundamental difference from anything else we might run with Workflows, the iPXE script (or its URL) is an input parameter.
Any other workflow would always want to boot into OSIE in workflow mode or otherwise be known statically in the Workflow Template.
But we just don't have a need to specify an iPXE script as part of the input for non-CustomIPXE OSes.

EquinixMetal supports Custom iPXE as the "most generic do what you want" option and so we need to be able to handle it in Tinkebell.
Custom Workflows (when supported by EM) should diminish the use of Custom iPXE, but I don't see it ever fully going away.
Its also a great option to just boot into [netboot.xyz] to debug/explore some new hardware and so still useful for Tinkerbell itself.

How to handle the 2 different types of OSes?
Explicitly like in 1Left?
Explicit and well defined is better than implicit or weakly defined, imo.
Path 1Right just feels wrong, but is how we do things to day in EquinixMetal.

Path 1Center is a bit of a mix, with benefits and warts from both sides.
The common/well specified options are well specified, but leaves the generic bits out for interpretation.

I lean towards trying out the 1Left path.

#### Fork 2: To bitflags or not bitflags?

Should we do bitflags in Instance.OperatingSystem.IP.flags (Path 2A) or break them out to bool `public` and `management` like currently done in [protos/packet][packet] (Path 2B)?

I don't particularly like "exploding" flags into fields, but that could just be my once-upon-a-time embedded developer thinking leaking.

I think another path/option is something like:
```protobuf
    enum Flags {
      NONE = 0;
      PUBLIC = 1;
      MANAGEMENT = 2;
    };
    repeated Flags flags = 5;
```
Where the code will just accumulate the flags seen.

This is a bit more formal of a definition of bitflags than "enums + int" but seems weird to me, I can see this being used though.
This path doesn't protect against repeated dupes, but I doubt that would be a big deal.

#### Extensions?

I can already think of a couple of extensions that we'd need for EquinixMetal:
* Licensing fields for some OSes
* Spot Market related fields
* IQN, used for block storage

And the 2 ways I can think of to handle the extensions:

1. A `bytes` field for embedded messages (JSON or encodded proto3).
   If json, we can merge the fields with Instance in Hegel for example.
   Hegel basically uses option 1 already in `Hardware.metadata` as a stringified json blob, its pretty inconvenient.
   I think the inconvenience mostly comes from having to grab everything from there instead of using the json blob to extend the common fields available in Instance.

2. Use an `Any extra` field.
   This seems nice from a serialization stand point, but would only really help for external services.
   For Hegel to make use of this we'd want to specify the protobuf fields being encoded and then there isn't really a big benefit from splitting out from Instance.

[metal3]: https://github.com/metal3-io/baremetal-operator/blob/master/docs/api.md#provisioning
[netboot.xyz]: https://netboot.xyz/
[packet]: https://github.com/tinkerbell/tink/tree/f5cdb83338d6961fb7c4c940918892b639126d0a/protos/packet
