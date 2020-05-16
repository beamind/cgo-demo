#include <stdio.h>
#include "../include/my_cfunc.h"
#include "../include/my_gfunc.h"
void main()
{
    int x0 = 2;
    int y0 = 10;
    int z0 = my_add(x0, y0);

    double x1 = (double)x0;
    double y1 = (double)y0;
    double z1 = MyPow(x1, y1);

    int z2 = MyProd(x0, y0);

    printf("使用c编译的so库: %d + %d = %d\n", x0, y0, z0);
    printf("使用go编译的so库: %f ** %f = %f\n", x1, y1, z1);
    printf("使用go编译的so库: %d * %d = %d\n", x0, y0, z2);
}