package connectRFC
/*
#include <stdio.h>
#include <stdlib.h>

void pathgetter() {
	system("echo %CLASSPATH% > classpathfile.txt");
	system("echo %GOPATH% > gopathfile.txt");
}

void executeSystem(char *s1,char *s2) {
    system(s1);
    system(s2);
}
*/
import "C"
import (
		"io/ioutil"
		"os"
		"path/filepath"
		"strings"
		"strconv"
		"bufio"
		"fmt"
	)

var functionName string
var logonLang string
var Client string
var instanceNo string
var abapServer string
var User string
var Passwd string
var varInputParams map[string]string
var varInputParamsType map[string]string
var varExportParams map[string]string
var varExportParamsType map[string]string
var structExportParams map[string]string
var tableParams map[string][]string

func Createfiles(lang string, client string, sysnr string, ashost string, user string, passwd string, functionname string) bool {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
    	panic(err)
    	return false
    }
	finalFormedString := "jco.client.lang=" + lang + "\njco.client.client=" + client + "\njco.client.sysnr=" + sysnr + "\njco.client.ashost=" + ashost + "\njco.client.user=" + user + "\njco.client.passwd=" + passwd;
	byteArray := []byte(finalFormedString)
	err = ioutil.WriteFile("ABAP_AS.jcoDestination",byteArray,0644)
	if err!=nil {
		panic(err)
		return false
	}
	C.pathgetter()
	bclasspath, err := ioutil.ReadFile("classpathfile.txt")
    if err != nil { 
    	panic(err) 
    	return false
    }
    bgopath, err := ioutil.ReadFile("gopathfile.txt")
    if err != nil { 
    	panic(err) 
    	return false
    }
    byteArray = []byte(functionname)
    err = ioutil.WriteFile("functionName.txt",byteArray,0644)
	if err!=nil {
		panic(err)
		return false
	}
//	now only variables
	finalFormedString = ""
	for key1 := range varInputParams {
		switch {
			case finalFormedString=="" : 
				finalFormedString = finalFormedString + key1 + "\n" + varInputParams[key1] + "\n"
			case finalFormedString!="" : 
				finalFormedString = finalFormedString + "\n" + key1 + "\n" + varInputParams[key1] + "\n"	
		}
	}	
	for key2 := range varInputParamsType {
		finalFormedString = finalFormedString + varInputParamsType[key2]
	}
	byteArray = []byte(finalFormedString)
	err = ioutil.WriteFile("inputParams.txt",byteArray,0644)
	if err!=nil {
		panic(err)
		return false
	}
	finalFormedString1 := "javac " + strings.Replace(string(bgopath)," \r\n","",-1) + "\\src\\com\\sapconnector\\deysub\\connectRFC\\connectToAbap.java -d " + dir
	finalFormedString2 := "java -cp \".;" + strings.Replace(string(bclasspath)," \r\n","",-1) + "\" connectToAbap"
/*	byteArray = []byte(finalFormedString)  
	err = ioutil.WriteFile("executeInWin.bat",byteArray,0644)
	if err!=nil {
		panic(err)
		return false
	}	*/
	exstring1 := C.CString(finalFormedString1)
	exstring2 := C.CString(finalFormedString2)
	C.executeSystem(exstring1,exstring2)
	b, err := ioutil.ReadFile("tempfile.txt")
    if err != nil { 
    	panic(err) 
    	return false
    }
    if string(b)!="success"{
    	return false
    }
	return true		 
}

func Connect(logon_language string, client string, instance_no string, abap_server_loc string, user string, password string) bool {
	logonLang = logon_language
	Client = client
	instanceNo = instance_no
	abapServer = abap_server_loc
	User = user
	Passwd = password
	return true
}

func GetFunction(function_name string) bool {
	functionName = function_name
	varInputParams = make(map[string]string)
	varInputParamsType = make(map[string]string)
	return true
}

func SetVariableImportParameterInt32(parameter_name string,parameter_value int32) bool {
	varInputParams[parameter_name] = string(parameter_value)
	varInputParamsType[parameter_name] = "integer"
	return true
}

func SetVariableImportParameterInt16(parameter_name string,parameter_value int16) bool {
	varInputParams[parameter_name] = string(parameter_value)
	varInputParamsType[parameter_name] = "short"
	return true
}

func SetVariableImportParameterInt64(parameter_name string,parameter_value int64) bool {
	varInputParams[parameter_name] = string(parameter_value)
	varInputParamsType[parameter_name] = "long"
	return true
}

func SetVariableImportParameterString(parameter_name string,parameter_value string) bool {
	varInputParams[parameter_name] = string(parameter_value)
	varInputParamsType[parameter_name] = "string"
	return true
}

func SetVariableImportParameterFloat64(parameter_name string,parameter_value float64) bool {
	varInputParams[parameter_name] = strconv.FormatFloat(parameter_value, 'f', 6, 64)
	varInputParamsType[parameter_name] = "double"
	return true
}

func GetVariableExportParameterInt32(parameter_name string) int {
	returnval, e := strconv.Atoi(varExportParams[parameter_name])
	if e != nil {
        fmt.Println(e)
    }
	return returnval
}

func GetVariableExportParameterInt16(parameter_name string) int {
	returnval, e := strconv.Atoi(varExportParams[parameter_name])
	if e != nil {
        fmt.Println(e)
    }
	return returnval
}

func GetVariableExportParameterInt64(parameter_name string) int {
	returnval, e := strconv.Atoi(varExportParams[parameter_name])
	if e != nil {
        fmt.Println(e)
    }
	return returnval
}

func GetVariableExportParameterString(parameter_name string) string {
	return varExportParams[parameter_name]
}

func GetVariableExportParameterFloat64(parameter_name string) float64 {
	returnval, e := strconv.ParseFloat(varExportParams[parameter_name],64)
	if e != nil {
        fmt.Println(e)
    }
	return returnval
}

func GetStructure(parameter_name string) map[string]string {
	file, err := os.Open("exportStructParameters.txt")
  	if err != nil {
  	}
  	defer file.Close()
	scanner := bufio.NewScanner(file)
	structstartcount := 0
	structname := ""
	fieldname := ""
	fieldval := ""
	thisstruct := 0
	structExportParams = make(map[string]string)
	for scanner.Scan() {
    	if scanner.Text()!="END;" && structstartcount!=0{
    		fieldname = scanner.Text()
    		scanner.Scan()
    		fieldval = scanner.Text()
    		if thisstruct==1{
    			structExportParams[fieldname] = fieldval
    		}
    		structstartcount = structstartcount + 1
    	}
    	if scanner.Text()=="END;" && structstartcount!=0{
    		structstartcount = 0
    		thisstruct = 0
    	}
    	if structstartcount==0{
    		structname = scanner.Text()
    		if structname==parameter_name{
    			thisstruct = 1
    		}
    		structstartcount = structstartcount + 1
    	}
	}
/*	for key, value := range structExportParams {
    	fmt.Println("Key:", key, "Value:", value)
	}*/
	return structExportParams	
}

func GetTable(parameter_name string) map[string][]string {
	tableParams = make(map[string][]string)
	file, err := os.Open("exportTableParams.txt")
  	if err != nil {
  	}
  	defer file.Close()
	scanner := bufio.NewScanner(file)
	tablereadcounter := 0
	name := ""
	key := ""
	count := 0
	tempmap := make(map[int]string)
	for scanner.Scan(){
		if tablereadcounter==0{
			name = scanner.Text()
			tablereadcounter = tablereadcounter + 1
			scanner.Scan()
		}
		if tablereadcounter==1{
			key = scanner.Text()
			if key!="ENDMETADATA;" && name==parameter_name{	
				tempmap[count] = key
				count = count + 1
			}
			if key=="ENDMETADATA;"{
				break
			}
		}
	}
	samefile, error := os.Open("exportTableParams.txt")
  	if error != nil {
  	}
  	defer samefile.Close()
	scannerAgain := bufio.NewScanner(samefile)
	flagstart := 0
	count = 0
	for scannerAgain.Scan(){
		if scannerAgain.Text()=="ENDMETADATA;"{
			flagstart = 1
			scannerAgain.Scan()
		}
		if flagstart==1 && scannerAgain.Text()!="END;"{
			tableParams[tempmap[count]] = append(tableParams[tempmap[count]], scannerAgain.Text())
			count = count + 1
			if count==len(tempmap){
				count = 0
			}
		}
	}
	return tableParams
}

func ChangeTableParameter(parameter_name string) bool {
	return true
}

func Execute(function_name string) bool {
	returnbool := Createfiles(logonLang,Client,instanceNo,abapServer,User,Passwd,function_name)
	//logic only for variables & structures
	varExportParams = make(map[string]string)
	varExportParamsType = make(map[string]string)
	file, err := os.Open("exportParameters.txt")
  	if err != nil {
    	return false
  	}
  	defer file.Close()
  	scanner := bufio.NewScanner(file)
  	for scanner.Scan() {
  		varKey := scanner.Text()
  		scanner.Scan()
  		varType := scanner.Text()
  		scanner.Scan()
  		varValue := scanner.Text()
  		varExportParams[varKey] = varValue
  		varExportParamsType[varKey] = varType
  	}
  	
    return returnbool
}

func CloseConnection(){
	os.Remove("inputParams.txt")
  	os.Remove("exportParameters.txt")
  	os.Remove("exportStructParameters.txt")
  	os.Remove("functionName.txt")
	os.Remove("tempfile.txt")
    os.Remove("classpathfile.txt")
    os.Remove("gopathfile.txt")
    os.Remove("exportTableParams.txt")
    os.Remove("ABAP_AS.jcoDestination")
    os.Remove("connectToAbap.class")
}