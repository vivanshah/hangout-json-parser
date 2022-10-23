package hangouts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadHangoutsFile(input *string) (Hangouts, error) {

	jsonFile, err := os.Open(*input)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	var hangouts Hangouts
	fmt.Println("Parsing hangouts file")
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(byteValue, &hangouts)
	if err != nil {
		fmt.Println("Error parsing input file: ", err.Error())
		return Hangouts{}, err
	}

	return hangouts, nil

}
