/**
    @Author: qiyou_wu
    @CreateDate: 2022/4/17
    @Description:
**/
package main

type Node struct {
	Value int //访问的请求数量

	Next *Node //指向下一个结点

	Pre *Node //指向上一个节点
}
