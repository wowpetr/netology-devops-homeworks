# Домашнее задание к занятию "7.6. Написание собственных провайдеров для Terraform."

## Задача 1. 
Давайте потренируемся читать исходный код AWS провайдера, который можно склонировать от сюда: 
[https://github.com/hashicorp/terraform-provider-aws.git](https://github.com/hashicorp/terraform-provider-aws.git).
Просто найдите нужные ресурсы в исходном коде и ответы на вопросы станут понятны.  

1. Найдите, где перечислены все доступные `resource` и `data_source`, приложите ссылку на эти строки в коде на гитхабе.   
    #### Решение
    Данные строки находятся в файле `internal/provider/provider.go`:
    * [Resources](https://github.com/hashicorp/terraform-provider-aws/blob/ab126a70873964b2aa7ea3839879128ce0a2dd97/internal/provider/provider.go#L944)
    * [Data sources](https://github.com/hashicorp/terraform-provider-aws/blob/ab126a70873964b2aa7ea3839879128ce0a2dd97/internal/provider/provider.go#L419)

2. Для создания очереди сообщений SQS используется ресурс `aws_sqs_queue` у которого есть параметр `name`. 
    * С каким другим параметром конфликтует `name`? Приложите строчку кода, в которой это указано.
        #### Решение
        Параметр `name` конфликтует с `name_prefix`:
        ```
        "name": {
            Type:          schema.TypeString,
            Optional:      true,
            Computed:      true,
            ForceNew:      true,
            ConflictsWith: []string{"name_prefix"},
        },
        "name_prefix": {
            Type:          schema.TypeString,
            Optional:      true,
            Computed:      true,
            ForceNew:      true,
            ConflictsWith: []string{"name"},
        },
        ```
        [Ссылка на код](https://github.com/hashicorp/terraform-provider-aws/blob/ab126a70873964b2aa7ea3839879128ce0a2dd97/internal/service/sqs/queue.go#L88)

    * Какая максимальная длина имени? 
        #### Решение
        Максимальная длина имени: 80 символов  
        
        [Ссылка на код](https://github.com/hashicorp/terraform-provider-aws/blob/ab126a70873964b2aa7ea3839879128ce0a2dd97/internal/service/sqs/queue.go#L431)

    * Какому регулярному выражению должно подчиняться имя? 
        #### Решение
        * `^[a-zA-Z0-9_-]{1,75}\.fifo$` для fifo очередей
        * `^[a-zA-Z0-9_-]{1,80}$` во всех остальных случаях  

        [Ссылка на код](https://github.com/hashicorp/terraform-provider-aws/blob/ab126a70873964b2aa7ea3839879128ce0a2dd97/internal/service/sqs/queue.go#L431)