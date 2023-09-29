// Copyright Letga Author(https://letga.net). All Rights Reserved.
// Apache License 2.0(https://github.com/lgcgo/letga-server/blob/main/LICENSE)

package tree

import (
	"errors"
	"reflect"
	"sort"
	"strings"
)

// Tree hanle
type Tree struct {
	NodeList map[string]*node // Node list
	NodeLink map[*node]*node  // Node relationship
}

// TreeNode structure
type TreeNode struct {
	Key       string      `json:"key"`
	ParentKey string      `json:"parentKey"`
	Title     string      `json:"title"`
	Weight    int         `json:"weight"`
	Source    interface{} `json:"source"`
	Children  []*TreeNode `json:"children"`
}

// Implement sort Inoder abstract component
func (t TreeNode) CalSortValue() int {
	return t.Weight
}

// Init tree with datas
func NewWithData(params interface{}) (*Tree, error) {
	var (
		nodeList    = make(map[string]*node, 0) // Node list
		paramsValue = reflect.ValueOf(params)
	)
	// Set node list
	for i := 0; i < paramsValue.Len(); i++ {
		//temp := &TreeNode{}
		itemElem := paramsValue.Index(i).Elem()
		itemKey := itemElem.FieldByName("Key")
		if !itemKey.IsValid() {
			return nil, errors.New("key field not exist")
		}
		itemParentKey := itemElem.FieldByName("ParentKey")
		if !itemParentKey.IsValid() {
			return nil, errors.New("parent key field not exist")
		}
		itemTitle := itemElem.FieldByName("Title")
		if !itemTitle.IsValid() {
			return nil, errors.New("title field not exist")
		}
		itemWeight := itemElem.FieldByName("Weight")
		if !itemWeight.IsValid() {
			return nil, errors.New("weight field not exist")
		}
		treeNode := &TreeNode{
			Key:       itemKey.String(),
			ParentKey: itemParentKey.String(),
			Title:     itemTitle.String(),
			Weight:    int(itemWeight.Int()),
			Source:    itemElem.Interface(),
			Children:  make([]*TreeNode, 0),
		}
		nodeList[itemKey.String()] = NewNode(treeNode)
	}
	return newTree(nodeList), nil
}

// Initial tree struct
func newTree(nodeList map[string]*node) *Tree {
	var nodeLink = make(map[*node]*node, 0) // Node relationship
	// Set node relationship list
	for _, v := range nodeList {
		parentKey := v.Noder.(*TreeNode).ParentKey
		if _, ok := nodeList[parentKey]; ok {
			nodeLink[v] = nodeList[parentKey]
		} else {
			nodeLink[v] = nil
		}
	}
	// Set node childs
	for _, n := range nodeList {
		var (
			nodes     = make(nodes, 0)
			treeNodes = make([]*TreeNode, 0)
		)
		// Get child nodes
		for k, v := range nodeLink {
			if n == v {
				nodes = append(nodes, k)
			}
		}
		sort.Sort(nodes)
		// data conversion
		for _, v := range nodes {
			treeNodes = append(treeNodes, v.Noder.(*TreeNode))
		}
		n.Noder.(*TreeNode).Children = treeNodes
	}
	return &Tree{nodeList, nodeLink}
}

// Get the parent Key set, excluding itself
func (t *Tree) GetParentKeys(key string) ([]string, error) {
	var (
		keys     = make([]string, 0)
		treeNode *TreeNode
		err      error
	)
	if treeNode, err = t.getTreeNode(key); err != nil {
		return nil, err
	}
	tempNode := treeNode
	for {
		if len(tempNode.ParentKey) == 0 {
			break
		}
		tempNode = t.NodeList[tempNode.ParentKey].Noder.(*TreeNode)
		keys = append(keys, tempNode.Key)
	}
	return keys, err
}

// Get the children Key set, excluding itself
func (t *Tree) GetChildKeys(key string) ([]string, error) {
	var (
		keys      = make([]string, 0)
		childrens = make(map[string]*TreeNode, 0)
		treeNode  *TreeNode
		err       error
	)
	if treeNode, err = t.getTreeNode(key); err != nil {
		return nil, err
	}
	// Initialize queue
	for _, v := range treeNode.Children {
		childrens[v.Key] = v
	}
	// level-traversal
	for {
		if len(childrens) == 0 {
			break
		}
		temp := childrens
		for k, v := range temp {
			keys = append(keys, v.Key)
			delete(childrens, k)
			if len(v.Children) > 0 {
				for _, vv := range v.Children {
					childrens[vv.Key] = vv
				}
			}
		}
	}
	return keys, nil
}

// Get the specified node
func (t *Tree) getTreeNode(key string) (*TreeNode, error) {
	n, ok := t.NodeList[key]
	if !ok {
		return nil, errors.New("node not exist")
	}
	return n.Noder.(*TreeNode), nil
}

// Get whole tree node
func (t *Tree) GetWholeTree() []*TreeNode {
	var (
		nodes = make(nodes, 0)
		res   []*TreeNode
	)
	for k, v := range t.NodeLink {
		if v == nil {
			nodes = append(nodes, k)
		}
	}
	sort.Sort(nodes)
	for _, v := range nodes {
		res = append(res, v.Noder.(*TreeNode))
	}
	return res
}

// Search
func (t *Tree) Search(search string, keys []string) ([]*TreeNode, error) {
	var (
		handler  *Tree
		nodeList = make(map[string]*node, 0) // Node list
		newKeys  []string
	)
	if len(keys) > 0 {
		for _, v := range keys {
			temKeys, _ := t.GetParentKeys(v)
			temKeys = append(temKeys, v)
			newKeys = append(newKeys, temKeys...)
		}
		newKeys = keysDeduplication(newKeys)
		for _, v := range newKeys {
			nodeList[v] = t.NodeList[v]
		}
		handler = newTree(nodeList)
	} else {
		handler = t
	}
	if len(search) > 0 {
		var (
			newNodeList = make(map[string]*node, 0) // Node list
			searchKeys  []string
		)
		for k, v := range handler.NodeList {
			if strings.Contains(v.Noder.(*TreeNode).Title, search) {
				searchKeys = append(searchKeys, k)
			}
		}
		for _, v := range searchKeys {
			searchParentKeys, _ := handler.GetParentKeys(v)
			searchKeys = append(searchKeys, searchParentKeys...)
		}
		searchKeys = keysDeduplication(searchKeys)
		for _, key := range searchKeys {
			newNodeList[key] = handler.NodeList[key]
		}
		handler = newTree(newNodeList)
	}
	return handler.GetWholeTree(), nil
}

// Keys deduplication
func keysDeduplication(keys []string) []string {
	set := make(map[string]struct{}, len(keys))
	j := 0
	for _, v := range keys {
		_, ok := set[v]
		if ok {
			continue
		}
		set[v] = struct{}{}
		keys[j] = v
		j++
	}
	return keys[:j]
}
