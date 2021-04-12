package ascii

import (
	"fmt"
	"testing"
)

func TestImageToGreyScales(t *testing.T) {
	fmt.Print(ImageToGreyScales("assets/sample-image.jpg"))
}
