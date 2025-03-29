package main // race condition состояние гонки
// если ввести go run -race main.go , то го нам выведет горутины в состоянии гонки
import (
	"fmt"
	"sync"
	"time"
)

type counter struct {
	count int
	mutex *sync.Mutex // позволяет блокировать/разблокировать и так избегать datarace
}

func (c *counter) inc() {
	c.mutex.Lock() // вызов этой команды блокирует переменную за какой-либо горутиной
	defer c.mutex.Unlock()
	c.count++
}

func (c *counter) value() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.count
}

func main() {
	c := counter{
		mutex: new(sync.Mutex),
	}

	for i := 0; i < 1000; i++ {
		go func() {
			c.inc()
		}()
	}

	time.Sleep(time.Second)

	fmt.Println(c.value())
}
