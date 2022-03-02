# kratos-shop
kratos 框架写商品微服务

本项目是一个使用 Kratos 框架创建的很简单的微服务商城项目。
> 注: 本项目中但凡 kratos 提供包,就不会自己封装第三方的包。

主要是为了学习 kratos 如何使用,尤其各种中间件之间的调用,包括微服务的一些技术点。

项目具体目录结构初步设计如下:

```
|-- kratos-shop
    |-- service
        |-- user // 用户服务 grpc
        |-- goods // 商品服务 grpc
        |-- order // 订单服务 grpc
        |-- inventory // 库存服务服务 grpc
    |-- shop // shop 商城服务 http (后期会考虑把订单单独拆出来)
        ├── api
        │   ├── service
        │   │   └── user 
        │   │       └── v1 // 用户服务的 proto
        │   │   └── goods
        │   │       └── v1 // 商品服务的 proto
        │   │           
        │   └── shop
        │       └── v1
        │           ├── error_reason.proto 
        │           ├── shop.proto 
    |-- admin // 后端管理系统 http
```


![扫码提建议](https://gitee.com/aliliin/picture/raw/master/2022-3-2/1646202139081-WechatIMG7.jpeg)


