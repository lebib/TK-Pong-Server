package main

// TODO: Implement the bounce thing and the physic will be done. Kind of.

import (
    "fmt"
    "math"
    "server/Geometry"
)

func main() {
    r1 := Geometry.Rect{Geometry.Point{X:10, Y:10}, 5, 5}
    r2 := Geometry.Rect{Geometry.Point{X:10, Y:10}, 5, 5}
    fmt.Println(r1.CollisionRR(r2))
    fmt.Println(r1.X, r1.Y)
    fmt.Println(Geometry.Round(10.1337, 0))
    fmt.Println(Geometry.GetPointFromAngle(math.Pi / 6, 3).X)
    p1 := Geometry.GetPointFromAngle(math.Pi/4, 3)
    fmt.Println(p1.X, p1.Y)
    fmt.Println(
        Geometry.LineCoefficient(
            Geometry.Point{10, 10},
            p1,
        ),
    )

    dir := Geometry.Direction{math.Pi / 4, 0.5}
    p2 := Geometry.ApplyTranslation(Geometry.Point{3, 3}, dir)
    fmt.Println(p2.X, p2.Y)
}

// func main() {
//     // p1 := Geometry.Point{X:0, Y:0}
//     // p2 := Geometry.Point{X:10, Y:10}
//     // r1 := Geometry.Rect{Coords:p1, W:5, H:5}
//     // r2 := Geometry.Rect{Coords:p2, W:5, H:5}
//     p3 := Geometry.Point{X:5, Y:5}
//     p4 := Geometry.GetPointFromAngle(-math.Pi / 2)
//     // p5 := Geometry.GetPointFromAngle(math.Pi)
//     // p6 := Geometry.GetPointFromAngle(-math.Pi / 2)
//     // p7 := Geometry.GetPointFromAngle(-math.Pi)

//     fmt.Println(p3.X, p3.Y)
//     fmt.Println(p4.X, p4.Y)
//     // fmt.Println(p5.X, p5.Y)
//     // fmt.Println(p6.X, p6.Y)
//     // fmt.Println(p7.X, p7.Y)
//     p8 := Geometry.Point{X:p3.X+p4.X, Y:p3.Y+p4.Y}
//     fmt.Println(p8.X, p8.Y)
//     fmt.Println(Geometry.LineCoefficient(p3, p8))
//     if math.IsInf(Geometry.LineCoefficient(p3, p8), 0) {
//         fmt.Println("Hello, World!")
//     }
// }