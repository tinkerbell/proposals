---
id: 0008
title: Explicit Certification For Components
status: published
authors: Jason DeTiberus <jdetiberus@packet.com>
---

## Summary

Currently certificate configuration for Tinkerbell is handled through a
combination of the `TINKERBELL_TLS_CERT` env variable and file based lookup.
The file based lookup will use the `TINKERBELL_CERTS_DIR` env variable if
defined, defaulting to `/certs/<facility>` if undefined. The certificate files
themselves are hardcoded to `bundle.pem`, which is expected to contain both the
CA certificate and the serving certificate, and `sserver-key.pem` for the
private key.

The current configuration methods do not offer a lot of flexibility for users,
especially if they are looking to run Tinkerbell in Kubernetes and leverage
automated certificate management tooling, such as
[cert-manager](https://cert-manager.io/), for automated certificate generation
and rotation.

This proposal is intended to provide generalized guidance on how Tinkerbell
components should expose certificate configuration to users.

## Goals and not Goals

Goals:

1. Provide common methods of exposing certificate configuration for Tinkerbell
services

Non-Goals:

1. Add tls support anywhere it is not currently enabled

## Content

Rather than using a combination of env variables and hard-coded defaults for
the certificate file lookups, we can leverage explictly configured command line
arguments for specifying the certificate path, and default certificate file
names to use.

### Expected configuration support for Tinkerbell components

- Provide an optional command line argument for specifying a PEM encoded CA certificate
  - `--ca-cert`
  - If not provided, it will treat the provided certificate as a CA/cert bundle
- Provide a command line argument for specifying a PEM encoded TLS certificate
  - If TLS support is optional, then this command line argument should also be optional
- Provide a command line argument for specifying a PEM encoded TLS private key
  - If TLS support is optional, then this command line argument should also be optional
- Provide an optional command line argument for specifying a PEM encoded TLS client certificate

## Alternatives

Keep the existing certificate handling, and implement workarounds for
deployments/tooling that do not match up with existing expectations.

Using Kubernetes and cert-manager as an example, this would require either:

- Having an external process that watches the cert-manager generated secret and
creating or updating a separate secret that is consumed by the Tinkerbell
deployment.
- Using an init container to create the expected files from the files that
are mounted using the cert-manager secret
