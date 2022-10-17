# Copyright 2022 Edgecast Inc., Licensed under the terms of the Apache 2.0 license.
# See LICENSE file in project root for terms.

terraform {
  required_providers {
    edgecast = {
       version = "0.5.9"
      source  = "github.com/terraform-providers/edgecast"
    }
  }
}

##########################################
# Variables
##########################################

variable "credentials" {
  type = object({
    api_token         = string
    ids_client_secret = string
    ids_client_id     = string
    ids_scope         = string
    api_address       = string
    ids_address       = string
  })
}

variable "account_number" {
  type = string
}

variable "MCC_ACCOUNT_EMAIL" {
  type = string
  default = "" # This prevents being prompted during integration test run
}

variable "cert_label_random" {
  type = string
  default = ""
}

##########################################
# Providers
##########################################

provider "edgecast" {
  api_token         = var.credentials.api_token
  ids_client_secret = var.credentials.ids_client_secret
  ids_client_id     = var.credentials.ids_client_id
  ids_scope         = var.credentials.ids_scope
  ids_address       = var.credentials.ids_address
  api_address       = var.credentials.api_address
}


