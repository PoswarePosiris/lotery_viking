package database

import (
	"fmt"
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

func Seed(clearData bool) error {
	dbService := New()
	db := dbService.GetDB()
	defer db.Close()

	if clearData {

		for _, table := range Tables {
			_, err := db.Exec(fmt.Sprintf("DELETE FROM %s", table))
			if err != nil {
				return err
			}
		}

		fmt.Println("Tables cleared")
	}

	dataImages := []models.Images{
		{BaseModel: models.BaseModel{ID: 1}, Name: "Image 1", Url: "https://picsum.photos/id/10/200/300", Format: "jpg"},
		{BaseModel: models.BaseModel{ID: 2}, Name: "Image 2", Url: "https://picsum.photos/id/25/200/300", Format: "png"},
		{BaseModel: models.BaseModel{ID: 3}, Name: "Image 3", Url: "https://picsum.photos/id/50/200/300", Format: "jpg"},
		{BaseModel: models.BaseModel{ID: 4}, Name: "Image 4 pub", Url: "https://picsum.photos/100/1/200/300", Format: "png"},
		{BaseModel: models.BaseModel{ID: 5}, Name: "Image 5 pub", Url: "https://picsum.photos/125/1/200/300", Format: "jpg"},
	}

	stmtImage, err := db.Prepare("INSERT INTO images (id, name, url, format) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmtImage.Close()

	for _, image := range dataImages {
		_, err = stmtImage.Exec(image.ID, image.Name, image.Url, image.Format)
		if err != nil {
			return err
		}
	}

	dataParameters := []models.Parameters{
		{BaseModel: models.BaseModel{ID: 1},
			NameLotery:    "Lottery 1",
			NameCasino:    "Casino de Sanary",
			DateStart:     "le 18 octobre 2024",
			DateEnd:       "le 24 decembre 2024",
			Status:        models.Scan,
			ClientData:    true,
			GeneralRules:  "Règles générales, Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi lacus turpis, finibus id semper sit amet, gravida vitae elit. Cras nec ante odio. Nam porta, erat vitae mollis pellentesque, metus orci rutrum arcu, sed tempus est nibh convallis turpis. Nulla eget semper elit, id scelerisque dolor. Fusce lobortis ex vel maximus dapibus. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac augue at mauris finibus dapibus. Sed eu aliquet augue. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Sed arcu nulla, vulputate a pharetra et, lacinia eget sapien. Aliquam pellentesque quam ac lacus dapibus finibus. Maecenas lobortis tincidunt lacinia. Mauris at arcu nec arcu molestie sagittis a et lectus. Quisque luctus viverra lorem quis pretium. Nam vel metus a velit pulvinar ornare. Interdum et malesuada fames ac ante ipsum primis in faucibus. Morbi rhoncus in neque ut mattis. Nulla vulputate aliquet nibh, eget venenatis lacus dapibus eu. Nam sem velit, imperdiet et erat sit amet, rhoncus laoreet dolor. Duis sodales tempor odio, ac imperdiet libero tempor et. Vivamus fermentum massa lacus, et sodales quam suscipit in. Vestibulum felis dui, facilisis id arcu pellentesque, fringilla elementum velit. Nulla iaculis gravida ligula, sed efficitur leo finibus in. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam nec venenatis ligula. Nam laoreet erat ac dui sollicitudin, id mollis tellus auctor. Quisque sodales lorem at felis laoreet, in feugiat velit dapibus. Maecenas mollis egestas auctor. Nullam malesuada neque ac diam aliquam finibus. Duis condimentum tempus enim, eget cursus lectus. Duis ac bibendum est, ut egestas ante.",
			SpecificRules: "Règles spécific, Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi lacus turpis, finibus id semper sit amet, gravida vitae elit. Cras nec ante odio. Nam porta, erat vitae mollis pellentesque, metus orci rutrum arcu, sed tempus est nibh convallis turpis. Nulla eget semper elit, id scelerisque dolor. Fusce lobortis ex vel maximus dapibus. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac augue at mauris finibus dapibus. Sed eu aliquet augue. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Sed arcu nulla, vulputate a pharetra et, lacinia eget sapien. Aliquam pellentesque quam ac lacus dapibus finibus. Maecenas lobortis tincidunt lacinia. Mauris at arcu nec arcu molestie sagittis a et lectus. Quisque luctus viverra lorem quis pretium. Nam vel metus a velit pulvinar ornare. Interdum et malesuada fames ac ante ipsum primis in faucibus. Morbi rhoncus in neque ut mattis. Nulla vulputate aliquet nibh, eget venenatis lacus dapibus eu. Nam sem velit, imperdiet et erat sit amet, rhoncus laoreet dolor. Duis sodales tempor odio, ac imperdiet libero tempor et. Vivamus fermentum massa lacus, et sodales quam suscipit in. Vestibulum felis dui, facilisis id arcu pellentesque, fringilla elementum velit. Nulla iaculis gravida ligula, sed efficitur leo finibus in. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam nec venenatis ligula. Nam laoreet erat ac dui sollicitudin, id mollis tellus auctor. Quisque sodales lorem at felis laoreet, in feugiat velit dapibus. Maecenas mollis egestas auctor. Nullam malesuada neque ac diam aliquam finibus. Duis condimentum tempus enim, eget cursus lectus. Duis ac bibendum est, ut egestas ante.",
			Secret:        "^[0-9]$",
			SecretLength:  10,
			HomePageId:    1,
			ScanPageId:    2,
			ResultPageId:  3,
		},
	}

	stmtParam, err := db.Prepare("INSERT INTO parameters (id, name_lotery, name_casino,  date_start, date_end, status, general_rules, specific_rules, secret, secret_length, home_page, scan_page, result_page) VALUES (?, ?, ?, ?,?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmtParam.Close()

	for _, parameters := range dataParameters {
		_, err = stmtParam.Exec(parameters.ID, parameters.NameLotery, parameters.NameCasino, parameters.DateStart, parameters.DateEnd, parameters.Status, parameters.GeneralRules, parameters.SpecificRules, parameters.Secret, parameters.SecretLength, parameters.HomePageId, parameters.ScanPageId, parameters.ResultPageId)
		if err != nil {
			return err
		}
	}

	dataPublicity := []models.PublicityImages{
		{
			ParameterId: 1,
			ImageId:     4,
		},
		{
			ParameterId: 1,
			ImageId:     5,
		},
	}

	stmtPublicity, err := db.Prepare("INSERT INTO publicity_images (parameter_id, image_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmtPublicity.Close()

	for _, publicity := range dataPublicity {
		_, err = stmtPublicity.Exec(publicity.ParameterId, publicity.ImageId)
		if err != nil {
			return err
		}
	}

	dataKiosks := []models.Kiosks{
		{
			Name:              "Kiosk 1 Sanary",
			Location:          "Sanary sur Mer",
			MacadressWifi:     "7C:0A:3F:A3:1B:90",
			MacadressEthernet: "C0:23:8D:A1:7B:6E",
			IdParameters:      1,
		},
		{
			Name:              "Kiosk 2 Sanary",
			Location:          "Sanary sur Mer",
			MacadressWifi:     "D8:A3:5C:E6:97:46",
			MacadressEthernet: "7C:0A:3F:F5:2F:EC",
			IdParameters:      1,
		},
	}

	stmtKiosks, err := db.Prepare("INSERT INTO kiosks (name, macadress_wifi, macadress_ethernet, location, id_parameters) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmtKiosks.Close()

	for _, kiosks := range dataKiosks {
		_, err = stmtKiosks.Exec(kiosks.Name, kiosks.MacadressWifi, kiosks.MacadressEthernet, kiosks.Location, kiosks.IdParameters)
		if err != nil {
			return err
		}
	}

	return nil
}
