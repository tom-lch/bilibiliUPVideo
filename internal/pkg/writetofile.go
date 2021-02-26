package pkg

import (
	"bilibiliSpider/rrtype"
	"bufio"
	"fmt"
	"os"
)

// 读取结构题写文件
func WriteUPInfo(upinfo *rrtype.RespUPInfo, aid string) {
	for _, videoELE := range upinfo.Data.VList {
		writeTotxt(aid, []string{videoELE.Title, "https://" + videoELE.PIC[2:], "https://www.bilibili.com/video/" + videoELE.BVID}...)
	}
}

func writeTotxt(aid string, datas ...string) {
	path := fmt.Sprintf("./txt/%s.txt", aid)
	write(path, datas...)
}

func checkFileExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func write(path string, datas ...string) {
	var file *os.File
	var err error
	file, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bufferWrite := bufio.NewWriter(file)
	if err != nil {
		panic(err)
	}
	for _, data := range datas {
		for _, v := range data {
			// 将数据写入缓冲区
			bufferWrite.WriteString(string(v))
		}
		bufferWrite.WriteString("\t")
	}
	bufferWrite.WriteString("\n")
	bufferWrite.Flush()
}
