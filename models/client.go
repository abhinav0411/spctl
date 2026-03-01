package models

import "net/http"

type Client struct {
	Session    *Session
	HTTPClient *http.Client
}
