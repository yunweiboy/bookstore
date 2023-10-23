#### 项目实现逻辑图
![bookstore项目实现逻辑图.png](png%2Fbookstore%E9%A1%B9%E7%9B%AE%E5%AE%9E%E7%8E%B0%E9%80%BB%E8%BE%91%E5%9B%BE.png)
#### 模拟客户端请求
##### 创建图书条目示例
    curl  -XPOST -H "Content-Type:application/json" -d '{"id": "978-7-111-55842", "name": "golang.txt", "authors":["tom"], "press":"nanjing published"}' localhost:8080/book
##### 更新图书条目
    curl  -XPOST -H "Content-Type:application/json" localhost:8080/book/978-7-111-55842
##### 获取某条图书条目
    curl  -XGET -H "Content-Type:application/json" localhost:8080/book/978-7-111-55842
##### 获取所有图书条目
    curl  -XGET -H "Content-Type:application/json" localhost:8080/book
##### 删除某条图书条目
    curl  -XDELETE -H "Content-Type:application/json" localhost:8080/book/978-7-111-55842
