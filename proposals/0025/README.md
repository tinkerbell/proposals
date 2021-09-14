---
id: 0025
title: Deprecate OSIE
status: discussion
authors: Thomas Stromberg
---

## Summary

Deprecate and archive the tinkerbell/osie repository

## Background

OSIE is the Operating System Installation Environment. OSIE has been in use by
Equinix Metal since before Tinkerbell was created. OSIE contains three major
components, only one of which has been in use by the Tinkerbell OSS community:

* A live Linux environment, built with Alpine Linux
* A Python script named osie-runner (deprecated by tink-worker)
* A Docker container of installation scripts (deprecated by tinkerbell/workflows)

While the latter two components were never actively used by Tinkerbell, the Linux
environment was the standard in-memory operating system for workflows execution
until Hook was introduced. In September 2021, the Tinkerbell sandbox switched from
OSIE to Hook for its live environment, meaning that OSIE is no longer in active use.

As OSIE came into existence previous to Tinkerbell, the repository is primarily
comprised of code that is irrelevant to the Tinkerbell community.
This tech debt is visible when you look at the code footprint of the two repositories
side-by-side:

![code size chart](chart.png "code size chart")

### Goals

* Bring the Tinkerbell community together behind a single live Linux environment
* Eliminate technical debt
* Eliminate confusion about which Live environment users should adopt

### Non-Goals

* Eliminate the ability for users to choose their own live Linux environment

## Proposal

* Update the tinkerbell/osie repository state from `Experimental` to `Deprecated`
* Commit to accepting bug fixes for the first 60 days, but not feature requests
* After 60-days, archive the tinkerbell/osie repository - which will make it
  read-only
* Allow users to continue to maintain their own local unsupported OSIE forks
  ad infinitum

## Alternatives

* Continue to maintain both Hook and OSIE
