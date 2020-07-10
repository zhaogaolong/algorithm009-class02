## LRU cache

> https://leetcode-cn.com/problems/lru-cache/

技术点：

- Hash table + 双端链表

时间复杂度：O(1), get、put， update
空间复杂度：O(cap), cap 是 LRU cache 的最大容量

```go
type LRUCache struct {
    data map[int]*Node
    cap int
    head *Node
    tail *Node
}

type Node struct{
    key int
    val int
    pre *Node
    next *Node
}


func Constructor(capacity int) LRUCache {
    cache := LRUCache{
        data: make(map[int]*Node),
        cap: capacity,
        head: &Node{},
        tail: &Node{},
    }
    cache.head.next = cache.tail
    cache.tail.pre = cache.head
    return cache
}


func (this *LRUCache) Get(key int) int {
    node, ok := this.data[key]
    if !ok{
        return -1
    }
    this.remove(node)
    this.putHead(node)
    return node.val
}


func (this *LRUCache) Put(key int, value int)  {
    node, ok := this.data[key]
    if ok{
        node.val = value
        this.remove(node)
        this.putHead(node)
        return
    }
    if len(this.data) >= this.cap{
        delKey := this.tail.pre.key
        this.remove(this.tail.pre)
        delete(this.data, delKey)
    }
    newNode := Node{
        key: key,
        val: value,
    }
    this.data[key] = &newNode
    this.putHead(&newNode)
}

func (this *LRUCache) remove(node *Node)  {
    node.pre.next = node.next
    node.next.pre = node.pre

}
func (this *LRUCache)putHead(node *Node)  {
    old := this.head.next
    this.head.next = node
    node.pre = this.head

    node.next = old
    old.pre = node
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
```
