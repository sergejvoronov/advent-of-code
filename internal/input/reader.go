package input

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	inputURL     = "https://adventofcode.com/%s/day/%s/input"
	inputFile    = "%s-%s.input"
	cacheDir     = "inputs"
	cookieName   = "session"
	cookieDomain = ".adventofcode.com"
)

type (
	Reader interface {
		ReadInput(year, day string) string
	}

	reader struct {
		SessionID string
	}
)

func NewReader(sessionID string) Reader {
	return &reader{SessionID: sessionID}
}

func (r *reader) ReadInput(year, day string) string {
	bytes, err := r.getInputData(r.SessionID, year, day)
	if err != nil {
		return fmt.Sprintf("Error reading input file: %s\n", err)
	}

	return string(bytes)
}

func (r *reader) getInputData(sessionID, year, day string) ([]byte, error) {
	filename := fmt.Sprintf(inputFile, year, day)
	inputFile := filepath.Join(cacheDir, filename)
	if _, err := os.Stat(inputFile); err != nil && os.IsNotExist(err) {
		req, err := r.createInputReq(sessionID, year, day)
		if err != nil {
			return nil, err
		}

		body, err := r.fetch(req)
		if err != nil {
			return nil, err
		}

		if err := ioutil.WriteFile(inputFile, body, 0644); err != nil {
			return nil, err
		}

		return body, nil
	}

	return ioutil.ReadFile(inputFile)
}

func (r *reader) createInputReq(sessionID, year, day string) (*http.Request, error) {
	url := fmt.Sprintf(inputURL, year, day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{
		Name:   cookieName,
		Value:  sessionID,
		Domain: cookieDomain,
		Path:   "/",
	})

	return req, nil
}

func (r *reader) fetch(req *http.Request) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
