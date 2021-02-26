package logic

import (
	"bilibiliSpider/internal/bhttp"
	"bilibiliSpider/internal/pkg"
	"bilibiliSpider/rrtype"
	"bilibiliSpider/stx"
	"fmt"
)

func GetUpInfo(stx *stx.StcConfig) {
	// 调用url获取到up的主页，然后获取所有video
	url := stx.UPList.URL
	upaids := stx.UPList.AIDs
	for _, aid := range upaids {
		var UPInfo rrtype.RespUPInfo
		geturl := fmt.Sprintf(url, aid, 1)
		fmt.Println(geturl)
		bhttp.ExecGet(&UPInfo, geturl, bhttp.Gethttp)
		// 将UPInfo写入文件中 UPInfo.
		pkg.WriteUPInfo(&UPInfo, aid)
		for i := 1; i*30 < UPInfo.Data.Count; i++ {
			geturl = fmt.Sprintf(url, aid, i+1)
			bhttp.ExecGet(&UPInfo, geturl, bhttp.Gethttp)
			// 将UPInfo写入文件中 UPInfo.
			pkg.WriteUPInfo(&UPInfo, aid)
		}
		fmt.Printf("%s 该用户易搜索完毕 \n", aid)
	}
	fmt.Println("所有UP主搜索完毕")
}
