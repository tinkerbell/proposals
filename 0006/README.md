---
id: 0006
title: Resiliency in Tinkerbell Provisioner.
status: Ideation
authors: Chitrabasu Khare <chitrabasukhare89@gmail.com>
---

## Summary:

This RFD proposes to have a resilient framework for Tinkerbell Provisioner.


## Goals:

- Bring resiliency in Tinkerbell. 
- Tinkerbell as a setup should be highly available.

## Content: 

- We would like to have a framework where Tinkerbell as a solution is fault-tolerant.
- All most all of the core services are managed by provisioner. So, it will be good to have a framework where provisioner is highly available.
- It can be a multinode provisioner solution with active-active or active-passive implementation.
- The Provisioner services should be partition tolerant during network failures or communication failure between provisioner nodes.
- Additional logic would be required to sync services running inside provisioner nodes.
- A separate machine to place an L3/L7 Load balancer or proxy would be required.


## Problems and Suggestions:

1) How does worker communicate with provisioner?
-  The worker will communicate to the load balancer (or proxy) and the request will be routed to a replica of provisioner.
-  Sessions need to be maintained inside Load balancer so that different requests of worker goes to the same replica of provisioner. 
-  How sessions can be maintained with the latest changes of GRPC streams, requires some investigation. 

2) Which services require sync among it replicas?
-  Currently, at least DB service needs a sync logic in backend. 
-  This can be achieved by having daemon agents running in both nodes and syncing the database. 

3) What happens when a request of a worker is being processed by a provisioner's replica and that replica goes down?
-  If DB is in sync then the state of the provisioner will be consistent. 
-  Additional logic inside load balancer would be required to ensure when traffic is routed to another replica, then contexts associated with the request are also changed. Workflow execution should continue will another replica now. 

4) How will boots DHCP request work with a load balancer?
-  The worker will request an IP from DHCP. Currently, GRPC request is made to get IP from the tink-server using the Provisioner IP. So, this should work the same way. 
-  Additional logic of working with GRPC request on load balancer would be required. 

## System Context Diagram

![provisioner_ha](provisioner_ha.png)
