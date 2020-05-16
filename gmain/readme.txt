1. cgo只能从main包中到处函数并生成头文件
2. 将go代码编译成静态链接库: go build -buildmode=c-archive -o ../lib/my_gfunc.a
3. 将go代码编译成动态链接库: go build -buildmode=c-shared -o ../lib/libmy_gfunc.so