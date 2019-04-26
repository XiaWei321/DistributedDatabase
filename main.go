package main


import (

	"./core"
)

func main(){


	lsi := core.LogicServiceImp{}
	isi := core.InitServiceImp{}
	isi.InitChannel()
	isi.InitRedisConnection()
	isi.InitEthereumConnection()
	isi.InitConfiguration("../conf/configuration.yaml")
	lsi.Watch()



}