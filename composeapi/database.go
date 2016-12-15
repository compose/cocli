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

//Version structure
type Version struct {
	Application string `json:"application"`
	Status      string `json:"status"`
	Preferred   bool   `json:"preferred"`
	Version     string `json:"version"`
}

//Database structure
type Database struct {
	DatabaseType string `json:"type"`
	Status       string `json:"status"`
	Embedded     struct {
		Versions []Version `json:"versions"`
	} `json:"_embedded"`
}

//DatabasesResponse structure (an array of Datacenter)
type DatabasesResponse struct {
	Embedded struct {
		Databases []Database `json:"applications"`
	} `json:"_embedded"`
}
