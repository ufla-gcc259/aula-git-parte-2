<?php

const AbsoluteZeroC = -273.15;
const FreezingC = 0;
const BoilingC = 100;

function Celsius($c) {
    return $c;
}

function Fahrenheit($f) {
    return $f;
}

function CToF($c) {
    return Fahrenheit($c * 9 / 5 + 32);
}

function FToC($f) {
    return Celsius(($f - 32) * 5 / 9);
}

?>
