---
id: 0029
title: Remove Postgres database support
status: Approved
authors: Chris Doherty <chris.doherty4@gmail.com>, Jacob Weinstock <jakobweinstock@gmail.com>
---

## Summary

Tinkerbell supports 2 backends for managing persistent data: Postgres and Kubernetes. The introduction of the Kubernetes backend also introduced a new way for consumers to interact with Tinkerbell. Instead of interacting with gRPC APIs users interact with custom resources such as the Workflow or Template [CRDs](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/) using `kubectl`. The remaining gRPC APIs on the Tink Server exist to facilitate Tink Worker-Server interaction. Finally, [Rufio](https://github.com/tinkerbell/rufio) was developed as it better complimented the Kubernetes backend and deployment model.

The Kubernetes backend has had demonstrated success through the [EKS Anywhere project](https://github.com/aws/eks-anywhere).

## Proposal

We propose dropping support for the Postgres backend in favor of the Kubernetes backend. Users may continue to interact with workflows and templates by submitting them directly to the Kubernetes cluster.

Supporting a single backend makes maintenance of Tinkerbell considerably simpler and helps avoid divergence. While a typical software architecture should isolate persistent storage technology from business logic, the Kubernetes backend introduced a new way of managing workflows - namely through Kubernetes controllers. This creates complexity with respect to ensuring parity between the 2 backends as they are fundamentally different.

Supporting only the Kubernetes backend has the caveat that Tinkerbell requires a Kubernetes cluster to function. Setups using [KCP](https://github.com/kcp-dev/kcp) may be possible should a user wish to deploy with raw Docker/Docker Compose. However, the cost of maintenance likely outweighs the need for a Kubernetes cluster particularly when options such as K3D, KinD or Minikube are available for local cluster deployment.