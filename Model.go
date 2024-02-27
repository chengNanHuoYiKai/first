package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/font"

	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	. "math/rand"
	"time"
)

//声明我方飞机变量
var   ship *Object=nil
var oneFont           font.Face

//初始化我方飞机对象
func shipAdd(){
	if ship==nil {
		var myplane Object
		myplane.init("myplane",267,750)
		ship=&myplane
	}
}
//飞机根据键盘移动
func shipMove(){
	//如果按下键盘左方向键我方飞机X减小 画面显示向左移动
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {

		ship.X -= 8
		if ship.X < 0 {//防止飞机飞出窗口左边界
			ship.X = 0
		}//如果按下键盘右方向键我方飞机X增大 画面显示向右移动
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		ship.X+= 8
		if ship.X >= 534 {//防止飞机飞出窗口右边界
			ship.X = 534
		}//如果按下键盘上方向键我方飞机Y坐标增大 画面显示向上移动
	}  else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		ship.Y -= 8
		if ship.Y <= 0 {//防止飞机飞出窗口上边界
			ship.Y = 0
		}//如果按下键盘下方向键我方飞机Y减小 画面显示向下移动
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		ship.Y += 8
		if ship.Y >= 750 {//防止飞机飞出窗口下边界
			ship.Y = 750
		}

	}
}
var enemy1List []*Object//声明敌机1的切片用于存储敌机1的对象
func enemy1Add() {
	Seed(time.Now().UnixNano())//是一个用于设置随机数种子
	random2 := Intn(200) // 生成0到200之间的随机数
	if (int)(random2) == 13 {//当生成的随机数为13时添加敌方飞机，也就是说有1/200的概率添加敌方飞机1
		var enemy1 Object
		random1 := Intn(550) // 生成0到550之间的随机数  确定敌机生成位置
		enemy1.init("enemy_new1",float64(random1),0)
		enemy1List = append(enemy1List,&enemy1 )
	}

}
func enemy1Move(){
	for _,enemy1 := range enemy1List {//遍历敌方第一钟飞机的切片
	       enemy1.Y++//使敌方飞机的Y坐标增加表现为敌方飞机向下移动
	}
}
func objectSliceDraw(screen *ebiten.Image, sl []*Object){//当想画出第一种飞机时，sl参数传为ennemy1List
	for _, sli := range sl {//遍历切片
		var op_slice ebiten.DrawImageOptions
		op_slice.GeoM.Translate(sli.X, sli.Y)//根据切片元素的x,y坐标来确定切片的位置
		screen.DrawImage(sli.image, &op_slice)//画出元素的图片

	}
}
var bulletList []*Object //定义存储子弹的切片
func bulletAdd() {

	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {//如果按下空格键就进入产生子弹的代码块
		//	按下空格子弹根据飞机确定位置
		var bullet1 Object

		if ship.HP >= 5 {//当我方飞机的血量大于等于5时 飞机一次发射两发子弹
			var bullet2 Object
			bullet1.init("bulletplus", ship.X+16, ship.Y)
			bullet2.init("bulletplus", ship.X+1, ship.Y)
			bulletList = append(bulletList, &bullet1)
			bulletList = append(bulletList, &bullet2)
		} else {//反之就发射一发子弹Y
			bullet1.init("bulletplus", ship.X+1, ship.Y-80)//ship.X+1  ,ship.Y-80 是子弹的x y 坐标使子弹恰好在飞机中间生成
			//子弹的坐标为 (ship.X-bullet.X)/2   ship.Y-bullet.Y
			bulletList = append(bulletList, &bullet1)//将新生成的飞机子弹添加到切片bulletList里
		}

	}

}
func bulletMove(){
	for _,bullet := range bulletList{
		bullet.Y-=5
	}
}
var bg *Object//定义背景的对象

func bgInit(){//背景产生的方法
	if bg==nil {
		var object Object
		object.init("bg",0,-700)
		bg=&object
	}
}
func bgMove() {//背景循环向下移动 营造我方飞机一直向前飞的效果
	bg.Y++
	if bg.Y >= -10 {
		bg.Y = -790
	}

}
var enemy2List []*Object//敌机2的切片
var enemy2Bullet []*Object//敌机2子弹的切片
var count=0//记录游戏的刷新次数
// 添加第二种敌机
func enemy2Add() {

	Seed(time.Now().UnixNano())
	random2 := Intn(600) // 生成0到1000之间的随机数
	if (int)(random2) == 13 {
		//	fmt.Println("画出了敌机2")
		var e Object
		random3 := Intn(300)//随机产生敌机2的x坐标
		e.init("enemy_new2", float64(random3), 120)
		enemy2List = append(enemy2List, &e)
	}

}

// 给第二种敌机添加子弹
func enemy2BulletAdd(){
	if count%60 == 0 {
		count++
		for _, e := range enemy2List {

			var enemybullet2 Object
			enemybullet2.init("enemybullet3", e.X+15, e.Y+67)
			enemy2Bullet = append(enemy2Bullet, &enemybullet2)
		}
	}
}

func enemy2BulletMove() {
	for _, b := range enemy2Bullet {
		b.Y += 1.5
	}
}
var boss2List []*Object//boss2的切片
var boss2Bullet []*Object //boss2的子弹切片
func boss2Add() {
	Seed(time.Now().UnixNano())
	if count%1000 == 0 &&count!=0{ //20秒左右产生一个boss2
		count++
		random4 := Intn(400) // 生成0到200之间的随机数
		var boss2 Object
		boss2.init("boos_new4", float64(random4), 112)
		boss2List = append(boss2List, &boss2)
	}
}
func boss2BulletAdd() {
	if count%200 == 0 {
		//生成了子弹
		count++
		for _, v := range boss2List {
			var bs2_bullet Object
			bs2_bullet.init("enemy_bullet2", v.X+75.5, v.Y+112)
			boss2Bullet = append(boss2Bullet, &bs2_bullet)
		}

	}
}

func boss2BulletMove() {
	for i, b := range boss2Bullet {

		//改变子弹坐标
		if ship.Y >= b.Y {//子弹会追踪飞机
			b.Y += 2
		} else {
			b.Y -= 2
		}
		if ship.X+float64(ship.width/2) >= b.X {
			b.X += 2
		} else {
			b.X -= 2
		}
		if b.Y+float64(b.height)-20 > ship.Y {//如果子弹在飞机的下方 子弹消失
			boss2Bullet = append(boss2Bullet[:i], boss2Bullet[i+1:]...)
			break
		}
	}
}
var boss1 *Object//定义大boss对象
var boss1Bullet []*Object//定义boss1子弹的切片
var score=0 //定义游戏分数
var boss1_speed=3.0
func Boss1Add() {
	if score >= 10 && boss1 == nil {
		Seed(time.Now().UnixNano())
		var boss Object
		x := Intn(359) + 1
		boss.init("boos_new1", float64(x), 10)
		boss1 = &boss
	}
	if boss1 != nil {
		//改变boss1的坐标
		boss1.X += boss1_speed
		if boss1.X <= 0 || boss1.X >= 360 {
			boss1_speed = -boss1_speed
		}

	}
}

// 添加boss1子弹
func boss1BulletAdd() {

	if count%70 == 0 && boss1 != nil {
		count++
		var bos1Bullet Object
		bos1Bullet.init("bullet_new2", boss1.X+110, boss1.Y+174)
		boss1Bullet = append(boss1Bullet, &bos1Bullet)
	}
}
func boss1BulletMove() {

	for i, bos1 := range boss1Bullet {
		bos1.Y += 2.0

		if bos1.Y > 800 {
			boss1Bullet = append(boss1Bullet[:i], boss1Bullet[i+1:]...)
			break
		}

	}

}
var  state             = true
var  start             = false
func pause() {
	if inpututil.IsKeyJustReleased(ebiten.KeyQ) {
		fmt.Println(state)
		state = !state

	}
}


func startGame() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		start = true
		ship=nil
		shipAdd()
		score = 0
		count = 0
		boss2List = boss2List[:0]

		enemy2Bullet = enemy2Bullet[:0]
		boss2Bullet = boss2Bullet[:0]
		boss1 = nil

		enemy2List = enemy2List[:0]
		boss1Bullet = boss1Bullet[:0]
		enemy1List = enemy1List[:0]
		bulletList = bulletList[:0]
	}
}

// 游戏开始界面
func startBg(screen *ebiten.Image) {
	var start Object
	start.init("start1", 0, 0)
	var op_bullet ebiten.DrawImageOptions
	op_bullet.GeoM.Translate(start.X, start.Y)
	screen.DrawImage(start.image, &op_bullet)
}

// 游戏结束
func gameOver(screen *ebiten.Image) {
	var over Object
	over.init("失败01", 22, 150)
	var op_bullet ebiten.DrawImageOptions
	op_bullet.GeoM.Translate(over.X, over.Y)
	screen.DrawImage(over.image, &op_bullet)

}
func victory(screen *ebiten.Image) {
	var victory Object
	victory.init("奖杯01", 50, 115)
	var op ebiten.DrawImageOptions
	op.GeoM.Translate(victory.X, victory.Y)
	screen.DrawImage(victory.image, &op)
}
