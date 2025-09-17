package main

import (
	"sync"
	"time"
)

type Node struct {
	data map[string]string
	mu   sync.RWMutex
}

type EventualConsistentStore struct {
	nodes map[string]*Node
}

func NewEventualConsistentStore() *EventualConsistentStore {
	return &EventualConsistentStore{
		nodes: map[string]*Node{
			"n1": {data: make(map[string]string)},
			"n2": {data: make(map[string]string)},
			"n3": {data: make(map[string]string)},
		},
	}
}

func (ecs *EventualConsistentStore) Write(key, val string) {
	primaryNode := ecs.nodes["n1"]
	primaryNode.mu.Lock()
	primaryNode.data[key] = val
	primaryNode.mu.Unlock()

	go ecs.replicateToOtherNodes(key, val)
}

func (ecs *EventualConsistentStore) replicateToOtherNodes(key, val string) {
	for nodeName, node := range ecs.nodes {
		time.Sleep(2 * time.Second)
		if nodeName == "n1" {
			continue
		}

		node.mu.Lock()
		node.data[key] = val
		node.mu.Unlock()
	}
}

func (ecs *EventualConsistentStore) Read(nodeID string) map[string]string {
	if node, exists := ecs.nodes[nodeID]; exists {
		node.mu.RLock()
		defer node.mu.RUnlock()

		result := make(map[string]string)
		for k, v := range node.data {
			result[k] = v
		}
		return result
	}
	return nil
}
