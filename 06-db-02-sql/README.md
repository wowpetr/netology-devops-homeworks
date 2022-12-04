# Домашнее задание к занятию "6.2. SQL"

## Задача 1

Используя docker поднимите инстанс PostgreSQL (версию 12) c 2 volume, 
в который будут складываться данные БД и бэкапы.

Приведите получившуюся команду или docker-compose манифест.

### Решение

```yaml
version: "3"

services:

  db:
    image: postgres:12
    container_name: postgres-db
    ports:
      - 5432:5432
    volumes:
      - ./pg_data:/var/lib/postgresql/data/pgdata
      - ./pg_backup:/backup
    environment:
      POSTGRES_PASSWORD: 123
      POSTGRES_DB: netology_db
      PGDATA: /var/lib/postgresql/data/pgdata
    networks:
      netology-network:
        ipv4_address: 172.22.0.10
    restart: always

networks:
  netology-network:
    driver: bridge
    ipam:
      config:
        - subnet: 172.22.0.0/24
```

## Задача 2

В БД из задачи 1: 
- создайте пользователя test-admin-user и БД test_db
- в БД test_db создайте таблицу orders и clients (спeцификация таблиц ниже)
- предоставьте привилегии на все операции пользователю test-admin-user на таблицы БД test_db
- создайте пользователя test-simple-user  
- предоставьте пользователю test-simple-user права на SELECT/INSERT/UPDATE/DELETE данных таблиц БД test_db

Таблица orders:
- id (serial primary key)
- наименование (string)
- цена (integer)

Таблица clients:
- id (serial primary key)
- фамилия (string)
- страна проживания (string, index)
- заказ (foreign key orders)

### Решение
Выполнил следующие запросы:
```sql
-- Создание базы данных
CREATE DATABASE test_db;

-- Создание пользователя test-admin-user
CREATE USER "test-admin-user" WITH PASSWORD '123';
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO "test-admin-user";

-- Создание пользователя test-simple-user
CREATE USER "test-simple-user";
GRANT SELECT,INSERT,UPDATE,DELETE ON ALL TABLES IN SCHEMA public TO "test-simple-user";

-- Создание таблицы orders
CREATE TABLE orders (
  id serial primary key,
  "наименование" varchar(100),
  "цена" int
);

-- Создание таблицы clients
CREATE TABLE clients (
  id serial primary key,
  "фамилия" varchar(50),
  "страна проживания" varchar(60),
  "заказ" int references orders(id) 
);
```
Приведите:
- итоговый список БД после выполнения пунктов выше,

    #### Решение
    ```
    test_db-# \l
                                    List of databases
        Name     |  Owner   | Encoding |  Collate   |   Ctype    |   Access privileges
    -------------+----------+----------+------------+------------+-----------------------
    netology_db | postgres | UTF8     | en_US.utf8 | en_US.utf8 |
    postgres    | postgres | UTF8     | en_US.utf8 | en_US.utf8 |
    template0   | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
                |          |          |            |            | postgres=CTc/postgres
    template1   | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +
                |          |          |            |            | postgres=CTc/postgres
    test_db     | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =Tc/postgres         +
                |          |          |            |            | postgres=CTc/postgres
    (5 rows)
    ```

- описание таблиц (describe)
    #### Решение
    ```
  test_db=# \d orders
                                        Table "public.orders"
      Column    |          Type          | Collation | Nullable |              Default
  --------------+------------------------+-----------+----------+------------------------------------
  id           | integer                |           | not null | nextval('orders_id_seq'::regclass)
  наименование | character varying(100) |           |          |
  цена         | integer                |           |          |
  Indexes:
      "orders_pkey" PRIMARY KEY, btree (id)
  Referenced by:
      TABLE "clients" CONSTRAINT "clients_заказ_fkey" FOREIGN KEY ("заказ") REFERENCES orders(id)

  test_db=# \d clients
                                          Table "public.clients"
        Column       |         Type          | Collation | Nullable |               Default
  -------------------+-----------------------+-----------+----------+-------------------------------------
  id                | integer               |           | not null | nextval('clients_id_seq'::regclass)
  фамилия           | character varying(50) |           |          |
  страна проживания | character varying(60) |           |          |
  заказ             | integer               |           |          |
  Indexes:
      "clients_pkey" PRIMARY KEY, btree (id)
  Foreign-key constraints:
      "clients_заказ_fkey" FOREIGN KEY ("заказ") REFERENCES orders(id)
    ```

- SQL-запрос для выдачи списка пользователей с правами над таблицами test_db
    #### Решение
    ```
    test_db=# SELECT * FROM information_schema.role_table_grants WHERE table_name='clients';
    grantor  |     grantee      | table_catalog | table_schema | table_name | privilege_type | is_grantable | with_hierarchy
    ----------+------------------+---------------+--------------+------------+----------------+--------------+----------------
    postgres | postgres         | test_db       | public       | clients    | INSERT         | YES          | NO
    postgres | postgres         | test_db       | public       | clients    | SELECT         | YES          | YES
    postgres | postgres         | test_db       | public       | clients    | UPDATE         | YES          | NO
    postgres | postgres         | test_db       | public       | clients    | DELETE         | YES          | NO
    postgres | postgres         | test_db       | public       | clients    | TRUNCATE       | YES          | NO
    postgres | postgres         | test_db       | public       | clients    | REFERENCES     | YES          | NO
    postgres | postgres         | test_db       | public       | clients    | TRIGGER        | YES          | NO
    postgres | test-admin-user  | test_db       | public       | clients    | INSERT         | NO           | NO
    postgres | test-admin-user  | test_db       | public       | clients    | SELECT         | NO           | YES
    postgres | test-admin-user  | test_db       | public       | clients    | UPDATE         | NO           | NO
    postgres | test-admin-user  | test_db       | public       | clients    | DELETE         | NO           | NO
    postgres | test-admin-user  | test_db       | public       | clients    | TRUNCATE       | NO           | NO
    postgres | test-admin-user  | test_db       | public       | clients    | REFERENCES     | NO           | NO
    postgres | test-admin-user  | test_db       | public       | clients    | TRIGGER        | NO           | NO
    postgres | test-simple-user | test_db       | public       | clients    | INSERT         | NO           | NO
    postgres | test-simple-user | test_db       | public       | clients    | SELECT         | NO           | YES
    postgres | test-simple-user | test_db       | public       | clients    | UPDATE         | NO           | NO
    postgres | test-simple-user | test_db       | public       | clients    | DELETE         | NO           | NO
    (18 rows)

    test_db=# SELECT * FROM information_schema.role_table_grants WHERE table_name='orders';
    grantor  |     grantee      | table_catalog | table_schema | table_name | privilege_type | is_grantable | with_hierarchy
    ----------+------------------+---------------+--------------+------------+----------------+--------------+----------------
    postgres | postgres         | test_db       | public       | orders     | INSERT         | YES          | NO
    postgres | postgres         | test_db       | public       | orders     | SELECT         | YES          | YES
    postgres | postgres         | test_db       | public       | orders     | UPDATE         | YES          | NO
    postgres | postgres         | test_db       | public       | orders     | DELETE         | YES          | NO
    postgres | postgres         | test_db       | public       | orders     | TRUNCATE       | YES          | NO
    postgres | postgres         | test_db       | public       | orders     | REFERENCES     | YES          | NO
    postgres | postgres         | test_db       | public       | orders     | TRIGGER        | YES          | NO
    postgres | test-admin-user  | test_db       | public       | orders     | INSERT         | NO           | NO
    postgres | test-admin-user  | test_db       | public       | orders     | SELECT         | NO           | YES
    postgres | test-admin-user  | test_db       | public       | orders     | UPDATE         | NO           | NO
    postgres | test-admin-user  | test_db       | public       | orders     | DELETE         | NO           | NO
    postgres | test-admin-user  | test_db       | public       | orders     | TRUNCATE       | NO           | NO
    postgres | test-admin-user  | test_db       | public       | orders     | REFERENCES     | NO           | NO
    postgres | test-admin-user  | test_db       | public       | orders     | TRIGGER        | NO           | NO
    postgres | test-simple-user | test_db       | public       | orders     | INSERT         | NO           | NO
    postgres | test-simple-user | test_db       | public       | orders     | SELECT         | NO           | YES
    postgres | test-simple-user | test_db       | public       | orders     | UPDATE         | NO           | NO
    postgres | test-simple-user | test_db       | public       | orders     | DELETE         | NO           | NO
    (18 rows)
    ```
- список пользователей с правами над таблицами test_db
    #### Решение
    ```
    test_db=# \dp
                                                Access privileges
    Schema  |       Name        |   Type   |         Access privileges          | Column privileges | Policies
    --------+-------------------+----------+------------------------------------+-------------------+----------
    public  | clients           | table    | postgres=arwdDxt/postgres         +|                   |
            |                   |          | "test-admin-user"=arwdDxt/postgres+|                   |
            |                   |          | "test-simple-user"=arwd/postgres   |                   |
    public  | clients_id_seq    | sequence |                                    |                   |
    public  | clients_заказ_seq | sequence |                                    |                   |
    public  | orders            | table    | postgres=arwdDxt/postgres         +|                   |
            |                   |          | "test-admin-user"=arwdDxt/postgres+|                   |
            |                   |          | "test-simple-user"=arwd/postgres   |                   |
    public  | orders_id_seq     | sequence |                                    |                   |
    (5 rows)
    ```
  



## Задача 3

Используя SQL синтаксис - наполните таблицы следующими тестовыми данными:

Таблица orders

|Наименование|цена|
|------------|----|
|Шоколад| 10 |
|Принтер| 3000 |
|Книга| 500 |
|Монитор| 7000|
|Гитара| 4000|

Таблица clients

|ФИО|Страна проживания|
|------------|----|
|Иванов Иван Иванович| USA |
|Петров Петр Петрович| Canada |
|Иоганн Себастьян Бах| Japan |
|Ронни Джеймс Дио| Russia|
|Ritchie Blackmore| Russia|

Используя SQL синтаксис:
- вычислите количество записей для каждой таблицы 
- приведите в ответе:
    - запросы 
    - результаты их выполнения.

### Решение

1. Выполнил запрос на заполнение таблицы `orders`:
```sql
INSERT INTO orders ("наименование", "цена") VALUES
  ('Шоколад', 10),
  ('Принтер', 3000),
  ('Книга', 500),
  ('Монитор', 7000),
  ('Гитара', 4000);
```
2. Запрос на заполнение таблицы `clients`:
```sql
INSERT INTO clients ("фамилия", "страна проживания") VALUES
  ('Иванов Иван Иванович', 'USA'),
  ('Петров Петр Петрович', 'Canada'),
  ('Иоганн Себастьян Бах', 'Japan'),
  ('Ронни Джеймс Дио', 'Russia'),
  ('Ritchie Blackmore', 'Russia')
```
3. Выполнил запросы получения количества записей таблиц orders и clients:
```
test_db=# SELECT COUNT(*) FROM clients;
 count
-------
     5
(1 row)

test_db=# SELECT COUNT(*) FROM orders;
 count
-------
     5
(1 row)
```

## Задача 4

Часть пользователей из таблицы clients решили оформить заказы из таблицы orders.

Используя foreign keys свяжите записи из таблиц, согласно таблице:

|ФИО|Заказ|
|------------|----|
|Иванов Иван Иванович| Книга |
|Петров Петр Петрович| Монитор |
|Иоганн Себастьян Бах| Гитара |

Приведите SQL-запросы для выполнения данных операций.
### Решение
```sql
test_db=# UPDATE clients SET "заказ" = 3 WHERE id = 1;
UPDATE 1
test_db=# UPDATE clients SET "заказ" = 4 WHERE id = 2;
UPDATE 1
test_db=# UPDATE clients SET "заказ" = 5 WHERE id = 3;
UPDATE 1
test_db=# select * from orders;
 id | наименование | цена
----+--------------+------
  1 | Шоколад      |   10
  2 | Принтер      | 3000
  3 | Книга        |  500
  4 | Монитор      | 7000
  5 | Гитара       | 4000
(5 rows)
test_db=# select * from clients;
 id |       фамилия        | страна проживания | заказ
----+----------------------+-------------------+-------
  4 | Ронни Джеймс Дио     | Russia            |
  5 | Ritchie Blackmore    | Russia            |
  1 | Иванов Иван Иванович | USA               |     3
  2 | Петров Петр Петрович | Canada            |     4
  3 | Иоганн Себастьян Бах | Japan             |     5
(5 rows)
```

Приведите SQL-запрос для выдачи всех пользователей, которые совершили заказ, а также вывод данного запроса.
### Решение
Способ 1 (только имена пользователей):
```
test_db=# SELECT фамилия FROM clients WHERE заказ is not NULL;
       фамилия
----------------------
 Иванов Иван Иванович
 Петров Петр Петрович
 Иоганн Себастьян Бах
(3 rows)
```
Способ 2 (имена пользователей с их заказами):
```
test_db=# SELECT c.фамилия,o.наименование FROM clients c INNER JOIN orders o ON c.заказ = o.id;
       фамилия        | наименование
----------------------+--------------
 Иванов Иван Иванович | Книга
 Петров Петр Петрович | Монитор
 Иоганн Себастьян Бах | Гитара
(3 rows)
```

## Задача 5

Получите полную информацию по выполнению запроса выдачи всех пользователей из задачи 4 
(используя директиву EXPLAIN).

Приведите получившийся результат и объясните что значат полученные значения.

### Решение
```
test_db=# EXPLAIN SELECT фамилия FROM clients WHERE заказ is not NULL;
                         QUERY PLAN
------------------------------------------------------------
 Seq Scan on clients  (cost=0.00..12.70 rows=269 width=118)
   Filter: ("заказ" IS NOT NULL)
(2 rows)

test_db=# EXPLAIN SELECT c.фамилия,o.наименование FROM clients c INNER JOIN orders o ON c.заказ = o.id;
                               QUERY PLAN
-------------------------------------------------------------------------
 Hash Join  (cost=17.20..30.62 rows=270 width=336)
   Hash Cond: (c."заказ" = o.id)
   ->  Seq Scan on clients c  (cost=0.00..12.70 rows=270 width=122)
   ->  Hash  (cost=13.20..13.20 rows=320 width=222)
         ->  Seq Scan on orders o  (cost=0.00..13.20 rows=320 width=222)
(5 rows)
```
В первом запросе мы видим последовательный скан поля `заказ` для проверки на NULL поскольку, во-первых, у нас нет индекса, а во-вторых, на данной таблице индекс будет не эффективен (я пробовал создать индекс на поле `заказ` и оптимизатор его не использует поскольку количество возвращаемых строк выше чем 10-15% общего количества строк).

Во втором запросе план запроса следующий: выполняется Hash Join который загружает подходящие записи с одной стороны объединения в хэш таблицу, которые в свою очередь проверяются с каждой записью с другой стороны объединения. Индексный поиск также, видимо, не используется, из-за размеров таблицы.

## Задача 6

Создайте бэкап БД test_db и поместите его в volume, предназначенный для бэкапов (см. Задачу 1).

Остановите контейнер с PostgreSQL (но не удаляйте volumes).

Поднимите новый пустой контейнер с PostgreSQL.

Восстановите БД test_db в новом контейнере.

Приведите список операций, который вы применяли для бэкапа данных и восстановления. 

### Решение
1. Подключился к докер контейнеру:
```bash
❯ docker exec -it postgres-db bash
```
2. Выполнил команду на создание резервной копии базы test_db:
```bash
root@1e7024037f6d:/# pg_dump -U postgres test_db > /backup/test_db.sql
```
3. Создал новую директорию для нового docker-compose, в которой создал  директорию `pg_data` и скопировал `pg_backup` из первого стека контейнеров.
4. Изменил docker-compose.yml с новыми сетевыми настройками и именем контейнера:
    ```yaml
    version: "3"

    services:

      db:
        image: postgres:12
        container_name: postgres-db-2
        ports:
          - 5433:5432
        volumes:
          - ./pg_data:/var/lib/postgresql/data/pgdata
          - ./pg_backup:/backup
        environment:
          POSTGRES_PASSWORD: 123
          POSTGRES_DB: test_db
          PGDATA: /var/lib/postgresql/data/pgdata
        networks:
          netology-network:
            ipv4_address: 172.23.0.10
        restart: always

    networks:
      netology-network:
        driver: bridge
        ipam:
          config:
            - subnet: 172.23.0.0/24
    ```
5. Запустил новый стек командой: `docker-compose up -d`.
6. Подключился в новый контейнер командой: `docker exec -it bash`.
7. Создал пользователей: `test-admin-user` и `test-simple-user`.
8. Восстановил базу командой:
9.    ```bash
      root@dab29f40fcd7:/# psql -h localhost -U posgtres test_db < /backup/test_db.sql
      ```