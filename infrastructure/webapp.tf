variable "resource_group_name" {
  type = string
}

variable "web_app_name" {
  type = string
}

variable "web_app_server_name" {
  type = string
}

variable "web_app_server_kind" {
  type = string
}

variable "web_app_server_sku" {
  type = string
}

variable "storage_account_name" {
  type = string
}

variable "storage_account_container_name" {
  type = string
}

variable "location" {
  type = string
}

# Configure the Azure provider
terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.34.0"
    }
  }
  backend "azurerm" {
    key                  = "terraform.tfstate"
  }
}
provider "azurerm" {
  features {}
}

# Get Resource Group data
data "azurerm_resource_group" "rg" {
  name                = var.resource_group_name
}

# Create the Linux App Service Plan
resource "azurerm_service_plan" "appserviceplan" {
  name                = var.web_app_server_name
  location            = var.location
  resource_group_name = var.resource_group_name
  os_type             = var.web_app_server_kind
  sku_name            = var.web_app_server_sku
}

# Create the web app, pass in the App Service Plan ID
resource "azurerm_linux_web_app" "webapp" {
  name                  = var.web_app_name
  location              = var.location
  resource_group_name   = var.resource_group_name
  service_plan_id       = azurerm_service_plan.appserviceplan.id
  https_only            = true
  site_config { 
    minimum_tls_version = "1.2"
    http2_enabled = true
    https_only = true
    always_on = false
    application_stack {
      go_version = 1.18
    }
  }
  app_settings = {
      "SCM_DO_BUILD_DURING_DEPLOYMENT" = "True"
  }
}
