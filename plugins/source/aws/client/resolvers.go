package client

import (
	"context"
	"reflect"

	"github.com/cloudquery/plugin-sdk/schema"
)

func ResolveAWSAccount(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return r.Set(c.Name, client.AccountID)
}

func ResolveAWSRegion(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return r.Set(c.Name, client.Region)
}

func ResolveAWSNamespace(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return r.Set(c.Name, client.AutoscalingNamespace)
}

func ResolveWAFScope(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	return r.Set(c.Name, meta.(*Client).WAFScope)
}

func ResolveTags(ctx context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	return ResolveTagField("Tags")(ctx, meta, r, c)
}

func ResolveTagField(fieldName string) func(context.Context, schema.ClientMeta, *schema.Resource, schema.Column) error {
	return func(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		var val reflect.Value

		if reflect.TypeOf(r.Item).Kind() == reflect.Ptr {
			val = reflect.ValueOf(r.Item).Elem()
		} else {
			val = reflect.ValueOf(r.Item)
		}

		if val.Kind() != reflect.Struct {
			panic("need struct type")
		}
		f := val.FieldByName(fieldName)
		if f.IsNil() {
			return r.Set(c.Name, map[string]string{}) // can't have nil or the integration test will make a fuss
		} else if f.IsZero() {
			panic("no such field " + fieldName)
		}
		data := TagsToMap(f.Interface())
		return r.Set(c.Name, data)
	}
}
