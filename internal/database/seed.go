package database

import (
	"fmt"
	"log"
	"lotery_viking/internal/models"
)

var Tables = []string{
	"users",
	"kiosks",
	"publicity_images",
	"parameters",
	"images",
	"rewards",
	"tickets",
}

func Seed() {
	dbService := New()
	db := dbService.GetDB()
	defer db.Close()

	for _, table := range Tables {
		_, err := db.Exec(fmt.Sprintf("DELETE FROM %s", table.name))
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Tables cleared")

	dataImages := []models.Images{
		{BaseModel: models.BaseModel{ID: 1}, Name: "Image 1", Url: "https://www.google.com", Format: "jpg"},
		{BaseModel: models.BaseModel{ID: 2}, Name: "Image 2", Url: "https://www.google.com", Format: "png"},
		{BaseModel: models.BaseModel{ID: 3}, Name: "Image 3", Url: "https://www.google.com", Format: "jpg"},
		{BaseModel: models.BaseModel{ID: 4}, Name: "Image 4 pub", Url: "https://www.google.com", Format: "png"},
		{BaseModel: models.BaseModel{ID: 5}, Name: "Image 5 pub", Url: "https://www.google.com", Format: "jpg"},
	}

	stmt, err := db.Prepare("INSERT INTO images (id, name, url, format) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, image := range dataImages {
		_, err = stmt.Exec(image.ID, image.Name, image.Url, image.Format)
		if err != nil {
			log.Fatal(err)
		}
	}

	dataParameters := []models.Parameters{
		{BaseModel: models.BaseModel{ID: 1},
			NameLotery:    "Lottery 1",
			NameCasino:    "Casino de Sanary",
			DateStart:     "le 18 octobre 2024",
			DateEnd:       "le 24 decembre 2024",
			Status:        "scan",
			GeneralRules:  "Règles générales, Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi lacus turpis, finibus id semper sit amet, gravida vitae elit. Cras nec ante odio. Nam porta, erat vitae mollis pellentesque, metus orci rutrum arcu, sed tempus est nibh convallis turpis. Nulla eget semper elit, id scelerisque dolor. Fusce lobortis ex vel maximus dapibus. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac augue at mauris finibus dapibus. Sed eu aliquet augue. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Sed arcu nulla, vulputate a pharetra et, lacinia eget sapien. Aliquam pellentesque quam ac lacus dapibus finibus. Maecenas lobortis tincidunt lacinia. Mauris at arcu nec arcu molestie sagittis a et lectus. Quisque luctus viverra lorem quis pretium. Nam vel metus a velit pulvinar ornare. Interdum et malesuada fames ac ante ipsum primis in faucibus. Morbi rhoncus in neque ut mattis. Nulla vulputate aliquet nibh, eget venenatis lacus dapibus eu. Nam sem velit, imperdiet et erat sit amet, rhoncus laoreet dolor. Duis sodales tempor odio, ac imperdiet libero tempor et. Vivamus fermentum massa lacus, et sodales quam suscipit in. Vestibulum felis dui, facilisis id arcu pellentesque, fringilla elementum velit. Nulla iaculis gravida ligula, sed efficitur leo finibus in. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam nec venenatis ligula. Nam laoreet erat ac dui sollicitudin, id mollis tellus auctor. Quisque sodales lorem at felis laoreet, in feugiat velit dapibus. Maecenas mollis egestas auctor. Nullam malesuada neque ac diam aliquam finibus. Duis condimentum tempus enim, eget cursus lectus. Duis ac bibendum est, ut egestas ante.",
			SpecificRules: "Règles spécific, Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi lacus turpis, finibus id semper sit amet, gravida vitae elit. Cras nec ante odio. Nam porta, erat vitae mollis pellentesque, metus orci rutrum arcu, sed tempus est nibh convallis turpis. Nulla eget semper elit, id scelerisque dolor. Fusce lobortis ex vel maximus dapibus. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac augue at mauris finibus dapibus. Sed eu aliquet augue. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Sed arcu nulla, vulputate a pharetra et, lacinia eget sapien. Aliquam pellentesque quam ac lacus dapibus finibus. Maecenas lobortis tincidunt lacinia. Mauris at arcu nec arcu molestie sagittis a et lectus. Quisque luctus viverra lorem quis pretium. Nam vel metus a velit pulvinar ornare. Interdum et malesuada fames ac ante ipsum primis in faucibus. Morbi rhoncus in neque ut mattis. Nulla vulputate aliquet nibh, eget venenatis lacus dapibus eu. Nam sem velit, imperdiet et erat sit amet, rhoncus laoreet dolor. Duis sodales tempor odio, ac imperdiet libero tempor et. Vivamus fermentum massa lacus, et sodales quam suscipit in. Vestibulum felis dui, facilisis id arcu pellentesque, fringilla elementum velit. Nulla iaculis gravida ligula, sed efficitur leo finibus in. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam nec venenatis ligula. Nam laoreet erat ac dui sollicitudin, id mollis tellus auctor. Quisque sodales lorem at felis laoreet, in feugiat velit dapibus. Maecenas mollis egestas auctor. Nullam malesuada neque ac diam aliquam finibus. Duis condimentum tempus enim, eget cursus lectus. Duis ac bibendum est, ut egestas ante.",
			Secret:        "^[0-9]$",
			SecretLength:  10,
			HomePageId:    1,
			ScanPageId:    2,
			ResultPageId:  3,
		},
	}

}
