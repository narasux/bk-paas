/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云 - PaaS 平台 (BlueKing - PaaS System) available.
 * Copyright (C) 2017 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 *	http://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 * to the current version of the project delivered to anyone in the future.
 */

package main

import (
	"os"

	"github.com/fatih/color"

	"github.com/TencentBlueKing/blueking-paas/client/cmd"
	"github.com/TencentBlueKing/blueking-paas/client/pkg/common/envs"
	"github.com/TencentBlueKing/blueking-paas/client/pkg/config"
)

func main() {
	// load global config ...
	if _, err := config.LoadConf(envs.ConfigFilePath); err != nil {
		color.Red("failed to load config, error: " + err.Error())
		os.Exit(1)
	}

	cmd.Execute()
}
