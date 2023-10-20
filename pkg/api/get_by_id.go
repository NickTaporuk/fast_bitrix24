package api

import "strconv"

// GetByID calls a Bitrix24 API method that returns a single record by ID
func (bx24 *ExtendedBitrix24) GetByID(method string, id int) (map[string]interface{}, error) {
	params := map[string]string{
		"ID": strconv.Itoa(id), // Convert int ID to string
	}
	return bx24.CallMethod(method, params)
}
