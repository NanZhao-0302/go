package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Person struct {
	name       string
	fatRate    float64
	curFatRate float64
}

type MutexFatRateRankMachine struct {
	score []Person
	regs  map[string]Person
	loc   sync.RWMutex
}

func NewMutexFatRateRankMachine(maxRegistrant int) *MutexFatRateRankMachine {
	return &MutexFatRateRankMachine{
		score: make([]Person, 0, maxRegistrant),
		regs:  map[string]Person{},
	}
}

func (machine *MutexFatRateRankMachine) register(p Person) {
	machine.loc.Lock()
	defer machine.loc.Unlock()

	_, ok := machine.regs[p.name]
	//if it does not exist in score, then register as new user
	if ok == false {
		for i, pl := range machine.score {
			if p.curFatRate <= pl.curFatRate {
				machine.score = insert(machine.score, p, i)
				machine.regs[p.name] = p
				return
			}
		}
		machine.score = append(machine.score, p)
		machine.regs[p.name] = p
	}
}
func (machine *MutexFatRateRankMachine) PrintScoreboardWithLock() {
	machine.loc.RLock()
	defer machine.loc.RUnlock()
	machine.PrintScoreboard()
}

func insert(a []Person, p Person, i int) []Person {

	return append(a[:i], append([]Person{p}, a[i:]...)...)
}

func (machine *MutexFatRateRankMachine) PrintScoreboard() {
	for i, person := range machine.score {
		fmt.Println(i+1, ": ", person.name, " ", person.curFatRate)
	}
}

func (machine *MutexFatRateRankMachine) getRank(name string) (int, bool) {
	machine.loc.RLock()
	defer machine.loc.RUnlock()

	for i, person := range machine.score {
		if person.name == name {
			return i + 1, true
		}
	}
	return 0, false
}

func (machine *MutexFatRateRankMachine) updateFatRate(p Person) (int, bool) {
	machine.loc.Lock()
	defer machine.loc.Unlock()

	_, ok := machine.regs[p.name]
	if ok == true {
		for i, pl := range machine.score {
			if p.name == pl.name {
				machine.score = remove(machine.score, i)
				break
			}
		}

		for i2, pl2 := range machine.score {
			if p.curFatRate <= pl2.curFatRate {
				machine.score = insert(machine.score, p, i2)
				return i2 + 1, true
			}
		}
		machine.score = append(machine.score, p)
		return len(machine.score), true

	} else {
		return 0, false
	}
}

func remove(a []Person, i int) []Person {
	return append(a[:i], a[i+1:]...)
}

func randFloat(min, max float64) float64 {
	if min < 0 {
		min = 0.0
	}
	return min + rand.Float64()*(max-min)
}

func main() {
	totalRegistrants := 1000

	//定义一个通道来捕获 Control C 以便在终止之前让所有程序执行
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	manager := NewMutexFatRateRankMachine(totalRegistrants)
	rand.Seed(time.Now().Unix())

	var wg sync.WaitGroup
	wg.Add(totalRegistrants)

	//为注册人注册
	for i := 0; i < totalRegistrants; i++ {
		go func(i int) {
			defer wg.Done()
			name := fmt.Sprintf("Person%d", i)
			base := randFloat(0, 0.4)
			manager.register(Person{name: name, fatRate: base, curFatRate: base})
		}(i)
	}
	wg.Wait()
	manager.PrintScoreboardWithLock()
	//注册完成

	//模拟超时情况
	reqCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer func() {
		fmt.Println("[Main]: canceling context")
		cancel()
	}()

	for _, person := range manager.regs {
		wg.Add(2)
		go func(person Person) {
			defer wg.Done()

			//Loop:
			for {
				select {
				case <-reqCtx.Done():
					fmt.Println("[", person.name, "]", "[Update FatRate]", "Timeout! Exiting")
					return
				default:
					//update
					minFatRate := person.fatRate - 0.2
					person.curFatRate = randFloat(minFatRate, person.fatRate+0.2)
					rank, ok := manager.updateFatRate(person)
					if ok {
						fmt.Println("[", person.name, "]", "[Update FatRate]", ",", rank, ",", person.curFatRate, " ", time.Now())
					} else {
						fmt.Println("[", person.name, "]", "[Update FatRate]", "  not found!")
					}

				}

			}
		}(person)

		go func(person Person) {
			defer wg.Done()

			for {
				select {
				case <-reqCtx.Done():
					fmt.Println("[", person.name, "]", "[Query Rank]", "Timeout! Exiting")
					return
				default:
					rank, ok := manager.getRank(person.name)
					if ok {
						fmt.Println("[", person.name, "]", "[Query Rank]", rank, ",", person.curFatRate, " ", time.Now())
					} else {
						fmt.Println("[", person.name, "]", "[Query Rank]", "  not found!")
					}

				}

			}
		}(person)

	}

	select {
	case <-done:
		fmt.Println("收到")
		cancel()
		wg.Wait()
	case <-reqCtx.Done():
		fmt.Println("超时")
		wg.Wait()
	}
	manager.PrintScoreboardWithLock()

	fmt.Println("程序结束啦!")
}
