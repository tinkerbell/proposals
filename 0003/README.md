---
id: 0003
title: Centralized Logging on Tinkerbell: 
status: POC done, Code change ready
authors: chitrabasu khare <chitrabasukhare89@gmail.com>
---


## Summary:
---------- 
To have a centralized logging framework for tinkerbell. 


## Overview: 
---------
- Tinkerbell will have the feature of centralized logging. 
- By default, the logging fetaure would be supported by rsyslog. 
- Logging feature should implemented in such a way that a user can modify the type of logging (rsyslog or fluentbit) as per his needs. This may require user to make additional changes on his end to enable that feature. 


## Goals: 
--------
Goals covers the feature Details:
- rsyslog is unix based package to capture the logs in a connected network. This entails to have a centralized logging using rsyslog. 
- rsyslog is based on the top of syslog. 
- rsyslog reads the log in syslog daemon service.
- rsyslog will run on a logging server. Description of this logging server will be provided as a part of logging configuration.
- Using some configuration files as rules, user can define the logs he/she in interested in and can direct them to a defined file.
- These file are core files which define how rsyslog centralized loging would work. 
- Since this implementation uses daemon services, it is necessary to run containerized tink services in root privileages. 
- All core tink services (including worker) would be using rsyslog implementation only and syslog as the log driver. It has been done to ensure logging of core services is correct and consistent. 
- docker-compose file have the logging configurations for tink services on provisioner. For tink-worker changes will be osie at workflow-helper.sh
- Action containers will have the logging as rsys or other (if user defined).
- tls and Log rotation support would be added. 


## Progress:
------------
This implementation of rsyslog has been verified with POC and been discussed indepth within Packet and Infracloud team.  


## Suggestions:
---------------
- We want to ensure at a basic minimum there is a centralized logging framework which exists for tinkerbell. In order to have generic implementation of configuring log-drivers below are some proposals.
 
1) How to make action containers run with logging drivers.
 
   a) Place the logging details inside workflow template. 
	version: "0.1"
	name: hello_world_workflow
	global_timeout: 600
	tasks:
	  - name: "hello world"
	    worker: "{{.device_1}}"
	    actions:
	      - name: "hello_world"
		image: hello-world
		timeout: 60  
	logging:
	  type: syslog  # log driver
	  config:
	    syslog-address : tcp://192.168.1.1:514
	    tag: "container_name/{{.Name}}"
	    
    - This will ensure all the worflow action containers have consistent way of logging informations. 
    - logging configuration would be specific to workflow containers. 
    
  
  b) Host level configuration. 
    - We can have a file for host level configuration. This will set environment variables which can be used in action.go while creating the containes. 
    - This will have a way where all the workflow template will have the common logging framework.
    - Problem
      - Design issue: logging here is specific to action containers. So having them in workflow temaplate make more sense other than in separate config file. 
      - maintaining multi node cluster 
      - updating the logging configuration. 
    


   Workaround 
   - Have configuration defined in action.go when action container are created. This will ensure we have bare centralized framwework ready.
   - Challenge:  
     - change in implementation when changes are done in provisioner. Like, moving whole infrastructure on a private network. 


## Future roadmap:
-----------------------
As a part of roadmap implementation certain changes in existing framework would be required to bring below proposal onboard.

1) Provisioner/Worker configuration
   Pushing logging configuration on the nodes via some template or config file. 

2) Centralized logging in Private infrastructure of Provisioner

3) Fault tolerance in logging. 
- Handling of scenarios where communication can break over network or nodes can go down. 
- Handling of scenarios where there can be distributed provisioner having services running across multiple nodes. 

    
4) Log files structure:
- Worker container Files to have workflow Id. In current framework this can be done. Changes would be required to attach id of workflow in the name of worker action containers. 
- Files for each node will be placed in the respective directory on logging server.


5) Implementation with fluentbit and other logging drivers:  
- Separate doc on how rsyslog, fluentbit would work. 


## Documentation:
-----------------
- Doc on working of logging framework with rsyslog implementation will be added in first release. 


## References:
---------------
https://docs.docker.com/config/containers/logging/syslog/#options
https://man7.org/linux/man-pages/man8/rsyslogd.8.html



