package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Id     int
	Status bool
}

type waitingPool []Player     // 等待池
type gameThread []Player      // 一条游戏进程
type playerPool [7]gameThread // 游戏池

// 生成100个游戏玩家，id为0-100，状态初始化为false表示不在游戏中
func generate100Players() (pool waitingPool) {
	var p Player
	for i := 0; i < 100; i++ {
		p.Id = i
		p.Status = false
		pool = append(pool, p)
	}
	fmt.Println("Already generate 100 players")
	return
}

// 洗牌算法，模拟Elo匹配
func shuffle(w waitingPool) []Player {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	outOfOrder := make([]Player, len(w))
	for i, randIndex := range r.Perm(len(w)) {
		outOfOrder[i] = w[randIndex]
	}
	return outOfOrder
}

// 删除十个位置，刷新等待池
func refreshWaitingPool(w []Player) (newPool []Player) {
	// newPool := make([]Player, len(w)-10)
	for i := 0; i < len(w); i++ {
		if w[i].Status == false {
			newPool = append(newPool, w[i])
		}
		continue
	}
	return newPool
}

// 匹配一场游戏进程
func match1Game(w waitingPool, g chan<- gameThread, newWaitingPool chan<- waitingPool) {

	outOfOrder := shuffle(w)

	// 匹配成功后状态设为true
	for i := 0; i < 10; i++ {
		outOfOrder[i].Status = true
	}

	// 将游戏中的玩家移出，更新等待池
	newWaitingPool <- refreshWaitingPool(outOfOrder)

	// 取前十个player作为一个游戏进程
	g <- outOfOrder[:10]
	fmt.Println("Now output players:", outOfOrder[:10])

	// 模拟游戏需要2秒
	time.Sleep(time.Second * 2)

}

// func main() {
// 	wP := generate100Players()
// 	g, nwP := match1Game(wP)
// 	fmt.Println("Here is what in gameLine:", len(g))
// 	fmt.Println("and these are waiting players:", len(nwP))
// }

func main() {
	waiting := make(chan waitingPool, 100)
	gaming := make(chan gameThread, 100)

	wP := generate100Players()
	// 游戏池中只开三个游戏进程
	for i := 1; i <= 3; i++ {
		go match1Game(wP, gaming, waiting)
	}

	for a := 1; a <= 9; a++ {

		<-gaming
		<-waiting
	}
	close(gaming)
	close(waiting)
}

// //这个是工作线程，处理具体的业务逻辑，将jobs中的任务取出，处理后将处理结果放置在results中。

// func worker(id int, jobs <-chan int, results chan<- int) {

// 	for j := range jobs {

// 		fmt.Println("worker", id, "processing job", j)
// 		fmt.Println(j * 2)

// 		time.Sleep(time.Second * 10)

// 		results <- j * 2

// 	}

// }

// func main() {

// 	//两个channel，一个用来放置工作项，一个用来存放处理结果。

// 	jobs := make(chan int, 100)

// 	results := make(chan int, 100)

// 	// 开启三个线程，也就是说线程池中只有3个线程，实际情况下，我们可以根据需要动态增加或减少线程。

// 	for w := 1; w <= 4; w++ {

// 		go worker(w, jobs, results)

// 	}

// 	// 添加9个任务后关闭Channel

// 	// channel to indicate that's all the work we have.

// 	for j := 1; j <= 9; j++ {

// 		jobs <- j

// 	}

// 	close(jobs)

// 	//获取所有的处理结果

// 	for a := 1; a <= 9; a++ {

// 		<-results

// 	}
// 	defer close(results)

// }
