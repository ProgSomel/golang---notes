package database

type Product struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	ImgURL string `json:"imageURL"`
}
var productList []Product

func Store(p Product) Product{
	p.ID = len(productList)+1
	productList = append(productList, p)

	return p
}

func List() []Product{
	return productList
}

func Get(productID int) *Product{
	for _, product := range productList{
		if product.ID == productID{
			return &product
		}
	}
	return nil
}

func Update(product Product){
	for idx, p := range productList{
		if p.ID == product.ID{
			productList[idx] = product
		}
	}
}

func Delete(productID int){
	var tempList []Product

	for _, p := range productList{
		if p.ID != productID{
			tempList = append(tempList, p)
		}
	}

	productList = tempList
}

func init(){
	prdct1 := Product{
		ID: 1,
		Title: "Orange",
		Description: "Orange is Red. I Love Orange",
		Price: 100,
		ImgURL: "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
	}
	prdct2 := Product{
		ID: 2,
		Title: "Apple",
		Description: "Apple is Green. I eat Apple",
		Price: 40,
		ImgURL: "https://static.vecteezy.com/system/resources/thumbnails/012/086/172/small_2x/green-apple-with-green-leaf-isolated-on-white-background-vector.jpg",
	}
	prdct3 := Product{
		ID: 3,
		Title: "Banana",
		Description: "Banana is Boring. I hate Banana",
		Price: 5,
		ImgURL: "https://www.dole.com/sites/default/files/media/2025-01/banana-cavendish_0.png",
	}
	productList = append(productList, prdct1)
	productList = append(productList, prdct2)
	productList = append(productList, prdct3)

}