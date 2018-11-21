package main

import (
    "crypto/sha256"
    "encoding/hex"
    "strconv"
    "strings"
    "fmt"
)

type node struct {
    id string
    objs []string
    num int
    hash uint64
}

var cnts []int

func addToRing (obj string, nodes []node) int {
    size := len(nodes)
    var objHash uint64
    objHash = calcHash(obj)
    fmt.Printf("obj:%d\n", objHash)
    var index int
    var min uint64

    if len(nodes) > 0 {
        min = nodes[0].hash - objHash
    }

    for i := 1; i < size; i++ {
        tmp := nodes[i].hash - objHash
        if  tmp >= 0 && tmp < min{
            min = tmp
            index = i
            //break
        }
    }

    fmt.Printf("%s added to %s\n", obj, nodes[index].id);
    nodes[index].num = cnts[index] + 1
    cnts[index] = cnts[index] + 1
    nodes[index].objs[cnts[index] - 1] = obj

    return 0
}


func calcHash(key string) uint64{
    h := sha256.New()
    h.Write([]byte(key))
    result := h.Sum(nil)
    s := fmt.Sprintf("%s%s","0x", hex.EncodeToString(result))
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
    objs := []string{"obj1", "obj2", "hahah" ,"obj4" ,"obj5", "obj6"}

    //calculate the position of every node
    for i := 0; i < nodeNum; i++ {
        s := []string{"node", strconv.Itoa(i)}
        var newNode node
        newNode.id = strings.Join(s, ":")
        newNode.num = 0
        newNode.hash = calcHash(newNode.id)
        newNode.objs = objs
        nodes = append(nodes, newNode)
    }

    //distribute data
    for i := 0; i < nodeNum; i++ {
        fmt.Printf("node[%d]%d\n", i, nodes[i].hash)
    }

    //fmt.Printf("len: %d\n", len(objs))
    for i := 0; i < len(objs); i++ {
        addToRing(objs[i], nodes)
    }

}
