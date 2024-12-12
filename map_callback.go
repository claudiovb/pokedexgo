package main

import (
	"fmt"
)

func commandMap(config *Config, _ ...string) error {

	locationArea, err := config.pokeapiClient.GetLocationsArea(config.nextApiUrl)
	if err != nil {
		return err
	}
	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}
	if locationArea.Next != nil {
		config.nextApiUrl = *locationArea.Next
	}
	if locationArea.Previous != nil {
		config.previousApiUrl = *locationArea.Previous
	}
	//Try hard mode:
	// results, ok := res["results"].([]interface{})
	// if !ok {
	// 	return fmt.Errorf("Error: 'results' field is not a slice of maps")
	// }

	// for _, result := range results {
	// 	resultMap, ok := result.(map[string]interface{})
	// 	if !ok {
	// 		continue
	// 	}
	// 	name, ok := resultMap["name"].(string)
	// 	if !ok {
	// 		continue
	// 	}
	// 	fmt.Println(name)
	// }

	//TODO parse the json output
	return nil
}

func commandMapBack(config *Config, _ ...string) error {
	locationArea, err := config.pokeapiClient.GetLocationsArea(config.previousApiUrl)
	if err != nil {
		return err
	}
	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}
	if locationArea.Next != nil {
		config.nextApiUrl = *locationArea.Next
	}
	if locationArea.Previous != nil {
		config.previousApiUrl = *locationArea.Previous
	}
	return nil
}
