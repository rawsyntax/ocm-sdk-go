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
	time "time"
)

// SampleBuilder contains the data and logic needed to build 'sample' objects.
//
// Sample of a metric.
type SampleBuilder struct {
	time  *time.Time
	value *float64
}

// NewSample creates a new builder of 'sample' objects.
func NewSample() *SampleBuilder {
	return new(SampleBuilder)
}

// Time sets the value of the 'time' attribute
// to the given value.
//
//
func (b *SampleBuilder) Time(value time.Time) *SampleBuilder {
	b.time = &value
	return b
}

// Value sets the value of the 'value' attribute
// to the given value.
//
//
func (b *SampleBuilder) Value(value float64) *SampleBuilder {
	b.value = &value
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *SampleBuilder) Copy(object *Sample) *SampleBuilder {
	if object == nil {
		return b
	}
	b.time = object.time
	b.value = object.value
	return b
}

// Build creates a 'sample' object using the configuration stored in the builder.
func (b *SampleBuilder) Build() (object *Sample, err error) {
	object = new(Sample)
	if b.time != nil {
		object.time = b.time
	}
	if b.value != nil {
		object.value = b.value
	}
	return
}
