package utils

import (
	"bytes"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"io"
	"os"
	"os/exec"
	"strings"
)

var UploadChannel chan bool

func UploadFile(filename string) (string, error) {
	// run ipfs add -r filename
	fmt.Println("执行命令： "+Conf.IPFS.ExecCommand)
	cmds := strings.Split(Conf.IPFS.ExecCommand," ")
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
	cmds := strings.Split(Conf.IPFS.ExecCommand," ")
	myhash := strings.Split(hash, "\000")
	finalhash := myhash[0]
	cmd := exec.Command(cmds[0], cmds[1], cmds[2], "ipfs", "get", finalhash, "-o="+filename)
	err := cmd.Run()

	return err
}

func CopyFile(from string, dest string){

	fromFile, err := os.Open(from)
	if err != nil {
		Log.Error("打开文件src失败: ",err)
	}
	defer fromFile.Close()

	destFile, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		Log.Error("打开文件dest失败: ",err)
	}
	defer destFile.Close()

	io.Copy(destFile, fromFile)

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

	CopyFile(Conf.DB.Rs.HistroyPath, Conf.IPFS.HostUploadPath)
	ipfsHash, err := UploadFile(Conf.IPFS.DockerUploadPath)
	if err != nil {
		Log.Error("上传文件到ipfs失败: ", err)
		return ""
	}
	return ipfsHash
}
















