# Ansible Playbook: site.yml

- [Tags](#tags)
- [Description](#description)
- [Requirements](#requirements)
- [Configure](#configure)
- [Install](#install)
- [License](#license)

## Tags
* 08-ansible-03-yandex
* 08-ansible-03-yandex-1 (latest)

## Description

Deploy [Clickhouse](https://github.com/ClickHouse/ClickHouse), [Vector](https://github.com/vectordotdev/vector) and [Lighthouse](https://github.com/VKCOM/lighthouse) stack using ansible as part of the homework for the DevOps Netology course.

## Requirements

- Ansible >= 2.14 (It might work on previous versions, but it is not guaranteed).

## Configure 

### Group Variables

#### All

| Name           | Default Value | Description                        |
| -------------- | ------------- | -----------------------------------|
| `ansible_ssh_common_args` | -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null | SSH common args |

#### Clickhouse 

| Name           | Default Value | Description                        |
| -------------- | ------------- | -----------------------------------|
| `clickhouse_version` | 23.1.2.9 | Clickhouse version |
| clickhouse_packages | [clickhouse-client, clickhouse-server, clickhouse-common-static] | Clickhouse packages to deploy |

#### Vector

| Name           | Default Value | Description                        |
| -------------- | ------------- | -----------------------------------|
| vector_version_major | 0.27.0 | Vector major version |

#### Lighthouse

| Name           | Default Value | Description                        |
| -------------- | ------------- | -----------------------------------|
| domain | mydomain.org | nginx site name |
| nginx_user | nginx | nginx user |

### Lighthouse config
Lighthouse config should be set via [config/lh.xml](config/lh.xml)
```xml
<?xml version="1.0"?>
<clickhouse>
  <listen_host>::</listen_host>
  <timezone>UTC</timezone>
</clickhouse>
```

### Nginx common template
[nginx.conf](./template/nginx.conf.j2)

### Nginx lighthouse site template
[site.yml](./template/lighthouse.conf.j2)

### Vector config template
[vector.toml](./template/vector.toml)

### Multi-node Installation
By default, this playbook will install only one node for each inventory group (clickhouse, vector and lighthouse).

inventory/prod.yml

```yaml
---
clickhouse:
  hosts:
    clickhouse-01:
vector:
  hosts:
    vector-01:
lighthouse:
  hosts:
    lighthouse-01:

```
## Install
```bash
ansible-playbook -i inventory/prod.yml site.yml
```

## License

This project is licensed under MIT License. See [LICENSE](./LICENSE) for more details.