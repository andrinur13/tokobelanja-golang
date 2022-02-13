package input

type CreateProduct struct {
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required"`
	Stock int    `json:"stock" binding:"required"`
}

type UpdateProduct struct {
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required"`
	Stock int    `json:"stock" binding:"required"`
}

type DeleteProduct struct {
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required"`
	Stock int    `json:"stock" binding:"required"`
}

type InputProduct struct {
	ID int `json:"id_product"`
}
