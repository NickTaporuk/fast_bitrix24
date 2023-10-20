package api

import "strconv"

// GetAll calls a Bitrix24 API method that supports pagination and returns all records
func (bx24 *ExtendedBitrix24) GetAll(method string) ([]map[string]interface{}, error) {
	var allResults []map[string]interface{}
	start := 0

	for {
		params := map[string]string{
			"start": strconv.Itoa(start),
		}
		result, err := bx24.CallMethod(method, params)
		if err != nil {
			return nil, err
		}

		records, ok := result["result"].([]map[string]interface{})
		if !ok || len(records) == 0 {
			break
		}

		allResults = append(allResults, records...)

		// Handle pagination (assuming the API uses a "next" field to indicate there are more records)
		if next, exists := result["next"].(float64); exists {
			start = int(next)
		} else {
			break
		}
	}

	return allResults, nil
}
