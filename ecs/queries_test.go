package ecs

import (
	"testing"
)

func TestQueryWithComponent(t *testing.T) {
	entities := initEntities()
	entities.registerComponent(HealthTypeId)
	entities.registerComponent(TestPositionTypeId)

	entities.createEntity().
		withComponent(Health{health: 100}).
		withComponent(TestPosition{x: 1, y: 1})

	query := Query{entities: &entities}
	query.withComponent(HealthTypeId).withComponent(TestPositionTypeId)

	if len(query.typeIds) != 2 {
		t.Fatalf("expected query.typeIds to be have a length of 1")
	}
	if query.queryMap != int32(0b00000011) {
		t.Fatalf("expected queryMap to be 3")
	}
}

func TestQueryRun(t *testing.T) {
	entities := initEntities()
	entities.registerComponent(HealthTypeId)
	entities.registerComponent(TestPositionTypeId)

	entities.createEntity().
		withComponent(Health{health: 100}).
		withComponent(TestPosition{x: 1, y: 1})

	entities.createEntity().
		withComponent(Health{health: 90})

	entities.createEntity().
		withComponent(TestPosition{x: 2, y: 2})

	entities.createEntity().
		withComponent(Health{health: 120}).
		withComponent(TestPosition{x: 3, y: 3})

	query := Query{entities: &entities}
	query.withComponent(HealthTypeId).withComponent(TestPositionTypeId)

	indexes, components := query.run()

	if indexes[0] != 0 || indexes[1] != 3 {
		t.Fatalf("Wrong indexes queried, expected [0,3] got[%d,%d]", indexes[0], indexes[1])
	}

	firstComponents := components[0]
	if firstComponents[0].typeId() != HealthTypeId {
		t.Fatalf(
			"Wrong typeId for first component, expected: %s got: %s",
			HealthTypeId,
			firstComponents[0].typeId(),
		)
	}

	if firstComponents[0].(Health).health != 100 {
		t.Fatalf(
			"Expected health: 100, got:  %d",
			firstComponents[0].(Health).health,
		)
	}

	if firstComponents[1].(Health).health != 120 {
		t.Fatalf(
			"Expected health: 100, got:  %d",
			firstComponents[1].(Health).health,
		)
	}

	secondComponents := components[1]
	if secondComponents[0].typeId() != TestPositionTypeId {
		t.Fatalf(
			"Wrong typeId for first component, expected: %s got: %s",
			TestPositionTypeId,
			secondComponents[0].typeId(),
		)
	}

	firstTestPostition := secondComponents[0].(TestPosition)
	if firstTestPostition.x != 1 || firstTestPostition.y != 1 {
		t.Fatalf(
			"Expected position: (1,1), got:  (%d,%d)",
			firstTestPostition.x,
			firstTestPostition.y,
		)
	}

	secondTestPosition := secondComponents[1].(TestPosition)
	if secondTestPosition.x != 3 || secondTestPosition.y != 3 {
		t.Fatalf(
			"Expected position: (3,3), got:  (%d,%d)",
			secondTestPosition.x,
			secondTestPosition.y,
		)
	}
}
