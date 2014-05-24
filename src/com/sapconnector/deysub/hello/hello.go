package main

import (
		"fmt"

		"com/sapconnector/deysub/connectRFC"
	)

func main() {
	connectRFC.Connect("EN","001","15","uxciebj.wdf.sap.corp","DEYSUB","algo..addict965431")
	connectRFC.GetFunction("STFC_CONNECTION")
	connectRFC.SetVariableImportParameterString("REQUTEXT","I am Rohit")
	returned3 := connectRFC.Execute("STFC_CONNECTION")
	if returned3==true{
		fmt.Printf("Boolean : %b\n",returned3)
	}
	fmt.Println(connectRFC.GetVariableExportParameterString("ECHOTEXT"))
	fmt.Println(connectRFC.GetVariableExportParameterString("RESPTEXT"))
	connectRFC.GetFunction("RFC_SYSTEM_INFO")
	returned3 = connectRFC.Execute("RFC_SYSTEM_INFO")
	if returned3==true{
		fmt.Printf("Boolean : %b\n",returned3)
	}
	test := connectRFC.GetStructure("RFCSI_EXPORT")
	fmt.Println(test["RFCPROTO"])
	connectRFC.GetFunction("BAPI_COMPANYCODE_GETLIST")
	returned4 := connectRFC.Execute("BAPI_COMPANYCODE_GETLIST")
	if returned4==true{
		fmt.Printf("Boolean : %b\n",returned4)
	}
	test1 := connectRFC.GetStructure("RETURN")
	fmt.Println(test1["LOG_MSG_NO"])
	returnMap := connectRFC.GetTable("COMPANYCODE_LIST")
	fmt.Println(returnMap["COMP_CODE"][0])
	connectRFC.CloseConnection()
}