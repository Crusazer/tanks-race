package collision

import (
	"math"
	m "github.com/Crusazer/tanks-race/pkg/math"
	"github.com/Crusazer/tanks-race/internal/physics/shapes"
)

type SATResult struct {
    Overlap    float64
    Normal     m.Vector2
    Contact1   m.Vector2
    Contact2   m.Vector2
    Edge1      m.Vector2
    Edge2      m.Vector2
}

func CheckCollision(rect1, rect2 *shapes.Rectangle) *SATResult {
    // Получаем оси для проверки
    axes1 := rect1.GetAxes()
    axes2 := rect2.GetAxes()
    
    // Объединяем все оси
    axes := append(axes1, axes2...)
    
    minOverlap := math.MaxFloat64
    smallestAxis := m.Vector2{}
    
    // Проверяем каждую ось
    for _, axis := range axes {
        // Проектируем оба прямоугольника на ось
        min1, max1 := rect1.Project(axis)
        min2, max2 := rect2.Project(axis)
        
        // Проверяем пересечение
        if max1 < min2 || max2 < min1 {
            // Нет пересечения - объект разделен
            return nil
        }
        
        // Вычисляем перекрытие
        overlap := math.Min(max1, max2) - math.Max(min1, min2)
        
        // Находим минимальное перекрытие
        if overlap < minOverlap {
            minOverlap = overlap
            smallestAxis = axis
        }
    }
    
    // Определяем правильное направление нормали
    centerDir := rect2.Center.Sub(rect1.Center)
    if smallestAxis.Dot(centerDir) < 0 {
        smallestAxis = smallestAxis.Scale(-1)
    }
    
    // Найдем точки контакта (simplified - берем центр пересечения)
    contact := findContactPoint(rect1, rect2)
    
    return &SATResult{
        Overlap:  minOverlap,
        Normal:   smallestAxis,
        Contact1: contact,
    }
}

// findContactPoint находит приближенную точку контакта
func findContactPoint(rect1, rect2 *shapes.Rectangle) m.Vector2 {
    // Берем среднюю точку между центрами
    mid := m.Vector2{
        X: (rect1.Center.X + rect2.Center.X) / 2,
        Y: (rect1.Center.Y + rect2.Center.Y) / 2,
    }
    
    return mid
}