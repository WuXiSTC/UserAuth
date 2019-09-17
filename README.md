# UserAuth

## Description

* 本项目是无锡市软件测试中心的用户验证服务
* 项目基于Go语言，数据库使用MySQL，并使用了Redis作为缓存

## Software Architecture

* 本项目由一系列API组成，分两种模式：
  * 主机模式：拥有数据库和缓存，可读可写，包含用户注册、用户验证、密码修改三个功能
  * 从机模式：仅有缓存，缓存与主机缓存同步，只读，仅包含用户验证功能

### 代码结构

本项目使用 iris 框架，按照典型的三层Web应用结构进行组织。从外到内对应的文件夹依次是`Controller`、`Service`、`Dao`。

* `Controller`：存放API的处理函数，与最外层的`Web API`直接相连，所有的Web请求皆直接传递到此文件夹的函数中，是对`Controller`层的`Web API`封装；
* `Service`：由`Controller`中各函数调用，处理网站业务逻辑，通过`Dao`层操作数据；
* `Dao`：封装一系列数据操作，向`Service`层隐藏数据库和缓存架构。

## Installation

### 准备

1. 安装redis
2. `go get github.com/kataras/iris`

   `go get github.com/go-sql-driver/mysql`

   `go get github.com/garyburd/redigo/redis`
3. `git clone https://gitee.com/WuXiSTC/UserAuth`

### 以主机模式启动

1. 安装mysql
2. 用Dao文件夹下的两个.sql文件建mysql数据库
3. 在配置文件中写上redis和mysql的地址，并配置为主机模式(详见配置文件说明)
4. `cd UserAuth`
5. `go run main.go`

### 以从机模式启动

1. 在别的地方以主机模式启动一个本应用
2. 在配置文件中写上以主机模式启动的那个应用的redis数据库地址和端口号(详见配置文件说明)
3. `cd UserAuth`
4. `go run main.go`

## Instructions

### 单主机模式

按照上文“以主机模式启动”的说明启动单个主机，这是通常的数据库+缓存模式

### 多主机模式

启动一个主机和多个从机，写任务交给主机，读任务交给从机，在此应用外面再套一层负载均衡或者用从机建一个CDN网络

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
