package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

func exp1() {
	s1 := pflag.NewFlagSet("exp1", pflag.ExitOnError)
	fmt.Println(s1.HasFlags())
}

func exp2() {
	s2 := pflag.NewFlagSet("exp1", pflag.ExitOnError)
	s2.Int("si", 0, "少年中国说")
	fmt.Println(s2.HasFlags())
}

//go:generate go run main.go
func main() {
	exp1()
	exp2()
}
