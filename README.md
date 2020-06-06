# topclient
淘宝开放平台SDK go-topclient

## sample

```golang
package main

import "github.com/daliyo/topclient"

import "github.com/daliyo/topclient/requests"

import "github.com/daliyo/topclient/responses"

import "fmt"

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	apiGateway, appKey, appSecret := "", "", ""
	client, err := topclient.NewTopClient(apiGateway, appKey, appSecret)
	checkErr(err)

	req := requests.NewTaobaoTimeGetRequest()
	res := responses.TaobaoTimeGetResponse{}

	err = client.Do(req, &res)
	checkErr(err)

	fmt.Println(res)
}
```
