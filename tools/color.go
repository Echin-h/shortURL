package tools

import (
	"fmt"
	"github.com/gookit/color"
)

func ColorPrint() {
	str := color.Red.Sprintf("This is red")
	fmt.Println(str)

}
