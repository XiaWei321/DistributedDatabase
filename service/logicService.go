package service


type LogicService interface {


	SendIpfsHashToEthereum(ipfsHash string)

	WatchRedisChannalChange()

	AcquireFileFromIpfs(ipfsHash string)
	
	WatchEthereumMessage()
}
