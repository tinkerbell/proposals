---
id: 0027
title: Tink Worker Rewrite
status: discussion
authors: Jacob Weinstock, Raj Dharwadkar, Scott Garman
---

## Summary

Tink worker's code base is quite fragile.

- Low unit test coverage
- Large, complex, and difficult to test functions
- Scattered calls to os.Exit and os.Getenv
- Tight coupling to the container runtime (docker)
- Difficult to modify or test tightly coupled functions and methods
- The action execution flow code is complex and difficult to follow and understand
- The Tink server interactions are difficult to follow and understand
- The Tink workflow status reporting call result is coupled to the worker exit status
- Lack of documentation and code comments

We proposal a rewrite of the code base to make it easier to understand and modify.

## Goals and not Goals

Goals

- high unit test coverage
- documented design philosophy
- easy to understand and modify code
- build to be able to handle different container runtimes (containerd, kubernetes, etc)
- backward compatible CLI interaction
- support for public container registries
- support for a single container registry with authentication
- implement the global timeout (fixes [#198](https://github.com/tinkerbell/tink/issues/198))
- drop support for [emphemeral data](https://docs.tinkerbell.org/workflows/working-with-workflows/#ephemeral-data)
- document 

Non-Goals

- creating an implementation for a different container runtime
- modifying Tink server

## Content

The core functionality of Tink worker will remain.

- querying Tink server for workflows
- executing workflows
- reporting action and workflow status

The CLI flags will remain the same so as to be backward compatible. The rewrite will be a drop in replacement for the existing Tink worker. No changes to Hook should be necessary.
The majority of the rewrite will focus on structuring the code in such a way to make it easier to understand and modify.
We will following the existing [Boot design philosophy](https://github.com/tinkerbell/boots/blob/main/docs/DESIGNPHILOSOPHY.md) tenets and incorporate 3 of our own.

- Engineer with clear and obvious layers of concern and purpose. -- [ref](https://github.com/ardanlabs/service/wiki#design-philosophy-review-and-culture)
- All proposed code changes should strive to be maintainable, manageable, and debug-able -- [ref](https://github.com/ardanlabs/service/wiki#design-philosophy-review-and-culture)
- All proposed code changes should strive to be the minimal code needed right now to solve the problem -- [ref](https://github.com/ardanlabs/service/wiki#design-philosophy-review-and-culture)

What will be changed?

- Container registry support
  - Currently, only a single container registry, that has authentication, is supported.
  - We will continue to support this but also will be adding support for any unauthenticated container registry.
- Ephemeral data
  - Currently, ephemeral data is a little known and little used feature. It is not documented in the code base. The only document we were able to find lives [here](https://docs.tinkerbell.org/workflows/working-with-workflows/#ephemeral-data). From the code stand point, this feature feels very unfinished. It is complicated to understand and maintain.
  - With this proposal we will be removing support for ephemeral data. We will poll the community for feedback on this.

## Test plan

The plan is have about >= 85% unit test coverage. We will also test with the [Sandbox](https://github.com/tinkerbell/sandbox) repository.
