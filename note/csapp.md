第一章 1.17

第四章 处理器体系结构 随便看看1.18

第五章 优化程序性能 1.19-1.20      1.21写lab

第六章 存储器层次结构 1.22

第七章 链接  1.23-1.25 1.26写lab

第八章 异常控制流 1.27-1.29 

第十章 系统级I/O  2.10-2.12 2.13写lab

第十一章 网络编程 学过

第十二章 并发编程 学过 2.14回顾一下

第九章 虚拟内存 2.15-2.17 2.18写lab



4.6 看网课中



# 深入理解计算机系统

## 第一章 计算机系统漫游

* 计算机系统由硬件和系统软件组成

### 1.1 信息就是位+上下文

1. 文本文件以字节序列的形式存储在文件中，只含有`ASCII`码,其他所有文件都成为二进制文件。
2. 数据都是由比特串表示  **everything is bits**
2. Byte and Bit
2. ！！number =>1
2. x<<8 = ? 逻辑左移和算术左移 
2. Unsigned and Signed number

### 1.2 程序被其他程序翻译成不同的格式

1. 高级语言转化为低级机器语言指令->按照可执行目标程序的格式打包->二进制磁盘文件

   ```
   linux> gcc -o hello hello.c
   ```

2.  编译系统：预处理器（文本），编译器（文本），汇编器（二进制），链接器（二进制）

   1. 预处理：处理宏定义
   2. 编译阶段：翻译位汇编语言
   3. 汇编阶段：翻译为机器语言
   4. 链接阶段：合并所有调用函数

### 1.3 了解编译系统如何工作的原因

1. 优化程序性能
2. 理解链接时出现的错误
3. 避免安全漏洞

### 1.4 处理器读并解释储存在内存中的指令

#### 1.4.1 系统的硬件组成

1. 总线：携带信息字节并负责在各个部件间传递
2. I/O设备：Input&Output
3. 主存：执行程序时，用来存放程序和程序处理的数据
4. 处理器（CPU）：核心->PC（程序计数器），解释（执行）存储在主存中的指令的引擎

#### 1.4.2 运行hello程序

1. 在shell中输入`./hello`，shell会将之后输入的字符读入寄存器，再存放到内存中，结束输入后，shell执行一系列指令来加载可执行的文件将其从磁盘复制到主存。
2. 直接存储器存取（`DMA`）

### 1.5 高速缓存cache

1. 针对处理器和主存之间读取数据效率的差异

### 1.6 储存设备形成层次结构

1. 从高到低：寄存器->`L1`高速缓存->`L2`高速缓存->`L3`高速缓存->主存->本地二级存储->远程二级存储（分布式文件系统，Web服务器）

### 1.7 操作系统管理硬件

1. 应用程序和硬件之间的一层插件
2. 功能：
   1. 防止硬件失控的应用程序滥用
   2. 向应用程序提供一个简单一致的机制来控制复杂的硬件

#### 1.7.1 进程

1. 概念：操作系统对一个正在运行的程序的一种抽象
2. 并发运行：一个进程的指令和另一个进程的指令交错执行/单核多核系统
3. 上下文：操作系统保持跟踪进程运行所需要的所有状态的所有信息
4. 从一个进程到另一个进程的转换是由操作系统内核管理的

#### 1.7.2 线程

1. 多个线程组成进程

#### 1.7.3  虚拟内存

1. 抽象概念：为每个进程提供一个假象（每个进程都在独占地使用主存）
2. 大量准确定义的区构成 程序代码和数据->堆->共享库->栈->内核虚拟内存

#### 1.7.4 文件

1. 就是字节序列

### 1.8 系统之间利用网络通信

1. 网络可视作一个I/O设备 

   本地客户端 ------- 远程服务器

### 1.9 

#### 1.9.1 Amdahl定律

想要显著加速整个系统，必须提升全系统中相当大的部分的速度

#### 1.9.2 并发和并行

1. 线程级并发 一个进程中执行多个控制流

   1. 超线程（同时多线程）

2. 指令级并行 同时执行多条指令的属性

3. 单指令、多数据并行 （SIMD并行）

   由特殊硬件实现，允许一条指令产生多个可以并行执行的操作

#### 1.9.3 计算机系统中抽象的重要性

1. 文件是I/O设备的抽象
2. 虚拟内存是对程序存储器的抽象
3. 进程是对一个正在运行的程序的抽象



## 第四章 处理器体系结构

### 4.1 `Y84-64` 指令集体系结构

#### 4.1.1 程序员可见的状态

1. `Y86-64`程序中的每一条指令都会被读取或者修改处理器状态的某些部分
2. 寄存器 条件码 内存 状态码（指示是否出现异常）
3. `Y84-64`程序用虚拟地址来引用内存位置

#### 4.1.2 `Y86-64`指令

1. `movq`:`irmovq`,`rrmovq`,`mrmovq`,`rnmovq` 立即数i、寄存器r、内存m
2. 整数操作指令
3. 跳转指令
4. 条件传送指令
5. `call` `ret` 函数调用
6. `pushq` `popq`出入栈
7. `halt` 暂停执行

#### 4.1.3 指令编码

1.  字节编码分为两个部分，高四位-代码，第四位-功能
2. 程序寄存器每一个都有一个寄存器标识（ID）
3. 指令编码可能附加常数字、寄存器指示符字节

#### 4.1.4 `Y86-64`异常

1. `HLT`、`ADR`非法地址、`INS`非法指令

#### 4.1.5 `Y86-64`程序

1. 不能使用立即数
2. 伪指令

### 4.2 逻辑设计和硬件控制语言`HCL`

#### 4.2.1 逻辑门

AND OR NOT

#### 4.2.2 组合电路和`HCL`布尔表达式

1. 组合电路 

   1.每个输出必须连接到1.一个系统输入2.某个存储器单元的输出3.某个逻辑门的输出

   2.多个逻辑门输出连接在一起可能会使线上的信号矛盾

   3.网必须无环

2. 多路复用器`MUX`

#### 4.2.3 字级的组合电路和`HCL`整数表达式

位级实现，字级相等

算数/逻辑单元

#### 4.2.4 集合关系

```
iexpr in {iexpr1, iexpr2, ..., iexprk}
```

#### 4.2.5 储存器和时钟

1. 组合电路从本质上讲，不存储任何信息，只简单地响应输入信号，产生等于输入地某个函数地输出。
2. 时序电路：时钟，存储器设备：1.时钟寄存器 2.随机访问存储器
3. 寄存器文件：多端口

### 4.3 `Y86-64`的顺序实现

#### 4.3.1 将处理组织成阶段

1. 取指：从内存读取指令字节
2. 译码从寄存器文件读入最多两个操作数
3. 执行：要么执行指令，要么加减栈指针
4. 访存：将数据写入内存
5. 写回：将结果写到寄存器文件
6. 更新PC：将PC设置成下一条指令的地址
7. 处理器无限循环，执行这些阶段

#### 4.3.2 `SEQ`硬件结构

1. 在`SEQ`中，所有的硬件单元的处理都在一个时钟周期内完成
2. 时钟寄存器，硬件单元，控制逻辑块，线路名字，数据连接

#### 4.3.3 `SEQ`的时序

1. 原则：从不回读（处理器从来不需要为了完成一条指令的执行而去读由该指令更新了的状态

#### 4.3.4 `SEQ`阶段的实现

1. 取指：指令内存硬件单元以PC作为第一个字节的地址，一次从内存中读取出10个字节，第一个字节为指令字节
2. 译码和回写：寄存器文件有四个端口，支持同时进行两个读和写
3. 执行阶段
4. 访存阶段：读写数据，内存读写的地址总是`valE`和`valA`
5. 更新PC阶段：依据指令的类型和是否要选择分支

### 4.4 流水线的通用原理

#### 4.4.1 计算流水线

1. 每个阶段完成指令的一部分
2. 电路延迟 10的-12次方（单位）秒
3. 吞吐量
4. 局限性： 1. 不一致的划分 2.流水线过深，收益反而下降

#### 4.4.4 带反馈的流水线系统

1.  针对一些相邻指令很可能是相关的程序

### 4.5 `Y86-64`的流水线实现

#### 4.5.1 重新安排计算阶段

1. 修改以后被称为`SEQ+`
2. 创建状态寄存器来保存在一条指令执行过程中计算出来的信号
3. 动态地计算PC，不存放程序计数器
4. 称为电路重定时

#### 4.5.2 插入流水线寄存器

1. 尝试在各个阶段插入流水线寄存器

#### 4.5.3 对信号进行重新排列和标号

## 第七章 链接

* 概念：将各种代码和数据片段手机并组合成一个单一文件的过程
* 分离编译的实现变为可能

### 7.1 编译器驱动程序

* 编译器驱动程序

* 源文件->（经过）翻译器->(被翻译为)可重定位目标文件->链接齐->完全链接的可执行目标文件

  ```shell
  cp main.c /tmp/main.i // 特别的，/tmp文件夹是linux中被用于存放各种临时文件的地方，cpp为c预处理器
  cc1 /tmp/main.i -Og -o /tmp/main.s // 特别的，cc1为C语言编译器，main.s为一个ASCII汇编语言文件
  as -o /tmp/main.o /tmp/main.s // 特别的，as为汇编器，main.o 为可重定位目标文件(relocatable object file)
  ld -o prog /tmp/main.o /tmp/sum.o // 特别的，ld为链接器程序，将必要文件组合起来
  ./prog // 运行程序
  ```

### 7.2 静态链接

* `ld`是一个静态链接器
* 链接器的两个主要任务
  * 符号解析
  * 重定位
* 目标文件是纯粹的字节块的集合，链接器所做的任务很少

### 7.3 目标文件

* 目标文件的三种形式：可重定位、可执行、共享
* 目标模块的定义：一个字节序列
* 目标文件：一种（以文件形式）（存放在磁盘）的目标模块
* unix、windows、mac的目标文件格式并不相同

### 7.4 可重定位的文件

* ELF文件格式
  * .`text` :已编译的机器代码
  * .`rodata`:`read only data`
  * .`data`:初始化的全局和静态c变量
  * .`bss`:未初始化的全局和静态c变量
  * .`symtab`：symbol_table
  * .`rel.text`:全局变量重定位信息
  * .`debug`：调试表
  * .`line`：源程序和机器指令的映射
  * .`strtab`:string_table

### 7.5 符号和符号表

* 三种符号：该模块定义的全局符号，别的模块定义的全局符号，局部符号（static）
* 非静态变量在栈中被管理
* 三个可重定位文件中的伪节：`UNDEF`未定义符号,`COMMON`未初始化的全局变量,`ABS`不被重定向的符号
* COMMON和.`bss`的区别

### 7.6 符号解析

* 全局符号的符号解析

#### 7.6.1 解析多重定义的全局符号

* 全局符号分为强和弱
* 三个规则
  1. 不允许同名强符号
  2. 有强则强
  3. 多个弱符号->随机选择
* gcc中这类错误只会产生一个警告，使用`-fno-common`或者`-Werror`选项将警告变成错误

#### 7.6.2 与静态库链接

* 静态库的概念：将所有相关的目标模块打包成为一个单独的文件
* 存档：Linux中一种特殊文件格式（.a)
* 创造静态库工具`AR`

#### 7.6.3 链接器使用静态库

* 链接器从左到右链接文件，如果链接过程中出现尚未定义但是被引用的符号则会报错！、

### 7.7 重定位

* 两步
  * 重定位节和符号定义
  * 重定位节中的符号引用

### 7.8 可执行目标文件

* ELF文件格式
  * .`init`
  * 其他与可重定位目标文件格式相同
  * .data有对齐要求

### 7.9 加载可执行的目标文件

* 加载：将程序复制到内存并运行的过程

### 7.10 动态链接共享库

* 目的：解决静态库频繁使用的函数对内存的重复占用造成浪费
* 动态链接
* 动态链接器
* 共享库也被称为共享目标，是一个目标模块
* 内存中.text节可被不同进程共享
* 通过加载技术使得运行时可以解析对共享库的引用

### 7.11 从应用程序中加载和链接共享库

* 例子：分发软件，构建高性能web服务器
* `dlopen`函数，为动态链接器提供一个简单的接口，允许应用程序在运行时加载和链接共享库

### 7.12 位置无关代码

1. `PIC`数据引用：数据段和代码段的距离总是不变->使用：全局偏移量表`GOT`
2. `PIC`函数调用：延迟绑定->GOT和过程链接表`PLT`协同工作

### 7.13 库打桩机制

* 不使用共享函数而是执行自己的代码

#### 7.13.1 编译时打桩

* `-I.`参数

#### 7.13.2 链接时打桩

* 静态链接器使用`--wrap`标志打桩

#### 7.13.3 运行时打桩

* `LD_PRELOAD`

### 7.14 处理目标文件的工具

* `AR`
* `STRINGS`
* `STRIP`
* `NM`
* `SIZE`
* `READELF`
* `OBJDUMP`
* `LDD`

## 第八章 异常控制流

### 8.1 异常

1. 概念：控制流中的突变
2. 异常表：一张跳转表，进行简介过程调用，到异常处理程序中，一共三种情况
   1. 返回给`curr`
   2. 返回给`next`
   3. 终止

#### 8.1.1 异常处理

1. 异常号：由处理器或者操作系统内核的设计者分配，每种异常异常号**唯一**
2. 异常表由操作系统分配，初始化
3. 过程
   1. 检测到异常
   2. 确认异常号
   3. 处理器触发异常->(通过异常表转到对应程序)
   4. 处理完成后通过“从中断返回”指令选择性地回到对应的程序中

#### 8.1.2 异常的类别

| 类别 | 原因                 | 异步/同步 | 返回行为    |
| ---- | -------------------- | --------- | ----------- |
| 中断 | 来自I/O设备的信号    | 异步      | next        |
| 陷阱 | 有意的异常           | 同步      | next        |
| 故障 | 潜在**可恢复**的错误 | 同步      | `curr`/终止 |
| 终止 | 不可恢复的错误       | 同步      | 终止        |

##### 1.中断

* 通过引脚发信号（中断时电压升高）
* 注意**异步**

##### 2.陷阱和系统调用

* **有意的**
* 用途：系统调用（用户程序和内核之间）
* `syscall`

##### 3.故障

* 由错误情况引起
* 被修正则返回`curr`否则返回到内核中的`abort`->终止
* 例子：缺页异常

##### 4.终止

* 不可恢复的错误引起
* 直接返回`abort`例程

#### 8.1.3 Linux/x86-64系统中的异常

1. 共256种
2. 故障和终止
   1. 除以0（异常号0）
   2. 只读段保护引起的故障（segmentation fault）（异常号13）
   3. 缺页->重复映射页面（异常号14）
   4. 硬件出错（异常号18）

3. 系统调用
   * 每个系统调用都有**唯一**整数号
   * C库实际上已经有很多封装的函数可以调用，不必直接调用`syscall`
   * 系统级函数的概念：系统调用以及相关联的包装函数
   * 所有参数都是通过寄存器传递

### 8.2 进程

* 经典定义：一个执行种程序的实例
* 进程的上下文：程序正确运行所需的状态集合（内核重新启动一个被抢占的进程所需的状态）

#### 8.2.1 逻辑控制流

* 程序计数器（PC）值
* 逻辑控制流：PC值的序列
* 实现：轮流使用处理器，称为多任务(时间分片)

#### 8.2.2 并发流

* 概念：不同的逻辑流的执行在时间上重叠
* 并行流：在不同的处理器核上运行的流

#### 8.2.3 私有地址空间

* 通用结构

#### 8.2.4 用户模式和内核模式

* 处理器通过某个控制寄存器上的模式位来控制进程可访问的地址空间范围和可执行的指令，此时称为内核模式，可以访问任何指令和任意内存位置
* 用户模式没有设置模式位，不允许执行特权命令，但可以通过异常进入内核模式
* Linux种的`/proc`

#### 8.2.5 上下文切换

* 本质：较高层形式的异常控制流
* 上下文的概念
* 调度：内核切换进程的决策
* 步骤
  1. 保存当前进程的上下文
  2. 恢复某个先前被抢占的进程被保存的上下文
  3. 将控制传递给这个新恢复的进程
* 系统调用时可能会发生上下文切换：比如当前进程阻塞
* 中断也可能会发生上下文切换：定时器
* 一个系统调用的例子：
  * read函数（用户模式->内核模式）
  * 等待磁盘读取数据（发生阻塞，此时切换上下文）
  * 数据到达，发送中断信号（内核再次切换上下文）

### 8.3 系统调用错误处理

* 错误处理使代码变得臃肿难以读懂？
* 通过错误报告函数简化->再通过错误处理包装函数进一步简化
* 一般而言小写为系统级函数，大写为包装函数

```c
void unix_error(char *msg){
    fprintf(stderr,"%s:%s\n",msg,strerror(errno));//errno 为设置的全局变量<errno.h>
    exit(0);
}

pid_t Fork(void){
	pid_t pid;
	
	if((pid = fork()) < 0){
		unix_error("Fork error");
	}
	return pid;
}
```

### 8.4 进程控制

#### 8.4.1 获取进程ID

* `getpid` | `getppid` 返回的id不相同

#### 8.4.2 创建和终止进程

* 进程总处于三种状态-运行、停止、终止
* fork函数创建的进程得到的是用户级别的虚拟地址空间，并且是父进程的副本；fork函数被调用一次，返回两次值，分别返回给父进程和子进程，其中在子进程中返回0（因为子进程id总是非零的，用于区分）

* 一个例子，在原书的基础上进行修改

```c
int main(){
    pid_t pid;
    int x = 1;

    pid = Fork();
    if (pid < 0){
        printf("error!");
    }
    if(pid == 0){
        printf("child : x = %d\nI am the chid process, my process ID is %d\n",++x,getpid());
        printf("father\n");
        exit(0);
    }
    printf("parent: x = %d\nI am the parent process, my process ID is %d\n",--x,getpid());
    printf("ok\n");
    exit(0);
}


结果：
parent: x = 0
I am the parent process, my process ID is 3846
child : x = 2
ok
I am the chid process, my process ID is 3860
father
```

* 可以发现
  * 有两个返回值：`pid`=0和`pid`>0
  * 他们具有相同但是独立的地址空间（x)
  * 共享文件：子进程继承了父进程所有打开文件，都输出在屏幕上 
  * 并发执行，是并发的独立进程，交替执行逻辑流指令
* 进程图（刻画程序语句的偏序的一种简单前趋图）

```c
int main(){
	Fork();
    Fork();
    printf("hello\n");
    exit(0);
}

结果：
hello
hello
hello
hello
```

* 上面这个程序，四个进程每一个都调用了一次`printf`函数

#### 8.4.3 回收子进程

* 这件事并不是由内核做，而是交给父进程做，内核只会传递子进程的状态
* 如果父进程终止了：内核会使用`init`进程来回收
* 僵死进程：一个终止但是还未被回收的进程
* `waitpid`函数
  * 判定等待集合的成员->`pid`
  * 修改默认行为->`options`
  * 检查已回收子进程的退出状态->`statusp`（指针）->`status`
  * 错误条件->`pid+errno`
* `wait`函数/简化版`waitpid`（默认`pid`为-1（父进程所有的子进程），`options`是0（不使用））

* 程序不会按照特定顺序回收子进程

* 一个例子

  ```c
  #include <sys/types.h> 
  #include <sys/wait.h> 
  #include <unistd.h> 
  int main() {     
      pid_t pc, pr;     
      pc=Fork();     
      if(pc<0)      
          printf("Error occured on forking.\n");     
      else if(pc==0)     
      {
          sleep(4);
          exit(0);     
      }         
      do     
      {     
          pr=waitpid(pc, NULL, WNOHANG);
          if(pr==0)      
          {         
              printf("No child exited\n");       
              sleep(1);      
          }     
      }while(pr==0);        
      if(pr==pc)      
          printf("successfully release child %d\n", pr);     
      else
          printf("some error occured\n"); 
  }
  
  输出结果：
  No child exited (1s)
  No child exited (1s)
  No child exited (1s)
  No child exited	(1s)
  successfully release child 8579
  ```

#### 8.4.4 让进程休眠

* `sleep`函数

* `pause`函数

  ```c
  // 8.5作业
  void snooze(int end){
      for(int i = 0; i < end;i++  ){
          printf("slept for %d of %d secs\n",i+1,end);
          sleep(1);
      }
      printf("end!");
  }
  ```

#### 8.4.5 加载并运行程序

* `execve`函数

* 调用一次，从不返回（只有出现错误才会返回）

* `argv`参数

* `envp`参数->每一个都是一个`k-v`键值对

* `main`函数的参数：

  * `argc`:`argv[]`的非空数量
  * `argv`：`argv[]`头指针
  * `envp`：`envp[]`的头指针

* `getenv`函数，在`envp[]`中搜索`k-v`键值对，返回指向v的指针

* `setenv`函数：存在则替代，不存在则创建

  ```c
  // 8.6 作业
  int main(int argc,char *argv[],char *envp[]){
      printf("Command-ine arguments:\n");
      for(int i = 0 ;argv[i]!=NULL;i++){
          printf("argv[%2d] : %s\n",i,argv[i]);
      }
  
      printf("Environment variables:\n");
      for(int i = 0;envp[i]!=NULL;i++){
          printf("envp[%2d] : %s\n",i,envp[i]);
      }
  }
  ```

#### 8.4.6 利用`fork`和`execve`运行程序

* 一个简单的shell实现，但是没有实现回收子进程的功能

### 8.5 信号

* 作用：允许进程和内核中断其他信号，通知进程系统发生了一个某种类型的事件
* 每种事件对应不同的信号

#### 8.5.1 信号术语

* 两个步骤
  * 发送信号->由内核实现（更新目的进程上下文的某种状态）
  * 接收信号->目的进程做出相应反应（一个信号最多只能被接受一次，接收了之后内核中地该信号的对应位就会被清除）
* 待处理信号：发送了还没被接收的信号，任何时候一种类型最多一个待处理信号，多的信号会被直接丢弃
* 目的进程可以选择性地阻塞某个信号

#### 8.5.2 发送信号

* 进程组
  * 每个进程都只属于一个进程组，默认父进程和子进程属于同一个进程组
  * `getpgrp`函数返回当前进程组id
  * `setpgid`函数改变进程的id，change（`pid`/第一个参数 to `pgid`/第二个参数（若为0则用`pid`作为进程组id））
* `/bin/kill`指令发送信号
* 从键盘发送信号
  * 作业(job):Unix shell表示对一条命令行求值而创建的进程
  * 前台作业，前台进程组
  * 后台作业，后台进程组
* `kill`函数
  * 参数`pid`，大于0，等于0，小于0三种情况
  * 参数`sig`（要发送的信号）
* `alarm`函数
  * 参数`secs`
  * 默认信号：`SIGALRM`
  * 默认进程：调用进程

#### 8.5.3 接收信号

* 进程模式切换（内核模式->用户模式）时，内核会检查进程的待处理信号集合，如果不为空则会选择一个信号强制让进程接收

* 信号预定义的默认行为

  * 进程终止
  * 进程终止并转储内存
  * 进程停止直到被重启
  * 忽略该信号

* `signal`函数可以修改信号的默认行为，信号`SIGSTOP`，`SIGKILL`不能修改

  * 参数`signum`（信号）

  * 参数handler：

    * `SIG_IGN`->忽略
    * `SIG_DFL`->恢复默认行为
    * 用户自定义的函数（信号处理程序）

  * 例子：

    ```c
    #include <stdio.h>
    #include <signal.h>
    void sigint_handler(int sig){
    
        printf("Caught SIGINT!\n");
        exit(0);
    }
    
    int main(){
        if(signal(SIGINT,sigint_handler) == SIG_ERR){
            printf("error");
        }
        pause();
    
        return 0;
    }
    
    输出：
    ^CCaught SIGINT!  
    ```

#### 8.5.4 阻塞和解除阻塞信号

* 隐式阻塞
  * 另一个信号引起的阻塞
* 显示阻塞
  * 使用`sigprocmask`函数和他的辅助函数来控制阻塞

#### 8.5.5 编写信号处理程序

* 安全的信号处理

  * 处理程序要尽可能简单
  * 在处理程序中只调用异步信号安全的函数
  * 信号处理中唯一安全输出是`write`函数
  * 保存和恢复`errno`，确保其他程序不会修改自己所需的`errno`
  * 阻塞所有的信号
  * 用`volatile`声明全局变量
  * 用`sig_atomic_t`声明标志，这种整型数据类型的读和写保证是原子的（不可中断）
  * 例子：

  ```c
  volatile sig_atomic_t g_flag = 0;
  
  void signal_handler(int signum) {
      g_flag = 1;
  }
  
  int main(void)
  {
      signal(SIGINT, signal_handler);
  
      while( !g_flag ){
          puts( "wait" );
          sleep(1);
      }
  
      puts( "exit" );
  }
  ```

* 正确的信号处理
  * 未处理的信号是不排队的（多了直接丢弃）->逆向思维：如果有未处理信号存在，那么肯定有至少一个信号到达了
  * 一定要回收僵死进程
* 可移植的信号处理
  * 系统调用可被中断
  * 函数语义可能不同
  * `sigaction`函数，过于复杂->`Signal`

#### 8.5.6 同步流以避免并发错误

* 同步的问题：竞争（race)
* 解决办法，在产生竞争之前，使用信号避免竞争->同时也可能因为信号没有及时解除引发新的问题

#### 8.5.7 显示地等待信号

* 一个常见的例子：比如shell会显示地等待用户输入指令
* pause函数和sleep函数都有缺陷->`signsuspend`函数
  * 等价于使用mask后使用pause，但是他是原子性的

### 8.6 非本地跳转

* 用户级异常控制流形式，从一个函数转移到另一个执行的函数
* `setjmp`函数和`longjmp`函数
  * 在`env`缓冲区中保存当前调用函数
  * `setjmp`函数只被调用一次，但是它返回多次，每个`longjmp`他都返回
  * `longjmp`函数不返回
  * 因为可能会跳过释放内存的过程而造成内存泄漏
* C++和JAVA中的软件异常
  * JAVA中的`try-catch`结构类似于`setjmp`函数
  * `throw`函数类似于`longjmp`函数

### 8.7 操作进程的工具

* `strace`、`ps`、`top`、`pmap`、`/proc`

## 第九章 虚拟内存

* 目的：为了更加有效管理内存并且减少出错
* 本质：对主存的抽象概念

### 9.1 物理和虚拟寻址

* 主存->一个连续地字节大小地单元组成地数组
* 物理寻址：`cpu`访问内存最自然的方式
* 虚拟寻址：cpu通过生成一个虚拟地址来访问主存
* 地址翻译：将虚拟地址翻译为物理地址
* 内存管理单元：利用主存中的查询表来翻译虚拟地址，该表由操作系统管理

### 9.2 地址空间

* 定义：是一个非负整数的有序集合
* 线性地址空间：地址空间中的整数是连续的
* 虚拟地址空间：`cpu`从地址空间生成的虚拟地址（n位地址空间->2^n虚拟地址空间）
* 单位：K->M->G->T->P->E

### 9.3 虚拟内存作为缓存的工具

* 物理页：物理内存被分割，页帧
* 虚拟页面的集合的三种状态：
  * 未分配的
  * 缓存的
  * 未缓存的

#### 9.3.1 DRAM缓存的组织结构

* DRAM是全相联的

#### 9.3.2 页表

* 一种数据结构：页表条目（PTE)的数组
  * PTE由一个有效位和一个地址字段组成（csapp中假设）
    * 有效位表示该虚拟页当前是否被缓存
  * 在[Page Table Entries in Page Table - GeeksforGeeks](https://www.geeksforgeeks.org/page-table-entries-in-page-table/)中，PTE由六个部分组成![Lightbox](https://media.geeksforgeeks.org/wp-content/uploads/Capture-24.png)
* 页表上面有虚拟地址和真实物理地址的映射
* 由操作系统维护内容
* 2^虚拟地址大小（n)/页大小(P)=nums（PTE）

#### 9.3.3 页命中

#### 9.3.4 缺页

* 缺页：虚拟内存中DRAM缓存不命中
* 缺页异常处理程序
* 触发缺页的过程
  * cpu调用未缓存数据
  * 从PTE的有效位判断出未缓存
  * 触发异常->调用对应异常处理程序
  * 内核将未缓存数据复制到选中的DRAM中
  * 异常处理程序返回
  * 重新启动调用数据的指令
  * 此时可以从DRAM中获取
* 交换、页面调度
* Linux中的`getrusage`函数

#### 9.3.5 分配页面

* 在磁盘上创建空间，更新对应PTE

#### 9.3.6 局部性

* 局部性使得不命中发生概率变得低
* 工作集、常驻集合

### 9.4 虚拟内存的内存管理

* 简化了内存管理
* 实际上，每个进程都有一个独立的页表
  * 多个虚拟页面可以映射到同一个共享物理页面上
* `VM`（virtual memory）的简化
  * 链接：允许每个进程使用相同的基本格式，不用管实际数据在何处
  * 加载：加载器为数据分配虚拟页，将页表条目只想目标文件中适当的位置
  * 共享：将适当虚拟页映射到相同的物理页面
  * 内存分配：从虚拟内存页面中分配

### 9.5 虚拟内存的内存保护

* 虚拟内存提供内存保护的方法
  * 只读段保护
  * 内核数据保护
  * 其他进程私有段数据保护
  * 共享页面保护
* 在`PTE`中添加许可位
* 违反许可条件->触发一般保护故障（“段错误”）

### 9.6 地址翻译

* 相关符号![image-20220518090707661](C:\Users\zhangxinyu\AppData\Roaming\Typora\typora-user-images\image-20220518090707661.png)

* 地址翻译是一个N元素的地址空间（VAS）中的元素和一个M元素的物理地址空间（PAS）中元素的映射
  * 要么对应到一个物理地址空间
  * 要么是一个空集
* 内存管理单元（`MMU`）利用页表来实现这种映射
  * CPU中的页表基址寄存器（一个控制寄存器）指向页表
  * 虚拟地址包括偏移和页号两部分
  * `MMU`利用虚拟页号（`VPN`）来选择适当的`PTE`
* `CPU`硬件执行步骤
  * 处理器生成虚拟地址，传递到`MMU`
  * `MMU`生成`PTE`，并向高速缓存/主存请求得到它
  * 高速缓存/主存向`MMU`返回`PTE`
  * `MMU`构造物理地址，并传递给高速缓存/主存
  * 高速缓存/主存返回请求数据给处理器
* 页面命中步骤：
  * 1-3步与CPU硬件执行前三步相同
  * `PTE`中有效位0，`MMU`触发异常，传递`CPU`控制到内核
  * 内核异常处理程序处理缺页异常
  * 更新内存中`PTE`
  * 返回原程序，再次执行原指令

#### 9.6.1 高速缓存和虚拟内存

* 大多数系统使用物理寻址的方式

#### 9.6.2 利用`TLB`加速地址翻译

* 试图增加命中->翻译后备缓冲器（`TLB`）它是一个小的虚拟寻址的缓存
* `TLB`命中的步骤
  * `CPU`产生虚拟地址
  * `MMU`从`TLB`中取出相应的`PTE`
  * `MMU`将虚拟地址翻译为物理地址，发送到高速缓存/主存
  * 高速缓存/主存返回给`cpu`
* `TLB`未命中
  * 新步骤：`MMU`会寻找对应`PTE`，新`PTE`可能会覆盖`TLB`中原本的条目

#### 9.6.3 多级页表

* 单独页表随着地址空间的大小而变化，过于庞大的页表对于地址空间是一个负担
* 压缩页表的常用方法之一就是使用层次结构的页表
* n个连续页表映射一个上级页表
* 实际上多级页表的地址翻译并不比单机页表慢很多

#### 9.6.4 端到端的地址翻译

* 一个存在`TLB`和`L1`的小系统的例子
* `TLB`根据`VPN`的位进行虚拟寻址
* 有效位为0时数据为什么都无所谓
* 考虑多种不命中或者缺页的情况

### 9.7 Intel Core `i7`/Linux内存系统

//稍稍有些看不懂，先搁置一下，看第十章

## 第十章 系统级I/O

* 输入输出是主存和外部设备之间复制数据的过程
* 文件元数据需要使用系统级别的I/O来获取
* 高级I/O函数编程有些情况存在冒险

### 10.1 Unix I/O

* 所有的I/O设备都被模型化为文件，输入输出被当作读写操作
* 使得所有输入输出都能以一种统一且一致的方式执行
  * 打开文件：程序向内核发出请求->内核返回一个描述符（整数类型）
    * shell创建每个进程时都会有三个打开文件：标准输入(0)，标准输出(1)，标准错误(2)
    * `<unistd.h>`描述符宏定义
  * 改变当前文件的位置：内核保存文件位置k（从文件开头开始的字节偏移量），初始为0
    * `seek`函数
  * 读写文件：读->将文件从当前位置开始指定数量字节复制到内存当中;写->将内存中指定数量的字节复制到文件中
    * `EOF`条件，应用程序能检测到
  * 关闭文件
    * 应用程序通知内核关闭文件
    * 内核释放数据结构，恢复描述符到描述符池中
    * 无论进程如何，内核最终都会关闭文件，释放资源

### 10.2 文件

* 文件的不同类型
  * 普通文件（.`txt`\.`dat`)
  * 目录：包含一组链接的文件
  * 套接字（socket）:与其他进程通信的文件
  * 其他文件类型
* Linux内核中所有文件在一个目录层次结构里面
* 每一个进程都有一个当前工作目录
* 路径名是一个字符串
  * 绝对路径
  * 相对路径

### 10.3 打开和关闭文件

* `open`函数

  ```c
  #include <sys/types.h>
  #include <sys/stat.h>
  #include <fcntl.h>
  
  int open(char *filename, int flags, mode_t mode);
  ```

  * 返回的是文件描述符
  * flags是访问方式（只读，只写，读写），可以通过|连接
  * mode指定了新文件的访问权限位
  * `umask`函数

* `close`函数关闭

### 10.4 读和写文件

* `read`函数
* `write`函数
* `size_t`为无符号`long`型，`ssize_t`为有符号`long`型，有符号型是因为错误处理要返回-1
* 不足值
  * 返回不足值的三种情况
    * 文件字节<读取片字节数
    * 从终端读取
    * 读写`socket`
  * 健壮的Web服务必须处理好不足值

### 10.5 用RIO包健壮地读写

* 提供了两类函数
  * 无缓冲输入输出
  * 带缓冲的输入

#### 10.5.1 无缓冲输入输出函数 

* `rio_readn`函数

* ```c
  ssize_t rio_readn(int fd, void *usrbuf, size_t n) 
  {
      size_t nleft = n;
      ssize_t nread;
      char *bufp = usrbuf;
  
      while (nleft > 0) {
  	if ((nread = read(fd, bufp, nleft)) < 0) {
  	    if (errno == EINTR) /* Interrupted by sig handler return */
  		nread = 0;      /* and call read() again */
  	    else
  		return -1;      /* errno set by read() */ 
  	} 
  	else if (nread == 0)
  	    break;              /* EOF */
  	nleft -= nread;
  	bufp += nread;
      }
      return (n - nleft);         /* return >= 0 */
  }
  ```

  #### 10.5.2 有缓冲的输入输出函数

* `rio_read`函数

  ```c
  static ssize_t rio_read(rio_t *rp, char *usrbuf, size_t n)
  {
      int cnt;
  
      while (rp->rio_cnt <= 0) {  /* Refill if buf is empty */
  	rp->rio_cnt = read(rp->rio_fd, rp->rio_buf, 
  			   sizeof(rp->rio_buf));
  	if (rp->rio_cnt < 0) {
  	    if (errno != EINTR) /* Interrupted by sig handler return */
  		return -1;
  	}
  	else if (rp->rio_cnt == 0)  /* EOF */
  	    return 0;
  	else 
  	    rp->rio_bufptr = rp->rio_buf; /* Reset buffer ptr */
      }
  
      /* Copy min(n, rp->rio_cnt) bytes from internal buf to user buf */
      cnt = n;          
      if (rp->rio_cnt < n)   
  	cnt = rp->rio_cnt;
      memcpy(usrbuf, rp->rio_bufptr, cnt);
      rp->rio_bufptr += cnt;
      rp->rio_cnt -= cnt;
      return cnt;
  }
  ```

* `rio_readlineb`函数

  ```c
  ssize_t rio_readlineb(rio_t *rp, void *usrbuf, size_t maxlen) 
  {
      int n, rc;
      char c, *bufp = usrbuf;
  
      for (n = 1; n < maxlen; n++) { 
          if ((rc = rio_read(rp, &c, 1)) == 1) {
  	    *bufp++ = c;
  	    if (c == '\n') {
                  n++;
       		break;
              }
  	} else if (rc == 0) {
  	    if (n == 1)
  		return 0; /* EOF, no data read */
  	    else
  		break;    /* EOF, some data was read */
  	} else
  	    return -1;	  /* Error */
      }
      *bufp = 0;
      return n-1;
  }
  ```

* `rio_read`函数是核心函数，它带有缓冲
  * 缓冲区：`rp->rio_cnt`处理不足值
  * 并且错误返回与原本的`read`函数保持一致

### 10.6 读取文件元数据

* `stat`函数

* `fstat`函数

* 头文件:

  ```c
  #include <unistd.h>
  #include <sys/stat.h>
  ```

* `st_mode`的常用宏定义

  * `S_ISREG`(m)->is a file?
  * `S_ISDIR`(m)->is a `dir`?
  * S_SISOCK(m)->is a socket?

* 查看文件是否有读写权限

  * `stat.st_mode` &`S_IRUSR`!=0->can read
  * `stat.st_mode` &`S_IWUSR`!=0->can write

### 10.7 读取目录内容

* `readdir`函数

  ```c
  #include <sys/types.h>
  #include <dirent.h>
  ```

  * `opendir`函数返回为一个目录流指针

  * `readdir`函数返回`dirent`的下一个指针，使用：

    ```c
    dep = readdir(streamp))!=NULL
    ```

  * `closedir`函数关闭流

### 10.8 共享文件

* 内核用三个数据结构来表示打开的文件
  * 描述符表，每个进程都有
  * 文件表，所有进程共享一张表
  * v-node表，所有进程共享
* 特殊情况：
  * open一个文件多次->此时不同的描述符对应不同的位置
  * 父子进程open同一个文件->此时父子进程共享文件位置

### 10.9 I/O重定向

* `dup2`函数（两个参数`oldfd`和`newfd`）
  * 将`oldfd`复制到`newfd`中
  * 此时newfd将共享oldfd的位置

### 10.10 标准I/O

* 标准I/O库将一个打开的文件模型化为一个流（指向FILE类型的结构的指针）
  * 流是对文件描述符和流缓冲区的抽象
  * 一个例子：`getc`函数
    * 调用`getc`函数
    * 返回文件下一个字符
    * 库调用`read`填充流缓冲区
    * 缓冲区将该字节返回给应用程序
    * 继续反复调用`getc`
    * 达到目的->只要缓冲区还有维度的字节，就能直接从流缓冲区得到数据服务
* `stdin`，`stdout`，`stderr`

### 10.11 I/O函数的使用

* 尽量使用标准I/O
* 不要使用`scanf`等为读取文本文件设计的函数来读取二进制文件
* 对网络套接字的I/O使用RIO函数库：原因->有限制
  * 跟在输出函数之后的输入函数，`fflush`等清空缓冲区的函数调用很重要
  * 跟在输入函数之后的输出函数：`fseek`等函数的调用很重要
  * 套接字对`lseek`函数是非法的，所以限制二出现了问题
    * 唯一办法是将读，写的流分开使用
    * 关闭第二个流的操作会失败造成内存泄漏
    * 这个问题再现成话的程序中会导致问题
  * 格式化使用`sprintf`

## 第十一章 网络编程

### 11.1 客户端-服务器编程模型

* 每个网络应用都是基于这个模型
* 一个应用：一个服务器**进程**，一个/多个客户端**进程**组成(而不是机器或者主机)
* 基本操作是事务
* 客户端-服务器事务的四步
  * 客户端在需要时向服务器发送**请求**->发起事务
  * 服务器收到请求，做出相应操作
  * 服务器发送**响应**给客户端，等待下一个请求
  * 客户端收到响应，处理响应
* 这里的事务仅仅是客户端和服务器执行的一系列步骤

### 11.2 网络

