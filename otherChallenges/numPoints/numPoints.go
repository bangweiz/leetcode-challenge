package numPoints

import "math"

func numPoints(points [][]int, r int) int {
	for i := 0; i < len(points) - 1; i++ {
		for j :=0; j < len(points) - 1 - i; j++ {
			if points[j][0] > points[j + 1][0] {
				points[j], points[j + 1] = points[j + 1], points[j]
			}
		}
	}

	centers := make([][]float64, 0)
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			dist := cal(points[i], points[j])
			if i != j && dist <= r * r * 4 {
				c1, c2 := findCenter(points[i], points[j], dist, r)
				centers = append(centers, c1)
				centers = append(centers, c2)
			}
		}
	}

	max := 0
	for i := range centers {
		count := 0
		for j := range points {
			if calFloat(centers[i], points[j]) <= math.Pow(float64(r), 2) {
				count++
			}
		}
		if max < count {
			max = count
		}
	}
	if max == 0 {
		return 1
	}
	return max
}

func calFloat(p1 []float64, p2 []int) float64  {
	l := p1[0] - float64(p2[0])
	h := p1[1] - float64(p2[1])
	return l * l + h * h
}

func cal(p1, p2 []int) int {
	l := p1[0] - p2[0]
	h := p1[1] - p2[1]
	return l * l + h * h
}

func findCenter(p1, p2 []int, dist, r int) (center1, center2 []float64) {
	center1 = make([]float64, 2)
	center2 = make([]float64, 2)
	var theta1, theta2 float64
	theta1 = math.Atan(float64(p2[1] - p1[1]) / float64(p2[0] - p1[0]))
	theta2 = math.Acos((math.Sqrt(float64(dist)) / 2) / float64(r))
	center1[0] = float64(p1[0]) + float64(r) * math.Cos(theta1 + theta2)
	center1[1] = float64(p1[1]) + float64(r) * math.Sin(theta1 + theta2)
	center2[0] = float64(p1[0]) + float64(r) * math.Cos(theta1 - theta2)
	center2[1] = float64(p1[1]) + float64(r) * math.Sin(theta1 - theta2)
	return
}