//总结：
//Node保存真正的数据
//LRUCache结构设计，快速定位Node的地址
//空节点head,tail的设计，保存在LRUCache struct中

type Node struct {
        key   int
        value int
        prev  *Node
        next  *Node
}

type LRUCache struct {
        Capacity int
        kvMap    map[int]*Node
        head     *Node
        tail     *Node
}

func Constructor(capacity int) LRUCache {
        // if capacity == 0 {
        //     var lru = LRUCache
        //     return lru
        // }

        lru := LRUCache{
                Capacity: capacity,
                kvMap:    make(map[int]*Node),
                head:     new(Node),
                tail:     new(Node),
        }

        lru.head.next = lru.tail
        lru.tail.prev = lru.head
        return lru
}

func (this *LRUCache) Get(key int) int {
        if node, ok := this.kvMap[key]; ok {
                this.moveToHead(node)
                return node.value
        }
        return -1
}

func (this *LRUCache) Put(key int, value int) {
        if node, ok := this.kvMap[key]; ok {
                node.value = value
                this.moveToHead(node)
        } else {
                node = new(Node)
                node.key = key
                node.value = value
                this.addNode(node)

                if this.Capacity < len(this.kvMap) {
                        this.popTail()
                }

        }

}

func (this *LRUCache) popTail() {
        this.removeNode(this.tail.prev)
}

func (this *LRUCache) moveToHead(node *Node) {
        this.removeNode(node)
        this.addNode(node)
}

func (this *LRUCache) removeNode(node *Node) {
        delete(this.kvMap, node.key)

        prev := node.prev
        next := node.next

        prev.next = next
        next.prev = prev
}

func (this *LRUCache) addNode(node *Node) {
        node.prev = this.head
        node.next = this.head.next

        this.head.next.prev = node
        this.head.next = node

        this.kvMap[node.key] = node
}


func main() {
        lru := Constructor(2)
        p1 := lru.Get(2)
        lru.Put(1, 1)
        lru.Put(2, 2)
        p2 := lru.Get(1)
        lru.Put(3, 3)
        p3 := lru.Get(2)

        fmt.Println(lru)
        //for k, v := range lru.kvMap {
        //      fmt.Println(k, v.value)
        //}
        fmt.Println(p1, p2, p3)

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */