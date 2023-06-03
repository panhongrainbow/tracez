package model

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
	"testing"
	"time"
)

var count int

func Test_Html_TEST(t *testing.T) {
	// 註冊路由處理函式
	http.HandleFunc("/", myHandler)
	http.HandleFunc("/test", myHandler2)

	// 啟動 server
	http.ListenAndServe(":3000", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	t, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 執行模板
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func myHandler2(w http.ResponseWriter, r *http.Request) {
	count++
	fmt.Println("Button clicked!")
}

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
