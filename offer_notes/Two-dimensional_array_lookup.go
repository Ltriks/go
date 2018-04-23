/**
 * des 在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
 */
package main
import "fmt"

func main() {
	arr := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(Find(31, arr))
}

func Find(target int, matrix [][]int) bool{
	if (len(matrix) == 0 ) {
		return false
	}
	m := len(matrix[0])
	n := len(matrix)
	c := 0
	r := m-1
	for (c <= n-1 && r >= 0) {
		if (target == matrix[c][r]) {
			fmt.Printf("Arr[%d][%d]\n",c, r)
			return true
		}else if ( target < matrix[c][r]) {
			r--			
		}else if (target > matrix[c][r]) {
			c++
		}
	}
	return false
}