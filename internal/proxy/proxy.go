package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type ReverseProxy struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
}

func NewReverseProxy(backendurl string) (*ReverseProxy, error) {
	parsedURL, err := url.Parse(backendurl)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return nil, fmt.Errorf("backend URL must include scheme and host")
	}

	rp := &ReverseProxy{target: parsedURL}
	rp.proxy = &httputil.ReverseProxy{
		Director:       rp.director,
		ModifyResponse: rp.modifyResponse,
		ErrorHandler:   rp.errorHandler,
	}
	return rp, nil
}

func (rp *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rp.proxy.ServeHTTP(w, r)
}

func (rp *ReverseProxy) director(req *http.Request) {
	req.URL.Scheme = rp.target.Scheme
	req.URL.Host = rp.target.Host
	req.Host = rp.target.Host

	req.Header.Set("X-Gateway", "enterprise-api-gateway")
	req.Header.Set("X-Forwarded-Time", time.Now().Format(time.RFC3339))

	if clientIP := req.RemoteAddr; clientIP != "" {
		req.Header.Set("X-Forwarded-For", clientIP)
	}
	req.Header.Set("X-Forwarded-Proto", req.URL.Scheme)
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))

}

func (rp *ReverseProxy) modifyResponse(resp *http.Response) error {
	resp.Header.Set("X-Gateway", "enterprise-api-gateway")
	resp.Header.Set("X-Processed-Time", time.Now().Format(time.RFC3339))
	resp.Header.Set("X-Backend-Server", rp.target.Host)
	return nil
}

func (rp *ReverseProxy) errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	statusCode := http.StatusBadGateway
	errMsg := "Backend service unreachable/unavailable"
	if err == http.ErrHandlerTimeout {
		statusCode = http.StatusGatewayTimeout
		errMsg = "Backend service timeout"
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, `{"error":"%s","details":"%s"}`, errMsg, err.Error())
}
