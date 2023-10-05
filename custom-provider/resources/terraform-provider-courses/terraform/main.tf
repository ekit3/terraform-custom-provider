terraform {
  required_providers {
    courses = {
      version = "0.1"
      source  = "hashicorp.com/ekite/courses"
    }
  }
}

provider "courses" {}

