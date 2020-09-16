package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	// path := "./app"
	// fileinfo, err := ioutil.ReadDir(path)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// if len(fileinfo) != 0 {
	// 	for _, v := range fileinfo {
	// 		filename := v.Name()
	// 		// if filename == "main.go"||filename=="build.exe"||filename=="__debug_bin" {
	// 		// 	continue
	// 		// }
	// 		if v.IsDir() {
	// 			err = os.RemoveAll(filepath.Join(path, filename))
	// 			if err != nil {
	// 				log.Fatalln(err)
	// 			}
	// 		} else {
	// 			err = os.Remove(filepath.Join(path, filename))
	// 			if err != nil {
	// 				log.Fatalln(err)
	// 			}
	// 		}
	// 	}
	// }
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("sh", "go", "build", `-tags=jsoniter`, `-ldflags="-s -w"`, "-o", "./app/app").Run()
	case "windows":
		//go build `-tags=jsoniter` -ldflags="-s -w" -o "./build/app/app.exe"
		err = exec.Command("powershell", "go", "build", `-tags=jsoniter`, `-ldflags="-s -w"`, "-o", "../build/app/app.exe").Run()
		//err = exec.Command("go", "build", `-ldflags="-s -w"`, "-o", "./build/app.exe").Run()
	default:
		err = fmt.Errorf("GOOS err")
	}
	if err != nil {
		log.Fatalln(err)
	}

	// createLog()
	// createConfig()
}

func createLog() {
	err := os.MkdirAll("./app/log/gin_log", os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = os.Create("./app/log/gin_log/system.log")
	if err != nil {
		log.Fatalln(err)
	}
}

func createConfig() {
	file, err := os.OpenFile("../config/config.json", 2, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	os.Mkdir("./app/config", os.ModePerm)
	json, err := os.Create("./app/config/config.json")
	if err != nil {
		log.Fatalln(err)
	}
	io.Copy(json, file)
}
