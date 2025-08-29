package math

type Matrix2 [2][2]float64

func (m Matrix2) MulVector(v Vector2) Vector2{
	return Vector2{
		X: m[0][0]*v.X + m[0][1]*v.Y,
		Y: m[1][0]*v.X + m[1][1]*v.Y,
	}
}

// Identity возвращает единичную матрицу
func Identity() Matrix2 {
	return Matrix2{{1, 0}, {0, 1}}
}

func Scale(s float64) Matrix2{
		return Matrix2{{s, 0}, {0, s}}
}