package utils

import (
	"bytes"
	"github.com/fsnotify/fsnotify"
	"log"
	"os/exec"
	"strings"
)

var UploadChannel chan bool

func UploadFile(filename string) (string, error) {
	// run ipfs add -r filename
	cmd := exec.Command("ipfs", "add", "-r", filename)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		// run: ipfs daemon
		cmdIpfsDaemon := exec.Command("ipfs", "daemon")
		cmdIpfsDaemon.Run()
		// try again
		cmd := exec.Command("ipfs", "add", "-r", filename)
		cmd.Stdout = &out
		err := cmd.Run()

		if err != nil {
			log.Print(err)
			return "", err
		}
	}
	out_str := strings.Split(out.String(), " ")
	hash := out_str[1]

	return hash, nil
}



func DownloadFile(hash string, filename string) (error){
	myhash := strings.Split(hash, "\000")
	finalhash := myhash[0]
	cmd := exec.Command("ipfs", "get", finalhash, "-o="+filename)
	err := cmd.Run()

	return err
}




func RedisCmdFileWatcher(){

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		Log.Error("创建文件watcher失败: ",err)
	}

	go func(){

		for {

			select{
				case event := <- watcher.Events:
					if event.Op & fsnotify.Write == fsnotify.Write {
						UploadChannel <- true
					}
				case err := <- watcher.Errors:
					Log.Error("获取文件状态失败: ",err)
			}

		}

	}()



}


















