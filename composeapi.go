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

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/parnurzeal/gorequest"
)

// Account structure
type Account struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// AccountResponse holding structure
type AccountResponse struct {
	Embedded struct {
		Accounts []Account `json:"accounts"`
	} `json:"_embedded"`
}

// Link structure for JSON+HAL links
type Link struct {
	HREF      string `json:"href"`
	Templated bool   `json:"templated"`
}

// Deployment structure
type Deployment struct {
	Errors struct {
		Error string `json:"error"`
	} `json:"errors"`
	ID                  string            `json:"id"`
	Name                string            `json:"name"`
	Type                string            `json:"type"`
	CreatedAt           time.Time         `json:"created_at"`
	ProvisionRecipeID   string            `json:"provision_recipe_id"`
	CACertificateBase64 string            `json:"ca_certificate_base64"`
	Connection          ConnectionStrings `json:"connection_strings"`
	Links               struct {
		ComposeWebUILink Link `json:"compose_web_ui"`
	} `json:"_links"`
}

// ConnectionStrings structure
type ConnectionStrings struct {
	Health   string   `json:"health"`
	SSH      string   `json:"ssh"`
	Admin    string   `json:"admin"`
	SSHAdmin string   `json:"ssh_admin"`
	CLI      []string `json:"cli"`
	Direct   []string `json:"direct"`
}

// DeploymentsResponse holding structure
type DeploymentsResponse struct {
	Embedded struct {
		Deployments []Deployment `json:"deployments"`
	} `json:"_embedded"`
}

// Recipe structure
type Recipe struct {
	ID           string `json:"id"`
	Template     string `json:"template"`
	Status       string `json:"status"`
	StatusDetail string `json:"status_detail"`
	AccountID    string `json:"account_id"`
	DeploymentID string `json:"deployment_id"`
	Name         string `json:"name"`

	CreatedAt time.Time `json:"created_at"`
	Embedded  struct {
		Recipes []Recipe `json:"recipes"`
	} `json:"_embedded"`
}

// Recipes structure (an array of Recipe)
type Recipes struct {
	Embedded struct {
		Recipes []Recipe `json:"recipes"`
	} `json:"_embedded"`
}

// Cluster structure
type Cluster struct {
	ID          string    `json:"id"`
	AccountID   string    `json:"account_id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Provider    string    `json:"provider"`
	Region      string    `json:"region"`
	Multitenant bool      `json:"multitenant"`
	AccountSlug string    `json:"account_slug"`
	CreatedAt   time.Time `json:"created_at"`
	Subdomain   string    `json:"subdomain"`
}

// ClustersResponse structure (an array of Cluster)
type ClustersResponse struct {
	Embedded struct {
		Clusters []Cluster `json:"clusters"`
	} `json:"_embedded"`
}

// User structure
type User struct {
	ID string `json:"id"`
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

func getAccount(rawmode bool) (*Account, error) {
	_, body, errs := gorequest.New().Get(apibase+"accounts").
		Set("Authorization", "Bearer "+apitoken).
		Set("Content-type", "json").
		End()

	if errs != nil {
		return nil, errs[0]
	}

	if rawmode {
		printJSON(body)
	}

	accountResponse := AccountResponse{}
	json.Unmarshal([]byte(body), &accountResponse)
	firstAccount := accountResponse.Embedded.Accounts[0]

	return &firstAccount, nil
}

func getDeployments(rawmode bool) (*[]Deployment, error) {
	_, body, errs := gorequest.New().Get(apibase+"deployments").
		Set("Authorization", "Bearer "+apitoken).
		Set("Content-type", "json").
		End()

	if errs != nil {
		return nil, errs[0]
	}

	if rawmode {
		printJSON(body)
	}

	deploymentResponse := DeploymentsResponse{}
	json.Unmarshal([]byte(body), &deploymentResponse)
	deployments := deploymentResponse.Embedded.Deployments

	return &deployments, nil
}

func getRecipe(rawmode bool, recipeid string) (*Recipe, error) {
	_, body, errs := gorequest.New().Get(apibase+"recipes/"+recipeid).
		Set("Authorization", "Bearer "+apitoken).
		Set("Content-type", "json").
		End()

	if errs != nil {
		return nil, errs[0]
	}

	if rawmode {
		printJSON(body)
	}

	recipe := Recipe{}
	json.Unmarshal([]byte(body), &recipe)

	return &recipe, nil
}

func getRecipesForDeployment(rawmode bool, deploymentid string) (*[]Recipe, error) {
	_, body, errs := gorequest.New().Get(apibase+"deployments/"+deploymentid+"/recipes").
		Set("Authorization", "Bearer "+apitoken).
		Set("Content-type", "json").
		End()

	if errs != nil {
		return nil, errs[0]
	}

	if rawmode {
		printJSON(body)
	}

	recipeResponse := Recipe{}
	json.Unmarshal([]byte(body), &recipeResponse)
	recipes := recipeResponse.Embedded.Recipes

	return &recipes, nil
}

func getClusters(rawmode bool) (*[]Cluster, error) {
	_, body, errs := gorequest.New().Get(apibase+"clusters").
		Set("Authorization", "Bearer "+apitoken).
		Set("Content-type", "json").
		End()

	if errs != nil {
		return nil, errs[0]
	}

	if rawmode {
		printJSON(body)
	}

	clustersResponse := ClustersResponse{}
	json.Unmarshal([]byte(body), &clustersResponse)
	clusters := clustersResponse.Embedded.Clusters

	return &clusters, nil
}

func getUser(rawmode bool) (*User, error) {
	_, body, errs := gorequest.New().Get(apibase+"user").
		Set("Authorization", "Bearer "+apitoken).
		Set("Content-type", "json").
		End()

	if errs != nil {
		return nil, errs[0]
	}

	if rawmode {
		printJSON(body)
	}

	user := User{}
	json.Unmarshal([]byte(body), &user)
	return &user, nil

}

type createDeploymentParams struct {
	Name         string `json:"name"`
	AccountID    string `json:"account_id"`
	ClusterID    string `json:"cluster_id,omitempty"`
	Datacenter   string `json:"datacenter,omitempty"`
	DatabaseType string `json:"type"`
	Version      string `json:"version,omitempty"`
	Units        int    `json:"units,omitempty"`
	SSL          bool   `json:"ssl,omitempty"`
	WiredTiger   bool   `json:"wired_tiger,omitempty"`
}

func createDeployment(rawmode bool,
	name string,
	dbtype string,
	accountid string,
	datacenter string,
	clusterid string) (*Deployment, error) {

	newdeployment := createDeploymentParams{Name: name,
		AccountID:    accountid,
		DatabaseType: dbtype,
		Datacenter:   datacenter,
		ClusterID:    clusterid,
	}

	if rawmode {
		tmpjson, err := json.Marshal(newdeployment)
		if err != nil {
			return nil, err
		}
		printJSON(string(tmpjson))
	}

	resp, body, errs := gorequest.New().Post(apibase+"deployments").
		Set("Authorization", "Bearer "+apitoken).
		Set("Content-type", "application/json; charset=utf-8").
		Send(newdeployment).
		End()

	if errs != nil {
		return nil, errs[0]
	}

	if rawmode {
		fmt.Println(resp)
		printJSON(body)
	}

	deployed := Deployment{}
	json.Unmarshal([]byte(body), &deployed)

	return &deployed, nil
}
