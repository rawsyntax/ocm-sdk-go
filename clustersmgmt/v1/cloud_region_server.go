/*
Copyright (c) 2019 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/openshift-online/ocm-sdk-go/errors"
)

// CloudRegionServer represents the interface the manages the 'cloud_region' resource.
type CloudRegionServer interface {

	// Get handles a request for the 'get' method.
	//
	// Retrieves the details of the region.
	Get(ctx context.Context, request *CloudRegionGetServerRequest, response *CloudRegionGetServerResponse) error
}

// CloudRegionGetServerRequest is the request for the 'get' method.
type CloudRegionGetServerRequest struct {
}

// CloudRegionGetServerResponse is the response for the 'get' method.
type CloudRegionGetServerResponse struct {
	status int
	err    *errors.Error
	body   *CloudRegion
}

// Body sets the value of the 'body' parameter.
//
//
func (r *CloudRegionGetServerResponse) Body(value *CloudRegion) *CloudRegionGetServerResponse {
	r.body = value
	return r
}

// SetStatusCode sets the status code for a give response and returns the response object.
func (r *CloudRegionGetServerResponse) SetStatusCode(status int) *CloudRegionGetServerResponse {
	r.status = status
	return r
}

// marshall is the method used internally to marshal responses for the
// 'get' method.
func (r *CloudRegionGetServerResponse) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// CloudRegionServerAdapter represents the structs that adapts Requests and Response to internal
// structs.
type CloudRegionServerAdapter struct {
	server CloudRegionServer
	router *mux.Router
}

func NewCloudRegionServerAdapter(server CloudRegionServer, router *mux.Router) *CloudRegionServerAdapter {
	adapter := new(CloudRegionServerAdapter)
	adapter.server = server
	adapter.router = router
	adapter.router.Methods("GET").Path("").HandlerFunc(adapter.getHandler)
	return adapter
}
func (a *CloudRegionServerAdapter) readCloudRegionGetServerRequest(r *http.Request) (*CloudRegionGetServerRequest, error) {
	var err error
	result := new(CloudRegionGetServerRequest)
	return result, err
}
func (a *CloudRegionServerAdapter) writeCloudRegionGetServerResponse(w http.ResponseWriter, r *CloudRegionGetServerResponse) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.status)
	err := r.marshal(w)
	if err != nil {
		return err
	}
	return nil
}
func (a *CloudRegionServerAdapter) getHandler(w http.ResponseWriter, r *http.Request) {
	req, err := a.readCloudRegionGetServerRequest(r)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to read request from client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
		return
	}
	resp := new(CloudRegionGetServerResponse)
	err = a.server.Get(r.Context(), req, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to run method Get: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
	err = a.writeCloudRegionGetServerResponse(w, resp)
	if err != nil {
		reason := fmt.Sprintf("An error occured while trying to write response for client: %v", err)
		errorBody, _ := errors.NewError().
			Reason(reason).
			ID("500").
			Build()
		errors.SendError(w, r, errorBody)
	}
}
func (a *CloudRegionServerAdapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}
