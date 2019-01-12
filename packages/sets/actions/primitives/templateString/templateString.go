package templateString

import (
    "bytes"
    "text/template"
)

type TemplateString string

func (y TemplateString) Template(dicts ...map[string]string) (string, error) {
    var dict map[string]string = map[string]string{}
    for _,d := range dicts {
        for k,v := range d {
            dict[k] = v
        }
    }
    var buf *bytes.Buffer = bytes.NewBuffer([]byte{})
    var t *template.Template
    if te,e := template.New("t").Parse(string(y)); e != nil {
        return "",e
    } else {
        t = te
    }
    if e := t.Execute(buf, dict); e != nil {
        return "",e
    }
    return buf.String(),nil
}
