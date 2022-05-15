# sudygo

    main 函数是作为goroutine 执行
    操作系统 线程 和 goroutine 的关系？
        线程被操作系统调度
        goroutine 实际上是被Go 运行时runtime来调度的，最终这些goroutine 它都可以映射到
        到某一个线程上，执行单元对操作系统来说它不认识goroutine,它只认识操作系统线程
        runtime 实际上就是负责把这些goroutine 调度到某一个Go runtime的一个逻辑处理器（P）
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



    控制反转、依赖注入
    PO 数据表结构
    DO 在PO的基础上逻辑
    BIZ 业务逻辑
        data
            | - PO
        biz
            业务逻辑
            | - Do
    pkg 与业务无关
    task 定时任务
    web http 业务数据格式规范


    protobuf 入门

    unit test 
        assert
    
    mockgen 
        gomock.newController
        expect().Bar()

    test suite

    pprof 工具分析 

    Tdd 测试驱动
        1. 容易编写并且启动测试的情况
        2. 场景复杂
    
 
    隔离
        动静隔离、读写分离
            动静隔离 
                1. cacheline
                2. 数据库mysql 表设计中避免 bufferpool 过期 datapage 可以缓存表的行
                    a. binlog 订阅
            读写分离
                1. 主从、replicaset、CQRS
                 
                    a. 由于CQRS 的本质是对于读写操作的分离，所以比较简单的Cqrs的做法是
                        CQ两端数据库表共享，CQ两端只是在上层代码上分离。
              轻重隔离
                1. 核心/非核心的故障域的差异隔离（机器资源、依赖资源）
               
              快慢隔离
                1. 
    # 故障域
        根据业务的轻重机器的资源是否需要独占和非独占，独占的情况下防止其他微服务的对主业务的影响。
    # topic
    # hystrix
    # resilience4j  
        早期转码集群被超大视频攻击、导致转码大量延迟
        缩略图服务，被大图实时缩略吃完CPU,导致正常的小图无法展示
        数据库实例cgroup未隔离，导致大sql引起的集体故障
        info日志量过大，导致异常error 日志采集延迟。


过载保护
    利特尔法则


如何计算接近峰值时的系统吞吐？
    cup: 使用一个单独的线程采样
    inflight 当前 服务中正在进行的请求的数量
    pass&RT: 最近 5s pass 每100ms采样窗口内成功请求的数量，RT为单个采样窗口中平均响应时间

限流
    限流是指在一段时间内，定义某个客户或者应用可以接收或处理多少个请求的技术（Auto Scaling）

        需要考虑的情况
            1. 令牌桶、漏桶针对单个节点，无法对分布式限流 
            2.QPS 限流
                a. 不同的请求可能需要数量迥异的资源来处理
                b. 某一种静态的QPS 限流不是特别准
            3. 每个用户设置限制
                a. 全局过载发生时候，针对某些“异常”进行控制
                b. 一定程度的“超卖”配额
            4. 按照优先级丢弃
            5. 拒绝请求也需要成本
            
    分布式限流
        redis incr  quota 表示速率

        最大最小公平分享(max-min fairness)

        critical_plus
        critical
        sheddable_plus
        sheddable
    熔断（circuit breakers）

        为了限制操作的持续时间，我们可以使用超时可以防止挂起操作并保证系统的响应
            1. 依赖的资源出现大量的错误
            2. 某个用户超过资源的配额时，后端任务会快速拒绝请求

        CAS 只有一个请求过来

        gutter

        positive feedback 用户总是积极重试，访问一个不达的服务

        jitter
    
    降级
        mttr 平均修复时间
    
    负载均衡
        p2c 算法，随机选取两个节点进行打分，选择更优的节点
        
    最佳实践
    sop
    dirt 载脏实验



    链路超时
        header 中传递go-timeout 进行超时时间向下游传递
        RPC 链路超时控制
            a. 通过RPC协议 timeout 传递超时时间
            b. 中断
            c. 确定超时时间 PM 确定
        业务链路控制超时


    重试 考虑的因素 重试次数 重试间隔 进程内重试跨进程重试
        1. 进程内重试
        2. 跨进程重试
            a. 数据库方案
            b. 分布式任务中心方案：直接存放数据库，定时轮训数据库
                在分布式任务中心注册一个任务，定时任务或者重复任务
            c. 延时消息方案: 发送一个延时消息给MQ,MQ会在约定的时间投递给消费者，消费者执行重试
        
    #### 布隆过滤器

    mysql binlog 同步技术 cannal

    ## 缓存模式 cache aside

        1. cache aside
        2. 
    
    ### redis sord set
        lazy (懒加载)
    

    可用性
        singleflight dns 回源