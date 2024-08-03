package main

import(
	"fmt"
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"

	"github.com/go-echarts/go-echarts/v2/components"
)

var (
	itemCntPie = 4
	seasons    = []string{"Spring", "Summer", "Autumn ", "Winter"}
)

func translado(i, j int) (int , int) {
	return i + 1 , j + 1
}

// generate random data for bar chart
func generateBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}
	return items
}

func generatePieItems() []opts.PieData {
	items := make([]opts.PieData, 0)
	for i := 0; i < itemCntPie; i++ {
		items = append(items, opts.PieData{Name: seasons[i], Value: rand.Intn(100)})
	}
	return items
}

func pieBase() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic pie example"}),
	)

	pie.AddSeries("pie", generatePieItems())
	return pie
}

func main() {
  x , y := translado( 1 , 2 )
	s := fmt.Sprintf("translado is %v , %v", x , y )
	fmt.Println(s)

	page := components.NewPage()
	page.AddCharts(
		pieBase(),
	)
	f, err := os.Create("pie.html")
	if err != nil {
		panic(err)
	}
	//page.Render(io.MultiWriter(f))
	page.Render(f)
	// create a new bar instance
		bar := charts.NewBar()
		// set some global options like Title/Legend/ToolTip or anything else
		bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
			Title:    "My first bar chart generated by go-echarts",
			Subtitle: "It's extremely easy to use, right?",
		}))

		// Put data into instance
		bar.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
			AddSeries("Category A", generateBarItems()).
			AddSeries("Category B", generateBarItems())
		// Where the magic happens
		//f, _ := os.Create("bar.html")
		//bar.Render(f)
}
