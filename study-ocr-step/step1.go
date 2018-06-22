package main

import(
	"fmt"
	"os"
	"image/png"
	"image"
	"image/color"
	"sort"
	"strconv"
)
//定义最大处理图片像素
const MAX_LENGTH  = 1000
//定义图片的长度和高度
var dx,dy int
//定义图片灰度的二维数组
var graySet [1000][1000]int
//定义二值化后图片图片的情况，0代表背景，1代表目标
var binarizationSet [1000][1000]int
var binarizationSetCopy [1000][1000]int

func main() {
	//1.图片灰度化
	newPath := Gray("ocr.png")

	//2.对图片二值化
	newPath = Binarization(newPath)
	//3.对图片进行长线剔除
	//4.平滑中值去噪
	newPath = Denoising(newPath)
	//5.图片细化
	//Refinement(newPath)
	fmt.Println("图片宽:"+strconv.Itoa(dx)+",高:"+strconv.Itoa(dy));
}

/**
 * 加权平均法 　　
 * 根据重要性及其它指标，将三个分量以不同的权值进行加权平均。
 * 由于人眼对绿色的敏感最高，对蓝色敏感最低，
 * 因此，按下式对RGB三分量进行加权平均能得到较合理的灰度图像。 　　
 * f(i,j)=0.30R(i,j)+0.59G(i,j)+0.11B(i,j))
 */
func Gray(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	pngFile, _ := png.Decode( file )
	bounds := pngFile.Bounds()
	dx = bounds.Max.X
	dy = bounds.Max.Y
	gray := image.NewGray(image.Rect(0, 0, dx, dy))
	var r32,g32,b32 uint32
	var r8,g8,b8 float32
	var newGray byte
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			//计算该点的灰度
			r32, g32, b32, _  = pngFile.At(x, y).RGBA()
			r8,g8,b8 = float32(r32>>8), float32(g32>>8), float32(b32>>8)
			newGray = uint8( 0.299 * r8 + 0.587 * g8 + 0.114 * b8 )
			fmt.Println(strconv.Itoa(x)+":"+strconv.Itoa(y))
			graySet[x][y] = int(newGray)
			gray.Set(x, y, color.Gray{newGray})
		}
	}
	newPath := "b.png"
	newFile, _ := os.Create(newPath)
	png.Encode(newFile, gray)
	return newPath
}

//图片二值化，阀值固定为180,固定<span style="font-family: Arial, Helvetica, sans-serif;">阀值，可以根据具体实验结果设置</span>

func Binarization(path string) string {
	newPath := "c.png"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	pngFile, _ := png.Decode( file )
	gray := image.NewGray(image.Rect(0, 0, dx, dy))
	var r32,g32,b32 uint32
	var r,g,b int
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			//计算该点的灰度
			r32, g32, b32, _  = pngFile.At(x, y).RGBA()
			r,g,b = int(r32>>8), int(g32>>8), int(b32>>8)
			if ( (r+g+b)/3 > 180 ) {
				binarizationSet[x][y] = 0
				gray.Set(x, y, color.Gray{uint8(255)})
			} else {
				binarizationSet[x][y] = 1
				gray.Set(x, y, color.Gray{uint8(0)})
			}
		}
	}
	newFile, _ := os.Create(newPath)
	png.Encode(newFile, gray)
	return newPath
}

//平滑中值去噪，3*3像素方阵中找灰度中值
func Denoising(path string) string {
	var middleGray uint8
	newPath := "d.png"
	gray := image.NewGray(image.Rect(0, 0, dx, dy))
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			middleGray = getMiddleGray( x, y )
			gray.Set(x, y, color.Gray{middleGray})
		}
	}
	newFile, _ := os.Create(newPath)
	png.Encode(newFile, gray)
	return newPath
}


//获取该点的灰度中值，如果数量为偶数，则去中间两个值得平均，否者直接去的中值
func getMiddleGray( x int, y int) uint8 {
	grayWindow := []int{0,0,0,0,0,0,0,0,0}
	if(x!=0 && y!=0 && x!=dx-1 && y!=dy-1) {
		grayWindow[0] = graySet[x-1][y-1]
		grayWindow[1] = graySet[x][y-1]
		grayWindow[2] = graySet[x+1][y-1]
		grayWindow[3] = graySet[x-1][y]
		grayWindow[4] = graySet[x][y]
		grayWindow[5] = graySet[x+1][y-1]
		grayWindow[6] = graySet[x-1][y+1]
		grayWindow[7] = graySet[x][y+1]
		grayWindow[8] = graySet[x+1][y-1]
		sort.Ints(grayWindow)
		return uint8(grayWindow[4])
	} else {
		return uint8(graySet[x][y])
	}
}

//Hilditch算法, 见http://www.cnblogs.com/xiaotie/archive/2010/08/12/1797760.html
//（I）：p 为1，即p不是背景；
//（2）：x1,x3,x5,x7不全部为1（否则把p标记删除，图像空心了）；
//（3）：x1~x8 中，至少有2个为1（若只有1个为1，则是线段的端点。若没有为1的，则为孤立点）；
//（4）：p的8连通联结数为1；
///      计算八联结的联结数，计算公式为：
///      (p6 - p6*p7*p0) + sigma(pk - pk*p(k+1)*p(k+2)), k = {0,2,4)
//（5）假设x3已经标记删除，那么当x3为0时，p的8联通联结数为1；
//（6）假设x5已经标记删除，那么当x5为0时，p的8联通联结数为1。
// x4 x3 x2
// x5 p  x1
// x6 x7 x8
//
// p3 p2 p1
// p4 p  p0
// p5 p6 p7
func Hilditch(x int, y int) int {
	neighor := []int{0,0,0,0,0,0,0,0}
	if(x!=0 && y!=0 && x!=dx-1 && y!=dy-1) {
		neighor[0] = binarizationSet[x+1][y]  //x1
		neighor[1] = binarizationSet[x+1][y-1] //x2
		neighor[2] = binarizationSet[x][y-1] //x3
		neighor[3] = binarizationSet[x-1][y-1] //x4
		neighor[4] = binarizationSet[x-1][y] //x5
		neighor[5] = binarizationSet[x-1][y+1] //x6
		neighor[6] = binarizationSet[x][y+1] //x7
		neighor[7] = binarizationSet[x+1][y+1] //x8
		if( binarizationSetCopy[x][y-1] == 1 && binarizationSet[x][y-1] == 0 ) {
			neighor[2] = 1
		}
		if( binarizationSetCopy[x-1][y] == 1 && binarizationSet[x-1][y] == 0 ) {
			neighor[4] = 1
		}
		number := 0
		for i:= 0 ; i<8 ; i++ {
			if neighor[i] == 1 {
				number ++
			}
		}
		//获取联通连接数
		unicomConnection := unicomConnectionCalc( neighor )
		//条件2
		if neighor[0]*neighor[2]*neighor[4]*neighor[6] == 0 {
			//条件3
			if number >= 2 {
				//条件4
				if unicomConnection == 1 {
					//条件5
					if( binarizationSetCopy[x][y-1] == 1 && binarizationSet[x][y-1] == 0 ) {
						neighor[2] = 0
						unicomConnection := unicomConnectionCalc( neighor )
						if( unicomConnection != 1 ) {    //条件不满足
							return 1
						}
						neighor[2] = 1
					}
					//条件6
					if( binarizationSetCopy[x-1][y] == 1 && binarizationSet[x-1][y] == 0 ) {
						neighor[4] = 0
						unicomConnection := unicomConnectionCalc( neighor )
						if( unicomConnection != 1 ) {    //条件不满足
							return 1
						}
					}
					return 0
				}
			}
		}
	}
	return 1
}


//Hilditch算法,用于获取联通数
func unicomConnectionCalc( neighor []int ) int {
	unicomConnection := neighor[6]-neighor[6]*neighor[7]*neighor[0]
	unicomConnection = unicomConnection + neighor[0]-neighor[0]*neighor[1]*neighor[2]
	unicomConnection = unicomConnection + neighor[2]-neighor[2]*neighor[3]*neighor[4]
	unicomConnection = unicomConnection + neighor[4]-neighor[4]*neighor[5]*neighor[6]
	return unicomConnection
}