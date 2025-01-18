# 基础镜像
FROM alpine:3.21.2

# 作者信息
LABEL authors="xiaoqidun"

# 复制程序
COPY gocos /bin/gocos

# 启动命令
ENTRYPOINT ["/bin/gocos"]
