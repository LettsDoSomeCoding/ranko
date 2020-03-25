terraform {
    backend "gcs" {
    bucket  = "tf-state-ranko"
    prefix  = "terraform/state"
    }
}
