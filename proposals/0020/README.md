---
id: 0020
title: Template Versioning
status: ideation
authors: Aman Parauliya <aman@infracloud.io>
---

## Summary

As per the proposal [0013], we are planning to induce workflow service as a part of the template which is used to create a workflow.

As we add new features, we may have to modify the existing template format. Therefore, from a future perspective, we should support template versioning so that we can support existing template formats.

## Goals and not Goals

### Goals

Provide template versioning so that current format or schema of the template should be compatible as well as with the new one.

### Not Goals

Current template schema should be depricated.

## Content

This feature will make the life of existing user a bit easier because of the following points:

- User doesn't need to migrate to new template schema since we have the support of existing templates.
- User will also have the flexibility of using the kind of template he/she would like to use or in another words whichever template schema is fit for his/her usecase.

## Lifecycle

With each of the release we can have a different version of template if template schema has been modified in that particular release. And based on the versions we can have different structure/functions in our code to work with different formats/schemas of a template which is a very important because actual workflow lies withtin the template itself.

Following is the example of two different template schemas:

### 1. Current template

```yaml
version: "0.1"
name: ubuntu_provisioning
global_timeout: 6000
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

Structure in which the above template will be parsed should look like this:

```go
type Workflow struct {
        Version       string `yaml:"version"`
        Name          string `yaml:"name"`
        ID            string `yaml:"id"`
        GlobalTimeout int    `yaml:"global_timeout"`
        Tasks         []Task `yaml:"tasks"`
}

```

### 2. Template after proposal [0013]

When proposal [0013] is implemented a new field in the template called as `service` will be introduced and the template will look like as following:

```yaml
version: "0.2"
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

Since there is a new field `services` inserted in the above template, There will be a new field needs to be added in the structre to parse the above template which will be as follows:

```go
type Workflow struct {
        Version       string `yaml:"version"`
        Name          string `yaml:"name"`
        ID            string `yaml:"id"`
        GlobalTimeout int    `yaml:"global_timeout"`
        Services      []Service `yaml:"services"`
        Tasks         []Task `yaml:"tasks"`
}

```

Now if we would like to support both of the above templates then we need to identify the struct in which we need to parse the template which we can identify based on the versions.

## Alternatives

- We can have the support only for latest template schema and ask user to migrate to the latest format of template schema each time we do a modification in the template schema while providing a stable release. In that case, there will not have the support for the existing schema of template which can affect to the existing users.
- Write a utility that can read the old templates and transform them as per the new schema.


[0013]: https://github.com/tinkerbell/proposals/pull/18
