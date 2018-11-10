package models

import (
	"fmt"

	"github.com/hpcloud/tail"
)

var Tails *tail.Tail

func NewTailf(filename string) {
	//filename := ".\\my.log"
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen: true,
		Follow: true,
		//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	Tails = tails
}
