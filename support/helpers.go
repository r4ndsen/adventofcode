package support

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/r4ndsen/adventofcode/cast"
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
	cookie := os.Getenv("SESSION_COOKIE")
	if cookie == "" {
		log.Fatal("SESSON_COOKIE env not set")
	}

	return cookie
}

func fetchInput(day int, file string) {
	log.Println("download input file for day:", day)

	url := fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day)

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	sessionCookie := http.Cookie{
		Name:  "session",
		Value: readCookie(),
	}
	req.AddCookie(&sessionCookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	out, err := os.Create(file)
	AssertNoError(err)

	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	AssertNoError(err)
}

func readLines(r io.ReadCloser) Input {
	buf := bytes.NewBuffer(nil)

	_, err := io.Copy(buf, r)
	AssertNoError(err)
	defer r.Close()

	return InputType(buf.Bytes())
}

func GetInput() Input {
	if day := os.Getenv("DAY"); day != "" {
		return GetInputFor(cast.ToInt(day))
	}

	log.Fatal("no day specified")
	return nil
}

func GetInputFor(day int) Input {
	dir := fmt.Sprintf("/tmp/adventofcode/2023/%d/", day)
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

func AssertNoError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
