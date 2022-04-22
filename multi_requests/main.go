package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

var waitgroup = sync.WaitGroup{}

// func MakeRequest(i int, url string, ch chan<- string) {
// 	start := time.Now()
// 	resp, _ := http.Get(url)
// 	secs := time.Since(start).Seconds()
// 	body, _ := ioutil.ReadAll(resp.Body)
// 	ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
// }

func MakeRequest(i int, endpoint []string, ch chan<- string) {
	start := time.Now()
	fmt.Println(start)
	method := "POST"

	payload := strings.NewReader("{\"query\":\"mutation PaymentSessionResolve($id: ID!, $authorizationExpiresAt: DateTime) {\\n  paymentSessionResolve(id: $id, authorizationExpiresAt: $authorizationExpiresAt) {\\n    paymentSession {\\n      id\\n      status {\\n        code\\n      }\\n      nextAction {\\n        action\\n        context {\\n          ... on PaymentSessionActionsRedirect {\\n            redirectUrl\\n          }\\n        }\\n      }\\n    }\\n    userErrors {\\n      field\\n      message\\n    }\\n  }\\n}\",\"variables\":{\"id\":\"" + endpoint[2] + "\"}}")

	client := &http.Client{}
	req, err := http.NewRequest(method, endpoint[0], payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("X-Shopify-Access-Token", endpoint[1])
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	end := time.Now()
	fmt.Println(end)

	// fmt.Printf("%?.2fs elapsed\n", time.Hour.Seconds())
	waitgroup.Done()
}

func main() {
	start := time.Now()
	endpointData := [][]string{
		{
			"https://mason-test-01.myshopify.com/payments_apps/api/2021-10/graphql.json",
			"xxx",
			"gid://shopify/PaymentSession/qun47Gkp3V2EAz3CV9hunlc-",
		},
		{
			"https://mason-test-01.myshopify.com/payments_apps/api/2021-10/graphql.json",
			"xxx",
			"gid://shopify/PaymentSession/d8T--POVJK5FCNfEZjr7vg9a",
		},
	}

	ch := make(chan string)
	waitgroup.Add(2)
	for i, endpoint := range endpointData {
		go MakeRequest(i, endpoint, ch)
	}

	// for range os.Args[1:] {
	// 	fmt.Println(<-ch)
	// }

	waitgroup.Wait()
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
