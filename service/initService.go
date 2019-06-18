package service

type InitService interface {

	InitChannel()

	InitRedisConnection()

	InitEthereumConnection()

}
