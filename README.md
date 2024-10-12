# Go开发——IM即时通讯系统



## 写在前面的话：

这个项目是我从B站上面找到的一个go开源项目， 虽然代码不知道放哪里了，但是教程比较详细并且是手写全程， 所以我觉得可以拿来学习学习。

主要学习方法是以视频为纲， 学习架构， 然后自己手敲代码。



## day0-2024-10-11

今天学习一天发现作者对框架的利用还算熟悉， 但是还没有到信手拈来的地步， 使用常有卡顿。并且编写的代码似乎不是Restful API的最佳实践， 我在作者的基础上做了一定改动， 让代码更规范和符合Restful API的最佳实践。

![image-20241011194546854](C:\Users\asus\AppData\Roaming\Typora\typora-user-images\image-20241011194546854.png)



`url` 应该是资源而不是动作， 动作要让`GET` `POST` 等动作来指定。

另外学会了`postman`

![image-20241011194650829](C:\Users\asus\AppData\Roaming\Typora\typora-user-images\image-20241011194650829.png)

这个比原作者使用的`swaggo`好用一些， 图形化界面对新手更友好。

目前来看这个项目有很明显的前后端分离的风格。



前端代码在`service`文件夹里， 这个文件夹包含从访问中读取信息， 以及返回信息的关键代码。

后端代码在`models` 文件夹里， 这个文件夹包含访问`mysql`（以及以后的`redis`）的代码。

初始化代码在`utils`文件夹里， 它目前主要包含了`mysql`的初始化代码。`DB`数据库对象， 后端的关键结构体在这个包里。



### 信息是怎么传递的

今天的编程感受就是， 信息的传递是编程中比较复杂的一部分。这里不再是平时C编程的键盘输入文件输入这么简单。任何输入的本质都是字节流， 但是通过API的包装让程序员可以输入输出结构体。这一点让编程简单了很多， 代价就是非常繁多的api， 让人不知道如何下手。



gin框架会对某个`url`绑定一个处理函数， 他对url的解析目前有两种：

![image-20241011195618454](C:\Users\asus\AppData\Roaming\Typora\typora-user-images\image-20241011195618454.png)



第一种就是在浏览器里输入， 第二种依然是按着格式输入， 然后我们可以获取`id`的值。

这里就是我说的复杂的地方。

现在我们知道id的值传进去了， 它的信息就在url里， 我们在编程的时候又该怎么获取呢？

用gin.Context的成员函数Param(s string)

![image-20241011195822357](C:\Users\asus\AppData\Roaming\Typora\typora-user-images\image-20241011195822357.png)



除了这种方式， 我们还有在url里输入?a1=123&a2=123233&...的方式传入信息， 这种又怎么读出来呢， 读到go里又是什么样的形式呢？

gin.Context提供了一些方式读出来, 今天用到的主要是json和form-Data

![image-20241011200305461](C:\Users\asus\AppData\Roaming\Typora\typora-user-images\image-20241011200305461.png)



ShouldBindJSON成员函数可以直接把json格式的输入读到map[string]interface{}里面

form-Data输入格式的则用PostForm函数根据Key读取Value

![image-20241011225615238](C:\Users\asus\AppData\Roaming\Typora\typora-user-images\image-20241011225615238.png)

### 后端的处理

后端其实就是接收前端传来的参数， 修改数据库里的内容。

重点是掌握增删查看数据库的框架`gorm`

今天主要使用的是`gorm`框架操作`mysql`数据库， 连接完毕可以返回一个gorm.db的指针结构体， 这是操作的核心。

连接部分基本是在`utils`包里完成的。

`utils`里用了`viper`外部库去设置`config`,  包括连接`mysql`的字符串等内容。

![image-20241011230139624](C:\Users\asus\AppData\Roaming\Typora\typora-user-images\image-20241011230139624.png)

这是初始化`mysql` 连接的函数， 顺便让每一步查询都有日志， 方便调试。
