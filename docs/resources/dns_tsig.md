---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ec_dns_tsig Resource - terraform-provider-edgecast"
subcategory: ""
description: |-
  
---

# ec_dns_tsig (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **account_number** (String) Account Number for the customer if not already specified in the provider configuration.
- **algorithm_name** (String) tsig encryption type:[HMAC-MD5,HMAC-SHA1,HMAC-SHA256,HMAC-SHA384,HMAC-SHA224,HMAC-SHA512]
- **alias** (String) Alias.
- **key_name** (String) tsig key name
- **key_value** (String) tsig value

### Optional

- **id** (String) The ID of this resource.

