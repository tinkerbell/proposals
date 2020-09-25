---
id: 0014
title: Workflow Services
status: published
authors: Manny <manny@packet.com>
---

## Summary

A Workflow Service is part of a Workflow, just like Tasks.
It is a container that starts before any Task and runs in the background.
Services are *not* required to continue running for the entire duration of the workflow.
Once the Workflow is Done (Finished/Errored/Timedout...) it will be killed/cleaned up.

Support for running workflow actions in the background.
This is to support [0013] in conjuction with [0015].

## Goals and not Goals

Run containers in detached/background so they may provides "services" to the further workflow actions/tasks.

## Content

By having containers running in the background we remove "global" code from boots and have it local to the workflow instead.
The same benefits could be had for serving of the installer/image files.
This is the initial/immediate use case, I would not be surprised to see others come up.

I initially thought about making this an attribute of the action.
That seemed a little bit too loose and easy to make bad choices with so changed it to a separate kind of "job" for the task.
@nathangoulding suggested a top-level attribute instead as it most likely makes sense to run this for the full life time of a workflow.
This makes sense to me.
It seems like the service should not out live its parent, so we would want to terminate services at the end of a task if we were still under a task.

#### Lifecycle

The services will be started before any Tasks.
All services will be required to be `Up` and `healthy` (if HEALTHCHECK is setup in the image) before any tasks are started.
Services are *not* required to continue running for the entire duration of the workflow.
Attempts to restart the service may occur if a service exits with an error.
A service being down will only affect the Workflow state if a Task/Action itself fails due to the missing service.
Once the Workflow is no longer Active (Errored/Timedout/Finished) all services will be terminated.

The Tink Event Bus will be needed in order to make use of Task/Action results.

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

## Alternatives

Keep code in boots, add more code to boots for more OSes.
Figure out versioning of files in an unmanaged http server.

[0013]: https://github.com/tinkerbell/proposals/pull/18
[0015]: https://github.com/tinkerbell/proposals/pull/20
