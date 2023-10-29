package stringutils

import (
	"io"
	"log"
	"strconv"
	"strings"
)

func ReadCloserString(r *io.ReadCloser) string {
	bytes, err := io.ReadAll(*r)
	if err != nil {
		log.Fatalf("Error while reading string from ReadCloser: %v\n", err)
	}
	s := string(bytes)
	return s
}

func DecodeHost(h string) (address string, port int) {
	h = strings.Replace(h, "http://", "", 1)
	s := strings.Split(h, ":")
	if len(s) != 2 {
		log.Fatalf("Error decoding host to address & port: %s\n", "Invalid host string.")
	}
	address = s[0]
	port, err := strconv.Atoi(s[1])
	if err != nil {
		log.Fatalf("Error parsing string to int: %v\n", err)
	}
	return address, port
}
