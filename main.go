package main

import (
	"./core"
	"./utils"
)

func main(){

	utils.GetYaml()
	utils.RedisCmdFileWatcher()
	lsi := core.LogicServiceImp{}
	isi := core.InitServiceImp{}
	isi.InitChannel()
	isi.InitRedisConnection()
	isi.InitEthereumConnection()
	lsi.WatchRedisChannalChange()



}