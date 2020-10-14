/*
Copyright 2018 The Kubernetes Authors.
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

package priorities

const (
	// EqualPriority defines the name of prioritizer function that gives an equal weight of one to all nodes.
	EqualPriority = "EqualPriority"
	// RequestedToCapacityRatioPriority defines the name of RequestedToCapacityRatioPriority.
	RequestedToCapacityRatioPriority = "RequestedToCapacityRatioPriority"
	// LeastRequestedPriority defines the name of prioritizer function that prioritize nodes by least
	// requested utilization.
	LeastRequestedPriority = "LeastRequestedPriority"
	// BalancedResourceAllocation defines the name of prioritizer function that prioritizes nodes
	// to help achieve balanced resource usage.
	BalancedResourceAllocation = "BalancedResourceAllocation"
)
