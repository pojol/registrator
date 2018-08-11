package main

import (
	_ "github.com/onecer/registrator/consul"
	_ "github.com/onecer/registrator/consulkv"
	_ "github.com/onecer/registrator/etcd"
	_ "github.com/onecer/registrator/skydns2"
	_ "github.com/onecer/registrator/zookeeper"
)
