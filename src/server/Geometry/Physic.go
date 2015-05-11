package Geometry
import "fmt"

////////////////////////////////////////////////////////////////////////////////

type Direction struct {
    A, Speed float64
}

type Vector struct {
    Point
    Direction
}

////////////////////////////////////////////////////////////////////////////////

func ApplyTranslation(point Point, dir Direction) Point {
    dirPoint := GetPointFromAngle(dir.A, 3)
    fmt.Println(dirPoint.X * dir.Speed)
    fmt.Println(dirPoint.Y * dir.Speed)
    point.X = point.X + dirPoint.X * dir.Speed
    point.Y = point.Y + dirPoint.Y * dir.Speed
    return point
}

func ApplyBounce(vector Vector)

////////////////////////////////////////////////////////////////////////////////