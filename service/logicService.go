package service


type LogicService interface {

	UploadAofFileToIpfs(filePath string)

	SendIpfsHashToEthereum(ipfsHash string)

	Watch()

	AcquireFileFromIpfs(ipfsHash string)

	RecoverRedisData(filePath string)

}
