package main

import "fmt"

func main() {
	cm := 5123456
	km := cm / 100000
	kmMod := cm % 100000
	hm := kmMod / 10000
	hmMod := kmMod % 10000
	dam := hmMod / 1000
	damMod := hmMod % 1000
	m := damMod / 100
	mMod := damMod % 100
	dm := mMod / 10
	cm = dm % 10

	fmt.Println("KM", km)
	fmt.Println("HM", hm)
	fmt.Println("DAM", dam)
	fmt.Println("M", m)
	fmt.Println("DM", dm)
	fmt.Println("CM", cm)
	fmt.Println("5 mod 10", 5%10)
}
