package main

import (
	"fmt"
)

type File struct {
	text string
}

func (f *File) read() {
	fmt.Println(f.text)
}
