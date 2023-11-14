package main

type Describer interface {
	Describe()
}


func main() {
	var d1 Describer
	d1.Describe()

}
