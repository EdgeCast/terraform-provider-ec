// Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0
// license. See LICENSE file in project root for terms.

package cps

import (
	"context"
	"log"

	"terraform-provider-edgecast/edgecast/api"

	"github.com/EdgeCast/ec-sdk-go/edgecast/cps/appendix"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kr/pretty"
)

func DataSourceCertReqStatuses() *schema.Resource {
	return &schema.Resource{
		ReadContext: DataSourceCertReqStatusesRead,
		Schema:      namedEntitySchema("Certificate Request Status"),
	}
}

func DataSourceCertReqStatusesRead(
	ctx context.Context,
	d *schema.ResourceData,
	m interface{},
) diag.Diagnostics {
	// Initialize CPS Service
	config := m.(**api.ClientConfig)
	cpsService, err := buildCPSService(**config)
	if err != nil {
		return diag.FromErr(err)
	}

	// Call Get Certificate Request Statuses
	params := appendix.NewAppendixGetCertificateStatusesParams()
	resp, err := cpsService.Appendix.AppendixGetCertificateStatuses(params)
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf(
		"[INFO] Retrieved Certificate Request Statuses: %# v\n",
		pretty.Formatter(resp))

	if resp != nil {
		flattened := FlattenNamedEntities(resp.HyperionCollectionNamedEntity)
		d.Set("items", flattened)
	}

	// always run
	d.SetId(getTimeStamp())

	return diag.Diagnostics{}
}