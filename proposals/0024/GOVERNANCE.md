# Governance

This document describes how the Liberal contribution governance structure is implemented in the Tinkerbell community.

## What is Liberal contribution?

Under a liberal contribution model, the people who do the most work [in the prescribed way] are recognized as most influential, but this is based on current work and not historic contributions.
Major project decisions are made based on a consensus seeking<sup>[1],[2]</sup> process (discuss major grievances) rather than pure vote, and strive to include as many community perspectives as possible.
Popular examples of projects that use a liberal contribution model include Node.js and Rust<sup>[3]</sup>.

## Community Roles

This doc outlines the various roles in the Tinkerbell community.

One of the overarching goals of having different roles in the community is to make sure to "allow good changes to enter the code with a minimum of fuss"<sup>[4]</sup>.
The following criteria is to be used when considering an individual for a role.

- a pattern of good judgement in regard to the code bases and to oneself
- demonstrated history of "playing nice" in the community
- amount and consistency of contributions and activity

| Role        | Responsibilities | Requirements | Defined by |
| -----       | ---------------- | ------------ | ------- |
| Contributor | Submit contributions | Signed DCO | Registered GitHub User |
| Member      | Issue and discussion administration, PR reviewer | History of activity, good judgement, contribution in the community and the repositories, and sponsored by 2 Committers | Tinkerbell GitHub org member |
| Committer   | Right to make changes to the code bases, merge pull requests, etc | Same as Member, shows leadership in the community and the repositories, and sponsored by 2 Maintainers | [OWNERS] file committer entry |
| Maintainer  | Voting privileges, set direction and priorities | Same as Committer, demonstrated responsibility and excellent technical judgement for the project, and voted in by existing Maintainers | [OWNERS] file maintainer entry |

## Member

Members are continuously active contributors in the community.
They can have issues and PRs assigned to them and pre-submit tests are automatically run for their PRs.
Members are expected to remain active contributors to the community.

**Defined by:** Member of the Tinkerbell GitHub organization

### Requirements

- Enabled [two-factor authentication] on their GitHub account
- Have made multiple contributions to the project or community.  Contributions may include, but are not limited to:
  - Authoring or reviewing PRs on GitHub
  - Filing or commenting on issues on GitHub
  - Contributing to community discussions (e.g. meetings, Slack, email discussion forums)
- Subscribed to [tinkerbell-contributors@googlegroups.com]
- Actively contributing to 1 or more repository.
- Sponsored by 2 committers. **Note the following requirements for sponsors**:
  - Sponsors must have close interactions with the prospective member - e.g. code/design/proposal review, coordinating on issues, etc.
  - Sponsors must be committers or maintainers in at least 1 OWNERS file in any repo in the [Tinkerbell org].
  - Sponsors should strive to be from multiple member companies to demonstrate integration across community.
- **[Open an issue]() [membership request] against the Tinkerbell/org repo**
  - Ensure your sponsors are @mentioned on the issue
  - Complete every item on the checklist ([preview the current version of the template](https://github.com/kubernetes/org/blob/main/.github/ISSUE_TEMPLATE/membership.yml))
  - Make sure that the list of contributions included is representative of your work on the project.
- Have your sponsoring reviewers reply confirmation of sponsorship: `+1`
- Once your sponsors have responded, your request will be reviewed by the [Tinkerbell GitHub Admin team]. Any missing information will be requested.

### Responsibilities and Privileges

- Can be assigned to issues and PRs; people can ask members for reviews with a `/cc @username`.
- Responsive to issues and PRs assigned to them
- Active owner of code they have contributed (unless ownership is explicitly transferred)
  - Code is well tested
  - Tests consistently pass
  - Addresses bugs or issues discovered after code is accepted
- Members can do `/lgtm` on open PRs.
- Tests can be run against their PRs automatically. No `/ok-to-test` needed.
- Members can do `/ok-to-test` for PRs that have a `needs-ok-to-test` label, and use commands like `/close` to close PRs as well.

## Committer

Committers are able to review code for quality and correctness on some part of Tinkerbell. They are knowledgeable about both the codebase and software engineering principles.

**Defined by:** committers entry in an OWNERS file in a repo owned by the Tinkerbell project.
Committer status is scoped to the repo/code in the Tinkerbell org.

>Note: Acceptance of code contributions requires at least one committer/maintainer.

### Requirements

The following applies to the repo or part of the codebase for which one would be a committer in an OWNERS file.

- member for at least 3 months
- Primary reviewer for at least 5 PRs to the codebase
- Reviewed or merged at least 20 substantial PRs to the codebase
- Knowledgeable about the codebase
- Sponsored by a committer/maintainer
  - With no objections from another committer/maintainer
  - Done through PR to update the OWNERS file
- May either self-nominate, be nominated by a committer/maintainer

### Responsibilities and Privileges

The following apply to the repo or the part of the codebase for which one would be a committer in an OWNERS file.

- Tests are automatically run for Pull Requests from members of the Tinkerbell GitHub organization
- --Code reviewer status may be a precondition to accepting large code contributions--
- Responsible for project quality control via code reviews
  - Focus on code quality and correctness, including testing
  - Focus on design principles
- Expected to be responsive to review requests as per community expectations
- Assigned PRs to review related to area of expertise
- Assigned bugs related to area of expertise
- May get a badge on PR and issue comments

## Inactive members

>This inactive members section follows the guideline defined in the Kubernetes community<sup>[5]</sup>

_Members are continuously active contributors in the community._

A core principle in maintaining a healthy community is encouraging active
participation. It is inevitable that people's focuses will change over time and
they are not expected to be actively contributing forever.

However, being a member of the Tinkerbell GitHub organization comes with
an [elevated set of permissions]. These capabilities should not be used by those
that are not familiar with the current state of the Tinkerbell project.

Therefore members with an extended period away from the project with no activity
will be removed from the Tinkerbell Github Organization and will be required to
go through the org membership process again after re-familiarizing themselves
with the current state.

### How inactivity is measured

Inactive members are defined as members of the Tinkerbell Organization
with **no** contributions across any repos within 8 months. This is
measured by the CNCF [DevStats project].

**Note:** Devstats does not take into account non-code contributions. If a
non-code contributing member is accidentally removed this way, they may open an
issue to quickly be re-instated.

After an extended period away from the project with no activity
those members would need to re-familiarize themselves with the current state
before being able to contribute effectively.

[1]: https://producingoss.com/html-chunk/consensus-democracy.html
[2]: https://en.wikipedia.org/wiki/Consensus-seeking_decision-making
[3]: https://opensource.guide/leadership-and-governance/#what-are-some-of-the-common-governance-structures-for-open-source-projects
[4]: https://producingoss.com/html-chunk/committers.html
[5]: https://github.com/kubernetes/community/blob/master/community-membership.md
