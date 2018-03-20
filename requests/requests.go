package requests

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"errors"
	"fmt"
)

//下载远程url内容
func DownloadString(remoteUrl string, queryValues url.Values) (body []byte, err error) {

	client := &http.Client{}
	body = nil
	uri, err := url.Parse(remoteUrl)
	if err != nil {
		return
	}
	if queryValues != nil {
		values := uri.Query()
		if values != nil {
			for k, v := range values {
				queryValues[k] = v
			}
		}
		uri.RawQuery = queryValues.Encode()
	}
	request, err := http.NewRequest("GET", uri.String(), nil)
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Add("Accept-Encoding", "gzip, deflate")
	request.Header.Add("Accept-Language", "zh-cn,zh;q=0.8,en-us;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "close")
	request.Header.Add("Host", uri.Host)
	request.Header.Add("Referer", uri.String())
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")

	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	if err != nil {
		return
	}

	if response.StatusCode == http.StatusOK {
		switch response.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(response.Body)
			for {
				buf := make([]byte, 1024)
				n, err := reader.Read(buf)

				if err != nil && err != io.EOF {
					return nil,err
				}

				if n == 0 {
					break
				}
				body = append(body, buf...)
			}
		default:
			body, _ = ioutil.ReadAll(response.Body)

		}
	}else{
		err = errors.New(fmt.Sprintf("bad status: %s", response.Status))
	}
	return
}

func DownloadAndSaveFile(remoteUrl, dstFile string) (error) {
	client := &http.Client{}
	uri, err := url.Parse(remoteUrl)
	if err != nil {
		return err
	}
	// Create the file
	out, err := os.Create(dstFile)
	if err != nil  {
		return err
	}
	defer out.Close()

	request, err := http.NewRequest("GET", uri.String(), nil)
	request.Header.Add("Connection", "close")
	request.Header.Add("Host", uri.Host)
	request.Header.Add("Referer", uri.String())
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")

	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()


	if resp.StatusCode == http.StatusOK {
		_, err = io.Copy(out, resp.Body)
	}else{
		return errors.New(fmt.Sprintf("bad status: %s", resp.Status))
	}
	return nil
}