// ��һ���������ֽ�������������:���� 90,��ӡ�� 90=2*3*3*5��
package main

import (
	"fmt"
)

func main() {
	var n, i int
	fmt.Printf("\nPlease input a number:\n")
	fmt.Scanf("%d", &n)
	fmt.Printf("%d=", n)
	for i = 2; i <= n; i++ {
		for n != i {
			if n%i == 0 {
				fmt.Printf("%d*", i)
				n = n / i
			} else {
				break
			}
		}
	}
	fmt.Printf("%d\n", n)
}
