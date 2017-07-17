package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type terraform struct {
	Modules []struct {
		Resources map[string]struct {
			Primary struct {
				Name       string
				Attributes map[string]interface{}
			}
		}
	}
}

func main() {

	// Import the parameters from flags
	tag := flag.String("t", "", "Tag to group ansible inventory, defaults to grouping by machine name")
	file := flag.String("f", "terraform.tfstate", "Speficy path to terraform state file, defaults to terraform.tfstate")
	flag.Parse()

	// Read in the terraform state file
	f, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Printf("Error: %v]n", err)
		os.Exit(1)
	}

	var jsonData terraform

	// Unmarshall the terraform state
	if err := json.Unmarshal(f, &jsonData); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	//
	if *tag != "" {
		jsonTag := "tags."
		jsonTag += *tag

		// Create a map of tags, if we're doing that
		jsonTagMap := make(map[string]bool)

		for _, value := range jsonData.Modules[0].Resources {
			if _, ok := value.Primary.Attributes[jsonTag]; !ok {
				fmt.Printf("Error: Tag name \"%s\" doesn't exist in the terraform statefile for %s\n", jsonTag, value.Primary.Attributes["name"])
				os.Exit(1)
			}
			if !jsonTagMap[value.Primary.Attributes[jsonTag].(string)] {
				jsonTagMap[value.Primary.Attributes[jsonTag].(string)] = true
			}
		}

		// Interate through the bool map of unique tags to group the inventory correctly
		for tag, _ := range jsonTagMap {
			i := 0
			for _, value := range jsonData.Modules[0].Resources {
				if value.Primary.Attributes[jsonTag] == tag {
					if i == 0 {
						fmt.Printf("[%v]\n%v\n", value.Primary.Attributes[jsonTag], value.Primary.Attributes["ips.0"])
					} else {
						fmt.Printf("%v\n", value.Primary.Attributes["ips.0"])
					}
					i++
				}
			}
		}
	} else {
		for _, value := range jsonData.Modules[0].Resources {
			fmt.Printf("[%v]\n%v\n", value.Primary.Attributes["name"], value.Primary.Attributes["ips.0"])
		}
	}
}
