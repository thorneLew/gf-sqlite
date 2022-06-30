package utils

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

func Command(cmd string, model string, filename string, f func()) error {
	logger := CmdLog(model, filename+".log")
	fmt.Println("父 协程 开始...")
	//c := exec.Command("cmd", "/C", cmd) 	// windows
	c := exec.Command("bash", "-c", cmd) // mac or linux
	// c.SysProcAttr = &syscall.SysProcAttr{Setpgid: false}
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup

	logger.Println(TASKSTART)
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			// 删除标准输出的换行符
			text := strings.TrimSuffix(readString, "\n")
			logger.Println(text)
		}
	}()
	err = c.Start()
	wg.Wait()
	if err := c.Process.Kill(); err != nil {
		return err
	}
	if err := c.Wait(); err != nil { //避免僵尸进程
		return err
	}
	f()
	logger.Println(TASKEND)
	return err
}

func CmdLog(model string, filename string) *glog.Logger {
	logger := glog.New()
	logger.SetConfigWithMap(g.Map{
		"path":        "./logs/" + model,
		"file":        filename,
		"level":       "all",
		"stdout":      false,
		"HeaderPrint": false,
	})
	return logger
}
