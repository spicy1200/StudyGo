# sudygo
/**
    main 函数是作为goroutine 执行
    操作系统 线程 和 goroutine 的关系？
        线程被操作系统调度
        goroutine 实际上是被Go 运行时runtime来调度的，最终这些goroutine 它都可以映射到
        到某一个线程上，执行单元对操作系统来说它不认识goroutine,它只认识操作系统线程
        go runtime 实际上就是负责把这些goroutine 调度到某一个Go runtime的一个逻辑处理器（P）
        P 像一个队列，把这些任务挂到队列上，然后相当于最终是在各个线程去队列里面捞到一个任务去执行
        通过go runtime 的特性使的我们可以调度数以万计的goroutine,以惊人的效率和能力执行并发运行

        runtime.GomaxProcs  设置逻辑处理器的个数

        go routine
        什么时间会结束
        有没有办法结束它

        log.Fatal 调用os.exit 不会调用 defers
            使用 log.Fatal 只能在 go init 或者 main 函数中

        errgroup


    当下面条件满足时，对变量V的读操作r是被允许看到对V的写操作w的:
        1.r不先行发生与w
        2.在w后r前没有对v的其他写操作
    为了保证对变量v的读操作r看到对v的写操作w,要保证w是r允许看到的唯一写操作。即当下面条件满足时，r被保证看到w:
        1. w先行发生于r
        2.其他对共享变量v的写操作要么在w前，要么在r后
    单个goroutine 中没有并发，所以上面两个定义相同的:

    读操作r看到最近一次的写操作w写入v的值
    当对个goroutine 访问共享变量v时，它们必须使用同步事件来建立先行发生这一条件保证读操作能看到需要的写操作
    （同步时间需要加互斥锁比如atomic机制）
        1. 对变量v的零值初始化在内存模型中表现的与写操作相同
        2. 对大于single machine word 的变量的读写操作表现的像以不确定顺序对多个single machine word 的变量的操作

        single machine word => （一个单一的机器字节）
    map 结构 hmap  cow => copy on write(写实copy) 机制 bgsave

    当一个map 它的存储内容会超出当个 single machine word 导致内部数据复制，回出现问题它不满足可见性 （可见性是指当多个线程访问同一个变量时，一个线程修改了这个变量的值，其他线程能够立即看得到修改的值）
    cow 一个新的进程它可能共享老的进程地址空间然后对里面的数据进行访问刷盘（拿到一个副本）如果有新的东西写进来时拷贝（写时拷贝）有变化的copy进来
        atomic.value  load store
    hmap
    mesi cpu通知核心把老的cache 刷掉

    share(分享) memory by communicationg(通信)


    data race 信息率
            是两个或多个goroutine 访问同一个资源(如变量或者数据结构)，并尝试对资源进行读写操而不考虑其他gorutine.
        这种类型的代码可以创建您见过的最疯狂的随机bug.通常需要大量的日志记录和运气才能找到这些类型的bug
    go build -race
    go test -race 

    go 的内存模型
        写入单个machine word 他是原子的
        但是 interface 的底层原理 interface 它实际上是两个machine word 的值 （type  data ）
    
    sync.atomic

    并发之原子性、可见性、有序性
    原子性
        即一个操作或者多个操作 要么全部执行并且执行的过程不会被任何因素打断，要么就都不执行。
    可见性
        可见性是指当多个线程访问同一个变量时，一个线程修改了这个变量的值，其他线程能够立即看得到修改的值
    有序行
        即程序执行的顺序按照代码的先后顺序执行




    data race （数据竞争）
        copy on write  为了解决 data race 读多写少的并发问题
    redis bgsave

    sync.atomic 
        copy-on-write 思路在微服务降级或者local cache 场景中经常使用。 写时复制指的是，写操作的时候全量老数据到一个新的对象中，
    携带上本次新写的数据，之后利用原子替换（atomic.value）,更新调用者的变量。来完成无锁访问共享数据


    mutex
        barging（冲突） 这种模式是为了提高吞吐量，当锁被释放时，它会唤醒第一个等待这，然后把锁给第一个等待者或者给第一个请求锁的人。
        handsoff 当释放时候，锁会一直持有直到第一个等待者准备好获取锁。它降低了吞吐量，因为锁被持有，即使另外一个goroutine 准备获取它
        spinning 自旋在等待队列为空或者应用程序重度使用锁时效果不错。parking和unparking goroutines 有不低的性能成本开销，相比自旋来说要慢得多

    errgroup
        并行工作流
        错误处理 或者 优雅降级
        context 传播和取消
        利用局部变量+闭包
    
    
    sync.pool 


    go concurrency patterns
        timing out
        moving on
        pipeline
        fan-out fan-in
        cancellation

**/