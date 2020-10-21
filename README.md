# Go API Learning
这是一篇按照《Go API 开发实战》构造出的学习笔记，仅供参考

## 教程1

### 准备阶段
* 如何安装和配置 Go 开发环境
* 如何安装和配置 Vim IDE

### 设计阶段
* API 构建技术选型
* API 基本原理
* API 规范设计

### 开发阶段
* 如何读取配置文件
* 如何管理和记录日志
* 如何做数据库的 CURD 操作
* 如何自定义错误 Code
* 如何读取和返回 HTTP 请求
* 如何进行业务逻辑开发
* 如何对请求插入自己的处理逻辑
* 如何进行 API 身份验证
* 如何进行 HTTPS 加密
* 如何用 Makefile 管理 API 源码
* 如何给 API 命令添加版本功能
* 如何管理 API 命令
* 如何生成 Swagger 在线文档

### 测试阶段
* 如何进行单元测试
* 如何进行性能测试（函数性能）
* 如何做性能分析
* API 性能测试和调优

### 部署阶段
* 如何用 Nginx 部署 API 服务
* 如何做 API 高可用

## 教程2
### RESTful API
本质是通过一套标准响应方式来发送信息对服务器上的东西进行操作(CRUD之类)。

可以套用标准HTTP协议，比如可以用GET查，POST增，PUT改，DELETE删。

### RPC
RPC相当于一个接口，这个接口实现的功能就类似调用服务器的函数然后输出结果。

比如把一个要执行的操作序列化成JSON格式，然后TCP发送到特定的API接口，然后服务器就会进行操作，并返回结果。

## 教程3
### 目录结构
```text
├── admin.sh                     # 进程的start|stop|status|restart控制文件
├── conf                         # 配置文件统一存放目录
│   ├── config.yaml              # 配置文件
│   ├── server.crt               # TLS配置文件
│   └── server.key
├── config                       # 专门用来处理配置和配置文件的Go package
│   └── config.go                
├── db.sql                       # 在部署新环境时，可以登录MySQL客户端，执行source db.sql创建数据库和表
├── docs                         # swagger文档，执行 swag init 生成的
│   ├── docs.go
│   └── swagger
│       ├── swagger.json
│       └── swagger.yaml
├── handler                      # 类似MVC架构中的C，用来读取输入，并将处理流程转发给实际的处理函数，最后返回结果
│   ├── handler.go
│   ├── sd                       # 健康检查handler
│   │   └── check.go 
│   └── user                     # 核心：用户业务逻辑handler
│       ├── create.go            # 新增用户
│       ├── delete.go            # 删除用户
│       ├── get.go               # 获取指定的用户信息
│       ├── list.go              # 查询用户列表
│       ├── login.go             # 用户登录
│       ├── update.go            # 更新用户
│       └── user.go              # 存放用户handler公用的函数、结构体等
├── main.go                      # Go程序唯一入口
├── Makefile                     # Makefile文件，一般大型软件系统都是采用make来作为编译工具
├── model                        # 数据库相关的操作统一放在这里，包括数据库初始化和对表的增删改查
│   ├── init.go                  # 初始化和连接数据库
│   ├── model.go                 # 存放一些公用的go struct
│   └── user.go                  # 用户相关的数据库CURD操作
├── pkg                          # 引用的包
│   ├── auth                     # 认证包
│   │   └── auth.go
│   ├── constvar                 # 常量统一存放位置
│   │   └── constvar.go
│   ├── errno                    # 错误码存放位置
│   │   ├── code.go
│   │   └── errno.go
│   ├── token
│   │   └── token.go
│   └── version                  # 版本包
│       ├── base.go
│       ├── doc.go
│       └── version.go
├── README.md                    # API目录README
├── router                       # 路由相关处理
│   ├── middleware               # API服务器用的是Gin Web框架，Gin中间件存放位置
│   │   ├── auth.go 
│   │   ├── header.go
│   │   ├── logging.go
│   │   └── requestid.go
│   └── router.go
├── service                      # 实际业务处理函数存放位置
│   └── service.go
├── util                         # 工具类函数存放目录
│   ├── util.go 
│   └── util_test.go
└── vendor                         # vendor目录用来管理依赖包
    ├── github.com
    ├── golang.org
    ├── gopkg.in
    └── vendor.json
```
当然，为了紧跟时代，我们使用go mod进行包管理。
方法:<kbd>go mod init</kbd>

## 教程4
虽然推荐使用vim进行编程，不过我还是使用GoLand+自动同步来完成远程开发。

使用Go版本1.14。

## 教程5
教程中监测的部分的源代码有微小的错误，在返回响应代码时出现了一些问题；进行修改，统一使用一个api来获取信息<kbd>/sd/monitor</kbd>，返回数据使用JSON格式。处理的过程中发现了使用JSON需要把结构体中的变量命名为大写字母开头，否则无法正常解析。

## 教程6
这一把我们要实现一个读取配置文件。虽然我个人倾向于使用JSON格式，不过还是用一把YAML吧。确实在Go环境中使用JSON格式有点麻烦。这回我们使用一个叫做Viper的工具来实现配置读取及配置热更新，而不是写死的参数。

## 教程7
在教程中，要求我们利用教程作者开发的日志包进行日志记录。于是，我们就用这个教程作者开发的日志包来进行日志填写。

## 教程8
在教程中，我们需要新建表来实现一个数据库的连接。从安全层面考虑，应该使用低权限非root账户来操作数据表。所以我新建了一个用户apiserver来完成数据库操作。

```sql
CREATE USER apiserver@localhost IDENTIFIED BY 'apiserver';
FLUSH privileges;
source db.sql
GRANT ALL ON db_apiserver.* TO apiserver@localhost;
FLUSH privileges;
```

## 教程9
我们操作一波，来个数据库连接。这里按照教程实现了对于两个数据库的连接(不过此时是同一个数据库)，利用了<kbd>Golang</kbd>的<kbd>gorm</kbd>库完成。

## 教程10
错误码编写我继续使用了教程自带的错误码。由于一般错误码都会有个统一规范，每个项目具体的错误码规范可能不太一样，所以我这一把就直接复用教程的错误码处理程序了，然后针对一些情况做了一点点修改。对错误码进行测试的时候我们可以用<kbd>curl</kbd>发包，不过这样操作有时候会比较麻烦。所以我使用Python编写了一个小的脚本进行发包探测。文件放在主目录下Request_Test.py，使用Python3编写，提供了<kbd>POST</kbd>和<kbd>GET</kbd>等自定义参数探测，也可以用来做一点压力测试。

## 教程11
感觉这一节的教程意义不大……就是把返回值包装了下放在函数中。其中<kbd>gin</kbd>可以通过检测网址中变量的形式来解析参数稍微有点意思，我只写了这一部分功能。其他的都没啥太多意思。

考虑到RESTful API的问题，我这次测试了两个方法：<kbd>PUT</kbd>和<kbd>POST</kbd>。同时我在这个版本使用了多级路由分发，探测HTTP Method并返回在后台。

## 教程12
这次考虑到自定义表前缀，为了方便书写，在<kbd>yaml</kbd>配置文件中新建参数<kbd>db_prefix</kbd>，使用这个作为表前缀。其他业务逻辑正常书写。

同时，我们使用教程提供的<kbd>auth.go</kbd>来进行账户密码加密。

<kbd>Create</kbd>添加<kbd>NewRecord</kbd>函数，判断是否是新行(用Where查询后判断，保证其不存在后可以插入)。

将SQL表的格式改成蛇形小写的形式，<kbd>UserModel</kbd>换成<kbd>User</kbd>，尽量保证和gorm的默认情况一致。

本来打算重写删除逻辑，但是懒得重新写可以解析<kbd>DELETE</kbd>中<kbd>body</kbd>部分的接口了，那么就略过。

## 教程13
这次需要增加一个中间件用来确定日志中的<kbd>Request-ID</kbd>,方便对整个请求链进行分析。所以这回我们就来加个中间件。

教程提供了<kbd>Logging</kbd>和<kbd>RequestId</kbd>这两个中间件，我们就直接复用了。

对于<kbd>Logging</kbd>这个中间件，修复了源代码中一些和这个工程具体实现不一样的地方，同时删除了多余的代码。

## 教程14
复用了教程<kbd>JWT-Token</kbd>部分的编写和<kbd>Auth</kbd>部分的编写，我觉得写得很不错所以就直接复用了。

现在写一个设置<kbd>Authorized</kbd>的小型登录API。

修改了<kbd>Token Invalid</kbd>的逻辑，增加了一个提示登录的<kbd>ErrLoginFirst</kbd>，代替<kbd>TokenInvalid</kbd>，更加具体。

## 教程15
按照教程的签名证书的方式给了个新的证书，使用<kbd>RunTLS</kbd>来运行服务器。

在<kbd>8088</kbd>端口上运行HTTPS服务器同时校验证书。

## 教程16
使用Makefile对工程进行了组织，比较简单。

基本就是直接复用了教程的内容；