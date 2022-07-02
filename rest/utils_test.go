package rest

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const characters201 = "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem. Nulla consequat massa quis enim. Donec pede justo, fringilla vel, aliquet nec, vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet a, venenatis vitae, justo. Nullam dictum felis eu pede mollis pretium. Integer tincidunt. Cras dapibus. Vivamus elementum semper nisi. Aenean vulputate eleifend tellus. Aenean leo ligula, porttitor eu, consequat vitae, eleifend ac, enim. Aliquam lorem ante, dapibus in, viverra quis, feugiat a, tellus. Phasellus viverra nulla ut metus varius laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies nisi vel augue. Curabitur ullamcorper ultricies nisi. Nam eget dui. Etiam rhoncus. Maecenas tempus, tellus eget condimentum rhoncus, sem quam semper libero, sit amet adipiscing sem neque sed ipsum. Nam quam nunc, blandit vel, luctus pulvinar, hendrerit id, lorem. Maecenas nec odio et ante tincidunt tempus. Donec vitae sapien ut libero venenatis faucibus. Nullam quis ante. Etiam sit amet orci eget eros faucibus tincidunt. Duis leo. Sed fringilla mauris sit amet nibh. Donec sodales sagittis magna. Sed consequat, leo eget bibendum sodales, augue velit cursus nunc, quis."

func TestGetBeginningOfDay(t *testing.T) {
	date := getBeginningOfDay(time.Now())
	assert.NotNil(t, date, "The `date` should not be nil")
	assert.Equal(t, "Local", date.Location().String(), "The `date.Location()` should be UTC")
}
func TestValidateUsername(t *testing.T) {
	err := validateUsername("mariiia")

	assert.Nil(t, err, "The `err` should be nil")

	err = validateUsername("")

	assert.NotNil(t, err, "The `err` should not be nil")

	err = validateUsername(characters201)

	assert.NotNil(t, err, "The `err` should not be nil")
}

func TestValidateTaskDailyLimit(t *testing.T) {
	err := validateTaskDailyLimit(1)

	assert.Nil(t, err, "The `err` should be nil")

	err = validateTaskDailyLimit(0)

	assert.NotNil(t, err, "The `err` should not be nil")
}

func TestValidateTaskTitle(t *testing.T) {
	err := validateTaskTitle("Sample title")

	assert.Nil(t, err, "The `err` should be nil")

	err = validateTaskTitle("")

	assert.NotNil(t, err, "The `err` should not be nil")

	err = validateTaskTitle(characters201)

	assert.NotNil(t, err, "The `err` should not be nil")

}

func TestTrimLowerUsername(t *testing.T) {
	username := trimLowerUsername(" Mariiia ")

	assert.Equal(t, "mariiia", username, "The `username` should not be `mariiia`")
}
