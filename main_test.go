package main

import "testing"

func Test_greetPerson(t *testing.T) {
	name := "Saya"
	mustName := "Hello Saya"

	if returnName := greetPerson(name); returnName != mustName {
		t.Errorf("GreetPerson () = %q, want %q ", returnName, mustName)
	}

}
