package main

import (
	"math"
)
//两个物体碰撞的方法
func hit(MyList *[]*Object, enemyList *[]*Object) {//假如传的参数为&bulletList  &enemy1List
	for i, my := range *MyList {//遍历我方子弹的切片
		for j, enemy := range *enemyList {//遍历敌方飞机1的切片
			/*
			（x1,y1）为某个子弹的中心坐标  （x1,y2）为某个敌方飞机的中心坐标
			 */
			x1 := my.X + float64(my.width)/2
			y1 := my.Y + float64(my.height)/2
			x2 := enemy.X + float64(enemy.width)/2
			y2 := enemy.Y + float64(enemy.height)/2
			//判断是否碰撞的条件
			if math.Abs(x2-x1) < float64(my.width+enemy.width)/2 && math.Abs(y2-y1) < float64(my.height+enemy.height)/2 {
				enemy.HP--
				if enemy.HP<=0{
					*enemyList = append((*enemyList)[:j], (*enemyList)[j+1:]...)//发生碰撞将该飞机从切片中移除

				}
				*MyList = append((*MyList)[:i], (*MyList)[i+1:]...)//发生碰撞后将该子弹从切片中移除
				score++
				return
			}
		}

	}
}
//写飞机（我方飞机或boss1）与其他子弹的碰撞

func planeHit(plane *Object,slice *[]*Object){

	for j, sl := range *slice {//遍历敌方飞机1的切片
		/*
			（x1,y1）为某个子弹的中心坐标  （x1,y2）为某个敌方飞机的中心坐标
		*/
		x1 := plane.X + float64(plane.width)/2
		y1 := plane.Y + float64(plane.height)/2
		x2 := sl.X + float64(sl.width)/2
		y2 := sl.Y + float64(sl.height)/2
		//判断是否碰撞的条件
		if math.Abs(x2-x1) < float64(plane.width+sl.width)/2 && math.Abs(y2-y1) < float64(plane.height+sl.height)/2 {
			*slice = append((*slice)[:j], (*slice)[j+1:]...)//发生碰撞将该元素从切片中移除
		    plane.HP--
			if plane.HP<=0 {
				start=!start
			}

			return
		}
	}


}