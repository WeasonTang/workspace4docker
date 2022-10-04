package main

import (
	"errors"
	"fmt"
	"os"
)

//使用一个结构体管理队列
type Queue struct {
	maxSize int
	array   [6]int
	front   int //指向队首
	rear    int //指向队尾
}

//添加数据队列
func (queue *Queue) AddQueue(val int) (err error) {
	//先判断队列是否已满
	if queue.rear == queue.maxSize-1 {
		return errors.New("the queue is full")
	}
	queue.rear++
	queue.array[queue.rear] = val
	return
}

//从队列取数据
func (queue *Queue) GetQueue() (val int, err error) {
	//先判断队列是否为空
	if queue.front == queue.rear {
		return -1, errors.New("the queue is empty")
	}
	queue.front++
	val = queue.array[queue.front]
	return val, err
}

//显示队列：找到队首遍历到队尾
func (queue *Queue) showQueue() {
	//先判断队列是否为空
	if queue.front == queue.rear {
		fmt.Println("the queue is empty")
		return
	}
	for i := queue.front + 1; i <= queue.rear; i++ {
		fmt.Printf("queue[%d]=%d\t", i, queue.array[i])
	}
	fmt.Println()
}

func main() {
	//创建队列
	queue := &Queue{
		maxSize: 4,
		front:   -1,
		rear:    -1,
	}

	var key string
	var val int
	for {
		fmt.Println("1.输入add 表示添加数据到队列")
		fmt.Println("2.输入get 表示从队列获取数据")
		fmt.Println("3.输入show 表示显式队列")
		fmt.Println("4.输入exit 表示退出队列")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要添加的值")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("add success")
			}
		case "get":
			val, err := queue.GetQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Printf("从队列取出%d\n", val)
			}
		case "show":
			queue.showQueue()
		case "exit":
			os.Exit(0)
		}
	}

}
