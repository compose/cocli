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

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/parnurzeal/gorequest"
)

var (
	apitoken = os.Getenv("COMPOSEAPITOKEN")
)

const (
	apibase = "https://api.compose.io/2016-07/"
)

// Link structure for JSON+HAL links
type Link struct {
	HREF      string `json:"href"`
	Templated bool   `json:"templated"`
}

func printJSON(jsontext string) {
	var tempholder map[string]interface{}

	if err := json.Unmarshal([]byte(jsontext), &tempholder); err != nil {
		log.Fatal(err)
	}
	indentedjson, err := json.MarshalIndent(tempholder, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(indentedjson))
}

//GetJSON Gets JSON string of content at an endpoint
func getJSON(endpoint string) (string, []error) {
	_, body, errs := gorequest.New().Get(apibase+endpoint).
		Set("Authorization", "Bearer "+apitoken).
		Set("Content-type", "json").
		End()

	return body, errs
}

//GetAccountJSON gets JSON string from endpoint
func GetAccountJSON() (string, []error) { return getJSON("accounts") }

//GetAccount Gets first Account struct from account endpoint
func GetAccount() (*Account, []error) {
	body, errs := GetAccountJSON()

	if errs != nil {
		return nil, errs
	}

	accountResponse := AccountResponse{}
	json.Unmarshal([]byte(body), &accountResponse)
	firstAccount := accountResponse.Embedded.Accounts[0]

	return &firstAccount, nil
}

//GetDeploymentsJSON returns raw deployment
func GetDeploymentsJSON() (string, []error) { return getJSON("deployments") }

//GetDeployments returns deployment structure
func GetDeployments() (*[]Deployment, []error) {
	body, errs := GetDeploymentsJSON()

	if errs != nil {
		return nil, errs
	}

	deploymentResponse := DeploymentsResponse{}
	json.Unmarshal([]byte(body), &deploymentResponse)
	deployments := deploymentResponse.Embedded.Deployments

	return &deployments, nil
}

//GetRecipeJSON Gets raw JSON for recipeid
func GetRecipeJSON(recipeid string) (string, []error) { return getJSON("recipes/" + recipeid) }

//GetRecipe gets status of Recipe
func GetRecipe(rawmode bool, recipeid string) (*Recipe, []error) {
	body, errs := GetRecipeJSON(recipeid)

	if errs != nil {
		return nil, errs
	}

	recipe := Recipe{}
	json.Unmarshal([]byte(body), &recipe)

	return &recipe, nil
}

//GetRecipesForDeploymentJSON returns raw JSON for getRecipesforDeployment
func GetRecipesForDeploymentJSON(deploymentid string) (string, []error) {
	return getJSON("deployments/" + deploymentid + "/recipes")
}

//GetRecipesForDeployment gets deployment recipe life
func GetRecipesForDeployment(deploymentid string) (*[]Recipe, []error) {
	body, errs := GetRecipesForDeploymentJSON(deploymentid)

	if errs != nil {
		return nil, errs
	}

	recipeResponse := Recipe{}
	json.Unmarshal([]byte(body), &recipeResponse)
	recipes := recipeResponse.Embedded.Recipes

	return &recipes, nil
}

//GetVersionsForDeploymentJSON returns raw JSON for getVersionsforDeployment
func GetVersionsForDeploymentJSON(deploymentid string) (string, []error) {
	return getJSON("deployments/" + deploymentid + "/versions")
}

//GetVersionsForDeployment gets deployment recipe life
func GetVersionsForDeployment(deploymentid string) (*[]VersionTransition, []error) {
	body, errs := GetVersionsForDeploymentJSON(deploymentid)

	if errs != nil {
		return nil, errs
	}

	versionsResponse := VersionsResponse{}
	json.Unmarshal([]byte(body), &versionsResponse)
	versionTransitions := versionsResponse.Embedded.VersionTransitions

	return &versionTransitions, nil
}

//GetClustersJSON gets clusters available
func GetClustersJSON() (string, []error) {
	return getJSON("clusters")
}

//GetClusters gets clusters available
func GetClusters() (*[]Cluster, []error) {
	body, errs := GetClustersJSON()

	if errs != nil {
		return nil, errs
	}

	clustersResponse := ClustersResponse{}
	json.Unmarshal([]byte(body), &clustersResponse)
	clusters := clustersResponse.Embedded.Clusters

	return &clusters, nil
}

//GetDatacentersJSON gets datacenters available as a string
func GetDatacentersJSON() (string, []error) {
	return getJSON("datacenters")
}

//GetDatacenters gets datacenters available as a Go struct
func GetDatacenters() (*[]Datacenter, []error) {
	body, errs := GetDatacentersJSON()

	if errs != nil {
		return nil, errs
	}

	datacenterResponse := DatacentersResponse{}
	json.Unmarshal([]byte(body), &datacenterResponse)
	datacenters := datacenterResponse.Embedded.Datacenters

	return &datacenters, nil
}

//GetDatabasesJSON gets databases available as a string
func GetDatabasesJSON() (string, []error) {
	return getJSON("databases")
}

//GetDatabases gets databases available as a Go struct
func GetDatabases() (*[]Database, []error) {
	body, errs := GetDatabasesJSON()

	if errs != nil {
		return nil, errs
	}

	datacenterResponse := DatabasesResponse{}
	json.Unmarshal([]byte(body), &datacenterResponse)
	databases := datacenterResponse.Embedded.Databases

	return &databases, nil
}

//GetUserJSON returns user JSON string
func GetUserJSON() (string, []error) {
	return getJSON("user")
}

//GetUser Gets information about user
func GetUser() (*User, []error) {
	body, errs := GetUserJSON()

	if errs != nil {
		return nil, errs
	}

	user := User{}
	json.Unmarshal([]byte(body), &user)
	return &user, nil
}

//CreateDeploymentJSON performs the call
func CreateDeploymentJSON(params CreateDeploymentParams) (string, []error) {
	_, body, errs := gorequest.New().Post(apibase+"deployments").
		Set("Authorization", "Bearer "+apitoken).
		Set("Content-type", "application/json; charset=utf-8").
		Send(params).
		End()
	return body, errs
}

//CreateDeployment creates a deployment
func CreateDeployment(params CreateDeploymentParams) (*Deployment, []error) {

	// This is a POST not a GET, so it builds its own request

	body, errs := CreateDeploymentJSON(params)

	if errs != nil {
		return nil, errs
	}

	deployed := Deployment{}
	json.Unmarshal([]byte(body), &deployed)

	return &deployed, nil
}
