package value

import (
	"github.com/lanvard/errors"
	"github.com/lanvard/support"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_bool_true_from_true(t *testing.T) {
	assert.True(t, support.NewValue(true).Bool())
}

func Test_bool_true_from_false(t *testing.T) {
	assert.False(t, support.NewValue(false).Bool())
}

func Test_bool_true_from_int_one(t *testing.T) {
	assert.True(t, support.NewValue(1).Bool())
}

func Test_bool_true_from_int_two(t *testing.T) {
	assert.False(t, support.NewValue(2).Bool())
}

func Test_bool_true_from_int_zero(t *testing.T) {
	assert.False(t, support.NewValue(0).Bool())
}

func Test_bool_true_from_string_one(t *testing.T) {
	assert.True(t, support.NewValue("1").Bool())
}

func Test_bool_true_from_string_true(t *testing.T) {
	assert.True(t, support.NewValue("true").Bool())
}

func Test_bool_true_from_string_on(t *testing.T) {
	assert.True(t, support.NewValue("on").Bool())
}

func Test_bool_true_from_string_yes(t *testing.T) {
	assert.True(t, support.NewValue("yes").Bool())
}

func Test_bool_true_with_error_should_panic(t *testing.T) {
	assert.PanicsWithError(t, "the error", func() {
		support.NewValueE("yes", errors.New("the error")).Bool()
	})
}

func Test_bool_with_false_and_error(t *testing.T) {
	value, err := support.NewValueE(false, errors.New("the error")).BoolE()

	assert.False(t, value)
	assert.EqualError(t, err, "the error")
}

func Test_bool_with_true_and_error(t *testing.T) {
	value, err := support.NewValueE(true, errors.New("the error")).BoolE()

	assert.True(t, value)
	assert.EqualError(t, err, "the error")
}
