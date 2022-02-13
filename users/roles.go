package users

import (
	"github.com/defintly/backend/database"
	"github.com/defintly/backend/types"
)

func GetPermissionsOfRole(roleId int) ([]*types.RolePermission, error) {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.RolePermissionType,
		"SELECT * FROM role_permissions WHERE id = $1", roleId)

	if err != nil {
		return nil, err
	}

	return slice.([]*types.RolePermission), err
}

func HasPermission(userId int, permission string) (bool, error) {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.IdInformationType,
		"SELECT user_role_mapping.user_id "+
			"FROM user_role_mapping "+
			"INNER JOIN roles ON user_role_mapping.role_id = roles.id"+
			"INNER JOIN role_permissions ON role_permissions.role_id = roles.id "+
			"WHERE user_role_mapping.user_id = $1 AND role_permissions.name = $2", userId, permission)

	if err != nil {
		return false, err
	}

	return len(slice.([]*types.IdInformation)) != 0, nil
}

func GetRoles() ([]*types.Role, error) {
	slice, err := database.QueryAsync(database.DefaultTimeout, types.RoleType, "SELECT * FROM roles")

	if err != nil {
		return nil, err
	}

	return slice.([]*types.Role), nil
}

func GetRolesIncludingPermissions() ([]*types.RoleIncludingPermissions, error) {
	simpleRoles, err := GetRoles()

	if err != nil {
		return nil, err
	}

	if len(simpleRoles) == 0 {
		return nil, nil
	}

	var roles []*types.RoleIncludingPermissions

	for _, role := range simpleRoles {
		slice, err := database.QueryAsync(database.DefaultTimeout, types.RolePermissionType,
			"SELECT * FROM role_permissions WHERE role = $1", role.Id)

		if err != nil {
			return nil, err
		}

		roles = append(roles, &types.RoleIncludingPermissions{
			Role:        role,
			Permissions: slice.([]*types.RolePermission),
		})
	}

	return roles, nil
}
