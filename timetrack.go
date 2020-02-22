//Package utif stands for Utility Functions, providing different functions to facilitate common tasks
package utif

import (
	"log"
	"regexp"
	"runtime"
	"time"
)

var _extractFnName = regexp.MustCompile(`^.*\.(.*)$`)

//TimeTrack allow to time given function by passing it start time. Caller function name will be extacted from runtime.
func TimeTrack(start time.Time) {
	elapsed := time.Since(start)

	// Skip this function, and fetch the PC and file for its parent
	pc, _, _, _ := runtime.Caller(1)

	// Retrieve a Function object this functions parent
	functionObject := runtime.FuncForPC(pc)

	// Regex to extract just the function name (and not the module path)

	name := _extractFnName.ReplaceAllString(functionObject.Name(), "$1")

	log.Printf("[timetrack] %s took %s\n", name, elapsed)
}

//Track given time elapsed since start time. fnname is name of function to print in log.
func Track(start time.Time, fnname string) {
	elapsed := time.Since(start)
	log.Printf("[timetrack] %s took %s\n", fnname, elapsed)
}
