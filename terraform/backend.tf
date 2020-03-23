terraform {
    backend "gcs" {
    bucket  = "opportune-lore-271620"
    prefix  = "terraform/state"
    }
}
