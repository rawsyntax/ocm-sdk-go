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
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// CloudProvidersClient is the client of the 'cloud_providers' resource.
//
// Manages the collection of cloud providers.
type CloudProvidersClient struct {
	transport http.RoundTripper
	path      string
	metric    string
}

// NewCloudProvidersClient creates a new client for the 'cloud_providers'
// resource using the given transport to sned the requests and receive the
// responses.
func NewCloudProvidersClient(transport http.RoundTripper, path string, metric string) *CloudProvidersClient {
	client := new(CloudProvidersClient)
	client.transport = transport
	client.path = path
	client.metric = metric
	return client
}

// List creates a request for the 'list' method.
//
// Retrieves the list of cloud providers.
func (c *CloudProvidersClient) List() *CloudProvidersListRequest {
	request := new(CloudProvidersListRequest)
	request.transport = c.transport
	request.path = c.path
	request.metric = c.metric
	return request
}

// CloudProvider returns the target 'cloud_provider' resource for the given identifier.
//
// Returns a reference to the service that manages an specific cloud provider.
func (c *CloudProvidersClient) CloudProvider(id string) *CloudProviderClient {
	return NewCloudProviderClient(
		c.transport,
		path.Join(c.path, id),
		path.Join(c.metric, "-"),
	)
}

// CloudProvidersListRequest is the request for the 'list' method.
type CloudProvidersListRequest struct {
	transport http.RoundTripper
	path      string
	metric    string
	query     url.Values
	header    http.Header
	order     *string
	page      *int
	search    *string
	size      *int
	total     *int
}

// Parameter adds a query parameter.
func (r *CloudProvidersListRequest) Parameter(name string, value interface{}) *CloudProvidersListRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *CloudProvidersListRequest) Header(name string, value interface{}) *CloudProvidersListRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Order sets the value of the 'order' parameter.
//
// Order criteria.
//
// The syntax of this parameter is similar to the syntax of the _order by_ clause of
// a SQL statement, but using the names of the attributes of the cloud provider
// instead of the names of the columns of a table. For example, in order to sort the
// clusters descending by name identifier the value should be:
//
// [source,sql]
// ----
// name desc
// ----
//
// If the parameter isn't provided, or if the value is empty, then the order of the
// results is undefined.
func (r *CloudProvidersListRequest) Order(value string) *CloudProvidersListRequest {
	r.order = &value
	return r
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *CloudProvidersListRequest) Page(value int) *CloudProvidersListRequest {
	r.page = &value
	return r
}

// Search sets the value of the 'search' parameter.
//
// Search criteria.
//
// The syntax of this parameter is similar to the syntax of the _where_ clause of a
// SQL statement, but using the names of the attributes of the cloud provider
// instead of the names of the columns of a table. For example, in order to retrieve
// all the cloud providers with a name starting with `A` the value should be:
//
// [source,sql]
// ----
// name like 'A%'
// ----
//
// If the parameter isn't provided, or if the value is empty, then all the clusters
// that the user has permission to see will be returned.
func (r *CloudProvidersListRequest) Search(value string) *CloudProvidersListRequest {
	r.search = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *CloudProvidersListRequest) Size(value int) *CloudProvidersListRequest {
	r.size = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *CloudProvidersListRequest) Total(value int) *CloudProvidersListRequest {
	r.total = &value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *CloudProvidersListRequest) Send() (result *CloudProvidersListResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *CloudProvidersListRequest) SendContext(ctx context.Context) (result *CloudProvidersListResponse, err error) {
	query := helpers.CopyQuery(r.query)
	if r.order != nil {
		helpers.AddValue(&query, "order", *r.order)
	}
	if r.page != nil {
		helpers.AddValue(&query, "page", *r.page)
	}
	if r.search != nil {
		helpers.AddValue(&query, "search", *r.search)
	}
	if r.size != nil {
		helpers.AddValue(&query, "size", *r.size)
	}
	if r.total != nil {
		helpers.AddValue(&query, "total", *r.total)
	}
	header := helpers.SetHeader(r.header, r.metric)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: http.MethodGet,
		URL:    uri,
		Header: header,
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = new(CloudProvidersListResponse)
	result.status = response.StatusCode
	result.header = response.Header
	if result.status >= 400 {
		result.err, err = errors.UnmarshalError(response.Body)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = result.unmarshal(response.Body)
	if err != nil {
		return
	}
	return
}

// CloudProvidersListResponse is the response for the 'list' method.
type CloudProvidersListResponse struct {
	status int
	header http.Header
	err    *errors.Error
	items  *CloudProviderList
	page   *int
	size   *int
	total  *int
}

// Status returns the response status code.
func (r *CloudProvidersListResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *CloudProvidersListResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *CloudProvidersListResponse) Error() *errors.Error {
	return r.err
}

// Items returns the value of the 'items' parameter.
//
// Retrieved list of cloud providers.
func (r *CloudProvidersListResponse) Items() *CloudProviderList {
	if r == nil {
		return nil
	}
	return r.items
}

// GetItems returns the value of the 'items' parameter and
// a flag indicating if the parameter has a value.
//
// Retrieved list of cloud providers.
func (r *CloudProvidersListResponse) GetItems() (value *CloudProviderList, ok bool) {
	ok = r != nil && r.items != nil
	if ok {
		value = r.items
	}
	return
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *CloudProvidersListResponse) Page() int {
	if r != nil && r.page != nil {
		return *r.page
	}
	return 0
}

// GetPage returns the value of the 'page' parameter and
// a flag indicating if the parameter has a value.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *CloudProvidersListResponse) GetPage() (value int, ok bool) {
	ok = r != nil && r.page != nil
	if ok {
		value = *r.page
	}
	return
}

// Size returns the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *CloudProvidersListResponse) Size() int {
	if r != nil && r.size != nil {
		return *r.size
	}
	return 0
}

// GetSize returns the value of the 'size' parameter and
// a flag indicating if the parameter has a value.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *CloudProvidersListResponse) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// Total returns the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *CloudProvidersListResponse) Total() int {
	if r != nil && r.total != nil {
		return *r.total
	}
	return 0
}

// GetTotal returns the value of the 'total' parameter and
// a flag indicating if the parameter has a value.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *CloudProvidersListResponse) GetTotal() (value int, ok bool) {
	ok = r != nil && r.total != nil
	if ok {
		value = *r.total
	}
	return
}

// unmarshal is the method used internally to unmarshal responses to the
// 'list' method.
func (r *CloudProvidersListResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(cloudProvidersListResponseData)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	r.items, err = data.Items.unwrap()
	if err != nil {
		return err
	}
	r.page = data.Page
	r.size = data.Size
	r.total = data.Total
	return err
}

// cloudProvidersListResponseData is the structure used internally to unmarshal
// the response of the 'list' method.
type cloudProvidersListResponseData struct {
	Items cloudProviderListData "json:\"items,omitempty\""
	Page  *int                  "json:\"page,omitempty\""
	Size  *int                  "json:\"size,omitempty\""
	Total *int                  "json:\"total,omitempty\""
}
