---
id: 0002
title: Set the mood around releasing code
status: accepted
authors: Gianluca Arbezzano <gianarb92@gmail.com>
---

## Summary

This proposal sets the expectation about what it means to release code part of
the Tinkerbell organisation.

## Goals and non-Goals

Goals are:

1. Define what a release is and what it is comprised of a release
2. Define standard policies and practices for releases that all Tinkerbell
   projects should follow
3. Define Quality Assurance polices for releases.

Non-Goals:

1. How each project should be released and what it should release
2. Define who own releases
4. Define a roadmap and how it should be presented to the community.
5. Support and deprecation lifecycle

## Content

A release is a tested and stable point in time snapshot from a codebase.

We follow the [Semantic Versioning (semver) 2.0.0](https://semver.org).
Summarized as follows:

Given a version number MAJOR.MINOR.PATCH, increment the:

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards compatible manner, and
* PATCH version when you make backwards compatible bug fixes.

Additional labels for pre-release and build metadata are available as extensions
to the MAJOR.MINOR.PATCH format.

### Schedule

Every subproject should be able to follow its own cadence in terms of how often
or regularly a new release should be cut, as long as it is documented and
shared with the community. It is a responsibility for the maintainer to figure
out a good strategy that better fits their workload, process and codebase.

`MAJOR` release right now are not even evaluated because the project is too
young. By definition breaking changes require a new `MAJOR` release, right now
any of our projects have a release yet, we are working hard to reach `v1.0.0`,
until there bc breaks are allowed but they will be notified with a proper label
`bc-break` attached to the pull request.

A new `MINOR` version for Tinkerbell should be released following a proper
schedule, right now I think every two months sounds reasonable looking at the
code written and merged in the last couple of weeks. In this way the
subprojects will have an intrinsic release target if they want to get their new
code included in the Tinkerbell release.

`PATCH` release for Tinkerbell can be done as required by the subprojects
bug fixes.

### GitHub release and release website

Git supports tag and GitHub has the concept of a release based on a particular
tag.

The artifact generated as part of the release will be published as part of the
release website, using as a guidance how [HashiCorp
Release](https://releases.hashicorp.com/) works.

I like the idea to use a release website vs having to push them as manifest in
GitHub release because it gives a single place where the community can
programmatically lookup what they are looking for, without having to look for
repositories, versions and so on.

As part of the Release page on GitHub we will publish a generated changelog.

### Distribution

Docker containers, binaries, packages such as `deb`, `rpm` requires
infrastructure such as:

1. A place where the binaries should be archived and easy do be downloaded from
   (bintray, s3)
2. Docker containers requires a registry

In order to gain consistency across projects and to build familiarity for the
community we will start as follow:

1. Docker containers will be pushed to `quay.io` under the `tinkerbell`
   organisation (this is already done)
2. Binaries and possibly `rpm`, `deb` will be pushed to
   [Bintray](http://bintray.com/) because it has a friendly open source plan.

Right now there is not a formal group or structure that manages the Tinkerbell
community needs in terms of infrastructure, servers and SaaS. We are building
this group but nothing is formalized yet. As soon as the group will be assembled
they will be the right target for questions, and requests about bintray, s3,
Packet devices. In the meantime you can ask on
[Slack](https://tinkerbell.org/community/slack/) in the #tinkerbell channel.

### Architectures and operating systems

The majority of the projects in the tinkerbell organisation right now are in Go,
it means that it is easy to compile targeting different architectures and
operating system. We also have an application in JavaScript (portal) and an
operating system distribution (osie). It is not part of the this proposal to
decide what Tinkerbell should support or not. But as a direction all the project
should setup their release workflow in a way that will support:

1. Multiple operating systems and architectures
2. Static binaries and artifacts
3. Docker containers

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

Tinkerbell itself is more like a meta-project that glues all the subprojects it
depends on. In terms of quality assurance we should have a strong and reliable
pipeline of E2E tests to guarantee that the new version for the subprojects are
working fine in integration.

The subproject Tinkerbell depends on will be pinned as part of the Tinkerbell
release but their version are a "suggestion". SemVer guarantees that all the
PATCH release are working as long as the MINOR release is the same.

Tinkerbell `v0.4.0` depends from boot `v1.2.0` every PATCH release for boot
should work.

SemVer guarantees that any MINOR release after the one pinned from Tinkerbell
should work as well.

Tinkerbell `v0.4.0` depends from boot `v1.2.0` every MINOR release for
boot should work with Tinkerbell `v0.4.0` for example:

* boot `v1.1.0` is not guaranteed to work
* boot `v1.2.0` should work
* boot `v1.2.4` should work
* boot `v1.23.0` should work.
* boot `v2.0.0` is not guaranteed to work

### Continuous Delivery

All subprojects should have a continuous delivery pipeline in place capable of
creating a release from a tag.

There should also be a well documented manual release process for the purposes
of local testing and troubleshooting.

### Documentation

The documentation should be versioned and it should follow the same release
cadence as Tinkerbell. It should be possible to jump between old version of the
Tinkerbell documentation. How many, and how is out of scope for this proposal.

## System-context-diagram

## APIs

## Alternatives

I didn't find a well explained release process that can be reused. Every project
has its own way even if, at the end the outcome is similar. I am listing here
some successful release process I am involved with, please add yours if
you have because I am sure this proposal will be a glue of the best experiences
we have as a team:

* [Kubernetes Release](https://github.com/kubernetes/sig-release/tree/master/release-team)
* [Docker release page on GitHub](https://github.com/docker/docker-ce/releases)
* [Kubernetes Release page on GitHub](https://github.com/kubernetes/kubernetes/releases)
