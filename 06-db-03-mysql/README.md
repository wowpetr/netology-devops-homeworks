# Домашнее задание к занятию "6.3. MySQL"

## Задача 1

Используя docker поднимите инстанс MySQL (версию 8). Данные БД сохраните в volume.

Изучите [бэкап БД](https://github.com/netology-code/virt-homeworks/tree/master/06-db-03-mysql/test_data) и 
восстановитесь из него.

Перейдите в управляющую консоль `mysql` внутри контейнера.

Используя команду `\h` получите список управляющих команд.

Найдите команду для выдачи статуса БД и **приведите в ответе** из ее вывода версию сервера БД.

Подключитесь к восстановленной БД и получите список таблиц из этой БД.

**Приведите в ответе** количество записей с `price` > 300.

В следующих заданиях мы будем продолжать работу с данным контейнером.

### Решение

1. Поднял docker-контейнер MySQL (версия 8) используя docker-compose файл:
    ```yaml
    version: '3.1'

    services:

    db:
        image: mysql:8
        container_name: mysql-db
        ports:
        - 3306:3306
        volumes:
        - ./mysql_data:/var/lib/mysql
        - ./mysql_backup:/backup
        restart: always
        environment:
        MYSQL_ROOT_PASSWORD: 123
        MYSQL_DATABASE: test_db

    ```
2. Подключился к созданному контейнеру и восстановился из файла резервной копии:
    ```bash
    ❯ docker exec -it mysql-db bash
    bash-4.4# mysql --password=123 < /backup/test_dump.sql test_db
    mysql: [Warning] Using a password on the command line interface can be insecure.
    ```  
3. Подключился к восстановленной базе и получил статус БД и версию сервера:
    ```
    ❯ docker exec -it -e LANG=C.UTF8 mysql-db mysql --password=123 test_db
    mysql> \s
    --------------
    mysql  Ver 8.0.31 for Linux on aarch64 (MySQL Community Server - GPL)

    Connection id:          17
    Current database:       test_db
    Current user:           root@localhost
    SSL:                    Not in use
    Current pager:          stdout
    Using outfile:          ''
    Using delimiter:        ;
    Server version:         8.0.31 MySQL Community Server - GPL
    Protocol version:       10
    Connection:             Localhost via UNIX socket
    Server characterset:    utf8mb4
    Db     characterset:    utf8mb4
    Client characterset:    latin1
    Conn.  characterset:    latin1
    UNIX socket:            /var/run/mysqld/mysqld.sock
    Binary data as:         Hexadecimal
    Uptime:                 57 min 0 sec

    Threads: 3  Questions: 85  Slow queries: 0  Opens: 208  Flush tables: 3  Open tables: 126  Queries per second avg: 0.024
    --------------
    ``` 
    Где версия сервера определена строкой:
    ```
    Server version:         8.0.31 MySQL Community Server - GPL
    ```
4.  Получил список таблиц восстановленной БД:
    ```
    mysql> show tables;
    +-------------------+
    | Tables_in_test_db |
    +-------------------+
    | orders            |
    +-------------------+
    1 row in set (0.01 sec)
    ```
5. Выполнил запрос для получения записей из таблицы `orders` c `price > 300`:
    ```
    mysql> select * from orders where price > 300;
    +----+----------------+-------+
    | id | title          | price |
    +----+----------------+-------+
    |  2 | My little pony |   500 |
    +----+----------------+-------+
    1 row in set (0.00 sec)
    ```

## Задача 2

Создайте пользователя test в БД c паролем test-pass, используя:
- плагин авторизации mysql_native_password
- срок истечения пароля - 180 дней 
- количество попыток авторизации - 3 
- максимальное количество запросов в час - 100
- аттрибуты пользователя:
    - Фамилия "Pretty"
    - Имя "James"

Предоставьте привилегии пользователю `test` на операции SELECT базы `test_db`.
    
Используя таблицу INFORMATION_SCHEMA.USER_ATTRIBUTES получите данные по пользователю `test` и **приведите в ответе к задаче**.

### Решение

1. Создал требуемого пользователя `test` командой:
    ```
    mysql> CREATE USER 'test'@'localhost'
        ->     IDENTIFIED WITH mysql_native_password BY 'test-pass'
        ->     WITH MAX_QUERIES_PER_HOUR 100
        ->     PASSWORD EXPIRE INTERVAL 180 DAY FAILED_LOGIN_ATTEMPTS 3
        ->     ATTRIBUTE '{"Имя": "James", "Фамилия": "Pretty"}';
    Query OK, 0 rows affected (0.02 sec)
    ```
2. Предоставил привилегии пользователю `test` на операции `SELECT` базы `test_db`:
    ```
    mysql> GRANT SELECT ON test_db.* TO 'test'@'localhost';
    Query OK, 0 rows affected, 1 warning (0.02 sec)
    
    mysql> SHOW GRANTS FOR 'test'@'localhost';
    +---------------------------------------------------+
    | Grants for test@localhost                         |
    +---------------------------------------------------+
    | GRANT USAGE ON *.* TO `test`@`localhost`          |
    | GRANT SELECT ON `test_db`.* TO `test`@`localhost` |
    +---------------------------------------------------+
    2 rows in set (0.00 sec)
    ```
3. Выполнил запрос на получение данных об атрибутах пользователя `test`:
    ```
    mysql> SELECT * FROM INFORMATION_SCHEMA.USER_ATTRIBUTES
        -> WHERE USER = 'test' AND HOST = 'localhost'\G
    *************************** 1. row ***************************
         USER: test
         HOST: localhost
    ATTRIBUTE: {"Имя": "James", "Фамилия": "Pretty"}
    1 row in set (0.01 sec)
    ```
## Задача 3

Установите профилирование `SET profiling = 1`.
Изучите вывод профилирования команд `SHOW PROFILES;`.

Исследуйте, какой `engine` используется в таблице БД `test_db` и **приведите в ответе**.
#### Решение
Вывод данной команды говорит о том, что для таблицы `orders` используется InnoDB:
```
mysql> SHOW TABLE STATUS WHERE Name = "Orders"\G
*************************** 1. row ***************************
           Name: orders
         Engine: InnoDB
        Version: 10
     Row_format: Dynamic
           Rows: 5
 Avg_row_length: 3276
    Data_length: 16384
Max_data_length: 0
   Index_length: 0
      Data_free: 0
 Auto_increment: 6
    Create_time: 2022-12-08 20:20:47
    Update_time: 2022-12-08 20:20:47
     Check_time: NULL
      Collation: utf8mb4_0900_ai_ci
       Checksum: NULL
 Create_options: 
        Comment: 
1 row in set (0.01 sec)
```
Измените `engine` и **приведите время выполнения и запрос на изменения из профайлера в ответе**:
- на `MyISAM`
- на `InnoDB`

#### Решение
```
mysql> SET default_storage_engine=MyISAM;
Query OK, 0 rows affected (0.00 sec)

mysql> SET default_storage_engine=InnoDB;
Query OK, 0 rows affected (0.00 sec)

mysql> SHOW PROFILES;
+----------+------------+-----------------------------------+
| Query_ID | Duration   | Query                             |
+----------+------------+-----------------------------------+
|        1 | 0.00091450 | show prifile                      |
|        2 | 0.00160975 | select * from mysql.user          |
|        3 | 0.00133525 | SHOW ENGINES                      |
|        4 | 0.00013650 | SHOW PROFILE QUERY 3              |
|        5 | 0.00027750 | SHOW PROFILE 3                    |
|        6 | 0.00238875 | SET default_storage_engine=MyISAM |
|        7 | 0.00056375 | SET default_storage_engine=InnoDB |
+----------+------------+-----------------------------------+
7 rows in set, 1 warning (0.01 sec)

mysql> SHOW PROFILE FOR QUERY 6;
+----------------+----------+
| Status         | Duration |
+----------------+----------+
| starting       | 0.001933 |
| Opening tables | 0.000171 |
| query end      | 0.000143 |
| closing tables | 0.000016 |
| freeing items  | 0.000038 |
| cleaning up    | 0.000089 |
+----------------+----------+
6 rows in set, 1 warning (0.01 sec)

mysql> SHOW PROFILE FOR QUERY 7;
+----------------+----------+
| Status         | Duration |
+----------------+----------+
| starting       | 0.000319 |
| Opening tables | 0.000129 |
| query end      | 0.000020 |
| closing tables | 0.000008 |
| freeing items  | 0.000032 |
| cleaning up    | 0.000055 |
+----------------+----------+
6 rows in set, 1 warning (0.01 sec)
```

## Задача 4 

Изучите файл `my.cnf` в директории /etc/mysql.

Измените его согласно ТЗ (движок InnoDB):
- Скорость IO важнее сохранности данных
- Нужна компрессия таблиц для экономии места на диске
- Размер буфера с незакомиченными транзакциями 1 Мб
- Буфер кеширования 30% от ОЗУ
- Размер файла логов операций 100 Мб

Приведите в ответе измененный файл `my.cnf`.

### Решение
Файл конфигурации MySQL - my.cnf:
```ini
[mysqld]

# Измененная часть по ТЗ
innodb_doublewrite=OFF
innodb_file_format=Barracuda
innodb-file-per-table=ON
innodb_buffer_pool_size=2355M # 30% (докеру выделено 8Гб)
innodb_log_buffer_size=1M
innodb_redo_log_capacity=100M
# Конец измененной части

skip-host-cache
skip-name-resolve
datadir=/var/lib/mysql
socket=/var/run/mysqld/mysqld.sock
secure-file-priv=/var/lib/mysql-files
user=mysql

pid-file=/var/run/mysqld/mysqld.pid

[client]
socket=/var/run/mysqld/mysqld.sock

!includedir /etc/mysql/conf.d/
```
Справочно: полный размер ОЗУ в контейнере mysql:
```bash
bash-4.4# cat /proc/meminfo | head -n 1
MemTotal:        8040056 kB
```