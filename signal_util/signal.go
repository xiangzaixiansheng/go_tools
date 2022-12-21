package signal_util

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type notify struct {
	cc chan os.Signal
}

func New() *notify {
	return &notify{
		cc: make(chan os.Signal, 1),
	}

}

// 捕捉信号
func (n *notify) WaitSignal() {
	signal.Notify(n.cc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	for {
		s := <-n.cc
		fmt.Println("收到信号 -- ", s)
		switch s {
		case syscall.SIGHUP:
			fmt.Println("收到终端断开信号, 忽略")
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL:
			shutdown()
		}
	}
}

// 应用退出处理方法
func shutdown() {
	defer func() {
		fmt.Println("已退出")
		os.Exit(0)
	}()

	fmt.Println("准备退出前的操作")
}

// 发送停止信号
func (n *notify) NotifyStop() {
	n.cc <- syscall.SIGINT
}
