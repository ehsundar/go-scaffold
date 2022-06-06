package api

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

func Invoke[T, R any](next func(context.Context, *T) (*R, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(T)

		if r.Header.Get("Content-Type") == "application/json" && r.ContentLength > 0 {
			d := json.NewDecoder(r.Body)
			err := d.Decode(req)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("error 1"))
				return
			}
		}

		vars := mux.Vars(r)
		mpDec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			DecodeHook:           nil,
			ErrorUnused:          false,
			ErrorUnset:           false,
			ZeroFields:           false,
			WeaklyTypedInput:     true,
			Squash:               false,
			Metadata:             nil,
			Result:               &req,
			TagName:              "json",
			IgnoreUntaggedFields: false,
			MatchName:            nil,
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error 4"))
			return
		}

		err = mpDec.Decode(vars)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("error 5"))
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
