package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func getGasPrice(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	// time.Sleep(time.Duration(i * 100 * int(time.Millisecond)))

	url := ""
	var jsonStr = []byte(`{"jsonrpc": "2.0", "method": "eth_sendRawTransaction", "params": ["unknown"], "id": 1}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func main() {
	var wg sync.WaitGroup
	s := time.Now().UnixMilli()
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go getGasPrice(&wg, i)
	}

	wg.Wait()
	res := time.Now().UnixMilli() - s

	fmt.Println(res)
}
