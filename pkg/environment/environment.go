package environment

type Environment struct {
	x_size         int
	y_size         int
	x_charging_pos int
	y_charging_pos int
	grid           [][]int
}

func NewEnvironment(x_size int, y_size int) Environment {
	grid := make([][]int, x_size)
	for i := range grid {
		grid[i] = make([]int, y_size)
	}
	return Environment{x_size: x_size, y_size: y_size, grid: grid, x_charging_pos: x_size - 1, y_charging_pos: y_size - 1}
}

func (e *Environment) GetGrid() [][]int {
	return e.grid
}

func (e *Environment) SetGrid(grid [][]int) {
	e.grid = grid
}

func (e *Environment) GetXSize() int {
	return e.x_size
}

func (e *Environment) GetYSize() int {
	return e.y_size
}

func (e *Environment) GetChargingPosition() (int, int) {
	return e.x_charging_pos, e.y_charging_pos
}
