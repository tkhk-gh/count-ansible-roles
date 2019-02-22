# count-ansible-roles

count-ansible-roles is count used ansible roles.

## requirements

- go 1.11
- gopkg.in/yaml.v2

## example

```
$ ls 
ansible.cfg     app.yml     inventory     roles     group_vars 

$ count-ansible-roles | sort
nginx: 1
python: 1
uwsgi: 1
```
