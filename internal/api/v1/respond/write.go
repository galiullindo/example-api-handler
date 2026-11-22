package respond

import "net/http"

func Write(w http.ResponseWriter, body []byte, code int) (int, error) {
	h := w.Header()
	h.Set("Content-Type", "application/json; charset=utf-8")
	h.Set("X-Content-Type-Options", "nosniff")

	w.WriteHeader(code)
	return w.Write(body)
}
