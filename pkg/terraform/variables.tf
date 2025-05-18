variable "project_id" {
  description = "The GCP project where secrets will be created"
  type        = string
}

variable "location" {
  description = "Your GCP region"
  type = string
}

variable "service_name" {
  description = "The name of the Cloud Run service"
  type        = string
  
}

variable "image" {
  description = "The container image to deploy"
  type        = string
  default     = "us-docker.pkg.dev/cloudrun/container/hello"
}

variable "ingress" {
  description = "The ingress settings for the Cloud Run service (INGRESS_TRAFFIC_ALL, INGRESS_TRAFFIC_INTERNAL_ONLY, INGRESS_TRAFFIC_INTERNAL_LOAD_BALANCER)"
  type        = string
  default     = "INGRESS_TRAFFIC_ALL"
}
variable "unauthenticated" {
  description = "Allow unauthenticated invocations"
  type        = bool
  default     = false
  
}