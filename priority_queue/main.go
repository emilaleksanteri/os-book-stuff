package main

import (
	"fmt"
)

type QueueNode struct {
	jobLen  int
	jobInfo string
	Next    *QueueNode
}

func (n *QueueNode) GetJobLen() int {
	return n.jobLen
}

func (n *QueueNode) GetJobInfo() string {
	return n.jobInfo
}

type PQueue struct {
	head *QueueNode
	size int
}

func (pq *PQueue) Size() int {
	return pq.size
}

func (pq *PQueue) Peak() *QueueNode {
	return pq.head
}

func (pq *PQueue) Insert(jobLen int, jobInfo string) {
	// we need to reorder queue to have highest priority first
	if pq.Size() == 0 {
		pq.head = &QueueNode{
			jobLen:  jobLen,
			jobInfo: jobInfo,
		}

		pq.size += 1
		return
	}

	if pq.Size() == 1 {
		head := pq.head
		insert := &QueueNode{
			jobLen:  jobLen,
			jobInfo: jobInfo,
		}
		if head.GetJobLen() <= jobLen {
			insert.Next = head
			pq.head = insert
		} else {
			pq.head.Next = insert
		}

		pq.size += 1
		return
	}

	toInsert := &QueueNode{
		jobLen:  jobLen,
		jobInfo: jobInfo,
	}
	head := pq.head
	pq.sortInsert(toInsert, head, head, true)
	pq.size += 1
}

func (pq *PQueue) sortInsert(inserted *QueueNode, prevNode *QueueNode, currNode *QueueNode, currIsHead bool) {
	if currIsHead && inserted.jobLen >= currNode.jobLen {
		pq.head = inserted
		inserted.Next = currNode
		return
	}

	if inserted.jobLen >= currNode.jobLen {
		prevNode.Next = inserted
		inserted.Next = currNode
		return
	}

	if currNode.Next == nil {
		currNode.Next = inserted
		return
	}

	next := currNode.Next
	pq.sortInsert(inserted, currNode, next, false)
}

func (pq *PQueue) Get() *QueueNode {
	if pq.Size() == 0 {
		return nil
	}

	if pq.Size() == 1 {
		pq.size -= 1
		return pq.head
	}

	pq.size -= 1
	hd := pq.head
	pq.head = hd.Next
	return hd
}

func (pq *PQueue) IsEmpty() bool {
	if pq.Size() == 0 {
		return true
	}

	return false
}

func main() {
	toInsert := []QueueNode{
		{jobLen: 10, jobInfo: "foo"},
		{jobLen: 9, jobInfo: "foo"},
		{jobLen: 7, jobInfo: "foo"},
		{jobLen: 6, jobInfo: "foo"},
		{jobLen: 5, jobInfo: "foo"},
	}

	testVal := QueueNode{
		jobLen:  8,
		jobInfo: "baz",
	}

	pq := PQueue{}
	for _, insert := range toInsert {
		pq.Insert(insert.jobLen, insert.jobInfo)
	}

	if pq.Size() != len(toInsert) {
		fmt.Printf("got %d, wanted: %d\n", pq.Size(), len(toInsert))
		return
	}

	pq.Insert(testVal.jobLen, testVal.jobInfo)
	if pq.Size() != len(toInsert)+1 {
		fmt.Printf("got %d, wanted: %d\n", pq.Size(), len(toInsert)+1)
	}

	all := []QueueNode{}
	for !pq.IsEmpty() {
		all = append(all, *pq.Get())
	}

	want := []QueueNode{
		{jobLen: 10, jobInfo: "foo"},
		{jobLen: 9, jobInfo: "foo"},
		{jobLen: 8, jobInfo: "baz"},
		{jobLen: 7, jobInfo: "foo"},
		{jobLen: 6, jobInfo: "foo"},
		{jobLen: 5, jobInfo: "foo"},
	}

	for i := range all {
		if all[i].jobLen != want[i].jobLen && all[i].jobInfo != want[i].jobInfo {
			fmt.Printf("oopsie queue not in order, got %+v and wanted %+v\n", all[i], want[i])
			break
		}
	}

}
