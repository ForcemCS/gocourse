## 说明

go course的相关内容

###  Go Compiler

一但将代码编译为机器代码，Runtime就会发作用，Runtime是管理程序执行的重要组件，它处理任务，例如：内存分配，垃圾收集(清理未使用的内存)

了解编译器和运行时如何一起工作对于掌握GO语言至关重要

如果go编译器将代码改为机器代码（机器不能直接理解吗？为什么还需要Runtime）

### 标准库
是指全面的package集合（预先编写的代码块），包含了大量的功能

###值拷贝和引用
int array struct是值类型，map slice pointer是引用类型