package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

type Ring struct {
	nodes  map[uint64]*Node
	vnodes map[uint64]*VirtualNode
}

type Node struct {
	id   string
	hash uint64
}

type VirtualNode struct {
	id   string
	hash uint64
}

func (r *Ring) InitRing() {

}

func (r *Ring) AddNode(key uint64) bool {
	node := &Node{
		id:   key,
		hash: calcHash(key),
	}
	r.nodes[node.hash] = node
	v0 := &VirtualNode{
		id:   key + "0",
		hash: calcHash(key + "0"),
	}
	r.vnodes[v0.hash] = v0

	v1 := &VirtualNode{
		id:   key + "1",
		hash: calcHash(key + "1"),
	}
	r.vnodes[v1.hash] = v1

	v2 := &VirtualNode{
		id:   key + "2",
		hash: calcHash(key + "2"),
	}
	r.vnodes[v2.hash] = v2

	v3 := &VirtualNode{
		id:   key + "3",
		hash: calcHash(key + "3"),
	}
	r.vnodes[v3.hash] = v3
}

func (r *Ring) RemoveNode(hash uint64) bool {
	delete(r.nodes, hash)

}

func (r *Ring) AddObject(obj string, nodes []node, virtualNodes []virtualNode) int {
	size := len(nodes)
	size1 := len(virtualNodes)

	var objHash uint64
	objHash = calcHash(obj)
	fmt.Printf("obj:%d\n", objHash)
	var virIndex int
	var actualIndex int
	var min uint64

	if len(nodes) > 0 {
		min = virtualNodes[0].hash - objHash
	}

	for i := 1; i < size1; i++ {
		tmp := virtualNodes[i].hash - objHash
		if tmp >= 0 && tmp < min {
			min = tmp
			virIndex = i
			//break
		}
	}

	if len(nodes) > 0 {
		min = nodes[0].hash - objHash
	}

	for i := 1; i < size; i++ {
		tmp := nodes[i].hash - virtualNodes[virIndex].hash
		if tmp >= 0 && tmp < min {
			min = tmp
			actualIndex = i
			//break
		}
	}

	fmt.Printf("%s added to %s\n", obj, nodes[actualIndex].id)
	nodes[actualIndex].num = cnts[actualIndex] + 1
	cnts[actualIndex] = cnts[actualIndex] + 1
	nodes[actualIndex].objs[cnts[actualIndex]-1] = obj

	return 0
}

func (r *Ring) RemoveObject() {

}

func (r *Ring) GetObject(key uint64) string {

}

func calcHash(key string) uint64 {
	h := sha256.New()
	h.Write([]byte(key))
	result := h.Sum(nil)
	s := fmt.Sprintf("%s%s", "0x", hex.EncodeToString(result))
	s1 := s[0:18]
	// fmt.Println(s1);
	var r uint64

	if res, err := strconv.ParseUint(s1, 0, 64); err == nil {
		//fmt.Println(res);
		r = uint64(res)
	}
	r1 := r % 4294967295
	//fmt.Printf("hash: %d\n", r1);
	return r1
}

func reDistribute(nodes []node) {

}

func main() {
	nodeNum := 4
	for i := 0; i < nodeNum; i++ {
		cnts = append(cnts, 0)
	}

	var nodes []node
	objs := []string{"obj1", "obj2", "hahah", "obj4", "obj5", "obj6"}

	//calculate the position of every node
	for i := 0; i < nodeNum; i++ {
		s := []string{"node", strconv.Itoa(i)}
		var newNode node
		newNode.id = strings.Join(s, "")
		newNode.num = 0
		newNode.hash = calcHash(newNode.id)
		newNode.objs = objs
		nodes = append(nodes, newNode)
	}

	var virtualNodes []virtualNode

	for i := 0; i < len(nodes); i++ {
		for j := 0; j < 4; j++ {
			s := []string{nodes[i].id, "#", strconv.Itoa(j)}
			ss := strings.Join(s, "")
			fmt.Printf("virname:%s\n", ss)
			var newVir virtualNode
			newVir.id = ss
			newVir.hash = calcHash(newVir.id)
			virtualNodes = append(virtualNodes, newVir)
		}
	}

	//distribute data
	for i := 0; i < nodeNum; i++ {
		fmt.Printf("node[%d]%d\n", i, nodes[i].hash)
	}

	//fmt.Printf("len: %d\n", len(objs))
	for i := 0; i < len(objs); i++ {
		addToRing(objs[i], nodes, virtualNodes)
	}

}
