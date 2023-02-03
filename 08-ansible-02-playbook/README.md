# Домашнее задание к занятию "8.2. Работа с Playbook"

## Подготовка к выполнению

1. (Необязательно) Изучите, что такое [clickhouse](https://www.youtube.com/watch?v=fjTNS2zkeBs) и [vector](https://www.youtube.com/watch?v=CgEhyffisLY)
2. Создайте свой собственный (или используйте старый) публичный репозиторий на github с произвольным именем.
3. Скачайте [playbook](./playbook/) из репозитория с домашним заданием и перенесите его в свой репозиторий.
4. Подготовьте хосты в соответствии с группами из предподготовленного playbook.

## Основная часть

1. Приготовьте свой собственный inventory файл `prod.yml`.
2. Допишите playbook: нужно сделать ещё один play, который устанавливает и настраивает [vector](https://vector.dev).
3. При создании tasks рекомендую использовать модули: `get_url`, `template`, `unarchive`, `file`.
4. Tasks должны: скачать нужной версии дистрибутив, выполнить распаковку в выбранную директорию, установить vector.
5. Запустите `ansible-lint site.yml` и исправьте ошибки, если они есть.
6. Попробуйте запустить playbook на этом окружении с флагом `--check`.
7. Запустите playbook на `prod.yml` окружении с флагом `--diff`. Убедитесь, что изменения на системе произведены.
8. Повторно запустите playbook с флагом `--diff` и убедитесь, что playbook идемпотентен.
9. Подготовьте README.md файл по своему playbook. В нём должно быть описано: что делает playbook, какие у него есть параметры и теги.
10. Готовый playbook выложите в свой репозиторий, поставьте тег `08-ansible-02-playbook` на фиксирующий коммит, в ответ предоставьте ссылку на него.

### Решение

Модифицировал playbook для установки и настройки `vector`:

```yaml
---
- name: Install Clickhouse
  hosts: clickhouse
  gather_facts: true
  tags: "clickhouse"
  handlers:
    - name: Start clickhouse service
      become: true
      ansible.builtin.service:
        name: clickhouse-server
        state: restarted
  tasks:
    - name: Get clickhouse distrib
      ansible.builtin.get_url:
        url: "https://packages.clickhouse.com/rpm/stable/{{ item }}-{{ clickhouse_version }}.{{ ansible_architecture }}.rpm"
        dest: "./{{ item }}-{{ clickhouse_version }}.rpm"
      with_items: "{{ clickhouse_packages }}"
    - name: Install clickhouse packages
      become: true
      ansible.builtin.yum:
        name:
          - clickhouse-common-static-{{ clickhouse_version }}.rpm
          - clickhouse-client-{{ clickhouse_version }}.rpm
          - clickhouse-server-{{ clickhouse_version }}.rpm
        disable_gpg_check: true
      notify: Start clickhouse service
    - name: Flush handlers
      ansible.builtin.meta: flush_handlers
    - name: Create database
      ansible.builtin.command: "clickhouse-client -q 'create database logs;'"
      register: create_db
      failed_when: create_db.rc != 0 and create_db.rc != 82
      changed_when: create_db.rc == 0

- name: Install Vector
  hosts: vector
  gather_facts: true
  tags: "vector"

  handlers:
    - name: Start vector service
      become: true
      ansible.builtin.service:
        name: vector
        state: restarted

  tasks:

    - name: Download Vector archive
      ansible.builtin.get_url:
        url: "https://packages.timber.io/vector/{{ vector_version_major }}/vector-{{ ansible_architecture }}-unknown-linux-musl.tar.gz"
        dest: "./vector-{{ ansible_architecture }}-unknown-linux-musl.tar.gz"

    - name: Ensure group 'vector' exists
      become: true
      ansible.builtin.group:
        name: vector

    - name: Add the user 'vector'
      become: true
      ansible.builtin.user:
        name: vector
        group: vector

    - name: Create the app directory for Vector
      become: true
      ansible.builtin.file:
        path: /opt/vector
        state: directory
        owner: vector
        group: vector

    - name: Create the data directory for Vector
      become: true
      ansible.builtin.file:
        path: /var/lib/vector
        state: directory
        owner: vector
        group: vector

    - name: Create the config directory for Vector
      become: true
      ansible.builtin.file:
        path: /etc/vector
        state: directory
        owner: vector
        group: vector

    - name: Extract vector
      become: true
      ansible.builtin.unarchive:
        src: "/home/wp/vector-{{ ansible_architecture }}-unknown-linux-musl.tar.gz"
        dest: /opt/vector
        remote_src: true
        extra_opts: [--strip-components=2]
        owner: vector
        group: vector

    - name: Template a file to /etc/vector/vector.toml
      become: true
      ansible.builtin.template:
        src: ./template/vector.toml
        dest: /etc/vector/vector.toml
        owner: root
        group: vector
        mode: '0644'
      notify: Start vector service

    - name: Copy vector executable file to /usr/bin
      become: true
      ansible.builtin.copy:
        src: /opt/vector/bin/vector
        dest: /usr/bin
        remote_src: true
        mode: preserve

    - name: Copy systemd service
      become: true
      ansible.builtin.copy:
        src: /opt/vector/etc/systemd/vector.service
        dest: /etc/systemd/system
        remote_src: true
        mode: preserve

```
#### Ссылка на playbook
[Ссылка](./playbook/)

#### Описание playbook

Playbook устанавливает на хосты описанные в inventory/prod.yml `Clickhouse` и `vector`:

inventory/prod.yml
```yaml
---
clickhouse:
  hosts:
    clickhouse-01:
      ansible_host: 51.250.90.17 
      ansible_user: wp
vector:
  hosts:
    vector-01:
      ansible_host: 84.201.134.234 
      ansible_user: wp
```

Данные хосты были развернуты в Yandex Cloud с помощью Terraform (код приложен [тут](./terraform/)).

В playbook использованы следующие теги: `clickhouse` и `vector`, которые соответственно позволяют запускать playbook отдельно для одного или другого хоста.

#### Используемые модули

* ansible.builtin.get_url (для закачки файлов установки)
* ansible.builtin.yum (установка clickhouse)
* ansible.builtin.service (запуск сервисов)
* ansible.builtin.command (создание базы данных clickhouse)
* ansible.builtin.group (создание группы пользователей vector)
* ansible.builtin.user (создание пользователя vector)
* ansible.builtin.file (создание директорий для vector)
* ansible.builtin.unarchive (распаковка архива vector)
* ansible.builtin.template (установка конфига /etc/vector/vector.toml)
* ansible.builtin.copy (копирование исполняемого файла vector и его systemd сервиса в требуемые директории)