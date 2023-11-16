/*


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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// IBMCloudKMSKeyEntryApplyConfiguration represents an declarative configuration of the IBMCloudKMSKeyEntry type for use
// with apply.
type IBMCloudKMSKeyEntryApplyConfiguration struct {
	CRKID         *string `json:"crkID,omitempty"`
	InstanceID    *string `json:"instanceID,omitempty"`
	CorrelationID *string `json:"correlationID,omitempty"`
	URL           *string `json:"url,omitempty"`
	KeyVersion    *int    `json:"keyVersion,omitempty"`
}

// IBMCloudKMSKeyEntryApplyConfiguration constructs an declarative configuration of the IBMCloudKMSKeyEntry type for use with
// apply.
func IBMCloudKMSKeyEntry() *IBMCloudKMSKeyEntryApplyConfiguration {
	return &IBMCloudKMSKeyEntryApplyConfiguration{}
}

// WithCRKID sets the CRKID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CRKID field is set to the value of the last call.
func (b *IBMCloudKMSKeyEntryApplyConfiguration) WithCRKID(value string) *IBMCloudKMSKeyEntryApplyConfiguration {
	b.CRKID = &value
	return b
}

// WithInstanceID sets the InstanceID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InstanceID field is set to the value of the last call.
func (b *IBMCloudKMSKeyEntryApplyConfiguration) WithInstanceID(value string) *IBMCloudKMSKeyEntryApplyConfiguration {
	b.InstanceID = &value
	return b
}

// WithCorrelationID sets the CorrelationID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CorrelationID field is set to the value of the last call.
func (b *IBMCloudKMSKeyEntryApplyConfiguration) WithCorrelationID(value string) *IBMCloudKMSKeyEntryApplyConfiguration {
	b.CorrelationID = &value
	return b
}

// WithURL sets the URL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the URL field is set to the value of the last call.
func (b *IBMCloudKMSKeyEntryApplyConfiguration) WithURL(value string) *IBMCloudKMSKeyEntryApplyConfiguration {
	b.URL = &value
	return b
}

// WithKeyVersion sets the KeyVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the KeyVersion field is set to the value of the last call.
func (b *IBMCloudKMSKeyEntryApplyConfiguration) WithKeyVersion(value int) *IBMCloudKMSKeyEntryApplyConfiguration {
	b.KeyVersion = &value
	return b
}