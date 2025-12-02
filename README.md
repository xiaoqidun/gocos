# GoCOS

腾讯云对象存储(COS)，Drone CI插件，[AWS S3](http://plugins.drone.io/drone-plugins/drone-s3/)的COS实现

# Drone CI配置例子

```yml
kind: pipeline
type: docker
name: default

steps:
  - name: upload
    image: ccr.ccs.tencentyun.com/xiaoqidun/gocos
    settings:
      secret_id:
        from_secret: secret_id
      secret_key:
        from_secret: secret_key
      bucket_url:
        from_secret: bucket_url
      source_path: build/release
      target_path: build/release
      strip_prefix: build/release
```

# Drone CI配置说明

### secret_id

API密钥管理获得的SecretId

### secret_key

API密钥管理获得的SecretKey

### bucket_url

存储桶概览中的访问域名

### source_path

DroneCI中文件的源位置

### target_path

存储桶中文件的目标位置

### strip_prefix

从文件的源位置剔除前缀