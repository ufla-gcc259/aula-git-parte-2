#include <stdio.h>
#include "tempconv.h"

int main(){

    printf("Que frio!: %.2f\n", AbsoluteZeroC);
    printf("Fervendo!: %s\n", formatFahrenheit(CToF((float)BoilingC)));
}