package mocks

import "flypack/models"

var Manifiests = &models.AllManifiestResponse{
	Manifiests: []*models.Manifiest{
		{
				ID: "1",
				ManifiestNumber: "1",
				DeliveryDate: "sdsa",
				City: "sdfdsf",
				Commune: "dsfsdf",
				SmallPackagesCapacity: 1,
				MediumPackagesCapacity: 2,
				XtraPackagesCapacity: 4,
				PackageList: []*models.ShipmentPackage{
					{
						Size: "M",
						OrderNumber: "1",
					},
				},
		},
				{
				ID: "2",
				ManifiestNumber: "2",
				DeliveryDate: "sdsa",
				City: "sdfdsf",
				Commune: "dsfsdf",
				SmallPackagesCapacity: 1,
				MediumPackagesCapacity: 2,
				XtraPackagesCapacity: 4,
				PackageList: []*models.ShipmentPackage{
					{
						Size: "M",
						OrderNumber: "1",
					},
				},
		},
	},
}