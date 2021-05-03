package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

const (
	NOTICE_COMMAND   = "/usr/bin/osascript"
	NOTICE_ARG_MODEL = `display notification "%s" with title "%s" `
	MAX_COUNT        = 12

	WORK_TIME  = 30
	RELAX_TIME = 5
)

type Notice struct {
	Title     string
	Content   string
	ShowCount int
}

func (n *Notice) Show() {
	for i := 0; i < n.ShowCount; i++ {
		command := exec.Command(NOTICE_COMMAND, "-e", fmt.Sprintf(NOTICE_ARG_MODEL, n.Content, n.Title))
		command.Run()

		if n.ShowCount == 1 {
			return
		}
		time.Sleep(time.Second * 10)
	}

}

func main() {
	runStartNotice := Notice{"运行成功", fmt.Sprintf("%d分钟后我回提醒您~", WORK_TIME), 1}
	relaxNotice := Notice{"休息一下", fmt.Sprintf("已经工作%d分钟了！活动一下身体吧", WORK_TIME), 3}
	workNotice := Notice{"开始工作啦", "加油，开始工作~", 1}

	showCount := 1
	log.Print("运行成功")
	runStartNotice.Show()

	for {
		time.Sleep(time.Minute * WORK_TIME)
		log.Print(fmt.Sprintf("第%d次运行", showCount))
		relaxNotice.Show()

		if showCount >= MAX_COUNT {
			return
		}
		showCount++

		time.Sleep(time.Minute * RELAX_TIME)
		workNotice.Show()
	}

}
