package gen

import (
	"fmt"
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

func NewDO(model interface{}) func(db *gorm.DB, tableName string, opts ...DOOption) *DO {
	modelType := reflect.TypeOf(model)
	if modelType.Kind() == reflect.Ptr {
		modelType = modelType.Elem()
	}

	return func(db *gorm.DB, tableName string, opts ...DOOption) *DO {
		d := DO{}

		// Force to copy Statement
		d.db = db.WithContext(db.Statement.Context)

		config := &DOConfig{}
		for _, opt := range opts {
			if opt != nil {
				if applyErr := opt.Apply(config); applyErr != nil {
					panic(applyErr)
				}
			}
		}

		d.DOConfig = config

		d.modelType = modelType
		err := d.db.Statement.Parse(model)
		if err != nil {
			panic(fmt.Errorf("Cannot parse model: %+v\n%w", model, err))
		}
		d.tableName = d.db.Statement.Schema.Table

		return &d
	}
}
