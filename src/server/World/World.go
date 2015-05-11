package World
import "server/Geometry"

////////////////////////////////////////////////////////////////////////////////

type Paddle struct {
    Geometry.Rect
}

type Ball struct {
    Geometry.Rect
    Geometry.Vector
}

type World struct {
    Geometry.Rect
    Ball
    Paddle1 Paddle
    Paddle2 Paddle
}

////////////////////////////////////////////////////////////////////////////////

func (p *Paddle) MovePaddle(dir int, celerity float64, width int) {
    if dir == 1 {
        Paddle.Rect.X = Paddle.Rect.X + celerity
    } else if dir == 2 {
        Paddle.Rect.X = Paddle.Rect.X - celerity
    }
    if Paddle.Rect.X < 0 {
        Paddle.Rect.X = 0
    } else if Paddle.Rect.X + Paddle.Rect.Width > width {
        Paddle.Rect.X = width - Paddle.Rect.Width
    }
}

////////////////////////////////////////////////////////////////////////////////