package bitbucket

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	bb "github.com/ktrysmt/go-bitbucket"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BITBUCKET_USERNAME", nil),
				Description: "Local user name for the bitbucket API. Can be specified with the `BITBUCKET_USERNAME` " +
					"environment variable.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("BITBUCKET_PASSWORD", nil),
				Description: "Password for the user accessing the API. Can be specified with the `BITBUCKET_PASSWORD` " +
					"environment variable.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"bitbucket_pipeline_variable":  resourcePipelineVariable(),
			"bitbucket_repository":         resourceRepository(),
			"bitbucket_branch_restriction": resourceBranchRestriction(),
			"bitbucket_project":            resourceProject(),
			"bitbucket_default_reviewer":   resourceDefaultReviewer(),
			"bitbucket_group_permission":   resourceGroupPermission(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"bitbucket_repository": dataSourceRepository(),
			"bitbucket_workspace":  dataSourceWorkspace(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (username != "") && (password != "") {
		c := bb.NewBasicAuth(username, password)

		return c, diags
	}

	c := bb.NewBasicAuth("", "")

	return c, diags
}
