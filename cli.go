package main

import (
	"fmt"
	"io"
	"time"
)

type CLI struct {
	outStream, errStream io.Writer
}

func (cli *CLI) Run(args []string) int {
	fmt.Fprintln(cli.outStream, "poe")

	next := make(chan *Runner)

	sr := NewRunner(100, []scenarioFn{TestScenario})
	go sr.start(next)
	timeoutCh := time.After(30 * time.Second)

L:
	for {
		select {
		case <-timeoutCh:
			break L
		case sr := <-next:
			go sr.run()
		}
	}

	fmt.Println("finished")
	fmt.Printf("success:%d, failure:%d", Success, Failure)
	return 1
}
