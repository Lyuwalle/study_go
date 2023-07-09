package main

import (
	"testing"
	"time"
)

//单元测试，使用go test -v命令测试
func TestPrint1(t *testing.T) {
	print1()
}
func TestGoPrint1(t *testing.T) {
	goPrint1()
	time.Sleep(1 * time.Millisecond)
}
func TestGoPrint2(t *testing.T) {
	goPrint2()
	time.Sleep(1 * time.Millisecond)
}
