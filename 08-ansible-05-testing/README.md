# Домашнее задание к занятию "5. Тестирование roles"

## Подготовка к выполнению
1. Установите molecule: `pip3 install "molecule==3.5.2"`
2. Выполните `docker pull aragast/netology:latest` -  это образ с podman, tox и несколькими пайтонами (3.7 и 3.9) внутри

## Основная часть

Наша основная цель - настроить тестирование наших ролей. Задача: сделать сценарии тестирования для vector. Ожидаемый результат: все сценарии успешно проходят тестирование ролей.

### Molecule

1. Запустите  `molecule test -s centos_7` внутри корневой директории clickhouse-role, посмотрите на вывод команды.
    #### Решение
    ```bash
    ~/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse ❯ molecule test -s centos_7
    INFO     centos_7 scenario test matrix: dependency, lint, cleanup, destroy, syntax, create, prepare, converge, idempotence, side_effect, verify, cleanup, destroy
    INFO     Performing prerun...
    INFO     Set ANSIBLE_LIBRARY=/Users/plosev/.cache/ansible-compat/7e099f/modules:/Users/plosev/.ansible/plugins/modules:/usr/share/ansible/plugins/modules
    INFO     Set ANSIBLE_COLLECTIONS_PATH=/Users/plosev/.cache/ansible-compat/7e099f/collections:/Users/plosev/.ansible/collections:/usr/share/ansible/collections
    INFO     Set ANSIBLE_ROLES_PATH=/Users/plosev/.cache/ansible-compat/7e099f/roles:/Users/plosev/.ansible/roles:/usr/share/ansible/roles:/etc/ansible/roles
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/hosts.yml linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/hosts
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/group_vars/ linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/group_vars
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/host_vars/ linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/host_vars
    INFO     Running centos_7 > dependency
    WARNING  Skipping, missing the requirements file.
    WARNING  Skipping, missing the requirements file.
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/hosts.yml linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/hosts
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/group_vars/ linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/group_vars
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/host_vars/ linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/host_vars
    INFO     Running centos_7 > lint
    COMMAND: yamllint .
    ansible-lint
    flake8


    Passed with production profile: 0 failure(s), 0 warning(s) on 0 files.
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/hosts.yml linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/hosts
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/group_vars/ linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/group_vars
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/host_vars/ linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/host_vars
    INFO     Running centos_7 > cleanup
    WARNING  Skipping, cleanup playbook not configured.
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/hosts.yml linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/hosts
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/group_vars/ linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/group_vars
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/host_vars/ linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/host_vars
    INFO     Running centos_7 > destroy
    INFO     Sanity checks: 'docker'

    PLAY [Destroy] *****************************************************************

    TASK [Set async_dir for HOME env] **********************************************
    ok: [localhost]

    TASK [Destroy molecule instance(s)] ********************************************
    changed: [localhost] => (item=centos_7)

    TASK [Wait for instance(s) deletion to complete] *******************************
    FAILED - RETRYING: [localhost]: Wait for instance(s) deletion to complete (300 retries left).
    ok: [localhost] => (item=centos_7)

    TASK [Delete docker networks(s)] ***********************************************
    skipping: [localhost]

    PLAY RECAP *********************************************************************
    localhost                  : ok=3    changed=1    unreachable=0    failed=0    skipped=1    rescued=0    ignored=0

    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/hosts.yml linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/hosts
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/group_vars/ linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/group_vars
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/host_vars/ linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/host_vars
    INFO     Running centos_7 > syntax

    playbook: /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/resources/playbooks/converge.yml
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/hosts.yml linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/hosts
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/group_vars/ linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/group_vars
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/host_vars/ linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/host_vars
    INFO     Running centos_7 > create

    PLAY [Create] ******************************************************************

    TASK [Set async_dir for HOME env] **********************************************
    ok: [localhost]

    TASK [Log into a Docker registry] **********************************************
    skipping: [localhost] => (item=None) 
    skipping: [localhost]

    TASK [Check presence of custom Dockerfiles] ************************************
    ok: [localhost] => (item={'capabilities': ['SYS_ADMIN'], 'command': '/usr/sbin/init', 'dockerfile': '../resources/Dockerfile.j2', 'env': {'ANSIBLE_USER': 'ansible', 'DEPLOY_GROUP': 'deployer', 'SUDO_GROUP': 'wheel', 'container': 'docker'}, 'image': 'centos:7', 'name': 'centos_7', 'privileged': True, 'tmpfs': ['/run', '/tmp'], 'volumes': ['/sys/fs/cgroup:/sys/fs/cgroup']})

    TASK [Create Dockerfiles from image names] *************************************
    changed: [localhost] => (item={'capabilities': ['SYS_ADMIN'], 'command': '/usr/sbin/init', 'dockerfile': '../resources/Dockerfile.j2', 'env': {'ANSIBLE_USER': 'ansible', 'DEPLOY_GROUP': 'deployer', 'SUDO_GROUP': 'wheel', 'container': 'docker'}, 'image': 'centos:7', 'name': 'centos_7', 'privileged': True, 'tmpfs': ['/run', '/tmp'], 'volumes': ['/sys/fs/cgroup:/sys/fs/cgroup']})

    TASK [Synchronization the context] *********************************************
    changed: [localhost] => (item={'capabilities': ['SYS_ADMIN'], 'command': '/usr/sbin/init', 'dockerfile': '../resources/Dockerfile.j2', 'env': {'ANSIBLE_USER': 'ansible', 'DEPLOY_GROUP': 'deployer', 'SUDO_GROUP': 'wheel', 'container': 'docker'}, 'image': 'centos:7', 'name': 'centos_7', 'privileged': True, 'tmpfs': ['/run', '/tmp'], 'volumes': ['/sys/fs/cgroup:/sys/fs/cgroup']})

    TASK [Discover local Docker images] ********************************************
    ok: [localhost] => (item=None)
    ok: [localhost]

    TASK [Build an Ansible compatible image (new)] *********************************
    ok: [localhost] => (item=molecule_local/centos:7)

    TASK [Create docker network(s)] ************************************************
    skipping: [localhost]

    TASK [Determine the CMD directives] ********************************************
    ok: [localhost] => (item=None)
    ok: [localhost]

    TASK [Create molecule instance(s)] *********************************************
    changed: [localhost] => (item=centos_7)

    TASK [Wait for instance(s) creation to complete] *******************************
    FAILED - RETRYING: [localhost]: Wait for instance(s) creation to complete (300 retries left).
    changed: [localhost] => (item=None)
    changed: [localhost]

    PLAY RECAP *********************************************************************
    localhost                  : ok=9    changed=4    unreachable=0    failed=0    skipped=2    rescued=0    ignored=0

    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/hosts.yml linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/hosts
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/group_vars/ linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/group_vars
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/host_vars/ linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/host_vars
    INFO     Running centos_7 > prepare
    WARNING  Skipping, prepare playbook not configured.
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/hosts.yml linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/hosts
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/group_vars/ linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/group_vars
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/host_vars/ linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/host_vars
    INFO     Running centos_7 > converge

    PLAY [Converge] ****************************************************************

    TASK [Gathering Facts] *********************************************************
    ok: [centos_7]

    TASK [Apply Clickhouse Role] ***************************************************

    TASK [clickhouse : Include OS Family Specific Variables] ***********************
    ok: [centos_7]

    TASK [clickhouse : include_tasks] **********************************************
    included: /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/tasks/precheck.yml for centos_7

    TASK [clickhouse : Requirements check | Checking sse4_2 support] ***************
    fatal: [centos_7]: FAILED! => {"changed": false, "cmd": ["grep", "-q", "sse4_2", "/proc/cpuinfo"], "delta": "0:00:00.234697", "end": "2023-02-22 13:48:17.713960", "msg": "non-zero return code", "rc": 1, "start": "2023-02-22 13:48:17.479263", "stderr": "", "stderr_lines": [], "stdout": "", "stdout_lines": []}

    PLAY RECAP *********************************************************************
    centos_7                   : ok=3    changed=0    unreachable=0    failed=1    skipped=0    rescued=0    ignored=0

    CRITICAL Ansible return code was 2, command was: ['ansible-playbook', '-D', '--inventory', '/Users/plosev/.cache/molecule/clickhouse/centos_7/inventory', '--skip-tags', 'molecule-notest,notest', '/Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/resources/playbooks/converge.yml']
    WARNING  An error occurred during the test sequence action: 'converge'. Cleaning up.
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/hosts.yml linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/hosts
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/group_vars/ linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/group_vars
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/host_vars/ linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/host_vars
    INFO     Running centos_7 > cleanup
    WARNING  Skipping, cleanup playbook not configured.
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/hosts.yml linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/hosts
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/group_vars/ linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/group_vars
    INFO     Inventory /Users/plosev/netology-devops-homeworks/08-ansible-04-role/playbook/roles/clickhouse/molecule/centos_7/../resources/inventory/host_vars/ linked to /Users/plosev/.cache/molecule/clickhouse/centos_7/inventory/host_vars
    INFO     Running centos_7 > destroy

    PLAY [Destroy] *****************************************************************

    TASK [Set async_dir for HOME env] **********************************************
    ok: [localhost]

    TASK [Destroy molecule instance(s)] ********************************************
    changed: [localhost] => (item=centos_7)

    TASK [Wait for instance(s) deletion to complete] *******************************
    FAILED - RETRYING: [localhost]: Wait for instance(s) deletion to complete (300 retries left).
    changed: [localhost] => (item=centos_7)

    TASK [Delete docker networks(s)] ***********************************************
    skipping: [localhost]

    PLAY RECAP *********************************************************************
    localhost                  : ok=3    changed=2    unreachable=0    failed=0    skipped=1    rescued=0    ignored=0
    ```
2. Перейдите в каталог с ролью vector-role и создайте сценарий тестирования по умолчанию при помощи `molecule init scenario --driver-name docker`.
    #### Решение
    ```bash
    ~/temp-netology-homeworks/08-ansible-04-role/vector-role ❯ molecule init scenario --driver-name docker
    INFO     Initializing new scenario default...
    INFO     Initialized scenario in /Users/plosev/temp-netology-homeworks/08-ansible-04-role/vector-role/molecule/default successfully.
    ```
3. Добавьте несколько разных дистрибутивов (centos:8, ubuntu:latest) для инстансов и протестируйте роль, исправьте найденные ошибки, если они есть.
    #### Решение
    Добавил `centos_7` и `ubuntu`:
    ```yaml
    platforms:
      - name: centos_7
        image: docker.io/pycontribs/centos:7
        pre_build_image: true
      - name: centos_8
        image: docker.io/pycontribs/centos:8
        pre_build_image: true
      - name: ubuntu
        image: docker.io/pycontribs/ubuntu:latest
        pre_build_image: true
    ```
    Реализовал запуск сервисов вручную для контейнеров:
    ```yaml
    ---
    - name: Start systemd vector service
    become: true
    ansible.builtin.service:
      name: vector
      state: restarted
    listen: "Start vector service"
    when: not ansible_module_running_in_container

    - name: Start vector service manually
    become: true
    ansible.builtin.shell: "vector --config /etc/vector/vector.toml >/dev/null 2>&1 &"
    listen: "Start vector service"
    when: ansible_module_running_in_container
    ```
    Результаты теста
    ```bash
    ~/temp-netology-homeworks/08-ansible-04-role/vector-role ❯ molecule test
    INFO     default scenario test matrix: dependency, lint, cleanup, destroy, syntax, create, prepare, converge, idempotence, side_effect, verify, cleanup, destroy
    INFO     Performing prerun...
    INFO     Set ANSIBLE_LIBRARY=/Users/plosev/.cache/ansible-compat/f5bcd7/modules:/Users/plosev/.ansible/plugins/modules:/usr/share/ansible/plugins/modules
    INFO     Set ANSIBLE_COLLECTIONS_PATH=/Users/plosev/.cache/ansible-compat/f5bcd7/collections:/Users/plosev/.ansible/collections:/usr/share/ansible/collections
    INFO     Set ANSIBLE_ROLES_PATH=/Users/plosev/.cache/ansible-compat/f5bcd7/roles:/Users/plosev/.ansible/roles:/usr/share/ansible/roles:/etc/ansible/roles
    INFO     Running default > dependency
    WARNING  Skipping, missing the requirements file.
    WARNING  Skipping, missing the requirements file.
    INFO     Running default > lint
    INFO     Lint is disabled.
    INFO     Running default > cleanup
    WARNING  Skipping, cleanup playbook not configured.
    INFO     Running default > destroy
    INFO     Sanity checks: 'docker'

    PLAY [Destroy] *****************************************************************

    TASK [Set async_dir for HOME env] **********************************************
    ok: [localhost]

    TASK [Destroy molecule instance(s)] ********************************************
    changed: [localhost] => (item=centos_7)
    changed: [localhost] => (item=centos_8)
    changed: [localhost] => (item=ubuntu)

    TASK [Wait for instance(s) deletion to complete] *******************************
    ok: [localhost] => (item=centos_7)
    ok: [localhost] => (item=centos_8)
    ok: [localhost] => (item=ubuntu)

    TASK [Delete docker networks(s)] ***********************************************
    skipping: [localhost]

    PLAY RECAP *********************************************************************
    localhost                  : ok=3    changed=1    unreachable=0    failed=0    skipped=1    rescued=0    ignored=0

    INFO     Running default > syntax

    playbook: /Users/plosev/temp-netology-homeworks/08-ansible-04-role/vector-role/molecule/default/converge.yml
    INFO     Running default > create

    PLAY [Create] ******************************************************************

    TASK [Set async_dir for HOME env] **********************************************
    ok: [localhost]

    TASK [Log into a Docker registry] **********************************************
    skipping: [localhost] => (item=None) 
    skipping: [localhost] => (item=None) 
    skipping: [localhost] => (item=None) 
    skipping: [localhost]

    TASK [Check presence of custom Dockerfiles] ************************************
    ok: [localhost] => (item={'image': 'docker.io/pycontribs/centos:7', 'name': 'centos_7', 'pre_build_image': True})
    ok: [localhost] => (item={'image': 'docker.io/pycontribs/centos:8', 'name': 'centos_8', 'pre_build_image': True})
    ok: [localhost] => (item={'image': 'docker.io/pycontribs/ubuntu:latest', 'name': 'ubuntu', 'pre_build_image': True})

    TASK [Create Dockerfiles from image names] *************************************
    skipping: [localhost] => (item={'image': 'docker.io/pycontribs/centos:7', 'name': 'centos_7', 'pre_build_image': True}) 
    skipping: [localhost] => (item={'image': 'docker.io/pycontribs/centos:8', 'name': 'centos_8', 'pre_build_image': True}) 
    skipping: [localhost] => (item={'image': 'docker.io/pycontribs/ubuntu:latest', 'name': 'ubuntu', 'pre_build_image': True}) 
    skipping: [localhost]

    TASK [Synchronization the context] *********************************************
    skipping: [localhost] => (item={'image': 'docker.io/pycontribs/centos:7', 'name': 'centos_7', 'pre_build_image': True}) 
    skipping: [localhost] => (item={'image': 'docker.io/pycontribs/centos:8', 'name': 'centos_8', 'pre_build_image': True}) 
    skipping: [localhost] => (item={'image': 'docker.io/pycontribs/ubuntu:latest', 'name': 'ubuntu', 'pre_build_image': True}) 
    skipping: [localhost]

    TASK [Discover local Docker images] ********************************************
    ok: [localhost] => (item={'changed': False, 'skipped': True, 'skip_reason': 'Conditional result was False', 'item': {'image': 'docker.io/pycontribs/centos:7', 'name': 'centos_7', 'pre_build_image': True}, 'ansible_loop_var': 'item', 'i': 0, 'ansible_index_var': 'i'})
    ok: [localhost] => (item={'changed': False, 'skipped': True, 'skip_reason': 'Conditional result was False', 'item': {'image': 'docker.io/pycontribs/centos:8', 'name': 'centos_8', 'pre_build_image': True}, 'ansible_loop_var': 'item', 'i': 1, 'ansible_index_var': 'i'})
    ok: [localhost] => (item={'changed': False, 'skipped': True, 'skip_reason': 'Conditional result was False', 'item': {'image': 'docker.io/pycontribs/ubuntu:latest', 'name': 'ubuntu', 'pre_build_image': True}, 'ansible_loop_var': 'item', 'i': 2, 'ansible_index_var': 'i'})

    TASK [Build an Ansible compatible image (new)] *********************************
    skipping: [localhost] => (item=molecule_local/docker.io/pycontribs/centos:7) 
    skipping: [localhost] => (item=molecule_local/docker.io/pycontribs/centos:8) 
    skipping: [localhost] => (item=molecule_local/docker.io/pycontribs/ubuntu:latest) 
    skipping: [localhost]

    TASK [Create docker network(s)] ************************************************
    skipping: [localhost]

    TASK [Determine the CMD directives] ********************************************
    ok: [localhost] => (item={'image': 'docker.io/pycontribs/centos:7', 'name': 'centos_7', 'pre_build_image': True})
    ok: [localhost] => (item={'image': 'docker.io/pycontribs/centos:8', 'name': 'centos_8', 'pre_build_image': True})
    ok: [localhost] => (item={'image': 'docker.io/pycontribs/ubuntu:latest', 'name': 'ubuntu', 'pre_build_image': True})

    TASK [Create molecule instance(s)] *********************************************
    changed: [localhost] => (item=centos_7)
    changed: [localhost] => (item=centos_8)
    changed: [localhost] => (item=ubuntu)

    TASK [Wait for instance(s) creation to complete] *******************************
    [WARNING]: Docker warning: The requested image's platform (linux/amd64) does
    not match the detected host platform (linux/arm64/v8) and no specific platform
    was requested
    changed: [localhost] => (item={'failed': 0, 'started': 1, 'finished': 0, 'ansible_job_id': '220065989858.13848', 'results_file': '/Users/plosev/.ansible_async/220065989858.13848', 'changed': True, 'item': {'image': 'docker.io/pycontribs/centos:7', 'name': 'centos_7', 'pre_build_image': True}, 'ansible_loop_var': 'item'})
    changed: [localhost] => (item={'failed': 0, 'started': 1, 'finished': 0, 'ansible_job_id': '786544837789.13869', 'results_file': '/Users/plosev/.ansible_async/786544837789.13869', 'changed': True, 'item': {'image': 'docker.io/pycontribs/centos:8', 'name': 'centos_8', 'pre_build_image': True}, 'ansible_loop_var': 'item'})
    changed: [localhost] => (item={'failed': 0, 'started': 1, 'finished': 0, 'ansible_job_id': '423880352359.13891', 'results_file': '/Users/plosev/.ansible_async/423880352359.13891', 'changed': True, 'item': {'image': 'docker.io/pycontribs/ubuntu:latest', 'name': 'ubuntu', 'pre_build_image': True}, 'ansible_loop_var': 'item'})

    PLAY RECAP *********************************************************************
    localhost                  : ok=6    changed=2    unreachable=0    failed=0    skipped=5    rescued=0    ignored=0

    INFO     Running default > prepare
    WARNING  Skipping, prepare playbook not configured.
    INFO     Running default > converge

    PLAY [Converge] ****************************************************************

    TASK [Gathering Facts] *********************************************************
    ok: [ubuntu]
    ok: [centos_8]
    ok: [centos_7]

    TASK [Include vector-role] *****************************************************

    TASK [vector-role : Get facts on current container] ****************************
    ok: [centos_7]
    ok: [ubuntu]
    ok: [centos_8]

    TASK [vector-role : Ensure group 'vector' exists] ******************************
    changed: [ubuntu]
    changed: [centos_8]
    changed: [centos_7]

    TASK [vector-role : Add the user 'vector'] *************************************
    changed: [ubuntu]
    changed: [centos_8]
    changed: [centos_7]

    TASK [vector-role : Create the app directory for Vector] ***********************
    changed: [centos_7]
    changed: [ubuntu]
    changed: [centos_8]

    TASK [vector-role : Create the data directory for Vector] **********************
    changed: [centos_7]
    changed: [ubuntu]
    changed: [centos_8]

    TASK [vector-role : Create the config directory for Vector] ********************
    changed: [centos_7]
    changed: [ubuntu]
    changed: [centos_8]

    TASK [vector-role : Download Vector archive] ***********************************
    changed: [centos_7]
    changed: [centos_8]
    changed: [ubuntu]

    TASK [vector-role : Extract vector] ********************************************
    changed: [ubuntu]
    changed: [centos_8]
    changed: [centos_7]

    TASK [vector-role : Template a file to /etc/vector/vector.toml] ****************
    changed: [centos_7]
    changed: [ubuntu]
    changed: [centos_8]

    TASK [vector-role : Copy vector executable file to /usr/bin] *******************
    changed: [centos_7]
    changed: [ubuntu]
    changed: [centos_8]

    TASK [vector-role : Copy systemd service] **************************************
    changed: [centos_7]
    changed: [ubuntu]
    changed: [centos_8]

    RUNNING HANDLER [vector-role : Start vector service] ***************************
    skipping: [centos_7]
    skipping: [centos_8]
    skipping: [ubuntu]

    PLAY RECAP *********************************************************************
    centos_7                   : ok=12   changed=10   unreachable=0    failed=0    skipped=1    rescued=0    ignored=0
    centos_8                   : ok=12   changed=10   unreachable=0    failed=0    skipped=1    rescued=0    ignored=0
    ubuntu                     : ok=12   changed=10   unreachable=0    failed=0    skipped=1    rescued=0    ignored=0

    INFO     Running default > idempotence

    PLAY [Converge] ****************************************************************

    TASK [Gathering Facts] *********************************************************
    ok: [ubuntu]
    ok: [centos_8]
    ok: [centos_7]

    TASK [Include vector-role] *****************************************************

    TASK [vector-role : Get facts on current container] ****************************
    ok: [centos_7]
    ok: [ubuntu]
    ok: [centos_8]

    TASK [vector-role : Ensure group 'vector' exists] ******************************
    ok: [centos_7]
    ok: [ubuntu]
    ok: [centos_8]

    TASK [vector-role : Add the user 'vector'] *************************************
    ok: [ubuntu]
    ok: [centos_8]
    ok: [centos_7]

    TASK [vector-role : Create the app directory for Vector] ***********************
    ok: [centos_7]
    ok: [ubuntu]
    ok: [centos_8]

    TASK [vector-role : Create the data directory for Vector] **********************
    ok: [centos_7]
    ok: [ubuntu]
    ok: [centos_8]

    TASK [vector-role : Create the config directory for Vector] ********************
    ok: [centos_7]
    ok: [ubuntu]
    ok: [centos_8]

    TASK [vector-role : Download Vector archive] ***********************************
    ok: [centos_7]
    ok: [ubuntu]
    ok: [centos_8]

    TASK [vector-role : Extract vector] ********************************************
    ok: [ubuntu]
    ok: [centos_8]
    ok: [centos_7]

    TASK [vector-role : Template a file to /etc/vector/vector.toml] ****************
    ok: [centos_7]
    ok: [ubuntu]
    ok: [centos_8]

    TASK [vector-role : Copy vector executable file to /usr/bin] *******************
    ok: [centos_7]
    ok: [ubuntu]
    ok: [centos_8]

    TASK [vector-role : Copy systemd service] **************************************
    ok: [centos_7]
    ok: [ubuntu]
    ok: [centos_8]

    PLAY RECAP *********************************************************************
    centos_7                   : ok=12   changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
    centos_8                   : ok=12   changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
    ubuntu                     : ok=12   changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0

    INFO     Idempotence completed successfully.
    INFO     Running default > side_effect
    WARNING  Skipping, side effect playbook not configured.
    INFO     Running default > verify
    INFO     Running Ansible Verifier

    PLAY [Verify] ******************************************************************

    TASK [Example assertion] *******************************************************
    ok: [centos_7] => {
        "changed": false,
        "msg": "All assertions passed"
    }
    ok: [centos_8] => {
        "changed": false,
        "msg": "All assertions passed"
    }
    ok: [ubuntu] => {
        "changed": false,
        "msg": "All assertions passed"
    }

    PLAY RECAP *********************************************************************
    centos_7                   : ok=1    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
    centos_8                   : ok=1    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
    ubuntu                     : ok=1    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0

    INFO     Verifier completed successfully.
    INFO     Running default > cleanup
    WARNING  Skipping, cleanup playbook not configured.
    INFO     Running default > destroy

    PLAY [Destroy] *****************************************************************

    TASK [Set async_dir for HOME env] **********************************************
    ok: [localhost]

    TASK [Destroy molecule instance(s)] ********************************************
    changed: [localhost] => (item=centos_7)
    changed: [localhost] => (item=centos_8)
    changed: [localhost] => (item=ubuntu)

    TASK [Wait for instance(s) deletion to complete] *******************************
    changed: [localhost] => (item=centos_7)
    changed: [localhost] => (item=centos_8)
    changed: [localhost] => (item=ubuntu)

    TASK [Delete docker networks(s)] ***********************************************
    skipping: [localhost]

    PLAY RECAP *********************************************************************
    localhost                  : ok=3    changed=2    unreachable=0    failed=0    skipped=1    rescued=0    ignored=0

    INFO     Pruning extra files from scenario ephemeral directory
    ```
4. Добавьте несколько assert'ов в verify.yml файл для  проверки работоспособности vector-role (проверка, что конфиг валидный, проверка успешности запуска, etc). Запустите тестирование роли повторно и проверьте, что оно прошло успешно.
    #### Решение
    Добавил проверку конфига и запуска процесса vector'а в verify.yml:
    ```yaml
    ---
    - name: Verify
      hosts: all
      gather_facts: false
      tasks:
        - name: Validate vector config
        ansible.builtin.shell: /usr/bin/vector validate /etc/vector/vector.toml 
        become: true
        register: vector_validate

        - name: Assert that the vector config is valid
        assert:
        that: vector_validate.rc == 0
          success_msg: Vector config is valid
          fail_msg: Vector config is not valid

        - name: Validate vector process
        ansible.builtin.shell: pgrep vector
        become: true
        register: vector_process
            
        - name: Assert that vector is running
        assert:
          that: vector_process.rc == 0
          success_msg: Vector is running
          fail_msg: Vector is not running
    ```
    Результат проверки:
    ```
    ~/temp-netology-homeworks/08-ansible-04-role/vector-role ❯ molecule verify
    INFO     default scenario test matrix: verify
    INFO     Performing prerun...
    INFO     Set ANSIBLE_LIBRARY=/Users/plosev/.cache/ansible-compat/f5bcd7/modules:/Users/plosev/.ansible/plugins/modules:/usr/share/ansible/plugins/modules
    INFO     Set ANSIBLE_COLLECTIONS_PATH=/Users/plosev/.cache/ansible-compat/f5bcd7/collections:/Users/plosev/.ansible/collections:/usr/share/ansible/collections
    INFO     Set ANSIBLE_ROLES_PATH=/Users/plosev/.cache/ansible-compat/f5bcd7/roles:/Users/plosev/.ansible/roles:/usr/share/ansible/roles:/etc/ansible/roles
    INFO     Running default > verify
    INFO     Running Ansible Verifier
    INFO     Sanity checks: 'docker'

    PLAY [Verify] ******************************************************************

    TASK [Validate vector config] **************************************************
    changed: [ubuntu]
    changed: [centos_7]
    changed: [centos_8]

    TASK [Assert that the vector config is valid] **********************************
    ok: [centos_7] => {
        "changed": false,
        "msg": "Vector config is valid"
    }
    ok: [centos_8] => {
        "changed": false,
        "msg": "Vector config is valid"
    }
    ok: [ubuntu] => {
        "changed": false,
        "msg": "Vector config is valid"
    }

    TASK [Validate vector process] *************************************************
    changed: [ubuntu]
    changed: [centos_8]
    changed: [centos_7]

    TASK [Assert that vector is running] *******************************************
    ok: [centos_7] => {
        "changed": false,
        "msg": "Vector is running"
    }
    ok: [centos_8] => {
        "changed": false,
        "msg": "Vector is running"
    }
    ok: [ubuntu] => {
        "changed": false,
        "msg": "Vector is running"
    }

    PLAY RECAP *********************************************************************
    centos_7                   : ok=4    changed=2    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
    centos_8                   : ok=4    changed=2    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
    ubuntu                     : ok=4    changed=2    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
    ```
5. Добавьте новый тег на коммит с рабочим сценарием в соответствии с семантическим версионированием.
    #### Решение
    [Ссылка репозиторий vector](https://github.com/wowpetr/ansible-vector/tree/0.2-beta)

### Tox

1. Добавьте в директорию с vector-role файлы из [директории](./example)
2. Запустите `docker run --privileged=True -v <path_to_repo>:/opt/vector-role -w /opt/vector-role -it aragast/netology:latest /bin/bash`, где path_to_repo - путь до корня репозитория с vector-role на вашей файловой системе.
3. Внутри контейнера выполните команду `tox`, посмотрите на вывод.
5. Создайте облегчённый сценарий для `molecule` с драйвером `molecule_podman`. Проверьте его на исполнимость.
6. Пропишите правильную команду в `tox.ini` для того чтобы запускался облегчённый сценарий.
8. Запустите команду `tox`. Убедитесь, что всё отработало успешно.
9. Добавьте новый тег на коммит с рабочим сценарием в соответствии с семантическим версионированием.

После выполнения у вас должно получится два сценария molecule и один tox.ini файл в репозитории. Не забудьте указать в ответе теги решений Tox и Molecule заданий. В качестве решения пришлите ссылку на  ваш репозиторий и скриншоты этапов выполнения задания. 

## Необязательная часть

1. Проделайте схожие манипуляции для создания роли lighthouse.
2. Создайте сценарий внутри любой из своих ролей, который умеет поднимать весь стек при помощи всех ролей.
3. Убедитесь в работоспособности своего стека. Создайте отдельный verify.yml, который будет проверять работоспособность интеграции всех инструментов между ними.
4. Выложите свои roles в репозитории.

В качестве решения пришлите ссылки и скриншоты этапов выполнения задания.
