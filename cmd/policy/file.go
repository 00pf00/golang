package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	memMax := int64(0)
	go func() {
		for {

			tmp := GetMemSize()
			if memMax < tmp {
				memMax = tmp
			}
			time.Sleep(1 * time.Second)

		}
	}()

	defer func() {
		fmt.Print(memMax)
	}()
	GetFileF()

}

func exec_shell(line string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", line)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

func GetMemSize() int64 {
	pid := os.Getpid()
	mem, err := exec_shell("cat /proc/" + strconv.Itoa(pid) + "/statm")
	if err != nil {
		fmt.Println("fail")
	}
	fmt.Println(mem)
	UB, err := exec_shell("getconf PAGESIZE")
	UB = strings.Replace(UB, "\n", "", -1)
	if err != nil {
		fmt.Print(err)
	}
	mems := strings.Split(mem, " ")
	mennum, err := strconv.ParseInt(mems[0], 10, 64)
	menunit, err := strconv.ParseInt(UB, 10, 64)
	return mennum * menunit

}

func GetFileM() {
	url := "https://mirrors.tuna.tsinghua.edu.cn/ubuntu-releases/18.04.2/ubuntu-18.04.2-desktop-amd64.iso"
	rep, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}
	md5h := md5.New()
	io.Copy(md5h, rep.Body)
	fmt.Print(md5h.Sum([]byte("")))
}

func GetFileF() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)
	path = path + "/file/"
	path = path + strconv.FormatInt(time.Now().Unix(), 10) + ".iso"
	file, err := os.Create(path)
	defer file.Close()
	url := "https://mirrors.tuna.tsinghua.edu.cn/ubuntu-releases/18.04.2/ubuntu-18.04.2-desktop-amd64.iso"
	rep, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}
	buf := make([]byte, 8*1024*128)
	io.CopyBuffer(file, rep.Body, buf)
	md5, err := exec_shell("md5sum " + path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(md5)
}
