---
id: 0005
title: tink-cli binary
status: abandoned 
authors: chitrabasu khare <chitrabasukhare89@gmail.com>
---

## Summary:

Tinkerbell should have tink-cli binary to execute tink-cli commands from outside the tink-cli docker container. Currently, execution of command can only be done from the inside container or via docker exec utility. 

## Goals:

- Flexibility to run tink-cli commands from outside docker container and without docker exec utility.

## Content: 

- tink-cli will be developed in the form of binary. This binary should support all Linux based platform.
- tink-cli binary would have all necessary commands to operate the provisioner and the Tinkrebell cluster.
- tink-cli would support authentication of the User.
- tink-cli would support execution of commands from the provisioner, outside of the provisioner or from any node where the binary is installed.
- tink-cli would be secure and easy to use. 
- tink-cli should be able to connect to tink-server running in the provisioner node. 
- tink-services to be exposed on a port of provisioner and tink-cli can communicate to tink-services back and forth via that port. 
- tink-cli binaries can be implemented on top of grpc or rest services.
- Please find a high-level diagram attached.

## Progress:

This proposal will be taken forward with PR https://github.com/tinkerbell/proposals/pull/5

## System Context Diagram
![tink-cli HLD](tink-cli_HLD.png)
