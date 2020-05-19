package numIslands

import "testing"

func TestNumsIsLands(t *testing.T) {
	grid := make([][]byte, 4)
	grid[0] = []byte{'1','1','0','1','0'}
	grid[1] = []byte{'1','1','0','0','1'}
	grid[2] = []byte{'1','1','0','0','0'}
	grid[3] = []byte{'0','0','1','1','0'}
	
	res := numIslands(grid)

	ans := 4

	if res != ans {
		t.Errorf("Failed, the res shoulbe %d, but it is %d.", ans, res)
	} else {
		t.Logf("Success, the res is %d", res)
	}
}