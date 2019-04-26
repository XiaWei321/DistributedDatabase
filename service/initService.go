package service

type InitService interface {

	InitChannel()

	InitRedisConnection()

	InitEthereumConnection()

	InitConfiguration(filename string)

}
