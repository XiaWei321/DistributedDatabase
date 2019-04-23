package core

import "github.com/kylelemons/go-gypsy/yaml"

type MonitorService interface {

	Watch()

	InitChannel()

	InitRedisConnection()

	InitEthereumConnection()

	InitConfiguration(filename string)(*yaml.File)

}