/*
Copyright 2019 The Kubernetes Authors.
Copyright 2020 Authors of Arktos - file modified.

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

package config

const (
	// RPC Server
	RpcPort = "50052"

	// HTTP Server
	HttpAddr   = "0.0.0.0"
	HttpPort   = 8663
	APIVersion = "v1"

	// ClusterController Info
	ClusterControllerIP   = "127.0.0.1"
	ClusterControllerPort = "50053"

	// OpenStack
	OpenStackUsername        = "admin"
	OpenStackPassword        = "secret"
	OpenStackDomainID        = "default"
	OpenStackScopProjectName = "admin"
	OpenStackScopDomainID    = "default"

	// Informer interval period
	FlavorInterval       = 600
	SiteResourceInterval = 10
	VolumePoolInterval   = 60
	VolumeTypeInterval   = 600
	EipPoolInterval      = 60

	// When the maximum number of unreachable requests is reached,
	// send the node unreachable status to the ClusterController
	MaxUnreachableNum = 3
)