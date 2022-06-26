package utils

import (
	"compress/flate"
	"compress/gzip"
	"github.com/google/brotli/go/cbrotli"
	"io/ioutil"
	"net/http"
)

func SetHeaders(request *http.Request, m *map[string]string) {
	for k, v := range *m {
		request.Header.Set(k, v)
	}
}


func GetEncoding(resp *http.Response) (encoding string) {
	if encodings, ok := resp.Header["Content-Encoding"]; !ok {
		return
	} else if len(encodings) > 0 {
		return encodings[0]
	}
	return
}


type ReaderType interface {
	Close() error
	Read(p []byte) (n int, err error)
}

func GetDataBytes(resp *http.Response, encoding string) (data []byte, err error) {
	var reader ReaderType = resp.Body
	if resp.Uncompressed || encoding != "" {
		switch encoding {
		case "br":
			reader = cbrotli.NewReader(reader)
		case "gzip":
			reader, err = gzip.NewReader(reader)
		case "deflate":
			reader = flate.NewReader(reader)
		}
		if err != nil {
			return
		}
		defer reader.Close()
	}
	data, err = ioutil.ReadAll(reader)
	return
}
