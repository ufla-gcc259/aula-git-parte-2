<?php

const BoilingC = 100;
const FreezingC = 0;
const AbsoluteZeroC = -273.15;




function Fahrenheit($f) {
    return $f;
}

function Celsius($c) {
    return $c;
}

function FToC($f) {
    return Celsius(($f - 32) * 5 / 9);
}

function CToF($c) {
    return Fahrenheit($c * 9 / 5 + 32);
}



?>