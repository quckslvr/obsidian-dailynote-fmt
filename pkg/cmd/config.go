package cmd

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadConfig(path string) {
	println(isConfig(path))
}

func isConfig(path string) (found bool) {
	dir := path + ".obsidian"
	if _, e := os.Open(dir); os.IsNotExist(e) {
		found = true
	}
	return
}
