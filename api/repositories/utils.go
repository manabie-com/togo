package repositories

import "go.mongodb.org/mongo-driver/bson/primitive"

// filterToObjectID modifies the filter value to ObjectID if the key of an item is in the slice of keys
func filterToObjectID(filter map[string]interface{}, keys ...string) error {
	for _, key := range keys {
		if id, ok := filter[key].(string); ok {
			idPrimitive, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				return err
			}
			filter[key] = idPrimitive
		}
	}

	return nil
}
