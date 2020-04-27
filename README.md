# topsdk
淘宝开放平台SDK go-topsdk

## sample

```golang
package main

import "github.com/daliyo/topsdk"

import "github.com/daliyo/topsdk/requests"

import "github.com/daliyo/topsdk/responses"

import "fmt"

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	apiGateway, appKey, appSecret := "", "", ""
	client, err := topsdk.NewTopClient(apiGateway, appKey, appSecret)
	checkErr(err)

	req := requests.NewTaobaoTimeGetRequest()
	res := responses.TaobaoTimeGetResponse{}

	err = client.Do(req, &res)
	checkErr(err)

	fmt.Println(res)
}
```
