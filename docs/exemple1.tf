terraform {
  required_providers {
    courses = {
      version = "0.1"
      source  = "hashicorp.com/ekite/courses"
    }
  }
}

provider "courses" {}

data "ekite_courses" "all" {
  provider = courses
}
output "all_courses" {
  value = data.ekite_courses.all.courses
}