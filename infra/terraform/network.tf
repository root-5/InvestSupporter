# VPC ネットワークの作成
resource "google_compute_network" "vpc_network" {
  name                    = "invest-supporter-vpc"
  auto_create_subnetworks = false
}

# サブネットを作成し、VPC ネットワークに関連付け
resource "google_compute_subnetwork" "subnet" {
  name          = "invest-supporter-subnet"
  ip_cidr_range = "10.0.1.0/24"
  region        = var.region
  network       = google_compute_network.vpc_network.id
}

# ファイアウォールルールの作成 (SSH とアプリケーション用)
resource "google_compute_firewall" "allow_ssh" {
  name    = "allow-ssh"
  network = google_compute_network.vpc_network.name

  allow {
    protocol = "tcp"
    ports    = ["22"]
  }

  source_ranges = var.allowed_ssh_ips
  target_tags   = ["ssh-enabled"]
}

# ファイアウォールルールの作成 (アプリケーション用)
resource "google_compute_firewall" "allow_app" {
  name    = "allow-app"
  network = google_compute_network.vpc_network.name

  allow {
    protocol = "tcp"
    ports    = ["7203"] # Docker Compose でホスト側に公開しているポート
  }

  source_ranges = ["0.0.0.0/0"]
  target_tags   = ["http-server"]
}
