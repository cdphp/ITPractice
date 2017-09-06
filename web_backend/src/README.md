### 本地部署
  * 安装环境(go)
    根据[官方文档](http://docscn.studygolang.com/doc/)安装最新版本的golang

  * 设置GOPATH
    以linux为例:($WorkDir表示你本机的工作目录)
    ```
    vim /etc/profile
    export GOPATH = $WorkDir/ITPractice/web_backend

    ```
### 项目说明

* 运行
    ```
    cd src
    go run main.go
    ```

* 项目结构
  ```
  configs         #配置
  controllers     #控制器
  lib             #库文件
  models          #数据模型
  main.go         #主文件
  ```
