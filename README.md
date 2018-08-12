# Registrator
## 问题描述

在这篇文章《[Docker Swarm Mode中部署SpringCloud微服务](https://uublog.com/article/20180805/docker-swarm-deploy-micro-services/)》之后

遇到了点新问题。

在运行时指定自己网络的时候，容器里面多个网络，多个IP，但是注册的不是固定注册某个网卡的IP作为服务IP。

这就导致有时候注册的IP不是属于overlay网络的IP。服务访问就会出现问题。

<!--more-->

如果是用rencher部署的话，它有一个label记录了容器IP。可以用`registrator`的`-useIpFromLabel`读取它label。而docker swarm并没有记录相关IP到label 也就不能用这个参数。

所以我fork了registrator的最新源码，做了点修改。让它支持可以传入在运行docker时，attach的network名字，根据这个网络名字在它container信息中取得它IP。

如我前面运行docker的服务的时候指定`--network=micro-service`

那么我运行`registrator`的时候，用我修改过的镜像`doubleshit/registrator:v7`指定`-useIpFromNetworkName=micro-service`参数即可。

如：

```bash
docker run -d \
    --restart=always \
    --name=registrator \
    --net=host \
    --volume=/var/run/docker.sock:/tmp/docker.sock \
    doubleshit/registrator:v7 \
    -cleanup \
    -internal \
    -ip <NODE_IP> \
    -useIpFromNetworkName=micro-service \
      consul://<NODE_IP>:8500
```

*commit的时候信息network写成了netword无视就好 哈哈*


Service registry bridge for Docker.

[![Circle CI](https://circleci.com/gh/gliderlabs/registrator.png?style=shield)](https://circleci.com/gh/gliderlabs/registrator)
[![Docker pulls](https://img.shields.io/docker/pulls/gliderlabs/registrator.svg)](https://hub.docker.com/r/gliderlabs/registrator/)
[![IRC Channel](https://img.shields.io/badge/irc-%23gliderlabs-blue.svg)](https://kiwiirc.com/client/irc.freenode.net/#gliderlabs)
<br /><br />

Registrator automatically registers and deregisters services for any Docker
container by inspecting containers as they come online. Registrator
supports pluggable service registries, which currently includes
[Consul](http://www.consul.io/), [etcd](https://github.com/coreos/etcd) and
[SkyDNS 2](https://github.com/skynetservices/skydns/).

Full documentation available at http://gliderlabs.com/registrator

## Getting Registrator

Get the latest release, master, or any version of Registrator via [Docker Hub](https://registry.hub.docker.com/u/gliderlabs/registrator/):

	$ docker pull gliderlabs/registrator:latest

Latest tag always points to the latest release. There is also a `:master` tag
and version tags to pin to specific releases.

## Using Registrator

The quickest way to see Registrator in action is our
[Quickstart](https://gliderlabs.com/registrator/latest/user/quickstart)
tutorial. Otherwise, jump to the [Run
Reference](https://gliderlabs.com/registrator/latest/user/run) in the User
Guide. Typically, running Registrator looks like this:

    $ docker run -d \
        --name=registrator \
        --net=host \
        --volume=/var/run/docker.sock:/tmp/docker.sock \
        gliderlabs/registrator:latest \
          consul://localhost:8500

## CLI Options
```
Usage of /bin/registrator:
  /bin/registrator [options] <registry URI>

  -cleanup=false: Remove dangling services
  -deregister="always": Deregister exited services "always" or "on-success"
  -internal=false: Use internal ports instead of published ones
  -ip="": IP for ports mapped to the host
  -resync=0: Frequency with which services are resynchronized
  -retry-attempts=0: Max retry attempts to establish a connection with the backend. Use -1 for infinite retries
  -retry-interval=2000: Interval (in millisecond) between retry-attempts.
  -tags="": Append tags for all registered services
  -ttl=0: TTL for services (default is no expiry)
  -ttl-refresh=0: Frequency with which service TTLs are refreshed
```

## Contributing

Pull requests are welcome! We recommend getting feedback before starting by
opening a [GitHub issue](https://github.com/gliderlabs/registrator/issues) or
discussing in [Slack](http://glider-slackin.herokuapp.com/).

Also check out our Developer Guide on [Contributing
Backends](https://gliderlabs.com/registrator/latest/dev/backends) and [Staging
Releases](https://gliderlabs.com/registrator/latest/dev/releases).

## Sponsors and Thanks

Big thanks to Weave for sponsoring, Michael Crosby for
[skydock](https://github.com/crosbymichael/skydock), and the Consul mailing list
for inspiration.

For a full list of sponsors, see
[SPONSORS](https://github.com/gliderlabs/registrator/blob/master/SPONSORS).

## License

MIT

<img src="https://ga-beacon.appspot.com/UA-58928488-2/registrator/readme?pixel" />
