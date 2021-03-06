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

// GithubIdentityProviderBuilder contains the data and logic needed to build 'github_identity_provider' objects.
//
// Details for `github` identity providers.
type GithubIdentityProviderBuilder struct {
	ca       *string
	clientID *string
	hostname *string
	teams    []string
}

// NewGithubIdentityProvider creates a new builder of 'github_identity_provider' objects.
func NewGithubIdentityProvider() *GithubIdentityProviderBuilder {
	return new(GithubIdentityProviderBuilder)
}

// CA sets the value of the 'CA' attribute
// to the given value.
//
//
func (b *GithubIdentityProviderBuilder) CA(value string) *GithubIdentityProviderBuilder {
	b.ca = &value
	return b
}

// ClientID sets the value of the 'client_ID' attribute
// to the given value.
//
//
func (b *GithubIdentityProviderBuilder) ClientID(value string) *GithubIdentityProviderBuilder {
	b.clientID = &value
	return b
}

// Hostname sets the value of the 'hostname' attribute
// to the given value.
//
//
func (b *GithubIdentityProviderBuilder) Hostname(value string) *GithubIdentityProviderBuilder {
	b.hostname = &value
	return b
}

// Teams sets the value of the 'teams' attribute
// to the given values.
//
//
func (b *GithubIdentityProviderBuilder) Teams(values ...string) *GithubIdentityProviderBuilder {
	b.teams = make([]string, len(values))
	copy(b.teams, values)
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *GithubIdentityProviderBuilder) Copy(object *GithubIdentityProvider) *GithubIdentityProviderBuilder {
	if object == nil {
		return b
	}
	b.ca = object.ca
	b.clientID = object.clientID
	b.hostname = object.hostname
	if len(object.teams) > 0 {
		b.teams = make([]string, len(object.teams))
		copy(b.teams, object.teams)
	} else {
		b.teams = nil
	}
	return b
}

// Build creates a 'github_identity_provider' object using the configuration stored in the builder.
func (b *GithubIdentityProviderBuilder) Build() (object *GithubIdentityProvider, err error) {
	object = new(GithubIdentityProvider)
	if b.ca != nil {
		object.ca = b.ca
	}
	if b.clientID != nil {
		object.clientID = b.clientID
	}
	if b.hostname != nil {
		object.hostname = b.hostname
	}
	if b.teams != nil {
		object.teams = make([]string, len(b.teams))
		copy(object.teams, b.teams)
	}
	return
}
