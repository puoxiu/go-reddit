

# 分布式ID生成-雪花算法
* 全局唯一
* 递增性
* 高性能
1. 

# 层
1. Controller---服务的入口，负责处理路由、参数校验、请求转发
2. Logic--逻辑处理，处理具体的事务  
3. Dao---数据库封装


# 参数校验
RE

# 用户认证

1. 原因：
* HTTP是无状态的，无法进行身份识别---》需要更上层的鉴权和认证

2.解决：
* cookie-session模式:
    * 过程
        * 客户端登录，通过cookie方式获取服务端sessionID
        * 客户端请求时，带上cokie，服务端认证
    * 问题
        * 占用资源
        * 扩展性差
        * CSRF攻击

* 基于token的无状态会话管理模式 
    * 过程
        * 客户端登录后获得一个唯一token，鉴权信息加密到token中
        * 访问需要权限的接口时，携带token，服务端只需要读取token中鉴权信息即可 进行验证
        * 
    * 流行方式 JWT（JSON Web Token）
        * JWT token格式
            header、payload、signature
    * go
        * JWT-go        

# 包循环引用的问题解决


# refreshtoken
作用：获取新的access token ，当之前的access token过期的时候
这个过期了就只能重新登陆了（一般可以设置30天过期）



# 登录设备数量限制-todo
redis: 存储: user_id --> token 进行验证


# 内存字节对齐
* 相同数据类型尽量放在一块


# 前后端id数值范围问题
* 后端id范围[2^64-1, 2^64-1], 前端id最大范围[-2^53-1, 2^53-1]
* 解决：-->字符串


# 投票
* 问题抽象：谁给哪个帖子投了什么票？
* 实际问题：
    1. 只能投一次
    2. 投的赞成还是反对

## redis常用数据类型
* string字符串
* Hash哈希
* List列表
* Set集合
* Sorted Set有序集合ZSet
* Bitmap位图
* 