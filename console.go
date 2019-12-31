package ilog

import (
	"fmt"
	"time"
)

type Console struct {

}

func (c *Console) Init() {

}

func (c *Console) Done() bool{
	return true
}

func (c *Console) Debug(v ...interface{}){
	fmt.Println("[DEBUG]", time.Now().Format("2006/01/02 15:04:05"),  v)
}

func (c *Console) Info(v ...interface{}){
	fmt.Println("[INFO]", time.Now().Format("2006/01/02 15:04:05"), v)
}

func (c *Console) Error(v ...interface{}){
	fmt.Println("[ERROR]", time.Now().Format("2006/01/02 15:04:05"), v)
}

func (c *Console) Fatal(v ...interface{}){
	fmt.Println("[FATAL]", time.Now().Format("2006/01/02 15:04:05"), v)
}

func LogConsole() ILog {
	return &Console{}
}