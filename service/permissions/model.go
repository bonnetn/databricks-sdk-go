// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package permissions

// all definitions in this file are in alphabetical order

type AccessControlRequest struct {
	GroupName string `json:"group_name,omitempty"`
	// getPermissionLevel defaults to CAN_ATTACH_TO when it's not defined (for
	// PUT/PATCH) so you should use .permissionLevel.isDefined to verify that it
	// actually exists
	PermissionLevel AccessControlRequestPermissionLevel `json:"permission_level,omitempty"`

	ServicePrincipalName string `json:"service_principal_name,omitempty"`

	UserName string `json:"user_name,omitempty"`
}

// getPermissionLevel defaults to CAN_ATTACH_TO when it's not defined (for
// PUT/PATCH) so you should use .permissionLevel.isDefined to verify that it
// actually exists
type AccessControlRequestPermissionLevel string

const AccessControlRequestPermissionLevelCanAttachTo AccessControlRequestPermissionLevel = `CAN_ATTACH_TO`

const AccessControlRequestPermissionLevelCanBind AccessControlRequestPermissionLevel = `CAN_BIND`

const AccessControlRequestPermissionLevelCanEdit AccessControlRequestPermissionLevel = `CAN_EDIT`

const AccessControlRequestPermissionLevelCanEditMetadata AccessControlRequestPermissionLevel = `CAN_EDIT_METADATA`

const AccessControlRequestPermissionLevelCanManage AccessControlRequestPermissionLevel = `CAN_MANAGE`

const AccessControlRequestPermissionLevelCanManageProductionVersions AccessControlRequestPermissionLevel = `CAN_MANAGE_PRODUCTION_VERSIONS`

const AccessControlRequestPermissionLevelCanManageRun AccessControlRequestPermissionLevel = `CAN_MANAGE_RUN`

const AccessControlRequestPermissionLevelCanManageStagingVersions AccessControlRequestPermissionLevel = `CAN_MANAGE_STAGING_VERSIONS`

const AccessControlRequestPermissionLevelCanRead AccessControlRequestPermissionLevel = `CAN_READ`

const AccessControlRequestPermissionLevelCanRestart AccessControlRequestPermissionLevel = `CAN_RESTART`

const AccessControlRequestPermissionLevelCanRun AccessControlRequestPermissionLevel = `CAN_RUN`

const AccessControlRequestPermissionLevelCanUse AccessControlRequestPermissionLevel = `CAN_USE`

const AccessControlRequestPermissionLevelCanView AccessControlRequestPermissionLevel = `CAN_VIEW`

const AccessControlRequestPermissionLevelCanViewMetadata AccessControlRequestPermissionLevel = `CAN_VIEW_METADATA`

const AccessControlRequestPermissionLevelIsOwner AccessControlRequestPermissionLevel = `IS_OWNER`

type AccessControlResponse struct {
	AllPermissions []Permission `json:"all_permissions,omitempty"`

	GroupName string `json:"group_name,omitempty"`

	ServicePrincipalName string `json:"service_principal_name,omitempty"`

	UserName string `json:"user_name,omitempty"`
}

type GetObjectPermissionsRequest struct {
	ObjectId string ` path:"object_id"`

	ObjectType string ` path:"object_type"`
}

type GetPermissionLevelsRequest struct {
	RequestObjectId string ` path:"request_object_id"`

	RequestObjectType string ` path:"request_object_type"`
}

type GetPermissionLevelsResponse struct {
	PermissionLevels []PermissionsDescription `json:"permission_levels,omitempty"`
}

type ObjectPermissions struct {
	AccessControlList []AccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`
}

type Permission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel PermissionPermissionLevel `json:"permission_level,omitempty"`
}

type PermissionPermissionLevel string

const PermissionPermissionLevelCanAttachTo PermissionPermissionLevel = `CAN_ATTACH_TO`

const PermissionPermissionLevelCanBind PermissionPermissionLevel = `CAN_BIND`

const PermissionPermissionLevelCanEdit PermissionPermissionLevel = `CAN_EDIT`

const PermissionPermissionLevelCanEditMetadata PermissionPermissionLevel = `CAN_EDIT_METADATA`

const PermissionPermissionLevelCanManage PermissionPermissionLevel = `CAN_MANAGE`

const PermissionPermissionLevelCanManageProductionVersions PermissionPermissionLevel = `CAN_MANAGE_PRODUCTION_VERSIONS`

const PermissionPermissionLevelCanManageRun PermissionPermissionLevel = `CAN_MANAGE_RUN`

const PermissionPermissionLevelCanManageStagingVersions PermissionPermissionLevel = `CAN_MANAGE_STAGING_VERSIONS`

const PermissionPermissionLevelCanRead PermissionPermissionLevel = `CAN_READ`

const PermissionPermissionLevelCanRestart PermissionPermissionLevel = `CAN_RESTART`

const PermissionPermissionLevelCanRun PermissionPermissionLevel = `CAN_RUN`

const PermissionPermissionLevelCanUse PermissionPermissionLevel = `CAN_USE`

const PermissionPermissionLevelCanView PermissionPermissionLevel = `CAN_VIEW`

const PermissionPermissionLevelCanViewMetadata PermissionPermissionLevel = `CAN_VIEW_METADATA`

const PermissionPermissionLevelIsOwner PermissionPermissionLevel = `IS_OWNER`

type PermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel PermissionsDescriptionPermissionLevel `json:"permission_level,omitempty"`
}

type PermissionsDescriptionPermissionLevel string

const PermissionsDescriptionPermissionLevelCanAttachTo PermissionsDescriptionPermissionLevel = `CAN_ATTACH_TO`

const PermissionsDescriptionPermissionLevelCanBind PermissionsDescriptionPermissionLevel = `CAN_BIND`

const PermissionsDescriptionPermissionLevelCanEdit PermissionsDescriptionPermissionLevel = `CAN_EDIT`

const PermissionsDescriptionPermissionLevelCanEditMetadata PermissionsDescriptionPermissionLevel = `CAN_EDIT_METADATA`

const PermissionsDescriptionPermissionLevelCanManage PermissionsDescriptionPermissionLevel = `CAN_MANAGE`

const PermissionsDescriptionPermissionLevelCanManageProductionVersions PermissionsDescriptionPermissionLevel = `CAN_MANAGE_PRODUCTION_VERSIONS`

const PermissionsDescriptionPermissionLevelCanManageRun PermissionsDescriptionPermissionLevel = `CAN_MANAGE_RUN`

const PermissionsDescriptionPermissionLevelCanManageStagingVersions PermissionsDescriptionPermissionLevel = `CAN_MANAGE_STAGING_VERSIONS`

const PermissionsDescriptionPermissionLevelCanRead PermissionsDescriptionPermissionLevel = `CAN_READ`

const PermissionsDescriptionPermissionLevelCanRestart PermissionsDescriptionPermissionLevel = `CAN_RESTART`

const PermissionsDescriptionPermissionLevelCanRun PermissionsDescriptionPermissionLevel = `CAN_RUN`

const PermissionsDescriptionPermissionLevelCanUse PermissionsDescriptionPermissionLevel = `CAN_USE`

const PermissionsDescriptionPermissionLevelCanView PermissionsDescriptionPermissionLevel = `CAN_VIEW`

const PermissionsDescriptionPermissionLevelCanViewMetadata PermissionsDescriptionPermissionLevel = `CAN_VIEW_METADATA`

const PermissionsDescriptionPermissionLevelIsOwner PermissionsDescriptionPermissionLevel = `IS_OWNER`

type SetObjectPermissionsRequest struct {
	AccessControlList []AccessControlRequest `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty" path:"object_id"`

	ObjectType string `json:"object_type,omitempty" path:"object_type"`
}

type UpdateObjectPermissionsRequest struct {
	AccessControlList []AccessControlRequest `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty" path:"object_id"`

	ObjectType string `json:"object_type,omitempty" path:"object_type"`
}
