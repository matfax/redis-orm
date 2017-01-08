package parser

import "errors"

type Relation struct {
	Name      string
	StoreType string
	ValueType string
	ModelType string
	//! fields
	Fields []*Field
	//! owner
	Obj *MetaObject
}

func NewRelation(obj *MetaObject) *Relation {
	return &Relation{
		Name: obj.Name,
		Obj:  obj,
	}
}

func (r *Relation) build() error {
	switch r.StoreType {
	case "pair", "set", "list":
		r.Fields = make([]*Field, 2)
		f1 := NewField()
		f1.Obj = r.Obj
		f1.Name = "Key"
		f1.Type = "string"
		f1.Flags.Add("primary")
		r.Fields[0] = f1

		f2 := NewField()
		f2.Obj = r.Obj
		f2.Name = "Value"
		f2.Type = r.ValueType
		r.Fields[1] = f2

	case "zset":
		r.Fields = make([]*Field, 3)
		f1 := NewField()
		f1.Obj = r.Obj
		f1.Name = "Key"
		f1.Type = "string"
		f1.Flags.Add("primary")
		r.Fields[0] = f1

		f2 := NewField()
		f2.Obj = r.Obj
		f2.Name = "Score"
		f2.Type = "float64"
		r.Fields[0] = f2

		f3 := NewField()
		f3.Obj = r.Obj
		f3.Name = "Value"
		f3.Type = r.ValueType
		r.Fields[0] = f3
	case "geo":
		r.Fields = make([]*Field, 4)
		f1 := NewField()
		f1.Obj = r.Obj
		f1.Name = "Key"
		f1.Type = "string"
		f1.Flags.Add("primary")
		r.Fields[0] = f1

		f2 := NewField()
		f2.Obj = r.Obj
		f2.Name = "Longitude"
		f2.Type = "float64"
		r.Fields[0] = f2

		f3 := NewField()
		f3.Obj = r.Obj
		f3.Name = "Latitude"
		f3.Type = "float64"
		r.Fields[0] = f3

		f4 := NewField()
		f4.Obj = r.Obj
		f4.Name = "Value"
		f4.Type = r.ValueType
		r.Fields[0] = f4
	default:
		return errors.New("unsupport `store` for relation")
	}
	return nil
}

func (r *Relation) Read(data map[interface{}]interface{}) error {
	for k, v := range data {
		key := k.(string)
		switch key {
		case "storetype":
			r.StoreType = v.(string)
		case "valuetype":
			r.ValueType = v.(string)
		case "modeltype":
			r.ModelType = v.(string)
		}
	}
	return r.build()
}
