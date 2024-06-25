package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("product 1", 10.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "product 1", p.Name)
	assert.Equal(t, 10.0, p.Price)

}

func TestProductWhenNameRequired(t *testing.T) {
	p, err := NewProduct("product 1", 10.0)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
