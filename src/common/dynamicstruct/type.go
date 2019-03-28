/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dynamicstruct

import (
	"fmt"
	"strings"
	
	"gopkg.in/go-playground/validator.v8"
)

// Field represent a field of dynamic struct
// FieldName is the field name that current field in target struct
type Field struct {
	Name      string
	Tags      []string
	FieldName string
}

func (fd *Field) GetTag() string {
	return strings.Join(fd.Tags, ",")
}

// DynamicStructure is a validate manager to dynamic struct, which is a common case in cmdb,
// users are about to add custom field to anything model.
type DynamicStructure struct {
	Fields []Field
	Raw    map[string]interface{}
}

func (ds *DynamicStructure) Validate() []error {
	errs := make([]error, 0)
	for _, field := range ds.Fields {
		val, ok := ds.Raw[field.FieldName]
		if ok == false {
			errs = append(errs, fmt.Errorf("field %s not found", field.FieldName))
			continue
		}
		var validate *validator.Validate
		validate = validator.New(nil)
		// errs := validate.Var(myEmail, "required,email")
		validate.Field(val, field.GetTag())
	}
	return errs
}
