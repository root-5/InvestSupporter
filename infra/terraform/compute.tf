resource "google_compute_instance" "app_server" {
  name         = "invest-supporter-app"
  machine_type = var.machine_type
  zone         = var.zone

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-13-trixie-v20251209" # Debian 13
      size  = 10 # GB単位
      type  = "pd-standard" # 標準永続ディスク、最安(HDD)
    }
  }

  # インスタンスを VPC ネットワークに接続
  network_interface {
    network    = google_compute_network.vpc_network.id
    subnetwork = google_compute_subnetwork.subnet.id
    # access_config を空で定義すると GCP がエフェメラル IP を自動割当
    access_config {}
  }

  # OS Login を有効化
  # GCP の IAM アカウントを Linux VM（主に Compute Engine）の OS ユーザー認証に使う仕組み
  metadata = {
    enable-oslogin = "TRUE"
  }

  # スタートアップスクリプト、インスタンス起動時に一度だけ root 権限で実行される
  metadata_startup_script = <<-EOF
    #!/bin/bash
    # Docker のインストール
    # https://matsuand.github.io/docs.docker.jp.onthefly/engine/install/debian/#install-using-the-repository
    apt-get update
    apt-get install -y ca-certificates curl gnupg lsb-release
    curl -fsSL https://download.docker.com/linux/debian/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
    curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
    apt-get update
    apt-get install docker-ce docker-ce-cli containerd.io docker-compose-plugin
    # ユーザーを docker グループに追加、sudo なしで docker コマンドを実行可能にする
    groupadd docker
    usermod -aG docker $USER
    newgrp docker
    # 再起動後も docker を自動起動する設定は Debian の場合デフォルトで有効
  EOF

  # サービスアカウントの設定 (説明を見る限り不要そう、必要なら戻す)
  # service_account {
  #   scopes = ["cloud-platform"]
  # }
}
