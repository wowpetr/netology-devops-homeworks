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

#### Описание playbook (play установки lighthouse)
```yaml
- name: Install Lighthouse
  become: true # повысить привилегии для всех задач
  hosts: lighthouse # по всем хостам группы lighthouse
  gather_facts: true # собирать факты
  tags: "lighthouse" # тег для маркировки задач хостов lighthouse

  pre_tasks: # вспомогательные задачи

    - name: Install nginx # установка nginx
      ansible.builtin.yum:
        name: nginx # имя пакета из репозитория
        state: present # обеспечить присутствие пакета

    - name: Copy custom nginx.conf # применение конфига nginx.conf
      ansible.builtin.template:
        src: ./template/nginx.conf.j2 # локальный путь до шаблона
        dest: /etc/nginx/nginx.conf # удаленный путь на целевой машине
      notify: Restart nginx # запрос перезапуска nginx

    - name: Set up nginx config directory # создание директорий sites-available и sites-enabled 
      ansible.builtin.file:
        path: /etc/nginx/{{ item }} # создавать директории в /etc/nginx
        state: directory # директории
      with_items: # цикл по указанным ниже директориям
        - sites-available
        - sites-enabled

    - name: Copy nginx lighthouse.conf # применить шаблон конфига для расположения (site) nginx
      ansible.builtin.template:
        src: ./template/lighthouse.conf.j2 # локальный путь до шаблона
        dest: /etc/nginx/sites-available/{{ domain }} # удаленный путь
        owner: '{{ nginx_user }}' # владелец файла конфига
        group: '{{ nginx_user }}' # группа файла конфига
        mode: 0644 # права доступа
      notify: Restart nginx # запрос перезапуска nginx

    - name: Activate the lighthouse nginx site # добавление симлинка на расположение сайта
      ansible.builtin.file:
        src: /etc/nginx/sites-available/{{ domain }} # источник для симлинка
        dest: /etc/nginx/sites-enabled/{{ domain }} # имя симлинка
        state: link # обеспечить присутствие линка

    - name: Install git # установка git
      ansible.builtin.yum:
        name: git # имя пакета
        state: present # обеспечить присутствие пакета

  handlers: # обработчики
    - name: Restart nginx # перезапуск сервиса nginx
      ansible.builtin.service:
        name: nginx # имя сервиса
        state: restarted # требовать перезапуска

  tasks: # основные задачи

    - name: Create www directory # создание директории расположения сайта /var/www/mydomain.com
      ansible.builtin.file:
        path: /var/www/{{ domain }} # путь создаваемой директории
        state: directory # что директория
        mode: 0775 # права доступа

    - name: Set SELinux context # добавление контекста SELinux для /var/www
      community.general.sefcontext:
        target: '/var/www/(/.*)?' # шаблон контекста
        setype: httpd_sys_content_t # тип контекста
        state: present # обеспечить присутствие контекста

    - name: Clone lighthouse # клонирование репозитория lighthouse
      ansible.builtin.git:
        repo: 'https://github.com/VKCOM/lighthouse.git' # репозиторий
        dest: /var/www/{{ domain }} # путь для клонирования
        version: master # ветка
        update: false # не обновлять

    - name: Set the permissions to all www files and directories # установка владельца и разрешений 755 на директории и 644 на файлы размещения /var/www
      ansible.builtin.file:
        path: /var/www # путь для установки разрешений
        recurse: true # рекурсивно 
        owner: '{{ nginx_user }}' # пользователь nginx
        group: '{{ nginx_user }}' # группа nginx
        mode: u=rwX,g=rX,o=rX # права доступа

    - name: Apply SELinux file context to www # восстановление SELinux разрешений для /var/www
      ansible.builtin.command: restorecon -irv /var/www # команда восстановления
```



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