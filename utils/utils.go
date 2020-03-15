package utils

import (
	"4rChd1Ts/config"
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
)

func DicFileToSlice(srcFilePath string, url string) ([]string, error) {
	result := GenerateByUrl(url)
	file, err := os.Open(srcFilePath)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		result = append(result, url+line)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}

	return result, nil
}

func GenerateByUrl(url string) []string {
	reg := regexp.MustCompile(`[a-zA-Z0-9]+`)
	sub := reg.FindAllString(url, -1)
	var dic []string
	for _, value1 := range sub {
		for _, value2 := range config.DefaultBackUpName {
			dic = append(dic, url+"/"+value1+value2)
		}
	}
	return dic
}

func WriteFile(dstFilePath string, contents string) error {
	f, err := os.OpenFile(dstFilePath, os.O_WRONLY, 0644)
	if err != nil {
		return err
	} else {
		n, _ := f.Seek(0, 2)
		_, err = f.WriteAt([]byte(contents), n)
	}
	defer f.Close()
	return nil
}
