package structgen

import (
	"fmt"
	"log"
	"strings"
)

type Name string

func (n Name) TitleName() string {
	if strings.HasSuffix(strings.ToLower(string(n)), "id") {
		nn := n[:len(n)-2] + "ID"
		return strings.Title(string(nn))
	}
	return strings.Title(string(n))
}

func (n Name) PluralName() string {
	return string(n) + "s"
}

type Schema struct {
	Structs      []*Struct
	SharedFields []*Field `yaml:"sharedFields"`
	SharedTags   []string `yaml:"sharedTags"`
}

func (s *Schema) Init() {
	for _, strct := range s.Structs {
		for i := range strct.Fields {
			strct.Fields[i].parent = strct
			for ti := range strct.Fields[i].Tags {
				strct.Fields[i].Tags[ti].parent = strct.Fields[i]
			}
		}
		for _, f := range s.SharedFields {
			field := f
			field.parent = strct
			strct.Fields = append(strct.Fields, field)
		}
	}

	for _, tag := range s.SharedTags {
		for _, stru := range s.Structs {
			for i, f := range stru.Fields {
				var alreadyExist = false
				for _, oldTag := range stru.Fields[i].Tags {
					if string(oldTag.Name) == tag {
						alreadyExist = true
					}
				}
				if alreadyExist {
					continue
				}
				stru.Fields[i].Tags = append([]*Tag{&Tag{
					Name:  Name(tag),
					Value: f.TagValue(), //string(f.Name),

					parent: f,
				}}, stru.Fields[i].Tags...)
			}
		}
	}

	for _, strct := range s.Structs {
		for _, f := range strct.Fields {
			f.parent = strct
			for _, t := range f.Tags {
				t.parent = f
			}
		}
	}
}

type Struct struct {
	Name
	Fields []*Field
	Tags   []string

	SharedFields []*Field
}

func (f Struct) EnumFields() (fields []*Field) {
	for _, field := range f.Fields {
		if field.Type == FieldEnum {
			fields = append(fields, field)
		}
	}
	return
}

const (
	FieldEnum = "enum"
	FieldTime = "time"
)

type Field struct {
	Name
	Type   string
	Consts []Const
	Tags   []*Tag

	parent *Struct
}

type Tag struct {
	Name
	Value string

	parent *Field
}

func (t Tag) FmtValue() string {
	val := t.Value
	if strings.Contains(t.Value, "{kind}") {
		log.Printf("%#v %#v", t.parent, t.parent)
		if t.parent != nil && t.parent.parent != nil {
			val = strings.Replace(t.Value, "{kind}", strings.ToLower(string(t.parent.parent.Name)), 1)
		}
	}
	return val
}

type Const struct {
	Name
}

func (f Field) ComputedType() string {
	if f.Type == "enum" {
		return fmt.Sprintf("%s%s", f.parent.TitleName(), f.TitleName())
	}
	if f.Type == "time" {
		return "time.Time"
	}
	if f.Type == "params" {
		return "struct{}"
	}
	return f.Type
}

func (f Field) TagValue() string {
	return string(f.Name)
}
