package main

import (
	"fmt"
)

//有两个类型 stockPosition 和 car，它们都有一个 getValue() 方法，
//我们可以定义一个具有此方法的接口 valuable。接着定义一个使用 valuable 类型作为参数的函数 showValue()，
//所有实现了 valuable 接口的类型都可以用这个函数。

type stockPosition struct {
	ticker     string  // 股票代码
	sharePrice float32 // 分享价格
	count      float32 // 数量
}

/*method to determine the value of a stock position*/
func (sp stockPosition) getValue() float32 {
	return sp.sharePrice * sp.count
}

type car struct {
	make  string  // 品牌
	model string  // 型号
	price float32 // 价格
}

/*method to determine the value of a car*/
func (c car) getValue() float32 {
	return c.price
}

/*defines a interface */
type valuable interface {
	getValue() float32
}

func showValue(value valuable) {
	// asset 资产
	fmt.Printf("Value of the asset is %f\n", value.getValue())
}

func main() {
	var o valuable = stockPosition{ticker: "Good", sharePrice: 355.93, count: 20}
	showValue(o)

	o = car{make: "BMW", model: "M3", price: 7888000}
	showValue(o)
}
