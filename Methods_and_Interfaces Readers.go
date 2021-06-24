package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (my MyReader) Read(b []byte) (int, error) {
	for i:=0;i<len(b);i++{
		b[i]=byte('A')
	}
	return 1, nil
}
func main() {
	reader.Validate(MyReader{})
}
