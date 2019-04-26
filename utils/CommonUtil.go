package utils

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)


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
