# OS Login 用 IAM 付与 (`gcloud compute ssh` コマンドでの接続に必要)
resource "google_project_iam_member" "os_login" {
  for_each = toset(var.oslogin_members)

  project = var.project_id
  role    = "roles/compute.osLogin"
  member  = each.value
}

# OS Admin Login 用 IAM 付与 (sudo 相当、`gcloud compute ssh` コマンドでの接続に必要)
resource "google_project_iam_member" "os_admin_login" {
  for_each = toset(var.osadminlogin_members)

  project = var.project_id
  role    = "roles/compute.osAdminLogin"
  member  = each.value
}
