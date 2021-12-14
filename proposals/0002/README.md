---
id: 0002
title: Set the mood around releasing code
status: accepted
authors: Gianluca Arbezzano <gianarb92@gmail.com>, Marques Johansson <mjohansson@equinix.com>
---

## Summary

This proposal defines release behaviors for projects in the Tinkerbell organisation.

## Goals and non-Goals

Goals are:

1. Define what a release is and what it is comprised of a release
2. Define standard policies and practices for releases that all Tinkerbell
   projects should follow
3. Define Quality Assurance polices for releases.

Non-Goals:

1. How each project should be released including pipelines and the artifacts to be released
2. Define who owns and maintains the release processes (CI/CD)
4. Define a roadmap and how it should be presented to the community.
5. Define a support and deprecation lifecycle

## Content

A release is a tested and stable point in time, a snapshot from a codebase.

The Tinkerbell project follows [Semantic Versioning (semver) 2.0.0](https://semver.org).
Semantic Versioning can be summarized as follows:

Given a version number MAJOR.MINOR.PATCH, increment the:

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards compatible manner, and
* PATCH version when you make backwards compatible bug fixes.

Additional labels for pre-release and build metadata are available as extensions
to the MAJOR.MINOR.PATCH format.

### Schedule

Every subproject should follow its own cadence in terms of how often
or regularly a new release is made. The process must be documented and
shared with the community. It is the responsibility of the maintainer to figure
out a good strategy that fits their projects' workload, process and codebase.

`MAJOR` releases should not be considered for young projects. Breaking changes
are inevitable at this stage and will require a new `MAJOR` release. Tinkerbell
projects will not reach `v1.0.0` until their APIs have stabilized. Breaking changes
should in v0 should be indicated with a `bc-break` label attached to the pull request
and a minor revision increment when released.

`MINOR` releases for Tinkerbell projects should be frequent to allow for changes
to be evaluated readily. It is especially important to release frequent minor revisions
for projects with a v0 major version.

`PATCH` release for Tinkerbell projects can be done as required by the subprojects
bug fixes.

### GitHub release and release website

Git supports tags and GitHub has the concept of a release based on a particular
tag.

The artifact generated as part of the release will be published as part of the
release website. Out-of-tree release artifacts, including changelogs, are preferable to artifacts and manifest
contained within Git repositories. For example, <https://github.com/tinkerbell/tink/releases>
gives a single place where the community can programmatically lookup the details and artifacts
of all or specific releases.

Within the GitHub release pages we will publish a changelog. The changelog format should
follow <https://keepachangelog.com/en/1.0.0/>.

### Distribution

OCI containers, binaries, packages such as `deb`, `rpm` require
infrastructure such as:

1. A place where the binaries should be archived and easily downloaded
2. OCI containers require a registry

In order to gain consistency across projects and to build familiarity for the
community we will start as follow:

1. OCI containers will be pushed to `quay.io` under the `tinkerbell`
   organisation (this is already done)
2. Binaries and possibly `rpm`, `deb` will be pushed to
   GitHub releases.

Right now there is not a formal group or structure that manages the Tinkerbell
community needs in terms of infrastructure, servers and SaaS. We are building
this group but nothing is formalized yet. As soon as the group will be assembled
they will be the right target for questions, and requests about bintray, s3,
Equinix Metal devices. In the meantime you can ask on
[Slack](https://tinkerbell.org/community/slack/) in the #tinkerbell channel.

### Architectures and operating systems

The majority of the projects in the tinkerbell organisation right now are in Go,
it means that it is easy to compile targeting different architectures and
operating system. We also have documentation, kernel build specs, and
operating systems (osie, hook). It is not part of the this proposal to
decide what Tinkerbell should support or not. But as a direction all the project
should setup their release workflow in a way that will support:

1. Multiple operating systems and architectures
2. Static binaries and artifacts
3. Publishing OCI artifacts to a repository

### Ownership

The contributors and maintainers for each project have the responsibility for
the quality of the release itself.

Tinkerbell as a project will have its group of interest around release
management that will take care of it (this is out of scope for this proposal but
has to be discussed in another one).

### Quality Assurance

Every project should be covered by tests: unit tests, functional tests and
integration as needed.

If all the test suite succeed the project can be released, otherwise tests
should be fixed first.

The documentation contains in tinkerbell.org covers the entire project and it
has to be reviewed before any Tinkerbell release. Each Tinkerbell MINOR release will
have its own documentation and users need a way to visualize the version they
are looking for.

### Tinkerbell

While Tinkerbell is a collection of projects, the version applied to the `tink` project
shall reflect the project version.

Dependent projects like `tinkerbell-docs` and the `sandbox` projects, which
make the other projects comprehendible and provable, must follow the major and minor
of the `tink` project to prevent confusion.  These projects should not introduce independent
resources because the major and minor versions must be pinned to `tink` version. 
Indepedent resources should be provided in a project that can be versioned independetly.

The `tink` project should adopt the latest releases of subprojects promptly.
A strong and reliable pipeline of E2E tests should guarantee that new versions of subprojects
are working with fine in a complete integration. Breaking changes to `tink`, should be avoided
when adopting subproject updates, including major revisions. When breaking changes are required,
the `tink` project should update its version accordingly. Breaking changes should be introduced
in prerelease versions for proving, for example `v2.0.0-alpha.1`.

SemVer guarantees that all the PATCH release are working as long as the MINOR release is the same.
For example, if `tink v0.4.0` depends on `boots v1.2.0`, `tink v0.4.0` should require no code or
interface changes to be built with and other `v1.2` version of `boots`.

SemVer guarantees that any MINOR release after the one pinned from Tinkerbell
should work as well.

Tinkerbell `v0.4.0` depends from boot `v1.2.0` every MINOR release for
boot should work with Tinkerbell `v0.4.0` for example:

* boot `v1.1.0` is not guaranteed to work
* boot `v1.2.0` should work
* boot `v1.2.4` should work
* boot `v1.23.0` should work.
* boot `v2.0.0` is not guaranteed to work

#### Documentation

The documentation should be versioned and it should follow the same release
cadence as Tinkerbell. The documentation website should offer the means to browse older
versions of the documentation. Likewise, users of the Git project should be able to
switch to the version of the documentation matching a tink release to find suitable
documentation.

### Continuous Delivery

All subprojects should have a continuous delivery pipeline in place capable of
creating a release from a tag.

There should also be a well documented manual release process for the purposes
of local testing and troubleshooting.

## System-context-diagram

## APIs

## Alternatives

Every project has its own approach even when the is similar. I am listing here
some successful release process I am involved with, please add yours if
you have because I am sure this proposal will be a glue of the best experiences
we have as a team:

* [Kubernetes Release](https://github.com/kubernetes/sig-release/tree/master/release-team)
* [Docker release page on GitHub](https://github.com/docker/docker-ce/releases)
* [Kubernetes Release page on GitHub](https://github.com/kubernetes/kubernetes/releases)
