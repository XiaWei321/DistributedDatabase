package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	DB Database `yaml:"database"`
	EC EthereumConfig `yaml:"ethereum-config"`
	IPFS Ipfs `yaml:"ipfs"`
}

type Database struct {
	Rs Redis `yaml:"redis"`
}

type Ipfs struct {

	ExecCommand string `yaml:"execute-command"`
	DockerUploadPath string `yaml:"docker-upload-path"`
	HostUploadPath string `yaml:"host-upload-path"`
}

type Redis struct {
	Address string `yaml:"redis-url"`
	Port string `yaml:"ethereum-url"`
	HistroyPath string `yaml:"histroy-path"`
	DownloadPath string `yaml:"download-path"`
	CommandFilePath string `yaml:"command-file-path"`
}

type EthereumConfig struct {
	Address string `yaml:"address"`
	Port string `yaml:"port"`
	EthereumAdminAccount string `yaml:"ethereum-admin-account"`
	EthereumAdminAddress string `yaml:"ethereum-admin-address"`
	EthereumAdminPassword string `yaml:"ethereum-admin-password"`
	EthereumContractAddress string `yaml:"ethereum-contract-address"`
}

var Conf Config

func GetYaml()(){

	//execpath, err := os.Executable()
	//configFilePath := filepath.Join(execpath,"../conf/conf.yaml")
	abspath, _ := filepath.Abs("../DistributedDatabase/conf/configuration.yaml")
	//abspath := "C:\\Users\\huyifan01\\Documents\\ID-Generator\\conf\\conf.yaml"
	yamlFile, err := ioutil.ReadFile(abspath)
	if err != nil{
		Log.Error("获取日志文件出错",err)
	}

	yaml.Unmarshal(yamlFile,&Conf)


}