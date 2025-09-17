package test

import (
	"gonum.org/v1/plot/plotutil"
	"log"
	"testing"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type DataPoint struct {
	Day   string
	Value float64
}

func TestPlot(t *testing.T) {
	// 示例数据 - 替换为您的实际数据
	data := []DataPoint{
		{"2023-01-01", 12.5},
		{"2023-01-02", -8.2},
		{"2023-01-03", 5.7},
		{"2023-01-04", -15.3},
		{"2023-01-05", 9.1},
		{"2023-01-06", -3.4},
		{"2023-01-07", 18.9},
		{"2023-01-08", -7.6},
		{"2023-01-09", 4.2},
		{"2023-01-10", -11.8},
	}

	// 创建图表
	p := plot.New()
	p.Title.Text = "数据象限图"
	p.X.Label.Text = "时间"
	p.Y.Label.Text = "数值"

	// 准备绘图数据
	points := make(plotter.XYs, len(data))
	for i, d := range data {
		// 解析字符串时间为time.Time
		t, err := time.Parse("2006-01-02", d.Day)
		if err != nil {
			log.Fatalf("无法解析时间: %v", err)
		}
		// 将时间转换为浮点数以便绘图
		points[i].X = float64(t.Unix())
		points[i].Y = d.Value
	}

	// 创建折线图（连接数据点）
	line, pointsScatter, err := plotter.NewLinePoints(points)
	if err != nil {
		log.Fatal(err)
	}

	// 设置线条和点的样式
	line.Color = plotutil.Color(0) // 使用第一种颜色
	pointsScatter.Color = plotutil.Color(0)
	pointsScatter.Shape = plotutil.Shape(0) // 使用第一种形状

	// 添加到图表
	p.Add(line, pointsScatter)

	// 添加象限线
	// 计算X轴和Y轴的中点
	xMin, xMax := p.X.Min, p.X.Max
	yMin, yMax := p.Y.Min, p.Y.Max
	xMid := (xMin + xMax) / 2
	yMid := (yMin + yMax) / 2

	// 添加垂直象限线
	vLine, _ := plotter.NewLine(plotter.XYs{{xMid, yMin}, {xMid, yMax}})
	vLine.Color = plotter.DefaultLineStyle.Color
	p.Add(vLine)

	// 添加水平象限线
	hLine, _ := plotter.NewLine(plotter.XYs{{xMin, yMid}, {xMax, yMid}})
	hLine.Color = plotter.DefaultLineStyle.Color
	p.Add(hLine)

	// 自定义X轴刻度显示为时间格式
	p.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02"}

	// 保存图表到文件
	if err := p.Save(10*vg.Inch, 6*vg.Inch, "quadrant_line_chart.png"); err != nil {
		log.Fatal(err)
	}

	log.Println("带连线的象限图已保存为 quadrant_line_chart.png")
}
