package core

import (
	"../utils"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/garyburd/redigo/redis"
	"github.com/kylelemons/go-gypsy/yaml"
)



type InitServiceImp struct {

}

type LogicServiceImp struct {

}

type RecieveAofReciept struct{

	AofIpfsHash string
}


var aofChannel chan RecieveAofReciept
var redisConnection redis.Conn
var ethereumConnection *rpc.Client
var config *yaml.File


func (isi InitServiceImp) InitConfiguration(filename string)(*yaml.File){

	conf,err := yaml.ReadFile(filename)

	if err != nil{
		fmt.Println(err)
		return nil
	}
	config = conf

	return config
}



func (isi InitServiceImp) InitRedisConnection(){

	databaseAddress, _ := config.Get("database.redis.address")
	conn, err := redis.Dial("tcp", databaseAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	redisConnection = conn
}

func (isi InitServiceImp) InitEthereumConnection(){

	ethereumAddress, _ := config.Get("ethereum.address")
	conn , err := rpc.Dial(ethereumAddress)
	if err!= nil {
		fmt.Println(err)
		return
	}
	ethereumConnection = conn
}


func (isi InitServiceImp) InitChannel(){

	aofChannel = make(chan RecieveAofReciept)
	utils.UploadChannel = make(chan bool)
}



func (lsi LogicServiceImp) UploadAofFileToIpfs()(string){

	filePath, _ := config.Get("database.redis.aofPath")
	fmt.Println("filepath: ",filePath)
	ipfsHash, err := utils.UploadFile(filePath)
	if(err != nil){
		fmt.Println(err)
		return ""
	}
	return ipfsHash
}

func (lsi LogicServiceImp) SendIpfsHashToEthereum(ipfsHash string) (txHash string){


	return ""

}






func (lsi LogicServiceImp) Watch(){

	for{

		aofReceipt := <-aofChannel

		fmt.Println(aofReceipt.AofIpfsHash)


	}


}

func (lsi LogicServiceImp) AcquireFileFromIpfs(ipfsHash string) bool{

	downloadPath, _ := config.Get("database.redis.downloadPath")
	err := utils.DownloadFile(ipfsHash,downloadPath)
	if err!= nil{
		fmt.Println(err)
		return false
	}
	return true
}

func (lsi LogicServiceImp) RecoverRedisData(filePath string){


}


