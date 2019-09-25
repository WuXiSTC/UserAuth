# UserAuth

## Description

* 本项目是无锡市软件测试中心的用户验证服务
* 项目基于Go语言，数据库使用MySQL，并使用了Redis作为缓存

## Software Architecture

* 本项目由一系列API组成，分两种模式：
  * 主机模式：拥有数据库和缓存，可读可写，包含用户注册、用户验证、密码修改三个功能
  * 从机模式：仅有缓存，从机缓存与主机缓存同步，只读，仅包含用户验证功能

### 源码组织结构

本项目使用 iris 框架，按照典型的三层Web应用结构进行组织。从外到内对应的文件夹依次是`Controller`、`Service`、`Dao`。

* `Controller`：存放API的处理函数，与最外层的`Web API`直接相连，所有的Web请求皆直接传递到此文件夹的函数中，是对`Controller`层的`Web API`封装；
* `Service`：由`Controller`中各函数调用，处理网站业务逻辑，通过`Dao`层操作数据；
* `Dao`：封装一系列数据操作，向`Service`层隐藏数据库和缓存架构。

`Dao`层中的每个部件都只负责一种数据操作，例如向数据库中插入的部件就只负责数据库和缓存的插入而不关心用户名冲突验证，用户名冲突验证由`Service`层的新建用户的部件负责完成。而`Service`层的新建用户部件可以涉及数据库和缓存的形式。

项目提供了预置的Docker封装方案，可直接打包为容器化微服务

## Installation

### 准备

1. 安装并运行MySQL和Redis
2. 用Dao文件夹下的两个.sql文件建好MySQL数据库
3. 在配置文件中写上Redis和MySQL数据库的地址(详见配置文件说明)

### 使用Docker启动

#### 编译(可选步骤)

（项目根目录下已提供了一个用下面👇这条指令编译好的可执行文件，因此这步可以省略）

```shell
docker run -it --rm -v "$(PWD):/app" yindaheng98/go-iris go build -v -o UserAuth
```

1. 在项目根目录运行上面👆的指令
2. 编译完成后会在项目根目录下生成可执行文件UserAuth

#### 打包运行

1. 在项目根目录下运行`docker build .`打包一个镜像
2. 直接运行`docker run [刚才打包好的镜像]`或者将编辑好的配置文件用`-v`挂载到`/home`目录下运行
3. 用`docker-compose`当然也是可以的

### 不使用docker启动

1. 先把go装好
2. `go get github.com/kataras/iris`

   `go get github.com/go-sql-driver/mysql`

   `go get github.com/garyburd/redigo/redis`
3. `go build`

## Instructions

### 单主机模式

主机模式是通常的数据库+缓存模式

1. 在配置文件中`DatabaseConfig.yaml`和`CacheConfig.yaml`中写上Redis和MySQL数据库的地址，在配置文件`SlaveConfig.yaml`中配置主机模式(详见配置文件说明)
2. 按照上一节的说明启动

### 多主机模式

启动一个主机和多个从机，写任务交给主机，读任务交给从机，在此应用外面再套一层负载均衡或者用从机建一个CDN网络做成可水平扩展的分布式系统。Redis数据库之间的连接需要其他安全手段保护。

1. 要求先在其他地方以主机模式部署启动好一个本应用
2. 在配置文件中`CacheConfig.yaml`中写上Redis数据库的地址，在配置文件`SlaveConfig.yaml`中配置从机模式(详见配置文件说明)
3. 按照上一节的说明启动

### 主机模式POST格式

1. /register：用户注册
   * ID：要注册的用户名
   * PASS：密码
2. /verify：用户验证
   * ID：要验证的用户名
   * PASS：要验证的密码
3. /update：修改密码
   * ID：要修改的用户名
   * PASS：原始密码
   * newPASS：新密码

#### 从机模式POST格式

1. /ping：用于检查从属服务器是否在线
   * 无参数
2. /verify：用户验证同主服务

### 主机模式返回值

所有的返回值均为JSON格式：`{"ok":true|false,"message":"返回信息"}`

1. /register：
   * 用户注册信息成功写入数据库：`ok`为`true`，`message`为空
   * 用户注册信息写入失败：`ok`为`false`，`message`为注册失败原因（“用户名已存在”或错误信息）
2. /verify：用户验证
   * 数据库中可以查找到用户名和密码对应记录：`ok`为`true`，`message`为空
   * 否则：`ok`为`false`，`message`为“用户名或密码错误”或错误信息
3. /update：修改密码
   * 用户密码信息成功修改：`ok`为`true`，`message`为空
   * 否则：`ok`为`false`，`message`为“用户名或密码错误”或错误信息

### 从机模式返回值

1. /ping：
   * 在线且请求来自主服务器：返回PONG
   * 其他情况无返回
2. /verify：用户验证同主服务

## Contribution

1. Fork the repository
2. Create Feat_xxx branch
3. Commit your code
4. Create Pull Request

## Gitee Feature

1. You can use Readme\_XXX.md to support different languages, such as Readme\_en.md, Readme\_zh.md
2. Gitee blog [blog.gitee.com](https://blog.gitee.com)
3. Explore open source project [https://gitee.com/explore](https://gitee.com/explore)
4. The most valuable open source project [GVP](https://gitee.com/gvp)
5. The manual of Gitee [https://gitee.com/help](https://gitee.com/help)
6. The most popular members  [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)
