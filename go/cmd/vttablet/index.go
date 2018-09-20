/*
Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"net/http"

	"vitess.io/vitess/go/vt/vtctl"
)

const (
	apiPrefix = "/api/"

	jsonContentType = "application/json; charset=utf-8"
)

// This is a separate file so it can be selectively included/excluded from
// builds to opt in/out of the redirect.

func init() {
	// Anything unrecognized gets redirected to the status page.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/debug/status", http.StatusFound)
	})
	http.HandleFunc(apiPrefix+"/health-check", func(w http.ResponseWriter, r *http.Request) {
		func() error {
			// JSON encode response.
			data, err := vtctl.MarshalJSON("ok")
			if err != nil {
				return fmt.Errorf("cannot marshal data: %v", err)
			}
			w.Header().Set("Content-Type", jsonContentType)
			w.Write(data)
			return nil
		}()
	})
}
