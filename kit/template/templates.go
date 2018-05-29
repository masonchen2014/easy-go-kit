package template

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

func init() {
	//create templateMap
	templateMap = make(map[string]*template.Template)
	//init FuncMap for templates
	fm := template.FuncMap{"GetVarListString": GetVarListString,
		"GenerateMethodString": GenerateMethodString,
		"UpperFirstLetter":     UpperFirstLetter,
		"LowerFirstLetter":     LowerFirstLetter,
	}
	//create templates and add it into the templateMap
	methodT := template.Must(template.New("method.tmpl").Funcs(fm).ParseFiles("template/templates/method.tmpl"))
	serviceT := template.Must(template.New("service.tmpl").Funcs(fm).ParseFiles("template/templates/service.tmpl"))
	structT := template.Must(template.New("struct.tmpl").Funcs(fm).ParseFiles("template/templates/struct.tmpl"))
	templateMap["methodT"] = methodT
	templateMap["serviceT"] = serviceT
	templateMap["structT"] = structT

}

var templateMap map[string]*template.Template

func generateTemplateString(templateName string, data interface{}) string {
	var b bytes.Buffer
	if t, ok := templateMap[templateName]; ok {
		err := t.Execute(&b, data)
		if err != nil {
			fmt.Println("executing template:", err)
			return ""
		}
		return b.String()
	}
	return ""
}

func GetVarListString(p []VarInfo) string {
	var str string
	for _, v := range p {
		str += v.Name + " " + v.Type + ","
	}
	str = strings.TrimRight(str, ",")
	return str
}

func UpperFirstLetter(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

func LowerFirstLetter(s string) string {
	return strings.ToLower(string(s[0])) + s[1:]
}

func GenerateServiceString(s *Service) string {
	return generateTemplateString("serviceT", s)
}

func GenerateMethodString(m *Method) string {
	return generateTemplateString("methodT", m)
}

func GenerateStructString(s *Struct) string {
	return generateTemplateString("structT", s)
}

type VarInfo struct {
	Name string
	Type string
}

type Method struct {
	Name        string
	ParamValues []VarInfo
	RetValues   []VarInfo
	Recvier     VarInfo
}

type Struct struct {
	Name    string
	Members []VarInfo
}

type Service struct {
	Name    string
	Methods []*Method
}
