package querybuilder

import "go.mongodb.org/mongo-driver/bson"

type Builder struct {
	filter bson.M
}

func New() *Builder {
	return &Builder{filter: bson.M{}}
}

func (b *Builder) Eq(key string, value any) *Builder {
	b.filter[key] = value
	return b
}

func (b *Builder) Regex(key string, value string) *Builder {
	b.filter[key] = bson.M{"$regex": value, "$options": "i"}
	return b
}

func (b *Builder) Build() bson.M {
	return b.filter
}
