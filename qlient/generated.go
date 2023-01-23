// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package qlient

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// AddDetectionAddDetection includes the requested fields of the GraphQL type Detection.
type AddDetectionAddDetection struct {
	Id      int     `json:"id"`
	ImageId int     `json:"imageId"`
	Person  float64 `json:"person"`
}

// GetId returns AddDetectionAddDetection.Id, and is useful for accessing the field via an interface.
func (v *AddDetectionAddDetection) GetId() int { return v.Id }

// GetImageId returns AddDetectionAddDetection.ImageId, and is useful for accessing the field via an interface.
func (v *AddDetectionAddDetection) GetImageId() int { return v.ImageId }

// GetPerson returns AddDetectionAddDetection.Person, and is useful for accessing the field via an interface.
func (v *AddDetectionAddDetection) GetPerson() float64 { return v.Person }

// AddDetectionResponse is returned by AddDetection on success.
type AddDetectionResponse struct {
	AddDetection AddDetectionAddDetection `json:"addDetection"`
}

// GetAddDetection returns AddDetectionResponse.AddDetection, and is useful for accessing the field via an interface.
func (v *AddDetectionResponse) GetAddDetection() AddDetectionAddDetection { return v.AddDetection }

// AddImageAddImage includes the requested fields of the GraphQL type Image.
type AddImageAddImage struct {
	Id       int    `json:"id"`
	Filename string `json:"filename"`
}

// GetId returns AddImageAddImage.Id, and is useful for accessing the field via an interface.
func (v *AddImageAddImage) GetId() int { return v.Id }

// GetFilename returns AddImageAddImage.Filename, and is useful for accessing the field via an interface.
func (v *AddImageAddImage) GetFilename() string { return v.Filename }

// AddImageResponse is returned by AddImage on success.
type AddImageResponse struct {
	AddImage AddImageAddImage `json:"addImage"`
}

// GetAddImage returns AddImageResponse.AddImage, and is useful for accessing the field via an interface.
func (v *AddImageResponse) GetAddImage() AddImageAddImage { return v.AddImage }

// DetectionDetection includes the requested fields of the GraphQL type Detection.
type DetectionDetection struct {
	Id     int                     `json:"id"`
	Person float64                 `json:"person"`
	Image  DetectionDetectionImage `json:"image"`
}

// GetId returns DetectionDetection.Id, and is useful for accessing the field via an interface.
func (v *DetectionDetection) GetId() int { return v.Id }

// GetPerson returns DetectionDetection.Person, and is useful for accessing the field via an interface.
func (v *DetectionDetection) GetPerson() float64 { return v.Person }

// GetImage returns DetectionDetection.Image, and is useful for accessing the field via an interface.
func (v *DetectionDetection) GetImage() DetectionDetectionImage { return v.Image }

// DetectionDetectionImage includes the requested fields of the GraphQL type Image.
type DetectionDetectionImage struct {
	Id       int    `json:"id"`
	Filename string `json:"filename"`
}

// GetId returns DetectionDetectionImage.Id, and is useful for accessing the field via an interface.
func (v *DetectionDetectionImage) GetId() int { return v.Id }

// GetFilename returns DetectionDetectionImage.Filename, and is useful for accessing the field via an interface.
func (v *DetectionDetectionImage) GetFilename() string { return v.Filename }

// DetectionResponse is returned by Detection on success.
type DetectionResponse struct {
	Detection DetectionDetection `json:"detection"`
}

// GetDetection returns DetectionResponse.Detection, and is useful for accessing the field via an interface.
func (v *DetectionResponse) GetDetection() DetectionDetection { return v.Detection }

// ImageFilenameImage includes the requested fields of the GraphQL type Image.
type ImageFilenameImage struct {
	Filename string `json:"filename"`
}

// GetFilename returns ImageFilenameImage.Filename, and is useful for accessing the field via an interface.
func (v *ImageFilenameImage) GetFilename() string { return v.Filename }

// ImageFilenameResponse is returned by ImageFilename on success.
type ImageFilenameResponse struct {
	Image ImageFilenameImage `json:"image"`
}

// GetImage returns ImageFilenameResponse.Image, and is useful for accessing the field via an interface.
func (v *ImageFilenameResponse) GetImage() ImageFilenameImage { return v.Image }

// ImageImage includes the requested fields of the GraphQL type Image.
type ImageImage struct {
	Id        int                 `json:"id"`
	Filename  string              `json:"filename"`
	Detection ImageImageDetection `json:"detection"`
}

// GetId returns ImageImage.Id, and is useful for accessing the field via an interface.
func (v *ImageImage) GetId() int { return v.Id }

// GetFilename returns ImageImage.Filename, and is useful for accessing the field via an interface.
func (v *ImageImage) GetFilename() string { return v.Filename }

// GetDetection returns ImageImage.Detection, and is useful for accessing the field via an interface.
func (v *ImageImage) GetDetection() ImageImageDetection { return v.Detection }

// ImageImageDetection includes the requested fields of the GraphQL type Detection.
type ImageImageDetection struct {
	Id     int     `json:"id"`
	Person float64 `json:"person"`
}

// GetId returns ImageImageDetection.Id, and is useful for accessing the field via an interface.
func (v *ImageImageDetection) GetId() int { return v.Id }

// GetPerson returns ImageImageDetection.Person, and is useful for accessing the field via an interface.
func (v *ImageImageDetection) GetPerson() float64 { return v.Person }

// ImageResponse is returned by Image on success.
type ImageResponse struct {
	Image ImageImage `json:"image"`
}

// GetImage returns ImageResponse.Image, and is useful for accessing the field via an interface.
func (v *ImageResponse) GetImage() ImageImage { return v.Image }

// __AddDetectionInput is used internally by genqlient
type __AddDetectionInput struct {
	ImageId int     `json:"imageId"`
	Person  float64 `json:"person"`
}

// GetImageId returns __AddDetectionInput.ImageId, and is useful for accessing the field via an interface.
func (v *__AddDetectionInput) GetImageId() int { return v.ImageId }

// GetPerson returns __AddDetectionInput.Person, and is useful for accessing the field via an interface.
func (v *__AddDetectionInput) GetPerson() float64 { return v.Person }

// __AddImageInput is used internally by genqlient
type __AddImageInput struct {
	Filename string `json:"filename"`
}

// GetFilename returns __AddImageInput.Filename, and is useful for accessing the field via an interface.
func (v *__AddImageInput) GetFilename() string { return v.Filename }

// __DetectionInput is used internally by genqlient
type __DetectionInput struct {
	Id int `json:"id"`
}

// GetId returns __DetectionInput.Id, and is useful for accessing the field via an interface.
func (v *__DetectionInput) GetId() int { return v.Id }

// __ImageFilenameInput is used internally by genqlient
type __ImageFilenameInput struct {
	Id int `json:"id"`
}

// GetId returns __ImageFilenameInput.Id, and is useful for accessing the field via an interface.
func (v *__ImageFilenameInput) GetId() int { return v.Id }

// __ImageInput is used internally by genqlient
type __ImageInput struct {
	Id int `json:"id"`
}

// GetId returns __ImageInput.Id, and is useful for accessing the field via an interface.
func (v *__ImageInput) GetId() int { return v.Id }

func AddDetection(
	ctx context.Context,
	client graphql.Client,
	imageId int,
	person float64,
) (*AddDetectionResponse, error) {
	req := &graphql.Request{
		OpName: "AddDetection",
		Query: `
mutation AddDetection ($imageId: Int!, $person: Float!) {
	addDetection(imageId: $imageId, person: $person) {
		id
		imageId
		person
	}
}
`,
		Variables: &__AddDetectionInput{
			ImageId: imageId,
			Person:  person,
		},
	}
	var err error

	var data AddDetectionResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func AddImage(
	ctx context.Context,
	client graphql.Client,
	filename string,
) (*AddImageResponse, error) {
	req := &graphql.Request{
		OpName: "AddImage",
		Query: `
mutation AddImage ($filename: String!) {
	addImage(filename: $filename) {
		id
		filename
	}
}
`,
		Variables: &__AddImageInput{
			Filename: filename,
		},
	}
	var err error

	var data AddImageResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func Detection(
	ctx context.Context,
	client graphql.Client,
	id int,
) (*DetectionResponse, error) {
	req := &graphql.Request{
		OpName: "Detection",
		Query: `
query Detection ($id: Int!) {
	detection(id: $id) {
		id
		person
		image {
			id
			filename
		}
	}
}
`,
		Variables: &__DetectionInput{
			Id: id,
		},
	}
	var err error

	var data DetectionResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func Image(
	ctx context.Context,
	client graphql.Client,
	id int,
) (*ImageResponse, error) {
	req := &graphql.Request{
		OpName: "Image",
		Query: `
query Image ($id: Int!) {
	image(id: $id) {
		id
		filename
		detection {
			id
			person
		}
	}
}
`,
		Variables: &__ImageInput{
			Id: id,
		},
	}
	var err error

	var data ImageResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}

func ImageFilename(
	ctx context.Context,
	client graphql.Client,
	id int,
) (*ImageFilenameResponse, error) {
	req := &graphql.Request{
		OpName: "ImageFilename",
		Query: `
query ImageFilename ($id: Int!) {
	image(id: $id) {
		filename
	}
}
`,
		Variables: &__ImageFilenameInput{
			Id: id,
		},
	}
	var err error

	var data ImageFilenameResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
