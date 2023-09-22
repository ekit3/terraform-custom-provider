terraform {
  required_providers {
    courses = {
      version = "0.0.4"
      source  = "hashicorp.com/ekite/courses"
    }
  }
}

# resource "ekite_cour" "mon_cour" {
#   provider = courses
#   name = "Cour créé depuis terraform"
#   time = 120
#   summary = "Un simple exemple que je viens de mettre à jours !"
# }
