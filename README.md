### go-zero超简单入门简易教程
本项目基于go-zero框架，除logic外其余大部分代码均由goctl生成，本项目由两个http子服务组成（两个主要用于展示service group的用法），项目结构按照名称定义划分。
```js
├── api      // 数据库脚本
├── db       // 数据库脚本
├── models   // 数据模型
├── services // 服务列表
    ├── article
    ├── user
├── utils    // 通用方法和工具
├── main.go  // 主入口
```

<br/>

##### 快速开始

1、环境依赖
* go ([点击链接🔗查看](https://go-zero.dev/docs/tasks))
* goctl ([点击链接🔗查看，需要先装go](https://go-zero.dev/docs/tasks/installation/goctl))
* mysql
* redis (虽然配置了，但项目逻辑并未使用，可以删掉这部分代码😜)

2、连接本地数据库执行数据库脚本
```bash
$ mysql -u root -p

$ use 你自己定义的数据库名字;

// 此处复制db/下的sql脚本执行
```


3、拉取项目代码
```bash
$ git clone git@github.com:Vanforu/go-zero-demo.git
```

4、打开项目，搜索所有的#TODO
将  #TODO: 这里自己改自己的对应配置 改成自己的配置内容

5、项目运行
```bash
# 整理依赖文件
$ go mod tidy
# 启动 go 程序
$ go run demo.go
```

关于该项目编写过程以及效果图
[【环境配置 - 文章链接地址】](https://blog.csdn.net/m0_37723113/article/details/138160859?spm=1001.2014.3001.5502)
[【项目表设计 - 文章链接地址】](https://blog.csdn.net/m0_37723113/article/details/138162985)
[【项目初始化 - 文章链接地址】](https://blog.csdn.net/m0_37723113/article/details/138164699)


最后
如果想0-1体验如何使用go-zero进行开发，删掉本项目中除api/和db/外的所有内容，参照[【项目表设计 - 文章链接地址】](https://blog.csdn.net/m0_37723113/article/details/138162985)和[【项目初始化 - 文章链接地址】](https://blog.csdn.net/m0_37723113/article/details/138164699)进行开发即可。
