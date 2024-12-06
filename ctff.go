package main

import (
	ccode "github.com/jurgen-kluft/ccode/ccode-fixer"
	cpkg "github.com/jurgen-kluft/ctff/package"
)

func main() {
	if ccode.Init() {
		pkg := cpkg.GetPackage()
		ccode.GenerateFiles(pkg)
		ccode.Generate(pkg, false, true)
	}
}
