/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/10/16 23:08
  */

package crawler

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli"

	"MagnetSearch/g"
	"MagnetSearch/utils"
)

func Search(ctx *cli.Context)  {
	keyword := "sakuramomo"
	numbers := 10
	sortWay := "ctime"
	detail  := true

	if ctx.IsSet("keyword") {
		keyword = ctx.String("keyword")
	}
	if ctx.IsSet("number") {
		numbers = ctx.Int("number")
	}
	if numbers > 100 || numbers < 0 {
		numbers = 10
	}
	if ctx.IsSet("sort") {
		switch ctx.Int("sort") {
		case 1:
			sortWay = "ctime"
		case 2:
			sortWay = "click"
		case 3:
			sortWay = "length"
		default:
			sortWay = "ctime"
		}
	}
	if !ctx.IsSet("detail") {
		detail = false
	}

	var resources []utils.ResourceInfo

	for i:=1; i<(numbers+9)/10+1; i++ {
		url := g.BITDOMAIN + "/search/" + keyword + "_" + sortWay + "_" + strconv.Itoa(i) + ".html"
		body, err := utils.HttpGetClient(url)
		if err != nil {
			fmt.Println("Error: search stop")
			return
		}
		resources = append(resources, utils.SearchHtml(string(body))...)
	}

	if numbers < len(resources) {
		resources = resources[:numbers]
	}

	for k, v := range resources {
		if detail == true {
			fmt.Println("序号:", k+1)
			fmt.Println("标题:", v.Name)
			fmt.Println("磁链:", v.Link)
			fmt.Println("大小:", v.Size)
			fmt.Println("日期:", v.Time)
			fmt.Println("热度:", v.Rank)
			fmt.Println()
		} else {
			fmt.Println(v.Link, "\t", v.Size, "\t", v.Time)
		}
	}
}