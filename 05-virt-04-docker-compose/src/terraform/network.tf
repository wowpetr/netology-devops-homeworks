# Network
resource "yandex_vpc_network" "netology-devops-network" {
  name = "net"
}

resource "yandex_vpc_subnet" "default" {
  name = "subnet"
  zone           = "ru-central1-b"
  network_id     = "${yandex_vpc_network.netology-devops-network.id}"
  v4_cidr_blocks = ["192.168.0.0/24"]
}
