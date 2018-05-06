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
		fmt.Println("Please provide template name")
		os.Exit(1)
	}

	templateName := os.Args[1]
	url := strings.Replace("https://github.com/mul14/boilerplate-%s/archive/master.zip", "%s", templateName, 1)

	fmt.Println("Fetching boilerplate from " + url)

	filename := download(url)

	exec.Command("unzip", filename).Run()

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
