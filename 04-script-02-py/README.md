# Домашнее задание к занятию "4.2. Использование Python для решения типовых DevOps задач"

## Обязательные задания

1. Есть скрипт:
	```python
	#!/usr/bin/env python3
	a = 1
	b = '2'
	c = a + b
	```
	* Какое значение будет присвоено переменной c?
	* Как получить для переменной c значение 12?
	* Как получить для переменной c значение 3?
     
    #### Решение:
    * Никакое значение не будет присвоено переменной c, так как возникнет исключение `TypeError: unsupported operand type(s) for +: 'int' and 'str'` о несовместимости типов для операции +.
    * `c = int(str(a) + b)`
    * `c = a + int(b)`
    
2. Мы устроились на работу в компанию, где раньше уже был DevOps Engineer. Он написал скрипт, позволяющий узнать, какие файлы модифицированы в репозитории, относительно локальных изменений. Этим скриптом недовольно начальство, потому что в его выводе есть не все изменённые файлы, а также непонятен полный путь к директории, где они находятся. Как можно доработать скрипт ниже, чтобы он исполнял требования вашего руководителя?

	```python
	#!/usr/bin/env python3

	import os

	bash_command = ["cd ~/netology/sysadm-homeworks", "git status"]
	result_os = os.popen(' && '.join(bash_command)).read()
	is_change = False
	for result in result_os.split('\n'):
		if result.find('modified') != -1:
			prepare_result = result.replace('\tmodified:   ', '')
			print(prepare_result)
			break

	```
     
    #### Решение:  
	```python
	#!/usr/bin/env python3

	import os
	
	rep_path = "~/netology-devops-homeworks"
	bash_command = [f"cd {rep_path}", "git status --porcelain"]
	result_os = os.popen(' && '.join(bash_command)).read()
	for result in result_os.split('\n'):
		if len(result) > 3 and result[1] == "M":
			filename = result[3:]
			prepare_result = os.path.abspath(filename)
			print(prepare_result)
    ``` 
    В скрипт добавлен ключ `--porcelain` команды `git status`, для облегчения разбора и для вывода полных путей файлов. Также удалена лишняя команда `break` из-за которой выводился только лишь один файл. Для получения абсолютных путей использовалась функция `os.path.abspath`.


3. Доработать скрипт выше так, чтобы он мог проверять не только локальный репозиторий в текущей директории, а также умел воспринимать путь к репозиторию, который мы передаём как входной параметр. Мы точно знаем, что начальство коварное и будет проверять работу этого скрипта в директориях, которые не являются локальными репозиториями.
    #### Решение:  
	```python
	#!/usr/bin/env python3

	import sys
	import os
	
	if len(sys.argv) > 1:
		repo_path = sys.argv[1]
	else:
		print("git repo path is not specified!")
		exit(1)
	
	if not os.path.exists(repo_path):
		print(f"git repo path {repo_path} does not exist!")
		exit(2)
	
	if not os.path.exists(repo_path + "/.git"):
		print(f"git repo path {repo_path} does not have a .git dirrectory!")
		exit(3)
	
	bash_command = [f"cd {repo_path}", "git status --porcelain"]
	result_os = os.popen(' && '.join(bash_command)).read()
	for result in result_os.split('\n'):
		if len(result) > 3 and result[1] == "M":
			filename = result[3:]
			prepare_result = os.path.abspath(filename)
			print(prepare_result)
    ```
    В скрипт импортирован модуль `sys` для разбора входных аргументов `sys.argv` и некоторый контроль задания правильного пути в первом аргументе при запуске скрипта. 
4. Наша команда разрабатывает несколько веб-сервисов, доступных по http. Мы точно знаем, что на их стенде нет никакой балансировки, кластеризации, за DNS прячется конкретный IP сервера, где установлен сервис. Проблема в том, что отдел, занимающийся нашей инфраструктурой очень часто меняет нам сервера, поэтому IP меняются примерно раз в неделю, при этом сервисы сохраняют за собой DNS имена. Это бы совсем никого не беспокоило, если бы несколько раз сервера не уезжали в такой сегмент сети нашей компании, который недоступен для разработчиков. Мы хотим написать скрипт, который опрашивает веб-сервисы, получает их IP, выводит информацию в стандартный вывод в виде: <URL сервиса> - <его IP>. Также, должна быть реализована возможность проверки текущего IP сервиса c его IP из предыдущей проверки. Если проверка будет провалена - оповестить об этом в стандартный вывод сообщением: [ERROR] <URL сервиса> IP mismatch: <старый IP> <Новый IP>. Будем считать, что наша разработка реализовала сервисы: drive.google.com, mail.google.com, google.com.
    #### Решение:  
	```python
	#!/usr/bin/env python3
	
	import socket
	import time
	
	services = ('drive.google.com', 'mail.google.com', 'google.com')
	results = dict()
	
	while True:
	
		for fqdn in services:
			ip = socket.gethostbyname(fqdn)
			prev_ip = results.get(fqdn)
			if prev_ip != ip and prev_ip is not None:
				print(f"[ERROR] {fqdn} IP mismatch: {prev_ip} {ip}")
			results[fqdn] = ip
	
		#print(results)
		time.sleep(5)
    ```