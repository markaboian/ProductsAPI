package product

import "Clase 11 - new project structure/internal/domain"

type Repository interface {
	GetAllProducts() []domain.Producto
}

