package utils

import (
	"bytes"
	"github.com/fsnotify/fsnotify"
	"os/exec"
	"strings"
)

var UploadChannel chan bool

func UploadFile(filename string) (string, error) {
	// run ipfs add -r filename
	cmds := strings.Split(Conf.IPFS.execCommand," ")
	cmd := exec.Command(cmds[0], cmds[1], cmds[2] ,"ipfs", "add", "-r", filename)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		// run: ipfs daemon
		cmdIpfsDaemon := exec.Command(cmds[0], cmds[1], cmds[2], "ipfs", "daemon")
		cmdIpfsDaemon.Run()
		// try again
		cmd := exec.Command(cmds[0], cmds[1], cmds[2], "ipfs", "add", "-r", filename)
		cmd.Stdout = &out
		err := cmd.Run()

		if err != nil {
			Log.Error("上传文件到IPFS失败: ", err)
			return "", err
		}
	}
	out_str := strings.Split(out.String(), " ")
	hash := out_str[1]

	return hash, nil
}



func DownloadFile(hash string, filename string) (error){
	cmds := strings.Split(Conf.IPFS.execCommand," ")
	myhash := strings.Split(hash, "\000")
	finalhash := myhash[0]
	cmd := exec.Command(cmds[0], cmds[1], cmds[2], "ipfs", "get", finalhash, "-o="+filename)
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
	Log.Debug("file path: ",Conf.DB.Rs.HistroyPath)
	err = watcher.Add(Conf.DB.Rs.HistroyPath)

	if err != nil{
		Log.Error("将文件加入监听列表失败: ", err)
	}


}


func UploadFileToIpfs()(string){

	filePath := Conf.DB.Rs.HistroyPath

	ipfsHash, err := UploadFile(filePath)
	if err != nil {
		Log.Error("上传文件到ipfs失败: ", err)
		return ""
	}
	return ipfsHash
}
















