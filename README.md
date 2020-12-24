# gocos
腾讯云对象存储，Drone CI插件

# Drone CI配置例子
```yml
kind: pipeline
type: docker
name: default

steps:
- name: upload
  image: xiaoqidun/gocos
  settings:
    secret_id:
      from_secret: secret_id
    secret_key:
      from_secret: secret_key
    bucket_url:
      from_secret: bucket_url
    source: build/release
    target: build/release
    strip_prefix: build/release
```
