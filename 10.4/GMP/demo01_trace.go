package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {

	//1. Create trace file
	f,err:=os.Create("./10.4/GMP/trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	//2.Start the trace
	err=trace.Start(f)
	if err != nil {
		panic(err)
	}
	// common task
	fmt.Println("Hello GMP")

	//3. stop the trace
	trace.Stop()

}


