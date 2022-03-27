# Docker

## 历史起源

* 底层技术上，dotCloud 平台利用了 Linux 容器技术。为了方便创建和管理这些容器，dotCloud 开发了一套内部工具，之后被命名为“Docker”。Docker就是这样诞生的

## Docker 

* Docker 引擎是用于运行和编排容器的基础设施工具。
* “Docker”一词也会用于指代开源 Docker 项目。其中包含一系列可以从 Docker 官网下载和安装的工具，比如 Docker 服务端和 Docker 客户端。
* 不过，该项目在 2017 年于 Austin 举办的 DockerCon 上正式命名为 Moby 项目。[moby/moby: Moby Project - a collaborative project for the container ecosystem to assemble container-based systems (github.com)](https://github.com/moby/moby)

## Docker安装

* windows

  * Windows 10 64位，Hyper-V和容器特性[在 Windows 10 上启用 Hyper-V | Microsoft Docs](https://docs.microsoft.com/zh-cn/virtualization/hyper-v-on-windows/quick-start/enable-hyper-v)

  * 想要运行Windows容器

    ```
    C:\Program Files\Docker\Docker> .\dockercli -SwitchDaemon
    ```

* Ubuntu 

  * ```
    wget -qO- https://get.docker.com/ | sh
    ```

## 核心技术

* 用途：通过镜像提供软件运行需要的条件，包括代码、运行时环境、系统工具、系统库、设置等，在镜像上拉取容器，保障测试环境和开发环境的一致性，减少因环境导致的运行软件冲突
* 使用：研发人员在将代码开发完成后，会将代码、相关运行环境构建镜像，测试人员在宿主机上下载服务的镜像，使用容器启动镜像后即可运行服务进行测试；测试无误后运维人员申请机器，拉取服务器的镜像，在一台或多台宿主机上可以同时运行多个容器，对用户提供服务。
* 实现原理:通过Namespace（命名空间）实现进程隔离、UnionFilesystem（联合文件系统）实现文件系统隔离、ControlGroups（控制组）实现资源隔离。

### namespace

* 当一个宿主机上运行多个服务，就可能出现资源的争夺，进程之间的互相影响。

* 将宿主机上同时运行的多个服务划分成每个独立的服务，自己单独进程运行。

* Linux `Namespace`提供了一种内核级别隔离系统资源的方法，通过将系统的全局资源放在不同的`Namespace`中，来实现资源隔离的目的。不同`Namespace`的程序，可以享有一份独立的系统资源。目前Linux中提供了六类系统资源的隔离机制，分别是：

  - `Mount`: 隔离文件系统挂载点

    - 挂载
      - bind mount:一个「文件系统」可以被挂载到多个「挂载点」
      - `s_mount`作为链表头，`mnt_instance`作为链表节点，将属于同一文件系统的挂载instance连接起来。而挂载instance通过内嵌的*`vfsmount`* 中的`mnt_sb`，指向了其所属的文件系统。
    - 挂载传播
    - 绑定挂载

  - `UTS`: 隔离主机名和域名信息

  - `IPC`: 隔离进程间通信

  - `PID`: 隔离进程的ID

  - `Network`: 隔离网络资源

  - `User`: 隔离用户和用户组的ID

  - `Namespace`的操作接口包括`clone()`、`setns()`、`unshare()`以及还有`/proc`下的部分文件。

    - `CLONE_NEWNS`: 用于指定`Mount Namespace`
    - `CLONE_NEWUTS`: 用于指定`UTS Namespace`
    - `CLONE_NEWIPC`: 用于指定`IPC Namespace`
    - `CLONE_NEWPID`: 用于指定`PID Namespace`
    - `CLONE_NEWNET`: 用于指定`Network Namespace`
    - `CLONE_NEWUSER`: 用于指定`User Namespace`

  - clone系统调用:可以通过`clone`系统调用来创建一个独立`Namespace`的进程：

    ```c
    int clone(int (*child_func)(void *), void *child_stack, int flags, void *arg);
    ```

    - `flags`参数控制创建进程时的特性

  - `setns()`函数可以把进程加入到指定的`Namespace`中：

    ```c
    int setns(int fd, int nstype);
    ```

  - `unshare()`系统调用用于将当前进程和所在的`Namespace`分离，并加入到一个新的`Namespace`中：

    ```c
    int unshare(int flags)
    ```

  - PID namespace

    - PID Namepace对进程pid重新标号，不同的namespace下的进程可以有同一个pid

    - 实现：树状结构，最顶层由系统初始化产生init

      ![img](https://pic2.zhimg.com/v2-3980c3d48d7535ced3afcb510fd5d8cd_b.jpg)

    - 数据结构源码

      ```c
      //include/linux/pid_namespace.h
      struct pid_namespace {
      	struct kref kref;
      	struct idr idr;
      	struct rcu_head rcu;
      	unsigned int pid_allocated;
      	struct task_struct *child_reaper;
      	struct kmem_cache *pid_cachep;
      	unsigned int level;
      	struct pid_namespace *parent;
      #ifdef CONFIG_PROC_FS
      	struct vfsmount *proc_mnt;
      	struct dentry *proc_self;
      	struct dentry *proc_thread_self;
      #endif
      #ifdef CONFIG_BSD_PROCESS_ACCT
      	struct fs_pin *bacct;
      #endif
      	struct user_namespace *user_ns;
      	struct ucounts *ucounts;
      	struct work_struct proc_work;
      	kgid_t pid_gid;
      	int hide_pid;
      	int reboot;        /* group exit code if this pidns was rebooted */
      	struct ns_common ns;
      } __randomize_layout;
      ```

    - 子进程pid

      ```c
      struct pid
      {
      	atomic_t count;
      	unsigned int level;   //该进程所属的pid_ns的level，也就表示了这个pid对象在多少个pid namespace中可见。
      	/* lists of tasks that use this pid */
      	struct hlist_head tasks[PIDTYPE_MAX]; //使用该pid结构体的进程描述符集合
      	struct rcu_head rcu;
      	struct upid numbers[1];  //存储每层的pid信息的变成数组，长度就是上面的level
      };
      struct upid {
      	int nr; //该层pid ns 的PID值
      	struct pid_namespace *ns; //该层pid ns结构体地址
      };
      ```

    - pid保存

      - ```c
        init_task_pid(struct task_struct *task, enum pid_type type, struct pid *pid)
        {
        	 task->pids[type].pid = pid;
        }
        
        ```

    - 获取pid...

    - ![这里写图片描述](https://img-blog.csdn.net/20151119162640491)

* docker做了什么？

  * Docker uses a technology called `namespace` to provide the isolated workspace called the *container*. When you run a container, Docker creates a set of *namespaces* for that container.
  * ![img](https://pic3.zhimg.com/80/v2-386c59d740af12c267cf80bc335c7996_720w.jpg)

###  CGroups

* 对于物理资源的隔离，如cpu、内存、磁盘IO、网络IO等，采用CGroups进行实现

  ![img](https://pic2.zhimg.com/80/v2-924d5782805b5251b2e897c75b400da9_720w.jpg)

### UnionFS（union file system）

* 镜像rootfs实现：把多个目录联合放在同一个目录下，而这些目录的物理位置是分开的。用户制作镜像的每一步操作就是多增加一个目录（docker中称之为层layer）。启动docker容器时通过UnionFS把相关的层全放在一个目录里，作为容器的根文件系统，而容器的启动就是可写层，来对docker镜像进行操作。

* **层（layer）**

  ```text
  $ cat /proc/mounts |grep e4e2f1159f512ab74a6afbfeca51413cc3b6a24e86caccf91e40a9d611ce0a9b
  none /var/lib/docker/aufs/mnt/e4e2f1159f512ab74a6afbfeca51413cc3b6a24e86caccf91e40a9d611ce0a9b aufs rw,relatime,si=63e50947768841ec,dio,dirperm1 0 0
  
  $ cat /sys/fs/aufs/si_63e50947768841ec/br[0-9]*
  /var/lib/docker/aufs/diff/e4e2f1159f512ab74a6afbfeca51413cc3b6a24e86caccf91e40a9d611ce0a9b=rw
  /var/lib/docker/aufs/diff/e4e2f1159f512ab74a6afbfeca51413cc3b6a24e86caccf91e40a9d611ce0a9b-init=ro
  /var/lib/docker/aufs/diff/974a7e81b15c1eb6ea6c3c66dfb50dfcdf7b99b1e6458e2d3dca9451e2414106=ro
  /var/lib/docker/aufs/diff/fd68755d715f47edc7f5ceaa2e5dc6788d4ca36a4d50f51a92a53045cd0b9fb1=ro
  /var/lib/docker/aufs/diff/0e1237afa6d0fff72d9fdd5f84ef7275b1a49448d7523d590686131a3b129496=ro
  /var/lib/docker/aufs/diff/440bf3d93514f6a35bd99d4ac098d9b709e878146e355c670bd8f1f533c185c5=ro
  /var/lib/docker/aufs/diff/57e27832290597d0c5f2dc2ab55d1c53a7aa8a2a40eb6d21d014ad1210b1bb6f=ro
  /var/lib/docker/aufs/diff/55da955ef5752f9c3d1810a7b23e0325dd7947a0c0aaecf6ae373f3e33979143=ro
  ```

  * ro+wh层
    * `xxx=ro`结尾
    * whiteout：某个上层目录覆盖了下层的相同名字的目录。用于隐藏低层分支的文件，也用于阻止readdir进入低层分支。
      * readdir是一个用来读取目录的函数
  * Init层
    * Init层是Docker项目单独生成的一个内部层，专门用来存放/etc/hosts、/etc/resolv.conf等信息。这些文件本来属于只读的系统镜像层的一部分，但是用户往往需要在启动容器时写入一些指定的值比如hostname，所以就需要在可读写层对它们进行修改，但是并不想commit掉，所以单独挂一层。
  * rw层
    * 进行写操作时，修改产生的内容会以增量的方式出现在这个层中
    * 删除ro-wh层文件时，会在rw层创建对应的个whiteout文件。
    * 当我们使用完了这个被修改过的容器之后，还可以使用docker commit和push指令，保存这个被修改过的可读写层，并上传到Docker Hub上，供其他人使用。而与此同时，原先的只读层里的内容则不会有任何变化。

* AUFS(Advance UnionFS)

  ```
  $ tree
  .
  |-- a
  |   |-- a.log
  |   `-- x.log
  `-- b
      |-- b.log
      `-- x.log
      
  $ mkdir mnt
  $ mount -t aufs -o dirs=./a:./b none ./mnt
  $ tree ./mnt
  ./mnt
  |-- a.log
  |-- b.log
  `-- x.log
  
  $ echo test > mnt/x.log 
  $ cat mnt/x.log 
  test
  $ cat a/x.log 
  test
  $ cat b/x.log 
  
  $ echo test > mnt/b.log 
  $ cat mnt/b.log 
  test
  $ cat b/b.log 
  $ cat a/b.log 
  test
  
  $ touch b/bb.log
  $ rm mnt/a.log
  $ rm mnt/bb.log
  $ ls -al mnt
  -rw-r--r-- 1 root root    0 Sep 19 23:11 b.log
  -rw-r--r-- 1 root root    0 Sep 19 23:11 x.log
  $ ls -al a
  -rw-r--r-- 1 root root    0 Sep 19 23:15 .wh.bb.log
  -rw-r--r-- 1 root root    0 Sep 19 23:11 b.log
  -rw-r--r-- 1 root root    0 Sep 19 23:11 x.log
  $ ls -al b
  -rw-r--r-- 1 root root    0 Sep 19 23:11 b.log
  -rw-r--r-- 1 root root    0 Sep 19 23:14 bb.log
  -rw-r--r-- 1 root root    0 Sep 19 23:11 x.log
  ```

  

  * 可写分支的[负载均衡](https://cloud.tencent.com/product/clb?from=10680)
  * 全兼容UnionFS
  * rr表示real-read-only，rr标记的是天生就是只读的分支，这样，AUFS可以提高性能，比如不再设置inotify来检查文件变动通知。
  * AUFS的whiteout的实现是通过在上层的可写的目录下建立对应的whiteout隐藏文件来实现的。