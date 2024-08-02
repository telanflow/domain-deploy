# Domain-Admin SSL Deploy

本项目为 [domain-admin](https://github.com/mouday/domain-admin) 的证书部署服务

- 支持Domain Admin项目证书自定义部署
- 通过控制3个header字段进行自定义部署
- 配置文件进行命令映射防止危险命令
- 通过请求头选择执行哪个命令

## 如何使用

安装可执行文件后：
```shell
# 修改 config.yml 配置文件，配置自定义token
vim config.yml

# 启动
./domain-deploy -c config.yml
```

Domain-Admin API部署配置：
```shell
接口地址：http://localhost:51000/issueCertificate

请求头：
{
    "Token": "xxxxxxx",
    "Deploy-Cmd": "cmd_simple_nginx",
    "Key-Save-Path": "/usr/local/nginx/conf/cert/"
}
```

## 安装为System服务

```shell
# 安装成linux服务
./domain-deploy install

# 执行成功后会创建文件：vim /etc/systemd/system/domain-deploy.service

# 启动
sudo systemctl start domain-deploy

# 停止
sudo systemctl stop domain-deploy

# 重启
sudo systemctl restart domain-deploy

# 状态
sudo systemctl status domain-deploy

```

## 手动编译

```shell
# 切换到项目根目录
make build

# 指定编译为目标系统
make linux
make windows
make darwin
make darwin_arm64

```
