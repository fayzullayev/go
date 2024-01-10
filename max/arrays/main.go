package main

import "fmt"

func main() {

	//websites := make(map[string]string)
	var websites map[string]string = map[string]string{
		"google": "https://wwww.google2.com",
	}

	fmt.Println(websites)

	websites["google"] = "https://wwww.google.com"
	websites["amazon"] = "https://wwww.aws.com"
	websites["facebook"] = "https://wwww.facebook.com"

	fmt.Printf("%+v\n", websites)

	value, ok := websites["amazon2"]
	fmt.Println(value, ok)

	delete(websites, "facebook")

	for k, v := range websites {
		fmt.Println(k, v)
	}

}
