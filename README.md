# cronsun [![Build Status](https://travis-ci.org/shunfei/cronsun.svg?branch=master)](https://travis-ci.org/shunfei/cronsun)

`cronsun` is a distributed cron-style job system. It's similar with `crontab` on stand-alone `*nix`.

[简体中文](README_ZH.md)

## Purpose

The goal of this project is to make it much easier to manage jobs on lots of machines and provides high availability.
`cronsun` is different from [Azkaban](https://azkaban.github.io/), [Chronos](https://mesos.github.io/chronos/), [Airflow](https://airflow.incubator.apache.org/).

## Features

- Easy manage jobs on multiple machines
- Managemant panel
- Mail service
- Multi-language support

## Architecture

```
                                                [web]
                                                  |
                                     --------------------------
           (add/del/update/exec jobs)|                        |(query job exec result)
                                   [etcd]                 [mongodb]
                                     |                        ^
                            --------------------              |
                            |        |         |              |
                         [node.1]  [node.2]  [node.n]         |
             (job exec fail)|        |         |              |
          [send mail]<-----------------------------------------(job exec result)

```


## Security

`cronsun` support security with `security.json` config. When `opne=true`， job command is only allow local files with special extension on the node.

```json
{
    "open": true,
    "#users": "allowed execution users",
    "users": [
        "www", "db"
    ],
    "#ext": "allowed execution file extension",
    "ext": [
        ".cron.sh", ".cron.py"
    ]
}
```

## Getting started

### Setup / installation

Install from binary [releases](https://github.com/shunfei/cronsun/releases), download and unzip.

Or build from source, require `go >= 1.7+`

```
go get -u github.com/shunfei/cronsun
cd $GOPATH/src/github.com/shunfei/cronsun
sh build.sh
```

### Run

1. Install [MongoDB](http://docs.mongodb.org/manual/installation/)
2. Install [etcd3](https://github.com/coreos/etcd)
3. Modify config in `conf` dir
4. Start Node: `./cronnode -conf conf/base.json`, start Web: `./cronweb -conf conf/base.json`
5. Opne `http://127.0.0.1:7079/ui/` with the browser

## Screenshot

**Brief**:

![](doc/img/Cronsun_dashboard_en.png)

**Exec result**:

![](doc/img/Cronsun_log_list_en.png)
![](doc/img/Cronsun_log_item_en.png)

**Job**:

![](doc/img/Cronsun_job_list_en.png)

![](doc/img/Cronsun_job_new_en.png)

**Node**:

![](doc/img/Cronsun_node_en.png)

## Credits

cron is base on [robfig/cron](https://github.com/robfig/cron)
