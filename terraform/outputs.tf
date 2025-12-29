# `terraform apply` 実行後や `terraform output` で参照可能な出力値の定義
# インフラを作成するという目的だと定義不要なファイル
# 開発者の利便性向上や json で取得しての自動化、CI/CD などで利用する場合に便利
# `terraform output <output_name>` で個別に参照可能
# json 形式で全て取得したい場合は `terraform output -json` を利用
# コンソール出力、tfstate ファイル、コマンド履歴などに平文が残ってしまうため、機密情報は出力しないこと
output "instance_name" {
  value = google_compute_instance.app_server.name
}

output "external_ip" {
  value = google_compute_instance.app_server.network_interface[0].access_config[0].nat_ip
}

output "ssh_command" {
  value = "gcloud compute ssh ${google_compute_instance.app_server.name} --zone ${var.zone} --tunnel-through-iap"
}
