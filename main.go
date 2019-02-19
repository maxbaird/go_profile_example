package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

/* Profiling flags to set at runtime*/
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")    //cpu profiling
var memprofile = flag.String("memprofile", "", "write memory profile to `file`") // memory profiling

func genVec(size int) []int {
	vec := make([]int, size)

	for i := 0; i < size; i++ {
		vec[i] = i
	}
	return vec
}

func squareVec(vec []int) []int {

	for i, e := range vec {
		vec[i] = e * e
	}
	return vec
}

func sumVec(vec []int) int {
	var sum int

	for _, e := range vec {
		sum = sum + e
	}
	return sum
}

func main() {
	/***** Start of profiling code block *****/
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	/***** End of profiling code block *****/

	/* Start of code and function calls to profile */
	vec := genVec(1024 * 1024)
	vec = squareVec(vec)
	sum := sumVec(vec)
	fmt.Printf("Sum: %d\n", sum)
	/* End of code and function calls to profile */

	/***** Start of profiling code block *****/
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
	/***** End of profiling code block *****/
}
