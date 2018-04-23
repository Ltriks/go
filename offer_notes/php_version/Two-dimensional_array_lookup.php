<?php

function Find($target, $array)
{
    if($array == null || !is_int($target) ){
        return false;
    }
    $m = count($array[0]);
    $n = count($array);
    $c = $n-1;
    $r = 0;
    while ($c >= 0 && $r <= $m-1) {
        if ($target == $array[$r][$c]) {
            return true;
        }else if ($target > $array[$r][$c]) {
            $r++;
        }else if ($target < $array[$r][$c]) {
            $c--;
        }
    }
    return false;
}

$arr = [
    [1, 4, 7, 11, 15],
    [2, 5, 8, 12, 19],
    [3, 6, 9, 16, 22],
    [10, 13, 14, 17, 24],
    [18, 21, 23, 26, 30],
];
var_dump(Find(8,$arr));