## DM modules
[dm_master_servers]
dm_master ansible_host=__IPPREFIX__.101

[dm_worker_servers]
dm-worker1 ansible_host=__IPPREFIX__.101  server_id=101 source_id="mysql-replica-01" mysql_host=mysql1 mysql_user=root mysql_password='' mysql_port=3306

dm-worker2 ansible_host=__IPPREFIX__.102  server_id=102 source_id="mysql-replica-02" mysql_host=mysql2 mysql_user=root mysql_password='' mysql_port=3306

[dm_portal_servers]
dm_portal ansible_host=__IPPREFIX__.101

## Monitoring modules
[prometheus_servers]
prometheus ansible_host=__IPPREFIX__.101

[grafana_servers]
; grafana ansible_host=__IPPREFIX__.101
; change to add specified port for test, ref: https://docs.pingcap.com/zh/tidb-data-migration/dev/deploy-a-dm-cluster-using-ansible#%E9%BB%98%E8%AE%A4%E6%9C%8D%E5%8A%A1%E7%AB%AF%E5%8F%A3
grafana ansible_host=__IPPREFIX__.101 grafana_port=3001

[alertmanager_servers]
alertmanager ansible_host=__IPPREFIX__.101

## Global variables
[all:vars]
cluster_name = test-cluster

ansible_user = tidb

dm_version = v1.0.6

deploy_dir = /home/tidb/deploy

grafana_admin_user = "admin"
grafana_admin_password = "admin"
