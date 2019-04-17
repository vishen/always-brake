package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

var (
	timeout      = time.Second * 5
	fatalMessage = "always-break has broken!"
	exitCode     = 1
)

func setFromEnv() {
	if t := os.Getenv("TIMEOUT_SECONDS"); t != "" {
		v, err := strconv.Atoi(t)
		if err != nil {
			log.Fatalf("unable to parse TIMEOUT_SECONDS %q to an integer: %v", t, err)
		}
		timeout = time.Second * time.Duration(v)
	}

	if m := os.Getenv("MESSAGE"); m != "" {
		fatalMessage = m
	}

	if e := os.Getenv("EXIT_CODE"); e != "" {
		v, err := strconv.Atoi(e)
		if err != nil {
			log.Fatal("unable to parse EXIT_CODE %q to an integer: %v", e, err)
		}
		exitCode = v
	}
}

func main() {
	setFromEnv()
	log.Printf("Waiting %v to break...", timeout)
	time.Sleep(timeout)
	log.Printf(fatalMessage)
	os.Exit(exitCode)
}
