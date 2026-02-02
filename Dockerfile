# 基础镜像
FROM alpine:3.23.3
# 作者信息
LABEL authors="xiaoqidun"

# 复制程序
COPY gocos /bin/gocos

# 启动命令
ENTRYPOINT ["/bin/gocos"]
