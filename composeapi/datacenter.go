// Copyright 2016 Compose, an IBM Company
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package composeapi

//Datacenter structure
type Datacenter struct {
	Region   string `json:"region"`
	Provider string `json:"provider"`
	Slug     string `json:"slug"`
}

//DatacentersResponse structure (an array of Datacenter)
type DatacentersResponse struct {
	Embedded struct {
		Datacenters []Datacenter `json:"datacenters"`
	} `json:"_embedded"`
}
