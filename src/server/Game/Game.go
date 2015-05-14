package Game

import "strconv"

type Player struct {
    Name string
    Score int64
}

type Game struct {
    P1 Player
    P2 Player
    World
}

////////////////////////////////////////////////////////////////////////////////

func (g *Game) Init() {
    g.P1 = Player{"Test", 1}
    g.P2 = Player{"John", 1}
    g.World = World{
        800, 600,
        Paddle{Rect{20, 0, 150, 15}},
        Paddle{Rect{20, 500, 150, 15}},
        Ball{Rect{50, 50, 15, 15}, Velocity{1, 1}},
    }
}

func (g *Game) MainLoop() {
    // 1- Get inputs
    // 2- Handle inputs
    // 3- Send outputs
}

func (g *Game) Init2Json() string {
    json := ""
    return json
}

func (g *Game) GameState2Json() string {
    p1_s := strconv.FormatInt(g.P1.Score, 10)
    p2_s := strconv.FormatInt(g.P2.Score, 10)
    p1_x := strconv.FormatFloat(g.World.Paddle1.X, 'f', 4, 64)
    p1_y := strconv.FormatFloat(g.World.Paddle1.Y, 'f', 4, 64)
    p2_x := strconv.FormatFloat(g.World.Paddle2.X, 'f', 4, 64)
    p2_y := strconv.FormatFloat(g.World.Paddle2.Y, 'f', 4, 64)
    json := ""
    json += "{"
    json += "\"Players\": [["
    json += p1_s + "," + p1_x + "," + p1_y + "],["
    json += p2_s + "," + p2_x + "," + p2_y + "]]}"
    return json
}

////////////////////////////////////////////////////////////////////////////////