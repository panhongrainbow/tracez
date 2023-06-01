package model

import (
	"fmt"
	"os/exec"
	"testing"
	"time"
)

func Test_PPROF_TEST(t *testing.T) {
	cmd := exec.Command("go", "test", "-v", "-bench=^Benchmark_Estimate_omitemptySample$", "-run=none", ".", "-benchtime=3s", "-memprofile", "mem.prof")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	// Start HTTP server for pprof
	go func() {
		cmd := exec.Command("go", "tool", "pprof", "-http=:8080", "mem.prof")
		err = cmd.Run()
		if err != nil {
			panic(err)
		}
	}()

	// Wait for HTTP server to start
	time.Sleep(time.Second)

	// Open web browser to view pprof results
	err = exec.Command("open", "http://localhost:8080").Start()
	if err != nil {
		panic(err)
	}

	// Wait for user to view pprof results
	fmt.Println("Press enter to exit...")
	time.Sleep(300 * time.Second)
}
