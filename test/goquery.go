/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/10/27 20:01
  */

package main

import (
	"strings"
	"log"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	html := `<body>

				<div>DIV1</div>
				<div class="name">DIV2</div>
				<span>SPAN</span>

			</body>
			`

	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))
	if err!=nil{
		log.Fatalln(err)
	}

	dom.Find("div[class]").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}