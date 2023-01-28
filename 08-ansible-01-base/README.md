# Домашнее задание к занятию "8.1. Введение в Ansible"

## Основная часть
1. Попробуйте запустить playbook на окружении из `test.yml`, зафиксируйте какое значение имеет факт `some_fact` для указанного хоста при выполнении playbook'a.
    #### Решение
    ```bash
    ❯ ansible-playbook -i inventory/test.yml site.yml 

    PLAY [Print os facts] ***************************************************************************************************************************************************************************************

    TASK [Gathering Facts] **************************************************************************************************************************************************************************************
    [WARNING]: Platform darwin on host localhost is using the discovered Python interpreter at /opt/homebrew/bin/python3.11, but future installation of another Python interpreter could change the meaning of
    that path. See https://docs.ansible.com/ansible-core/2.14/reference_appendices/interpreter_discovery.html for more information.
    ok: [localhost]

    TASK [Print OS] *********************************************************************************************************************************************************************************************
    ok: [localhost] => {
        "msg": "MacOSX"
    }

    TASK [Print fact] *******************************************************************************************************************************************************************************************
    ok: [localhost] => {
        "msg": 12
    }

    PLAY RECAP **************************************************************************************************************************************************************************************************
    localhost                  : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0 
    ```
    Ответ: playbook вернул `some_fact` равным 12.

2. Найдите файл с переменными (group_vars) в котором задаётся найденное в первом пункте значение и поменяйте его на 'all default fact'.
    #### Решение
    Файл с переменной задающей `some_fact` - `group_vars/all/examp.yml`.  Заменил значение на `some_fact` на требуемое:
    ```yml
    ---
    some_fact: "all default fact"
    ```
3. Воспользуйтесь подготовленным (используется `docker`) или создайте собственное окружение для проведения дальнейших испытаний.
    #### Решение
    Запустил два контейнера docker:
    ```bash
    ❯ docker run -it -d --name centos7 --rm centos:7
    ❯ docker run -it -d --name ubuntu --rm ubuntu
    ```
    Так как в ubuntu docker не установлен python, то подключился к контейнеру и установил его:
    ```bash
    ❯ docker exec -it ubuntu bash
    root@42714b38e72f:/# apt-get update
    root@42714b38e72f:/# apt-get install python3
    ```
4. Проведите запуск playbook на окружении из `prod.yml`. Зафиксируйте полученные значения `some_fact` для каждого из `managed host`.
    #### Решение
    ```bash
    ❯ ansible-playbook -i inventory/prod.yml site.yml

    PLAY [Print os facts] ***************************************************************************************************************************************************************************************

    TASK [Gathering Facts] **************************************************************************************************************************************************************************************
    ok: [ubuntu]
    ok: [centos7]

    TASK [Print OS] *********************************************************************************************************************************************************************************************
    ok: [centos7] => {
        "msg": "CentOS"
    }
    ok: [ubuntu] => {
        "msg": "Ubuntu"
    }

    TASK [Print fact] *******************************************************************************************************************************************************************************************
    ok: [centos7] => {
        "msg": "el"
    }
    ok: [ubuntu] => {
        "msg": "deb"
    }

    PLAY RECAP **************************************************************************************************************************************************************************************************
    centos7                    : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   
    ubuntu                     : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0 
    ```
    Получено значение `el` для centos7 и `deb` для ubuntu.

5. Добавьте факты в `group_vars` каждой из групп хостов так, чтобы для `some_fact` получились следующие значения: для `deb` - 'deb default fact', для `el` - 'el default fact'.
    #### Решение
    Изменены файлы:  
    group_vars/deb/examp.yml
    ```yml
    ---
    some_fact: "deb default fact"
    ```
    group_vars/el/examp.yml
    ```yml
    ---
    some_fact: "el default fact"
    ```
6.  Повторите запуск playbook на окружении `prod.yml`. Убедитесь, что выдаются корректные значения для всех хостов.
    #### Решение
    ```bash
    ❯ ansible-playbook -i inventory/prod.yml site.yml

    PLAY [Print os facts] ***************************************************************************************************************************************************************************************

    TASK [Gathering Facts] **************************************************************************************************************************************************************************************
    ok: [ubuntu]
    ok: [centos7]

    TASK [Print OS] *********************************************************************************************************************************************************************************************
    ok: [centos7] => {
        "msg": "CentOS"
    }
    ok: [ubuntu] => {
        "msg": "Ubuntu"
    }

    TASK [Print fact] *******************************************************************************************************************************************************************************************
    ok: [centos7] => {
        "msg": "el default fact"
    }
    ok: [ubuntu] => {
        "msg": "deb default fact"
    }

    PLAY RECAP **************************************************************************************************************************************************************************************************
    centos7                    : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   
    ubuntu                     : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0 
    ```

7. При помощи `ansible-vault` зашифруйте факты в `group_vars/deb` и `group_vars/el` с паролем `netology`.
    #### Решение
    ```bash
    ❯ ansible-vault encrypt group_vars/deb/examp.yml
    ❯ ansible-vault encrypt group_vars/el/examp.yml
    ```
8. Запустите playbook на окружении `prod.yml`. При запуске `ansible` должен запросить у вас пароль. Убедитесь в работоспособности.
    #### Решение
    ```bash
    ❯ ansible-playbook -i inventory/prod.yml site.yml --ask-vault-pass 
    Vault password: 

    PLAY [Print os facts] ***************************************************************************************************************************************************************************************

    TASK [Gathering Facts] **************************************************************************************************************************************************************************************
    ok: [ubuntu]
    ok: [centos7]

    TASK [Print OS] *********************************************************************************************************************************************************************************************
    ok: [centos7] => {
        "msg": "CentOS"
    }
    ok: [ubuntu] => {
        "msg": "Ubuntu"
    }

    TASK [Print fact] *******************************************************************************************************************************************************************************************
    ok: [centos7] => {
        "msg": "el default fact"
    }
    ok: [ubuntu] => {
        "msg": "deb default fact"
    }

    PLAY RECAP **************************************************************************************************************************************************************************************************
    centos7                    : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   
    ubuntu                     : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
    ```
9. Посмотрите при помощи `ansible-doc` список плагинов для подключения. Выберите подходящий для работы на `control node`.
    #### Решение
    ```
    ❯ ansible-doc -t connection -l
    ansible.builtin.local          execute on controller                                                                                                                                                    
    ansible.builtin.paramiko_ssh   Run tasks via python ssh (paramiko)                                                                                                                                      
    ansible.builtin.psrp           Run tasks over Microsoft PowerShell Remoting Protocol                                                                                                                    
    ansible.builtin.ssh            connect via SSH client binary                                                                                                                                            
    ansible.builtin.winrm          Run tasks over Microsoft's WinRM                                                                                                                                         
    ansible.netcommon.grpc         Provides a persistent connection using the gRPC protocol                                                                                                                 
    ansible.netcommon.httpapi      Use httpapi to run command on network appliances                                                                                                                         
    ansible.netcommon.libssh       Run tasks using libssh for ssh connection                                                                                                                                
    ansible.netcommon.netconf      Provides a persistent connection using the netconf protocol                                                                                                              
    ansible.netcommon.network_cli  Use network_cli to run command on network appliances                                                                                                                     
    ansible.netcommon.persistent   Use a persistent unix socket for connection                                                                                                                              
    community.aws.aws_ssm          execute via AWS Systems Manager                                                                                                                                          
    community.docker.docker        Run tasks in docker containers                                                                                                                                           
    community.docker.docker_api    Run tasks in docker containers                                                                                                                                           
    community.docker.nsenter       execute on host running controller container                                                                                                                             
    community.general.chroot       Interact with local chroot                                                                                                                                               
    community.general.funcd        Use funcd to connect to target                                                                                                                                           
    community.general.iocage       Run tasks in iocage jails                                                                                                                                                
    community.general.jail         Run tasks in jails                                                                                                                                                       
    community.general.lxc          Run tasks in lxc containers via lxc python library                                                                                                                       
    community.general.lxd          Run tasks in lxc containers via lxc CLI                                                                                                                                  
    community.general.qubes        Interact with an existing QubesOS AppVM                                                                                                                                  
    community.general.saltstack    Allow ansible to piggyback on salt minions                                                                                                                               
    community.general.zone         Run tasks in a zone instance                                                                                                                                             
    community.libvirt.libvirt_lxc  Run tasks in lxc containers via libvirt                                                                                                                                  
    community.libvirt.libvirt_qemu Run tasks on libvirt/qemu virtual machines                                                                                                                               
    community.okd.oc               Execute tasks in pods running on OpenShift                                                                                                                               
    community.vmware.vmware_tools  Execute tasks inside a VM via VMware Tools                                                                                                                               
    containers.podman.buildah      Interact with an existing buildah container                                                                                                                              
    containers.podman.podman       Interact with an existing podman container                                                                                                                               
    kubernetes.core.kubectl        Execute tasks in pods running on Kubernetes

    ```

    `ansible.builtin.local` используется для работы на Ansible controller.
    
10. В `prod.yml` добавьте новую группу хостов с именем  `local`, в ней разместите localhost с необходимым типом подключения.
    #### Решение
    Добавил новую группу `local` в `inventory/prod.yml`:
    ```yml
    ---
    el:
        hosts:
        centos7:
            ansible_connection: docker
    deb:
        hosts:
        ubuntu:
            ansible_connection: docker
    local:
        hosts:
        local-test-1:
            ansible_connection: local
    ```
11. Запустите playbook на окружении `prod.yml`. При запуске `ansible` должен запросить у вас пароль. Убедитесь что факты `some_fact` для каждого из хостов определены из верных `group_vars`.
    #### Решение
    ```bash
    ❯ ansible-playbook -i inventory/prod.yml site.yml --ask-vault-pass
    Vault password: 

    PLAY [Print os facts] ***************************************************************************************************************************************************************************************

    TASK [Gathering Facts] **************************************************************************************************************************************************************************************
    [WARNING]: Platform darwin on host local-test-1 is using the discovered Python interpreter at /opt/homebrew/bin/python3.11, but future installation of another Python interpreter could change the meaning
    of that path. See https://docs.ansible.com/ansible-core/2.14/reference_appendices/interpreter_discovery.html for more information.
    ok: [local-test-1]
    ok: [ubuntu]
    ok: [centos7]

    TASK [Print OS] *********************************************************************************************************************************************************************************************
    ok: [centos7] => {
        "msg": "CentOS"
    }
    ok: [ubuntu] => {
        "msg": "Ubuntu"
    }
    ok: [local-test-1] => {
        "msg": "MacOSX"
    }

    TASK [Print fact] *******************************************************************************************************************************************************************************************
    ok: [centos7] => {
        "msg": "el default fact"
    }
    ok: [ubuntu] => {
        "msg": "deb default fact"
    }
    ok: [local-test-1] => {
        "msg": "all default fact"
    }

    PLAY RECAP **************************************************************************************************************************************************************************************************
    centos7                    : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   
    local-test-1               : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   
    ubuntu                     : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0

    ```

## Необязательная часть

1. При помощи `ansible-vault` расшифруйте все зашифрованные файлы с переменными.
    #### Решение
    ```bash
    ❯ ansible-vault decrypt group_vars/deb/examp.yml
    ❯ ansible-vault decrypt group_vars/el/examp.yml 
    ```
2. Зашифруйте отдельное значение `PaSSw0rd` для переменной `some_fact` паролем `netology`. Добавьте полученное значение в `group_vars/all/exmp.yml`.
    #### Решение
    ```bash
    ❯ ansible-vault encrypt_string 'PaSSw0rd'
    New Vault password: 
    Confirm New Vault password: 
    Encryption successful
    !vault |
            $ANSIBLE_VAULT;1.1;AES256
            62306330653262386235353236633531646164643537313839663234633066663063616538633535
            3166666335653734323539393830383063343966306663360a623061646262626665396132346438
            36303432383066383339343763373566636132623437336534666436346637353330656131616164
            6336396432393830610a663963376362386265616365313636396363616664343838363666663863
            3230 
    ```

3. Запустите `playbook`, убедитесь, что для нужных хостов применился новый `fact`.
    #### Решение
    Новое значение `some_fact` не применится так как у нас нет хостов в группе `all`. У нас `centos7` в группе `el`, `ubuntu` в группе `deb`, а новый хост в группе `local`. Поэтому, видимо, ошибка в задании.
    ```
    ❯ ansible-playbook -i inventory/prod.yml site.yml --ask-vault-pass
    Vault password: 

    PLAY [Print os facts] ***************************************************************************************************************************************************************************************

    TASK [Gathering Facts] **************************************************************************************************************************************************************************************
    [WARNING]: Platform darwin on host local-test-1 is using the discovered Python interpreter at /opt/homebrew/bin/python3.11, but future installation of another Python interpreter could change the meaning
    of that path. See https://docs.ansible.com/ansible-core/2.14/reference_appendices/interpreter_discovery.html for more information.
    ok: [local-test-1]
    ok: [ubuntu]
    ok: [centos7]

    TASK [Print OS] *********************************************************************************************************************************************************************************************
    ok: [centos7] => {
        "msg": "CentOS"
    }
    ok: [ubuntu] => {
        "msg": "Ubuntu"
    }
    ok: [local-test-1] => {
        "msg": "MacOSX"
    }

    TASK [Print fact] *******************************************************************************************************************************************************************************************
    ok: [centos7] => {
        "msg": "el default fact"
    }
    ok: [ubuntu] => {
        "msg": "deb default fact"
    }
    ```