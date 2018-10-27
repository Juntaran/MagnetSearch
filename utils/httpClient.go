/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/10/27 19:41
  */

package utils

import (
	"log"
	"io/ioutil"
	"net/http"

	"magnetSearch/g"
)

func HttpGetClient(domain string) ([]byte, error) {
	req, err := http.NewRequest("GET", domain, nil)
	if err != nil {
		log.Println("NewRequest Error:", err)
	}
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("User-Agent", g.RandUA())

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println("DefaultClient Error:", err)
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	//log.Println("body:", string(body))
	return body, nil
}