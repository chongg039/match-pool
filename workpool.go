package main

import (
	"log"
	"sync"
	"time"
)

// 玩家
type Player struct {
	ID     int
	Status bool
}

// 玩家池
type PlayerPool struct {
	Pool []Player
}

type GameLine []Player

// 工人线程
type Worker struct {
}

// func refreshWaitingPool(w []Player) (newPool []Player) {
// 	// newPool := make([]Player, len(w)-10)
// 	for i := 0; i < len(w); i++ {
// 		if w[i].Status == false {
// 			newPool = append(newPool, w[i])
// 		}
// 		continue
// 	}
// 	return newPool
// }

// func refreshPool(pool PlayerPool, player <-chan Player) {

// }

func work(wg *sync.WaitGroup, players <-chan Player, game chan<- GameLine) {
	defer wg.Done()
	counter := 1
	var onegame GameLine
	for p := range players {
		if counter <= 10 {
			onegame = append(onegame, p)
			counter++
			continue
		} else {
			log.Println("Now create one game:", onegame)
			game <- onegame
		}
		time.Sleep(time.Second * 4)
	}
}

func main() {
	players := make(chan Player, 100)
	games := make(chan GameLine, 100)

	var wg sync.WaitGroup

	for i := 0; i < 1; i++ {
		go work(&wg, players, games)
	}

	// 初始化生成100个玩家并放入players
	var p Player
	for i := 0; i < 100; i++ {
		p.ID = i
		p.Status = false
		players <- p
		wg.Add(1)
	}
	close(players)

	for i := 0; i < 2; i++ {
		<-games
	}

	wg.Wait()
}
