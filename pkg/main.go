/**
    @Author: qiyou_wu
    @CreateDate: 2022/4/17
    @Description:
**/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Restrictedflag bool //限流开关标记

func main() {

	Restrictedflag = false
	head := initLinkList()

	go sendReq()

	go setFlowRate(head)

	time.Sleep(60 * time.Second)
}

func sendReq() {
	for {
		if !Restrictedflag {
			fmt.Println("发送成功")
		} else {
			r := rand.Intn(100)
			if r < 70 {
				fmt.Println("发送失败")
			}
			fmt.Println("发送成功")
		}
		time.Sleep(1 * time.Second)
	}
}

//模拟设置每秒的访问数量，用0-100之间的随机数表示
func setFlowRate(Head *Node) {
	node := Head
	for {
		//获取当前流量
		r := rand.Intn(100)
		fmt.Println(fmt.Sprintf("此刻的流量为%d", r))
		//将流量设置到滑动窗口
		node.Value = r
		//统计历史数据，由于算上历史数据触发限流有点苛刻，暂时这里不使用
		_, flag := countHistory(node)
		node = node.Next
		//如果是刚开始的几秒中历史数据不具备意义
		if !flag {
			time.Sleep(1 * time.Second)
			continue
		}
		if r > 80 && Restrictedflag == false {
			fmt.Println("开始限流")
			Restrictedflag = true
		}
		time.Sleep(1 * time.Second)
	}
}

//统计前三秒的历史数据
func countHistory(node *Node) (int, bool) {
	res := 0
	for i := 0; i < 3; i++ {
		node = node.Pre
		if node.Value < 0 {
			return -1, false
		}
		res = res + node.Value
	}
	return res, true
}

//初始化滑动窗口（用环形链表）
func initLinkList() *Node {
	head := new(Node)
	head.Value = 1
	cash := head
	for i := 0; i < 10; i++ {
		node := new(Node)
		node.Value = -1
		head.Next = node
		node.Pre = head
		head = node
		if i == 9 {
			head.Next = cash
			cash.Pre = head
		}
	}
	return cash
}
