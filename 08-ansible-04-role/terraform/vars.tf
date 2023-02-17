variable "yc_token" {
  type      = string
  default   = ""
  sensitive = true
}

variable "yc_cloud_id" {
  type      = string
  default   = ""
  sensitive = true
}

variable "yc_folder_id" {
  type      = string
  default   = ""
  sensitive = true
}

variable "yc_region" {
  type    = string
  default = "ru-central1-a"
}