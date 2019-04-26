package main

import (
	"../core"
	"fmt"
)

func main(){

	lsi := core.LogicServiceImp{}
	isi := core.InitServiceImp{}
	isi.InitConfiguration("../conf/configuration.yaml")
	ipfsHash := lsi.UploadAofFileToIpfs()
	fmt.Println("ipfsHash: ",ipfsHash)

	result := lsi.AcquireFileFromIpfs(ipfsHash)
	if result{
		fmt.Println("success")
	}else{
		fmt.Println("fail")
	}
}
