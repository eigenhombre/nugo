package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func goroot() string {
	return os.Getenv("GOPATH")
}

func mkdir(path string) error {
	return os.Mkdir(path, 0755)
}

func usage() string {
	return `Usage: nugo <projectname>`
}

// dirName returns the directory name of the path.
func dirName(path string) string {
	i := len(path) - 1
	for i >= 0 && path[i] == '/' {
		i--
	}
	for i >= 0 && path[i] != '/' {
		i--
	}
	return path[:i+1]
}

func writeFile(name, content string) {
	// Make parent directories as needed.
	err := os.MkdirAll(dirName(name), 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	f, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func sh(args ...string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	gr := goroot()
	if gr == "" {
		fmt.Println("GOPATH not set")
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		fmt.Println(usage())
		os.Exit(1)
	}
	proj := os.Args[1]
	path := gr + "/" + proj
	err := mkdir(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = os.Chdir(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	writeFile("./Makefile", Makefile(proj))
	year, _, _ := time.Now().Date()
	writeFile("./README.md", readme(proj, year))
	writeFile("./Dockerfile", Dockerfile(proj))
	writeFile("./main.go", main_go())
	writeFile("./main_test.go", unit_test_go())
	writeFile("./.github/workflows/build.yml", build())
	writeFile("./.gitignore", gitignore(proj))
	writeFile("./LICENSE", license(year))
	sh("go", "mod", "init", fmt.Sprintf("github.com/eigenhombre/%s", proj))
	sh("go", "mod", "tidy")
	sh("git", "init")
	f, err := os.Open(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	names, err := f.Readdirnames(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println()
	for i, name := range names {
		if i%4 == 0 {
			fmt.Print("\t")
		}
		fmt.Printf(" %12s", name)
		if i%4 == 3 {
			fmt.Print("\n")
		}
	}
	fmt.Println()
	fmt.Println("OK")
}
