package httpapi

type GetAccountRequest struct {
	Address            string             `json:"address"`
	Balance            int                `json:"balance"`
	CreateTime         int64              `json:"create_time"`
	NetWindowSize      int                `json:"net_window_size"`
	NetWindowOptimized bool               `json:"net_window_optimized"`
	AccountResource    AccountResource    `json:"account_resource"`
	OwnerPermission    OwnerPermission    `json:"owner_permission"`
	ActivePermission   []ActivePermission `json:"active_permission"`
	FrozenV2           []FrozenV2         `json:"frozenV2"`
	AssetOptimized     bool               `json:"asset_optimized"`
}

type AccountResource struct {
	EnergyWindowSize      int  `json:"energy_window_size"`
	EnergyWindowOptimized bool `json:"energy_window_optimized"`
}

type Keys struct {
	Address string `json:"address"`
	Weight  int    `json:"weight"`
}

type OwnerPermission struct {
	PermissionName string `json:"permission_name"`
	Threshold      int    `json:"threshold"`
	Keys           []Keys `json:"keys"`
}

type ActivePermission struct {
	Type           string `json:"type"`
	ID             int    `json:"id"`
	PermissionName string `json:"permission_name"`
	Threshold      int    `json:"threshold"`
	Operations     string `json:"operations"`
	Keys           []Keys `json:"keys"`
}

type FrozenV2 struct {
	Type string `json:"type,omitempty"`
}
