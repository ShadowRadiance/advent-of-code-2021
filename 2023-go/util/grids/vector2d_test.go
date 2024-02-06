package grids

import "testing"

func TestPosition_Mod_LargePositive(t *testing.T) {
	initial := Vector2D[int]{X: 200, Y: 400}
	expected := Vector2D[int]{X: 200 % 49, Y: 400 % 49}
	actual := initial.ModWrap(49, 49)
	if expected != actual {
		t.Errorf("Expected: %+v, Got %+v\n", expected, actual)
	}
}
func TestPosition_Mod_SmallNegative(t *testing.T) {

	initial := Vector2D[int]{X: -1, Y: -1}
	expected := Vector2D[int]{X: 48, Y: 48}
	actual := initial.ModWrap(49, 49)
	if expected != actual {
		t.Errorf("Expected: %+v, Got %+v\n", expected, actual)
	}
}
func TestPosition_Mod_NoChange(t *testing.T) {
	initial := Vector2D[int]{X: 1, Y: 1}
	expected := Vector2D[int]{X: 1, Y: 1}
	actual := initial.ModWrap(49, 49)
	if expected != actual {
		t.Errorf("Expected: %+v, Got %+v\n", expected, actual)
	}
}
func TestPosition_Mod_LargeNegatives(t *testing.T) {
	initial := Vector2D[int]{X: -100, Y: -200}
	expected := Vector2D[int]{X: 47, Y: 45}
	actual := initial.ModWrap(49, 49)
	if expected != actual {
		t.Errorf("Expected: %+v, Got %+v\n", expected, actual)
	}

}
