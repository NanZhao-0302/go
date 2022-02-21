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

const totalRegistrants int = 1000

type Persons struct {
	name       string
	fatRate    float64
	curFatRate float64
}

type FatRateRankMachine struct {
	score []Persons
	regs  map[string]Persons
}

func fatRateMachineOperator(in <-chan func(*FatRateRankMachine), done <-chan struct{}) {
	machine := &FatRateRankMachine{
		score: make([]Persons, 0, totalRegistrants),
		regs:  map[string]Persons{},
	}

	for {
		select {
		case <-done:
			return
		case f := <-in:
			f(machine)
		}
	}
}

type ChannelFatRateRankMachine chan func(*FatRateRankMachine)

func NewChannelFatRateRankMachine() (ChannelFatRateRankMachine, func()) {
	ch := make(ChannelFatRateRankMachine, 1000)
	done := make(chan struct{})
	go fatRateMachineOperator(ch, done)
	return ch, func() {
		close(done)
	}
}

func (c ChannelFatRateRankMachine) getRegistrants() map[string]Persons {
	var registrants map[string]Persons
	done := make(chan struct{})
	c <- func(machine *FatRateRankMachine) {
		registrants = machine.regs
		close(done)
	}
	<-done
	return registrants
}

func (c ChannelFatRateRankMachine) UpdateFatrate(p Persons) (int, bool) {
	var val int
	var ok1 bool
	done := make(chan struct{})
	c <- func(machine *FatRateRankMachine) {
		defer close(done)
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
					val = i2 + 1
					ok1 = true
					return
				}
			}

			machine.score = append(machine.score, p)
			val = len(machine.score)
			ok1 = true
			return

		} else {
			val = 0
			ok1 = false
			return
		}
	}
	<-done
	return val, ok1

}

func (c ChannelFatRateRankMachine) getRank(p Persons) (int, bool) {
	var val int
	var ok bool
	done := make(chan struct{})

	c <- func(machine *FatRateRankMachine) {
		defer close(done)
		for i, person := range machine.score {
			if person.name == p.name {
				val = i + 1
				ok = true
				return
			}
		}
		val = 0
		ok = false
	}

	<-done

	return val, ok
}

func (c ChannelFatRateRankMachine) register(p Persons) {

	c <- func(machine *FatRateRankMachine) {
		_, ok := machine.regs[p.name]
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

}

func insert(a []Persons, p Persons, i int) []Persons {

	return append(a[:i], append([]Persons{p}, a[i:]...)...)
}

func (c ChannelFatRateRankMachine) PrintScoreboard() {
	done := make(chan struct{})
	c <- func(machine *FatRateRankMachine) {
		for i, person := range machine.score {
			fmt.Println(i+1, ": ", person.name, " ", person.curFatRate)
		}
		close(done)
	}
	<-done

}

func remove(a []Persons, i int) []Persons {
	return append(a[:i], a[i+1:]...)
}

func randFloat(min, max float64) float64 {
	if min < 0 {
		min = 0.0
	}
	return min + rand.Float64()*(max-min)
}

func main() {

	//定义一个通道
	signaldone := make(chan os.Signal, 1)
	signal.Notify(signaldone, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	manager, finish := NewChannelFatRateRankMachine()
	rand.Seed(time.Now().Unix())

	var wg sync.WaitGroup
	wg.Add(totalRegistrants)

	//注册
	for i := 0; i < totalRegistrants; i++ {
		go func(i int) {
			defer wg.Done()
			name := fmt.Sprintf("Persons%d", i)
			base := randFloat(0, 0.4)
			manager.register(Persons{name: name, fatRate: base, curFatRate: base})
		}(i)
	}
	wg.Wait()
	manager.PrintScoreboard()
	//结束注册

	reqCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		fmt.Println("[Main]: canceling context")
		cancel()
	}()

	registrants := manager.getRegistrants()

	for _, person := range registrants {
		wg.Add(2)
		go func(person Persons) {
			defer wg.Done()

			for {
				select {
				case <-reqCtx.Done():
					fmt.Println("[", person.name, "]", "[Update FatRate]", "Timeout! Exiting")
					return
				default:
					minFatRate := person.fatRate - 0.2
					person.curFatRate = randFloat(minFatRate, person.fatRate+0.2)
					rank, ok := manager.UpdateFatrate(person)
					if ok {
						fmt.Println("[", person.name, "]", "[Update FatRate]", ",", rank, ",", person.curFatRate, " ", time.Now())
					} else {
						fmt.Println("[", person.name, "]", "[Update FatRate]", "  not found!")
					}

				}

			}
		}(person)

		go func(person Persons) {
			defer wg.Done()
			for {
				select {
				case <-reqCtx.Done():
					fmt.Println("[", person.name, "]", "[Query Rank]", "Timeout! Exiting")
					return
				default:
					rank, ok := manager.getRank(person)
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
	case <-signaldone:
		fmt.Println("收到")
		cancel()
		wg.Wait()

	case <-reqCtx.Done():
		fmt.Println("超时")
		wg.Wait()

	}

	manager.PrintScoreboard()
	finish()

	fmt.Println("程序完成!")
}
