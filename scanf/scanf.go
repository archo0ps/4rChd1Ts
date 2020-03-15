package scanf

import (
	"4rChd1Ts/config"
	"4rChd1Ts/help"
	"4rChd1Ts/utils"
	"4rChd1Ts/utils/Pool"
	"github.com/fatih/color"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func HttpGet(url string) error {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	if !config.RandomUserAgent {
		req.Header.Set("User-Agent", config.DefaultUserAgent)
	} else {
		rand.Seed(time.Now().UnixNano())
		req.Header.Set("User-Agent", config.UserAgents[rand.Intn(len(config.UserAgents))])
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode == http.StatusOK ||
		resp.StatusCode == http.StatusFound {
		color.Green("[%d]:%s", resp.StatusCode, url)
	} else if resp.StatusCode == http.StatusForbidden ||
		resp.StatusCode == http.StatusServiceUnavailable {
		color.Yellow("[%d]:%s", resp.StatusCode, url)
	}
	if help.O != "" && resp.StatusCode != http.StatusNotFound {
		err = utils.WriteFile(help.O, "["+strconv.Itoa(resp.StatusCode)+"]"+url+"\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func StartScanf(url string, dicPath string) error {
	if url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}
	if help.O != "" {
		_, err := os.Stat(help.O)
		if err != nil {
			file, err := os.Create(help.O)
			if err != nil {
				return err
			}
			file.Close()
		}
	}

	dic, err := utils.DicFileToSlice(dicPath, url)
	if err != nil {
		return err
	}

	p := Pool.NewPool(config.DefaultMaxWorker)

	go func() {
		for _, value := range dic {
			t := Pool.NewTask(HttpGet, value)
			p.EntryChannel <- t
		}
	}()

	p.Run()
	return nil
}
