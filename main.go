package main


import (

	"./core"
)

func main(){


	msi := core.MonitorServiceImp{}

	msi.InitChannel()
	msi.InitIpfsConnection()
	msi.InitRedisConnection()
	msi.InitConfiguration("../conf/configuration.yaml")
	msi.Watch()



}