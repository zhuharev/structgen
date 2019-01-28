package structgen

import (
	"fmt"
	"strings"
)

type Name string

func (n Name) TitleName() string {
	if strings.HasSuffix(string(n), "id") {
		nn := strings.TrimRight(string(n), "id")
		nn = nn + "ID"
		return nn
	}
	return strings.Title(string(n))
}

func (n Name) PluralName() string {
	return string(n) + "s"
}

type Schema struct {
	Structs      []*Struct
	SharedFields []Field  `yaml:"sharedFields"`
	SharedTags   []string `yaml:"sharedTags"`
}

func (s *Schema) Init() {

	for _, strct := range s.Structs {
		for i := range strct.Fields {
			strct.Fields[i].parent = strct
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
				stru.Fields[i].Tags = append([]Tag{Tag{
					Name:  Name(tag),
					Value: string(f.Name),
				}}, stru.Fields[i].Tags...)
			}
		}
	}
}

type Struct struct {
	Name
	Fields []Field
	Tags   []string

	SharedFields []Field
}

func (f Struct) EnumFields() (fields []Field) {
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
	Tags   []Tag

	parent *Struct
}

type Tag struct {
	Name
	Value string
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
	return f.Type
}
