### GRPC golang实现
#### Protocol buffer介绍
##### 概念
&emsp;&emsp;protobuf是谷歌开源的一种数据格式，适合高性能，对响应速度有要求的数据传输场景。因为protobuf是二进制数据格式，需要编码和解码。数据本身不具备可读性。因此只能反序列化之后得到真正的可读的数据。
- 序列化后体积相比json和xml小很多，适合网络传输；
- 支持跨平台多语言
- 消息格式升级的兼容性好
- 序列化和反序列化速度很快
##### 安装
- 安装通用编译器protoc
  - ubuntu系统下安装
    ```bash
    sudo apt install proto-compiler
    ```
  - fedora系统下安装
    ```bash
    sudo dnf install proto-compiler
    ```
  - 二进制包下载安装=>[官方包](https://github.com/protocolbuffers/protobuf/releases)
- 安装go的protobuf插件
    ```bash
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

    cat >> ~/.bashrc << EOF

    export PATH=\$PATH:$(go env GOPATH)/bin

    EOF

    source ~/.bashrc
    ```
##### .proto扩展文件
- 基本语法说明
  1. 详见pb/userinfo.proto
  2. 注意看注释
- 字段类型映射
  | proto type | 说明 | C++ type | python type | go type |
  | :---: | :---: | :---: | :---: | :---: |
  | double |  | double | float | float64 |
  | float |  | float | float | float32 |
  | int32 | 使用变长编码，对于负值的效率很低 | int32 | int | int32 |
  | uint32 | 使用变长编码 | uint32 | int/long | uint32 |
  | uint64 | 使用变长编码 | uint64 | int/long | uint64 |
  | sint32 | 使用变长编码，负值时比int32高效很多 | int32 | int | int32 |
  | sint64 | 使用变长编码，负值时比int64高效很多 | int64 | int/long | int64 |
  | fixed32 | 总是4字节，如果数值总是比228大的话，这个类型比uint32高效 | uint32 | int | uint32 |
  | fixed64 | 总是8字节，如果数值总是比256大的话，这个类型比uint64高效 | uint64 | int/long | uint64 |
  | sfixed32 | 总是4字节 | int32 | int | int32 |
  | sfixed64 | 总是8字节 | int64 | int/long | int64 |
  | bool |  | bool | bool | bool |
  | string | 必须是UTF-8或者7-bit ASCII编码 | string | str/unicode | string |
  | bytes | 可能包含任意顺序的字节数据 | string | str | []byte |
- 生成go文件
  - 方式一（新版protoc-gen-go）
    ```bash
    # 1、命令会生成分别生成2个go文件
    # 2、一个是protobuf相关的解编码go文件
    # 3、一个是grpc的service,client接口go文件
    # 4、--go_out/--go-grpc_out 指定生成文件的根路径，且根路径必需存在（这个根据路径将与.proto文件中go_package中路径组合）
    # 5、paths=source_relative表示使用相对路径

    protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        helloworld/helloworld.proto
    ```
  - 方式二（旧版protoc-gen-go）
    ```bash
    # 1、命令会生成一个含protobuf,grpc的组合go文件
    # 2、--go_out 指定生成文件的根路径，且根路径必需存在（这个根据路径将与.proto文件中go_package中路径组合）

    protoc --go_out=plugins=grpc:./ pb/helloworld.proto
    ```
#### GRPC实例模板
##### 基本原理
- 待补充
##### 服务端
- 代码解析
  1. 详见sc/service.go
  2. 注意看注释
- 重点说明
  1. GRPC在初始化时，要注册proto中定义的service(服务)；
  2. 通过protoc生成go代码后，这个服务是一个接口类型；
  3. 这个接口定义了proto文件中service的所有rpc方法；
  4. 所以服务端需要根据自身的业务功能实现proto中定义的rpc相关的接口；
  5. proto3与proto2有一定区别，本文是使用proto3
##### 客户端
- 代码解析
  1. 详见sc/client.go
  2. 注意看注释
- 重点说明
#### 运行
- 生成protobuf,grpc的go代码
  ```bash
  cd ./pb
  ./gen-go.sh
  ```
- 编绎并运行demo
  ```bash
  cd ../
  go run main.go
  ```
