package courses

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Cour struct {
	Id      string `json:"_id"`
	Name    string `json:"name"`
	Time    int    `json:"time"`
	Summary string `json:"summary"`
}

func dataSourceCour() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCourCreate,
		ReadContext:   resourceCourRead,
		UpdateContext: resourceCourUpdate,
		DeleteContext: resourceCourDelete,

		Schema: map[string]*schema.Schema{
			"_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"time": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"summary": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCourCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	time := d.Get("time").(int)
	summary := d.Get("summary").(string)

	data := map[string]interface{}{
		"name":    name,
		"time":    time,
		"summary": summary,
	}

	jsonData, err := json.Marshal(data)

	// https://www.uuidtools.com/api/generate/v1/count/uuid_count
	resp, err := http.Post("http://localhost:4000/cours", "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	cour := &Cour{}
	err = json.NewDecoder(resp.Body).Decode(&cour)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(cour.Id)
	return diags
}

func resourceCourRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	return nil
}

func resourceCourUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	return nil
}

func resourceCourDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	return nil
}
