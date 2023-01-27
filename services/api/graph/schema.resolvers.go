package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"

	"github.com/moritztng/perceptionOS/services/api/graph/model"
)

// Image is the resolver for the image field.
func (r *detectionResolver) Image(ctx context.Context, obj *model.Detection) (*model.Image, error) {
	image := r.DB.GetImage(uint(obj.ImageID))
	return &model.Image{ID: int(image.ID), Filename: image.Filename}, nil
}

// Detection is the resolver for the detection field.
func (r *imageResolver) Detection(ctx context.Context, obj *model.Image) (*model.Detection, error) {
	fmt.Printf("hi")
	fmt.Print(obj.ID)
	detection := r.DB.GetDetectionOfImage(uint((*obj).ID))
	fmt.Print(detection)
	return &model.Detection{ID: int(detection.ID), ImageID: obj.ID, Person: float64(detection.Person)}, nil
}

// AddImage is the resolver for the addImage field.
func (r *mutationResolver) AddImage(ctx context.Context, filename string) (*model.Image, error) {
	image := r.DB.AddImage(filename)
	return &model.Image{ID: int(image.ID), Filename: image.Filename}, nil
}

// AddDetection is the resolver for the addDetection field.
func (r *mutationResolver) AddDetection(ctx context.Context, imageID int, person float64) (*model.Detection, error) {
	// AddDetection is the resolver for the addDetection field.
	detection := r.DB.AddDetection(uint(imageID), float32(person))
	return &model.Detection{ID: int(detection.ID), ImageID: imageID, Person: float64(detection.Person)}, nil
}

// Image is the resolver for the image field.
func (r *queryResolver) Image(ctx context.Context, id int) (*model.Image, error) {
	image := r.DB.GetImage(uint(id))
	return &model.Image{ID: int(image.ID), Filename: image.Filename}, nil
}

// Detection is the resolver for the detection field.
func (r *queryResolver) Detection(ctx context.Context, id int) (*model.Detection, error) {
	detection := r.DB.GetDetection(uint(id))
	return &model.Detection{ID: int(detection.ID), ImageID: int(detection.ImageID), Person: float64(detection.Person)}, nil
}

// TakeImage is the resolver for the takeImage field.
func (r *queryResolver) TakeImage(ctx context.Context) (string, error) {
	r.MessageProducer.PublishImageRequest()
	return "", nil
}

// Detection returns DetectionResolver implementation.
func (r *Resolver) Detection() DetectionResolver { return &detectionResolver{r} }

// Image returns ImageResolver implementation.
func (r *Resolver) Image() ImageResolver { return &imageResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type detectionResolver struct{ *Resolver }
type imageResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }