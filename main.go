package main

import (
	"flag"
	"fmt"
	"os"
)

type Args struct {
	Time   int
	IsList bool
}

var flags Args

func init() {
	const (
		usageTime   = "working minitus"
		usageIsList = "flag of display config"
	)

	flag.IntVar(&flags.Time, "time", 25, usageTime)
	flag.IntVar(&flags.Time, "t", 25, usageTime+" (shorthand)")
	flag.BoolVar(&flags.IsList, "list", false, usageIsList)
	flag.BoolVar(&flags.IsList, "l", false, usageIsList+" (shorthand)")
	flag.Parse()
}

func main() {
	args := flag.Args()

	if len(args) <= 0 {
		fmt.Println("no command")
		os.Exit(1)
	}

	switch args[0] {
	case "start":
		// start timer
		// create timer file to tmp/current
		fmt.Println("start!")
	case "stop":
		// stop timer
		// update file of tmp/current
		fmt.Println("stop!")
	case "clear":
		// clear timer
		// update file of tmp/current
		fmt.Println("clear!")
	case "config":
		// show/edit config file
		// update $HOME/.config/chiri/config.yml
		fmt.Println("config!")
	case "stat":
		// display statistics like top command
		fmt.Println("stat!")
	}
}
