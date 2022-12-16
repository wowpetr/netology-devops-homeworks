# Домашнее задание к занятию "6.4. PostgreSQL"

## Задача 1

Используя docker поднимите инстанс PostgreSQL (версию 13). Данные БД сохраните в volume.

Подключитесь к БД PostgreSQL используя `psql`.

Воспользуйтесь командой `\?` для вывода подсказки по имеющимся в `psql` управляющим командам.

**Найдите и приведите** управляющие команды для:
- вывода списка БД
- подключения к БД
- вывода списка таблиц
- вывода описания содержимого таблиц
- выхода из psql

### Решение 
1. Использовал следующий docker-compose файл для запуска контейнера PostgreSQL:
    ```yml
    version: "3"

    services:

    db:
        image: postgres:13
        container_name: postgres13-db
        ports:
            - 5432:5432
        volumes:
            - ./pg_data:/var/lib/postgresql/data/pgdata
            - ./pg_backup:/backup
        environment:
            POSTGRES_PASSWORD: 123
            PGDATA: /var/lib/postgresql/data/pgdata
        restart: always
    ```
2. Подключился к контейнеру и командой:
    ```bash
    ❯ docker exec -it postgres13-db psql -U postgres
    ```
3. Команда для вывода списка БД:
    ```
    \l
    ```
4. Команда подключения к БД (пример):
    ```
    \c test_database
    ```
5. Команда для вывода списка таблиц:
    ```
    \dt
    ```
6. Команда для вывода описания содержимого таблиц (пример):
    ```
    \dS orders
    ```
7. Команда выхода из `psql`:
    ```
    \q
    ```
    Или Ctrl-D.

## Задача 2

Используя `psql` создайте БД `test_database`.

Изучите [бэкап БД](https://github.com/netology-code/virt-homeworks/tree/master/06-db-04-postgresql/test_data).

Восстановите бэкап БД в `test_database`.

Перейдите в управляющую консоль `psql` внутри контейнера.

Подключитесь к восстановленной БД и проведите операцию ANALYZE для сбора статистики по таблице.

Используя таблицу [pg_stats](https://postgrespro.ru/docs/postgresql/12/view-pg-stats), найдите столбец таблицы `orders` 
с наибольшим средним значением размера элементов в байтах.

**Приведите в ответе** команду, которую вы использовали для вычисления и полученный результат.

### Решение 
1. Создал БД в psql командой:
    ```sql
    CREATE DATABASE test_database;
    ```
2. Восстановил резервную копию командой:
    ```bash
    root@00dfa3cf427a:/# psql -U postgres test_database < /backup/test_dump.sql
    ```
3. Зашел в управляющую консоль `psql` командой:
    ```bash
    ❯ docker exec -it postgres13-db psql -U postgres test_database
    ```
4. Выполнил команду `ANALYZE orders;`
   
5. Столбцом с наибольшим средним значением таблицы `orders` в байтах является `title`, а его значением - 16 байт:
    ```
    test_database=# SELECT * FROM pg_stats WHERE tablename = 'orders' and avg_width = (SELECT MAX(avg_width) FROM pg_stats WHERE tablename = 'orders');
    -[ RECORD 1 ]----------+--------------------------------------------------------------------------------------------------------------------------------------------------
    schemaname             | public
    tablename              | orders
    attname                | title
    inherited              | f
    null_frac              | 0
    avg_width              | 16
    n_distinct             | -1
    most_common_vals       | 
    most_common_freqs      | 
    histogram_bounds       | {"Adventure psql time",Dbiezdmin,"Log gossips","Me and my bash-pet","My little database","Server gravity falls","WAL never lies","War and peace"}
    correlation            | -0.3809524
    most_common_elems      | 
    most_common_elem_freqs | 
    elem_count_histogram   |
    ```

## Задача 3

Архитектор и администратор БД выяснили, что ваша таблица orders разрослась до невиданных размеров и
поиск по ней занимает долгое время. Вам, как успешному выпускнику курсов DevOps в нетологии предложили
провести разбиение таблицы на 2 (шардировать на orders_1 - price>499 и orders_2 - price<=499).

Предложите SQL-транзакцию для проведения данной операции.

Можно ли было изначально исключить "ручное" разбиение при проектировании таблицы orders?
### Решение
```
test_database=# BEGIN;
BEGIN
test_database=*# SELECT * INTO orders_1 FROM orders WHERE price > 499;
SELECT 3
test_database=*# SELECT * INTO orders_2 FROM orders WHERE price <= 499;
SELECT 5
test_database=*# END;
COMMIT
test_database=# \dt
          List of relations
 Schema |   Name   | Type  |  Owner   
--------+----------+-------+----------
 public | orders   | table | postgres
 public | orders_1 | table | postgres
 public | orders_2 | table | postgres
(3 rows)

test_database=# SELECT * FROM orders_1;
 id |       title        | price 
----+--------------------+-------
  2 | My little database |   500
  6 | WAL never lies     |   900
  8 | Dbiezdmin          |   501
(3 rows)

test_database=# SELECT * FROM orders_2;
 id |        title         | price 
----+----------------------+-------
  1 | War and peace        |   100
  3 | Adventure psql time  |   300
  4 | Server gravity falls |   300
  5 | Log gossips          |   123
  7 | Me and my bash-pet   |   499
(5 rows)
```
Опционально, можно еще удалить еще таблицу `orders` в транзакции или после нее отдельно командой: `DROP TABLE orders;`, также можно добавить блокировку `LOCK TABLE` таблицы `orders` при необходимости в начале транзакции.
## Задача 4

Используя утилиту `pg_dump` создайте бекап БД `test_database`.


Как бы вы доработали бэкап-файл, чтобы добавить уникальность значения столбца `title` для таблиц `test_database`?

### Решение
Выполнил команду для создания резервной копии базы `test_database`:
```bash
root@00dfa3cf427a:/# pg_dump -U postgres test_database > /backup/test_database_1.sql
```
Для добавления уникальности столбца `title` я бы изменил следующий блок:
```sql
ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);
```
заменив его на:
```sql
ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id),
    ADD CONSTRAINT title_unique UNIQUE (title);
```