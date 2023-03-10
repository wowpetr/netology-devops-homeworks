# Домашнее задание к занятию 8.6 «Создание собственных модулей»

## Основная часть

Ваша цель — написать собственный module, который вы можете использовать в своей role через playbook. Всё это должно быть собрано в виде collection и отправлено в ваш репозиторий.

**Шаг 1.** В виртуальном окружении создайте новый `my_own_module.py` файл.

**Шаг 2.** Наполните его содержимым:

Возьмите это наполнение [из статьи](https://docs.ansible.com/ansible/latest/dev_guide/developing_modules_general.html#creating-a-module).

**Шаг 3.** Заполните файл в соответствии с требованиями Ansible так, чтобы он выполнял основную задачу: module должен создавать текстовый файл на удалённом хосте по пути, определённом в параметре `path`, с содержимым, определённым в параметре `content`.

**Шаг 4.** Проверьте module на исполняемость локально.
#### Решение
payload.json
```json
{
    "ANSIBLE_MODULE_ARGS": {
        "path": "/Users/plosev/example.txt",
        "content": "Created with ansible module create_file\n"

    }
}
```
Команда проверки:
```bash
python -m ansible.modules.create_file payload.json | jq .
```
Результат:
```json
{
  "changed": false,
  "path": "/Users/plosev/example.txt",
  "uid": 501,
  "gid": 20,
  "owner": "plosev",
  "group": "staff",
  "mode": "0644",
  "state": "file",
  "size": 57,
  "invocation": {
    "module_args": {
      "path": "/Users/plosev/example.txt",
      "content": "Created with ansible module create_file\n"
    }
  }
}
```

**Шаг 5.** Напишите single task playbook и используйте module в нём.

**Шаг 6.** Проверьте через playbook на идемпотентность.
#### Решение

Написал и использовал Playbook:
```yml
---
- name: Test module
  hosts: localhost
  tasks:
    - name: Copy file
      create_file:
        path: "~/example.txt"
        content: "Created with create_file ansible module\n"
```
Запуск с удаленным файлом:
```
~/temp-netology-homeworks/08-ansible-06-module/ansible ❯ ansible-playbook test_module_1.yml
[WARNING]: No inventory was parsed, only implicit localhost is available
[WARNING]: provided hosts list is empty, only localhost is available. Note that the implicit localhost does not match 'all'

PLAY [Test module] ***********************************************************************************************************************************************************************************

TASK [Gathering Facts] *******************************************************************************************************************************************************************************
ok: [localhost]

TASK [Copy file] *************************************************************************************************************************************************************************************
changed: [localhost]

PLAY RECAP *******************************************************************************************************************************************************************************************
localhost                  : ok=2    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   
```
Повторный запуск:
```
~/temp-netology-homeworks/08-ansible-06-module/ansible ❯ ansible-playbook test_module_1.yml
[WARNING]: No inventory was parsed, only implicit localhost is available
[WARNING]: provided hosts list is empty, only localhost is available. Note that the implicit localhost does not match 'all'

PLAY [Test module] ***********************************************************************************************************************************************************************************

TASK [Gathering Facts] *******************************************************************************************************************************************************************************
ok: [localhost]

TASK [Copy file] *************************************************************************************************************************************************************************************
ok: [localhost]

PLAY RECAP *******************************************************************************************************************************************************************************************
localhost                  : ok=2    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
```

**Шаг 7.** Выйдите из виртуального окружения.

**Шаг 8.** Инициализируйте новую collection: `ansible-galaxy collection init my_own_namespace.yandex_cloud_elk`.

**Шаг 9.** В эту collection перенесите свой module в соответствующую директорию.

**Шаг 10.** Single task playbook преобразуйте в single task role и перенесите в collection. У role должны быть default всех параметров module.

**Шаг 11.** Создайте playbook для использования этой role.

**Шаг 12.** Заполните всю документацию по collection, выложите в свой репозиторий, поставьте тег `1.0.0` на этот коммит.

**Шаг 13.** Создайте .tar.gz этой collection: `ansible-galaxy collection build` в корневой директории collection.

**Шаг 14.** Создайте ещё одну директорию любого наименования, перенесите туда single task playbook и архив c collection.

**Шаг 15.** Установите collection из локального архива: `ansible-galaxy collection install <archivename>.tar.gz`.
#### Решение
```bash
~/temp-netology-homeworks/08-ansible-06-module/wowpetr/devops1 ❯ ansible-galaxy collection install wowpetr-devops1-1.0.0.tar.gz
Starting galaxy collection install process
Process install dependency map
Starting collection install process
Installing 'wowpetr.devops1:1.0.0' to '/Users/plosev/.ansible/collections/ansible_collections/wowpetr/devops1'
wowpetr.devops1:1.0.0 was installed successfully
```

**Шаг 16.** Запустите playbook, убедитесь, что он работает.
#### Решение
```bash
~/temp-netology-homeworks/08-ansible-06-module/wowpetr/devops1 ❯ ansible-playbook wowpetr.devops1
[WARNING]: No inventory was parsed, only implicit localhost is available
[WARNING]: provided hosts list is empty, only localhost is available. Note that the implicit localhost does not match 'all'

PLAY [Test module] ***********************************************************************************************************************************************************************************

TASK [Gathering Facts] *******************************************************************************************************************************************************************************
ok: [localhost]

TASK [wowpetr.devops1.test-create-file : Copy text file] *********************************************************************************************************************************************
changed: [localhost]

PLAY RECAP *******************************************************************************************************************************************************************************************
localhost                  : ok=2    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
```

**Шаг 17.** В ответ необходимо прислать ссылки на collection и tar.gz архив, а также скриншоты выполнения пунктов 4, 6, 15 и 16.
#### Решение
- [Ссылка на репозиторий коллекции](https://github.com/wowpetr/devops1-ansible-collection/tree/1.0.0)
- [wowpetr-devops1-1.0.0.tar.gz](wowpetr-devops1-1.0.0.tar.gz)