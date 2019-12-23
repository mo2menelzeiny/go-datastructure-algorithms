package stacksandqueues

import "container/list"

const (
	ANIMAL_DOG = "DOG"
	ANIMAL_CAT = "CAT"
)

type Animal struct {
	AnimalType string
	Name       string
	order      int
}

type AnimalShelter struct {
	dogsQueue list.List
	catsQueue list.List
}

func (as *AnimalShelter) Size() int {
	return as.catsQueue.Len() + as.dogsQueue.Len()
}

func (as *AnimalShelter) Enqueue(a Animal) {
	a.order = as.Size() + 1
	switch a.AnimalType {
	case ANIMAL_CAT:
		as.catsQueue.PushBack(a)
	case ANIMAL_DOG:
		as.dogsQueue.PushBack(a)
	}
}

func (as *AnimalShelter) DequeueAny() Animal {
	cat := as.catsQueue.Front().Value.(Animal)
	dog := as.dogsQueue.Front().Value.(Animal)
	if cat.order > dog.order {
		return as.catsQueue.Remove(as.catsQueue.Front()).(Animal)
	} else {
		return as.dogsQueue.Remove(as.dogsQueue.Front()).(Animal)
	}
}

func (as *AnimalShelter) DequeueAnimal(animal string) Animal {
	switch animal {
	case ANIMAL_CAT:
		return as.catsQueue.Remove(as.catsQueue.Front()).(Animal)
	default:
		return as.dogsQueue.Remove(as.dogsQueue.Front()).(Animal)
	}
}
