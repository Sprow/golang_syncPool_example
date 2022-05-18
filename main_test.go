package main

import (
	"bytes"
	"fmt"
	pool "syncPool/joinstrings"

	"testing"
)

func BenchmarkWithPool1000(b *testing.B) {
	book, err := getText("text_1000_strings.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	work := func() {
		buf := pool.GetBuffer()
		defer pool.PutBuffer(buf)
		_, err = JoinString(buf, " ", book...)
		if err != nil {
			b.Error(err)
			return
		}
	}
	for i := 0; i < b.N; i++ {
		work()
	}
}

func BenchmarkWithoutPool1000(b *testing.B) {
	book, err := getText("text_1000_strings.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	work := func() {
		buf := &bytes.Buffer{}

		_, err = JoinString(buf, " ", book...)
		if err != nil {
			b.Error(err)
			return
		}
	}
	for i := 0; i < b.N; i++ {
		work()
	}
}

func BenchmarkWithPool10000(b *testing.B) {
	book, err := getText("text_10000_strings.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	work := func() {
		buf := pool.GetBuffer()
		defer pool.PutBuffer(buf)
		_, err = JoinString(buf, " ", book...)
		if err != nil {
			b.Error(err)
			return
		}
	}
	for i := 0; i < b.N; i++ {
		work()
	}
}

func BenchmarkWithoutPool10000(b *testing.B) {
	book, err := getText("text_10000_strings.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	work := func() {
		buf := &bytes.Buffer{}

		_, err = JoinString(buf, " ", book...)
		if err != nil {
			b.Error(err)
			return
		}
	}
	for i := 0; i < b.N; i++ {
		work()
	}
}
