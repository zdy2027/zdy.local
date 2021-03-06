package weedOP

import (
	"io"
	"net/http"
	"encoding/json"
	"errors"
	"io/ioutil"
	"fmt"
	"bytes"
	"mime/multipart"
	"net/textproto"
	"strings"
	"net/url"
	"os"
	"github.com/Kenshin/curl"
	"github.com/ddliu/go-httpclient"
)

type uploadResp struct {
	Fid      string
	FileName string
	FileUrl  string
	Size     int64
	Error    string
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func createFormFile(writer *multipart.Writer, fieldname, filename, mime string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fieldname), escapeQuotes(filename)))
	if len(mime) == 0 {
		mime = "application/octet-stream"
	}
	h.Set("Content-Type", mime)
	return writer.CreatePart(h)
}

func makeFormData(filename, mimeType string, content io.Reader) (formData io.Reader, contentType string, err error) {
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	part, err := createFormFile(writer, "file", filename, mimeType)
	//log.Println(filename, mimeType)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(part, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	formData = buf
	contentType = writer.FormDataContentType()
	//log.Println(contentType)
	writer.Close()

	return
}


func upload(url string, contentType string, formData io.Reader) (r *uploadResp, err error) {
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}
	resp, err := client.Post(url, contentType, formData)
	if err != nil{
		return
	}
	defer resp.Body.Close()
	upload := new(uploadResp)
	if err = decodeJson(resp.Body, upload); err != nil {
		return
	}

	if upload.Error != "" {
		err = errors.New(upload.Error)
		return
	}

	r = upload

	return
}

func decodeJson(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func del(url string) error {
	client := http.Client{}
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		txt, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(txt))
	}
	return err
}

func Upload(filename, mimeType string, file io.Reader, args string) (fid string, size int64, err error) {
	data, contentType, err := makeFormData(filename, mimeType, file)
	if err != nil {
		return
	}

	u := url.URL{
		Scheme: "http",
		Host:   args,
		Path:   "/submit",
		RawQuery: url.Values{}.Encode(),
	}

	resp, err := upload(u.String(), contentType, data)
	//resp, err := upload(u.String(), mimeType, file)
	if err == nil {
		fid = resp.Fid
		size = resp.Size
	}

	return
}

func UploadFile(url,filename string)(fid string, size int64, err error){
	client := httpclient.NewHttpClient()
	res, err := client.WithHeader("Connection", "close").WithOption(httpclient.OPT_TIMEOUT, 30).Post("http://"+url+"/submit", map[string]string {
		"@file": filename,
	})
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic error")
			fmt.Println(err)
		}
	}()
	if res != nil {
		defer res.Body.Close()
	}
	if httpclient.IsTimeoutError(err) {
		return "",0,errors.New("timeouterror")
	}
	upload := new(uploadResp)
	if err = decodeJson(res.Body, upload); err != nil {
		return "",0,errors.New("decodeJson error")
	}
	if upload.Error != "" {
		err = errors.New(upload.Error)
		return "",0,err
	}

	fid = upload.Fid
	size = upload.Size
	err = nil
	return
}

func DownloadWeed(url, filename string) error{
	code, res, err := curl.Get(url)
	if code != 0 {
		return err
	}

	// close
	defer res.Body.Close()
	if res.StatusCode == 200 {
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		defer file.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		s := buf.String()
		file.WriteString(s)
		return nil
	}else {
		return errors.New("status code not 200")
	}
}