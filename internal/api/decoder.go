package api

import (
	"context"
	"encoding/json"
	"net/http"
)

func Invoker[T, R any](next func(context.Context, *T) (*R, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(T)
		d := json.NewDecoder(r.Body)
		err := d.Decode(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error 1"))
			return
		}

		resp, err := next(r.Context(), req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error 2"))
			return
		}

		w.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(w)
		err = enc.Encode(resp)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error 3"))
			return
		}

		return
	}
}
