package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

var (
	BIN                 string = "gomamayo"
	VERSION             string = getVersion()
	CURRENT_REVISION, _        = sh.Output("git", "rev-parse", "--short", "HEAD")
	BUILD_LDFLAGS       string = "-s -w -X main.revision=" + CURRENT_REVISION
)

// func init() {
// 	VERSION = getVersion()
// 	CURRENT_REVISION, _ = sh.Output("git", "rev-parse", "--short", "HEAD")
// }

func getVersion() string {
	_, err := exec.LookPath("gobump")
	if err != nil {
		fmt.Println("installing gobump")
		sh.Run("go", "install", "github.com/x-motemen/gobump/cmd/gobump@latest")
	}
	v, _ := sh.Output("gobump", "show", "-r", "./cmd/gomamayo/")
	return v
}

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	// mg.Deps(InstallDeps)
	fmt.Println("Building...")
	bin := BIN
	if runtime.GOOS == "windows" {
		bin += ".exe"
	}
	cmd := exec.Command("go", "build", "-trimpath", "-ldflags="+BUILD_LDFLAGS, "-o", bin, "./cmd/gomamayo/")
	return cmd.Run()
}

// A custom install step if you need your bin someplace other than go/bin
func Install() error {
	mg.Deps(Build)
	fmt.Println("Installing...")
	cmd := exec.Command("go", "install", "-ldflags="+BUILD_LDFLAGS, "./cmd/gomamayo/")
	return cmd.Run()
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	cmd := exec.Command("go", "get", "github.com/stretchr/piglatin")
	return cmd.Run()
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("goxz")
	bin := BIN
	if runtime.GOOS == "windows" {
		bin += ".exe"
	}
	os.RemoveAll(bin)
}

func ShowVersion() {
	fmt.Println(getVersion())
}

func Cross() {
	_, err := exec.LookPath("goxz")
	if err != nil {
		fmt.Println("installing goxz")
		sh.Run("go", "install", "github.com/Songmu/goxz/cmd/goxz@latest")
	}
	// v, _ := sh.Output("gobump", "show", "-r", "./cmd/gomamayo/")
	sh.Run("goxz", "-n", BIN, "-pv=v"+VERSION, "./cmd/gomamayo/")
}

func Bump() {
	_, err := exec.LookPath("gobump")
	if err != nil {
		fmt.Println("installing gobump")
		sh.Run("go", "install", "github.com/x-motemen/gobump/cmd/gobump@latest")
	}
	sh.Run("gobump", "up", "./cmd/gomamayo/")
}

func Upload() {
	_, err := exec.LookPath("ghr")
	if err != nil {
		fmt.Println("installing ghr")
		sh.Run("go", "install", "github.com/tcnksm/ghr@latest")
	}
	sh.Run("ghr", "-draft", "v"+VERSION, "goxz")
}
