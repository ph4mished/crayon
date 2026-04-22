//to test if true color fallback works
package main
import(
	"fmt"
	//"crayon"
	"github.com/ph4mished/crayon"
)

func main(){
//	col := crayon.Parse("[fg=#ff5fd7]HELLO [fg=#d7875f]WORLD[reset]")
	 rgb := crayon.Parse("[fg=rgb(230,240,abc)]HELLO[reset]")
	 c_256 := crayon.Parse("[fg=rgb(255,255,255)]HELLO[fg=231]WORLD[reset]")
	 hex := crayon.Parse("[fg=#875f00]HELLO[reset]")
	 try_it := crayon.Parse("Start [fg=cyan]Middle[reset] End")//"Start [fg=cyan Middle [reset] End")
	 
	fmt.Println("RGB: ", rgb.Sprint())
	fmt.Println("256: ", c_256.Sprint())
	fmt.Println("HEX: ", hex.Sprint())
	fmt.Println(try_it)
}
