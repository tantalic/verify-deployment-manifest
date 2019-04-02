package main

import (
	"html"
	"net/http"
	"net/url"

	"github.com/toolhouse/deployment-manifest/pkg/deployment"

	badge "github.com/narqo/go-badge"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	http.ListenAndServe(":8000", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {

	u, err := manifestURLFromReq(r)
	if err != nil {
		errorResp(w, "error", "invalid url")
		return
	}

	env := environmentURLFromReq(r)

	manifest, err := deployment.FetchManifest(u)
	if err != nil {
		errorResp(w, env, "no manifest")
		return
	}

	validResp(w, env, manifest.Version())
}

func manifestURLFromReq(r *http.Request) (string, error) {
	u := "https:/" + r.URL.Path
	_, err := url.ParseRequestURI(u)

	return u, err
}

func environmentURLFromReq(r *http.Request) string {
	env := r.URL.Query().Get("env")
	if env == "" {
		return "deployed"
	}

	return env
}

func errorResp(w http.ResponseWriter, label, value string) {
	resp(w, label, value, badge.ColorYellow)
}

func validResp(w http.ResponseWriter, env, value string) {
	resp(w, env, value, badge.ColorBlue)
}

func resp(w http.ResponseWriter, label, value string, color badge.Color) {
	w.Header().Set("Content-Type", "image/svg+xml")
	badge.Render(html.EscapeString(label), html.EscapeString(value), color, w)
}
