---
id: 0007
title: PlaywithTink and the user persona
status: discussion
authors: Gianluca Arbezzano <gianarb92@gmail.com>
---

## Summary

The setup workflows and tutorial we have today in my opinion are for
contributors. For developers who want to code for Tinkerbell.

Because the process checkouts the `tinkerbell/tink` repository, and it runs the latest version of
dependencies like: osie, tink-server, tink-worker, boots and so on. This is a
big problem today because all those components are under development, and we
can't guarantee that what it is in master works. Even less in integration with
the other components. The entropy is too high.

## Goals and non-Goals

Goals:

1. Figure out a way to serve a pinned "version" of Tinkerbell that does not
   use the current master version but something that is tested.
2. Avoid compiling or building code or docker images in favor of using already
   available solutions.

Non-Goals:

1. Tag a release for project, we are not ready for that and there is proposal
   0002 under discussion about that.
2. Getting to something that will work forever, we need something that works
   NOW.
3. Make the contributor experience better

## Content

I made a prototype
[gianarb/playwithtink](https://github.com/gianab/playwithtink).

For what concerns this proposal I would like to highlight two different
personas:

1. Contributors who want to write code for Tinkerbell to make it better. Those
   people usually checkout the repository and they are not scared but willing to
   look at all the cool code written in there. And they want an easy way to run
   their changes and the version of the dependencies they need to test
2. Users are willing for stability and they don't want to see code that is not
   useful to actually bring them to the joy of playing with Tinkerbell.

I think we are not doing a great job to serve those two kind of people with what
they want.

This proposal is the first attempt to provide something different for those
personas. As a contributors and developers we developed what is available as
part of the `tinkerbell/tink` repo, and we are developing Tinkerbell with it, so
it works enough for what I can tell.

I don't like to propose a solution for our user that is too far away from we
have in tink at the moment because:

1. It will help contributors to become user and even more important vice-versa.
2. We will be able to reuse some code (vagrant e2e tests) to validate this
   solution
3. It is validated by the community already.
4. Not enough time
5. I think we will learn from having two different solutions for contributors
   and users those will evolve by themself as soon as they will be by their own.

We can have a new repository called: `tinkerbell/playwithtink` it will contain
a setup.sh copy/paste or equivalent script and a `docker-compose` file like the
one we serve in `tinkerbell/tink` but with pinned docker images.

We can replace or re-write setup guides similar to the one we already have, but
they won't checkout the `tinkerbell/tink` repository, they will clone
`tinkerbell/playwithtink` (as a release if we want extra safety) in this way the
version in use will stay the same no matter what happens in the underline
repositories.

I see this solution as a temporary one to get something "stable" to users until
we get a proper release lifecycle up and running.

Now that I have my hands dirty with the prototype I think this can be used as
baseline for the tinkerbell.org documentation:

* https://tinkerbell.org/docs/setup/packet-with-terraform/
* https://tinkerbell.org/docs/setup/local-with-vagrant/

`playwithtink` works in the same way but other than cloning the repo `tink`
users will clone `playwithtink`.

### Playwithtink and tink/deploy are similar

Yes they are. For now. There is a bit of friction today when changing the
content of `tink/deploy` as a contributor/maintainer because we know users that
never used tink before relay on that setup to play with it.

My expectation is that as soon as we will have proper releases, packages for the
various distribution and so on life will get easier for both users and
contributors.

In the meantime this is a starting point to segment what we provide for
contributors vs what we provide for users.

### Call for maintainers and owners

We need a maintainer for this repository and also somebody that can actually
make it to happen.

The reponsability for the maintainers will be to bump versions up and to tag
playwithtink releases that users can refer to when trying tinkerbell and when
opening a PR. From a release we should be able to identify what they run and
replicate their issues quickly.
