package composeapi

import (
	"time"
)

// Deployment structure
type Deployment struct {
	Errors struct {
		Error string `json:"error,omitempty"`
	} `json:"errors,omitempty"`
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

//CreateDeploymentParams Parameters to be completed before creating a deployment
type CreateDeploymentParams struct {
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
