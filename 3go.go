// ��ӡ�����еġ�ˮ�ɻ�����,��ν��ˮ�ɻ�������ָһ����λ��,���λ
// ���������͵��ڸ�����������:153 ��һ����ˮ�ɻ�����,��Ϊ 153=1 �����η�
// +5 �����η�+3 �����η���
package main

import (
	"fmt"
)

func main() {
	for i := 100; i < 999; i++ {
		baiwei := i / 100
		shiwei := (i - baiwei*100) / 10
		gewei := i % 10
		if i == baiwei*baiwei*baiwei+shiwei*shiwei*shiwei+gewei*gewei*gewei {
			fmt.Println(i)
		}
	}
}
