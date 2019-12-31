package ilog

import (
	"bytes"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type File struct {
	log 	*log.Logger
	wg 		sync.WaitGroup
	file 	*os.File
	rotate 	bool
	logfile string
}

func (l *File) Init() {
	if strings.EqualFold(l.logfile, "") {
		l.logfile = DefaultFile
	}
	if l.rotate {
		go l.rotateFile(l.logfile)
		l.logfile = l.rotateName(l.logfile)
	}
	file, err := os.OpenFile(l.logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND,0)
	if err != nil {
		panic("open log file error !")
	}
	l.file = file
	l.log = log.New(file, "", log.LstdFlags)
}

func (l *File) rotateName(logfile string) string {
	f := strings.Split(logfile, ".")
	var s bytes.Buffer
	s.WriteString(strings.Join(f[0:len(f) - 1], "."))
	s.WriteString("-")
	s.WriteString(time.Now().Format("2006-01-02"))
	s.WriteString(".")
	s.WriteString(f[len(f) - 1])
	return s.String()
}

func (l *File) rotateFile(logfile string) {
	filename := logfile
	for {
		now := time.Now()
		// 计算下一个零点
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
		t := time.NewTimer(next.Sub(now))
		<-t.C
		name := l.rotateName(filename)
		for {
			file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_APPEND,0)
			if err == nil {
				l.file.Close()
				l.file = file
				l.log = log.New(file, "", log.LstdFlags)
				break
			}
		}
	}
}

func (l *File) Done() bool {
	l.wg.Wait()
	l.file.Close()
	return true
}

func (l *File) Debug(v ...interface{}) {
	l.wg.Add(1)
	go func() {
		defer l.wg.Done()
		l.log.SetPrefix("[Debug] ")
		l.log.Println(v...)
	}()
}

func (l *File) Info(v ...interface{}) {
	l.wg.Add(1)
	go func() {
		defer l.wg.Done()
		l.log.SetPrefix("[INFO] ")
		l.log.Println(v...)
	}()
}

func (l *File) Error(v ...interface{}) {
	l.wg.Add(1)
	go func() {
		defer l.wg.Done()
		l.log.SetPrefix("[ERROR] ")
		l.log.Println(v...)
	}()
}

func (l *File) Fatal(v ...interface{}) {
	l.wg.Add(1)
	go func() {
		defer l.wg.Done()
		l.log.SetPrefix("[Fatal] ")
		l.log.Println(v...)
	}()
}

func LogFile() ILog {
	return &File{}
}