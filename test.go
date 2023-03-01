package main
import "fmt"

type void struct{}
var member void

func has(set map[string]void, entry string) bool {
	for key, _ := range set {
		if (key == entry) {
		fmt.Println("True")
			return true
		}
	}

	fmt.Println("False")
	return false
}
func main() {
    set := make(map[string]void)
    set["apple"] = member
    set["orange"] = member
    set["mango"] = member
    set["paw-paw"] = member
    // fmt.Println(set)
	// set["mango"]
    fmt.Println("The range is:")
	for k := range set {
		fmt.Println(k)
	}

	has(set, "paw-paw")

}
