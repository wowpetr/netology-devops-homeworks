resource "yandex_compute_instance" "clickhouse-01" {
  name        = "clickhause-01"
  platform_id = "standard-v3"

  resources {
    cores  = 2
    memory = 2
  }

  boot_disk {
    initialize_params {
      image_id = "fd8151sv1q69mchl804a" # centos-stream-8-v20230109
    }
  }

  network_interface {
    subnet_id = yandex_vpc_subnet.subnet-1.id
    nat = true
  }

  metadata = {
    user-data = "${file("./meta.txt")}"
  }
}

resource "yandex_compute_instance" "vector-01" {
  name        = "vector-01"
  platform_id = "standard-v3"

  resources {
    cores  = 2
    memory = 2
  }

  boot_disk {
    initialize_params {
      image_id = "fd8151sv1q69mchl804a" # centos-stream-8-v20230109
    }
  }

  network_interface {
    subnet_id = yandex_vpc_subnet.subnet-1.id
    nat = true
  }

  metadata = {
    user-data = "${file("./meta.txt")}"
  }
}

resource "yandex_compute_instance" "lighthouse-01" {
  name        = "lighthouse-01"
  platform_id = "standard-v3"

  resources {
    cores  = 2
    memory = 2
  }

  boot_disk {
    initialize_params {
      image_id = "fd8151sv1q69mchl804a" # centos-stream-8-v20230109
    }
  }

  network_interface {
    subnet_id = yandex_vpc_subnet.subnet-1.id
    nat = true
  }

  metadata = {
    user-data = "${file("./meta.txt")}"
  }
}

resource "yandex_vpc_network" "net-1" {
}

resource "yandex_vpc_subnet" "subnet-1" {
  network_id     = yandex_vpc_network.net-1.id
  v4_cidr_blocks = ["10.0.1.0/28"]
}

output "external_ip_addres_clickhouse-01" {
  value = yandex_compute_instance.clickhouse-01.network_interface.0.nat_ip_address
}
output "external_ip_addres_vector-01" {
  value = yandex_compute_instance.vector-01.network_interface.0.nat_ip_address
}
output "external_ip_addres_lighthouse-01" {
  value = yandex_compute_instance.lighthouse-01.network_interface.0.nat_ip_address
}