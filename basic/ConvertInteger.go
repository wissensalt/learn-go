package main

func main() {
	var value32 int32 = 10000
	var value64 = int64(value32)
	var value8 = int8(value64)

	println("Value 32 : ", value32)
	println("Value 64 : ", value64)
	println("Value 8 : ", value8)
}
