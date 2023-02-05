# Домашнее задание к занятию "8.3. Использование Yandex Cloud"

## Подготовка к выполнению

1. Подготовьте в Yandex Cloud три хоста: для `clickhouse`, для `vector` и для `lighthouse`.

Ссылка на репозиторий LightHouse: https://github.com/VKCOM/lighthouse

## Основная часть

1. Допишите playbook: нужно сделать ещё один play, который устанавливает и настраивает lighthouse.
2. При создании tasks рекомендую использовать модули: `get_url`, `template`, `yum`, `apt`.
3. Tasks должны: скачать статику lighthouse, установить nginx или любой другой webserver, настроить его конфиг для открытия lighthouse, запустить webserver.
4. Приготовьте свой собственный inventory файл `prod.yml`.
5. Запустите `ansible-lint site.yml` и исправьте ошибки, если они есть.
6. Попробуйте запустить playbook на этом окружении с флагом `--check`.
7. Запустите playbook на `prod.yml` окружении с флагом `--diff`. Убедитесь, что изменения на системе произведены.
8. Повторно запустите playbook с флагом `--diff` и убедитесь, что playbook идемпотентен.
9. Подготовьте README.md файл по своему playbook. В нём должно быть описано: что делает playbook, какие у него есть параметры и теги.
10. Готовый playbook выложите в свой репозиторий, поставьте тег `08-ansible-03-yandex` на фиксирующий коммит, в ответ предоставьте ссылку на него.

### Решение

#### Ссылка на playbook
[Ссылка](./playbook/)


#### Ссылка на terraform
[Ссылка](./terraform/)

В playbook использованы следующие теги: `clickhouse`, `vector` и `lighthouse`, которые соответственно позволяют запускать playbook отдельно для соответствующего хоста.

#### Используемые модули

* ansible.builtin.yum (установка nginx)
* ansible.builtin.service (запуск сервисов)
* ansible.builtin.file (создание директорий для nginx и clickhouse, а также установки разрешений файлов и директорий)
* ansible.builtin.template (установка конфигов nginx)
* ansible.builtin.copy (копирование конфигов nginx, lighthouse)
* ansible.builtin.git (клонирование репозитория lighhouse)
* community.general.sefcontext (установка контеста SELinux)
* ansible.builtin.command (восстановление контекста SELinux)