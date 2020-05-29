学习笔记

链表翻转定式

```go
for curr != nil{
    tmp = curr.Next
    curr.Next = pre
    pre = curr
    curr = tmp
}
return pre
```
