terraform {
  required_providers {
    courses = {
      version = "0.3.18"
      source  = "hashicorp.com/ekite/courses"
    }
  }
}

# data "ekite_courses" "all" {
#   provider = courses
# }

# output "courses" {
#   value = data.ekite_courses.all
# }
# resource "ekite_cour" "example" {
#   provider = courses
#   name = "test from tf update"
#   time = 60
#   summary = "summary from tf"
# }
