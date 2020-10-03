---
id: 0010
title: Workflow behavior for on_timeout and on_failure
status: abandoned 
authors: Gaurav Gahlot <gauravgahlot0107@gmail.com>
---

## Summary:

When a workflow action runs out of time or fails, consequently its parent task and workflow are marked as timeout or failed.
In such scenarios, either `on_timeout` or `on_failure` actions must be executed.
Currently, both `on_timeout` or `on_failure` are not actions but a command that is executed after an action timeout or fails.

## Goals:

The primary goals are:
* `on_timeout` and `on_failure` to be:
  * removed from action level
  * added as root level elements in a workflow template
  * converted to actions, instead of accepting commands
* `on_timeout` action should be executed when an action (irrespective of its task) runs out of time
* `on_failure` action should be executed when an action (irrespective of its task) fails

## Content: 

- When a workflow action runs out of time or fails, consequently its parent task and workflow are marked as timeout or failed.
- In such scenarios, either `on_timeout` or `on_failure` actions must be executed.
- Currently, both `on_timeout` or `on_failure` are not actions but a command that is executed after an action timeout or fails.
- Both the entities are supported at the action level.
- As we convert them to a workflow action, they will be removed from action level and added as root level elements in a workflow.
- They will support all the elements that make an action and will take the following form:

```
...
on_failure:
  image: on-failure
  timeout: 90
    volumes:
      - ./host-path:/container-path
    environment:
      key: value
on_timeout:
  image: on-timeout
  timeout: 90
    volumes:
      - ./host-path:/container-path
    environment:
      key: value
...
```
- The complete sample template for a workflow with `on_timeout` and `on_failure` can be found [here](sample.tmpl)


## Progress:

Post discussion, this proposal will be taken forward with implementation under the [Tink repository](https://github.com/tinkerbell/tink/).
