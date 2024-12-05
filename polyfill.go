package gen

import (
	tmpl "github.com/rockcookies/go-gen/internal/template"
	"gorm.io/gorm"
	"reflect"
)

type Tmpl struct {
	Header                      string
	DIYMethod                   string
	CRUDMethod                  string
	CRUDMethodTest              string
	DIYMethodTestBasic          string
	DIYMethodTest               string
	Model                       string
	ModelMethod                 string
	DefaultQuery                string
	QueryMethod                 string
	QueryMethodTest             string
	TableQueryStruct            string
	TableQueryStructWithContext string
	TableQueryIface             string
}

func newTmpl() *Tmpl {
	return &Tmpl{
		Header:                      tmpl.Header,
		DIYMethod:                   tmpl.DIYMethod,
		CRUDMethod:                  tmpl.CRUDMethod,
		CRUDMethodTest:              tmpl.CRUDMethodTest,
		DIYMethodTestBasic:          tmpl.DIYMethodTestBasic,
		DIYMethodTest:               tmpl.DIYMethodTest,
		Model:                       tmpl.Model,
		ModelMethod:                 tmpl.ModelMethod,
		DefaultQuery:                tmpl.DefaultQuery,
		QueryMethod:                 tmpl.QueryMethod,
		QueryMethodTest:             tmpl.QueryMethodTest,
		TableQueryStruct:            tmpl.TableQueryStruct,
		TableQueryStructWithContext: tmpl.TableQueryStructWithContext,
		TableQueryIface:             tmpl.TableQueryIface,
	}
}

func NewDO(model interface{}) func(db *gorm.DB, tableName string) *DO {
	modelType := reflect.TypeOf(model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}

	return func(db *gorm.DB, tableName string) *DO {
		d := DO{}

		d.DOConfig = &DOConfig{}
		d.db = db
		d.modelType = modelType
		d.tableName = tableName

		return &d
	}
}
