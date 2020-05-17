1. 将c/c++代码编译动态链接库
step 1: 编译: gcc -c -fPIC -o ../lib/my_cfunc.o my_cfunc.c
step 2: 生成共享库: gcc -shared -o ../lib/libmy_cfunc.so ../lib/my_cfunc.o
2. 编译时使用so库: gcc -o main main.c -I ../include/ -lmy_cfunc -lmy_gfunc -lmy_cppfunc -L../lib/
3. 运行可执行文件时如需用到so库，需在通过环境变量LD_LIBRARY_PATH指定so库路径
4. cgo只能从main包中到处函数并生成头文件
5. 将go代码编译成静态链接库: go build -buildmode=c-archive -o ../lib/my_gfunc.a
6. 将go代码编译成动态链接库: go build -buildmode=c-shared -o ../lib/libmy_gfunc.so