package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/ilyatos/logic.stress/pkg/client"
	"github.com/ilyatos/logic.stress/pkg/helpers"
	"github.com/joho/godotenv"
	"log"
	"net"
	"net/url"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type atomicCount uint32

func (ac *atomicCount) increment() {
	atomic.AddUint32((*uint32)(ac), 1)
}

var failedLaunchesCount atomicCount
var usersCount uint
var launchesCount uint
var host string

func init() {
	flag.UintVar(&usersCount, "u", 1, "Number of users")
	flag.UintVar(&launchesCount, "l", 1, "Number of launches that will be performed for each user")
	flag.StringVar(&host, "host", "http://logic.hacktory.ai", "Logic host to stress")
	flag.Parse()
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	u, err := url.Parse(host)
	if err != nil {
		log.Fatalln(err)
	}

	c := client.NewClient(u, os.Getenv("LOGIC_AUTH_TOKEN"), nil)

	wg := new(sync.WaitGroup)
	for i := 1; i <= int(usersCount); i++ {
		wg.Add(1)
		go run(c, "test"+strconv.Itoa(i)+"a", wg)
	}
	wg.Wait()
	fmt.Println("FINISHED! Failed starts:", failedLaunchesCount)
}

func run(c *client.Client, subdomain string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < int(launchesCount); i++ {
		user, err := c.GetUser(subdomain)
		if err != nil {
			log.Fatalln(err)
		}

		sl := &client.LabStart{
			Subdomain:  user.Subdomain,
			TemplateId: 1,
			IP:         net.IPv4(100, 99, 98, 97),
		}
		err = c.StartLab(sl)
		if err != nil {
			log.Fatalln(err)
		}

		err = waitForCompletedStatus(c, user, i, true)
		if err != nil {
			log.Println(err)
		} else {
			err = c.StopLab(user)
			if err != nil {
				log.Fatalln(err)
			}
		}

		waitForCompletedStatus(c, user, i, false)
	}
}

func waitForCompletedStatus(c *client.Client, user *client.User, launch int, isStarting bool) error {
	ticker := time.NewTicker(2 * time.Second)
	for range ticker.C {
	GetLabStatus:
		st, err := c.GetLabStatus(user)
		if err != nil {
			log.Println(user, err)
			goto GetLabStatus
		}

		helpers.PrintLabState(user, launch, st)
		if isStarting && (st.State == "stop" || st.State == "stopping" || st.State == "stopped") {
			failedLaunchesCount.increment()
			return errors.New("unexpected stopping is occurred")
		}
		if st.Status == 100 {
			ticker.Stop()
			break
		}
	}

	return nil
}
