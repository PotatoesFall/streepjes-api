// Code generated by "enumer -type Role -linecomment -sql -json"; DO NOT EDIT.

//
package users

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

const _RoleName = "not_authorizedbartenderadmin"

var _RoleIndex = [...]uint8{0, 14, 23, 28}

func (i Role) String() string {
	if i < 0 || i >= Role(len(_RoleIndex)-1) {
		return fmt.Sprintf("Role(%d)", i)
	}
	return _RoleName[_RoleIndex[i]:_RoleIndex[i+1]]
}

var _RoleValues = []Role{0, 1, 2}

var _RoleNameToValueMap = map[string]Role{
	_RoleName[0:14]:  0,
	_RoleName[14:23]: 1,
	_RoleName[23:28]: 2,
}

// RoleString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func RoleString(s string) (Role, error) {
	if val, ok := _RoleNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Role values", s)
}

// RoleValues returns all values of the enum
func RoleValues() []Role {
	return _RoleValues
}

// IsARole returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Role) IsARole() bool {
	for _, v := range _RoleValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Role
func (i Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Role
func (i *Role) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Role should be a string, got %s", data)
	}

	var err error
	*i, err = RoleString(s)
	return err
}

func (i Role) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *Role) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	str, ok := value.(string)
	if !ok {
		bytes, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("value is not a byte slice")
		}

		str = string(bytes[:])
	}

	val, err := RoleString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
