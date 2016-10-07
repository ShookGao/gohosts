package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
)

func main() {

	if runtime.GOOS == "windows" {
		os.Chdir("c:/windows/system32/drivers/etc")
		done()
	} else {
		os.Chdir("/etc")
		done()
	}
}

func isExist(name string) bool {
	_, err := os.Stat(name)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func done() {
	fmt.Println("Read remote data...")
	res, err := http.Get("https://raw.githubusercontent.com/racaljk/hosts/master/hosts")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()
	hostsByte, _ := ioutil.ReadAll(res.Body)
	if !isExist("hosts.bak") {
		os.Rename("hosts", "hosts.bak")
	}
	fmt.Println("Write remote data...")
	err = ioutil.WriteFile("hosts", hostsByte, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Write done!")
}
