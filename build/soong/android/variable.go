package android
type Product_variables struct {
        Requires_synchronous_setsurface struct {
		Cflags []string
        }
        Disable_ashmem_tracking struct {
		Cflags []string
        }
	Allows_invalid_pthread struct {
		Cppflags []string
	}
	Egl_needs_handle struct {
		Cppflags []string
	}
	Egl_workaround_bug_10194508 struct {
		Cppflags []string
	}
	Exynos4_enhancements struct {
		Cflags []string
		Cppflags []string
	}
	Has_legacy_camera_hal1 struct {
		Cflags []string
        }
	Needs_legacy_camera_hal1_dyn_native_handle struct {
		Cppflags []string
	}
	Needs_text_relocations struct {
		Cppflags []string
	}
	Target_shim_libs struct {
		Cppflags []string
	}
	Uses_generic_camera_parameter_library struct {
		Srcs []string
	}
	Uses_nvidia_enhancements struct {
		Cppflags []string
	}
	Uses_qcom_bsp_legacy struct {
		Cppflags []string
	}
	Uses_qti_camera_device struct {
		Cppflags []string
		Shared_libs []string
	}
}

type ProductVariables struct {
	Requires_synchronous_setsurface    *bool `json:",omitempty"`
	Disable_ashmem_tracking    *bool `json:",omitempty"`
	Allows_invalid_pthread    *bool `json:",omitempty"`
	Egl_needs_handle    *bool `json:",omitempty"`
	Egl_workaround_bug_10194508    *bool `json:",omitempty"`
	Exynos4_enhancements    *bool `json:",omitempty"`
	Has_legacy_camera_hal1  *bool `json:",omitempty"`
	Needs_legacy_camera_hal1_dyn_native_handle  *bool `json:",omitempty"`
	Needs_text_relocations  *bool `json:",omitempty"`
	Specific_camera_parameter_library  *string `json:",omitempty"`
	Target_shim_libs  *string `json:",omitempty"`
	Uses_generic_camera_parameter_library  *bool `json:",omitempty"`
	Uses_nvidia_enhancements  *bool `json:",omitempty"`
	Uses_qcom_bsp_legacy  *bool `json:",omitempty"`
	Uses_qti_camera_device  *bool `json:",omitempty"`
}
