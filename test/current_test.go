package test

import (
	"fmt"
	"mytool/internal/code"
	"testing"
	"time"
)

func TestCurrentGenerator(t *testing.T){
	start :=time.Now()
	curr:=code.ConcurrencyGenerator{}
	curr.Run()
	fmt.Println("共耗时:",time.Since(start))
}
