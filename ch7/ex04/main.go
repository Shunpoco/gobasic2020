package main

type MyReader interface {
	Read(p []byte) (n int, err error)
}
