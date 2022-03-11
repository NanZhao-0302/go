package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

func main() {
	count := 1000
	scope := FatRateScope{
		min: 0.0,
		max: 0.4,
	}
	chanPerson := make(chan *Person, count)
	wg := sync.WaitGroup{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(num int) {
			defer wg.Done()

			PersonName := fmt.Sprintf("Person-%v", num)
			FatRate := rand.Float64() * (scope.max - scope.min)
			fmt.Println(FatRate)
			p := &Person{}
			p.register(PersonName, FatRate)
			jsonPerson := &PersonJson{
				Name:    PersonName,
				FatRate: FatRate,
			}

			JsonFile(p, jsonPerson)
			chanPerson <- p
		}(i)
	}
	wg.Wait()
	close(chanPerson)
	persons := updateFatRate(chanPerson, scope)

	for index, newPerson := range persons {
		fmt.Println(newPerson.name, ",当前排名第：", index)
	}
	time.Sleep(3 * time.Second)
}

func regJsonToFile(p *Person, person *PersonJson) {
	data, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
	}
	if str, err := os.Getwd(); err == nil {
		path := str
		if err := p.appendPersonIntoFile(data, path+"/homework/five/person.json"); err != nil {
			log.Println("写入文件出错：", err)
		}
	} else {
		log.Fatal("路径出现问题")
	}
}
