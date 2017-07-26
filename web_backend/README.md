## 系统后台
  采用Go的Restful设计思想去定义接口(PHP太麻烦了，不适合)

### 数据库设计
  [数据库文件](https://github.com/ITPai/ITPractice/tree/master/web_backend/schema/db_structure.sql)

### 接口规范
  * Request(http请求)
    * header：
    ```
    Accept: application/json;version=1.0
    Content-Type: application/json
    Token:p9sut9o7nmmog3qu99ki5lfvo7 //需要的时候

    ```
    * body：
    ```
    {
      "username":"testuser",
      "password":"testpass"
    }
    ```

  * Resources(资源)
    * example：
    ```
    GET /users      获取users列表
    GET /users/1    获取id为1的user信息
    POST /users     添加user
    PUT /users/1    修改id为1的user信息
    DELETE /users/1 删除id为1的user信息
    ```
    资源其实就是url的显示方式，不同的url代表不同的资源。而url映射到路由的方法：

    | url | method | action | 说明 |
    | ------| ------ | ------ | ------ |
    | /users | GET | index | 获取列表信息 |
    | /users/:id | GET | info | 获取详情 |
    | /users | POST | add | 添加信息 |
    | /users/:id | PUT | update | 修改信息 |
    | /users/:id | DELETE | delete | 删除信息 |

### 本地部署
  * 安装环境(go)
    根据[官方文档](http://docscn.studygolang.com/doc/)安装最新版本的golang
  * 获取mux包
    ```
    go get "github.com/gorilla/mux"
    ```
  * 运行
    ```
    go run src/v1.0/*.go
    ```
