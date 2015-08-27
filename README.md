## **CfClientMonitor** [![Build Status](https://drone.io/github.com/maciejmrowiec/cf_client_monitor/status.png)](https://drone.io/github.com/maciejmrowiec/cf_client_monitor/latest) 

New Relic monitoring for cfengine client performance.

#### Features

*disksize/cfengine* - cfengine installation disk size


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

**Depends on commands:** 

* du