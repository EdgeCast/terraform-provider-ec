// Copyright Edgecast, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.
package waf

import (
	"context"
	"errors"
	"fmt"
	"log"
	"terraform-provider-ec/ec/api"
	"terraform-provider-ec/ec/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	sdkwaf "github.com/edgecast/ec-sdk-go/edgecast/waf"
)

func ResourceRateRule() *schema.Resource {

	return &schema.Resource{
		CreateContext: ResourceRateRuleCreate,
		ReadContext:   ResourceRateRuleRead,
		UpdateContext: ResourceRateRuleUpdate,
		DeleteContext: ResourceRateRuleDelete,

		Schema: map[string]*schema.Schema{
			"account_number": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Identifies your account by its customer account number.",
			},
			"duration_sec": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntAtLeast(1),
				Description: "Indicates the length, in seconds, of the rolling window that tracks the number of requests eligible for rate limiting. \\\n" +
					"The rate limit formula is calculated through the num and duration_sec properties as indicated below. \\\n" +
					"    `num` requests per `duration_sec` \\\n    Valid values are: \\\n    `1 | 5 | 10 | 30 | 60 | 120 | 300`",
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: "Indicates whether this rate rule will be enforced. \\\n" +
					"Valid values are: \n" +
					"    * true: Disabled. This rate limit will not be applied to traffic.\n" +
					"    * false: Enabled. Traffic is restricted to this rate limit.",
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
				Description:  "Assigns a name to this access rule.",
			},
			"num": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntAtLeast(1),
				Description: "Indicates the rate limit value. This value identifies the number of requests that will trigger rate limiting. \\\n" +
					"The rate limit formula is calculated through the num and duration_sec properties as indicated below. \\\n" +
					"`num` requests per `duration_sec`",
			},
			"keys": {
				Type:     schema.TypeSet,
				Optional: true,
				Description: "Indicates the method by requests will be grouped for the purposes of this rate rule. \\\n" +
					"Valid values are: \n" +
					"    * Missing / Empty Array: If the `keys` property is not defined or set to an empty array, " +
					"all requests will be treated as a single group for the purpose of rate limiting. \n" +
					"    * IP: Indicates that requests will be grouped by IP address. Each unique IP address is considered a separate group. \n" +
					"    * USER_AGENT: Indicates that requests will be grouped by a client's user agent. " +
					"Each unique combination of IP address and user agent is considered a separate group. \n",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"condition_group": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "Indicates the system-defined alphanumeric ID of a condition group. Example: `12345678-90ab-cdef-ghij-klmnopqrstuvwxyz1`",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Indicates the name of a condition group.",
							Optional:    true,
						},
						"condition": {
							Type:        schema.TypeSet,
							Description: "Contains a list of match conditions.",
							Optional:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target": {
										Type:        schema.TypeSet,
										Description: "Describes the type of match condition.",
										Optional:    true,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Required: true,
													Type:     schema.TypeString,
													Description: "Determines how requests will be identified. \\\n" +
														"    Valid values are: `FILE_EXT | REMOTE_ADDR | REQUEST_HEADERS | REQUEST_METHOD | REQUEST_URI`",
												},
												"value": {
													Optional: true,
													Type:     schema.TypeString,
													Description: "type: REQUEST_HEADERS Only \\\n" +
														"Indicates the name of the request header through which requests will be identified. \\\n" +
														"    Valid values are: `Host | Referer | User-Agent`",
												},
											},
										},
									},
									"op": {
										Type:        schema.TypeSet,
										Description: "Contains the match condition's properties",
										Optional:    true,
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"is_case_insensitive": {
													Optional:    true,
													Type:        schema.TypeBool,
													Description: "Indicates whether the comparison between the request and the values property is case-sensitive.",
												},
												"is_negated": {
													Optional:    true,
													Type:        schema.TypeBool,
													Description: "Indicates whether this match condition will be satisfied when the request matches or does not match the value defined by the values property.",
												},
												"type": {
													Required: true,
													Type:     schema.TypeString,
													Description: "Indicates how the system will interpret the comparison between the request and the `values` property. Valid values are: \\\n" +
														"    * EM: Requires that the request's attribute be set to one of the value(s) defined in the `values` property. \n" +
														"    * IPMATCH: Requires that the request's IP address either be contained by an IP block or be an exact match to an IP address defined in the `values` property. \\\n" +
														"    *Note: Only use IPMATCH with the REMOTE_ADDR match condition.* \n" +
														"    * RX: Requires that the request's attribute be an exact match to the regular expression defined in the `value` property. \n",
												},
												"value": {
													Optional: true,
													Type:     schema.TypeString,
													Description: "type: REQUEST_HEADERS Only \\\n" +
														"Indicates the name of the request header through which requests will be identified. \\\n" +
														"    Valid values are: `Host | Referer | User-Agent`",
												},
												"values": {
													Type:     schema.TypeSet,
													Optional: true,
													Description: "type: EM and IPMATCH Only \\\n" +
														"Identifies one or more values used to identify requests that are eligible for rate limiting. \\\n" +
														"If you are identifying traffic via a URL path (REQUEST_URI), then you should specify a URL path " +
														"pattern that starts directly after the hostname. Exclude a protocol or a hostname when defining this property. \\\n" +
														"Sample values: \\\n    /marketing \\\n    /800001/mycustomerorigin \\\n" +
														"*Note:If you are matching requests by IP address, make sure to use standard IPv4 and CIDR notation.*",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				Description: "Contains the set of condition groups associated with a rule.",
			},
		},
	}
}

func ResourceRateRuleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	config := m.(**api.ClientConfig)

	wafService, err := buildWAFService(**config)

	if err != nil {
		return diag.FromErr(err)
	}

	accountNumber := d.Get("account_number").(string)

	log.Printf("[INFO] Creating WAF Rate Rule for Account >> %s", accountNumber)

	rateRule := sdkwaf.RateRule{
		Name:        d.Get("name").(string),
		CustomerID:  accountNumber,
		Disabled:    d.Get("disabled").(bool),
		Num:         d.Get("num").(int),
		DurationSec: d.Get("duration_sec").(int),
	}

	if v, ok := d.GetOk("keys"); ok {
		if keys, ok := helper.ConvertInterfaceToStringArray(v); ok {
			rateRule.Keys = *keys
		} else {
			return diag.Errorf("Error reading 'keys''")
		}
	}

	conditionGroups, err := ExpandConditionGroups(d.Get("condition_group"))

	if err != nil {
		return diag.Errorf("error parsing condition_groups: %+v", err)
	}

	rateRule.ConditionGroups = *conditionGroups

	log.Printf("[DEBUG] Customer ID: %+v \n", rateRule.CustomerID)
	log.Printf("[DEBUG] Disabled: %+v\n", rateRule.Disabled)
	log.Printf("[DEBUG] DurationSec: %+v\n", rateRule.DurationSec)
	log.Printf("[DEBUG] Name: %+v\n", rateRule.Name)
	log.Printf("[DEBUG] Num: %+v\n", rateRule.Num)
	log.Printf("[DEBUG] Keys: %+v\n", rateRule.Keys)
	log.Printf("[DEBUG] ConditionGroups: %+v\n", rateRule.ConditionGroups)

	resp, err := wafService.AddRateRule(rateRule)

	if err != nil {
		d.SetId("")
		return diag.FromErr(err)
	}

	log.Printf("[INFO] %+v", resp)

	d.SetId(resp.ID)

	return ResourceRateRuleRead(ctx, d, m)
}

func ResourceRateRuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	config := m.(**api.ClientConfig)
	accountNumber := d.Get("account_number").(string)
	ruleID := d.Id()

	log.Printf("[INFO] Retrieving rate rule %s for account number %s", ruleID, accountNumber)

	wafService, err := buildWAFService(**config)

	if err != nil {
		return diag.FromErr(err)
	}

	resp, err := wafService.GetRateRule(accountNumber, ruleID)

	if err != nil {
		d.SetId("")
		return diag.FromErr(err)
	}

	log.Printf("[INFO] Successfully retrieved rate rule %s: %+v", ruleID, resp)

	d.SetId(resp.ID)
	d.Set("account_number", accountNumber)
	d.Set("duration_sec", resp.DurationSec)
	d.Set("disabled", resp.Disabled)
	d.Set("name", resp.Name)
	d.Set("num", resp.Num)
	d.Set("keys", resp.Keys)

	flattenedConditionGroups := FlattenConditionGroups(resp.ConditionGroups)

	d.Set("condition_group", flattenedConditionGroups)

	return diags
}

func ResourceRateRuleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	return diags
}

func ResourceRateRuleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics
	return diags
}

// ExpandConditionGroups converts user-provided Terraform configuration data into the Condition Group API Model
func ExpandConditionGroups(attr interface{}) (*[]sdkwaf.ConditionGroup, error) {

	if set, ok := attr.(*schema.Set); ok {

		items := set.List()
		groups := make([]sdkwaf.ConditionGroup, 0)

		for _, item := range items {
			curr := item.(map[string]interface{})

			group := sdkwaf.ConditionGroup{
				ID:   curr["id"].(string),
				Name: curr["name"].(string),
			}

			conditions, err := ExpandConditions(curr["condition"])

			if err != nil {
				return nil, fmt.Errorf("error parsing conditions: %v", err)
			}

			group.Conditions = *conditions

			groups = append(groups, group)
		}

		return &groups, nil

	} else {
		return nil, errors.New("attr input was not a *schema.Set")
	}
}

// ExpandConditions converts user-provided Terraform configuration data into the Condition API Model
func ExpandConditions(attr interface{}) (*[]sdkwaf.Condition, error) {

	if set, ok := attr.(*schema.Set); ok {

		items := set.List()
		conditions := make([]sdkwaf.Condition, 0)

		for _, item := range items {
			curr := item.(map[string]interface{})

			// The properties for target and op are stored as a map in a 1-item set
			targetMap, err := helper.GetMapFromSet(curr["target"])

			if err != nil {
				return nil, err
			}

			opMap, err := helper.GetMapFromSet(curr["op"])

			if err != nil {
				return nil, err
			}

			condition := sdkwaf.Condition{}

			if targetType, ok := targetMap["type"]; ok {
				condition.Target.Type = targetType.(string)
			}

			if targetValue, ok := targetMap["value"]; ok {
				condition.Target.Value = targetValue.(string)
			}

			if opType, ok := opMap["type"]; ok {
				condition.OP.Type = opType.(string)
			}

			if opValue, ok := opMap["value"]; ok {
				condition.OP.Value = opValue.(string)
			}

			if opValues, ok := opMap["values"]; ok {
				if arr, ok := helper.ConvertInterfaceToStringArray(opValues); ok {
					condition.OP.Values = *arr
				}
			}

			if v, ok := opMap["is_case_insensitive"]; ok {
				boolValue := v.(bool)
				condition.OP.IsCaseInsensitive = &boolValue
			}

			if v, ok := opMap["is_negated"]; ok {
				boolValue := v.(bool)
				condition.OP.IsNegated = &boolValue
			}

			conditions = append(conditions, condition)
		}

		return &conditions, nil

	} else {
		return nil, errors.New("attr input was not a *schema.Set")
	}
}

// FlattenConditionGroups converts the ConditionGroup API Model
// into a format that Terraform can work with
func FlattenConditionGroups(conditionGroups []sdkwaf.ConditionGroup) []map[string]interface{} {

	flattened := make([]map[string]interface{}, 0)

	for _, cg := range conditionGroups {
		m := make(map[string]interface{})

		m["id"] = cg.ID
		m["name"] = cg.Name
		m["condition"] = FlattenConditions(cg.Conditions)

		flattened = append(flattened, m)
	}

	return flattened
}

// FlattenConditions converts the Condition API Model
// into a format that Terraform can work with
func FlattenConditions(conditions []sdkwaf.Condition) []map[string]interface{} {

	flattened := make([]map[string]interface{}, 0)

	for _, c := range conditions {
		m := make(map[string]interface{})

		m["op"] = FlattenOP(c.OP)
		m["target"] = FlattenTarget(c.Target)

		flattened = append(flattened, m)
	}

	return flattened
}

// FlattenOP converts the OP API Model
// into a format that Terraform can work with
func FlattenOP(op sdkwaf.OP) []map[string]interface{} {

	m := make(map[string]interface{})

	if op.IsNegated != nil {
		m["is_negated"] = *(op.IsNegated)
	}

	if op.IsCaseInsensitive != nil {
		m["is_case_insensitive"] = *(op.IsNegated)
	}

	m["type"] = op.Type
	m["value"] = op.Value
	m["values"] = op.Values

	// We return a collection of just 1 item
	// Since we defined OP as a 1-item set in the schema
	return []map[string]interface{}{m}
}

// FlattenTarget converts the Target API Model
// into a format that Terraform can work with
func FlattenTarget(target sdkwaf.Target) []map[string]interface{} {

	m := make(map[string]interface{})

	m["type"] = target.Type
	m["value"] = target.Value

	// We return a collection of just 1 item
	// Since we defined Target as a 1-item set in the schema
	return []map[string]interface{}{m}
}
