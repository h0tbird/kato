# The Káto system

**Káto** (from Greek *κάτω*: 'down', 'below', 'underneath') is an opinionated system which governs diverse computing workloads and work-flows.
Like in catabolism (from Greek *κάτω* káto, 'downward' and *βάλλειν* ballein, 'to throw'), the *Káto* system is the catalyst used to breakdown complex monolithic platforms into its fundamental microservices.

</br>

<img src="https://raw.githubusercontent.com/h0tbird/coreseed/master/imgs/kato.png"
 alt="Booddies logo" title="Booddies" align="right" width="69%" height="69%"/>

**Distinctive attributes**

- Geolocation
- Multidatacenter
- Cloud agnostic
- Variable costs
- Hardware abstraction
- Endo/exo-elasticity
- Microservices
- Containerization
- Task scheduling
- CI/CD pipelines
- Service discovery
- Load balancing
- High availability
- Self-healing

</br>

# Overview

[CoreOS](https://coreos.com/) is the foundation on which *Káto* is built. It provides the fundamental components used to assemble container-based distributed systems: [etcd](https://github.com/coreos/etcd) is used for consensus and discovery, [fleet](https://github.com/coreos/etcd) is a distributed init system, [flannel](https://github.com/coreos/flannel) is used for virtual networking and [rkt](https://github.com/coreos/rkt) and [docker](https://github.com/docker/docker) are container engines.

All this *CoreOS* goodies are used to bootstrap a [Mesos](https://github.com/apache/mesos) cluster. *Mesos* is a distributed systems kernel which abstracts compute resources away from machines. Accordingly, it provides schedulers (or frameworks in *Mesos* parlance) which can run on top in order to utilise the exposed compute resources.

[Marathon](https://github.com/mesosphere/marathon) is one of such frameworks. It is a cluster-wide init and control system for long-running applications. Other frameworks like [Jenkins](https://github.com/jenkinsci/mesos-plugin) and [Elasticsearch ](https://github.com/mesos/elasticsearch) share the same cluster resources.

[Ceph](https://github.com/ceph/ceph-docker) is a distributed object store and file system designed to provide excellent performance, reliability and scalability. It provides a multi-host persistent data storage for containers.

# Components

|Component|Current Version|Container|
|---|:---:|:---:|
|[CoreOS](https://coreos.com)|[alpha](https://coreos.com/releases/)|-|
|[Mesos](http://mesos.apache.org)|[0.27.2](https://git-wip-us.apache.org/repos/asf?p=mesos.git;a=blob_plain;f=CHANGELOG;hb=0.27.2)|[![Docker Pulls](https://img.shields.io/docker/pulls/mesosphere/mesos-master.svg)](https://hub.docker.com/r/mesosphere/mesos-master/)|
|[Mesos-DNS](http://mesosphere.github.io/mesos-dns)|[0.5.2](https://github.com/mesosphere/mesos-dns/releases/tag/v0.5.2)|[![Docker Pulls](https://img.shields.io/docker/pulls/mesosphere/mesos-dns.svg)](https://hub.docker.com/r/mesosphere/mesos-dns/)|
|[Marathon](https://mesosphere.github.io/marathon)|[0.15.3](https://github.com/mesosphere/marathon/releases/tag/v0.15.3)|[![Docker Pulls](https://img.shields.io/docker/pulls/mesosphere/marathon.svg)](https://hub.docker.com/r/mesosphere/marathon/)|
|[Zookeeper](https://zookeeper.apache.org)|[3.4.8](https://zookeeper.apache.org/doc/r3.4.8/)|[![Docker Pulls](https://img.shields.io/docker/pulls/h0tbird/zookeeper.svg)](https://hub.docker.com/r/h0tbird/zookeeper/)|
|[go-dnsmasq](https://github.com/janeczku/go-dnsmasq)|[1.0.0](https://github.com/janeczku/go-dnsmasq/releases/tag/1.0.0)|[![Docker Pulls](https://img.shields.io/docker/pulls/janeczku/go-dnsmasq.svg)](https://hub.docker.com/r/janeczku/go-dnsmasq/)|
|[cAdvisor](https://github.com/google/cadvisor)|[0.22.0](https://github.com/google/cadvisor/releases/tag/v0.22.0)|[![Docker Pulls](https://img.shields.io/docker/pulls/google/cadvisor.svg)](https://hub.docker.com/r/google/cadvisor/)|
|[Ceph](http://ceph.com)|[9.2.0](https://github.com/h0tbird/docker-ceph/releases/tag/v9.2.0-2)|[![Docker Pulls](https://img.shields.io/docker/pulls/h0tbird/ceph.svg)](https://hub.docker.com/r/h0tbird/ceph/)|

## Install
```
go get github.com/h0tbird/kato
go build -o ${GOPATH}/bin/katoctl github.com/h0tbird/kato
```

## 1. Deploy Káto's infrastructure

- [x] [Vagrant](https://github.com/h0tbird/coreseed/blob/master/docs/vagrant.md)
- [x] [Amazon EC2](https://github.com/h0tbird/coreseed/blob/master/docs/ec2.md)
- [x] [Packet.net](https://github.com/h0tbird/coreseed/blob/master/docs/packet.md)

## 2. Start the stack
```
cd /etc/fleet
fleetctl submit zookeeper\@.service
fleetctl start zookeeper@1 zookeeper@2 zookeeper@3
fleetctl start mesos-master.service mesos-dns.service
fleetctl start marathon.service cadvisor.service
fleetctl start dnsmasq.service mesos-node.service
fleetctl start ceph-mon.service ceph-osd.service
```

## License
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
