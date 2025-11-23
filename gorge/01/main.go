package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	logs := []string{
		"GET /home 200",
		"GET /products 200",
		"POST /products 201",
		"GET /home 200",
		"GET /about 404",
		"GET /home 500",
		"GET /products 200",
		"GET /products 200",
		"GET /about 404",
	}

	paths := map[string]int{}
	statusCodes := map[string]int{}

	for _, value := range logs {
		logElements := strings.Split(value, " ")
		path := logElements[1]
		statusCode, error := strconv.Atoi(logElements[2])
		if error != nil {
			continue
		}

		_, pok := paths[path]
		if pok {
			paths[path] += 1
		} else {
			paths[path] = 1
		}

		strStatusCode := strconv.Itoa(statusCode)
		_, sok := statusCodes[strStatusCode]
		if sok {
			statusCodes[strStatusCode] += 1
		} else {
			statusCodes[strStatusCode] = 1
		}
	}

	maxPath := ""
	maxPathCount := 0
	for key, value := range paths {
		if maxPathCount < value {
			maxPath = key
			maxPathCount = value
		}
	}

	fmt.Println(statusCodes)
	fmt.Println(paths)
	fmt.Printf("一番アクセスが多いパスは「%s」で%d回\n", maxPath, maxPathCount)
}
