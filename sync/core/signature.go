package core

import (
	"dbsync/sync/contract/strategy/diff"
	"fmt"
	"github.com/viant/toolbox"
)

//Signature represents dataset signature
type Signature struct {
	id                        string
	CountValue                interface{}
	MaxValue                  interface{}
	MinValue                  interface{}
	UniqueIDCountValue        interface{}
	NotNullUniqueIDCountValue interface{}
}

//Min returns min value
func (c *Signature) Min() int {
	if c.MinValue == nil {
		return 0
	}
	return toolbox.AsInt(c.MinValue)
}

//Max returns max value
func (c *Signature) Max() int {
	if c.MaxValue == nil {
		return 0
	}
	return toolbox.AsInt(c.MaxValue)
}

//Count returns count value
func (c *Signature) Count() int {
	if c == nil || c.CountValue == nil {
		return 0
	}
	return toolbox.AsInt(c.CountValue)
}

//UniqueIDCount returns unique count value
func (c *Signature) IsUniqueIDDefined() bool {
	if c.UniqueIDCountValue == nil {
		return false
	}
	return true
}

//UniqueIDCount returns unique count value
func (c *Signature) UniqueIDCount() int {
	if c.UniqueIDCountValue == nil {
		return 0
	}
	return toolbox.AsInt(c.UniqueIDCountValue)
}

//NotNullUniqueIDCount returns non unique count value
func (c *Signature) NotNullUniqueIDCount() int {
	if c.NotNullUniqueIDCountValue == nil {
		return 0
	}
	return toolbox.AsInt(c.NotNullUniqueIDCountValue)
}

//IsEqual returns true if signatures are equal
func (c *Signature) IsEqual(signature *Signature) bool {
	if signature == nil {
		return false
	}
	return c.Count() == signature.Count() &&
		c.Max() == signature.Max() &&
		c.Min() == signature.Min()
}

//ValidateIDConsistency checks if dist
func (c *Signature) ValidateIDConsistency() error {
	if c.id == "" {
		return nil
	}

	if !c.IsUniqueIDDefined() {
		return nil
	}
	if c.Count() == c.UniqueIDCount() {
		return nil
	}
	if c.Count() != c.NotNullUniqueIDCount() {
		return fmt.Errorf("unique column has NULL Source, rowCount: %v, unique ID count: %v ", c.Count(), c.NotNullUniqueIDCount())
	}

	return fmt.Errorf(" data has unique ID duplicates, rowCount: %v, unique ID count: %v ", c.Count(), c.UniqueIDCount())
}

//NewSignatureFromRecord creates a new signature from a record
func NewSignatureFromRecord(idKey string, record Record) *Signature {
	if idKey == "" {
		return &Signature{
			id:         idKey,
			CountValue: record.Get(diff.AliasCount),
		}
	}
	return &Signature{
		id:                        idKey,
		CountValue:                record.Get(diff.AliasCount),
		MinValue:                  record.Get(fmt.Sprintf(diff.AliasMinIdTemplate, idKey)),
		MaxValue:                  record.Get(fmt.Sprintf(diff.AliasMaxIdTemplate, idKey)),
		UniqueIDCountValue:        record.Get(fmt.Sprintf(diff.AliasUniqueIDCountTemplate, idKey)),
		NotNullUniqueIDCountValue: record.Get(fmt.Sprintf(diff.AliasNonNullIDCountTemplate, idKey)),
	}
}
