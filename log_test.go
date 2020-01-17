package ilog

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func Test_ConsoleLog(t *testing.T) {
	Init(nil)
	Info("a", "b", "c")
}

func Test_FileLog(t *testing.T) {
	Init(LogFile)
	Info("a", "b", "c")
	for {
		if Done() {
			break
		}
	}
	checkFile(t, "itea.log")
}

func Test_FileLog_Rotate(t *testing.T) {
	Init(LogFile, EnableRotate())
	Info("a", "b", "c")
	for {
		if Done() {
			break
		}
	}
	checkFile(t, fmt.Sprintf("itea-%s.log", time.Now().Format("2006-01-02")))
}

func Test_FileLog_Rotate_File(t *testing.T) {
	Init(LogFile, EnableRotate(), WithFile("aaa.log"))
	Info("a", "b", "c")
	for {
		if Done() {
			break
		}
	}
	checkFile(t, fmt.Sprintf("aaa-%s.log", time.Now().Format("2006-01-02")))
}

func checkFile(t *testing.T, file string) {
	_, err := os.Stat(file)
	if err != nil {
		if os.IsExist(err) {
			t.Errorf("log file %s not found", file)
		}
		t.Error("log file error : ", err)
		return
	}

	defer os.Remove(file)

	dat, err := ioutil.ReadFile(file)
	if err != nil {
		t.Error("read log file error : ", err)
		return
	}

	now := time.Now().Format("2006/01/02 15:04:05")
	expect := fmt.Sprintf("[INFO] %s %s", now, "a b c")

	if strings.Trim(string(dat), "\r\n") != expect {
		t.Errorf("log content error, expect %s, get %s", expect, string(dat))
	}
}