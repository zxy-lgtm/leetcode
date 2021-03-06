## Kubernetes对象

### 对象spec

* 具有spec的对象必须设置内容：期望状态（*Desired State*）
* 期望状态：kubernetes集群工作负载的理想状态

### 对象status

* 描述对象*当前状态*
* 由kubernetes系统和组件设置和更新
* 控制平面（control plane）一直主动管理状态趋向理想状态

### 描述kubernetes对象

* 在建立对象时，要提供spec去描述期望状态，以及一些基本信息
* 使用[API](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md)创建对象时，需要提供json格式的信息
* 但是大部分时候只需要.yaml文件，它会自己转换
* 一些必须信息

```yaml
apiVersion	-必需字段，API版本
kind		-必需字段，创建的对象类型
metadata	-必需字段，唯一性标识对象的一些数据
	name	-metadata中的数据......
spec		-必需字段，期望状态......
```

* 一个例子

  ```yaml
  apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: nginx-deployment
  spec:
    selector: #核心分组语句，通过这个客户端/用户可以识别一组对象
      matchLabels:
        app: nginx
    replicas: 2 # tells deployment to run 2 pods matching the template
    template:
      metadata:
        labels:
          app: nginx
      spec:
        containers:
        - name: nginx
          image: nginx:1.14.2
          ports:
          - containerPort: 80
  ```

  * selector
    * 定义Deployment如何查找要管理的Pods
    * 基于等值需求：
      * `=`、`==`、`!=`
      * 例如Pod指定节点选择标准
    * 基于集合需求：
      * `in`、`notin` 和 `exists`
  * replicas
    * replicaSet：副本控制器，在这其中负责启动两个`nginx` Pods
  * template
    * Pods的`metadata`
    * Pods的`spec`

### Kubernetes对象管理

* tips：[应该只使用一种技术来管理 Kubernetes 对象。混合和匹配技术作用在同一对象上将导致未定义行为。](https://kubernetes.io/zh/docs/concepts/overview/working-with-objects/object-management/#%E6%8C%87%E4%BB%A4%E5%BC%8F%E5%91%BD%E4%BB%A4)

* 指令式命令

  ```shell
  kubectl create deployment ......
  ```

* 指令式对象配置

  * 需要编写yaml文件

  ```shell
  kubectl create -f file.yaml
  kubectl delete -f file.yaml
  kubectl replace -f file.yaml
  ```

  * tips：`replace` 指令式命令将现有规范替换为新提供的规范，并放弃对配置文件中 **缺少的对象**的所有更改。

* 声明式对象配置

  ```shell
  kubectl diff -f dir/   #首先查看将要进行的修改
  kubectl apply -f dir/
  
  递归处理
  kubectl diff -R -f dir/   
  kubectl apply -R -f dir/
  ```

* 各有优缺点

### name 和 ID

* 每一个对象都有一个name：标识唯一性
* 每个Kubernetes对象都有都有一个ID和一个name标识唯一性
* Pod和Deployment同名可以接受
* 非唯一属性
  * labels
  * annotation

#### name

* 引用资源url中的对象
* 常用四种资源命名约束
  * DNS子域名
  * RFC1123标签名
  * RFC1035标签名
  * 路径分段名称
    * 某些资源类型要求名称能被安全地用作路径中的片段。 其名称不能是 `.`、`..`，也不可以包含 `/` 或 `%` 这些字符

#### UIDs

* Kubernetes系统生成，全局唯一

### namespace

* 提供一种机制，将同一集群中的资源划分为相互隔离的组。

* 仅对带有namespace的对象适用

* 创建和删除：[名字空间的管理指南文档](https://kubernetes.io/zh/docs/tasks/administer-cluster/namespaces/)

  * 使用ns.yaml文建
  * [建立新的名字空间](https://kubernetes.io/zh/docs/tasks/administer-cluster/namespaces/#creating-a-new-namespace)
  * [删除名字空间](https://kubernetes.io/zh/docs/tasks/administer-cluster/namespaces/#deleting-a-namespace)

* 查看

  ```
  kubectl get namespace
  ```

* Kubernetes会创建四个初始名字空间

  * default - 没有指明namespace对象的默认空间
  * kube-system - Kubernetes系统创建对象使用的namespace
  * kube-public - 自动创建的，主要用于集群，用于一些集群共享资源
  * kube-node-lease -  此名字空间用于与各个节点相关的 [租约（Lease）](https://kubernetes.io/docs/reference/kubernetes-api/cluster-resources/lease-v1/)对象。 节点租期允许 kubelet 发送[心跳](https://kubernetes.io/zh/docs/concepts/architecture/nodes/#heartbeats)（heartbeats），由此控制平面（control plane）能够检测到节点故障。

* 为请求设置命名空间

  ```shell
  指令 --namespace // 本次指令生效
  kubectl config set-context --current --namespace=<名字空间名称> // 一直生效
  ```

#### namespace and DNS

* 当一个服务被创建，Kubernetes会创建一个形式为 `<服务名称>.<名字空间名称>.svc.cluster.local`的DNS条目

#### 许多对象类型并没有在namespace中

* node

* persistentVolumes

* namespace 

* 查看哪些 Kubernetes 资源在名字空间中，哪些不在名字空间中：

  ```shell
  # 位于名字空间中的资源
  kubectl api-resources --namespaced=true
  
  # 不在名字空间中的资源
  kubectl api-resources --namespaced=false
  ```

#### control plane Automatic label

### Labels and Selectors

* *标签（Labels）* 是附加到 Kubernetes 对象（比如 Pods）上的**键值对**
* 标签能够支持高效的查询和监听操作
* selector前面

#### API 

* LIST 、 WATCH 过滤

  * 同样有基于等值和集合
  * 写法也相同

* 设置引用

  *  [`services`](https://kubernetes.io/zh/docs/concepts/services-networking/service/) 和 [`replicationcontrollers`](https://kubernetes.io/zh/docs/concepts/workloads/controllers/replicationcontroller/)、[`pods`](https://kubernetes.io/zh/docs/concepts/workloads/pods/)
  * 对象的标签在文件中都是使用映射定义，只基于等值需求

* 支持基于集合需求的资源

  *  [`Job`](https://kubernetes.io/zh/docs/concepts/workloads/controllers/job/)、 [`Deployment`](https://kubernetes.io/zh/docs/concepts/workloads/controllers/deployment/)、 [`Replica Set`](https://kubernetes.io/zh/docs/concepts/workloads/controllers/replicaset/) 和 [`DaemonSet`](https://kubernetes.io/zh/docs/concepts/workloads/controllers/daemonset/)

  * 一个例子：

    ```
    selector:
      matchLabels:
        component: redis
      matchExpressions:
        - {key: tier, operator: In, values: [cache]}
        - {key: environment, operator: NotIn, values: [dev]}
    ```

  *  有效的运算符包括 `In`、`NotIn`、`Exists` 和 `DoesNotExist`。

### 注解（Annotations）

* 为对象附加元数据
  * 键值对
  * **Note:**Map 中的键和值必须是字符串。 换句话说，你不能使用数字、布尔值、列表或其他类型的键或值。
* 可以由注解来记录的一些例子
  * 声明性配置所管理的字段
  * 构建、发布或着镜像信息
  * 日志记录、监控、分析或者审计仓库的指针
  * 调试目的的客户端或者工具信息
  * 指令
  * 一些工具的数据信息

### Finalizers

* 带有namespace的键
* 提醒控制器清理被删除的对象拥有的资源
* 可用于垃圾回收
* 工作原理
  * 在 `metadata.finalizers`中定义值
  * 在删除资源是，服务器会注意到该字段的值
    * 将你开始执行删除的时间添加到 `metadata.deletionTimestamp` 字段。
    * 在`metadata.finalizers`为空之前禁止删除
    * 返回202状态码
    * 控制器会在满足finalier之后删除此键
    * `metadata.finalizers`为空时自动删除deletiontimestamp字段
* Tips:避免手动移除finalizer

### 字段选择器

* 根据字段筛选出对应的Pod
* 支持字段
  * name
  * namespace
  * 其他随着资源类型不同而变化
* 支持操作符：`=`、`==` 和 `!=` （`=` 和 `==` 的意义是相同的）
* 链式选择器，用`，`分隔
* 跨资源类型选择

### [常用标签与例子](https://kubernetes.io/zh/docs/concepts/overview/working-with-objects/common-labels/)

## 架构

### 节点（Node）

* Kubernetes 通过将容器放入在节点（Node）上运行的 Pod 中来执行工作负载。
* 由控制面管理

### RuntimeClass

* 运行多容器运行时还需要解决以下几个问题：
  - 集群里有哪些可用的容器运行时？
  - 如何为 Pod 选择合适的容器运行时？
  - 如何让 Pod 调度到装有指定容器运行时的节点上？
  - 容器运行时在运行容器时会产生有一些业务运行以外的额外开销，这种「额外开销」需要怎么统计？
* 为了解决这些问题推出了RuntimeClass
* 使用
  1. 配置RuntimeClass
  2. 在Pod spec中指定runtimeClassName来使用
  3. 如果未指定则使用默认的RuntimeHandler
* CRI配置（ Container Runtime Interface）
  * 实现了运行时和 Kubernetes 的解耦
  * cintainerd  在`/etc/containerd/config.toml` 配置文件来配置，handler 需要配置在 runtimes 块中
  * cri-o  在`/etc/crio/crio.conf` 配置文件来配置， handler 需要配置在 crio.runtime 表下面
* 调度 
  * scheduling字段，需要设置挑选器，运行时取交集
  * 阻止某些需要特定 RuntimeClass 的 pod，可以在 `tolerations` 中指定
* 回调
  * PostStart
  * PreStop

