package main

import (
	"encoding/json"
	"flag"
	"log"
	"mztlive/classgenerator/language"
)

// const jsonString string = `{"recipient":{"type":"1","address_id":"7964","name":"铜锣湾渣fit人","country_code":"852","phone":"91234567","country_code_2":"852","phone_2":"","address":"香港, 香港銅鑼灣軒尼詩道铜锣湾广场一期12座1888室","coordinate":"114.1828103000,22.2802646000","service_type":"1"},"sender":{"type":"1","address_id":"7962","name":"大飞","country_code":"852","phone":"91234567","address":"香港, 香港旺角旺角花園街12座1888室","coordinate":"114.1708026000,22.3210526000","service_type":"1","method":""},"remark":{"overview":"","text":"asldfklajfljaklsfa","merchant":"adfafd"},"box_size_code":"S","box_quantity":"1","total_weight":"1","cod":{"service_fee":"10","cash":"10","total":"20.00"},"merchant_id":"","client_order_id":"","client_ref_id":"adlfjalsjflajlk","client_shipment_code":"","shipment_code":"","pickup":{"datetime":"2022-04-08 12:00-16:00"},"delivery":{"datetime":""}}`

// const jsonString string = `{"foo": [{"key": 1}, {"key": 2}],"two": {"three": 4, "four": 5}, "hello": [1,2,3], "world": [[1,2,3]]}`

var PhpClassNamespace string
var SaveFolder string
var jsonString string

func main() {
	flag.StringVar(&PhpClassNamespace, "namespace", "Example", "PHP Class Namespace")
	flag.StringVar(&SaveFolder, "f", "./", "Save Folder")
	flag.StringVar(&jsonString, "json", "", "json string")
	flag.Parse()

	if jsonString == "" {
		log.Fatalf("json string is empty.\n")
	}

	jsonMap := make(map[string]interface{})

	err := json.Unmarshal([]byte(jsonString), &jsonMap)
	if err != nil {
		log.Fatalf("json format failed. reason: %s\n", err)
	}

	classGenerator := language.NewPHPClassGenerator(PhpClassNamespace, SaveFolder)
	classGenerator.Generate(jsonMap, "Example")
}
