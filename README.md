# board_backend
up在线协作白板后端的实现

项目的服务端代码，由于项目架构设计之初希望大部分功能尽可能在客户端实现，使客户端实现直接的分布式通信，服务器不过多参与其过程。

目前服务器的作用仅仅作为图床与附件服务器使用。

# 使用说明

```shell
# clone 项目
git clone https://github.com/SIT-board/board_backend.git
# 进入 board_backend 目录
cd board_backend
# 使用 go mod 并安装go依赖包
go generate

# linux 编译运行
go build -o server main.go 
./server

# windows
go build -o server.exe main.go
./server.exe
```
应用会默认监听  8080 端口，默认保存文件的目录是 ./files 。如需修改配置，请在 [config.yaml](config.yaml)中修改

# API 文档
## 上传文件
### 请求
```json
POST /file/upload
Content-Type: multipart/form-data; boundary=<Boundary>

--<Boundary>
Content-Disposition: form-data; name="file";filename="filename"
<FileContent>
```
### 响应
```json
{
  "url":"<scheme>://<host>/file/download/<2006_01_02>_<uuid>_<filename>"
}
```
### 参数说明
- <scheme> ：http 或者 https
- <host> :  server 的host
- <2006_01_02> :  上传时间的格式化
- <uuid> : 生成的uuid
- <filename> ： form 表单中的filename

## 下载文件
```json
GET /file/download/<filename>
```
