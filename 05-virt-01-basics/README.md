# Домашнее задание к занятию "5.1. Основы виртуализации"

## Задача 1

Вкратце опишите, как вы поняли - в чем основное отличие паравиртуализации и виртуализации на основе ОС.

### Решение
При паравиртуализации используется модифицированная гостевая операционная система, которая распознает свое виртуальное состояние и активно взаимодействует с гипервизором для организации доступа к аппаратному обеспечению. Этот подход значительно повышает производительность, однако для этого гостевые операционные системы должны быть сильно модифицированы, причем эти модификации зависят от конкретного гипервизора.  

Виртуализация на уровне операционной системы, или контейнеризация, - это другой подход к изоляции, в котором не используется гипервизор. Вместо этого используются функции ядра, которые позволяют изолировать процессы от остальной системы. Поскольку для контейнеризации не требуется виртуализация, накладные расходы на ресурсы для виртуализации на уровне операционной системы низки. Большинство реализаций предлагают почти естественную производительность. Для контейнеров есть зависимость от операционной системы. Например, можно создавать контейнеры Linux только на Unix-подобных системах и нельзя на них развернуть контейнеры Windows. 

Истинная виртуальная машина имеет ядро операционной системы, процесс инициализации, драйверы для взаимодействия с оборудованием. С другой стороны, контейнер является просто фасадом операционной системы.

## Задача 2

Выберите тип один из вариантов использования организации физических серверов, 
в зависимости от условий использования.

Организация серверов:
- физические сервера
- паравиртуализация
- виртуализация уровня ОС

Условия использования:

- Высоконагруженная база данных, чувствительная к отказу
- Различные Java-приложения
- Windows системы для использования Бухгалтерским отделом 
- Системы, выполняющие высокопроизводительные расчеты на GPU

Опишите, почему вы выбрали к каждому целевому использованию такую организацию.

### Решение
1. Высоконагруженую БД, чувствительной к отказу, мне кажется, нужно размещать в виртуальной машине так как при сбое можно легко вернуться к предыдущему состоянию из снапшотов либо быстро переместить ее на другой физический сервер. 
2. Java-приложения, видимо, лучше разместить в контейнерах для их изоляции друг от друга.
3. Windows системы для бухгалтерии наверное лучше разместить на виртуальных машинах при возможности пробросить аппаратный ключ защиты. Если же это невозможно, то придется размещать на физических серверах. Часто такие системы - это монолитные приложения с базой данных которые обслуживаются сторонними заказчиками, и поэтому проще всего дать им доступ к виртуальной машине. Здесь также система снапшотов виртуальной машины поможет быстро вернуть работоспособность при сбоях.
4. Думаю, что для высокопроизводительных расчетов на GPU нужно использовать физические серверы из-за прямого доступа к оборудованию достигается приемлемая производительность.

## Задача 3

Как вы думаете, возможно ли совмещать несколько типов виртуализации на одном сервере?
Приведите пример такого совмещения.

### Решение
Совмещать различные типы виртуализации на одном сервере возможно. Например, паравиртуализирированные виртуальные машины и контейнеры могут сосуществовать рядом. Также, внутри виртуальных машин можно развертывать контейнеры и другие виртуальные машины. Например, можно создать виртуальную машину Linux Debain на Yandex.Cloud внутри которой использовать контейнеры Docker. 