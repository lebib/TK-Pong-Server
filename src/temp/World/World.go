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

func (p *Paddle) MovePaddle(celerity float64, width int) {
    Paddle.Rect.X = Paddle.Rect.X + celerity
    
    if Paddle.Rect.X < 0 {
        Paddle.Rect.X = 0
    } else if Paddle.Rect.X + Paddle.Rect.Width > width {
        Paddle.Rect.X = width - Paddle.Rect.Width
    }
}

////////////////////////////////////////////////////////////////////////////////