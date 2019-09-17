# UserAuth

#### Description

* 本项目是无锡市软件测试中心的用户验证服务
* 项目基于Go语言，数据库使用MySQL，并使用了Redis作为缓存

#### Software Architecture

* 本项目由一系列API组成，包含用户注册、用户验证、密码修改三个功能

#### Installation

1. 安装redis，使运行于`localhost:6379`
2. `go get github.com/kataras/iris`

   `go get github.com/go-sql-driver/mysql`

   `go get github.com/garyburd/redigo/redis`
3. `git clone https://gitee.com/WuXiSTC/UserAuth`
4. `cd UserAuth`
5. `go run main.go`

#### Instructions

##### POST列表

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

##### 返回值列表

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

#### Contribution

1. Fork the repository
2. Create Feat_xxx branch
3. Commit your code
4. Create Pull Request

#### Gitee Feature

1. You can use Readme\_XXX.md to support different languages, such as Readme\_en.md, Readme\_zh.md
2. Gitee blog [blog.gitee.com](https://blog.gitee.com)
3. Explore open source project [https://gitee.com/explore](https://gitee.com/explore)
4. The most valuable open source project [GVP](https://gitee.com/gvp)
5. The manual of Gitee [https://gitee.com/help](https://gitee.com/help)
6. The most popular members  [https://gitee.com/gitee-stars/](https://gitee.com/gitee-stars/)