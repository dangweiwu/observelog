package collect

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"observelog/internal/context"
	"strings"
)

type RepBody struct {
	Code   int `json:"code"`
	Status []struct {
		Name       string `json:"name"`
		Successful int    `json:"successful"`
		Failed     int    `json:"failed"`
	} `json:"status"`
}

type ReqObserve struct {
	ctx *context.AppCtx
}

func NewReqObserve(ctx *context.AppCtx) *ReqObserve {
	return &ReqObserve{ctx}
}

func (this ReqObserve) Send(data string) (oknum, failed int, err error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/%s/%s/_multi", this.ctx.Config.Observe.Host, this.ctx.Config.Observe.Org, this.ctx.Config.Observe.Stream), strings.NewReader(data))
	if err != nil {
		return 0, 0, err
	}
	req.SetBasicAuth(this.ctx.Config.Observe.Username, this.ctx.Config.Observe.Password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return 0, 0, errors.New(fmt.Sprintf("%d", resp.StatusCode))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}
	var rp RepBody
	if err := json.Unmarshal(body, &rp); err != nil {
		return 0, 0, err
	}

	//fmt.Println("resbody:", string(body))
	return rp.Status[0].Successful, rp.Status[0].Failed, nil
}
