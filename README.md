
在grpc目录下执行
protoc --proto_path=. --go_out=plugins=grpc:./user user.proto

* --protp_path是proto文件的目录(user.proto)

* --go_out 是输出的 proto文件编译成的结合grpc和protobuf的代码的输出目录、(./user)当前目录的user文件夹
后面跟的是要编译的proto文件

* 这个demo中的userInsert是模拟了一下连接数据库，这里使用的是MongoDB 例如：本地装个docker mongo 把db挂载出来
```
docker pull mongo  
docker run --name=mongo -d -p 27017:27017 -v /Users/zhangyong/Documents/docker/mongo:/data/db mongo
```