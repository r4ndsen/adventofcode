package support

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

func Split(delim byte, content []byte) [][]byte {
	result := make([][]byte, 0)

	prevIdx := 0
	for i, v := range content {
		if v == delim {
			result = append(result, content[prevIdx:i])
			prevIdx = i + 1
		}

		if i == len(content)-1 {
			result = append(result, content[prevIdx:i])
			break
		}
	}

	return result
}

func Trim(chr byte, content []byte) []byte {
	for i := len(content) - 1; i >= 0; i-- {
		if content[i] != chr {
			return append(content[:i+2])
		}
	}

	return content
}

func readCookie() string {
	envs, err := godotenv.Read("../.env")

	if err != nil {
		log.Fatal(err)
	}

	cookie, ok := envs["COOKIE"]
	if !ok {
		log.Fatal("cookie not found in env")
	}

	return cookie
}

func fetchInput(day int, file string) {
	log.Println("download input file for day:", day)

	url := fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", day)

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Cookie", readCookie())

	c := http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	out, err := os.Create(file)
	Check(err)

	defer out.Close()

	io.Copy(out, resp.Body)
}

func readLines(r io.ReadCloser) [][]byte {
	defer r.Close()

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	result := make([][]byte, 0)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			break
		}

		result = append(result, []byte(text))
	}

	return result
}

func GetInputFor(day int) [][]byte {
	dir := fmt.Sprintf("/tmp/adventofcode/%d/", day)
	filePath := dir + "input.txt"
	var err error
	f, err := os.Open(filePath)

	if err == nil {
		fileInfo, _ := f.Stat()

		if fileInfo.Size() == 0 {
			err = errors.New("empty file")
		}
	}

	if err != nil {
		fmt.Println(err)
		err = os.MkdirAll(dir, 0700)
		if err != nil {
			log.Fatalf("failed to create dir: %v", err)
		}

		fetchInput(day, filePath)

		return GetInputFor(day)
	}

	log.Println("read cached input file for day:", day)

	return readLines(f)
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
