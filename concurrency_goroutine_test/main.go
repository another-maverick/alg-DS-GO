package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

var (
	inputJson = `{
   "gophers": [{
         "name": "A",
         "sleep": 1,
         "eat": 1
      },
      {
         "name": "B",
         "sleep": 4,
         "eat": 3
      },
      {
         "name": "C",
         "sleep": 1,
         "eat": 4
      },
      {
         "name": "D",
         "sleep": 5,
         "eat": 2
      },
      {
         "name": "E",
         "sleep": 2,
         "eat": 3
      }
   ],
   "totalFood": 30
}`
)

type Gopher struct {
	Name string
	Sleep time.Duration
	Eat int
}

type Farm struct {
	Gophers []Gopher `json:"gophers"`
	TotalFood int `json:"totalfood"`
	sync.Mutex
}

func (f *Farm) eatFood(gopher *Gopher) error{
	f.Lock()
	defer f.Unlock()

	if f.TotalFood < gopher.Eat {
		log.Printf("Gopher %s wants to eat %v food. There is not enough food!", gopher.Name, gopher.Eat)
		return errors.New("Not enough food!")
	}

	f.TotalFood -= gopher.Eat
	log.Printf("%s Gopher will be eating %v quantity of food. %v will be left", gopher.Name, gopher.Eat, f.TotalFood)
	return nil
}

func (g *Gopher) gopherLive(farm *Farm, messages chan string){
	for{
		time.Sleep(time.Second * g.Sleep)
		if err := farm.eatFood(g); err != nil{
			messages <- fmt.Sprintf("Gopher %s has done eating. It dies. ",g.Name)
			return
		}
	}

}



func main() {
	gopherFarm := &Farm{}
	if err := json.Unmarshal([]byte(inputJson), &gopherFarm); err != nil {
		log.Fatal("cannot unmarshall the input JSON ", err)
	}

	messages := make(chan string)

	for i, gopher := range gopherFarm.Gophers{
		log.Printf("gopher %s is joining the farm -- ", gopher.Name)
		go gopherFarm.Gophers[i].gopherLive(gopherFarm, messages)
	}

	log.Println("Waiting.")

	for range gopherFarm.Gophers {
		msg, ok :=  <-messages
		if !ok {
			fmt.Println("channel has been closed!")
		}
		log.Println(msg)

	}
	log.Println("Done Waiting!")

}
