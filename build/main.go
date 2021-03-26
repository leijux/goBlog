package main

import (
	"crypto/md5"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

const (
	win = "windows"
	lin = "linux"

	GOOS = win
)

var AppFs afero.Fs

func main() {
	AppFs = afero.NewOsFs()
	var err error
	switch GOOS {
	case lin:
		err = linuxBuild()
	case win:
		err = winBuild()
	default:
		log.Fatalln("GOOS err")
	}
	if err != nil {
		log.Printf("%+v", err)
	}
	createLog()
	createConfig()
}

func winBuild() error {
	gox := exec.Command("gox", "-os", "windows", "-arch", "amd64")
	err := gox.Run()
	if err != nil {
		return errors.Wrap(err, "winBuild")
	}
	return nil
	//strbuf := strings.NewReader(`gox -os "windows" -arch amd64`)
	//pw := exec.Command("powershell")
	//pw.Stdin = strbuf
	//if err := pw.Run(); err != nil {
	//	return errors.Wrap(err, "win build err")
	//}
	//return nil
}

func linuxBuild() error {
	strbuf := strings.NewReader(`gox -os "linux" -arch amd64`)
	pw := exec.Command("powershell")
	pw.Stdin = strbuf
	if err := pw.Run(); err != nil {
		return errors.Wrap(err, "linux build err")
	}
	return nil
}

func createLog() {
	_, err := AppFs.Stat("./build/app/log/gin_log/")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			_ = AppFs.MkdirAll("./build/app/log/gin_log/", os.ModePerm)
		}
	}
}

const configPath = "./build/app/config/"

func createConfig() {
	_, err := AppFs.Stat(filepath.Join(configPath, "config.json"))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			_ = AppFs.MkdirAll(configPath, os.ModePerm)
			file, err := AppFs.Create(filepath.Join(configPath, "config.json"))
			if err != nil {
				log.Println(err)
			}
			defer file.Close()
			filejson, err := AppFs.Open("./config/config.json")
			if err != nil {
				log.Println(err)
			}
			defer filejson.Close()
			_, _ = io.Copy(file, filejson)
			return
		}
	}

	file, err := AppFs.Open(filepath.Join(configPath, "config.json"))
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	filejson, err := AppFs.Open("./config/config.json")
	if err != nil {
		log.Println(err)
	}
	defer filejson.Close()
	if getMd5(file) != getMd5(filejson) {
		_, _ = io.Copy(file, filejson)
	}
}
func getMd5(file io.Reader) string {
	md5h := md5.New()
	_, _ = io.Copy(md5h, file)
	return string(md5h.Sum(nil)) //md5
}
