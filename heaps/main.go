package main

type MaxHeap struct {
  array [] int
}

func (h *MaxHeap) Insert(key int) {
  h.array = append(h.array, key)
  h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) maxHeapifyUp(index int) {

}

func main ()  {
  
}

