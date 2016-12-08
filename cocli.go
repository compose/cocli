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
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("cocli", "A Compose CLI application")

	rawmodeflag             = app.Flag("raw", "Output raw JSON responses").Default("false").Bool()
	fullcaflag              = app.Flag("fullca", "Show all of CA Certificates").Default("false").Bool()
	showcmd                 = app.Command("show", "Show attribute")
	showaccountcmd          = showcmd.Command("account", "Show account details")
	showdeploymentscmd      = showcmd.Command("deployments", "Show deployments")
	showrecipecmd           = showcmd.Command("recipe", "Show recipe")
	showrecipeid            = showrecipecmd.Arg("recid", "Recipe ID").String()
	showrecipescmd          = showcmd.Command("recipes", "Show recipes for a deployment")
	showrecipesdeploymentid = showrecipescmd.Arg("depid", "Deployment ID").String()
	showclusterscmd         = showcmd.Command("clusters", "Show available clusters")
	showuser                = showcmd.Command("user", "Show current associated user")

	createcmd                  = app.Command("create", "Create...")
	createdeploymentcmd        = createcmd.Command("deployment", "Create deployment")
	createdeploymentname       = createdeploymentcmd.Arg("name", "New Deployment Name").String()
	createdeploymenttype       = createdeploymentcmd.Arg("type", "New Deployment Type").String()
	createdeploymentcluster    = createdeploymentcmd.Flag("cluster", "Cluster ID").String()
	createdeploymentdatacenter = createdeploymentcmd.Flag("datacenter", "Datacenter location").String()

	apitoken = os.Getenv("COMPOSEAPITOKEN")
)

const (
	apibase = "https://api.compose.io/2016-07/"
)

func main() {
	if apitoken == "" {
		log.Fatal("COMPOSEAPITOKEN environment variable not set")
	}

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case "show account":
		account, err := getAccount(*rawmodeflag)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%15s: %s\n", "ID", account.ID)
		fmt.Printf("%15s: %s\n", "Name", account.Name)
		fmt.Printf("%15s: %s\n", "Slug", account.Slug)
		fmt.Println()

	case "show deployments":
		deployments, err := getDeployments(*rawmodeflag)

		if err != nil {
			log.Fatal(err)
		}

		for _, v := range *deployments {
			fmt.Printf("%15s: %s\n", "ID", v.ID)
			fmt.Printf("%15s: %s\n", "Name", v.Name)
			fmt.Printf("%15s: %s\n", "Type", v.Type)
			fmt.Printf("%15s: %s\n", "Created At", v.CreatedAt)
			fmt.Printf("%15s: %s\n", "Web UI Link", getLink(v.Links.ComposeWebUILink))
			fmt.Println()
		}
	case "show recipe":
		recipe, err := getRecipe(*rawmodeflag, *showrecipeid)

		if err != nil {
			log.Fatal(err)
		}

		printRecipe(*recipe)
	case "show recipes":
		recipes, err := getRecipesForDeployment(*rawmodeflag, *showrecipesdeploymentid)

		if err != nil {
			log.Fatal(err)
		}

		for _, v := range *recipes {
			printRecipe(v)
			fmt.Println()
		}

	case "show clusters":
		clusters, err := getClusters(*rawmodeflag)
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range *clusters {
			printCluster(v)
			fmt.Println()
		}
	case "show user":
		user, err := getUser(*rawmodeflag)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%15s: %s\n", "ID", user.ID)
		fmt.Println()
	case "create deployment":
		account, err := getAccount(false)
		if err != nil {
			log.Fatal(err)
		}

		if *createdeploymentdatacenter == "" && *createdeploymentcluster == "" {
			log.Fatal("Must supply either a --cluster id or --datacenter region")
		}

		deployment, err := createDeployment(*rawmodeflag,
			*createdeploymentname,
			*createdeploymenttype,
			account.ID,
			*createdeploymentdatacenter,
			*createdeploymentcluster)
		if err != nil {
			log.Fatal(err)
		}

		if deployment.Errors.Error != "" {
			fmt.Printf("Error: %s\n", deployment.Errors.Error)
		} else {
			printDeployment(*deployment)
		}
	}
}

func getLink(link Link) string {
	return strings.Replace(link.HREF, "{?embed}", "", -1) // TODO: This should mangle the HREF properly
}

func printRecipe(recipe Recipe) {
	fmt.Printf("%15s: %s\n", "ID", recipe.ID)
	fmt.Printf("%15s: %s\n", "Template", recipe.Template)
	fmt.Printf("%15s: %s\n", "Status", recipe.Status)
	fmt.Printf("%15s: %s\n", "Status Detail", recipe.StatusDetail)
	fmt.Printf("%15s: %s\n", "Account ID", recipe.AccountID)
	fmt.Printf("%15s: %s\n", "Deployment ID", recipe.DeploymentID)
	fmt.Printf("%15s: %s\n", "Name", recipe.Name)
	fmt.Printf("%15s: %d\n", "Child Recipes", len(recipe.Embedded.Recipes))

}

func printCluster(cluster Cluster) {
	fmt.Printf("%15s: %s\n", "ID", cluster.ID)
	fmt.Printf("%15s: %s\n", "Account ID", cluster.AccountID)
	fmt.Printf("%15s: %s\n", "Account Slug", cluster.AccountSlug)
	fmt.Printf("%15s: %s\n", "Name", cluster.Name)
	fmt.Printf("%15s: %s\n", "Type", cluster.Type)
	fmt.Printf("%15s: %t\n", "Multitenant", cluster.Multitenant)
	fmt.Printf("%15s: %s\n", "Provider", cluster.Provider)
	fmt.Printf("%15s: %s\n", "Region", cluster.Region)
	fmt.Printf("%15s: %s\n", "Created Ad", cluster.CreatedAt)
	fmt.Printf("%15s: %s\n", "Subdomain", cluster.Subdomain)
}

func printDeployment(deployment Deployment) {
	fmt.Printf("%15s: %s\n", "ID", deployment.ID)
	fmt.Printf("%15s: %s\n", "Name", deployment.Name)
	fmt.Printf("%15s: %s\n", "Type", deployment.Type)
	fmt.Printf("%15s: %s\n", "Created At", deployment.CreatedAt)
	if deployment.ProvisionRecipeID != "" {
		fmt.Printf("%15s: %s\n", "Prov Recipe ID", deployment.ProvisionRecipeID)
	}
	if deployment.CACertificateBase64 != "" {
		if *fullcaflag {
			fmt.Printf("%15s: %s\n", "CA Certificate", deployment.CACertificateBase64)
		} else {
			fmt.Printf("%15s: %s...\n", "CA Certificate", deployment.CACertificateBase64[0:32])
		}
	}
	fmt.Printf("%15s: %s\n", "Web UI Link", getLink(deployment.Links.ComposeWebUILink))
	fmt.Printf("%15s: %s\n", "Health", deployment.Connection.Health)
	fmt.Printf("%15s: %s\n", "SSH", deployment.Connection.SSH)
	fmt.Printf("%15s: %s\n", "Admin", deployment.Connection.Admin)
	fmt.Printf("%15s: %s\n", "SSHAdmin", deployment.Connection.SSHAdmin)
	fmt.Printf("%15s: %s\n", "CLI Connect", deployment.Connection.CLI)
	fmt.Printf("%15s: %s\n", "Direct Connect", deployment.Connection.Direct)

}
