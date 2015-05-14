package Game

type Rect struct {
    X, Y, W, H float64
}

type Paddle struct {
    Rect
}

type Velocity struct {
    X, Y float64
}

type Ball struct {
    Rect
    Velocity
}

type World struct {
    W, H float64
    Paddle1 Paddle
    Paddle2 Paddle
    Ball
}

////////////////////////////////////////////////////////////////////////////////

func (p *Paddle) Move(speed float64, rect Rect) {

}

func (b *Ball) Move(rect Rect) {

}

func (b *Ball) BounceEdge(width int64) {
    
}

func (b *Ball) BouncePaddle(p Paddle) {

}

func CollisionRR(r1, r2 Rect) bool {
    return false
}

// func (p *Paddle) Move(speed float64, rect Rect) {
//     p.Rect.X = p.Rect.X + speed
//     if p.Rect.X < 0 {
//         p.Rect.X = 0
//     } else if p.Rect.X + p.Rect.W > rect.W {
//         p.Rect.X = rect.W - p.Rect.W
//     }
// }

// ////////////////////////////////////////////////////////////////////////////////