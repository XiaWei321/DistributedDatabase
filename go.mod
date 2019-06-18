module DistributedDatabase

require (
	DistributedDatabase/core v0.0.0
	DistributedDatabase/utils v0.0.0
	github.com/allegro/bigcache v1.2.1 // indirect
	github.com/aristanetworks/goarista v0.0.0-20190607111240-52c2a7864a08 // indirect
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/ethereum/go-ethereum v1.8.27 // indirect
	github.com/garyburd/redigo v1.6.0 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/kylelemons/go-gypsy v0.0.0-20160905020020-08cad365cd28 // indirect
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7 // indirect
	github.com/rs/cors v1.6.0 // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
	golang.org/x/net v0.0.0-20190613194153-d28f0bde5980 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
)

replace DistributedDatabase/utils v0.0.0 => ./utils

replace DistributedDatabase/core => ./core
