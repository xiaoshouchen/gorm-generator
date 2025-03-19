# gorm-generator

## Introduction

this project is used to generate gorm model and repository from mysql database and other databases.

## Usage

### install

```shell
go install  github.com/xiaoshouchen/gorm-generator@latest
```

### run

```shell
gorm-generator
```


## config

```json
{
  "databases": [
    {
      "scheme": "panda-trip",
      "package_name": "model",
      "connect": {
        "type": "mysql",
        "host": "localhost",
        "port": "3306",
        "user": "root",
        "password": "password"
      },
      "table_filter_option": "all",
      "tables": [],
      "cache_expiration": [
        {
          "name": "users",
          "time": 3600
        }
      ]
    }
  ]
}
```
