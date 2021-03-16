package main

import (
	_ "github.com/pojol/registrator/consul"
	_ "github.com/pojol/registrator/consulkv"
	_ "github.com/pojol/registrator/etcd"
	_ "github.com/pojol/registrator/skydns2"
	_ "github.com/pojol/registrator/zookeeper"
)
