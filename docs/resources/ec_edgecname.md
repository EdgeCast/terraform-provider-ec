---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ec_edgecname Resource - terraform-provider-edgecast"
subcategory: ""
description: |-
  
---

# ec_edgecname (Resource)



## Example Usage

```terraform
resource "ec_edgecname" "edgecname_edgecast_origin" {
    account_number = "<account_number>"
    name = "<cname_url>"
    dir_path = "/ec/origin/path"
    enable_custom_reports = 1 # True = 1, False = 0
    media_type_id = 3 # HTTPLarge = 3, HTTP Small = 8, ADN = 14
    origin_id = -1
}

resource "ec_edgecname" "edgecname_customer_origin" {
    account_number = "<account_number>"
    name = "<cname_url>"
    dir_path = "/origin/path/to/content"
    enable_custom_reports = 1 # True = 1, False = 0
    media_type_id = 3 # HTTPLarge = 3, HTTP Small = 8, ADN = 14
    origin_id = 100
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **account_number** (String) Account Number for the customer if not already
					specified in the provider configuration.
- **media_type_id** (Number) Identifies the Delivery Platform on which the
					edge CNAME will be created. 
					3:Http Large, 8:HTTP Small, 14: ADN
- **name** (String) Sets the name that will be assigned to the edge
					CNAME. It should only contain lower-case alphanumeric
					characters, dashes, and periods. The name specified for
					this parameter should also be defined as a CNAME record
					on a DNS server. The CNAME record defined on the DNS server 
					should point to the CDN hostname
					(e.g., wpc.0001.edgecastcdn.net) for the platform
					identified by the "platform" parameter
- **origin_id** (Number) Identifies whether an edge CNAME will be created
					for a CDN origin server or a customer origin server. 
					Valid values: 
					-1: Indicates that you would like to create an
					edge CNAME for our CDN storage service,
					CustomerOriginID: Specifying an ID for an existing
					customer origin configuration indicates that you would
					like to create an edge CNAME for that customer origin

### Optional

- **dir_path** (String) Identifies a location on the origin server. This
					string should specify the relative path from the root
					folder of the origin server to the desired location. Set
					this parameter to blank to point the edge CNAME to the
					root folder of the origin server.
- **enable_custom_reports** (Number) Determines whether hits and data transferred
					statistics will be tracked for this edge CNAME. Logged
					data can be viewed through the Custom Reports module.
					Valid values are:
					0: Disabled (Default Value).
					1: Enabled. CDN activity on this edge CNAME will be logged.
- **id** (String) The ID of this resource.

### Read-Only

- **origin_string** (String) Indicates the origin identifier, the account
					number, and the relative path associated with the edge CNAME.

