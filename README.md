
## Teamdo - 团队任务管理

### 如何初始化本地开发环境
1. 配置数据库域名`127.0.0.1 db.dev.com`
2. 确保能以root:root访问本地数据库
3. 创建数据库`create database teamdo`
4. 授权数据库`grant all on teamdo.* to teamdo@'127.0.0.1' identified by 'root'`
5. 执行`go run commands/cmd.go orm syncdb -v`安装数据库
6. 执行`start_service.bat`启动服务
7. 访问http://127.0.0.1:6021/console/console/，成功获取api console页面

### 如何集成到Ningx？
1. 在hosts文件中添加如下域名
```
127.0.0.1 devapi.vxiaocheng.com
127.0.0.1 db.dev.com
```

2. 编辑Nginx的`nginx.conf`文件，在api.weapp.com配置中添加location
```
    location /teamdo/ {
      #重写去掉url中server name部分
      rewrite ^/teamdo(.*)/ /$1 break;
      #http 协议
      proxy_pass http://127.0.0.1:7001;
    }
```