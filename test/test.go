package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main(){

	cmd := exec.Command("docker", "exec", "distributeddatabase-ethereum","netstat")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil{
		fmt.Println("执行出错: ", err)
	}
	fmt.Println(out.String())
}
