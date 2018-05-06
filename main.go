package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Template name is required")
		fmt.Println("Usage: " + green("boil <name>"))
		os.Exit(1)
	}

	templateName := os.Args[1]
	url := strings.Replace("https://github.com/mul14/boilerplate-%s/archive/master.zip", "%s", templateName, 1)

	fmt.Println("Fetching boilerplate from " + url)

	filename := download(url)

	exec.Command("unzip", filename).Run()

	os.Remove(filename)

	fmt.Println("Done")
}

func download(url string) string {
	file, err := ioutil.TempFile("", "")
	if err != nil {
		log.Fatal(err.Error())
	}

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer response.Body.Close()

	io.Copy(file, response.Body)

	return file.Name()
}

func green(string string) string {
	return "\x1b[32;1m" + string + "\x1b[0m"
}
