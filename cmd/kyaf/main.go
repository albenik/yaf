package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/albenik/yaf"
)

type multivar []string

func (f multivar) String() string {
	return strings.Join(f, ",")
}

func (f *multivar) Set(s string) error {
	*f = append(*f, s)
	return nil
}

func main() {
	incArg := make(multivar, 0)
	flag.Var(&incArg, "i", "Include only specified resources")

	excArg := make(multivar, 0)
	flag.Var(&excArg, "x", "Exclude specified resources")

	flag.Parse()

	arg := incArg
	exc := false
	if len(arg) == 0 {
		arg = excArg
		exc = true
	}
	if len(arg) == 0 {
		fmt.Fprintln(os.Stderr, "ERROR: only one type of filter allowed")
		os.Exit(1)
	}

	filters, err := yaf.ParseArgument(arg)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		os.Exit(1)
	}

	if err = yaf.Filter(os.Stdout, os.Stdin, filters, exc); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		os.Exit(1)
	}
}
