package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/draw"
	"image/jpeg"
	"net/http"
	"os"
	"strconv"
	"time"
	// "runtime"
)

func main() {
	// runtime.GOMAXPROCS(runtime.NumCPU())
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", upload)
	mux.HandleFunc("/mosaic", mosaic)
	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	// building up the source tile database
	TILESDB = tilesDB()
	fmt.Println("Mosaic server started.")
	server.ListenAndServe()
}

func upload(w http.ResponseWriter, r *http.Request) {
	path, _ := os.Executable()
	fmt.Println(path)
	t, err := template.ParseFiles("/Users/lyuxiyang/Documents/gitProjectgo/study_go/go_web/chap9_concurrency/mosaic/upload.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, nil)
}

//处理器函数，⾸先，程序会获取⽤户上传的⽬标图⽚，并从表单中获取瓷砖图⽚的尺⼨；接着，程序会对⽬标图⽚进⾏解码，并创建出⼀张全新的、
//空⽩的⻢赛克图⽚；之后，程序会复制⼀份瓷砖图⽚数据库，并为每张瓷砖图⽚设置起始点（source point），
//⽽这⼀起始点将在稍后的代码中被image/draw 包所使⽤。在完成了上述的准备⼯作之后，程序就可以开始对⽬标图⽚分割出的各张瓷砖图⽚尺
//⼨的⼦图⽚进⾏迭代了。
func mosaic(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()
	// get the content from the POSTed form
	r.ParseMultipartForm(10485760) // max body in memory is 10MB
	//获取⽤户上传的⽬标图⽚，以及瓷砖图⽚的尺⼨
	file, _, _ := r.FormFile("image")
	defer file.Close()
	tileSize, _ := strconv.Atoi(r.FormValue("tile_size"))
	// decode and get original image 对⽤户上传的⽬标图⽚进⾏解码
	original, _, _ := image.Decode(file)
	bounds := original.Bounds()
	// create a new image for the mosaic
	newimage := image.NewNRGBA(image.Rect(bounds.Min.X, bounds.Min.X, bounds.Max.X, bounds.Max.Y))
	// build up the tiles database 复制瓷砖图数据库
	db := cloneTilesDB()
	// source point for each tile, which starts with 0, 0 of each tile 对⽤户上传的⽬标图⽚进⾏解码
	sp := image.Point{0, 0}
	//对⽬标图⽚分割出的每张⼦图进⾏迭代
	for y := bounds.Min.Y; y < bounds.Max.Y; y = y + tileSize {
		for x := bounds.Min.X; x < bounds.Max.X; x = x + tileSize {
			// use the top left most pixel as the average color
			r, g, b, _ := original.At(x, y).RGBA()
			color := [3]float64{float64(r), float64(g), float64(b)}
			// get the closest tile from the tiles DB
			nearest := nearest(color, &db)
			file, err := os.Open(nearest)
			if err == nil {
				img, _, err := image.Decode(file)
				if err == nil {
					// resize the tile to the correct size
					t := resize(img, tileSize)
					tile := t.SubImage(t.Bounds())
					tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)
					// draw the tile into the mosaic
					draw.Draw(newimage, tileBounds, tile, sp, draw.Src)
				} else {
					fmt.Println("error:", err, nearest)
				}
			} else {
				fmt.Println("error:", nearest)
			}
			file.Close()
		}
	}

	buf1 := new(bytes.Buffer)
	//将图⽚编码为JPEG 格式，然后通过base64字符串将其传输⾄浏览器
	jpeg.Encode(buf1, original, nil)
	originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	buf2 := new(bytes.Buffer)
	jpeg.Encode(buf2, newimage, nil)
	mosaic := base64.StdEncoding.EncodeToString(buf2.Bytes())

	t1 := time.Now()
	images := map[string]string{
		"original": originalStr,
		"mosaic":   mosaic,
		"duration": fmt.Sprintf("%v ", t1.Sub(t0)),
	}
	t, _ := template.ParseFiles("/Users/lyuxiyang/Documents/gitProjectgo/study_go/go_web/chap9_concurrency/mosaic/results.html")
	t.Execute(w, images)

}
