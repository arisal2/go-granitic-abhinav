package main

import gy "github.com/graniticio/granitic-yaml/v2"
import "go-granitic-abhinav/bindings"

func main() {
	gy.StartGraniticWithYaml(bindings.Components())
}
