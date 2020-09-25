---
id: 0007
title: Sandbox and the user persona
status: published
authors: Gianluca Arbezzano <gianarb92@gmail.com>
---

## Summary

The Tinkerbell setup, workflows, and tutorials available today are made for
contributors. These are for developers who want to contribute to Tinkerbell.

The current process checks out the `tinkerbell/tink` repository and runs the
latest versions of the dependencies: osie, tink-server, tink-worker, boots, and
so on. This is a big problem today because these components are all under active
development. We can't guarantee that the `master` branches work at all, and even
less so when integrating the various components. The level of entropy is too
high.

## Goals and non-Goals

Goals:

1. Figure out a way to serve a pinned "version" of Tinkerbell that uses known
   working dependency revisions in favor of the master branch.
2. Use pre-built binaries and Docker images to avoid any need for compiling code
   or building images.

Non-Goals:

1. Tagging project releases. We are not ready for that. [Proposal
   0002](https://github.com/tinkerbell/proposals/tree/master/0002) has captured
   our plans for that.
2. Getting something that will work forever. We need something that works NOW.
3. Making the contributor experience better.

## Content

I made a prototype at
[gianarb/sandbox](https://github.com/gianarb/playwithtink).

For the concerns of this proposal I would like to highlight two different
personas:

1. Contributors who want to write code for Tinkerbell to make it better. These
   people usually checkout the repository and they are not scared, but rather
   willing, to look at all the cool code in Tinkerbell. They want an easy way to
   run their changes with the versions of the dependencies they want to test.
2. Users wanting for stability. They don't want to see code.  It does not bring
   them joy.  They want to actually use Tinkerbell.

I think we are not doing a great job of serving what these personas want.

For the contributors and developers that developed `tinkerbell/tink`, and
continue to develop Tinkerbell, the existing patterns have worked well enough.
This proposal is the first attempt to identify and differentiate the needs of
these personas and offer an improved solution.

I don't like the idea of proposing a solution for users that is too far away
from we have in tink at the moment because:

1. Contributors should experience Tinkerbell as users do, and vice-versa.
2. We may be able to reuse code (vagrant e2e tests) to validate our solution
3. What is currently in place has been validated by the community already.
4. There's not enough time.
5. If we offer different solutions for the contributor and user experiences, I
   believe these solutions will evolve independently and we will have two very
   different solutions to support.

I propose a new repository called: `tinkerbell/sandbox`.  It will contain a
`setup.sh` script and a `docker-compose` file copied from, or equivalent to, the
ones served in `tinkerbell/tink`. The difference is that the docker images will
be pinned to known image SHAs.

We can replace or re-write setup guides similar to the ones we already have, but
they won't checkout the `tinkerbell/tink` repository, they will clone
`tinkerbell/sandbox` (where we can create releases for extra stability). In this
way the versions in use will stay the same no matter what happens to the
underlying repositories.

I see this solution as a temporary one to give something "stable" to users until
we get a proper release lifecycle up and running.

Now that I have my hands dirty with the prototype, I think this approach can be
used as an improved baseline for the tinkerbell.org documentation:

* https://tinkerbell.org/docs/setup/packet-with-terraform/
* https://tinkerbell.org/docs/setup/local-with-vagrant/

The `sandbox` works in the same way.  Instead of cloning the `tink` repo, users
will clone `sandbox`.

## Concerns

### sandbox and tink/deploy are similar

Yes, they are, for now. There is a bit of friction today, for contributors and
maintainers, when changing the contents of
[`tink/deploy`](https://github.com/tinkerbell/tink/tree/master/deploy).  We know
first time users rely on the `master` branch of `tink` and we don't want their
first experience with Tinkerbell to be a bad one.

My expectation is that as soon as we have proper releases, packages for the
various distributions, and so on, life will get easier for both users and
contributors.

In the meantime, this is a starting point to isolate what contributors work on
from what we provide users.

### Call for maintainers and owners

We need maintainers for this `sandbox` repository, including people that have
the both the will and necessary privileges to make it happen.

The responsibility of maintainers will be to bump versions and tagging sandbox
releases that users can refer to when trying Tinkerbell and when filing issues.
With issue reports that include the release version, we should be able to
identify and replicate issues quickly.
