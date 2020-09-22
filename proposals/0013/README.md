---
id: 0013
title: Self Contained Workflows
status: ideation
authors: Manny <manny@packet.com>
---

## Summary

A method to specify http(s) based assets as part of the workflow.

## Goals and not Goals

Fully self-contained, versioned workflows.
This includes serving up of tink assets like osie initrd, kernel, modules, scripts.
We would also want to use this feature to serve static OS files locally with out having to manage a webserver's directories.
The code in boots tasked with dynamically creating/serving OS specific files (ipxe script, kickstart, Flatcar support file) would be replaced with this.

## Content

This proposal is acting mostly as the super-proposal tying in the two related proposals ([0014] and [0015]) where most of the work will take place.
If/once the related proposals are accepted and implemented the goals of this proposal will be met by the mere fact that self-contained-workflows is possible.

An example using my preferred syntax option of [0015]:

```yaml
version: "0.1"
name: ubuntu_provisioning
global_timeout: 6000
services:
  - name: serve boot files
    worker: {{provisioner1}}
    image: osie-boot-files-server:v1.0.42
    labels:
      - "traefik.http.routers.myrouter.rule=Host(`{{workflowid}}.workflows.tinkerbell.local`)"
      - "traefik.http.routers.myrouter.rule=Path(`/kernel`)"
      - "traefik.http.routers.myrouter.rule=Path(`/initrd`)"
      - "traefik.http.routers.myrouter.rule=Path(`/modules`)"
      - "traefik.http.routers.myrouter.rule=Path(`/auto.ipxe`)"

  - name: password receiver
    worker: "{{provisioner1}}"
    image: password-receiver:v1.0.42
    labels:
      - "traefik.http.routers.myrouter.rule=Host(`{{workflowid}}.workflows.tinkerbell.local`)"
      - "traefik.http.routers.myrouter.rule=Path(`/key`)"
      - "traefik.http.routers.myrouter.rule=Path(`/password`)"

tasks:
  - name: os-installation
    worker: {{.device_1}}
    volumes:
      - /dev:/dev
      - /dev/console:/dev/console
      - /lib/firmware:/lib/firmware:ro
    environment:
      MIRROR_HOST: <MIRROR_HOST_IP>
    actions:
      - name: disk-wipe
        image: disk-wipe:v1.0.42
        timeout: 90
      - name: disk-partition
        image: disk-partition:v1.0.42
        timeout: 600
        environment:
          MIRROR_HOST: <MIRROR_HOST_IP>
        volumes:
          - /statedir:/statedir
      - name: install-root-fs
        image: install-root-fs:v1.0.42
        timeout: 600
...
```

The `osie boot files` service will be running most likely along side Boots, Hegel and Tink-Server though not necessarily.
When `{{.device_1}}` powers on Boots will do its DHCP/TFTP dance.
When Boots is ready to set the filename url for the auto.ipxe script it will discover that there is an action ready to serve the auto.ipxe file (how exactly?).
Boots will set the `filename` dhcp option to the value discovered.
From this point on, the workflow can be fully self contained/versioned.
The `osie boot files` service will be able to write the auto.ipxe file so that the kernel files are fetched from itself.
The workflow will be able to use an image that will fetch the image from https://{{workflowid}}.workflows.tinkerbell.local, so on and so forth.

Questions:

1. How will Boots detect auto.ipxe is being served by a service container?
   1. Will it search for a label containing `auto.ipxe`?
   1. Or a special flag type label
2. If we go with 1.ii do we specify the url or expect that it will always match a pattern?
   1. boots.auto.ipxe=https://{{workflowid}}.workflows.tinkerbell.local/ipxe-script.ipxe
   1. boots.auto.ipxe (boots will force the filename=https://{{workflowid}}.workflows.tinkerbell.local/auto.ipxe)

   Option 2 seems easier to get wrong, for example by not setting the correct proxy path label.

## Alternatives

#### Status Quo

Keep the [installers](https://github.com/tinkerbell/boots/tree/master/installers) in place, add more as OS that need them come up.

This isn't really an option.

Boot's code is already pretty convoluted with layering violations providing things for networkd and similar.
Our VMWare ESXi setup isn't really ideal (afaik) and to make it better we'd need to add even more special file serving to boots.
Boots has to know way too many details of both the underlying hardware and the intended OS to get the job done.

#### Just Version Static Assets And Fetch Them

We have already tried/half-hacked attempts at similar functionality in Packet as this proposal allows.
Custom OSIE Versions is used to specify an alternate OSIE (not /current) to allow for real-world testing mostly.
This only works for OSIE though.
VMWare, NixOS, and all of the other OSes Packet supports can only really be updated en-masse/live-to-production with our current setup.

This also means boots will still have code to generate dynamic bits.
This is already hard to maintain and will only get worse.

#### Run More Services

Just run services that know how to serve/generate these files with your own orchestrator.

This is ok, but then we'd have to still deal with backwards compatibility, upgrade for new versions, figure out a way to experiment....
All this would be for free if we had fully self contained workflows.

[0014]: https://github.com/tinkerbell/proposals/pull/19
[0015]: https://github.com/tinkerbell/proposals/pull/20
