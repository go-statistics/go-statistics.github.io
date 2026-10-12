// Command figure draws the landing page's figure, static/img/figure.svg:
//
//	cd figure && go run . > ../static/img/figure.svg
//
// The page's terminal shows this file from func main on, read at build
// time, and the CI draws the figure again and fails if a byte differs.
package main

import (
	"fmt"
	"os"

	"github.com/go-ndarray/ndarray/random"
	"github.com/go-plot/plot"
	"github.com/go-statistics/stats"
)

func main() {
	// 10 000 gamma draws from seed 42: the
	// same draws SciPy makes from that seed.
	truth := stats.Gamma{A: 2.5, Loc: 0, Scale: 1.5}
	x, err := truth.Rvs(random.DefaultRNG(42), 10000)
	if err != nil {
		panic(err)
	}

	// The maximum-likelihood fit, location
	// held at 0, and a Kolmogorov–Smirnov
	// test of the draws against the fit.
	fixed := map[string]float64{"Loc": 0}
	fit, err := stats.FitGamma(x, &stats.FitOptions{Fixed: fixed})
	if err != nil {
		panic(err)
	}
	ks, err := stats.KSTest1Samp(x, fit.CDF, nil)
	if err != nil {
		panic(err)
	}

	// The densities on a grid.
	t := make([]float64, 400)
	pdf := make([]float64, len(t))
	fitted := make([]float64, len(t))
	for i := range t {
		t[i] = 16 * float64(i) / float64(len(t)-1)
		pdf[i], fitted[i] = truth.PDF(t[i]), fit.PDF(t[i])
	}

	fig, axs := plot.Subplots(1, 1, plot.FigSize(6.4, 4))
	ax := axs[0][0]
	ax.Hist(x, plot.Bins(60), plot.Density(),
		plot.Alpha(0.35), plot.Label("10 000 draws"))
	ax.Plot(t, pdf, plot.Label("Gamma(2.5, scale 1.5)"))
	ax.Plot(t, fitted, plot.LineStyle("--"),
		plot.Label(fmt.Sprintf("fit: a = %.3f, scale = %.3f",
			fit.A, fit.Scale)))
	ax.SetTitle(fmt.Sprintf(
		"Kolmogorov–Smirnov against the fit: D = %.4f, p = %.2f",
		ks.Statistic, ks.PValue))
	ax.SetXLim(0, 16)
	ax.SetXLabel(`$x$`)
	ax.Legend(plot.Loc("upper right"))
	if err := fig.SaveSVG(os.Stdout); err != nil {
		panic(err)
	}
}
