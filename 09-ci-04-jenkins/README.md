# Домашнее задание к занятию 10 «Jenkins»

## Подготовка к выполнению

1. Создать два VM: для jenkins-master и jenkins-agent.
2. Установить Jenkins при помощи playbook.
3. Запустить и проверить работоспособность.
4. Сделать первоначальную настройку.

## Основная часть

1. Сделать Freestyle Job, который будет запускать `molecule test` из любого вашего репозитория с ролью.
### Решение:
![](./img/freestyle-prop1.png)
![](./img/freestyle-prop2.png)
![](./img/freestyle-output1.png)
![](./img/freestyle-output2.png)
![](./img/freestyle-status.png)

2. Сделать Declarative Pipeline Job, который будет запускать `molecule test` из любого вашего репозитория с ролью.
### Решение:
![](./img/decpipeline-script.png)
![](./img/dpipeline-status.png)

3. Перенести Declarative Pipeline в репозиторий в файл `Jenkinsfile`.
### Решение:
![](./img/dpipeline-scm-prop.png)
![](./img/dpipeline-scm-status.png)

4. Создать Multibranch Pipeline на запуск `Jenkinsfile` из репозитория.
### Решение:
![](./img/mpipeline-prop.png)
![](./img/mpipeline-status.png)
![](./img/mpipeline-scanlog.png)

5. Создать Scripted Pipeline, наполнить его скриптом из [pipeline](./pipeline).
6. Внести необходимые изменения, чтобы Pipeline запускал `ansible-playbook` без флагов `--check --diff`, если не установлен параметр при запуске джобы (prod_run = True). По умолчанию параметр имеет значение False и запускает прогон с флагами `--check --diff`.
7. Проверить работоспособность, исправить ошибки, исправленный Pipeline вложить в репозиторий в файл `ScriptedJenkinsfile`.
8. Отправить ссылку на репозиторий с ролью и Declarative Pipeline и Scripted Pipeline.

## Необязательная часть

1. Создать скрипт на groovy, который будет собирать все Job, завершившиеся хотя бы раз неуспешно. Добавить скрипт в репозиторий с решением и названием `AllJobFailure.groovy`.
2. Создать Scripted Pipeline так, чтобы он мог сначала запустить через Yandex Cloud CLI необходимое количество инстансов, прописать их в инвентори плейбука и после этого запускать плейбук. Мы должны при нажатии кнопки получить готовую к использованию систему.
