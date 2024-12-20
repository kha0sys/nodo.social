package dto

import (
	"time"
	"github.com/kha0sys/nodo.social/functions/domain/models"
	"github.com/kha0sys/nodo.social/functions/domain/models/contact"
)

// ProductDTO representa los datos de un producto para transferencia
type ProductDTO struct {
	ID              string             `json:"id"`
	Name            string             `json:"name"`
	Description     string             `json:"description"`
	Price           float64            `json:"price"`
	StoreID         string             `json:"storeId"`
	NodeID          string             `json:"nodeId"`
	Images          []string           `json:"images"`
	Contact         contact.ContactInfo `json:"contact"`
	DonationPercent int               `json:"donationPercent"`
	ApprovalStatus  string            `json:"approvalStatus"`
	CreatedAt       time.Time         `json:"createdAt"`
	UpdatedAt       time.Time         `json:"updatedAt"`
}

// ToModel convierte el DTO a un modelo Product
func (dto *ProductDTO) ToModel() *models.Product {
	return &models.Product{
		ID:              dto.ID,
		Name:            dto.Name,
		Description:     dto.Description,
		Price:          dto.Price,
		StoreID:        dto.StoreID,
		NodeID:         dto.NodeID,
		Images:         dto.Images,
		Contact:        dto.Contact,
		DonationPercent: dto.DonationPercent,
		ApprovalStatus: dto.ApprovalStatus,
		CreatedAt:      dto.CreatedAt,
		UpdatedAt:      dto.UpdatedAt,
	}
}

// FromProductModel crea un DTO a partir de un modelo Product
func FromProductModel(product *models.Product) *ProductDTO {
	return &ProductDTO{
		ID:              product.ID,
		Name:            product.Name,
		Description:     product.Description,
		Price:          product.Price,
		StoreID:        product.StoreID,
		NodeID:         product.NodeID,
		Images:         product.Images,
		Contact:        product.Contact,
		DonationPercent: product.DonationPercent,
		ApprovalStatus: product.ApprovalStatus,
		CreatedAt:      product.CreatedAt,
		UpdatedAt:      product.UpdatedAt,
	}
}
