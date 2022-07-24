#include <stdio.h>
#include "./tempconv/tempconv.h"

int main() {
  printf("Que frio! %.2f°C", AbsoluteZeroC);
  const float boilingTempF = CToF(BoilingC);
  printf("Fervendo! %.2f°F", boilingTempF);
}