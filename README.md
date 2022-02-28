## SpiderWeb 项目 go 版本


### 当前环境配置
#### 前端
- node 12.1.0
- npm 6.9.0

#### 后端
- go 1.16.6


### 项目启动
- 采用前后端分离服务启动
#### 前端
```shell
cd frontend
npm install
# 默认启动0.0.0.0:8001 可自行修改配置文件
npm run serve
```
#### 后端
```shell
# 默认启动0.0.0.0:8000,可自行修改conf/service_x.conf,同时修改前端对应配置文件
# linux
./fresh -c ./fresh.conf
# win
./fresh.exe -c ./fresh.conf
```

#### 访问
- 浏览器输入ip:port, eg:localhost:8001 **port为前端服务端口**