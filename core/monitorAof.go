package core

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/garyburd/redigo/redis"
	"github.com/kylelemons/go-gypsy/yaml"

)

type MonitorServiceImp struct{

}

type RecieveAofReciept struct{

	AofIpfsHash string
}


var aofChannel chan RecieveAofReciept
var redisConnection redis.Conn
var ethereumConnection *rpc.Client
var config *yaml.File


func (msi MonitorServiceImp) InitConfiguration(filename string){

	conf,err := yaml.ReadFile(filename)

	if err != nil{
		fmt.Println(err)
		return
	}
	config = conf
}



func (msi MonitorServiceImp) InitRedisConnection(){

	databaseAddress, _ := config.Get("database.redis.address")
	conn, err := redis.Dial("tcp", databaseAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	redisConnection = conn
}

func (msi MonitorServiceImp) InitEthereumConnection(){

	ethereumAddress, _ := config.Get("ethereum.address")
	conn , err := rpc.Dial(ethereumAddress)
	if err!= nil {
		fmt.Println(err)
		return
	}
	ethereumConnection = conn
}






func (msi MonitorServiceImp) InitChannel(){

	aofChannel = make(chan RecieveAofReciept)

}


func (msi MonitorServiceImp) Watch(){

	for{

		aofReceipt := <-aofChannel

		fmt.Println(aofReceipt.AofIpfsHash)

	}


}
