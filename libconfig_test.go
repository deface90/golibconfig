package libconfig

import (
	"testing"
)

func TestConfig_ReadString(t *testing.T) {
	str := "a=\"1\";"

	config := NewConfig()
	err := config.ReadString(str)
	if err != nil {
		t.Fatal("error reading string")
	}
}

func TestConfig_Lookup(t *testing.T) {
	str := "aint: 1;\n" +
		"astring: \"test\";\n" +
		"afloat: 1.0;\n" +
		"abool: true;\n"

	config := NewConfig()
	err := config.ReadString(str)
	if err != nil {
		t.Fatal("error reading string")
	}

	aint, err := config.LookupInt("aint")
	if err != nil {
		t.Fatal("error getting setting" + err.Error())
	}
	if aint != 1 {
		t.Fatal("wrong int value")
	}

	astr, err := config.LookupString("astring")
	if err != nil {
		t.Fatal("error getting setting" + err.Error())
	}
	if astr != "test" {
		t.Fatal("wrong string value")
	}

	afloat, err := config.LookupFloat("afloat")
	if err != nil {
		t.Fatal("error getting setting" + err.Error())
	}
	if afloat != 1.0 {
		t.Fatal("wrong float value")
	}

	abool, err := config.LookupBool("abool")
	if err != nil {
		t.Fatal("error getting setting" + err.Error())
	}
	if abool != true {
		t.Fatal("wrong bool value")
	}
}

func TestSetting_IsNil(t *testing.T) {
	confstr := "struct: { aint=1; astring=\"test\"; afloat=13.37; abool=true; }"
	config := NewConfig()
	err := config.ReadString(confstr)
	if err != nil {
		t.Fatal("error loading config " + err.Error())
	}

	setting := config.Lookup("struct")
	if setting.IsNil() {
		t.Fatalf("wrong is nil value")
	}

	setting = config.Lookup("notpresent")
	if !setting.IsNil() {
		t.Fatalf("wrong is nil value")
	}
}

func TestLookupSetting(t *testing.T) {
	confstr := "struct: { aint=1; astring=\"test\"; afloat=13.37; abool=true; }"
	config := NewConfig()
	err := config.ReadString(confstr)
	if err != nil {
		t.Fatal("error loading config " + err.Error())
	}

	setting := config.Lookup("struct")

	// and now test the settings
	aint, err := setting.LookupInt("aint")
	if err != nil {
		t.Fatal("error getting setting " + err.Error())
	}
	if aint != 1 {
		t.Fatal("wrong int value")
	}

	astr, err := setting.LookupString("astring")
	if err != nil {
		t.Fatal("error getting setting " + err.Error())
	}
	if astr != "test" {
		t.Fatal("wrong string value")
	}

	afloat, err := setting.LookupFloat("afloat")
	if err != nil {
		t.Fatal("error getting setting " + err.Error())
	}
	if afloat != 13.37 {
		t.Fatal("wrong float value")
	}

	abool, err := setting.LookupBool("abool")
	if err != nil {
		t.Fatal("error getting setting " + err.Error())
	}
	if abool != true {
		t.Fatal("wrong bool value")
	}
}

func TestSetSetting(t *testing.T) {
	confstr := "struct: { aint=1; astring=\"test\"; afloat=13.37; abool=true; }"
	config := NewConfig()
	err := config.ReadString(confstr)
	if err != nil {
		t.Fatal("error loading config " + err.Error())
	}
	setting := config.Lookup("struct")

	// Read / Write / Read back of all types

	//Int
	aint, err := setting.LookupInt("aint")
	if err != nil {
		t.Fatal("error getting setting " + err.Error())
	}
	if aint != 1 {
		t.Fatal("wrong int value")
	}
	set := config.Lookup("struct.aint")
	set.SetInt(10)
	aint, err = setting.LookupInt("aint")
	if aint != 10 {
		t.Fatal("wrong int value")
	}

	//String
	astr, err := setting.LookupString("astring")
	if err != nil {
		t.Fatal("error getting setting " + err.Error())
	}
	if astr != "test" {
		t.Fatal("wrong string value")
	}

	set = config.Lookup("struct.astring")
	set.SetString("hello")
	astr, err = setting.LookupString("astring")
	if astr != "hello" {
		t.Fatal("wrong string value")
	}

	//Float
	afloat, err := setting.LookupFloat("afloat")
	if err != nil {
		t.Fatal("error getting setting " + err.Error())
	}
	if afloat != 13.37 {
		t.Fatal("wrong float value")
	}
	set = config.Lookup("struct.afloat")
	set.SetFloat(10.5)
	afloat, err = setting.LookupFloat("afloat")
	if aint != 10 {
		t.Fatal("wrong float value")
	}

	//Bool
	abool, err := setting.LookupBool("abool")
	if err != nil {
		t.Fatal("error getting setting " + err.Error())
	}
	if abool != true {
		t.Fatal("wrong bool value")
	}
	set = config.Lookup("struct.abool")
	set.SetBool(false)
	abool, err = setting.LookupBool("abool")
	if abool != false {
		t.Fatal("wrong bool value")
	}
}

func TestAddStringSetting(t *testing.T) {

	confstr := "struct: { aint=1; astring=\"test\"; afloat=13.37; abool=true; }"
	config := NewConfig()
	err := config.ReadString(confstr)
	if err != nil {
		t.Fatal("error loading config " + err.Error())
	}

	//Create new setting
	newsetting := config.AddStringSettingToParent("struct", "anotherstring")
	if newsetting == nil {
		t.Fatal("error adding string setting")
	}
	//Set value
	newsetting.SetString("hello")

	//Readback
	setting := config.Lookup("struct")
	astr, err := setting.LookupString("anotherstring")
	if err != nil {
		t.Fatal("error getting setting " + err.Error())
	}

	if astr != "hello" {
		t.Fatal("wrong string value")
	}

}

func TestAddIntSetting(t *testing.T) {

	confstr := "struct: { aint=1; astring=\"test\"; afloat=13.37; abool=true; }"
	config := NewConfig()
	err := config.ReadString(confstr)
	if err != nil {
		t.Fatal("error loading config " + err.Error())
	}

	//Create new setting
	newsetting := config.AddIntSettingToParent("struct", "anotherint")
	if newsetting == nil {
		t.Fatal("error adding int setting")
	}
	//Set value
	newsetting.SetInt(22)

	//Readback
	setting := config.Lookup("struct")
	aint, err := setting.LookupInt("anotherint")
	if err != nil {
		t.Fatal("error getting int " + err.Error())
	}

	if aint != 22 {
		t.Fatal("wrong string value")
	}
}

func TestAddFloatSetting(t *testing.T) {

	confstr := "struct: { aint=1; astring=\"test\"; afloat=13.37; abool=true; }"
	config := NewConfig()
	err := config.ReadString(confstr)
	if err != nil {
		t.Fatal("error loading config " + err.Error())
	}

	//Create new setting
	newsetting := config.AddFloatSettingToParent("struct", "anotherfloat")
	if newsetting == nil {
		t.Fatal("error adding float setting")
	}
	//Set value
	newsetting.SetFloat(16.89)

	//Readback
	setting := config.Lookup("struct")
	afloat, err := setting.LookupFloat("anotherfloat")
	if err != nil {
		t.Fatal("error getting float " + err.Error())
	}

	if afloat != 16.89 {
		t.Fatal("wrong float value")
	}

}

func TestAddBoolSetting(t *testing.T) {

	confstr := "struct: { aint=1; astring=\"test\"; afloat=13.37; abool=true; }"
	config := NewConfig()
	err := config.ReadString(confstr)
	if err != nil {
		t.Fatal("error loading config " + err.Error())
	}

	//Create new setting
	newsetting := config.AddBoolSettingToParent("struct", "anotherbool")
	if newsetting == nil {
		t.Fatal("error adding bool setting")
	}
	//Set value
	newsetting.SetBool(false)

	//Readback
	setting := config.Lookup("struct")
	abool, err := setting.LookupBool("anotherbool")
	if err != nil {
		t.Fatal("error getting bool " + err.Error())
	}

	if abool != false {
		t.Fatal("wrong bool value")
	}

}
