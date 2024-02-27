package main

import (
	//"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)
//planeHp 根据飞机名称映射飞机的生命值
var planeHp = map[string]int{
	"myplane":3,
	 "boos_new1":10,
	"boos_new4":5,
}
type Object struct {
	name   string   //表示对象的名称
	image  *ebiten.Image //表示对象对应的图片
	X      float64//对象的X坐标
	Y      float64//对象的Y坐标
	HP   int//对象的生命值
	width  int//对象的宽度  将来用于碰撞检测
	height int//对象的高度
}
func (object *Object) init(name string, X, Y float64) {//对象的初始化方法
	object.name = name
	object.X = X
	object.Y = Y
	object.HP = planeHp[object.name]//从planeHp map集合里映射飞机对象的生命值
	object.getImg()//调用getImg方法初始化对象的图片

}
func (object *Object) getImg() {
	//使用ebitenutil.NewImageFromFile函数从文件系统中加载图像，并将图像赋值给对象的image字段
	img, _, err1 := ebitenutil.NewImageFromFile("image/" + object.name+".png")
	object.width, object.height = img.Size()//初始化对象的宽高

	if err1 != nil {
		log.Fatal(err1)
	}
	object.image = img
}

//对象的图片绘画方法
func (object *Object)ObjectDraw(screen *ebiten.Image){
	var op_objecte ebiten.DrawImageOptions//ebiten.DrawImageOptions变量，用于配置绘制图像的选项
	op_objecte.GeoM.Translate(object.X, object.Y)//设置图象的位置
	////使用screen.DrawImage方法将对象的图像绘制到目标图像上，传递了对象的图像object.image和绘制选项&op_objecte
	screen.DrawImage(object.image, &op_objecte)
}
