# Dev Note 
> A file to Note something while developing

## Go dependency  -- go module  
> Need to update Golang version to 1.13+
- Configure environment variable to set it on
```shell
 export GO111MODULE=on
```
- Initilization
```shell
 go mod init Goit
```
***Reference:*** 
```shell
go mod download #下载 go.mod 文件中指明的所有依赖
go mod tidy #整理现有的依赖，删除未使用的依赖。
go mod graph #查看现有的依赖结构
go mod init #生成 go.mod 文件 (Go 1.13 中唯一一个可以生成 go.mod 文件的子命令)
go mod edit #编辑 go.mod 文件
go mod vendor #导出现有的所有依赖 (事实上 Go modules 正在淡化 Vendor 的概念)
go mod verify #校验一个模块是否被篡改过
go clean -modcache #清理所有已缓存的模块版本数据。
go mod #查看所有 go mod 的使用命令。
```

