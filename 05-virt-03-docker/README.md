
# Домашнее задание к занятию "3. Введение. Экосистема. Архитектура. Жизненный цикл Docker контейнера"

## Задача 1

Сценарий выполнения задачи:

- создайте свой репозиторий на https://hub.docker.com;
- выберете любой образ, который содержит веб-сервер Nginx;
- создайте свой fork образа;
- реализуйте функциональность:
запуск веб-сервера в фоне с индекс-страницей, содержащей HTML-код ниже:
```
<html>
<head>
Hey, Netology
</head>
<body>
<h1>I’m DevOps Engineer!</h1>
</body>
</html>
```
Опубликуйте созданный форк в своем репозитории и предоставьте ответ в виде ссылки на https://hub.docker.com/username_repo.

### Решение

1. Зарегистрировал аккаунт на https://hub.docker.com/.
2. Выбрал официальный образ `nginx:1.32.2`.
3. Создал файлы:  
    Dockerfile
    ```Dockerfile
    FROM nginx:1.23.2
    COPY ./index.html /usr/share/nginx/html/index.html
    ```
    index.html
    ```html
    <html>
    <head>
    Hey, Netology
    </head>
    <body>
    <h1>I'm DevOps Engineer!</h1>
    </body>
    </html>
    ```
    Выполнил сборку образа командой:
    ```bash
    docker build -t wowpetr/nginx:1.32.2 .
    ```
    Опубликовал образ:
    ```bash
    docker push wowpetr/nginx:1.32.2
    ```
4. Запустил созданный образ и проверил работоспособность сайта:
    ```bash
    ~/temp-netology-homeworks/05-virt-03-docker ❯ docker run -it --rm -d -p 8080:80 wowpetr/nginx:1.32.2
    ~/temp-netology-homeworks/05-virt-03-docker ❯ curl localhost:8080
    <html>
    <head>
    Hey, Netology
    </head>
    <body>
    <h1>I'm DevOps Engineer!</h1>
    </body>
    </html>
    ```

Ссылка на созданный репозиторий с загруженным образом:  
https://hub.docker.com/r/wowpetr/nginx

## Задача 2

Посмотрите на сценарий ниже и ответьте на вопрос:
"Подходит ли в этом сценарии использование Docker контейнеров или лучше подойдет виртуальная машина, физическая машина? Может быть возможны разные варианты?"

Детально опишите и обоснуйте свой выбор.

Сценарий:

- Высоконагруженное монолитное java веб-приложение;
    ##### Решение
    Использовать Docker возможно так как то, что оно высоконагруженное никак не влияет поскольку Docker не снижает быстродействие, а то что оно монолитное даже лучше так как будет один процесс. Можно использовать образ `openjdk` для этой цели.
  
- Nodejs веб-приложение;
    ##### Решение
    Использовать Docker возможно и даже существует необходимый образ `node` на базе которого это можно осуществить.
- Мобильное приложение c версиями для Android и iOS;
    ##### Решение
    Для iOS точно невозможно, так как это противоречит пользовательским соглашениям Apple. Что касается Android, то я находил статьи которые позволяют это сделать для Web приложений, например используя [docker-android](https://github.com/budtmo/docker-android) образы. Думаю, использование виртуальных машин более оправдано.
- Шина данных на базе Apache Kafka;
    ##### Решение
    Использование Docker возможно, существуют соответствующие образы на DockerHub (bitnami/kafka).
- Elasticsearch кластер для реализации логирования продуктивного веб-приложения - три ноды elasticsearch, два logstash и две ноды kibana;
    ##### Решение
    Развернуть все эти контейнеры можно так как существуют соответствующие образы на DockerHub, однако, потребуется настройка работы в режиме кластера.
- Мониторинг-стек на базе Prometheus и Grafana;
    ##### Решение
    Использование Docker возможно, существуют соответствующие образы на DockerHub.
- MongoDB, как основное хранилище данных для java-приложения;
    ##### Решение
    Использование Docker возможно, существуют соответствующие образы на DockerHub (mongo).
- Gitlab сервер для реализации CI/CD процессов и приватный (закрытый) Docker Registry.
    ##### Решение
    Использование Docker возможно, существуют соответствующие образы на DockerHub (gitlab-runner, registry).

## Задача 3

- Запустите первый контейнер из образа ***centos*** c любым тэгом в фоновом режиме, подключив папку ```/data``` из текущей рабочей директории на хостовой машине в ```/data``` контейнера;
- Запустите второй контейнер из образа ***debian*** в фоновом режиме, подключив папку ```/data``` из текущей рабочей директории на хостовой машине в ```/data``` контейнера;
- Подключитесь к первому контейнеру с помощью ```docker exec``` и создайте текстовый файл любого содержания в ```/data```;
- Добавьте еще один файл в папку ```/data``` на хостовой машине;
- Подключитесь во второй контейнер и отобразите листинг и содержание файлов в ```/data``` контейнера.

#### Решение
1. Запустил первый контейнер образа `centos:7`, подключив директорию data:  
```bash
~/temp-netology-homeworks/05-virt-03-docker ❯ docker run -it -d --rm -v $PWD/data:/data --name centos7 centos:7
c184dc6d8bb414ac4787d727eee32a6297b9db8d1a973839703faced0771dfe4
```
2. Запустил второй контейнер образа `debian:10`, подключив директорию data:  
```bash
~/temp-netology-homeworks/05-virt-03-docker ❯ docker run -it -d --rm -v $PWD/data:/data --name debian10 debian:10
Unable to find image 'debian:10' locally
10: Pulling from library/debian
3ba81f4c3c21: Pull complete
Digest: sha256:e83854c9fb469daa7273d73c43a3fe5ae2da77cb40d3d34282a9af09a9db49f9
Status: Downloaded newer image for debian:10
96d4256c33fe1614d4eb007cf62f3ec2fd012ecb1484b739843e8c656d51a61f
```
3. Вошел в первый контейнер и создал файл в директории /data:
```
~/temp-netology-homeworks/05-virt-03-docker ❯ docker exec -it centos7 bash
[root@c184dc6d8bb4 /]# echo "123456" > /data/test.txt
[root@c184dc6d8bb4 /]# cat /data/test.txt
123456
```
4. Создал еще один файл в директории /data:
```
~/temp-netology-homeworks/05-virt-03-docker ❯ cd data
~/temp-netology-homeworks/05-virt-03-docker/data ❯ echo "123" > test2.txt
```
5. Вошел в второй контейнер и отобразил содержимое директории /data:
```
~/temp-netology-homeworks/05-virt-03-docker/data ❯ docker exec -it debian10 bash
root@96d4256c33fe:/# cd /data
root@96d4256c33fe:/data# ls
test.txt  test2.txt
root@96d4256c33fe:/data# cat *
123456
123
```
    