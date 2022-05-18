package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	pool "syncPool/joinstrings"
)

func main() {
	book, err := getText("text_1000_strings.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := pool.GetBuffer()
	defer pool.PutBuffer(buf)
	res, err := JoinString(buf, " | ", book...)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

}

func JoinString(buf *bytes.Buffer, splitBy string, strings ...string) (res string, err error) {
	if _, err := buf.WriteString(strings[0]); err != nil { // first string
		return "", err
	}
	for _, str := range strings[1:] { // 2+ strings
		if _, err := buf.WriteString(splitBy); err != nil { // split strings by splitBy variable
			return "", err
		}
		if _, err := buf.WriteString(str); err != nil {
			return "", err
		}
	}
	res = buf.String()
	return
}

func getText(path string) ([]string, error) {
	var fileSlice []string
	file, err := os.Open(path)
	if err != nil {
		return fileSlice, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileSlice = append(fileSlice, scanner.Text())
	}
	return fileSlice, nil
}
