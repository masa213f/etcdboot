# etcdboot

A deploy tool of etcd cluster.
You can easy to create and update systemd services for an etcd cluster with `etcdboot`.

## Usage

1. Create a cluster setting file as follows.

```
name: cluster
members:
  - name: member1
    clientURL: http://127.0.0.1:12379
    peerURL: http://127.0.0.1:12380
  - name: member2
    clientURL: http://127.0.0.1:22379
    peerURL: http://127.0.0.1:22380
  - name: member3
    clientURL: http://127.0.0.1:32379
    peerURL: http://127.0.0.1:32380
```

2. Run `etcdboot create`.

```
$ sudo etcdboot create -c example/cluster.yaml -m member1
create service file: /etc/systemd/system/etcd.service
```
