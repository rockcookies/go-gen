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

func (d *DO) UnsafeSetDB(db *gorm.DB) {
	d.db = db
}

func (d *DO) UnsafeSetAlias(alias string) {
	d.alias = alias
}

func (d *DO) UnsafeSetModelType(modelType reflect.Type) {
	d.modelType = modelType
}

func (d *DO) UnsafeSetTableName(tableName string) {
	d.tableName = tableName
}
