package imagemasker

import "testing"

func TestFindAbsoluteMatch(t *testing.T) {
	// Test that a mask finds the correct position in the image
	findAbsoluteMatch(`./test_images/peppers.png`, `./test_images/peppers_crop.png`)
}
