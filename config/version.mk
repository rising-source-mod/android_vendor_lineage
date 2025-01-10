LINEAGE_VERSION := RisingOS-$(RISING_BUILD_VERSION)

# Display version
LINEAGE_DISPLAY_VERSION := RisingOS-$(RISING_DISPLAY_VERSION)

# LineageOS version properties
PRODUCT_SYSTEM_PROPERTIES += \
    ro.lineage.version=$(LINEAGE_VERSION) \
    ro.lineage.display.version=$(LINEAGE_DISPLAY_VERSION) \
    ro.lineage.build.version=$(RISING_VERSION) \
    ro.lineage.releasetype=$(RISING_BUILDTYPE) \
    ro.modversion=$(LINEAGE_VERSION)
