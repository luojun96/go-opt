package main

type DecoratedVisitor struct {
	visitor    Visitor
	decorators []VisitorFunc
}

func NewDecoratedVisitor(v Visitor, fn ...VisitorFunc) Visitor {
	if len(fn) == 0 {
		return v
	}
	return DecoratedVisitor{visitor: v, decorators: fn}
}

// Visit implements Visitor
func (v DecoratedVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		if err != nil {
			return err
		}

		if err := fn(info, nil); err != nil {
			return err
		}

		for i := range v.decorators {
			if err := v.decorators[i](info, nil); err != nil {
				return err
			}
		}

		return nil
	})
}
