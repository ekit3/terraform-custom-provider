package courses

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

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
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
	var diags diag.Diagnostics
	client := &http.Client{Timeout: 10 * time.Second}

	url := "http://localhost:4000/cours/" + d.Id()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()
	cour := &Cour{}
	err = json.NewDecoder(r.Body).Decode(&cour)
	if err != nil {
		return diag.FromErr(err)
	}
	d.Set("name", cour.Name)
	d.Set("time", cour.Time)
	d.Set("summary", cour.Summary)

	return diags
}

func resourceCourUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
	name := d.Get("name").(string)
	timeobj := d.Get("time").(int)
	summary := d.Get("summary").(string)

	data := map[string]interface{}{
		"name":    name,
		"time":    timeobj,
		"summary": summary,
	}

	jsonData, err := json.Marshal(data)
	log.Println("ICI!!!!!")
	log.Println(jsonData)

	url := "http://localhost:4000/cours/" + d.Id()

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	log.Println(resp.StatusCode)
	//defer resp.Body.Close()
	d.Set("last_updated", time.Now().Format(time.RFC850))
	return resourceCourRead(ctx, d, m)
}

func resourceCourDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := &http.Client{}
	url := "http://localhost:4000/cours/" + d.Id()
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	d.SetId("")
	return diags
}
