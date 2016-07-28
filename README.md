## Kubernetes RESTful APIs test

## Plan:

1. What is RESTful API?
   RESTful API是一种通信架构，全称是“表现层状态转移”，在2000年由HTTP规范的主要编写人Roy Thomas Fielding提出。
   RESTful API的优点是结构清晰，易于扩展。有三个特征：
   
   - 一个URI代表一个资源
   - client和server之间传递它的某种表现层
   - client通过HTTP的动词（GET、POST、PUT、DELETE）来完成对资源的操作

2. How does Kubernetes use their restful apis?
   通过它自有的组件kube-apiserver提供API服务，它运行在master节点上，包含：
   
   - 端口8080：HTTP请求，非认证
   - 端口6443: HTTPS请求，采用基于Token或者证书等的安全认证

   除此之外，Kubernetes还可以通过代理方式访问API：

```shell
   kubectl proxy --port=8081 &
```

3. How to write a golang program to do this?
   在上面的例子rest-api中，实现了对Kubernetes RESTful API的访问及包裹，这种实现更像是"api-proxy"的方法：
   
   client ---> api-proxy ---> k8s api
   client      api-proxy <--- k8s api
   client <--- api-proxy      k8s api

   这种方式存在两个缺陷：
   - 相比直接访问k8s api中间加了一层肯定是慢了，api-proxy收到用户的1个请求后可能需要对k8s api发送多个请求，并将请求结果进行分析和重组再返回给用户。解决这个问题的思路是：给api-proxy添加缓存
   - 每次添加新的资源和新的路由规则，都需要重写程序。这明显是低效的，应该找个地方可以注册资源和路由规则，这样就无需改动服务器程序了。解决思路：将资源和对应的路由规则存储在键值对存储中，服务器每次调用去这个存储里查找相应的规则。

## EXTENSION:

1. How does a ccloud looks like? UI? Functions? Arch? and so on.
   这个问题应该分为下面两个问题：
   **为什么要有容器云？**
  
   容器云相比现有的虚机云，它的优势主要来自容器的优势。不管是Docker容器，还是Rkt等其他容器，它们都带来了轻量、快速、规范的使用方法。这一切都指向一个目的：解放生产力，提高生产效率。

   sxxx
   
   推广新技术并不是件容易的事，在电子计算器刚出现时，也有不少人对它的准确性和可靠性持怀疑态度，对于这种看似变戏法的小物件拒不接受。它的推广者说这个小玩意儿计算非常迅速，精度也高，但固执者们仍只愿相信自己手动算数的结果。慢慢地，人们发现对于大数据量的计算，手工计算已经无法满足效率的要求，电子计算器终于走上了舞台。电子计算器的优势一直在那里，只是缺少一个展现自我的机会。同样，对组织内部的旧程序、旧系统，无论是微服务式改造，还是单纯更迭为Docker容器，也遭受着难以想象的巨大阻力。除开利益的纠纷，一种新技术推广甚难，我猜情形应该是这样的：旧系统运行良好，没有太多无法解决的通电。


   **如果有，它应该是什么样子？界面？结构？技术栈？**

