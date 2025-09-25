package repo

type Product struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	ImgURL string `json:"imageURL"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(product Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

func (r *productRepo)  Create(p Product) (*Product, error) {
	
	p.ID = len(r.productList)+1
	r.productList = append(r.productList, &p)

	return &p, nil
}

func (r *productRepo) Get(productID int) (*Product, error) {

	for _, product := range r.productList{
		if product.ID == productID{
			return product, nil
		}
	}
	return nil, nil
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func (r *productRepo) Delete(productID int) error {
	var tempList []*Product

	for _, p := range r.productList{
		if p.ID != productID{
			tempList = append(tempList, p)
		}
	}

	r.productList = tempList
	return nil
}

func (r *productRepo) Update(product Product) (*Product, error) {

	for idx, p := range r.productList{
		if p.ID == product.ID{
			r.productList[idx] = &product
		}
	}
	return &product, nil
}

//? constructor or constructor function
func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProducts(repo)
	return repo
}

func generateInitialProducts(r *productRepo){
	prdct1 := &Product{
		ID: 1,
		Title: "Orange",
		Description: "Orange is Red. I Love Orange",
		Price: 100,
		ImgURL: "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
	}
	prdct2 := &Product{
		ID: 2,
		Title: "Apple",
		Description: "Apple is Green. I eat Apple",
		Price: 40,
		ImgURL: "https://static.vecteezy.com/system/resources/thumbnails/012/086/172/small_2x/green-apple-with-green-leaf-isolated-on-white-background-vector.jpg",
	}
	prdct3 := &Product{
		ID: 3,
		Title: "Banana",
		Description: "Banana is Boring. I hate Banana",
		Price: 5,
		ImgURL: "https://www.dole.com/sites/default/files/media/2025-01/banana-cavendish_0.png",
	}
	r.productList = append(r.productList, prdct1)
	r.productList = append(r.productList, prdct2)
	r.productList = append(r.productList, prdct3)

}