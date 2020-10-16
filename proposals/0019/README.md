---
id: 0019
title: Use PR# as the Proposa Number
status: ideation
authors: Manny <mmendez@equinix.com>
---

## Summary

Lets avoid conflicts / race conditions when allocating proposal numbers by letting GitHub handle that.

## Goals and not Goals

* Avoid multiple PRs using the same proposal number.
* Avoid needing to look up latest proposal number when starting a new proposal.

## Content

Lets just make the GitHub PR# be the proposal number.
This will mean we'll likely see jumps in proposals since not every PR will be a proposal, but I think thats fine.

This is a break from the Oxide RFD setup and I'm ok with that.
Oxide's RFDs aren't public and the branching/numbering scheme seems to reflect that:
- checking origin for largest branch to figure out the next rfd number
- wording about pushing to master

We have a different scenario and sticking with the numbering format is a bit more difficult than necessary.

## Alternatives

Deal with clashes when they happen.
