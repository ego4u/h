package h

import (
	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
)

type JsonGet struct {
	Obj     any
	exprMap map[string]jp.Expr
}

func NewJsonGet(jsonString string) (*JsonGet, error) {
	obj, err := oj.ParseString(jsonString)
	if err != nil {
		return nil, err
	}
	return &JsonGet{
		Obj:     obj,
		exprMap: map[string]jp.Expr{},
	}, nil
}

func (r *JsonGet) Get(jsonPath string) ([]any, error) {
	x, ok := r.exprMap[jsonPath]
	if !ok {
		var err error
		x, err = jp.ParseString(jsonPath)
		if err != nil {
			return nil, err
		}
		r.exprMap[jsonPath] = x
	}
	return x.Get(r.Obj), nil
}

func (r *JsonGet) Get0(jsonPath string) any {
	v, err := r.Get(jsonPath)
	Poe(err)
	if len(v) == 0 {
		return nil
	}
	return v[0]
}

func (r *JsonGet) GetA(jsonPath string) any {
	v, err := r.Get(jsonPath)
	Poe(err)
	return v
}
