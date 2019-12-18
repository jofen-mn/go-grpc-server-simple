package main

import (
	"runtime"
	"sync"
	"time"
	"go-grpc-server-simple/inf"
	"math/rand"
	"strconv"
	"strings"
	"log"
	"google.golang.org/grpc"
	"fmt"
	"golang.org/x/net/context"
	"encoding/json"
)

var (
	wg sync.WaitGroup
)

const (
	networkType = "tcp"
	server = "127.0.0.1"
	port = "41005"
	parallel = 50
	times = 10
)


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	now := time.Now()
	//for i := 0; i < int(parallel) ; i++  {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		exe()
	//	}()
	//}
	//
	//wg.Wait()

	exe()

	log.Printf("time taken: %.2f ", time.Now().Sub(now).Seconds())
	fmt.Println("hello world.")
}

func exe()  {
	conn, _ := grpc.Dial(server + ":" + port)
	defer conn.Close()
	client := inf.NewDataClient(conn)

	for i := 0; i < int(times); i++ {
		getUser(client)
	}

}

func getUser(client inf.DataClient)  {
	request := &inf.UserRq{}
	r := rand.Intn(parallel)
	request.Id = int32(r)

	if client == nil {
		fmt.Println("client is nil.")
		return
	}

	if request == nil {
		fmt.Println("request is nil.")
		return
	}

	temStr, _ := json.Marshal(client)

	log.Printf("client: %s", string(temStr))
	log.Println()
	log.Printf("request: %+v", request)
	log.Println()

	response, err := client.GetUser(context.Background(), request)

	if err != nil {
		log.Println("getUser error: %s ", err.Error())
		return
	}

	if id, _ := strconv.Atoi(strings.Split(response.Name, ":")[0]); id != r {
		log.Printf("response error  %#v", response)
	}
}
