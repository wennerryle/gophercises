package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/goccy/go-yaml"
)

const (
	MIME_JSON = "application/json"
	MIME_YAML = "application/yaml"
)

func parseFormat(w http.ResponseWriter, req *http.Request, dst any) error {
	mediaType := strings.TrimSpace(
		strings.Split(
			req.Header.Get("Content-Type"),
			";",
		)[0],
	)

	switch mediaType {
	case MIME_JSON:
		req.Body = http.MaxBytesReader(w, req.Body, mb2b(1))
		return json.NewDecoder(req.Body).Decode(dst)
	case MIME_YAML:
		req.Body = http.MaxBytesReader(w, req.Body, mb2b(1))
		return yaml.NewDecoder(req.Body).Decode(dst)
	default:
		return fmt.Errorf("unexpected Content-Type: %s", mediaType)
	}
}
