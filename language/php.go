package language

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"sync"
)

type PHPClassGenerator struct {
	Namespace  string
	SaveFolder string
}

func NewPHPClassGenerator(namespace, saveFolder string) *PHPClassGenerator {
	if len(namespace) == 0 {
		namespace = "Example"
	}

	if len(saveFolder) == 0 {
		saveFolder = "./"
	}

	return &PHPClassGenerator{
		Namespace:  namespace,
		SaveFolder: saveFolder,
	}
}

func (php *PHPClassGenerator) Generate(jsonMap map[string]interface{}, className string) {
	wt := sync.WaitGroup{}
	wt.Add(1)
	go php.classGenerator(jsonMap, className, &wt)
	wt.Wait()
	fmt.Println("PHP Class Generate Success")
}

func (php *PHPClassGenerator) classGenerator(jsonMap map[string]interface{}, className string, wt *sync.WaitGroup) {
	defer wt.Done()

	err := createFloder(php.SaveFolder)
	file, err := openFile(php.SaveFolder + "/" + className + ".php")
	if err != nil {
		log.Fatalf("open file failed. reason: %s\n", err)
	}

	defer file.Close()

	propertyBuffer := &bytes.Buffer{}
	funcsBuffer := &bytes.Buffer{}
	classBuffer := &bytes.Buffer{}

	classBuffer.WriteString(fmt.Sprintf("<?php\n\n namespace %s;\n\n class %s  {\n", php.Namespace, className))
	for k, v := range jsonMap {
		propertyKey := toCamelCase(k, false)
		funcNameKey := toCamelCase(k, true)
		switch v.(type) {
		case map[string]interface{}:
			php.writeBuffer(propertyBuffer, funcsBuffer, "\\"+php.Namespace+"\\"+toCamelCase(k, true), propertyKey, funcNameKey)
			wt.Add(1)
			go php.classGenerator(v.(map[string]interface{}), toCamelCase(k, true), wt)
		case []interface{}:
			php.arrayPropertyGenerator(k, v, propertyBuffer, funcsBuffer, propertyKey, funcNameKey, wt)
		default:
			typeString := reflect.TypeOf(v).String()
			php.writeBuffer(propertyBuffer, funcsBuffer, typeString, propertyKey, funcNameKey)
		}

	}

	propertyBuffer.WriteTo(classBuffer)
	funcsBuffer.WriteTo(classBuffer)
	classBuffer.WriteString("}\n")
	classBuffer.WriteTo(file)
}

func (php *PHPClassGenerator) arrayPropertyGenerator(k string, v interface{}, propertyBuffer, funcsBuffer *bytes.Buffer, propertyKey, funcNameKey string, wt *sync.WaitGroup) {
	vv := v.([]interface{})
	vvFirst := vv[0]
	switch vvFirst.(type) {
	case map[string]interface{}:
		php.writeBuffer(propertyBuffer, funcsBuffer, "\\"+php.Namespace+"\\"+toCamelCase(k, true)+"[]", propertyKey, funcNameKey)
		wt.Add(1)
		go php.classGenerator(vvFirst.(map[string]interface{}), toCamelCase(k, true), wt)
	case []interface{}:
		php.writeBuffer(propertyBuffer, funcsBuffer, "array", propertyKey, funcNameKey)
	default:
		typeString := reflect.TypeOf(vvFirst).String()
		php.writeBuffer(propertyBuffer, funcsBuffer, typeString+"[]", propertyKey, funcNameKey)
	}
}

func (php *PHPClassGenerator) writeBuffer(propertyBuffer, funcsBuffer *bytes.Buffer, propertyType, propertyKey, funcNameKey string) {
	propertyBuffer.WriteString(fmt.Sprintf("\t/**\n\t* @var %s\n\t*/\n\tprivate $%s;\n\n", propertyType, propertyKey))
	funcsBuffer.WriteString(fmt.Sprintf("\t/**\n\t* @return %s\n\t*/\n\tpublic function get%s() {\n\t\treturn $this->%s;\n\t}\n\n", propertyType, funcNameKey, propertyKey))
}
