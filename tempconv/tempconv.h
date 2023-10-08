#include <iostream>
using namespace std;

const float AbsoluteZeroC = -273.15;
const float FreezingC = 0;
const float BoilingC = 100;

float CToF(float c)
{
	return c * 9 / 5 + 32;
}

float FToC(float f)
{
	return (f - 32) * 5 / 9;
}