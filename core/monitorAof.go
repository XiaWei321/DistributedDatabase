package core

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/kylelemons/go-gypsy/yaml"
)

type MonitorServiceImp struct{

}

type RecieveAofReciept struct{

	AofIpfsHash string
}


var aofChannel chan RecieveAofReciept
var connection redis.Conn
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
	}
	connection = conn
}

func (msi MonitorServiceImp) InitIpfsConnection(){

	ipfsAddress, _ := config.Get("ipfs.address")

	

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
