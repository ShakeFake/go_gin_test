# go_gin_test
## 打包
    go build -o test .

## 镜像
    docker build -t yunkai:test .

## 镜像运行
    docker run -d -p 8091:8091 --name test yunkai:test