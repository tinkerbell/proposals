---
id: 0028
title: Make hardware data available to Kubernetes workflows
status: published
authors: Srikar Ganti <srikarsganti@gmail.com>, Chris Doherty <chris.doherty4@gmail.com>
---

## Context

Cluster API Provider Tinkerbell is used to provision Kubernetes clusters using a Tinkerbell stack. It watches for the TinkerbellMachine custom resource and generates a Workflow and Template custom resource that are used by the Workflow controller for rendering Tink Worker actions. The Workflow Controller currently supports a string-string map used as data for rendering the template.

## Problem

CAPI enforces a constraint where it uses a single template for all nodes within a group. For example, TinkerbellMachine’s for the control plane are all created using the same template. Often, the Tinkerbell actions defined in a template action override require machine specific information, such as disk, that is unavailable when the workflow is rendered. This forces consumers to specify a single value or leverage an out-of-band mechanism for creating an actions template.

## Proposal

We propose adding a hardwareRef field to the [Workflow CRD](https://github.com/tinkerbell/tink/blob/main/pkg/apis/core/v1alpha1/workflow_types.go#L18). CAPT [will be responsible for populating the hardwareRef](https://github.com/tinkerbell/cluster-api-provider-tinkerbell/blob/main/controllers/machine.go#L661) as it acquires the hardware to be used. The Workflow Controller will use the hardwareRef to retrieve the Hardware, serialize it and make it available to the template during rendering. The hardware data will be namespaced under a .hardware key. The existing `hardwareMap` will remain untouched and defensive code will be added to ensure the hardwareMap does not specify a hardware key that could overwrite the .hardware.


## Example template

This example illustrates the use of disks retrieved from the hardware type used in a template.

```yaml
tasks:
  - name: foo
    image: write2disk
    environment: 
      DEST_DISK: {{ .hardware.disks[0] }}  
```

```yaml
apiVersion: "tinkerbell.org/v1alpha1"
kind: Workflow
metadata:
  name: wf1
  namespace: default
spec:
  templateRef: debian
  hardwareRef:
  	name: hardware-1
  	namespace: hardware-namespace
  hardwareMap:
    device_1: 3c:ec:ef:4c:4f:54
```

## Cons

- The `Hardware` type is similar to a POJO and represents a storage structure. By exposing it directly to templates we create a data coupling between templates and the data storage structure making it difficult to evolve the storage structure. Given the consumer owns the template and Tinkerbell owns the storage structure its likely we'll want to evolve the storage format that may result in breaking template definitions. 
- Tinkerbell currently supports actions that run on different Tink Workers. A single `hardwareRef` doesn't cover workflow use cases that leverage multiple Tink Workers. We could address this in the future but it it would need to take a wider view of Tinkerbell components and ensure the commonly observed `device_x` key in the `hardwareMap` isn't integral to identifying hardware resources.

