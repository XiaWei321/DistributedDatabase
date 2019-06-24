package core

import "C"
import (
	"DistributedDatabase/utils"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/garyburd/redigo/redis"
	"sync"
	"time"
)



type InitServiceImp struct {

}

type LogicServiceImp struct {

}

type RecieveAofReciept struct{

	AofIpfsHash string
}

var redisConnection redis.Conn
var ethereumConnection *ethclient.Client
var mutex sync.Mutex

var aofChannel chan RecieveAofReciept
var transactionChannel chan string
var messageChannel chan string



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

	if ethereumConnection == nil {
		mutex.Lock()
		if ethereumConnection == nil {
			conn, err := ethclient.Dial(utils.Conf.EC.EthereumUrl)
			if err != nil{
				utils.Log.Error("获取以太坊连接失败: ",err)
			}else{
				ethereumConnection = conn

			}
		}
		mutex.Unlock()
	}
}


func (isi InitServiceImp) InitChannel(){

	aofChannel = make(chan RecieveAofReciept)
	utils.UploadChannel = make(chan bool)
	messageChannel = make(chan string)
	utils.MergeInstructChannel = make(chan []utils.Instruction)
	transactionChannel = make(chan string)
}




func (lsi LogicServiceImp) SendIpfsHashToEthereum(ipfsHash string) (txHash string){


	return ""

}



func (lsi LogicServiceImp) WatchRedisChannalChange(){


	go func(){
		flag := true
		for{
			<- utils.UploadChannel
			if flag {
				ipfsHash := utils.UploadFileToIpfs()
				utils.Log.Info("修改后的Redis历史记录文件为: ", ipfsHash)
				encodeHash := utils.EncryptTransactionInput(ipfsHash)
				to := common.HexToAddress("")
				from := common.HexToAddress(utils.Conf.EC.EthereumAdminAddress)
				message := utils.NewMessage(from, &to, "0x10","0x"+encodeHash, "0x295f05","0x77359400")
				txHash := utils.SendTransaction(ethereumConnection,&message, utils.Conf.EC.EthereumAdminPassword,context.TODO())
				utils.Log.Info("提交的交易Hash为: ", txHash)
				transactionChannel <- txHash
				utils.WaitUtilNoPendingTransactions(ethereumConnection)
			}
			flag = !flag


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


func (lsi LogicServiceImp) WatchEthereumMessage(){


	for {

		txHash := <- transactionChannel
		go func(){
			for{
				reciept := utils.GetTrasnactionReciept(ethereumConnection, txHash)
				if reciept != nil {
					break
				}
				time.Sleep(time.Duration(100)*time.Millisecond)
			}
			//TODO: 解析交易里面的值

		}()

	}

}

func WatchMergedInstructions(){


	go func(){

		instructions := <- utils.MergeInstructChannel
		go utils.WriteInstructionsToFile(instructions)
	}()

}
