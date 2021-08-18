![img](https://geekr.gstatics.cn/wp-content/uploads/2021/03/1_gYcSAua4-0jX5b2h-ocxLA.jpeg)

### 链表

链表是一种数据结构，和数组不同，链表并不需要一块连续的内存空间，它通过「指针」将一组零散的内存块串联起来使用，如图所示：

![数组和链表的内存分布](https://laravelacademy.org/storage/uploads/images/gallery/2019-10/scaled-1680-/FrpTrH8Z5fVbZ6lXAIg1b2qNAojU.png)

#### 单链表

链表有多种类型，最简单的是单链表，单链表是最原生的链表，其结构如图所示：

![单链表](https://laravelacademy.org/storage/uploads/images/gallery/2019-10/scaled-1680-/Fmib5GjUDyM9HmZNFA-Swn_Gaeoe.png)

单链表中有两个节点比较特殊，分别是第一个节点和最后一个节点。我们通常把第一个节点叫作头节点，把最后一个结点叫作尾节点。

其中，<font color='red'>头节点用来记录链表的基地址</font>，有了它，我们就可以遍历得到整条链表。而尾节点特殊的地方是：指针不是指向下一个结点，而是指向一个空地址 NULL，表示这是链表上最后一个节点。

对于其他普通节点而言，每个节点至少使用两个内存空间：一个用于存储实际数据，另一个用于存储下一个元素的指针，从而形成出一个节点序列，构建链表。

对单链表而言，理论上来说，插入和删除节点的时间复杂度是 O(1)，查询节点的时间复杂度是 O(n)。

#### 双向链表

![双向链表](https://laravelacademy.org/storage/uploads/images/gallery/2019-10/scaled-1680-/FhRu8_YO4DQWaDv5OlqNLp8RW6B7.png)

#### 循环链表

![循环链表](https://laravelacademy.org/storage/uploads/images/gallery/2019-10/scaled-1680-/Fuicu0DP-ePWw08V7CRCFLwr2UF4.png)

##### 约瑟夫环问题

设编号为1, 2, ...n 的n个人围坐一圈，约定编号为k (1<=k<=n)的人从1开始报数，数到m的那个人出列，它的下一位又从1开始报数，数到m的那个人又出列，以此类推，直到所有人出列为止，由此产生一个出队编号的序列。

思路：

根据用户的输入，生成一个小孩出圈的顺序

n = 5, 即有5个人

k = 1, 从第一个人开始报数

m = 2, 数2下

- 创建一个辅助指针tail，事先指向环形链表的最后一个节点

- 小孩报数时，先让first和tail指针移动k-1次

- 当小孩报数时，让first和tail指针同事移动m-1次

- 这是就可以将first指向的小孩节点出圈

  
