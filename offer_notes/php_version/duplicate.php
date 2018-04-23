<?php 	

function duplicate(array $nums)
{
	if ($nums == []) {
		return false;
	}
	for ($i = 0; $i < count($nums); $i++) { 
		while ($nums[$i] != $i && $nums[$i] != $nums[$nums[$i]]) {
			swap($i, $nums[$i], $nums);
		}
		if ($nums[$i] != $i && $nums[$i] == $nums[$nums[$i]]) {
			return true;
		}
	}
}

function swap(int $i, int $j, array &$nums){ // 地址& 不加原数组不改变
	$tmp = $nums[$i];
	$nums[$i] = $nums[$j];
	$nums[$j] = $tmp;
}

var_dump(duplicate([2, 3, 1, 0, 2, 5]));