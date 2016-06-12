package stationcrawler

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
	"strconv"
)

func readJson(url string, jsonStruct interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyBytes, &jsonStruct)
	if err != nil {
		return err
	}
	return nil
}

func parseUnixString(unix string) (time.Time, error) {
	i, err := strconv.ParseInt(unix, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(i, 0), nil
}

func isNow(start, end time.Time) bool {
	now := time.Now()
	if now.After(start) && now.Before(end) {
		return true
	}
	return false
}
