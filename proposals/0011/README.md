---
id: 0011
title: how to announce a breaking change for Tinkerbell
status: published
authors: Gianluca Arbezzano <gianarb92@gmail.com>
---

## Summary

A breaking change (short bc break) is a planned update that breaks compatibility
for a feature or a piece of code. What was working before, won't work the same
way after a bc break. They are sometime needed to simplify a functionality, remove
feature not needed or to fix wrong decisions.

At this stage, where the various services are not released and do not have
semver we can accept BC break.

We need to have a consolidated way to communicate them.

## Goals and no-Goals

Goal:

* Clarify the way we can communicate bc break across services in Tinkerbell.
* A solution with low overhead that has to work NOW.

No-Goal:

* Explain how SemVer manages Bc-Break or how we will communicate bc-break when
  SemVer will be in use.
* Explain how a component should communicate bc-break to their community.

## Content

Even if we do not expect many BC break the possibility to remove or change a
functionality is a luxury that won't stay for long. We can do it right now,
because we did not release a stable release yet and community is coming with
more and more use cases and requests that will may be better to achieve giving
only one way to do things.

Even if we don't have a day for the first major release of Tinkerbell or for the
other component I doubt it will be next week, this means that BC break are
expected and the way we communicate them makes the difference between an happy
community manager and somebody who looks for a different solution.

Every PR that introduce a BC break will get labeled with a `bc-break` label. In
this way we will be able to build a workflow that programmatically will extract
the BC-Break for every component.

Every PR has to contain as part of the description a section titled:
"how to migrate" teaching the user about how it should pass from a previous
version to a new one when possible. If not possible because the section will
explain why we remove the entire functionality and will offer an alternative if
it exists.

### How To Migrate

The "how to migrate" section has to be placed in the PR description. Here a few
things you should include:

1. You have to be clear about the required steps.
2. You can write a bash script if necessary that can be used to perform
   the migration.
3. This chapter is the documentation users will rely on when performing an
   update. You should explain how to validate if the migration went right.
4. If a bash script is not enough feel free to implement a new binary, even
   better if it has a `--dry-run` approach that can be used to validate the
   action that will be performed without having to do them straight away.
5. We know that updates are crucial, and critical, sometime it is better to have
   the best written format of what it is required vs a too magic bash script
6. Something like
   [coccinelle](https://en.wikipedia.org/wiki/Coccinelle_(software) or [`go
   fix`](https://golang.org/cmd/fix/) should be evaluated as well.

### Share And Aggregate BC Breaks

With the idea that every component should be independent and interchangeable I
don't want to write as part of this PR how boots, osie will have to communicate
their BC break to the outside, but we need to agree at minimum to be able to
communicate them to the all Tinkerbell community.

[tinkerbell/sandbox](https://github.com/tinkerbell/sandbox) is the way the
community interacts with Tinkerbell following the documentation (not true yet
but very soon). The sandbox project will follow a semver and it will bump a
major release when required, in general it should report as part of the
changelog the bc-break from the underline component.

## System-context-diagram

## APIs

## Alternatives
