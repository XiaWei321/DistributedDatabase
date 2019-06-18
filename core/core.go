package core

import "C"
import (
	"DistributedDatabase/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/garyburd/redigo/redis"
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




func (isi InitServiceImp) InitRedisConnection(){

	databaseAddress := utils.Conf.DB.Rs.Address
	conn, err := redis.Dial("tcp", databaseAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	redisConnection = conn
}

func (isi InitServiceImp) InitEthereumConnection(){

	ethereumAddress := utils.Conf.EC.Address
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




func (lsi LogicServiceImp) SendIpfsHashToEthereum(ipfsHash string) (txHash string){


	return ""

}



func (lsi LogicServiceImp) WatchRedisChannalChange(){


	go func(){

		for{
			<- utils.UploadChannel
			ipfsHash := utils.UploadAofFileToIpfs()
			utils.Log.Info("修改后的Redis历史记录文件为: ", ipfsHash)
		}

	}()

}


func (lsi LogicServiceImp) AcquireFileFromIpfs(ipfsHash string) bool{

	downloadPath := utils.Conf.DB.Rs.DownloadPath
	err := utils.DownloadFile(ipfsHash,downloadPath)
	if err!= nil{
		fmt.Println(err)
		return false
	}
	return true
}


