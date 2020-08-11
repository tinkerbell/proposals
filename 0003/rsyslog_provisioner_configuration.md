## Configuration of rsyslog daemon for docker and container

## Changes on rsyslog
```
/etc/rsyslog.conf:
$ModLoad imtcp
$InputTCPServerRun 514
```

ls /etc/rsyslog.d/docker_daemon.conf
ls /etc/rsyslog.d/docker_container.conf

## Changes on docker daemon
```
sudo vi /etc/rsyslog.d/docker_daemon.conf
$template DockerLogs, "/var/log/dockerlfs/daemon.log"
if $programname startswith 'dockerd' then -?DockerLogs
& stop
```

## Changes on fetching docker configuration
```
sudo vi /etc/rsyslog.d/docker_container.conf
$template DockerContainerLogs,"/var/log/dockerlfs/%hostname%_%syslogtag:R,ERE,1,ZERO:.*container_name/([^\[]+)--end%.log"
if $syslogtag contains 'container_name'  then -?DockerContainerLogs
& stop
```

## Restarting service
```
sudo service rsyslog restart/status
```
