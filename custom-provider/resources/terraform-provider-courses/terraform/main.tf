terraform {
  required_providers {
    hashicups = {
      version = "0.1"
      source  = "hashicorp.com/ekite/courses"
    }
  }
}

provider "courses" {}

