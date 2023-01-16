// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"context"
	"errors"
	"fmt"
	"image/color"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/MichaelMure/git-bug/api/graphql/models"
	"github.com/MichaelMure/git-bug/entities/bug"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type LabelResolver interface {
	Name(ctx context.Context, obj *bug.Label) (string, error)
	Color(ctx context.Context, obj *bug.Label) (*color.RGBA, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Label_name(ctx context.Context, field graphql.CollectedField, obj *bug.Label) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Label_name(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Label().Name(rctx, obj)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Label_name(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Label",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Label_color(ctx context.Context, field graphql.CollectedField, obj *bug.Label) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Label_color(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Label().Color(rctx, obj)
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*color.RGBA)
	fc.Result = res
	return ec.marshalNColor2ᚖimageᚋcolorᚐRGBA(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Label_color(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Label",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "R":
				return ec.fieldContext_Color_R(ctx, field)
			case "G":
				return ec.fieldContext_Color_G(ctx, field)
			case "B":
				return ec.fieldContext_Color_B(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Color", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _LabelConnection_edges(ctx context.Context, field graphql.CollectedField, obj *models.LabelConnection) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_LabelConnection_edges(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Edges, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]*models.LabelEdge)
	fc.Result = res
	return ec.marshalNLabelEdge2ᚕᚖgithubᚗcomᚋMichaelMureᚋgitᚑbugᚋapiᚋgraphqlᚋmodelsᚐLabelEdgeᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_LabelConnection_edges(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "LabelConnection",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "cursor":
				return ec.fieldContext_LabelEdge_cursor(ctx, field)
			case "node":
				return ec.fieldContext_LabelEdge_node(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type LabelEdge", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _LabelConnection_nodes(ctx context.Context, field graphql.CollectedField, obj *models.LabelConnection) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_LabelConnection_nodes(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Nodes, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.([]bug.Label)
	fc.Result = res
	return ec.marshalNLabel2ᚕgithubᚗcomᚋMichaelMureᚋgitᚑbugᚋentitiesᚋbugᚐLabelᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_LabelConnection_nodes(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "LabelConnection",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "name":
				return ec.fieldContext_Label_name(ctx, field)
			case "color":
				return ec.fieldContext_Label_color(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Label", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _LabelConnection_pageInfo(ctx context.Context, field graphql.CollectedField, obj *models.LabelConnection) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_LabelConnection_pageInfo(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.PageInfo, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*models.PageInfo)
	fc.Result = res
	return ec.marshalNPageInfo2ᚖgithubᚗcomᚋMichaelMureᚋgitᚑbugᚋapiᚋgraphqlᚋmodelsᚐPageInfo(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_LabelConnection_pageInfo(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "LabelConnection",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "hasNextPage":
				return ec.fieldContext_PageInfo_hasNextPage(ctx, field)
			case "hasPreviousPage":
				return ec.fieldContext_PageInfo_hasPreviousPage(ctx, field)
			case "startCursor":
				return ec.fieldContext_PageInfo_startCursor(ctx, field)
			case "endCursor":
				return ec.fieldContext_PageInfo_endCursor(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type PageInfo", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _LabelConnection_totalCount(ctx context.Context, field graphql.CollectedField, obj *models.LabelConnection) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_LabelConnection_totalCount(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.TotalCount, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(int)
	fc.Result = res
	return ec.marshalNInt2int(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_LabelConnection_totalCount(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "LabelConnection",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type Int does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _LabelEdge_cursor(ctx context.Context, field graphql.CollectedField, obj *models.LabelEdge) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_LabelEdge_cursor(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Cursor, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_LabelEdge_cursor(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "LabelEdge",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _LabelEdge_node(ctx context.Context, field graphql.CollectedField, obj *models.LabelEdge) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_LabelEdge_node(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Node, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(bug.Label)
	fc.Result = res
	return ec.marshalNLabel2githubᚗcomᚋMichaelMureᚋgitᚑbugᚋentitiesᚋbugᚐLabel(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_LabelEdge_node(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "LabelEdge",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "name":
				return ec.fieldContext_Label_name(ctx, field)
			case "color":
				return ec.fieldContext_Label_color(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Label", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var labelImplementors = []string{"Label"}

func (ec *executionContext) _Label(ctx context.Context, sel ast.SelectionSet, obj *bug.Label) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, labelImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Label")
		case "name":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Label_name(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return innerFunc(ctx)

			})
		case "color":
			field := field

			innerFunc := func(ctx context.Context) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Label_color(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			}

			out.Concurrently(i, func() graphql.Marshaler {
				return innerFunc(ctx)

			})
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var labelConnectionImplementors = []string{"LabelConnection"}

func (ec *executionContext) _LabelConnection(ctx context.Context, sel ast.SelectionSet, obj *models.LabelConnection) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, labelConnectionImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("LabelConnection")
		case "edges":

			out.Values[i] = ec._LabelConnection_edges(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "nodes":

			out.Values[i] = ec._LabelConnection_nodes(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "pageInfo":

			out.Values[i] = ec._LabelConnection_pageInfo(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "totalCount":

			out.Values[i] = ec._LabelConnection_totalCount(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var labelEdgeImplementors = []string{"LabelEdge"}

func (ec *executionContext) _LabelEdge(ctx context.Context, sel ast.SelectionSet, obj *models.LabelEdge) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, labelEdgeImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("LabelEdge")
		case "cursor":

			out.Values[i] = ec._LabelEdge_cursor(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "node":

			out.Values[i] = ec._LabelEdge_node(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNLabel2githubᚗcomᚋMichaelMureᚋgitᚑbugᚋentitiesᚋbugᚐLabel(ctx context.Context, sel ast.SelectionSet, v bug.Label) graphql.Marshaler {
	return ec._Label(ctx, sel, &v)
}

func (ec *executionContext) marshalNLabel2ᚕgithubᚗcomᚋMichaelMureᚋgitᚑbugᚋentitiesᚋbugᚐLabelᚄ(ctx context.Context, sel ast.SelectionSet, v []bug.Label) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNLabel2githubᚗcomᚋMichaelMureᚋgitᚑbugᚋentitiesᚋbugᚐLabel(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalNLabelConnection2githubᚗcomᚋMichaelMureᚋgitᚑbugᚋapiᚋgraphqlᚋmodelsᚐLabelConnection(ctx context.Context, sel ast.SelectionSet, v models.LabelConnection) graphql.Marshaler {
	return ec._LabelConnection(ctx, sel, &v)
}

func (ec *executionContext) marshalNLabelConnection2ᚖgithubᚗcomᚋMichaelMureᚋgitᚑbugᚋapiᚋgraphqlᚋmodelsᚐLabelConnection(ctx context.Context, sel ast.SelectionSet, v *models.LabelConnection) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._LabelConnection(ctx, sel, v)
}

func (ec *executionContext) marshalNLabelEdge2ᚕᚖgithubᚗcomᚋMichaelMureᚋgitᚑbugᚋapiᚋgraphqlᚋmodelsᚐLabelEdgeᚄ(ctx context.Context, sel ast.SelectionSet, v []*models.LabelEdge) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNLabelEdge2ᚖgithubᚗcomᚋMichaelMureᚋgitᚑbugᚋapiᚋgraphqlᚋmodelsᚐLabelEdge(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalNLabelEdge2ᚖgithubᚗcomᚋMichaelMureᚋgitᚑbugᚋapiᚋgraphqlᚋmodelsᚐLabelEdge(ctx context.Context, sel ast.SelectionSet, v *models.LabelEdge) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._LabelEdge(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
