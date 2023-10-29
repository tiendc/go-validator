package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Group(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		i := 10
		errs := Validate(
			Group(
				NumGT(&i, 9),
				NumEQ(&i, 10),
			),
		)
		assert.Equal(t, 0, len(errs))
	})

	t.Run("failure", func(t *testing.T) {
		i := 10
		errs := Validate(
			Group(
				NumGT(&i, 9),
				NumGTE(&i, 11),
				NumIn(&i, 0),
			),
		)
		assert.Equal(t, 1, len(errs))
		assert.Equal(t, "group", errs[0].Type())
		assert.Equal(t, 2, len(errs[0].Unwrap()))
	})
}

func Test_OneOf(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		i := 10
		errs := Validate(
			OneOf(
				NumIn(&i, 0, 1, 2),
				NumGT(&i, 9),
				NumGT(&i, 5),
				NumLT(&i, 0),
			),
		)
		assert.Equal(t, 0, len(errs))
	})

	t.Run("failure", func(t *testing.T) {
		i := 10
		errs := Validate(
			OneOf(
				NumIn(&i, 0, 1, 2),
				NumGT(&i, 10),
				NumLT(&i, 0),
			),
		)
		assert.Equal(t, 1, len(errs))
		assert.Equal(t, "one_of", errs[0].Type())
	})
}

func Test_ExactOneOf(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		i := 10
		errs := Validate(
			ExactOneOf(
				NumIn(&i, 0, 1, 2),
				NumGT(&i, 9),
				NumLT(&i, 0),
			),
		)
		assert.Equal(t, 0, len(errs))
	})

	t.Run("failure", func(t *testing.T) {
		i := 10
		errs := Validate(
			ExactOneOf(
				NumIn(&i, 0, 1, 2),
				NumGT(&i, 9),
				NumGT(&i, 5),
				NumGT(&i, 3),
				NumLT(&i, 0),
			),
		)
		assert.Equal(t, 1, len(errs))
		assert.Equal(t, "exact_one_of", errs[0].Type())
	})
}

func Test_NotOf(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		i := 10
		errs := Validate(
			NotOf(
				NumEQ(&i, 9),
				NumGT(&i, 11),
				NumLT(&i, 10),
			),
		)
		assert.Equal(t, 0, len(errs))
	})

	t.Run("failure", func(t *testing.T) {
		i := 10
		errs := Validate(
			NotOf(
				NumEQ(&i, 9),
				NumGTE(&i, 10),
				NumLT(&i, 10),
			),
		)
		assert.Equal(t, 1, len(errs))
		assert.Equal(t, "not_of", errs[0].Type())
	})
}

func Test_If(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		i := 10
		errs := Validate(
			If(i == 10).OnError(SetCustomKey("i must be 10")),
		)
		assert.Equal(t, 0, len(errs))
	})

	t.Run("failure", func(t *testing.T) {
		i := 11
		errs := Validate(
			If(i == 10).OnError(SetCustomKey("i must be 10")),
		)
		assert.Equal(t, 1, len(errs))
		assert.Equal(t, "if", errs[0].Type())
	})
}

func Test_When(t *testing.T) {
	t.Run("success then", func(t *testing.T) {
		i1 := 1
		i2 := 10
		errs := Validate(
			When(NumEQ(&i1, 1)).Then(
				NumGT(&i2, 1),
				NumLT(&i2, 100),
			).Else(
				NumEQ(&i2, 1),
			),
		)
		assert.Equal(t, 0, len(errs))
	})

	t.Run("success else", func(t *testing.T) {
		i1 := 1
		i2 := 10
		errs := Validate(
			When(NumEQ(&i1, 2)).Then(
				NumEQ(&i2, 100),
			).Else(
				NumEQ(&i2, 10),
			),
		)
		assert.Equal(t, 0, len(errs))
	})

	t.Run("failure then", func(t *testing.T) {
		i1 := 1
		i2 := 10
		errs := Validate(
			When(NumEQ(&i1, 1)).Then(
				NumGT(&i2, 100),
			).Else(
				NumEQ(&i2, 10),
			),
		)
		assert.Equal(t, 1, len(errs))
		assert.Equal(t, "gt", errs[0].Type())
	})

	t.Run("failure else", func(t *testing.T) {
		i1 := 1
		i2 := 10
		errs := Validate(
			When(NumEQ(&i1, 2)).Then(
				NumGT(&i2, 100),
			).Else(
				NumEQ(&i2, 100),
			),
		)
		assert.Equal(t, 1, len(errs))
		assert.Equal(t, "eq", errs[0].Type())
	})
}

func Test_Case(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		i1 := 1
		i2 := 10
		errs := Validate(
			Case(
				When(NumEQ(&i1, 1)).Then(NumGT(&i2, 0)),
				When(NumEQ(&i1, 2)).Then(NumLT(&i2, 100)),
			).Default(
				NumEQ(&i2, 1),
			),
		)
		assert.Equal(t, 0, len(errs))
	})

	t.Run("success case default", func(t *testing.T) {
		i1 := 3
		i2 := 10
		errs := Validate(
			Case(
				When(NumEQ(&i1, 1)).Then(NumGT(&i2, 0)),
				When(NumEQ(&i1, 2)).Then(NumLT(&i2, 100)),
			).Default(
				NumEQ(&i2, 10),
			),
		)
		assert.Equal(t, 0, len(errs))
	})

	t.Run("failure case", func(t *testing.T) {
		i1 := 2
		i2 := 100
		errs := Validate(
			Case(
				When(NumEQ(&i1, 1)).Then(NumGT(&i2, 0)),
				When(NumEQ(&i1, 2)).Then(NumLT(&i2, 10)),
			).Default(
				NumEQ(&i2, 10),
			),
		)
		assert.Equal(t, 1, len(errs))
		assert.Equal(t, "lt", errs[0].Type())
	})

	t.Run("failure case default", func(t *testing.T) {
		i1 := 3
		i2 := 100
		errs := Validate(
			Case(
				When(NumEQ(&i1, 1)).Then(NumGT(&i2, 0)),
				When(NumEQ(&i1, 2)).Then(NumLT(&i2, 10)),
			).Default(
				NumEQ(&i2, 10),
			),
		)
		assert.Equal(t, 1, len(errs))
		assert.Equal(t, "eq", errs[0].Type())
	})
}
