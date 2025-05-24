output "cloud_run_url" {
  value = google_cloud_run_v2_service.pigen_cloudrun.uri
  description = "The URL of the deployed Cloud Run service"
}

output "service_name" {
  value = google_cloud_run_v2_service.pigen_cloudrun.name
  description = "The name of the deployed Cloud Run service"
  
}