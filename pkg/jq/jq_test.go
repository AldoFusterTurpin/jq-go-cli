package jq

import "testing"

func TestValidateIdentityOperator(t *testing.T) {
	t.Run("return no error and true", func(t *testing.T) {
		in := "."
		want := true

		b, err := ValidateIdentityOperator(in)
		if err != nil {
			t.Error(err)
		}

		if b != want {
			t.Errorf("got %v, but wanted %v", b, want)
		}
	})

	t.Run("return no error and false due to empty string", func(t *testing.T) {
		in := ""
		want := false

		b, err := ValidateIdentityOperator(in)
		if err != nil {
			t.Error(err)
		}

		if b != want {
			t.Errorf("got %v, but wanted %v", b, want)
		}
	})

	t.Run("return no error and false due to wrong identity operator being two dots", func(t *testing.T) {
		in := ".."
		want := false

		b, err := ValidateIdentityOperator(in)
		if err != nil {
			t.Error(err)
		}

		if b != want {
			t.Errorf("got %v, but wanted %v", b, want)
		}
	})
}

func TestValidateAndGetIndexFromArrayExp(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		in := ".[0]"
		expectedIndex := 0
		index, err := ValidateAndGetIndexFromArrayExp(in)
		if err != nil {
			t.Error(err)
		}

		if index != expectedIndex {
			t.Errorf("got %v, but wanted %v", index, expectedIndex)
		}
	})

	t.Run("error due to duplicate expression", func(t *testing.T) {
		in := ".[0].[0]"
		expectedIndex := -1
		index, err := ValidateAndGetIndexFromArrayExp(in)
		if err == nil {
			t.Errorf("expected error but got no error")
		}

		if index != expectedIndex {
			t.Errorf("got %v, but wanted %v", index, expectedIndex)
		}
	})

	t.Run("error due to simple string instead of expected pattern", func(t *testing.T) {
		in := "wrong"
		expectedIndex := -1
		index, err := ValidateAndGetIndexFromArrayExp(in)
		if err == nil {
			t.Error("expected error but got no error")
		}

		if index != expectedIndex {
			t.Errorf("got %v, but wanted %v", index, expectedIndex)
		}
	})
	t.Run("error due to missing number", func(t *testing.T) {
		in := ".[]"
		expectedIndex := -1
		index, err := ValidateAndGetIndexFromArrayExp(in)
		if err == nil {
			t.Error("expected error but got no error")
		}

		if index != expectedIndex {
			t.Errorf("got %v, but wanted %v", index, expectedIndex)
		}
	})

	t.Run("also error", func(t *testing.T) {
		in := ".[_]"
		expectedIndex := -1
		index, err := ValidateAndGetIndexFromArrayExp(in)
		if err == nil {
			t.Error("expected error but got no error")
		}

		if index != expectedIndex {
			t.Errorf("got %v, but wanted %v", index, expectedIndex)
		}
	})
}

func TestGetIthElementFromArray(t *testing.T) {
	t.Run("get second element", func(t *testing.T) {
		f := []interface{}{
			"firstElement",
			"secondElement",
		}
		want := "secondElement"

		got, err := GetIthElementFromArray(f, 1)
		if err != nil {
			t.Error(err)
		}

		if got != want {
			t.Errorf("got %v, but wanted %v", got, want)
		}
	})

	t.Run("get index out of bound returns error", func(t *testing.T) {
		f := []interface{}{
			"firstElement",
			"secondElement",
		}

		got, err := GetIthElementFromArray(f, 2)
		if err == nil {
			t.Error(err)
		}

		if got != nil {
			t.Errorf("got %v, but wanted %v", got, nil)
		}
	})

	t.Run("get index when input is not an array returns error", func(t *testing.T) {
		f := map[string]interface{}{
			"firstElement":  "not an array",
			"secondElement": "hello hello ",
		}

		got, err := GetIthElementFromArray(f, 2)
		if err == nil {
			t.Error(err)
		}

		if got != nil {
			t.Errorf("got %v, but wanted %v", got, nil)
		}
	})

	t.Run("get index when is negative returns ", func(t *testing.T) {
		f := []interface{}{
			"firstElement",
			"secondElement",
		}

		got, err := GetIthElementFromArray(f, -1)
		if err == nil {
			t.Error(err)
		}

		if got != nil {
			t.Errorf("got %v, but wanted %v", got, nil)
		}
	})
}
