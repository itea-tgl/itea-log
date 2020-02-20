package ilog

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var rotateItems []*item

type item struct {
	filename string
	prefix string
	log *log.Logger
	file *os.File
}

func NewItem(p string, o option) *item {
	i := &item{filename: o.logfile, prefix: p}
	if o.divide {
		o.logfile = i.formatFilename()
	}
	i.logPrefix()
	if o.rotate {
		rotateFile(i)
		o.logfile = i.rotateName()
	}
	file, err := os.OpenFile(o.logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND,0)
	if err != nil {
		panic("open log file error !")
	}
	i.file = file
	i.log = log.New(file, i.prefix, log.LstdFlags)
	return i
}

func (i *item) rotateName() string {
	l := len(i.filename)
	if i.filename[l-4:] == ".log" {
		i.filename = i.filename[:l-4]
	}
	var s bytes.Buffer
	s.WriteString(i.filename)
	s.WriteString("-")
	s.WriteString(time.Now().Format("15-04"))
	s.WriteString(".log")
	return s.String()
}

func rotateFile(i *item) {
	rotateItems = append(rotateItems, i)
	if len(rotateItems) > 1 {
		return
	}
	go func() {
		fmt.Println(1111111111111111111)
		for {
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Second * 60)
			next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
			for _, i := range rotateItems {
				go func(i *item) {
					name := i.rotateName()
					for {
						file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_APPEND,0)
						if err == nil {
							i.log = log.New(file, i.prefix, log.LstdFlags)
							err = i.file.Close()
							if err != nil {
								log.Println("former log file close error : ", err)
							}
							i.file = file
							break
						}
					}
				}(i)
			}
		}
	}()
}

func (i *item) formatFilename() string {
	i.filename = fmt.Sprintf("%s-%s", strings.ToLower(i.prefix), i.filename)
	return i.filename
}

func (i *item) logPrefix() {
	i.prefix = fmt.Sprintf("[%s] ", i.prefix)
}

func (i *item) print(v ...interface{}) {
	i.log.Println(v...)
}

func (i *item) close() {
	err := i.file.Close()
	if err != nil {
		log.Println(err)
	}
}