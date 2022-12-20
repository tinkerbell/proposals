---
id: 0031
title: Prevention of invalid Hardware configurations and mitigation of accidental application state overwrites
status: prediscussion
authors: Chris Doherty <chris.doherty4@gmail.com>, Aravind Ramalingam
---

# Prevention of invalid Hardware configurations and mitigation of accidental application state overwrites

## Summary

The Kubernetes backend of Tinkerbell uses Custom Resource Defitions (CRDs) to persist application
state. 
The Tink repository defines the core CRDs including the `Hardware` type.
`Hardware` contains both static data about the hardware, and application state updated by the 
Tinkerbell stack.
Operators of Tinkerbell submit additional hardware to the cluster using `kubectl`. 
When hardware is submitted it is not gated by any validation. 
This makes it possible for operators to submit hardware with invalid data such as duplicate MAC 
address' or IPs and, more importantly, inadvertent modifications to application state.

When provisioning Kubernetes clusters using Cluster API Provider Tinkerbell (CAPT), the 
`Hardware` resources is the source of truth that aids the system in determining whether a piece of 
hardware is provisioned and subsequently in-use.
If an operator submits the same hardware resource with `kubectl` and overwrites the application
state (likely inadvertently) CAPT will identify the hardware as unprovisioned and begin remedial
action to bring the node to a provisioned state.
This has consequences for cluster workloads, and ultimately end users, as they are abruptly 
terminated.

## Goals

This proposal seeks to prevent invalid hardware objects from being submitted to the Kubernetes 
cluster responsible for storing Tinkerbell custom resources subsequently mitigating the risk of 
human operators inadvertantly overwriting important application state.

## Proposal

Kubernetes defines the concept of [admission control][admissionControl] via built-in admission
controllers. 
By default, Kubernetes enables 2 admission controllers that can be used to extend cluster admission 
control:  `MutatingAdmissionWebhook` and `ValidatingAdmissionWebhook`.
Both admission controllers call out to a configurable set of webhooks that are responsible for  
deciding if a request to create, update or delete a resource should be submitted to the cluster. 
Mutating webhooks are called first and can modify the contents of an object while validating
webhooks check object contents only (see the [reinvocation policy][reinvocationPolicy] for  
additional details on admission behavior).

We will introduce both mutating and validating webhooks that defaults the `Hardware` resources 
application state on _create_ operations.
This removes the responsibility of defaulting application state from operators when submitting new 
resources and ensures we have an opportunity to check for duplicate MACs and IP configurations.

To perform duplicate data checks in the validating webhook it is necessary to retrieve all existing 
`Hardware` from the  cluster. 
Using the [custom admission webhook][customAdmissionWebhook] pattern we can gain access to a client
to retrieve `Hardware` from the cluster.

### Restricting mutation of application state fields

In addition to the basic validations performed by the validating webhook, we will add logic to the 
validating webhook that ensures only entites with permission are modifying application state fields 
on update.
For each request to be validated the webhook receives an `AdmissionReview` (Appendix A).
The `AdmissionReview` contains a `request.userInfo.groups` field that includes all groups

## Related Work Items

### Separation of static data from application state

The community has discussed at various points the sepration of static hardware data from 
application state currently represented on the `Hardware` resource. 
This could be achieved with 2 distinct CRDs with 1 being owned by human operators and the other
by the system. 
This proposal seeks to compliment that direction as we will still require mechanisms to restrict
what subjects can mutate application state, however the risk is reduced given human operators
need a more intentional action to mutate application state.

## Appendix

### A. `AdmissionReview` Object

```yaml
apiVersion: admission.k8s.io/v1
kind: AdmissionReview
request:
  # Random uid uniquely identifying this admission call
  uid: 705ab4f5-6393-11e8-b7cc-42010a800002

  # Fully-qualified group/version/kind of the incoming object
  kind:
    group: autoscaling
    version: v1
    kind: Scale

  # Fully-qualified group/version/kind of the resource being modified
  resource:
    group: apps
    version: v1
    resource: deployments

  # subresource, if the request is to a subresource
  subResource: scale

  # Fully-qualified group/version/kind of the incoming object in the original request to the API server.
  # This only differs from `kind` if the webhook specified `matchPolicy: Equivalent` and the
  # original request to the API server was converted to a version the webhook registered for.
  requestKind:
    group: autoscaling
    version: v1
    kind: Scale

  # Fully-qualified group/version/kind of the resource being modified in the original request to the API server.
  # This only differs from `resource` if the webhook specified `matchPolicy: Equivalent` and the
  # original request to the API server was converted to a version the webhook registered for.
  requestResource:
    group: apps
    version: v1
    resource: deployments

  # subresource, if the request is to a subresource
  # This only differs from `subResource` if the webhook specified `matchPolicy: Equivalent` and the
  # original request to the API server was converted to a version the webhook registered for.
  requestSubResource: scale

  # Name of the resource being modified
  name: my-deployment

  # Namespace of the resource being modified, if the resource is namespaced (or is a Namespace object)
  namespace: my-namespace

  # operation can be CREATE, UPDATE, DELETE, or CONNECT
  operation: UPDATE

  userInfo:
    # Username of the authenticated user making the request to the API server
    username: admin

    # UID of the authenticated user making the request to the API server
    uid: 014fbff9a07c

    # Group memberships of the authenticated user making the request to the API server
    groups:
      - system:authenticated
      - my-admin-group
    # Arbitrary extra info associated with the user making the request to the API server.
    # This is populated by the API server authentication layer and should be included
    # if any SubjectAccessReview checks are performed by the webhook.
    extra:
      some-key:
        - some-value1
        - some-value2

  # object is the new object being admitted.
  # It is null for DELETE operations.
  object:
    apiVersion: autoscaling/v1
    kind: Scale

  # oldObject is the existing object.
  # It is null for CREATE and CONNECT operations.
  oldObject:
    apiVersion: autoscaling/v1
    kind: Scale

  # options contains the options for the operation being admitted, like meta.k8s.io/v1 CreateOptions, UpdateOptions, or DeleteOptions.
  # It is null for CONNECT operations.
  options:
    apiVersion: meta.k8s.io/v1
    kind: UpdateOptions

  # dryRun indicates the API request is running in dry run mode and will not be persisted.
  # Webhooks with side effects should avoid actuating those side effects when dryRun is true.
  # See http://k8s.io/docs/reference/using-api/api-concepts/#make-a-dry-run-request for more details.
  dryRun: False
```

[admissionControl]: https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/
[customAdmissionWebhook]: https://book.kubebuilder.io/reference/webhook-for-core-types.html
[reinvocationPolicy]: https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/#reinvocation-policy