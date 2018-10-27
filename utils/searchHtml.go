/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/10/27 19:55
  */

package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ResourceInfo struct {
	Name string
	Time string
	Size string
	Rank int
	Link string
}

func SearchHtml(html string) []ResourceInfo {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var resources []ResourceInfo

	dom.Find(".media-body").Each(func(i int, selection *goquery.Selection) {
		name := strings.TrimSpace(selection.Find("h4").Find("a").Text())
		s := selection.Find(".media-more")
		time := s.Find(".label-success").Text()
		size := s.Find(".label-warning").Text()
		rank := s.Find(".label-primary").Text()
		link := ""
		href, IsExist := s.Find("a").Attr("href")
		if IsExist == true {
			href = strings.TrimSpace(href)
			if strings.Contains(href, "index/magnet") {
				href = strings.Split(href, "index")[1]
				href = href[1:]
			}
			if strings.Contains(href, "&") {
				href = strings.Split(href, "&")[0]
			}
			link = href
		}

		var resourceInfo = new(ResourceInfo)
		resourceInfo.Name = name
		resourceInfo.Time = time
		resourceInfo.Rank, _ = strconv.Atoi(rank)
		resourceInfo.Size = size
		resourceInfo.Link = link

		resources = append(resources, *resourceInfo)
	})

	return resources
}
