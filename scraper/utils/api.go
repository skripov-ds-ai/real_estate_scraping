package utils

import (
	"net/url"
)

// https://my.matterport.com/api/mp/models/graph

func CreateHeaders(userAgent, referer string) (m map[string]string) {
	//accept: */*
	//accept-encoding: gzip, deflate, br
	//accept-language: en-US,en;q=0.9,ru;q=0.8,ja;q=0.7
	//content-length: 2953
	//content-type: application/json
	//dnt: 1
	//origin: https://my.matterport.com
	//referer: https://my.matterport.com/show/?play=1&m=s2q4VDSQsbY
	//sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="102", "Google Chrome";v="102"
	//sec-ch-ua-mobile: ?0
	//sec-ch-ua-platform: "Linux"
	//sec-fetch-dest: empty
	//sec-fetch-mode: cors
	//sec-fetch-site: same-origin
	//user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36
	//x-matterport-application-name: showcase
	//x-matterport-application-version: 3.1.69.3-0-g204b0700e6
	//h.Set("User-Agent", userAgent)
	//h.Set("Accept", "*/*")
	m = map[string]string{
		"accept":          "*/*",
		"accept-encoding": "gzip, deflate, br",
		"accept-language": "en-US,en;q=0.9,ru;q=0.8,ja;q=0.7",
		//"content-length":  "2953",
		"content-type": "application/json",
		"dnt":          "1",
		"origin":       "https://my.matterport.com",
		"referer":      referer,
		//"sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="102", "Google Chrome";v="102"
		//"sec-ch-ua-mobile: ?0
		//"sec-ch-ua-platform: "Linux"
		//"sec-fetch-dest: empty
		//"sec-fetch-mode: cors
		//"sec-fetch-site: same-origin
		"user-agent":                       userAgent,
		"x-matterport-application-name":    "showcase",
		"x-matterport-application-version": "3.1.69.3-0-g204b0700e6",
	}
	//for k, v := range m {
	//	h.Set(k, v)
	//}
	return
}

func GetModelId(urlString string) (modelId *string, err error) {
	//	https://my.matterport.com/show/?play=1&m=yFHoSPfUWZF
	u, err := url.Parse(urlString)
	if err != nil {
		return
	}

	m, _ := url.ParseQuery(u.RawQuery)
	if result, ok := m["m"]; ok && len(result) > 0 {
		modelId = &result[0]
	}
	return
}

func CreatePayload(modelId string) (payload map[string]interface{}) {
	//payload
	//{operationName: "GetModelDetails", variables: {modelId: "s2q4VDSQsbY"},…}
	//operationName: "GetModelDetails"
	//query: "query GetModelDetails($modelId: ID!) {\n  model(id: $modelId) {\n    id\n    name\n    organization\n    visibility\n    discoverable\n    created\n    ...ModelStatus\n    image {\n      ...PhotoDetails\n      __typename\n    }\n    publication {\n      ...PublicationDetails\n      __typename\n    }\n    options {\n      ...ModelOptions\n      __typename\n    }\n    ...ModelAvailableAssets\n    __typename\n  }\n}\n\nfragment ModelStatus on Model {\n  state\n  __typename\n}\n\nfragment PhotoDetails on Photo {\n  id\n  label\n  classification\n  category\n  height\n  width\n  created\n  modified\n  status\n  filename\n  format\n  url\n  resolutions\n  type\n  origin\n  validUntil\n  thumbnailUrl: resizeUrl(resolution: thumbnail)\n  presentationUrl: resizeUrl(resolution: presentation)\n  snapshotLocation {\n    viewMode\n    position {\n      x\n      y\n      z\n      __typename\n    }\n    rotation {\n      x\n      y\n      z\n      w\n      __typename\n    }\n    zoom\n    floorVisibility {\n      id\n      meshId\n      sequence\n      __typename\n    }\n    anchor {\n      id\n      pano {\n        id\n        placement\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment PublicationDetails on ModelPublicationDetails {\n  address\n  published\n  lastPublished\n  presentedBy\n  summary\n  externalUrl\n  contact {\n    name\n    email\n    phoneNumber\n    __typename\n  }\n  __typename\n}\n\nfragment ModelOptions on ModelOptions {\n  backgroundColor\n  dollhouseEnabled\n  dollhouseLabelsEnabled\n  floorSelectEnabled\n  floorplanEnabled\n  highlightReelEnabled\n  labelsEnabled\n  measurements\n  socialSharingEnabled\n  spaceSearchEnabled\n  tourButtonsEnabled\n  tourDollhousePanSpeed\n  tourFastTransitionsEnabled\n  tourPanAngle\n  tourPanDirection\n  tourPanSpeed\n  tourTransitionSpeed\n  tourTransitionTime\n  tourZoomDuration\n  unitType\n  vrEnabled\n  __typename\n}\n\nfragment ModelAvailableAssets on Model {\n  assets {\n    meshes(formats: \"dam\", resolutions: [\"50k\", \"500k\"], compressions: none) {\n      ...MeshDetails\n      __typename\n    }\n    textures {\n      ...TextureDetails\n      __typename\n    }\n    tilesets {\n      ...TilesetDetails\n      __typename\n    }\n    __typename\n  }\n  lod: policy(name: \"spaces.chunked.mesh.lod\") {\n    ... on PolicyOptions {\n      options\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment MeshDetails on Mesh {\n  id\n  status\n  filename\n  format\n  resolution\n  url\n  validUntil\n  __typename\n}\n\nfragment TextureDetails on Texture {\n  id\n  status\n  format\n  resolution\n  quality\n  urlTemplate\n  validUntil\n  __typename\n}\n\nfragment TilesetDetails on Tileset {\n  status\n  tilesetVersion\n  url\n  urlTemplate\n  validUntil\n  tilesetDepth\n  tilesetPreset\n  __typename\n}\n"
	//variables: {modelId: "s2q4VDSQsbY"}
	payload = make(map[string]interface{})
	payload["operationName"] = "GetModelDetails"
	payload["query"] = `query GetModelDetails($modelId: ID!) {
	  model(id: $modelId) {
		id
		name
		organization
		visibility
		discoverable
		created
		...ModelStatus
		image {
		  ...PhotoDetails
		  __typename
		}
		publication {
		  ...PublicationDetails
		  __typename
		}
		options {
		  ...ModelOptions
		  __typename
		}
		...ModelAvailableAssets
		__typename
	  }
	}
	
	fragment ModelStatus on Model {
	  state
	  __typename
	}
	
	fragment PhotoDetails on Photo {
	  id
	  label
	  classification
	  category
	  height
	  width
	  created
	  modified
	  status
	  filename
	  format
	  url
	  resolutions
	  type
	  origin
	  validUntil
	  thumbnailUrl: resizeUrl(resolution: thumbnail)
	  presentationUrl: resizeUrl(resolution: presentation)
	  snapshotLocation {
		viewMode
		position {
		  x
		  y
		  z
		  __typename
		}
		rotation {
		  x
		  y
		  z
		  w
		  __typename
		}
		zoom
		floorVisibility {
		  id
		  meshId
		  sequence
		  __typename
		}
		anchor {
		  id
		  pano {
			id
			placement
			__typename
		  }
		  __typename
		}
		__typename
	  }
	  __typename
	}
	
	fragment PublicationDetails on ModelPublicationDetails {
	  address
	  published
	  lastPublished
	  presentedBy
	  summary
	  externalUrl
	  contact {
		name
		email
		phoneNumber
		__typename
	  }
	  __typename
	}
	
	fragment ModelOptions on ModelOptions {
	  backgroundColor
	  dollhouseEnabled
	  dollhouseLabelsEnabled
	  floorSelectEnabled
	  floorplanEnabled
	  highlightReelEnabled
	  labelsEnabled
	  measurements
	  socialSharingEnabled
	  spaceSearchEnabled
	  tourButtonsEnabled
	  tourDollhousePanSpeed
	  tourFastTransitionsEnabled
	  tourPanAngle
	  tourPanDirection
	  tourPanSpeed
	  tourTransitionSpeed
	  tourTransitionTime
	  tourZoomDuration
	  unitType
	  vrEnabled
	  __typename
	}
	
	fragment ModelAvailableAssets on Model {
	  assets {
		meshes(formats: "dam", resolutions: ["50k", "500k"], compressions: none) {
		  ...MeshDetails
		  __typename
		}
		textures {
		  ...TextureDetails
		  __typename
		}
		tilesets {
		  ...TilesetDetails
		  __typename
		}
		__typename
	  }
	  lod: policy(name: "spaces.chunked.mesh.lod") {
		... on PolicyOptions {
		  options
		  __typename
		}
		__typename
	  }
	  __typename
	}
	
	fragment MeshDetails on Mesh {
	  id
	  status
	  filename
	  format
	  resolution
	  url
	  validUntil
	  __typename
	}
	
	fragment TextureDetails on Texture {
	  id
	  status
	  format
	  resolution
	  quality
	  urlTemplate
	  validUntil
	  __typename
	}
	
	fragment TilesetDetails on Tileset {
	  status
	  tilesetVersion
	  url
	  urlTemplate
	  validUntil
	  tilesetDepth
	  tilesetPreset
	  __typename
	}`
	payload["variables"] = map[string]string{
		"modelId": modelId,
	}
	//	payload["operationName"] = "getSnapShots"
	//	payload["query"] = `query GetSnapshots($modelId: ID!) {
	//		model(id: $modelId) {
	//			id
	//			assets {
	//				photos {
	//					...PhotoDetails
	//					__typename
	//			}
	//			__typename
	//		}
	//		__typename
	//	}
	//	}
	//
	//	fragment PhotoDetails on Photo {
	//		id
	//		label
	//		classification
	//		category
	//		height
	//		width
	//		created
	//		modified
	//		status
	//		filename
	//		format
	//		url
	//		resolutions
	//		type
	//		origin
	//		validUntil
	//		thumbnailUrl: resizeUrl(resolution: thumbnail)
	//		presentationUrl: resizeUrl(resolution: presentation)
	//		snapshotLocation {
	//			viewMode
	//			position {
	//				x
	//				y
	//				z
	//				__typename
	//		}
	//		rotation {
	//			x
	//			y
	//			z
	//			w
	//			__typename
	//		}
	//		zoom
	//		floorVisibility {
	//			id
	//			meshId
	//			sequence
	//			__typename
	//		}
	//		anchor {
	//			id
	//			pano {
	//				id
	//				placement
	//				__typename
	//		}
	//		__typename
	//	}
	//	__typename
	//	}
	//	__typename
	//	}
	//`
	return
}
