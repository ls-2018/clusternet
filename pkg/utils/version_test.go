/*
Copyright 2022 The Clusternet Authors.

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

package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSATokenAutoGenerated(t *testing.T) {
	tests := []struct {
		name              string
		kubeServerVersion string
		want              bool
		err               error
	}{
		// bad versions
		{
			kubeServerVersion: "unknown",
			err:               fmt.Errorf("could not parse \"unknown\" as version"),
		},
		// no auto-generation of secret-based service account token
		{
			kubeServerVersion: KubeV1240Alpha4.String(),
			want:              false,
		},
		{
			kubeServerVersion: "v1.24.0",
			want:              false,
		},
		{
			kubeServerVersion: "v1.24.1",
			want:              false,
		},
		{
			kubeServerVersion: "v1.24.0+k3s1",
			want:              false,
		},
		{
			kubeServerVersion: "v1.24.0-rc2+k3s1",
			want:              false,
		},
		{
			kubeServerVersion: "v1.24.0-alpha.5",
			want:              false,
		},
		{
			kubeServerVersion: "v1.24.0-beta.2",
			want:              false,
		},
		{
			kubeServerVersion: "v1.24.0-rc.0",
			want:              false,
		},
		{
			kubeServerVersion: "v1.25.0",
			want:              false,
		},

		// auto-generation of legacy secret-based service account token
		{
			kubeServerVersion: "v1.24.0-alpha.3",
			want:              true,
		},
		{
			kubeServerVersion: "v1.24.0-alpha.2+k3s1",
			want:              true,
		},
		{
			kubeServerVersion: "v1.23.1",
			want:              true,
		},
		{
			kubeServerVersion: "v1.23.1-alpha.0",
			want:              true,
		},
		{
			kubeServerVersion: "v1.23.5-rc.1",
			want:              true,
		},
		{
			kubeServerVersion: "v1.23.6+k3s1",
			want:              true,
		},
		{
			kubeServerVersion: "v1.22.10",
			want:              true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saTokenAutoGenerated, err := SATokenAutoGenerated(tt.kubeServerVersion)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("kubeServerVersion %s, error expected %v but got %v", tt.kubeServerVersion, tt.err, err)
			}
			if saTokenAutoGenerated != tt.want {
				t.Errorf("kubeServerVersion %s, saTokenAutoGenerated expected %v but got %v", tt.kubeServerVersion, tt.want, saTokenAutoGenerated)
			}
		})
	}
}

func TestEndpointSliceV1beta1Promoted(t *testing.T) {
	tests := []struct {
		name              string
		kubeServerVersion string
		want              bool
		err               error
	}{
		// bad versions
		{
			kubeServerVersion: "unknown",
			err:               fmt.Errorf("could not parse \"unknown\" as version"),
		},
		// EndpointSlice v1beta1 promoted
		{
			kubeServerVersion: KubeV1170Beta2.String(),
			want:              true,
		},
		{
			kubeServerVersion: "v1.17.0",
			want:              true,
		},
		{
			kubeServerVersion: "v1.17.1",
			want:              true,
		},
		{
			kubeServerVersion: "v1.17.0+k3s1",
			want:              true,
		},
		{
			kubeServerVersion: "v1.17.0-rc1+k3s1",
			want:              true,
		},
		{
			kubeServerVersion: "v1.17.0-beta.3",
			want:              true,
		},
		{
			kubeServerVersion: "v1.17.0-rc.0",
			want:              true,
		},
		{
			kubeServerVersion: "v1.18.0",
			want:              true,
		},

		// EndpointSlice v1beta1 not promoted
		{
			kubeServerVersion: "v1.17.0-beta.1",
			want:              false,
		},
		{
			kubeServerVersion: "v1.17.0-alpha.2+k3s1",
			want:              false,
		},
		{
			kubeServerVersion: "v1.16.1",
			want:              false,
		},
		{
			kubeServerVersion: "v1.16.1-alpha.0",
			want:              false,
		},
		{
			kubeServerVersion: "v1.16.5-rc.1",
			want:              false,
		},
		{
			kubeServerVersion: "v1.16.6+k3s1",
			want:              false,
		},
		{
			kubeServerVersion: "v1.16.10",
			want:              false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			endpointSliceV1beta1Promoted, err := EndpointSliceV1beta1Promoted(tt.kubeServerVersion)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("kubeServerVersion %s, error expected %v but got %v", tt.kubeServerVersion, tt.err, err)
			}
			if endpointSliceV1beta1Promoted != tt.want {
				t.Errorf("kubeServerVersion %s, endpointSliceV1beta1Promoted expected %v but got %v", tt.kubeServerVersion, tt.want, endpointSliceV1beta1Promoted)
			}
		})
	}
}

func TestEndpointSliceV1Promoted(t *testing.T) {
	tests := []struct {
		name              string
		kubeServerVersion string
		want              bool
		err               error
	}{
		// bad versions
		{
			kubeServerVersion: "unknown",
			err:               fmt.Errorf("could not parse \"unknown\" as version"),
		},
		// EndpointSlice v1 promoted
		{
			kubeServerVersion: KubeV1210Beta1.String(),
			want:              true,
		},
		{
			kubeServerVersion: "v1.25.0",
			want:              true,
		},
		{
			kubeServerVersion: "v1.22.1",
			want:              true,
		},
		{
			kubeServerVersion: "v1.22.0-alpha.4",
			want:              true,
		},
		{
			kubeServerVersion: "v1.21.1",
			want:              true,
		},
		{
			kubeServerVersion: "v1.21.0",
			want:              true,
		},
		{
			kubeServerVersion: "v1.21.0+k3s1",
			want:              true,
		},
		{
			kubeServerVersion: "v1.21.0-rc1+k3s1",
			want:              true,
		},
		{
			kubeServerVersion: "v1.21.0-rc.0",
			want:              true,
		},

		// EndpointSlice v1 not promoted
		{
			kubeServerVersion: "v1.21.0-beta.0",
			want:              false,
		},
		{
			kubeServerVersion: "v1.17.0-alpha.2+k3s1",
			want:              false,
		},
		{
			kubeServerVersion: "v1.20.10",
			want:              false,
		},
		{
			kubeServerVersion: "v1.21.0-alpha.0",
			want:              false,
		},
		{
			kubeServerVersion: "v1.16.5-rc.1",
			want:              false,
		},
		{
			kubeServerVersion: "v1.16.6+k3s1",
			want:              false,
		},
		{
			kubeServerVersion: "v1.16.10",
			want:              false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			endpointSliceV1Promoted, err := EndpointSliceV1Promoted(tt.kubeServerVersion)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("kubeServerVersion %s, error expected %v but got %v", tt.kubeServerVersion, tt.err, err)
			}
			if endpointSliceV1Promoted != tt.want {
				t.Errorf("kubeServerVersion %s, EndpointSliceV1Promoted expected %v but got %v", tt.kubeServerVersion, tt.want, endpointSliceV1Promoted)
			}
		})
	}
}
