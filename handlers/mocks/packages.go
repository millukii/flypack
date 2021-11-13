package mocks

import "flypack/models"

var ShipmentPackages = &models.AllShipmentPackagesResponse{

	Count: 6,
	PackagesList:[]*models.ShipmentPackage{
		{
			Id: "1",
			Size: "X",
			OrderNumber: "1",
		},
				{
						Id: "2",
			Size: "X",
			OrderNumber: "2",
		},
				{
						Id: "3",
			Size: "X",
			OrderNumber: "3",
		},
				{
						Id: "4",
			Size: "X",
			OrderNumber: "4",
		},
				{
						Id: "5",
			Size: "X",
			OrderNumber: "5",
		},
				{
						Id: "6",
			Size: "X",
			OrderNumber: "5",
		},
	},

}
