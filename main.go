package main

import (
	"DistributedDatabase/core"
	"DistributedDatabase/utils"
)

func main(){

	utils.GetYaml()
	utils.RedisCmdFileWatcher()
	lsi := core.LogicServiceImp{}
	isi := core.InitServiceImp{}
	isi.InitChannel()
	//isi.InitRedisConnection()
	//isi.InitEthereumConnection()
	lsi.WatchRedisChannalChange()
	lsi.WatchEthereumMessage()


}