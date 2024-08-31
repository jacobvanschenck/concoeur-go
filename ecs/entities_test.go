package ecs

import (
	"testing"
)

type Health struct {
	health int
}

const HealthTypeId = "healthTypeId"

func (health Health) typeId() TypeId {
	return HealthTypeId
}

type TestPosition struct {
	x int
	y int
}

const TestPositionTypeId = "testPositionTypeId"

func (testPosition TestPosition) typeId() TypeId {
	return TestPositionTypeId
}

type AnotherComponent struct {
	x int
}

const AnotherComponentTypeId = "anotherComponentTypeId"

func (another AnotherComponent) typeId() string {
	return AnotherComponentTypeId
}

func initEntities() Entities {
	return Entities{
		components:  make(map[TypeId][]Component),
		bitMasks:    make(map[TypeId]int32),
		insertIndex: 0,
		entityMap:   []int32{},
	}
}

func TestRegisterComponent(t *testing.T) {
	entities := initEntities()

	// Call the function we want to test
	entities.registerComponent(HealthTypeId)

	// Check that the component was registered correctly
	if len(entities.components) != 1 {
		t.Fatalf("expected 1 component, got %d", len(entities.components))
	}

	if components, exists := entities.components[HealthTypeId]; !exists ||
		len(components) != 0 {
		t.Fatalf("expected component to be registered for %s", HealthTypeId)
	}

	expectedBitMask := int32(0b00000001)
	if bitMask, exists := entities.bitMasks[HealthTypeId]; !exists ||
		bitMask != int32(expectedBitMask) {
		t.Fatalf(
			"expected bitmask to be %d for %s, got %d",
			expectedBitMask,
			HealthTypeId,
			bitMask,
		)
	}

	entities.registerComponent(TestPositionTypeId)

	expectedBitMask = int32(0b00000010)
	if bitMask, exists := entities.bitMasks[TestPositionTypeId]; !exists ||
		bitMask != expectedBitMask {
		t.Fatalf(
			"expected bitmask to be %d for %s, got %d",
			expectedBitMask,
			TestPositionTypeId,
			bitMask,
		)
	}

	entities.registerComponent(AnotherComponentTypeId)

	expectedBitMask = int32(0b00000100)
	if bitMask, exists := entities.bitMasks[AnotherComponentTypeId]; !exists ||
		bitMask != expectedBitMask {
		t.Fatalf(
			"expected bitmask to be %d for %s, got %d",
			expectedBitMask,
			AnotherComponentTypeId,
			bitMask,
		)
	}
}

func TestCreateEntity(t *testing.T) {
	entities := initEntities()
	if len(entities.entityMap) != 0 {
		t.Fatalf("expected entityMap to be empty")
	}
	entities.createEntity()
	if len(entities.entityMap) != 1 {
		t.Fatalf("expected entityMap to be have a length of 1")
	}
	entities.createEntity()
	if len(entities.entityMap) != 1 {
		t.Fatalf("expected entityMap to be have a length of 1 since no components added")
	}
}

func TestCreateEntityWithComponent(t *testing.T) {
	entities := initEntities()
	if len(entities.entityMap) != 0 {
		t.Fatalf("expected entityMap to be have a length of 1")
	}

	entities.registerComponent(HealthTypeId)
	entities.registerComponent(TestPositionTypeId)

	entities.createEntity().withComponent(Health{health: 100})

	if len(entities.entityMap) != 1 {
		t.Fatalf("expected entityMap to be have a length of 1")
	}

	entities.createEntity().
		withComponent(Health{health: 100}).
		withComponent(TestPosition{x: 1, y: 1})

	entities.createEntity().
		withComponent(TestPosition{x: 1, y: 1})
	if len(entities.entityMap) != 3 {
		t.Fatalf("expected entityMap to be have a length of 1 since no components added")
	}

	expectedBitMask := int32(0b00000001)
	if bitMask := entities.entityMap[0]; bitMask != expectedBitMask {
		t.Fatalf(
			"expected bitmask to be %d for first entity, got %d",
			expectedBitMask,
			bitMask,
		)
	}

	expectedBitMask = int32(0b00000011)
	if bitMask := entities.entityMap[1]; bitMask != expectedBitMask {
		t.Fatalf(
			"expected bitmask to be %d for second entity, got %d",
			expectedBitMask,
			bitMask,
		)
	}

	expectedBitMask = int32(0b00000010)
	if bitMask := entities.entityMap[2]; bitMask != expectedBitMask {
		t.Fatalf(
			"expected bitmask to be %d for third entity, got %d",
			expectedBitMask,
			bitMask,
		)
	}
}
