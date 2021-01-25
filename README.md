
## Teamdo - 团队任务管理 (配套学习golang服务ddd架构规范)

### 如何初始化本地开发环境
1. 配置数据库域名`127.0.0.1 db.dev.com`
2. 确保能以root:root访问本地数据库
3. 创建数据库`create database teamdo`
4. 授权数据库`grant all on teamdo.* to teamdo@'127.0.0.1' identified by 'root'`
5. 配置本地环境变量
   ```
   BEEGO_RUNMODE = dev
   BEEGO_MODE = dev
   ENABLE_DEV_TEST_RESOURCE = 1
   ```
6. 执行`go run commands/cmd.go orm syncdb -v`安装数据库
7. 执行`start_service.bat`启动服务

### 如何集成到Ningx？
1. 在hosts文件中添加如下域名
```
127.0.0.1 devapi.vxiaocheng.com
127.0.0.1 db.dev.com
```

2. 编辑Nginx的`nginx.conf`文件, 增加server配置
```
server {
    listen       80;
    server_name  devapi.vxiaocheng.com;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    access_log  logs/api_gateway.access.log;
    
    location /teamdo/ {
      #重写去掉url中server name部分
      rewrite ^/teamdo(.*)/ /$1 break;
      proxy_pass http://127.0.0.1:7001;
    }
}
```

### 需求
```
1. 用户可以注册
2. 用户注册后可以登录，并查看已参与的项目列表
3. 用户可以管理项目（即用户创建了项目成为项目管理员）
4. 项目管理员可以邀请其他用户加入项目（未邀请的用户无法查看项目）
5. 项目管理员可以管理泳道（即任务状态列），包括创建、删除、编辑、排序等
6. 项目管理员可以管理任务（非项目管理员无法创建和移动任务），包括创建、移动、删除，编辑等
7. 项目管理员可以指定某个用户为一个任务的执行者
8. 任务的执行者可以移动任务到各泳道中
9. 项目管理员可以为任务添加子任务，一个完成的任务，必须是所有子任务都已完成
10. 项目的所有成员可以在每个任务中发表评论
11. 泳道中的任务是按优先级和已耗时间排序的，越紧急的，越接近过期时间的排在上面
12. 项目成员可以查看各任务的状态迁移日志，内容包括，操作人、操作内容、操作时间
13. 项目管理员可以查看项目的统计图表，包含每日、每月任务状态
```