// �ж� 101-200 ֮���ж��ٸ�����,���������������
package main

import (
	"fmt"
	"math"
)

func main() {
	count := 0 //��������
	for i := 101; i <= 200; i++ {
		mid := int(math.Sqrt(float64(i)))
		for j := 2; j <= mid; j++ {
			if i%j == 0 {
				//fmt.Println(i)//���������ļ�Ϊ����
				count = count + 1
				break
			}
		}
	}
	fmt.Println(100 - count)
}
