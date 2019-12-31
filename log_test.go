package ilog

import (
	"testing"
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
}

func Test_FileLog_Rotate(t *testing.T) {
	Init(LogFile, EnableRotate())
	Info("a", "b", "c")
	for {
		if Done() {
			break
		}
	}
}

func Test_FileLog_Rotate_File(t *testing.T) {
	Init(LogFile, EnableRotate(), WithFile("aaa.log"))
	Info("a", "b", "c")
	for {
		if Done() {
			break
		}
	}
}