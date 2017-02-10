package testfunc

import "math"

type SearchDomain struct {
	UpperBound []float64
	LowerBound []float64
}

type GlobalMinimum struct {
	X     []float64
	Value float64
}

func Rastrigin(x []float64) (y float64) {
	a := 10
	y = a * len(x)
	for _, xi := range x {
		y = y + (xi*xi - a*math.Cos(2*math.Pi))
	}
	return
}

func Ackley(x, y, float64) (z float64) {
	z = -20 * math.Exp(-0.2*math.Sqrt(0.5*(x*x+y*y)))
	z = z - math.Exp(0.5*(math.Cos(2*math.Pi*x)+math.Cos(2*math.Pi*y))) + math.E + 20
	return
}

func Sphere(x []float64) (y float64) {
	for _, xi := range x {
		y = y + (xi * xi)
	}
	return
}

func Rosenbrock(x []float64) (z float64) {
	for i := 0; i < len(x)-1; i++ {
		z = z + (100*math.Pow(x[i+1]-x[i]*x[i], 2) + math.Pow(x[i]-1, 2))
	}
	return
}

func Beale(x, y, float64) (z float64) {
	z = math.Pow(1.5-x+x*y, 2) + math.Pow(2.25-x+x*y*y, 2) + math.Pow(2.625-x+x*y*y*y, 2)
	return
}

func GoldsteinPrice(x, y, float64) (z float64) {
	z = (1 + math.Pow(x+y+1, 2)*(19-14*x+3*x*x-14*y+6*x*y+3*y*y))
	z = z * (30 + math.Pow(2*x-3*y, 2)*(18-32*x+12*x*x+48*y-36*x*y+27*y*y))
	return
}

func Booth(x, y, float64) (z float64) {
	z = math.Pow(x+2*y-7, 2) + math.Pow(2*x+y-5, 2)
	return
}

func BukinN6(x, y, float64) (z float64) {
	z = 100*math.Sqrt(math.Abs(y-0.01*x*x)) + 0.01*math.Abs(x+10)
	return z
}

func Matyas(x, y, float64) (z float64) {
	z = 0.26*(x*x+y*y) - 0.48*x*y
	return
}

func LeviN13(x, y, float64) (z float64) {
	z = math.Pow(math.Sin(3*math.Pi*x), 2)
	z = z + math.Pow(x-1, 2)*(1+math.Pow(math.Sin(3*math.Pi*y), 2))
	z = z + math.Pow(y-1, 2)*(1+math.Pow(math.Sin(2*math.Pi*y), 2))
	return
}

func ThreeHumpCamel(x, y, float64) (z float64) {
	z = 2*x*x - 1.05*math.Pow(x, 4) + math.Pow(x, 6)/6 + x*y + y*y
	return
}

func Easom(x, y, float64) (z float64) {
	z = -math.Cos(x) * math.Cos(y) * math.Exp(-(math.Pow(x-math.Pi, 2) + math.Pow(y-math.Pi, 2)))
	return
}

func CrossInTray(x, y, float64) (z float64) {
	z = math.Abs(100 - math.Sqrt(x*x+y*y)/math.Pi)
	z = math.Abs(math.Sin(x)*math.Sin(y)*math.Exp(z)) + 1
	z = -0.0001 * math.Pow(z, 0.1)
	return
}

func Eggholder(x, y, float64) (z float64) {
	z = -(y + 47) * math.Sin(math.Sqrt(math.Abs(x/2+(y+47))))
	z = z - x*math.Sin(math.Sqrt(math.Abs(x-(y+47))))
}

func HolderTable(x, y, float64) (z float64) {
	z = math.Abs(1 - math.Sqrt(x*x+y*y)/math.Pi)
	z = math.Sin(x) * math.Cos(y) * math.Exp(z)
	if z > 0 {
		z = -z
	}
	return
}

func McCormick(x, y, float64) (z float64) {
	z = math.Sin(x+y) + math.Pow(x-y, 2) - 1.5*x + 2.5*y + 1
}

func SchafferN2(x, y, float64) (z float64) {
	z = math.Pow(math.Sin(x*x-y*y), 2) - 0.5
	z = z / math.Pow(1+0.001*(x*x+y*y), 2)
	z = z + 0.5
	return
}

func SchafferN4(x, y, float64) (z float64) {
	z = math.Pow(math.Cos(math.Sin(math.Abs(x*x-y*y), 2))) - 0.5
	z = z / math.Pow(1+0.001*(x*x+y*y), 2)
	z = z + 0.5
	return
}

func StyblinskiTang(x []float64) (y float64) {
	for _, xi := range x {
		y = y + (math.Pow(xi, 4) - 16*xi*xi + 5*xi)
	}
	y = y / 2
	return
}
