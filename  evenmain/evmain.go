package main

import (
	"even"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

type Work struct {
	Workname string
	Workid   int
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	for i := 0; i < 100; i++ {
		fmt.Println(even.Even(i), even.Odd(i))
	}
}
func server(workChan <-chan *Work) {
	for work := range workChan {
		go safelyDo(work) // start the goroutine for that work
	}
}

func safelyDo(work *Work) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Work failed with %s in %v", err, work)
		}
	}()
	do(work)
}
func do(work *Work) {

}
