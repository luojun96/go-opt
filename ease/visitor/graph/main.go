package main

func main() {
	c := Circle{10}
	r := Rectangle{100, 200}
	t := Triangle{100, 100, 100}
	shapes := []Shape{c, r, t}

	for _, s := range shapes {
		s.accept(JsonVisitor)
		s.accept(XmlVisitor)
	}
}
