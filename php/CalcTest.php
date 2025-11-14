<?php

function add($a, $b) {
    return $a + $b;
}

function sub($a, $b) {
    return $a - $b;
}

function mul($a, $b) {
    return $a * $b;
}

function div($a, $b) {
    if ($b == 0) {
        throw new Exception("division by zero");
    }
    return $a / $b;
}

// Это ваш тестовый код
assert(add(2, 3) === 5, "add test failed");
assert(sub(5, 3) === 2, "sub test failed");
assert(mul(4, 5) === 20, "mul test failed");
assert(div(10, 2) === 5, "div test failed");

echo "All tests passed!\n";

?>