/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云 - PaaS 平台 (BlueKing - PaaS System) available.
 * Copyright (C) 2017 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 *     http://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 * to the current version of the project delivered to anyone in the future.
 */

package vcs

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Types", func() {

	Context("Test Files", func() {
		It("AsTree", func() {
			files := Files{
				{Action: FileActionAdded, Path: "README.md"},
				{Action: FileActionAdded, Path: "webfe/static/example.css"},
				{Action: FileActionAdded, Path: "webfe/static/example.js"},
				{Action: FileActionDeleted, Path: "api/main.py"},
				{Action: FileActionDeleted, Path: "webfe/templates/example.html"},
				{Action: FileActionModified, Path: "backend/main.go"},
				{Action: FileActionModified, Path: "backend/types.go"},
				{Action: FileActionModified, Path: "docs/example.txt"},
			}

			excepted := DirTree{
				Name: "/",
				Dirs: []*DirTree{
					{
						Name: "api", Files: Files{
							{Action: FileActionDeleted, Path: "main.py"},
						},
					},
					{
						Name: "backend", Files: Files{
							{Action: FileActionModified, Path: "main.go"},
							{Action: FileActionModified, Path: "types.go"},
						},
					},
					{
						Name: "docs", Files: Files{
							{Action: FileActionModified, Path: "example.txt"},
						},
					},
					{
						Name: "webfe", Dirs: []*DirTree{
							{
								Name: "static", Files: Files{
									{Action: FileActionAdded, Path: "example.css"},
									{Action: FileActionAdded, Path: "example.js"},
								},
							},
							{
								Name: "templates", Files: Files{
									{Action: FileActionDeleted, Path: "example.html"},
								},
							},
						},
					},
				},
				Files: Files{
					{Action: FileActionAdded, Path: "README.md"},
				},
			}
			Expect(files.AsTree()).To(Equal(excepted))
		})
	})
})
