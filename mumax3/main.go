package main

import "github.com/mumax/3/mainpkg"

// redirected to mainpkg to allow go input files
// (these will call mainpkg.Init, but not mainpkg.RunFiles)
func main() {
	mainpkg.Main()
}
