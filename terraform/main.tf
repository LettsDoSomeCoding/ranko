provider "google" {
    region = var.region
}

data "google_project" "project" {}

resource "google_storage_bucket" "tfstate_bucket" {
  name          = "tf-state-ranko"
  location      = var.region
  labels        = {
    "app" = "ranko"
  }
  versioning {
      enabled = true
  }
}

resource "google_container_registry" "ranko_registry" {
  location = "EU"
}