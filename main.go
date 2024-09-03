package main

import (
	"ddb/storage/schemas"
	"fmt"
)

func main() {
	fmt.Println(
		"    ,---,        ,---,        ,---,. \n" +
			"  .'  .' `\\    .'  .' `\\    ,'  .'  \\\n" +
			",---.'     \\ ,---.'     \\ ,---.' .' | \n" +
			"|   |  .`\\  ||   |  .`\\  ||   |  |: | \n" +
			":   : |  '  |:   : |  '  |:   :  :  / \n" +
			"|   ' '  ;  :|   ' '  ;  ::   |    ;  \n" +
			"'   | ;  .  |'   | ;  .  ||   :     \\ \n" +
			"|   | :  |  '|   | :  |  '|   |   . | \n" +
			"'   : | /  ; '   : | /  ; '   :  '; | \n" +
			"|   | '` ,/  |   | '` ,/  |   |  | ;  \n" +
			";   :  .'    ;   :  .'    |   :   /   \n" +
			"|   ,.'      |   ,.'      |   | ,'    \n" +
			"'---'        '---'        `----'      \n")

	scs, err := schemas.New("./storage/schemas.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(scs.All())
}
