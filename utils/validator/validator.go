package validator

import (
	"fmt"
	"ginblog/utils/errmsg"
	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

//该包主要做数据验证 validator是go自带的数据验证工具
func Validate(data interface{}) (string, int) {
	validate := validator.New()
	uni := ut.New(zh_Hans_CN.New()) //翻译中文
	trans, _ := uni.GetTranslator("zh_Hans_CN")

	err := zh.RegisterDefaultTranslations(validate, trans) //对翻译方法进行注册
	if err != nil {
		fmt.Println("err:", err)
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string { //将模型的标签映射到返回
		lable := field.Tag.Get("label")
		return lable
	})

	err = validate.Struct(data) //验证结构体
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) { //类型断言写法
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCESS
}
