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
