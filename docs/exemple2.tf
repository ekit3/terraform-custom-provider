terraform {
  required_providers {
    courses = {
      version = "0.0.2"
      source  = "hashicorp.com/ekite/courses"
    }
  }
}

resource "ekite_cour" "mon_cour" {
  provider = courses
  name = "Cour créé depuis terraform"
  time = 60
  summary = "Un simple exemple"
}
