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

func (v *Vector) HandleBounceOnBorder() {
    v.Direction.angle = (-1)*v.Direction.angle
}

func (v *Vector) HandleBounceOnRectangle(rect Geometry.Rect) {
    // The angle is chose in function of where the paddle is touched.
    // More the ball is near of the borders, more the angle will be low.
    // and the speed will increase.
    // So for this purpose I need to know on each % of the paddle the ball
    // will collide.

    //we now there is a collision between the ball and the paddle.
    //the coordinates of the vector is the coordinates of the ball's middle.
    //So we just have to check the y coordinates.

    y := v.Point.Y
    percent := y * 100 / rect.width
    
}

////////////////////////////////////////////////////////////////////////////////