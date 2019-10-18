package main

import (
	"fmt"
	"os/exec"
)

func main() {
	concurrentExecutions := 100

	outChan := make(chan string)

	go func() {
		for ix := 0; ix < concurrentExecutions; ix++ {
			go func() {
				out, err := exec.Command("node", "app.js").Output()
				fmt.Println("current iteration, ", ix)
				if err != nil {
					panic(err)
				}
				outChan <- string(out)
			}()
		}
	}()

	for ix := 0; ix < concurrentExecutions; ix++ {
		fmt.Println(<-outChan)
	}
}
