// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ManagedIOSStoreApp Contains properties and inherited properties for an iOS store app that you can manage with an Intune app protection policy.
type ManagedIOSStoreApp struct {
	// ManagedApp is the base model of ManagedIOSStoreApp
	ManagedApp
	// BundleID The app's Bundle ID.
	BundleID *string `json:"bundleId,omitempty"`
	// AppStoreURL The Apple AppStoreUrl.
	AppStoreURL *string `json:"appStoreUrl,omitempty"`
	// ApplicableDeviceType The iOS architecture for which this app can run on.
	ApplicableDeviceType *IOSDeviceType `json:"applicableDeviceType,omitempty"`
	// MinimumSupportedOperatingSystem The value for the minimum supported operating system.
	MinimumSupportedOperatingSystem *IOSMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`
}