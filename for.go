package main

import "fmt"


func main() {

	// 省略了循环条件，死循环写法，加了 break 则循环终止，有没有分号 ; 都可以
	for {
		fmt.Println("Hello World")
		break
	}
	for ; ; {
		fmt.Println("Hello World")
		break
	}

	fmt.Println("\r")

	// 经典for循环写法
	for i := 0; i <= 3; i ++ {
		fmt.Println(i)
	}

	fmt.Println("\r")

	// 前置、后置语句为空，sum<=7 左右的 ; 可以不写
	sum := 1
	fmt.Println(sum)
	for ; sum<=7 ; {
		sum += sum
		fmt.Println(sum)
	}

	fmt.Println("\r")

	sum = 1
    for sum <= 2 {
		sum += sum;
		fmt.Println(sum);
	}

	fmt.Println("\r")


}