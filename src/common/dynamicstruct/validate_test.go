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

import "testing"

func TestDynamicStructure_Validate(t *testing.T) {
	raw := map[string]interface{}{
		"field1": "abc",
	}
	ds := DynamicStructure{}
	ds.Fields = append(ds.Fields, Field{
		Name: "testtest",
		FieldName: "field1",
		tags: []string{"required", "lt=5", "gt=1"},
	})
	ds.Raw = raw
	err := ds.Validate()
	t.Errorf("validation result, err: %+v", err)
}
