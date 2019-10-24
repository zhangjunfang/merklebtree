package merklebtree_go

import (
	"crypto/md5"
	"encoding/hex"
)

type MerkleBtree struct {
	hashmap map[string]int64
	Root
}

type Root struct {
	Hash string
}

func (mbtree *MerkleBtree) ComputeHash(){
	var s string
	for key, _ := range mbtree.hashmap {
		signByte := []byte(key)
		hash := md5.New()
		hash.Write(signByte)
		s = s + hex.EncodeToString(hash.Sum(nil))
	}
	signByte := []byte(s)
	hash := md5.New()
	hash.Write(signByte)
	mbtree.Hash = hex.EncodeToString(hash.Sum(nil))
}

func (mbtree *MerkleBtree) BuildWithKeyValue(kv KeyVersion) {
	mbtree.hashmap[kv.Key] = kv.Version
	mbtree.ComputeHash()
}

func (mbtree *MerkleBtree) Delete(key string) {
	delete(mbtree.hashmap, key)
	mbtree.ComputeHash()
}

func (mbtree *MerkleBtree) Serach(key string) SearchResult {
	version, existed := mbtree.hashmap[key]
	return SearchResult{Key: key, Version: version, Existed: existed}
}

type SearchResult struct {
	Key     string
	Version int64
	Existed bool
}

type KeyVersion struct {
	Key     string
	Version int64
}

func NewMBTree() *MerkleBtree {
	btree := MerkleBtree{hashmap: make(map[string]int64), Root: Root{Hash: ""}}
	return &btree
}
