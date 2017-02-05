package elements

import (
	"../common"
	"strings"
)

func CreateElementText(input map[string]string, elmType string) string {
	inputType := `text`
	if elmType == "NUMBER" {
		inputType = `number`
	}

	attrRequired := ``
	classRequired := ``
	if input["required"] == "required" {
		attrRequired = ` required="required"`
		classRequired = ` <span class="kintone-required">*</span>`
	}

	t := `<div class="form-group"><label><h4>{{.label}}` + classRequired + `</h4><input type="` + inputType + `" class="form-control ` + strings.ToLower(elmType) + `" id="_kintone_control_{{.code}}" name="_kintone_control_{{.code}}" data-field_type="` + strings.ToUpper(elmType) + `" data-expression="" value="{{.defaultValue}}" maxlength="" min="" placeholder=""` + attrRequired + `></label></div>`
	return common.EvalTemplate(t, input)
}

func CreateElementTextarea(input map[string]string) string {
	attrRequired := ``
	classRequired := ``
	if input["required"] == "required" {
		attrRequired = ` required="required"`
		classRequired = ` <span class="kintone-required">*</span>`
	}

	t := `<div class="form-group"><h4>{{.label}}` + classRequired + `</h4><textarea class="form-control multi_line_text" id="_kintone_control_{{.code}}" name="_kintone_control_{{.code}}" data-field_type="MULTI_LINE_TEXT"` + attrRequired + `></textarea></div>`
	return common.EvalTemplate(t, input)
}

func createLabel(input map[string]string, elmType string) string {
	dataRequired := ` data-required="false"`
	classRequired := ``
	if input["required"] == "required" {
		dataRequired = ` data-required="true"`
		classRequired = ` <span class="kintone-required">*</span>`
	}
	return `<div class="form-group ` + strings.ToLower(elmType) + `" id="_kintone_control_{{.code}}" data-field_type="` + strings.ToUpper(elmType) + `"` + dataRequired + `><h4>{{.label}}` + classRequired + `</h4>`
}

func CreateElementRadio(input map[string]string, slice []string) string {
	t := createLabel(input, "radio_button")
	for pos, _ := range slice {
		t += `<div class="radio"><label><input type="radio" name="_kintone_control_{{.code}}" class="form-control" value="` + slice[pos] + `">` + slice[pos] + `</label></div>`
	}
	return common.EvalTemplate(t+"</div>", input)
}

func CreateElementCheckbox(input map[string]string, slice []string) string {
	t := createLabel(input, "check_box")
	for pos, _ := range slice {
		t += `<div class="checkbox"><label><input type="checkbox" name="_kintone_control_{{.code}}[]" class="form-control" value="` + slice[pos] + `">` + slice[pos] + `</label></div>`
	}
	return common.EvalTemplate(t+"</div>", input)
}

func CreateElementMultiSelect(input map[string]string, elmType string, slice []string) string {
	multiple := ``
	if elmType == "MULTI_SELECT" {
		multiple = `multiple="multiple" size="5"`
	}

	attrRequired := ``
	classRequired := ``
	if input["required"] == "required" {
		attrRequired = ` required="required"`
		classRequired = ` <span class="kintone-required">*</span>`
	}

	t := `<div class="form-group"><label><h4>{{.label}}` + classRequired + `</h4><select class="form-control ` + strings.ToLower(elmType) + `" id="_kintone_control_{{.code}}" name="_kintone_control_{{.code}}" data-field_type="` + strings.ToUpper(elmType) + `"` + attrRequired + multiple + `>`

	for pos, _ := range slice {
		t += "<option>" + slice[pos] + "</option>"
	}

	return common.EvalTemplate(t+"</select></label></div>", input)
}
