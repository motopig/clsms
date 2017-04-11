package blue

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func Send(Mobile string, Content string) (error, int) {
	MERGEURL := GETURL + "?account=" + ACCOUNT + "&pswd=" + PASSWORD + "&mobile=%s&msg=%s"
	strUrl := fmt.Sprintf(MERGEURL, Mobile, Content)
	failed := 1
	success := 0
	r, err := http.NewRequest("GET", strUrl, nil)
	if err != nil {
		return err, failed
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err, failed
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		// 超过运营商条数限制
		return errors.New("resp.StatusCode!=http.StatusOK: " + strconv.Itoa(resp.StatusCode)), failed
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		return err, failed
	}

	ret := strings.Split(string(data), ",")

	if len(ret[1]) > int(1) {
		code, _ := strconv.Atoi(ret[1])
		return errors.New(msgs[code]), failed
	}

	return nil, success
}
