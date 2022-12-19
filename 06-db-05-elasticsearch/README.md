# Домашнее задание к занятию "6.5. Elasticsearch"

## Задача 1

В этом задании вы потренируетесь в:
- установке elasticsearch
- первоначальном конфигурировании elastcisearch
- запуске elasticsearch в docker

Используя докер образ [elasticsearch:7](https://hub.docker.com/_/elasticsearch) как базовый:

- составьте Dockerfile-манифест для elasticsearch
- соберите docker-образ и сделайте `push` в ваш docker.io репозиторий
- запустите контейнер из получившегося образа и выполните запрос пути `/` c хост-машины

Требования к `elasticsearch.yml`:
- данные `path` должны сохраняться в `/var/lib` 
- имя ноды должно быть `netology_test`

В ответе приведите:
- текст Dockerfile манифеста
- ссылку на образ в репозитории dockerhub
- ответ `elasticsearch` на запрос пути `/` в json виде

Подсказки:
- при сетевых проблемах внимательно изучите кластерные и сетевые настройки в elasticsearch.yml
- при некоторых проблемах вам поможет docker директива ulimit
- elasticsearch в логах обычно описывает проблему и пути ее решения
- обратите внимание на настройки безопасности такие как `xpack.security.enabled` 
- если докер образ не запускается и падает с ошибкой 137 в этом случае может помочь настройка `-e ES_HEAP_SIZE`
- при настройке `path` возможно потребуется настройка прав доступа на директорию

Далее мы будем работать с данным экземпляром elasticsearch.

### Решение
Dockerfile
```docker
FROM elasticsearch:7.17.8
COPY --chown=elasticsearch:elasticsearch elasticsearch.yml /usr/share/elasticsearch/config/
COPY --chown=elasticsearch:elasticsearch heap.options /usr/share/elasticsearch/config/jvm.options.d/
RUN chown elasticsearch:root /var/lib
```
elasticsearch.yml
```yml
cluster.name: netology
node.name: netology_test 
node.roles: [ master, data, ingest ]
cluster.initial_master_nodes: [ netology_test ]
path.data: /var/lib
network.host: 0.0.0.0
http.port: 9200
discovery.seed_hosts: [ 127.0.0.1 ]
```
heap.options
```
-Xms1g
-Xmx1g
```
Сборка образа:
```bash
❯ docker build -t wowpetr/elasticsearch:7 .
```
Публикация образа:
```bash
❯ docker push wowpetr/elasticsearch:7
```
#### Ссылка на образ
https://hub.docker.com/r/wowpetr/elasticsearch

#### Ответ от Elasticsearch на запрос пути `/`
```bash
❯ docker run --rm -it -d --name elastic7 -p 9200:9200 wowpetr/elasticsearch:7              
cccd2d09fb87fb3a5b2e14fc8bc6311d3f40bca20e649eb5aedcc6c8fdb50ae1
❯ curl -X GET "localhost:9200/"
{
  "name" : "netology_test",
  "cluster_name" : "netology",
  "cluster_uuid" : "T1EnAj6RTl2T_U7ElSsmIQ",
  "version" : {
    "number" : "7.17.8",
    "build_flavor" : "default",
    "build_type" : "docker",
    "build_hash" : "120eabe1c8a0cb2ae87cffc109a5b65d213e9df1",
    "build_date" : "2022-12-02T17:33:09.727072865Z",
    "build_snapshot" : false,
    "lucene_version" : "8.11.1",
    "minimum_wire_compatibility_version" : "6.8.0",
    "minimum_index_compatibility_version" : "6.0.0-beta1"
  },
  "tagline" : "You Know, for Search"
}
```
#### Удаление всех индексов
```bash
❯ curl -X DELETE "localhost:9200/_all?pretty"                                                                                      ✘ INT
{
  "acknowledged" : true
}
```

## Задача 2

В этом задании вы научитесь:
- создавать и удалять индексы
- изучать состояние кластера
- обосновывать причину деградации доступности данных

Ознакомтесь с [документацией](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html) 
и добавьте в `elasticsearch` 3 индекса, в соответствии со таблицей:

| Имя | Количество реплик | Количество шард |
|-----|-------------------|-----------------|
| ind-1| 0 | 1 |
| ind-2 | 1 | 2 |
| ind-3 | 2 | 4 |

Получите список индексов и их статусов, используя API и **приведите в ответе** на задание.

Получите состояние кластера `elasticsearch`, используя API.

Как вы думаете, почему часть индексов и кластер находится в состоянии yellow?

Удалите все индексы.

**Важно**

При проектировании кластера elasticsearch нужно корректно рассчитывать количество реплик и шард,
иначе возможна потеря данных индексов, вплоть до полной, при деградации системы.

### Решение
#### Добавление индексов
```bash
❯ curl -X PUT "localhost:9200/ind-1?pretty" -H 'Content-Type: application/json' -d' { "settings": { "number_of_shards": 1, "number_of_replicas": 0 } } '
{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "ind-1"
}
❯ curl -X PUT "localhost:9200/ind-2?pretty" -H 'Content-Type: application/json' -d' { "settings": { "number_of_shards": 2, "number_of_replicas": 1 } } '
{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "ind-2"
}
❯ curl -X PUT "localhost:9200/ind-3?pretty" -H 'Content-Type: application/json' -d' { "settings": { "number_of_shards": 4, "number_of_replicas": 2 } } '
{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "ind-3"
}
```
#### Индексы с их статусами
```bash
❯ curl -X GET "localhost:9200/_cat/indices?pretty" 
green  open .geoip_databases Kgxr7VWFSIuM3ZV2J2wvdA 1 0 40 0 38.2mb 38.2mb
green  open ind-1            HbrFVFyaQjCUCxCmRhCIBg 1 0  0 0   226b   226b
yellow open ind-3            QTr0VZpCRXa_xhdD69hltA 4 2  0 0   904b   904b
yellow open ind-2            XL8wqEuwS662dcATBCZmpg 2 1  0 0   452b   452b
```
#### Состояние кластера
```bash
❯ curl -X GET "localhost:9200/_cluster/health?pretty"
{
  "cluster_name" : "netology",
  "status" : "yellow",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 10,
  "active_shards" : 10,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 10,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 50.0
}
```
Часть индексов в состоянии `Yellow` так как имеются `UNASSIGNED` реплики, которые не могут быть назначены поскольку у нас кластер состоит из одного узла:
```bash
❯ curl -X GET "localhost:9200/_cat/shards?pretty" 
.ds-ilm-history-5-2022.12.19-000001                           0 p STARTED              172.17.0.2 netology_test
ind-1                                                         0 p STARTED     0   226b 172.17.0.2 netology_test
.ds-.logs-deprecation.elasticsearch-default-2022.12.19-000001 0 p STARTED              172.17.0.2 netology_test
ind-2                                                         1 p STARTED     0   226b 172.17.0.2 netology_test
ind-2                                                         1 r UNASSIGNED                      
ind-2                                                         0 p STARTED     0   226b 172.17.0.2 netology_test
ind-2                                                         0 r UNASSIGNED                      
.geoip_databases                                              0 p STARTED    40 38.2mb 172.17.0.2 netology_test
ind-3                                                         2 p STARTED     0   226b 172.17.0.2 netology_test
ind-3                                                         2 r UNASSIGNED                      
ind-3                                                         2 r UNASSIGNED                      
ind-3                                                         1 p STARTED     0   226b 172.17.0.2 netology_test
ind-3                                                         1 r UNASSIGNED                      
ind-3                                                         1 r UNASSIGNED                      
ind-3                                                         3 p STARTED     0   226b 172.17.0.2 netology_test
ind-3                                                         3 r UNASSIGNED                      
ind-3                                                         3 r UNASSIGNED                      
ind-3                                                         0 p STARTED     0   226b 172.17.0.2 netology_test
ind-3                                                         0 r UNASSIGNED                      
ind-3                                                         0 r UNASSIGNED
```
## Задача 3

В данном задании вы научитесь:
- создавать бэкапы данных
- восстанавливать индексы из бэкапов

Создайте директорию `{путь до корневой директории с elasticsearch в образе}/snapshots`.

Используя API [зарегистрируйте](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-register-repository.html#snapshots-register-repository) 
данную директорию как `snapshot repository` c именем `netology_backup`.

**Приведите в ответе** запрос API и результат вызова API для создания репозитория.

Создайте индекс `test` с 0 реплик и 1 шардом и **приведите в ответе** список индексов.

[Создайте `snapshot`](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-take-snapshot.html) 
состояния кластера `elasticsearch`.

**Приведите в ответе** список файлов в директории со `snapshot`ами.

Удалите индекс `test` и создайте индекс `test-2`. **Приведите в ответе** список индексов.

[Восстановите](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-restore-snapshot.html) состояние
кластера `elasticsearch` из `snapshot`, созданного ранее. 

**Приведите в ответе** запрос к API восстановления и итоговый список индексов.

Подсказки:
- возможно вам понадобится доработать `elasticsearch.yml` в части директивы `path.repo` и перезапустить `elasticsearch`

### Решение

#### Регистрация репозитория
Дополнил elasticsearch.yml путем для репозиториев:
```yml
path.repo: ["/mount/backups"]
```
Дополнил Dockerfile командами создания пути репозитория:
```docker
RUN mkdir -p /mount/backups
RUN chown elasticsearch:elasticsearch /mount/backups
```
После сборки нового образа и запуска контейнера произвел регистрацию репозитория:

```bash
❯ curl -X PUT "localhost:9200/_snapshot/netology_backup?pretty" -H 'Content-Type: application/json' -d'
{
  "type": "fs",
  "settings": {
    "location": "/mount/backups/snapshots"
  }
}
'

{
  "acknowledged" : true
}
```
#### Создание индекса `test` 
```bash
❯ curl -X PUT "localhost:9200/test?pretty" -H 'Content-Type: application/json' -d' { "settings": { "number_of_shards": 1, "number_of_replicas": 0 } } '
{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "test"
}
❯ curl -X GET "localhost:9200/_cat/indices?pretty"
green open .geoip_databases xMXnElRARLei4U4tUJFufQ 1 0 40 0 38.2mb 38.2mb
green open test             ReoqnVUlSb-PEoHamZry_g 1 0  0 0   226b   226b
```
#### Создание snapshot'a кластера
```bash
❯ curl -X PUT "localhost:9200/_snapshot/netology_backup/%3Cmy_snapshot_%7Bnow%2Fd%7D%3E?wait_for_completion=true&pretty"
{
  "snapshot" : {
    "snapshot" : "my_snapshot_2022.12.19",
    "uuid" : "X5bnxVTcSiOsdEwz8saEow",
    "repository" : "netology_backup",
    "version_id" : 7170899,
    "version" : "7.17.8",
    "indices" : [
      "test",
      ".ds-ilm-history-5-2022.12.19-000001",
      ".geoip_databases",
      ".ds-.logs-deprecation.elasticsearch-default-2022.12.19-000001"
    ],
    "data_streams" : [
      "ilm-history-5",
      ".logs-deprecation.elasticsearch-default"
    ],
    "include_global_state" : true,
    "state" : "SUCCESS",
    "start_time" : "2022-12-19T21:35:54.837Z",
    "start_time_in_millis" : 1671485754837,
    "end_time" : "2022-12-19T21:35:55.838Z",
    "end_time_in_millis" : 1671485755838,
    "duration_in_millis" : 1001,
    "failures" : [ ],
    "shards" : {
      "total" : 4,
      "failed" : 0,
      "successful" : 4
    },
    "feature_states" : [
      {
        "feature_name" : "geoip",
        "indices" : [
          ".geoip_databases"
        ]
      }
    ]
  }
}
```
#### Список файлов в директории после создания snapshot'а
```bash
❯ docker exec -it elastic7 ls -lh /mount/backups/snapshots
total 48K
-rw-rw-r-- 1 elasticsearch root 1.5K Dec 19 21:35 index-0
-rw-rw-r-- 1 elasticsearch root    8 Dec 19 21:35 index.latest
drwxrwxr-x 6 elasticsearch root 4.0K Dec 19 21:35 indices
-rw-rw-r-- 1 elasticsearch root  29K Dec 19 21:35 meta-X5bnxVTcSiOsdEwz8saEow.dat
-rw-rw-r-- 1 elasticsearch root  721 Dec 19 21:35 snap-X5bnxVTcSiOsdEwz8saEow.dat
```
#### Удаление индекса `test` и создание индекса `test2`
```bash
❯ curl -X DELETE "localhost:9200/test?pretty"           
{
  "acknowledged" : true
}
```
```bash
❯ curl -X PUT "localhost:9200/test2?pretty" -H 'Content-Type: application/json' -d' { "settings": { "number_of_shards": 1, "number_of_replicas": 0 } } '
{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "test2"
}
❯ curl -X GET "localhost:9200/_cat/indices?pretty"
green open .geoip_databases xMXnElRARLei4U4tUJFufQ 1 0 40 0 38.2mb 38.2mb
green open test2            Z-GInPoWTaCVPBeRxMyLTg 1 0  0 0   226b   226b
```
#### Восстановление кластера в исходное состояние из snapshot'а
##### Подготовительные действия перед восстановлением (отключение индексации и отключение различных служб elasticsearch)
```bash
❯ curl -X PUT "localhost:9200/_cluster/settings?pretty" -H 'Content-Type: application/json' -d'
{
  "persistent": {
    "ingest.geoip.downloader.enabled": false
  }
}
'

{
  "acknowledged" : true,
  "persistent" : {
    "ingest" : {
      "geoip" : {
        "downloader" : {
          "enabled" : "false"
        }
      }
    }
  },
  "transient" : { }
}
❯ curl -X POST "localhost:9200/_ilm/stop?pretty"

{
  "acknowledged" : true
}
❯ curl -X POST "localhost:9200/_ml/set_upgrade_mode?enabled=true&pretty"

{
  "acknowledged" : true
}
❯ curl -X PUT "localhost:9200/_cluster/settings?pretty" -H 'Content-Type: application/json' -d'
{
  "persistent": {
    "xpack.monitoring.collection.enabled": false
  }
}
'

{
  "acknowledged" : true,
  "persistent" : {
    "xpack" : {
      "monitoring" : {
        "collection" : {
          "enabled" : "false"
        }
      }
    }
  },
  "transient" : { }
}
❯ curl -X POST "localhost:9200/_watcher/_stop?pretty"

{
  "acknowledged" : true
}
❯ curl -X PUT "localhost:9200/_cluster/settings?pretty" -H 'Content-Type: application/json' -d'
{
  "persistent": {
    "action.destructive_requires_name": false
  }
}
'

{
  "acknowledged" : true,
  "persistent" : {
    "action" : {
      "destructive_requires_name" : "false"
    }
  },
  "transient" : { }
}
❯ curl -X DELETE "localhost:9200/_data_stream/*?expand_wildcards=all&pretty"

{
  "acknowledged" : true
}
❯ curl -X DELETE "localhost:9200/*?expand_wildcards=all&pretty"

{
  "acknowledged" : true
}
```
##### Команда восстановления
```bash
❯ curl -X POST "localhost:9200/_snapshot/netology_backup/my_snapshot_2022.12.19/_restore?pretty" -H 'Content-Type: application/json' -d'
{
  "indices": "*",
  "include_global_state": true
}
'

{
  "accepted" : true
}
```
##### Действия после восстановления 
```bash
❯ curl -X PUT "localhost:9200/_cluster/settings?pretty" -H 'Content-Type: application/json' -d'
{
  "persistent": {
    "ingest.geoip.downloader.enabled": true
  }
}
'

{
  "acknowledged" : true,
  "persistent" : {
    "ingest" : {
      "geoip" : {
        "downloader" : {
          "enabled" : "true"
        }
      }
    }
  },
  "transient" : { }
}
❯ curl -X POST "localhost:9200/_ilm/start?pretty"

{
  "acknowledged" : true
}
❯ curl -X POST "localhost:9200/_ml/set_upgrade_mode?enabled=false&pretty"

{
  "acknowledged" : true
}
❯ curl -X PUT "localhost:9200/_cluster/settings?pretty" -H 'Content-Type: application/json' -d'
{
  "persistent": {
    "xpack.monitoring.collection.enabled": true
  }
}
'

{
  "acknowledged" : true,
  "persistent" : {
    "xpack" : {
      "monitoring" : {
        "collection" : {
          "enabled" : "true"
        }
      }
    }
  },
  "transient" : { }
}
❯ curl -X POST "localhost:9200/_watcher/_start?pretty"

{
  "acknowledged" : true
}
❯ curl -X PUT "localhost:9200/_cluster/settings?pretty" -H 'Content-Type: application/json' -d'
{
  "persistent": {
    "action.destructive_requires_name": null
  }
}
'

{
  "acknowledged" : true,
  "persistent" : { },
  "transient" : { }
}
```
 ##### Индексы после восстановления
```bash
❯ curl -X GET "localhost:9200/_cat/indices?pretty"
green open .geoip_databases            KkgOdxp-T8yZCLM23Wqh9A 1 0 40 0  38.2mb  38.2mb
green open test                        K6L4QmkXQuWU5zDImqtLPw 1 0  0 0    226b    226b
```
Как можно видеть, состояние кластера на момент сделанного snapshot'а возвращено и поэтому индекс `test2` отсутствует.