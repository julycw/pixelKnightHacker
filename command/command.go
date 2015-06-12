package command

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	serverUrl = "http://s1.3dmmxwk.u77.com/service/main.ashx"
)

func send(serverUrl, sid, userid, target string, params map[string]string) ([]byte, error) {
	v := url.Values{}
	v.Set("sid", sid)
	v.Set("t", target)
	v.Set("userid", userid)
	if params != nil {
		for key, value := range params {
			v.Set(key, value)
		}
	}

	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	if req, err := http.NewRequest("POST", serverUrl, body); err != nil {
		return []byte{}, err
	} else {
		req.Header.Set("Accept", "*/*")
		// req.Header.Set("Accept-Encoding", "gzip, deflate")
		req.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
		req.Header.Set("Connection", "keep-alive")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Cookie", "ASP.NET_SessionId=xxxzpi3mxxz21wmzquznlzlc")
		req.Header.Set("RA-Sid", "7AE3CEA6-20150105-053821-b1e0d6-379bf4")
		req.Header.Set("RA-Ver", "2.10.4")
		req.Header.Set("Referer", "http://res.mxwk.90tank.com/Res1024/BraveRistOfficial.swf?2.2.12")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.118 Safari/537.36")

		resp, err := client.Do(req)
		if err != nil { //发送
			return []byte{}, err
		}
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		return data, nil
	}

	return []byte{}, nil
}
