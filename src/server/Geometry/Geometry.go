package Geometry
import "math"

////////////////////////////////////////////////////////////////////////////////

type Point struct  {
    X, Y float64
}

type Rect struct {
    Point
    W, H float64
}

////////////////////////////////////////////////////////////////////////////////

func (r1 *Rect) CollisionRR(r2 Rect) bool {
    if r1.X + r1.W < r2.X || r1.X > r2.X + r2.W {
        if r1.Y + r1.H < r2.Y || r1.Y > r2.Y + r2.H {
            return false;
        }
    } 
    return true
}

func (r *Rect) CollisionPR(p Point) bool {
    if p.X < r.X || p.X > r.X + r.W {
        if p.Y < r.Y || p.Y > r.Y + r.H {
            return false;
        }
    }
    return true;
}

////////////////////////////////////////////////////////////////////////////////

func Round(val float64, precision int) float64 {
    buffer := math.Pow(10, float64(precision))
    return TT(val * buffer) / buffer
}

func TT(val float64) float64  {
    return math.Floor(val + 0.5)
}

func GetPointFromAngle(angle float64, precision int) Point {
    return Point{ Round(math.Cos(angle), precision),
        Round(math.Sin(angle), precision),
    }
}

func LineCoefficient(p1, p2 Point) float64 {
    return (p2.Y - p1.Y) / (p2.X - p1.X)
}

func DegreesToRadians(degrees float64) float64 {

}

////////////////////////////////////////////////////////////////////////////////