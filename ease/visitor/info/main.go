package main

func main() {
	info := Info{}
	var v Visitor = &info
	v = LogVisitor{v}
	v = NameVisitor{v}
	v = OtherThingsVisitor{v}

	loadFile := func(info *Info, err error) error {
		info.Name = "Luo Jun"
		info.Namespace = "Free"
		info.OtherThings = "I'm a developer."
		return nil
	}
	v.Visit(loadFile)
}
