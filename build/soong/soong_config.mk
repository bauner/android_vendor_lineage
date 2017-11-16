# Insert new variables inside the Lineage structure
lineage_soong:
	$(hide) mkdir -p $(dir $@)
	$(hide) (\
	echo '{'; \
	echo '"Lineage": {'; \
	echo '    "Egl_needs_handle": $(if $(filter true,$(BOARD_EGL_NEEDS_HANDLE_VALUE)),true,false),'; \
	echo '    "Egl_workaround_bug_10194508": $(if $(filter true,$(BOARD_EGL_WORKAROUND_BUG_10194508)),true,false),'; \
	echo '    "Exynos4_enhancements": $(if $(filter true,$(TARGET_NEEDS_EXYNOS4_ENHANCEMENTS)),true,false),'; \
	echo '    "Needs_text_relocations": $(if $(filter true,$(TARGET_NEEDS_PLATFORM_TEXT_RELOCATIONS)),true,false),'; \
	echo '    "Specific_camera_parameter_library": "$(TARGET_SPECIFIC_CAMERA_PARAMETER_LIBRARY)",'; \
	echo '    "Target_shim_libs": "$(subst $(space),:,$(TARGET_LD_SHIM_LIBS))",'; \
	echo '    "Uses_generic_camera_parameter_library": $(if $(TARGET_SPECIFIC_CAMERA_PARAMETER_LIBRARY),false,true),'; \
	echo '    "Uses_qcom_bsp_legacy": $(if $(filter true,$(TARGET_USES_QCOM_BSP_LEGACY)),true,false)'; \
	echo '},'; \
	echo '') > $(SOONG_VARIABLES_TMP)
