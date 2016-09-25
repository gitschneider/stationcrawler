package stationcrawler

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
	"strconv"
	"errors"
	"bytes"
)

func NoSongNowError() error {
	return errors.New("No Song is played at the moment!")
}

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

	// There could be a response with BOM, which is actually neither allowed
	// nor can be parsed. Trimming is totally fine and needed
	bodyBytes = bytes.TrimPrefix(bodyBytes, []byte{239, 187, 191})

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