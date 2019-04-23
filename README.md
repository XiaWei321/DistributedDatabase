# DistributedDatabase
- Go Version: v1.11.9
- Ethereum Version: v1.8.27
- Docker Version: v1.13.1
- Redis Version: 

在DistributedDatabase目录下使用docker-compose up 命令一键启动`redis`,`ethereum`以及`ipfs`
> **注意：启动ethereum的docker之后我们需要进入容器初始化以太坊的账户。并且redis容器启动之后需要进入redis容器执行systemctl restart redis**

