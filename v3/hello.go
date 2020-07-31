package hello

import (
	"fmt"
	"math/rand"
	"time"
)

func SayHello(language string) {
	switch language {
	case "cn":
		fmt.Println("求知若渴，求贤若愚。")
	case "en":
		fmt.Println("Stay hungry,stay foolish.")
	default:
		fmt.Println("only support cn and en")
	}
}

func Hello() {
	fmt.Println("Hello World!")
}

func randInt64(min, max int64) int64 {
	if min < 0 {
		return 0
	}
	if min >= max {
		return min
	}
	rand.Seed(time.Now().Unix())
	return rand.Int63n(max-min) + min
}

var poem []string = []string{
	`
	君不见，黄河之水天上来，奔流到海不复回。

	君不见，高堂明镜悲白发，朝如青丝暮成雪。

	人生得意须尽欢，莫使金樽空对月。

	天生我材必有用，千金散尽还复来。

	烹羊宰牛且为乐，会须一饮三百杯。

	岑夫子，丹丘生，将进酒，杯莫停。

	与君歌一曲，请君为我倾耳听。

	钟鼓馔玉不足贵，但愿长醉不复醒。

	古来圣贤皆寂寞，惟有饮者留其名。

	陈王昔时宴平乐，斗酒十千恣欢谑。

	主人何为言少钱，径须沽取对君酌。

	五花马，千金裘，呼儿将出换美酒，与尔同销万古愁。
	`,
	`
	如何让你遇见我
	在我最美丽的时刻

	为这
	我已在佛前求了五百年
	求佛让我们结一段尘缘
	佛於是把我化做一棵树
	长在你必经的路旁

	阳光下
	慎重地开满了花
	朵朵都是我前世的盼望

	当你走近
	请你细听
	那颤抖的叶
	是我等待的热情

	而当你终於无视地走过
	在你身後落了一地的
	朋友啊
	那不是花瓣
	那是我凋零的心
	`,
}

func RelaxMoment() {
	var index = randInt64(0, int64(len(poem)))
	fmt.Println(poem[index])
}
