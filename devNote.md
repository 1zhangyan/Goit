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
## Goit 存储对象的创建blob object
### git 中的存储对象 blob  
以一个文件为例子，git add命令会将修改后的文件计算一个sha1签名，将这个签名作为文件名新建文件，将用户修改的文件内容使用deflate压缩，作为新文件的文件内容，将新文件存储在object目录下。
具体操作上，为了便于寻找，git将object目录下的这个文件，按照sha1签名的前两位进行新建目录，将后面位作为真正的文件名存储。这样，sha1值前两位相等的文件会被放在同一个文件目录下方。  

### goit 实现该方法的简单例子
```go

//压缩
func DoZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

//解压缩
func DoZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r , _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}

//sha1函数
func Dohashs(src []byte) []byte {
	h := sha1.New()
	h.Write(src)
	return h.Sum(nil)
}

func main() {
var filePath string = "test.txt"
content,err := ioutil.ReadFile(filePath)
if err != nil {
    fmt.Println("Read fail" , err)
}

header:= "blob " + strconv.Itoa(len(content))+"\x00"
fmt.Println(header)
newcontent := append([]byte(header) , content...)
fmt.Printf("header is  %x \n",header)
fmt.Printf("sha1 is %x \n",Dohashs(newcontent))
fmt.Printf("zlib newcontent is %x \n" , DoZlibCompress(newcontent))
fmt.Printf("zlib content is %x \n" , DoZlibCompress(content))

os.Mkdir("testgit/"+hex.EncodeToString(Dohashs(newcontent))[:2] ,os.ModePerm)

fp, err := os.Create("testgit/"+hex.EncodeToString(Dohashs(newcontent))[:2]+"/"+hex.EncodeToString(Dohashs(newcontent))[2:])
if err != nil {
    fmt.Println(err)
    return
}
defer fp.Close()
buf := new(bytes.Buffer)
binary.Write(buf, binary.LittleEndian,  DoZlibCompress(newcontent))
fp.Write(buf.Bytes())

}

```
上述代码按照git的实现思路使用golang硬加了一个object 对象,
为了测试我们的对象和真正的 git blob 对象是否相兼容，我们可以新建 testgit 这个 git repo，使用 git 的原生命令查看该对象。
```shell
mkdir testgit
cd testgit
git init
#创建了一个testgit repo
#执行上面的 go 代码,上面的代码是测试代码，路径已经写死了，go 代码文件名为 test.go ,文件位置和 testgit 文件夹在同一个目录下
vi test.txt #将hello+回车写入test.txt
go run test.go
cd testgit
git cat-file -t ce013625030ba8dba906f756967f9e9ca394464a #查看对象类型
git cat-file -p ce013625030ba8dba906f756967f9e9ca394464a #查看对象内容
```
