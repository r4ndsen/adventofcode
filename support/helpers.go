package support

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
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

	_, err = io.Copy(out, resp.Body)
	Check(err)
}

func readLines(r io.ReadCloser) [][]byte {
	buf := bytes.NewBuffer(nil)

	_, err := io.Copy(buf, r)
	Check(err)
	defer r.Close()

	return bytes.Split(buf.Bytes(), []byte("\n"))
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

	log.Printf("read cached input file for day: %v - %s\n", day, filePath)

	return readLines(f)
}

func IsInt(input string) bool {
	if len(input) == 0 {
		return false
	}

	_, err := strconv.Atoi(input)
	return err == nil
}

func ToInt(input string) int {
	val, err := strconv.Atoi(input)
	Check(err)
	return val
}

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
