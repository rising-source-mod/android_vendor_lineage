# Build fingerprint
ifneq ($(BUILD_FINGERPRINT),)
ADDITIONAL_SYSTEM_PROPERTIES += \
    ro.build.fingerprint=$(BUILD_FINGERPRINT)
endif

# crDroid System Version
ADDITIONAL_SYSTEM_PROPERTIES += \
    ro.rising.build.version=$(RISING_BUILD_VERSION) \
    ro.rising.display.version=$(RISING_DISPLAY_VERSION) \
    ro.rising.version=$(RISING_VERSION) \
    ro.modversion=$(RISING_VERSION) \
    ro.risinglegal.url=https://risingtechoss.github.io/legal.php

# LineageOS Platform SDK Version
ADDITIONAL_SYSTEM_PROPERTIES += \
    ro.lineage.build.version.plat.sdk=$(LINEAGE_PLATFORM_SDK_VERSION)

# LineageOS Platform Internal Version
ADDITIONAL_SYSTEM_PROPERTIES += \
    ro.lineage.build.version.plat.rev=$(LINEAGE_PLATFORM_REV)
