package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	a := emoji.Sprint(":world_map:")
	return "Hello " + a + "!"
}
