package core

import "github.com/kylelemons/go-gypsy/yaml"

type MonitorService interface {

	Watch()

	InitChannel()

	InitRedisConnection()

	InitIpfsConnection()

	InitConfiguration(filename string)(*yaml.File)

}