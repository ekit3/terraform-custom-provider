package courses

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"ekite_cour": dataSourceCour(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"ekite_courses": dataSourceCourses(),
		},
	}
}
