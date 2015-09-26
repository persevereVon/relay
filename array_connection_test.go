package graphql_relay_test

import (
	"github.com/chris-ramon/graphql-go/testutil"
	"github.com/sogko/graphql-relay-go"
	"reflect"
	"testing"
)

var arrayConnectionTestLetters = []interface{}{
	"A", "B", "C", "D", "E",
}

func TestConnectionFromArray_HandlesBasicSlicing_ReturnsAllElementsWithoutFilters(t *testing.T) {
	args := graphql_relay.NewConnectionArguments(nil)

	expected := &graphql_relay.Connection{
		Edges: []*graphql_relay.Edge{
			&graphql_relay.Edge{
				Node:   "A",
				Cursor: "YXJyYXljb25uZWN0aW9uOjA=",
			},
			&graphql_relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&graphql_relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&graphql_relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
			&graphql_relay.Edge{
				Node:   "E",
				Cursor: "YXJyYXljb25uZWN0aW9uOjQ=",
			},
		},
		PageInfo: graphql_relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjA=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjQ=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := graphql_relay.ConnectionFromArray(arrayConnectionTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, connection result diff: %v", testutil.Diff(expected, result))
	}
}
func TestConnectionFromArray_HandlesBasicSlicing_RespectsASmallerFirst(t *testing.T) {
	// Create connection arguments from map[string]interface{},
	// which you usually get from types.GQLParams.Args
	filter := map[string]interface{}{
		"first": 2,
	}
	args := graphql_relay.NewConnectionArguments(filter)

	// Alternatively, you can create connection arg the following way.
	// args := graphql_relay.NewConnectionArguments(filter)
	// args.First = 2

	expected := &graphql_relay.Connection{
		Edges: []*graphql_relay.Edge{
			&graphql_relay.Edge{
				Node:   "A",
				Cursor: "YXJyYXljb25uZWN0aW9uOjA=",
			},
			&graphql_relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
		},
		PageInfo: graphql_relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjA=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjE=",
			HasPreviousPage: false,
			HasNextPage:     true,
		},
	}

	result := graphql_relay.ConnectionFromArray(arrayConnectionTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, connection result diff: %v", testutil.Diff(expected, result))
	}
}
func TestConnectionFromArray_HandlesBasicSlicing_RespectsAnOverlyLargeFirst(t *testing.T) {

	filter := map[string]interface{}{
		"first": 10,
	}
	args := graphql_relay.NewConnectionArguments(filter)

	expected := &graphql_relay.Connection{
		Edges: []*graphql_relay.Edge{
			&graphql_relay.Edge{
				Node:   "A",
				Cursor: "YXJyYXljb25uZWN0aW9uOjA=",
			},
			&graphql_relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&graphql_relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&graphql_relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
			&graphql_relay.Edge{
				Node:   "E",
				Cursor: "YXJyYXljb25uZWN0aW9uOjQ=",
			},
		},
		PageInfo: graphql_relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjA=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjQ=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := graphql_relay.ConnectionFromArray(arrayConnectionTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, connection result diff: %v", testutil.Diff(expected, result))
	}
}
func TestConnectionFromArray_HandlesBasicSlicing_RespectsASmallerLast(t *testing.T) {

	filter := map[string]interface{}{
		"last": 2,
	}
	args := graphql_relay.NewConnectionArguments(filter)

	expected := &graphql_relay.Connection{
		Edges: []*graphql_relay.Edge{
			&graphql_relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
			&graphql_relay.Edge{
				Node:   "E",
				Cursor: "YXJyYXljb25uZWN0aW9uOjQ=",
			},
		},
		PageInfo: graphql_relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjM=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjQ=",
			HasPreviousPage: true,
			HasNextPage:     false,
		},
	}

	result := graphql_relay.ConnectionFromArray(arrayConnectionTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, connection result diff: %v", testutil.Diff(expected, result))
	}
}
func TestConnectionFromArray_HandlesBasicSlicing_RespectsAnOverlyLargeLast(t *testing.T) {

	filter := map[string]interface{}{
		"last": 10,
	}
	args := graphql_relay.NewConnectionArguments(filter)

	expected := &graphql_relay.Connection{
		Edges: []*graphql_relay.Edge{
			&graphql_relay.Edge{
				Node:   "A",
				Cursor: "YXJyYXljb25uZWN0aW9uOjA=",
			},
			&graphql_relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&graphql_relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&graphql_relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
			&graphql_relay.Edge{
				Node:   "E",
				Cursor: "YXJyYXljb25uZWN0aW9uOjQ=",
			},
		},
		PageInfo: graphql_relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjA=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjQ=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := graphql_relay.ConnectionFromArray(arrayConnectionTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, connection result diff: %v", testutil.Diff(expected, result))
	}
}

func TestConnectionFromArray_HandlesPagination_RespectsFirstAndAfter(t *testing.T) {

	filter := map[string]interface{}{
		"first": 2,
		"after": "YXJyYXljb25uZWN0aW9uOjE=",
	}
	args := graphql_relay.NewConnectionArguments(filter)

	expected := &graphql_relay.Connection{
		Edges: []*graphql_relay.Edge{
			&graphql_relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&graphql_relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
		},
		PageInfo: graphql_relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjI=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjM=",
			HasPreviousPage: false,
			HasNextPage:     true,
		},
	}

	result := graphql_relay.ConnectionFromArray(arrayConnectionTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, connection result diff: %v", testutil.Diff(expected, result))
	}
}
func TestConnectionFromArray_HandlesPagination_RespectsFirstAndAfterWithLongFirst(t *testing.T) {

	filter := map[string]interface{}{
		"first": 10,
		"after": "YXJyYXljb25uZWN0aW9uOjE=",
	}
	args := graphql_relay.NewConnectionArguments(filter)

	expected := &graphql_relay.Connection{
		Edges: []*graphql_relay.Edge{
			&graphql_relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&graphql_relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
			&graphql_relay.Edge{
				Node:   "E",
				Cursor: "YXJyYXljb25uZWN0aW9uOjQ=",
			},
		},
		PageInfo: graphql_relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjI=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjQ=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := graphql_relay.ConnectionFromArray(arrayConnectionTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, connection result diff: %v", testutil.Diff(expected, result))
	}
}
func TestConnectionFromArray_HandlesPagination_RespectsLastAndBefore(t *testing.T) {
	filter := map[string]interface{}{
		"last":   2,
		"before": "YXJyYXljb25uZWN0aW9uOjM=",
	}
	args := graphql_relay.NewConnectionArguments(filter)

	expected := &graphql_relay.Connection{
		Edges: []*graphql_relay.Edge{
			&graphql_relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&graphql_relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
		},
		PageInfo: graphql_relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjE=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjI=",
			HasPreviousPage: true,
			HasNextPage:     false,
		},
	}

	result := graphql_relay.ConnectionFromArray(arrayConnectionTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, connection result diff: %v", testutil.Diff(expected, result))
	}
}
func TestConnectionFromArray_HandlesPagination_RespectsLastAndBeforeWithLongLast(t *testing.T) {
	filter := map[string]interface{}{
		"last":   10,
		"before": "YXJyYXljb25uZWN0aW9uOjM=",
	}
	args := graphql_relay.NewConnectionArguments(filter)

	expected := &graphql_relay.Connection{
		Edges: []*graphql_relay.Edge{
			&graphql_relay.Edge{
				Node:   "A",
				Cursor: "YXJyYXljb25uZWN0aW9uOjA=",
			},
			&graphql_relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&graphql_relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
		},
		PageInfo: graphql_relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjA=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjI=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := graphql_relay.ConnectionFromArray(arrayConnectionTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, connection result diff: %v", testutil.Diff(expected, result))
	}
}
func TestConnectionFromArray_HandlesPagination_RespectsFirstAndAfterAndBefore_TooFew(t *testing.T) {
	filter := map[string]interface{}{
		"first":  2,
		"after":  "YXJyYXljb25uZWN0aW9uOjA=",
		"before": "YXJyYXljb25uZWN0aW9uOjQ=",
	}
	args := graphql_relay.NewConnectionArguments(filter)

	expected := &graphql_relay.Connection{
		Edges: []*graphql_relay.Edge{
			&graphql_relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&graphql_relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
		},
		PageInfo: graphql_relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjE=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjI=",
			HasPreviousPage: false,
			HasNextPage:     true,
		},
	}

	result := graphql_relay.ConnectionFromArray(arrayConnectionTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, connection result diff: %v", testutil.Diff(expected, result))
	}
}
func TestConnectionFromArray_HandlesPagination_RespectsFirstAndAfterAndBefore_TooMany(t *testing.T) {
	filter := map[string]interface{}{
		"first":  4,
		"after":  "YXJyYXljb25uZWN0aW9uOjA=",
		"before": "YXJyYXljb25uZWN0aW9uOjQ=",
	}
	args := graphql_relay.NewConnectionArguments(filter)

	expected := &graphql_relay.Connection{
		Edges: []*graphql_relay.Edge{
			&graphql_relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&graphql_relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&graphql_relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
		},
		PageInfo: graphql_relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjE=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjM=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := graphql_relay.ConnectionFromArray(arrayConnectionTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, connection result diff: %v", testutil.Diff(expected, result))
	}
}
func TestConnectionFromArray_HandlesPagination_RespectsFirstAndAfterAndBefore_ExactlyRight(t *testing.T) {
	filter := map[string]interface{}{
		"first":  3,
		"after":  "YXJyYXljb25uZWN0aW9uOjA=",
		"before": "YXJyYXljb25uZWN0aW9uOjQ=",
	}
	args := graphql_relay.NewConnectionArguments(filter)

	expected := &graphql_relay.Connection{
		Edges: []*graphql_relay.Edge{
			&graphql_relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&graphql_relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&graphql_relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
		},
		PageInfo: graphql_relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjE=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjM=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := graphql_relay.ConnectionFromArray(arrayConnectionTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, connection result diff: %v", testutil.Diff(expected, result))
	}
}
func TestConnectionFromArray_HandlesPagination_RespectsLastAndAfterAndBefore_TooFew(t *testing.T) {
	filter := map[string]interface{}{
		"last":   2,
		"after":  "YXJyYXljb25uZWN0aW9uOjA=",
		"before": "YXJyYXljb25uZWN0aW9uOjQ=",
	}
	args := graphql_relay.NewConnectionArguments(filter)

	expected := &graphql_relay.Connection{
		Edges: []*graphql_relay.Edge{
			&graphql_relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&graphql_relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
		},
		PageInfo: graphql_relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjI=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjM=",
			HasPreviousPage: true,
			HasNextPage:     false,
		},
	}

	result := graphql_relay.ConnectionFromArray(arrayConnectionTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, connection result diff: %v", testutil.Diff(expected, result))
	}
}
func TestConnectionFromArray_HandlesPagination_RespectsLasttAndAfterAndBefore_TooMany(t *testing.T) {
	filter := map[string]interface{}{
		"last":   4,
		"after":  "YXJyYXljb25uZWN0aW9uOjA=",
		"before": "YXJyYXljb25uZWN0aW9uOjQ=",
	}
	args := graphql_relay.NewConnectionArguments(filter)

	expected := &graphql_relay.Connection{
		Edges: []*graphql_relay.Edge{
			&graphql_relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&graphql_relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&graphql_relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
		},
		PageInfo: graphql_relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjE=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjM=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := graphql_relay.ConnectionFromArray(arrayConnectionTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, connection result diff: %v", testutil.Diff(expected, result))
	}
}
func TestConnectionFromArray_HandlesPagination_RespectsLastAndAfterAndBefore_ExactlyRight(t *testing.T) {
	filter := map[string]interface{}{
		"last":   3,
		"after":  "YXJyYXljb25uZWN0aW9uOjA=",
		"before": "YXJyYXljb25uZWN0aW9uOjQ=",
	}
	args := graphql_relay.NewConnectionArguments(filter)

	expected := &graphql_relay.Connection{
		Edges: []*graphql_relay.Edge{
			&graphql_relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&graphql_relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&graphql_relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
		},
		PageInfo: graphql_relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjE=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjM=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := graphql_relay.ConnectionFromArray(arrayConnectionTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, connection result diff: %v", testutil.Diff(expected, result))
	}
}

func TestConnectionFromArray_HandlesCursorEdgeCases_ReturnsNoElementsIfFirstIsZero(t *testing.T) {
	filter := map[string]interface{}{
		"first": 0,
	}
	args := graphql_relay.NewConnectionArguments(filter)

	expected := &graphql_relay.Connection{
		Edges:    []*graphql_relay.Edge{},
		PageInfo: graphql_relay.PageInfo{},
	}

	result := graphql_relay.ConnectionFromArray(arrayConnectionTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, connection result diff: %v", testutil.Diff(expected, result))
	}
}
func TestConnectionFromArray_HandlesCursorEdgeCases_ReturnsAllElementsIfCursorsAreInvalid(t *testing.T) {
	filter := map[string]interface{}{
		"before": "invalid",
		"after":  "invalid",
	}
	args := graphql_relay.NewConnectionArguments(filter)

	expected := &graphql_relay.Connection{
		Edges: []*graphql_relay.Edge{
			&graphql_relay.Edge{
				Node:   "A",
				Cursor: "YXJyYXljb25uZWN0aW9uOjA=",
			},
			&graphql_relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&graphql_relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&graphql_relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
			&graphql_relay.Edge{
				Node:   "E",
				Cursor: "YXJyYXljb25uZWN0aW9uOjQ=",
			},
		},
		PageInfo: graphql_relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjA=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjQ=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := graphql_relay.ConnectionFromArray(arrayConnectionTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, connection result diff: %v", testutil.Diff(expected, result))
	}
}
func TestConnectionFromArray_HandlesCursorEdgeCases_ReturnsAllElementsIfCursorsAreOnTheOutside(t *testing.T) {
	filter := map[string]interface{}{
		"before": "YXJyYXljb25uZWN0aW9uOjYK",     // ==> offset: int(6)
		"after":  "YXJyYXljb25uZWN0aW9uOi0xCg==", // ==> offset: int(-1)
	}
	args := graphql_relay.NewConnectionArguments(filter)

	expected := &graphql_relay.Connection{
		Edges: []*graphql_relay.Edge{
			&graphql_relay.Edge{
				Node:   "A",
				Cursor: "YXJyYXljb25uZWN0aW9uOjA=",
			},
			&graphql_relay.Edge{
				Node:   "B",
				Cursor: "YXJyYXljb25uZWN0aW9uOjE=",
			},
			&graphql_relay.Edge{
				Node:   "C",
				Cursor: "YXJyYXljb25uZWN0aW9uOjI=",
			},
			&graphql_relay.Edge{
				Node:   "D",
				Cursor: "YXJyYXljb25uZWN0aW9uOjM=",
			},
			&graphql_relay.Edge{
				Node:   "E",
				Cursor: "YXJyYXljb25uZWN0aW9uOjQ=",
			},
		},
		PageInfo: graphql_relay.PageInfo{
			StartCursor:     "YXJyYXljb25uZWN0aW9uOjA=",
			EndCursor:       "YXJyYXljb25uZWN0aW9uOjQ=",
			HasPreviousPage: false,
			HasNextPage:     false,
		},
	}

	result := graphql_relay.ConnectionFromArray(arrayConnectionTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, connection result diff: %v", testutil.Diff(expected, result))
	}
}

func TestConnectionFromArray_HandlesCursorEdgeCases_ReturnsNullIfCursorsIsConsecutive(t *testing.T) {
	filter := map[string]interface{}{
		"before": "YXJyYXljb25uZWN0aW9uOjM=", // ==> offset: int(3)
		"after":  "YXJyYXljb25uZWN0aW9uOjI=", // ==> offset: int(2)
	}
	args := graphql_relay.NewConnectionArguments(filter)

	expected := &graphql_relay.Connection{
		Edges:    []*graphql_relay.Edge{},
		PageInfo: graphql_relay.PageInfo{},
	}

	result := graphql_relay.ConnectionFromArray(arrayConnectionTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, connection result diff: %v", testutil.Diff(expected, result))
	}
}
func TestConnectionFromArray_HandlesCursorEdgeCases_ReturnsNoElementsIfCursorsCross(t *testing.T) {
	filter := map[string]interface{}{
		"before": "YXJyYXljb25uZWN0aW9uOjI=", // ==> offset: int(2)
		"after":  "YXJyYXljb25uZWN0aW9uOjQ=", // ==> offset: int(4)
	}
	args := graphql_relay.NewConnectionArguments(filter)

	expected := &graphql_relay.Connection{
		Edges:    []*graphql_relay.Edge{},
		PageInfo: graphql_relay.PageInfo{},
	}

	result := graphql_relay.ConnectionFromArray(arrayConnectionTestLetters, args)
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("wrong result, connection result diff: %v", testutil.Diff(expected, result))
	}
}
func TestConnectionFromArray_CursorForObjectInConnection_ReturnsAnEdgeCursor_GivenAnArrayAndAMemberObject(t *testing.T) {
	letterBCursor := graphql_relay.CursorForObjectInConnection(arrayConnectionTestLetters, "B")
	expected := graphql_relay.ConnectionCursor("YXJyYXljb25uZWN0aW9uOjE=")
	if !reflect.DeepEqual(letterBCursor, expected) {
		t.Fatalf("wrong result, cursor result diff: %v", testutil.Diff(expected, letterBCursor))
	}
}
func TestConnectionFromArray_CursorForObjectInConnection_ReturnsEmptyCursor_GivenAnArrayAndANonMemberObject(t *testing.T) {
	letterFCursor := graphql_relay.CursorForObjectInConnection(arrayConnectionTestLetters, "F")
	if letterFCursor != "" {
		t.Fatalf("wrong result, expected empty cursor, got: %v", letterFCursor)
	}
}