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

package phase

import (
	"context"
	"os/exec"
	"path/filepath"
)

// MakeDetectorStep build the detector step
// detector will generate group.toml(to groupPath) and plan.toml(to planPath) based on order.toml(in orderPath) and source code(in appDir)
func MakeDetectorStep(
	ctx context.Context,
	lifecycleDir, appDir, orderPath, groupPath, planPath, layersDir, logLevel string, uid, gid uint32,
) Step {
	args := []string{
		"-app", appDir,
		"-order", orderPath,
		"-group", groupPath,
		"-plan", planPath,
		"-layers", layersDir,
		"-log-level", logLevel,
	}
	cmd := exec.CommandContext(ctx, filepath.Join(lifecycleDir, "detector"), args...)
	return makeStep("Detect", "Detecting Buildpacks...", cmd, WithUser(uid, gid))
}
