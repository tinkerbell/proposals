---
id: 0021
title: The Tinkerbell Action Hub
status: accepted
authors: Gianluca Arbezzano <ciao@gianarb.it>
---

## Summary

This proposal is about having a hub for "certified" and reusable actions and
workflows.  It will run on `hub.tinkerbell.org` and I pictured it as a read only
registry with a static site in front of it. More about the actual implementation
later and concrete tools selected moving forward.

## Goals and not Goals

Goal:
1. Define goals and requirement for the Hub
2. Explain how we can implement the Hub MVP

No-Goal:

1. I don't want to define what `certified` means

## Content

Tinkerbell uses workflows to describe the work that has to be done in a piece of
hardware. A workflow is composed of a set of actions, the smallest unit of work
inside a workflow.

A workflow is a descriptive specification shaped as a YAML file; actions are OCI
or Docker images run as container via Docker.

A key concept for Tinkerbell is reusability. Action is written, tested, and
built once, and they can be reused across many workflows.

Reusability improves reliability and stability because more users will stress
and push the limit of a specific workflow or Action. The community wants to
build a healthy mechanism to share workflows and actions.

As part of this proposal, we will stay focused on Action because, as I wrote, it
is the smallest unit of work, and it is packaged as an OCI image, a very well
known and established packaging format.

Having a well-established way to share actions and workflows is essential
because domain experts will share their experience. Operating system experts,
for example, will be able to maintain and share their Actions and Workflows, all
the community will benefit from that.

Actions look like this: Wipe a disk Partitioning a disk Interact with a BMC to
restart a server Sending a message on Slack at some point during workflow
execution

The Hub will work as a catalog for Actions first and later for Workflows.

The Tinkerbell community wants to stay independent, and we identified the
[ArtifactHub](https://artifacthub.io/) project sponsored by CNCF (as Tinkerbell
itself) as the right partner for building the Action Hub. I shared our
requirements and idea during an ArtifactHub contributors' meeting, and they are
happy to collaborate with us.

I identified a couple of needs: We need a place (a git repository) where actions
can be built and shared in collaboration. This is the source of truth and the
place where contributors collaborate on all the Actions the community wants to
support. Via CI/CD, we can validate, build, and push images to a registry.  We
need an image registry where our Actions will be released to. We have to find
one; it can be Quay, Docker Hub, our one, or something else. I think it is
essential to have multi-arch support out of the box because we want to run
actions across ARM, AMD, x84, etc.  We need a searchable catalog on a website
where users can discover and use the available actions. This is what ArtifactHub
is.  Actions are "just Docker images," so it is not essential at this point to
have a place where external contributors can push their Action because any
available registry makes it possible already. But ArtifactHub is not limited to
the Tinkerbell community; everybody will add their actions and mark them as
"Tinkerbell action." We are okay with that, but only the ones stored in our
repository will be maintained by the community.

Falco is a popular Security scanning tool; they have rules that describe what to
look for. ArtifactHub recently added support for [Falco
rules](https://artifacthub.io/packages/search?page=1&kind=1); this is a list of
all the Falco rules available in ArtifactHub. You can filter by
[organization](https://artifacthub.io/packages/search?page=1&kind=1&org=falco)
getting only the one released by the Falco community.

The idea is the same, we will get a Tinkerbell Action category, and the one we
support will be part of the Tinkerbell organization.

---

A good side effect of having Actions part of a repository that has many more
artifacts is discoverability. Using metadata and research in the right way will
bring people to Tinkerbell looking for helm charts, for example.
---

ArtifactHub is a website that scrapes and renders metadata provided by a git
repository. We have to implement a CI/CD system that will build, validate, and
push actions, plus it will have to generate the metadata file that ActionHub
periodically scrape. This is an example of metadata file:

```yaml
# Artifact Hub package metadata file
version: A SemVer 2 version (required)
name: The name of the package (only alphanum, no spaces, dashes allowed) (required)
displayName: The name of the package nicely formatted (required)
createdAt: The date this package was created (RFC3339 layout) (required)
description: A short description of the package (required)
logoPath: Path to the logo file relative to the package directory (optional, but it improves package visibility)
digest: String that uniquely identifies this package version (optional)
license: SPDX identifier of the package license (https://spdx.org/licenses/) (optional)
homeURL: The URL of the project home page (optional)
appVersion: The version of the app that this contains (optional)
containersImages: # (optional)
  - name: Image identifier (optional)
    image: The format should match ${REGISTRYHOST}/${USERNAME}/${NAME}:${TAG}
    whitelisted: When set to true, this image won't be scanned for security vulnerabilities
containsSecurityUpdates: Whether this package version contains security updates (optional, boolean)
operator: Whether this package is an Operator (optional, boolean)
deprecated: Whether this package is deprecated (optional, boolean)
keywords: # (optional)
  - A list of keywords about this package
  - Using one or more categories names as keywords will improve package visibility
links: # (optional)
  - name: Title of the link (required for each link)
    url: URL of the link (required for each link)
readme: | # (optional)
  Package documentation in markdown format

  Content added here will be rendered on Artifact Hub
install: | # (optional)
  Brief install instructions in markdown format

  Content added here will be displayed when the INSTALL button on the package details page is clicked.
changes: # (optional)
  - A list of changes introduced in this package version
  - Use one entry for each of them
maintainers: # (optional)
  - name: The maintainer name (required for each maintainer)
    email: The maintainer email (required for each maintainer)
provider: # (optional)
  name: The name of the individual, company, or service that provides this package (optional)
ignore: # (optional, entries use .gitignore syntax)
  - lib
```

https://github.com/artifacthub/hub/blob/master/docs/metadata/artifacthub-pkg.yml

Having a clear separation between the metadata used by ArtifactHub and the
actual artifact (the image itself) will sound a complication, but it gives us
more control and flexibility between ongoing work (PRs on specific actions) and
their release cycle (ActionHub artifact).

## System-context-diagram

A contributor will interface with a traditional repository, currently located in
[github.com/gianarb/hub](https://github.com/gianarb/hub), it will be moved to its
permanent location [github.com/tinkerbell/hub](https://github.com/tinkerbell/hub).

I will explain the layout using the command `tree` and commenting various
directories:

```terminal
./hub $ tree -L 2
.
├── Makefile
├── README.md
    # The actions directories contains a subfolder for every aciton, in this
    # case we have only one called disk-wipe
├── actions
│   └── disk-wipe
    # The manifest related to artifacts live in their own subdirectory
    # because at some point we will add workflows as well.
│   └── actions
    # this contains utility code that empowers validation, and automation for
    # this repository in form of CLI usually.
├── cmd
├── go.mod
├── go.sum
    # CLi scripts used for automation purpose
├── hack
    # Go code
├── pkg
└── shell.nix
```

The generated manifest are loaded in a separate branch called in our case
`artifacthub-manifests`.

All the automation is based on GitHub Action, and it will be covered in issues
and PRs in the hub repository. But there is already a command available that you
can run via:

```terminal
go run cmd/gen/main.go generate
```

It will generate the artifacthub-manifests starting from the `./actions`
directory.

Another command that we will write is something like `actionhub build
./actions/disk-wipe` it will build, push and tests single actions using
[buildkit](https://github.com/moby/buildkit).

All of those workflows will be glued as GitHub Actions workflows.

## APIs

## Alternatives

I think the end solution is what we need, I can't see alternatives for the high
level implementation. Alternatives can be discussed for the specific tools such
like registries, git repository structure or ArtifactHub.
