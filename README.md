## **CfClientMonitor** [![Build Status](https://drone.io/github.com/maciejmrowiec/cf_client_monitor/status.png)](https://drone.io/github.com/maciejmrowiec/cf_client_monitor/latest) 

New Relic monitoring for cfengine client performance.

#### Features

* **disksize/cfengine** - cfengine installation disk size

* **cpu/loadaverage/1m** - system wide loadaverage 1 minute
* **cpu/loadaverage/5m** - system wide loadaverage 5 minutes
* **cpu/loadaverage/15m** - system wide loadaverage 15 minutes

* **cpu/tasks/active** - number of currently runnable kernel scheduling entities (processes, threads)
* **cpu/tasks/total** - number of kernel scheduling entities that currently exist on the system

* **memory/rss/average/command_name** - RSS memory usage, sampled by 1s, is the NewRelic sampling inerval is larger, 1s samples are averaged. Unit: [KB]


#### Installation

###### Dependencies

Requires golang toolchain.

```
sudo apt-get install golang
```

###### Build

```
go get github.com/yvasiyarov/newrelic_platform_go
go get github.com/maciejmrowiec/cf_client_monitor
go build
```

#### Usage

*  -interval=1: Sampling interval [min]
*  -key="": Newrelic license key (required)
*  -verbose=false: Verbose mode

To deamonize in backgrund you can use:

```
nohup ./cfclientmonitor -key=<my_newrelic_key> >/dev/null 2>&1 &
```

**Depends on:** 

* 'du' command
* /proc/loadavg
* 'ps -aux' with 6th column RSS and 11th column command name