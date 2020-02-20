package ilog

import (
	"sync"
)

const (
	DefaultFile = "itea.log"
	TypeInfo	= "INFO"
	TypeError 	= "ERROR"
	TypeDebug 	= "DEBUG"
	TypeFatal 	= "FATAL"
)

var logLevel []string

func init() {
	logLevel = []string{
		TypeInfo, TypeError,
		//TypeDebug, TypeFatal,
	}
}

type option struct {
	rotate  bool
	divide  bool
	logfile string
}

type File struct {
	logs 	map[string]*item
	wg 		sync.WaitGroup
	option 	*option
}

func (f *File) Init() {
	f.logs = make(map[string]*item)

	for _, l := range logLevel {
		f.logs[l] = NewItem(l, *f.option)
	}
}

func (f *File) Done() bool {
	f.wg.Wait()
	for _, item := range f.logs {
		item.close()
	}
	return true
}

func (f *File) Debug(v ...interface{}) {
	f.wg.Add(1)
	go func() {
		defer f.wg.Done()
		f.logs[TypeDebug].print(v...)
	}()
}

func (f *File) Info(v ...interface{}) {
	f.wg.Add(1)
	go func() {
		defer f.wg.Done()
		f.logs[TypeInfo].print(v...)
	}()
}

func (f *File) Error(v ...interface{}) {
	f.wg.Add(1)
	go func() {
		defer f.wg.Done()
		f.logs[TypeError].print(v...)
	}()
}

func (f *File) Fatal(v ...interface{}) {
	f.wg.Add(1)
	go func() {
		defer f.wg.Done()
		f.logs[TypeFatal].print(v...)
	}()
}

func (f *File) enableRotate() {
	f.option.rotate = true
}

func (f *File) enableDivide() {
	f.option.divide = true
}

func (f *File) withFile(file string) {
	f.option.logfile = file
}

func LogFile() ILog {
	return &File{option:&option{
		logfile:DefaultFile,
	}}
}