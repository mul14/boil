package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"io"
	"os/exec"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Template name is required")
		fmt.Println("Usage: " + green("boil <name>"))
		os.Exit(1)
	}

	run(os.Args[1])
}

func run(templateName string) {
	url := strings.Replace(
		"https://github.com/mul14/%s/archive/master.zip",
		"%s",
		templateName,
		1,
	)

	fmt.Println("Fetching boilerplate from " + url)

	filename := download(url)

	extract(filename)

	os.Remove(filename)
}

func download(url string) string {
	file, err := ioutil.TempFile("", "")
	if err != nil {
		fmt.Println(err.Error())
	}

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	if response.StatusCode >= 200 && response.StatusCode <= 299 {
		defer response.Body.Close()
		io.Copy(file, response.Body)
		fmt.Println(green("Boilerplate has been downloaded"))
	} else {
		fmt.Println(red("Boilerplate not found"))
		os.Exit(1)
	}

	return file.Name()
}

func extract(filename string) {
	err := exec.Command("unzip", filename).Run()
	if err != nil {
		printError(err.Error())
		os.Exit(1)
	}
	fmt.Println(green("Boilerplate has been extracted"))
}

func green(string string) string {
	return "\x1b[32;1m" + string + "\x1b[0m"
}

func red(string string) string {
	return "\x1b[31;1m" + string + "\x1b[0m"
}

func printError(message string) {
	fmt.Println(red(message))
}
