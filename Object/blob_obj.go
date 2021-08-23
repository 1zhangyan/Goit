package Object

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

type BlobObject struct {
	name string ;
	len  int64 ;
	content []byte ;
}

func doZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()
}

//解压缩
func doZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r , _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}

//sha1函数
func dohashs(src []byte) []byte {
	h := sha1.New()
	h.Write(src)
	return h.Sum(nil)
}

func AddBlob (repoPath string , fileName string){
	content,err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Read fail" , err)
	}
	header:= "blob " + strconv.Itoa(len(content))+"\x00"
	newcontent := append([]byte(header) , content...)
	obejctPath := repoPath+"/.git/objects/"
	os.Mkdir(obejctPath+hex.EncodeToString(dohashs(newcontent))[:2] ,os.ModePerm)
	fmt.Printf("makdir in "+obejctPath+hex.EncodeToString(dohashs(newcontent))[:2]+"\n")
	fp,_:= os.Create(obejctPath+hex.EncodeToString(dohashs(newcontent))[:2]+"/"+hex.EncodeToString(dohashs(newcontent))[2:])
	fmt.Println("Create file in " + obejctPath+hex.EncodeToString(dohashs(newcontent))[:2]+"/"+hex.EncodeToString(dohashs(newcontent))[2:]+"\n")
	defer fp.Close()
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian,  doZlibCompress(newcontent))
	fp.Write(buf.Bytes())
}




