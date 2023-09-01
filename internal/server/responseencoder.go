package server

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	http1 "net/http"
	"strings"
)

const (
	baseContentType = "application"
)

type Reply struct {
	Code    int          `json:"code"`
	Result  *interface{} `json:"result,omitempty"`
	Message string       `json:"Message"`
	Reason  string       `json:"reason"`
}

// ContentType returns the content-type with base prefix.
func ContentType(subtype string) string {
	return strings.Join([]string{baseContentType, subtype}, "/")
}

// ResponseEncoder encodes the object to the HTTP response.
func ResponseEncoder(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if v == nil {
		return nil
	}
	if rd, ok := v.(http.Redirector); ok {
		url, code := rd.Redirect()
		http1.Redirect(w, r, url, code)
		return nil
	}
	codec, _ := http.CodecForRequest(r, "Accept")

	item := &Reply{
		Result:  &v,
		Message: "ok",
	}
	data, err := codec.Marshal(item)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", ContentType(codec.Name()))
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}

// ErrorEncoder encodes the error to the HTTP response.
func ErrorEncoder(w http.ResponseWriter, r *http.Request, err error) {
	se := errors.FromError(err)
	codec, _ := http.CodecForRequest(r, "Accept")

	errBody := &Reply{
		Code:    int(se.Code),
		Message: se.Message,
		Reason:  se.Reason,
	}

	body, err := codec.Marshal(se)
	if se.Code > 600 {
		w.WriteHeader(http1.StatusOK)
		body, err = codec.Marshal(errBody)
	} else {
		w.WriteHeader(int(se.Code))
	}
	if err != nil {
		w.WriteHeader(http1.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", ContentType(codec.Name()))

	_, _ = w.Write(body)
}
