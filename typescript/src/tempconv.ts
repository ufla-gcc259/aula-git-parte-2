import { Celsius } from './scales/Celsius';
import { Fahrenheit } from './scales/Fahrenheit';

export const ABSOLUTE_ZERO_C = new Celsius(-273.15);

export const FREEZING_C = new Celsius(0);

export const BOILING_C = new Celsius(100);

export const celsius_to_fahrenheit = (c: number) => new Fahrenheit((c * 9) / 5 + 32);

export const fahrenheit_to_celsius = (f: number) => new Celsius(((f - 32) * 5) / 9);
