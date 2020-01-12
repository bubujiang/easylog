package cfun

import "html/template"

func ShowModulesTags(mt map[string][]string) template.HTML  {
	moduleHtml := `<div class="form-group col-md-6">
					<label for="module">模块</label>
					<select class="form-control" style="width: 100%;" id="module" name="module" onchange="$('.tags').hide();$('#'+$(this).val()+'_tags').show();">`
	tagsHtml := ``
	for module, tags := range mt {
		moduleHtml += `<option value="`+module+`">`+module+`</option>`
		if tagsHtml == "" {
			tagsHtml += `<div class="form-group col-md-6 tags" id="`+module+`_tags">`
		} else {
			tagsHtml += `<div class="form-group col-md-6 tags" id="`+module+`_tags" style="display: none;">`
		}
		tagsHtml += `<label for="`+module+`tags">Tag</label>
					<select class="form-control" style="width: 100%;" id="`+module+`tags" name="tag">`
		for _, tag := range tags {
			tagsHtml += `<option value="`+tag+`">`+tag+`</option>`
		}
		tagsHtml += `</select></div>`
		//fmt.Printf("%s\t%d\n", name, age)
	}
	moduleHtml += `</select></div>`

	html := `<div class="form-row">`
	html += moduleHtml + tagsHtml
	html += `</div>`

	return template.HTML(html)
}
