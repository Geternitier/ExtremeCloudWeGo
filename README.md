# ExtremeCloudWeGo 说明文档

version: 0.0

本文档介绍了ExtremeCloudWeGo项目，包括以下内容：

- 小组成员
- 接口说明
- IDL管理
- 部署步骤

## 小组成员
- 211250036 黄天宇
- 211250005 周楚皓
- 211250033 时国皓

## 接口说明
1. Method=POST   absolutePath=/add-student-info
2. Method=GET    absolutePath=/query

### 接口：登记学生信息

```
POST http://127.0.0.1:8888/add-student-info
```

#### 请求参数

请求参数为JSON格式，包含学生信息，字段如下：

- id i32: 学号
- name string: 姓名
- college College: 学院
  - name string: 名称
  - address string: 地址
- email list<string>: 邮箱（可选）

#### 响应参数

响应参数为JSON格式，包含以下字段：

- success bool: 是否成功
- message string: 信息

#### 示例

请求：

```
POST http://127.0.0.1:8888/add-student-info HTTP/1.1
Host: Exvs
Content-Type: application/json

{
    "id": 7802, 
    "name":"Gundam", 
    "college": {
        "name": "Earth Federation Force", 
        "address": "Jabro"
    }
}
```

响应：

```
HTTP/1.1 200 OK
Content-Type: application/json

{
    "success": "true",
    "message": "Add 7802 GUNDAM to Database."
}
```

### 接口：询问学生信息

查询n号学生信息：
```http request
GET http://127.0.0.1:8888/query?id=n
```

#### 请求参数

无

#### 响应参数

响应参数为JSON格式，包含学生信息，字段如下：

- id i32: 学号
- name string: 姓名
- college College: 学院
    - name string: 名称
    - address string: 地址
- email list<string>: 邮箱（可选）

#### 示例

请求：
```http request
GET http://127.0.0.1:8888/query?id=7802
```

响应：
```http request
HTTP/1.1 200 OK
Content-Type: application/json

{
    "id": 7802, 
    "name":"Gundam", 
    "college": {
        "name": "Earth Federation Force", 
        "address": "Jabro"
    }
}
```

## IDL管理

仅针对IDL更新服务编写IDL文件(idl/IDL.thrift)，http端使用该IDL产生相应kitex客户端，调用其IDL更新服务进行IDL热更新
- 这里http端使用kitex调用客户端向kitex服务端发送请求并获得响应

目前更新方式为单独启动一个goroutine，每10s更新一次IDL

## 部署步骤

1. 安装Go语言环境
2. 安装Kitex和Hertz，过程详见CloudWeGo官网
3. 运行etcd数据库
4. 运行servers/main.go，启动学生服务
5. 运行http/main.go，启动http服务
