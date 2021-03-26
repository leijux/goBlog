package config

import "testing"

func Test_configGetString(t *testing.T) {
	s := GetString("gin.open")
	if s == "" {
		t.Error("vlue is empty")
	}
	t.Log(s)
}

func Test_configSet(t *testing.T) {
	err := Set("gin.prot", ":8081")
	if err != nil {
		t.Error(err)
	}
	if GetString("gin.prot") != ":8081" {
		t.Fatal("set err")
	}
	_ = Set("gin.prot", ":8080")
}
