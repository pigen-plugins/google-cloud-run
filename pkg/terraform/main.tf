provider "google" {
  project = var.project_id
}

terraform {
  backend "gcs" {}
}

resource "google_cloud_run_v2_service" "pigen_cloudrun" {
  name     = var.service_name
  location = var.location
  ingress = var.ingress

  template {
    containers {
      image = var.image
    }
  }
}

resource "google_cloud_run_v2_service_iam_member" "unauth_invoker" {
  count    = var.unauthenticated ? 1 : 0

  project  = var.project_id
  location = var.location
  name     = google_cloud_run_v2_service.pigen_cloudrun.name

  role     = "roles/run.invoker"
  member   = "allUsers"
}

