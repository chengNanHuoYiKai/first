package main

import (
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"strconv"

	//"github.com/hajimehoshi/ebiten/ebitenutil"
	//"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	//"github.com/hajimehoshi/ebiten/text"
	//"golang.org/x/image/font"
	//"golang.org/x/image/font/opentype"
	//
	// */"image/color"
	//. "strconv"
	//
	//
	//"log"
	//
	//"github.com/hajimehoshi/ebiten/v2"
	//
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	_ "golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
	_ "strconv"
)

func init() {

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	oneFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
}
//定义了一个Game结构体 实现了游戏引擎里的Game接口
type Game struct{}



//下面实现了Game结构体里三个方法 Update()  Draw()   Layout()

//Draw()相当于MVC架构里面的V view  用于渲染显示图形画面
func (g *Game) Draw(screen *ebiten.Image) {
	if !start { // start=false 游戏未开始 游戏失败 游戏胜利
		if ship!=nil&&ship.HP<=0 {
			gameOver(screen)
		} else if boss1!=nil&&boss1.HP<=0 {
			victory(screen)
		}else {
			startBg(screen)
		}


	}else {

		if bg!=nil {
			bg.ObjectDraw(screen) //将背景的Draw方法写在Draw的最上面一行，因为在画图时最上面的代码图象在最里面一层

		}

if ship!=nil {
	ship.ObjectDraw(screen)
}


		objectSliceDraw(screen, enemy1List)
		objectSliceDraw(screen, bulletList)
		objectSliceDraw(screen, enemy2List)   //enemy2List 调用Draw方法
		objectSliceDraw(screen, enemy2Bullet) //enemy2Bullet 调用Draw方法
		objectSliceDraw(screen, boss2List)    //boss2List 调用Draw方法
		objectSliceDraw(screen, boss2Bullet)  //boss2Bullet 调用Draw方法
		objectSliceDraw(screen, boss1Bullet)
		if boss1 != nil {
			boss1.ObjectDraw(screen)
		}
		//画出分数
		str := strconv.Itoa(score)
		text.Draw(screen, "score:"+str, oneFont, 10, 20, color.White)

		// text.Draw((screen),"score:"+str, oneFont, 10, 20, color.White)
		//text.Draw(screen, "score:"+str, oneFont, 10, 20, color.White)

}
}
//Update()相当与MVC里面的C control  用于将来控制我们飞机元素的添加和移动
func (g *Game) Update() error {
	if !start {
		startGame()
		return nil
	}
	pause()

	if !state {
		//pause()
		return nil
	}
	count++//每执行一次Update()刷新次数加一
	bgInit()
	shipAdd()
	bgMove()

	shipMove()
	enemy1Add()
	enemy1Move()
	bulletAdd()
	bulletMove()
	enemy2Add()
	enemy2BulletAdd()
    enemy2BulletMove()
	boss2Add()
	boss2BulletAdd()
	boss2BulletMove()
	Boss1Add()
	boss1BulletAdd()
	boss1BulletMove()

	hit(&bulletList,&enemy1List)//传入我方子弹切片的地址  敌方飞机1的切片地址
	hit(&bulletList,&enemy2List)
	hit(&bulletList,&boss2List)
	hit(&bulletList,&boss2Bullet)
	hit(&bulletList,&boss1Bullet)
	planeHit(ship,&enemy2Bullet)
	planeHit(ship,&boss2Bullet)
	planeHit(ship,&boss1Bullet)
	if boss1!=nil {
		planeHit(boss1,&bulletList)
	}

	return nil
}


//设置窗口大小  宽600 高800
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 600, 800
}

func main() {
	//设置窗口大小
	ebiten.SetWindowSize(600, 800)
	//设置窗口的标题
	ebiten.SetWindowTitle("飞机大战")
	/*接着调用 ebiten.RunGame(&Game{}) 启动游戏循环，并将 &Game{} 作为参数传递给 RunGame() 函数
	以每秒60帧循环执行Game接口的 Update()方法 Draw()方法
	，以创建并运行游戏对象。如果在运行游戏过程中出现错误，将通过 log.Fatal() 打印错误信息并终止程序的执行。
	 */
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}