package utils

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
)



var mutex sync.Mutex
var file1Instructions []Instruction
var file2Instructions []Instruction
var mergeInstructions []Instruction
var MergeInstructChannel chan []Instruction
var UploadChannel chan bool



type Instruction struct {

	Id string
	DataInstruction string

}



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


// 调用该函数时需要加锁
func FileMerge(file1 string, file2 string, dest string){


	sourceFile, err := os.Open(file1)
	if err != nil {
		Log.Error("打开文件1失败: ",err)
	}
	defer sourceFile.Close()
	sourceFile2, err := os.Open(file2)
	if err != nil {
		Log.Error("打开文件2失败: ",err)
	}
	defer sourceFile2.Close()


	r1 := bufio.NewReader(sourceFile)
	r2 := bufio.NewReader(sourceFile2)

	readInstructionsToarray(*r1, file1Instructions)
	readInstructionsToarray(*r2, file2Instructions)
	mergeInstructionsAndWriteToFile()

}

func readInstructionsToarray(r bufio.Reader, array []Instruction){

	for{

		content,_,e := r.ReadLine()
		if e == io.EOF {
			break
		}
		newInstruct := ""
		splits := strings.Split(string(content), ",")
		for i:=1; i< len(splits); i++{
			newInstruct = newInstruct+splits[i]+" "
		}

		array = append(array, Instruction{
			Id: splits[0],
			DataInstruction: newInstruct,
		})

	}


}

func mergeInstructionsAndWriteToFile(){

	st1, st2 := 0, 0

	for st1 < len(file1Instructions) && st2 < len(file2Instructions) {
		if file1Instructions[st1].Id < file2Instructions[st2].Id{
			mergeInstructions = append(mergeInstructions, file1Instructions[st1])
			st1++
		}else{
			mergeInstructions = append(mergeInstructions, file2Instructions[st2])
			st2++
		}
	}

	for st1 < len(file1Instructions) {
		mergeInstructions = append(mergeInstructions, file1Instructions[st1])
		st1++
	}
	for st2 < len(file2Instructions) {
		mergeInstructions = append(mergeInstructions, file2Instructions[st2])
		st2++
	}
	MergeInstructChannel <- mergeInstructions
}



func WriteInstructionsToFile(instructions []Instruction){

	mutex.Lock()
	file, err := os.OpenFile("../test/testFile", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		Log.Error("指令写入文件失败: ", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i:=0; i<len(instructions); i++ {
		fmt.Fprintln(writer, instructions[i])
	}
	writer.Flush()
	mutex.Unlock()
}



func DecryptTransactionInput(input string) string {

	input = input[2:]

	decodeStr, _ := hex.DecodeString(input)

	return string(decodeStr)

}

func EncryptTransactionInput(input string) string {

	encodeStr := hex.EncodeToString([]byte(input))

	return encodeStr
}










