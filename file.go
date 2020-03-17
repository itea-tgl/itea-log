package ilog

import (
	"strings"
	"sync"
)

const (
	DefaultFile  = "itea.log"
	DefaultClean = 30
	TypeInfo	 = "INFO"
	TypeError 	 = "ERROR"
	TypeDebug 	 = "DEBUG"
	TypeFatal 	 = "FATAL"
)

var (
	logLevel	[]string
	directory 	string
	clean 		int
)

func init() {
	logLevel = []string{
		TypeInfo, TypeError,
		//TypeDebug, TypeFatal,
	}
	directory = "./"
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

	dir(f.option.logfile)

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

func (f *File) withFile(s string) {
	if string(s[0]) == "/" {
		s = s[1:]
	}
	f.option.logfile = s
}

func (f *File) fileKeep(n int) {
	clean = n
}

func dir(s string) {
	a := strings.Split(s, "/")
	if l := len(a); l > 1 {
		a[l-1] = ""
		directory = directory + strings.Join(a, "/")
	}
	//log.Fatalln(directory)
}

func LogFile() ILog {
	clean = DefaultClean
	return &File{option:&option{
		logfile:DefaultFile,
	}}
}