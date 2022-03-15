---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "ec_waf_scopes Resource - terraform-provider-edgecast"
subcategory: ""
description: |-
  
---

# ec_waf_scopes (Resource)



## Example Usage

```terraform
resource "ec_waf_scopes" "scopes1" {
  account_number = var.account_number

  scope {
    host {
      is_case_insensitive = false
      type                = "EM"
      values              = ["site1.com/path1","site2.com"]
    }

    limit {
      id                   = "<Rate Rule ID #1>"
      name = "limit action custom"
      duration_sec         = 60
      enf_type             = "CUSTOM_RESPONSE"
      status               = 404
      response_body_base64 = "SGVsbG8sIHdvcmxkIQo="
      response_headers = {
        "header1" = "val1"
        "header2" = "val2"
        "header3" = "val4"
      }
    }

    limit {
      id           = "<Rate Rule ID #2>"
      name          = "limit action drop request"
      duration_sec = 10
      enf_type     = "DROP_REQUEST"
    }

    limit {
      id           = "<Rate Rule ID #3>"
      name = "limit action redirect"
      duration_sec = 10
      enf_type     = "REDIRECT_302"
      url          = "https://mysite.com/redirected2"
    }

    path {
      is_case_insensitive = false
      is_negated          = false
      type                = "GLOB"
      value               = "*"
    }

    acl_audit_action {
      enf_type = "ALERT"
    }

    acl_audit_id = "<Access Rule ID>"

    acl_prod_action {
      name = "acl action"
      enf_type      = "ALERT"
    }

    acl_prod_id = "<Access Rule ID>"

    profile_audit_action {
      enf_type = "ALERT"
    }

    profile_audit_id = "<Managed Rule ID>"

    profile_prod_action {
      name = "managed rule action"
      enf_type             = "CUSTOM_RESPONSE"
      response_body_base64 = "SGVsbG8sIHdvcmxkIQo="
      response_headers = {
        "header1" = "val1"
        "header2" = "val2"
        "header3" = "val3"
      }
      status = 404
    }

    profile_prod_id = "<Managed Rule ID>"

    rules_audit_action {
      enf_type = "ALERT"
    }

    rules_audit_id = "<Custom Rule ID>"

    rules_prod_action {
      name = "custom rule action"
      enf_type             = "BLOCK_REQUEST"
    }

    rules_prod_id = "<Custom Rule ID>"

    bots_prod_id = "<Bots Rule ID>"

    bots_prod_action {
      name = "bots action"
      enf_type = "BROWSER_CHALLENGE"
      valid_for_sec = 60
    }
  }

}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **scope** (Block List, Min: 1) (see [below for nested schema](#nestedblock--scope))

### Optional

- **account_number** (String) Identifies your account by its customer account number.
- **id** (String) The ID of this resource.

<a id="nestedblock--scope"></a>
### Nested Schema for `scope`

Optional:

- **acl_audit_action** (Block Set, Max: 1) Describe the type of action that will take place when the access rule defined within the acl_audit_id property is violated. Refer to the Audit Action section for property details. (see [below for nested schema](#nestedblock--scope--acl_audit_action))
- **acl_audit_id** (String) Indicates the system-defined ID for the access rule that will audit production traffic for this Security Application Manager configuration.
- **acl_prod_action** (Block Set, Max: 1) Describes the type of action that will take place when the access rule defined within the acl_prod_id property is violated. Refer to the Prod Action section for property details. (see [below for nested schema](#nestedblock--scope--acl_prod_action))
- **acl_prod_id** (String) Indicates the system-defined ID for the access rule that will be applied to production traffic for this Security Application Manager configuration.
- **bots_prod_action** (Block Set, Max: 1) Describes the type of action that will take place when the bots rule defined within the bots_prod_id property is violated. Refer to the Prod Action section for property details. (see [below for nested schema](#nestedblock--scope--bots_prod_action))
- **bots_prod_id** (String) Indicates the system-defined ID for the bots rule that will be applied to production traffic for this Security Application Manager configuration.
- **host** (Block Set, Max: 1) Describes a hostname match condition. Refer to the URL and Path section for property details. (see [below for nested schema](#nestedblock--scope--host))
- **limit** (Block List) (see [below for nested schema](#nestedblock--scope--limit))
- **name** (String) Indicates the name assigned to the Security Application Manager configuration. Default Value: 'name'
- **path** (Block Set, Max: 1) Describes a URL match condition.Refer to the URL and Path section for property details. (see [below for nested schema](#nestedblock--scope--path))
- **profile_audit_action** (Block Set, Max: 1) Describes the type of action that will take place when the managed rule defined within the profile_audit_id property is violated. Refer to the Audit Action  section for property details. (see [below for nested schema](#nestedblock--scope--profile_audit_action))
- **profile_audit_id** (String) Indicates the system-defined ID for the managed rule that will audit production traffic for this Security Application Manager configuration.
- **profile_prod_action** (Block Set, Max: 1) Describes the type of action that will take place when the managed rule defined within the profile_prod_id property is violated. Refer to the Prod Action section for property details. (see [below for nested schema](#nestedblock--scope--profile_prod_action))
- **profile_prod_id** (String) Indicates the system-defined ID for the managed rule that will be applied to production traffic for this Security Application Manager configuration.
- **rules_audit_action** (Block Set, Max: 1) Describes the type of action that will take place when the custom rule set defined within the rules_audit_id property is violated. Refer to the Audit Action section for property details. (see [below for nested schema](#nestedblock--scope--rules_audit_action))
- **rules_audit_id** (String) Indicates the system-defined ID for the custom rule set that will audit production traffic for this Security Application Manager configuration.
- **rules_prod_action** (Block Set, Max: 1) Describes the type of action that will take place when the custom rule set defined within the rules_prod_id property is violated. Refer to the Prod Action section for property details. (see [below for nested schema](#nestedblock--scope--rules_prod_action))
- **rules_prod_id** (String) Indicates the system-defined ID for the custom rule set that will be applied to production traffic for this Security Application Manager configuration.

<a id="nestedblock--scope--acl_audit_action"></a>
### Nested Schema for `scope.acl_audit_action`

Required:

- **enf_type** (String)

Optional:

- **name** (String)


<a id="nestedblock--scope--acl_prod_action"></a>
### Nested Schema for `scope.acl_prod_action`

Required:

- **enf_type** (String)

Optional:

- **name** (String)
- **response_body_base64** (String)
- **response_headers** (Map of String)
- **status** (Number)
- **url** (String)
- **valid_for_sec** (Number)


<a id="nestedblock--scope--bots_prod_action"></a>
### Nested Schema for `scope.bots_prod_action`

Required:

- **enf_type** (String)

Optional:

- **name** (String)
- **response_body_base64** (String)
- **response_headers** (Map of String)
- **status** (Number)
- **url** (String)
- **valid_for_sec** (Number)


<a id="nestedblock--scope--host"></a>
### Nested Schema for `scope.host`

Required:

- **type** (String)

Optional:

- **is_case_insensitive** (Boolean)
- **is_negated** (Boolean)
- **value** (String)
- **values** (List of String)


<a id="nestedblock--scope--limit"></a>
### Nested Schema for `scope.limit`

Required:

- **duration_sec** (Number) Indicates the length of time, in seconds, that the action defined within this object will be applied to a client that violates the rate rule identified by the id property.\
\
Valid values are: 10 | 60 | 300
- **enf_type** (String) Indicates the type of action that will be applied to rate limited requests.\
\
Valid values are:ALERT: Alert OnlyREDIRECT_302: Redirect (HTTP 302)CUSTOM_RESPONSE: Custom ResponseDROP_REQUEST: Drop Request (503 Service Unavailable response with a retry-after of 10 seconds)
- **id** (String) Indicates the system-defined ID for the rate limit configuration that will be applied to this Security Application Manager configuration.

Optional:

- **name** (String) Indicates the name assigned to this enforcement action.
- **response_body_base64** (String) Note: Only valid when ENFType is set to CUSTOM_RESPONSE \
\
Indicates the response body that will be sent to rate limited requests. This value is Base64 encoded.
- **response_headers** (Map of String) Note: Only valid when ENFType is set to CUSTOM_RESPONSE\
\
Contains the set of headers that will be included in the response sent to rate limited requests.
- **status** (Number) Note: Only valid when ENFType is set to CUSTOM_RESPONSE\
\
Indicates the HTTP status code (e.g., 404) for the custom response sent to rate limited requests.
- **url** (String) Note: Only valid when ENFType is set to REDIRECT_302\
\
Indicates the URL to which rate limited requests will be redirected.


<a id="nestedblock--scope--path"></a>
### Nested Schema for `scope.path`

Required:

- **type** (String)

Optional:

- **is_case_insensitive** (Boolean)
- **is_negated** (Boolean)
- **value** (String)
- **values** (List of String)


<a id="nestedblock--scope--profile_audit_action"></a>
### Nested Schema for `scope.profile_audit_action`

Required:

- **enf_type** (String)

Optional:

- **name** (String)


<a id="nestedblock--scope--profile_prod_action"></a>
### Nested Schema for `scope.profile_prod_action`

Required:

- **enf_type** (String)

Optional:

- **name** (String)
- **response_body_base64** (String)
- **response_headers** (Map of String)
- **status** (Number)
- **url** (String)
- **valid_for_sec** (Number)


<a id="nestedblock--scope--rules_audit_action"></a>
### Nested Schema for `scope.rules_audit_action`

Required:

- **enf_type** (String)

Optional:

- **name** (String)


<a id="nestedblock--scope--rules_prod_action"></a>
### Nested Schema for `scope.rules_prod_action`

Required:

- **enf_type** (String)

Optional:

- **name** (String)
- **response_body_base64** (String)
- **response_headers** (Map of String)
- **status** (Number)
- **url** (String)
- **valid_for_sec** (Number)

