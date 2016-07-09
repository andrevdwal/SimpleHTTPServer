package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	internalErr = 100
)

var (
	port = flag.Int("p", 5050, "port to listen on")
	root = flag.String("d", "_public", "root directory to serve")
	linf = log.New(os.Stdout, "", 0)
	lerr = log.New(os.Stderr, "ERROR ", 0)
	version = "0.0.0" // set when built
)

func main() {
	linf.Println(banner)
	linf.Println(version)

	flag.Parse()

	if _, err := os.Stat(*root); os.IsNotExist(err) {
		linf.Println("Creating directory:", *root)
		if err = os.MkdirAll(*root, os.ModePerm); err != nil {
			lerr.Println(err)
			os.Exit(internalErr)
		}
	}

	http.Handle("/", http.FileServer(http.Dir(*root)))
	go func() {
		lerr.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
	}()

	linf.Printf("Address:   http://localhost:%d\n", *port)
	linf.Printf("Directory: %s\n", *root)

	// wait
	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-wait

	linf.Println("Good bye!")
}

var banner = `  ___ _            _       _  _ _____ _____ ___   ___
 / __(_)_ __  _ __| |___  | || |_   _|_   _| _ \ / __| ___ _ ___ _____ _ _
 \__ \ | '  \| '_ \ / -_) | __ | | |   | | |  _/ \__ \/ -_) '_\ V / -_) '_|
 |___/_|_|_|_| .__/_\___| |_||_| |_|   |_| |_|   |___/\___|_|  \_/\___|_|
             |_|`
