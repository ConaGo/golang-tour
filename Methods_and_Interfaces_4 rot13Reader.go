package main

import (
	"io"
	"os"
	"strings"
	//"fmt"
)

type rot13Reader struct {
	r io.Reader
}
func (me *rot13Reader) Read(b []byte) (int, error) {
	n, err := me.r.Read(b)
	for i:=0; i < n; i++ {
		c := &b[i]
		isCapital := *c >= 'A' && *c <= 'Z'
		if !isCapital && (*c < 'a' || *c > 'z') {
			//Spaces -> do nothing
		} else {
			b[i] += 13
			if isCapital && *c > 'Z' || !isCapital && *c > 'z' {
				b[i] -= 26 // Overflow -> go back len(alphabet) chars
			}
		}
	}
	//fmt.Println(n)
	return n, err
}
