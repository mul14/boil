package main

import (
	"fmt"
	"io"
	"io/ioutil"
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
	var destinationPath string

	if len(os.Args) == 2 {
		fmt.Println("Extract boilerplate in current directory? [y/N]")
		var response string
		fmt.Scanln(&response)
		if response == "y" {
			destinationPath = "."
		} else {
			fmt.Println("Cancel")
			os.Exit(0)
		}
	} else {
		destinationPath = os.Args[2]
	}

	run(templateName, destinationPath)
}

func run(templateName string, destinationPath string) {
	url := strings.Replace(
		"https://github.com/mul14/boilerplate-%s/archive/master.zip",
		"%s",
		templateName,
		1,
	)

	fmt.Println("Fetching boilerplate from " + url)

	filename := download(url)

	extract(filename, destinationPath)

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

func extract(filename string, destinationPath string) {
	err := exec.Command("unzip", "-j", filename, "-d", destinationPath).Run()
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
