# golang 后台项目demo
---

### v2.2.0
- 新增library/routine/封装协程相关库功能
- 在library/routine下新增Group，在标准实验库errgroup基础上提供协程数量控制功能
- 基本使用示例
- ```go
    // 创建一个最大并发数为1000的协程分组
    group := routine.NewGroup(1000)
    for i := 0; i < 10000; i++ {
        j := i
        // 启动协程执行打印，受最大协程数量限制，当分组运行中的协程数量达到最大限制时，
        // 该语句将等待已占用协程的释放，协程的占用和释放在封装内部自动完成
        group.Go(func() error {
            fmt.Println(j)
            return nil
        });
    }
    // Wait将阻塞等待所有已启动的协程完全退出
    group.Wait()

- NewGroupWithContext可以包装传入的context，实际用法和errgroup相同

---

### v2.1.2
- 优化library/http库客户端封装实现，防止潜在的响应体未关闭导致的资源泄露风险
- 新加业务示例代码Hello，示例http客户端基本用法

---

### v2.1.1
- 修复library/http库客户端封装在返回状态码不为200时响应体不被关闭导致的连接泄露问题

---

### v2.1.0
- 更新library/http库客户端封装实现

---

### v2.0.2
- 修复library/clean在初始化监听时阻塞主进程的BUG

---

### v2.0.1
- 修复library/conf.Container关闭时可能触发的死锁问题
- 重写library/clean: 简化实现&新增对清理函数的支持

---

### v2.0.0
- 调整library/mysql: 缓存并重用prepare，提升sql执行效率，__兼容现有业务代码__
- 调整library/log及使用范例，拆分错误和信息日志，取消轮转支持（可由程序外部工具实现）
- 其它library包优化和调整
