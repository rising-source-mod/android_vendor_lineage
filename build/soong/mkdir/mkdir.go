// Copyright 2025 The LineageOS Project.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mkdir

import (
	"android/soong/android"

	"path/filepath"
)

func MkdirFactory() android.Module {
	module := &Mkdir{}
	module.AddProperties(&module.properties)
	android.InitAndroidArchModule(module, android.DeviceSupported, android.MultilibCommon)
	return module
}

type MkdirProperties struct {
	Dir string `android:"arch_variant"`
}

type Mkdir struct {
	android.ModuleBase

	properties MkdirProperties

	installDir android.InstallPath
	output android.ModuleOutPath
}

func (this *Mkdir) GenerateAndroidBuildActions(ctx android.ModuleContext) {
	if filepath.Clean(this.properties.Dir) != this.properties.Dir {
		ctx.PropertyErrorf("dir", "Should be a clean filepath")
		return
	}

	out := android.PathForModuleOut(ctx, "out.txt")
	android.WriteFileRuleVerbatim(ctx, out, "")
	this.output = out

	this.installDir = android.PathForModuleInstall(ctx, "", this.properties.Dir)
}

func (this *Mkdir) AndroidMkEntries() []android.AndroidMkEntries {
	return []android.AndroidMkEntries{{
		Class: "FAKE",
		// Need at least one output file in order for this to take effect.
		OutputFile: android.OptionalPathForPath(this.output),
		Include: "vendor/lineage/build/core/mkdir.mk",
		ExtraEntries: []android.AndroidMkExtraEntriesFunc{
			func(ctx android.AndroidMkExtraEntriesContext, entries *android.AndroidMkEntries) {
				entries.SetString("LOCAL_SOONG_INSTALL_DIR", this.installDir.String())
			},
		},
	}}
}
