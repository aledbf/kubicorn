// Copyright © 2017 The Kubicorn Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cluster

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	Cloud_Amazon       = "amazon"
	Cloud_Azure        = "azure"
	Cloud_Google       = "google"
	Cloud_Baremetal    = "baremetal"
	Cloud_DigitalOcean = "digitalocean"
)

type Cluster struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Name              string `json:"name,omitempty"`
	ServerPools       []*ServerPool `json:"serverPools,omitempty"`
	Cloud             string `json:"cloud,omitempty"`
	Location          string `json:"location,omitempty"`
	Ssh               *Ssh `json:"ssh,omitempty"`
	Network           *Network `json:"network,omitempty"`
	Values            *Values `json:"values,omitempty"`
	KubernetesApi     *KubernetesApi `json:"kubernetesApi,omitempty"`
}

func NewCluster(name string) *Cluster {
	return &Cluster{
		Name: name,
	}
}
