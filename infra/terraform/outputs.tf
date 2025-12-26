output "instance_name" {
  value = google_compute_instance.app_server.name
}

output "external_ip" {
  value = google_compute_address.static_ip.address
}

output "ssh_command" {
  value = "gcloud compute ssh ${google_compute_instance.app_server.name} --zone ${var.zone}"
}
