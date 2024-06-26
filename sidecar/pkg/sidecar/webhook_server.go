// This file is part of MinIO Operator
// Copyright (c) 2024 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package sidecar

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/minio/operator/pkg/common"

	"github.com/gorilla/mux"
)

// Used for registering with rest handlers (have a look at registerStorageRESTHandlers for usage example)
// If it is passed ["aaaa", "bbbb"], it returns ["aaaa", "{aaaa:.*}", "bbbb", "{bbbb:.*}"]
func restQueries(keys ...string) []string {
	var accumulator []string
	for _, key := range keys {
		accumulator = append(accumulator, key, "{"+key+":.*}")
	}
	return accumulator
}

func configureWebhookServer(c *Controller) *http.Server {
	router := mux.NewRouter().SkipClean(true).UseEncodedPath()

	router.Methods(http.MethodPost).
		Path(common.WebhookAPIBucketService + "/{namespace}/{name:.+}").
		HandlerFunc(c.BucketSrvHandler).
		Queries(restQueries("bucket")...)

	router.NotFoundHandler = http.NotFoundHandler()

	s := &http.Server{
		Addr:           "127.0.0.1:" + common.WebhookDefaultPort,
		Handler:        router,
		ReadTimeout:    time.Minute,
		WriteTimeout:   time.Minute,
		MaxHeaderBytes: 1 << 20,
	}

	return s
}

func configureProbesServer(c *Controller, tenantTLS bool) *http.Server {
	router := mux.NewRouter().SkipClean(true).UseEncodedPath()

	router.Methods(http.MethodGet).
		Path("/ready").
		HandlerFunc(readinessHandler(tenantTLS))

	router.NotFoundHandler = http.NotFoundHandler()

	s := &http.Server{
		Addr:           "0.0.0.0:4444",
		Handler:        router,
		ReadTimeout:    time.Minute,
		WriteTimeout:   time.Minute,
		MaxHeaderBytes: 1 << 20,
	}

	return s
}

func readinessHandler(tenantTLS bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		schema := "https"
		if !tenantTLS {
			schema = "http"
		}
		// we only check against the local instance of MinIO
		url := schema + "://localhost:9000"
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to create request: %s", err), http.StatusInternalServerError)
			return
		}

		// we do insecure skip verify because we are checking against the local instance and don't care for the
		// certificate
		client := &http.Client{
			Timeout: time.Millisecond * 500,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}

		response, err := client.Do(request)
		if err != nil {
			http.Error(w, fmt.Sprintf("HTTP request failed: %s", err), http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()

		if response.StatusCode == 403 {
			fmt.Fprintln(w, "Readiness probe succeeded.")
		} else {
			http.Error(w, fmt.Sprintf("Readiness probe failed: expected status 403, got %d", response.StatusCode), http.StatusServiceUnavailable)
		}
	}
}
