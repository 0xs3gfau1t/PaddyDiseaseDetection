// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"segFault/PaddyDiseaseDetection/ent/disease"
	"segFault/PaddyDiseaseDetection/ent/diseaseidentified"
	"segFault/PaddyDiseaseDetection/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DiseaseIdentifiedCreate is the builder for creating a DiseaseIdentified entity.
type DiseaseIdentifiedCreate struct {
	config
	mutation *DiseaseIdentifiedMutation
	hooks    []Hook
}

// SetLocation sets the "location" field.
func (dic *DiseaseIdentifiedCreate) SetLocation(s string) *DiseaseIdentifiedCreate {
	dic.mutation.SetLocation(s)
	return dic
}

// SetSeverity sets the "severity" field.
func (dic *DiseaseIdentifiedCreate) SetSeverity(i int) *DiseaseIdentifiedCreate {
	dic.mutation.SetSeverity(i)
	return dic
}

// SetNillableSeverity sets the "severity" field if the given value is not nil.
func (dic *DiseaseIdentifiedCreate) SetNillableSeverity(i *int) *DiseaseIdentifiedCreate {
	if i != nil {
		dic.SetSeverity(*i)
	}
	return dic
}

// SetCreatedAt sets the "created_at" field.
func (dic *DiseaseIdentifiedCreate) SetCreatedAt(t time.Time) *DiseaseIdentifiedCreate {
	dic.mutation.SetCreatedAt(t)
	return dic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dic *DiseaseIdentifiedCreate) SetNillableCreatedAt(t *time.Time) *DiseaseIdentifiedCreate {
	if t != nil {
		dic.SetCreatedAt(*t)
	}
	return dic
}

// SetPhotos sets the "photos" field.
func (dic *DiseaseIdentifiedCreate) SetPhotos(s []string) *DiseaseIdentifiedCreate {
	dic.mutation.SetPhotos(s)
	return dic
}

// SetStatus sets the "status" field.
func (dic *DiseaseIdentifiedCreate) SetStatus(d diseaseidentified.Status) *DiseaseIdentifiedCreate {
	dic.mutation.SetStatus(d)
	return dic
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dic *DiseaseIdentifiedCreate) SetNillableStatus(d *diseaseidentified.Status) *DiseaseIdentifiedCreate {
	if d != nil {
		dic.SetStatus(*d)
	}
	return dic
}

// SetID sets the "id" field.
func (dic *DiseaseIdentifiedCreate) SetID(u uuid.UUID) *DiseaseIdentifiedCreate {
	dic.mutation.SetID(u)
	return dic
}

// SetUploadedByID sets the "uploaded_by" edge to the User entity by ID.
func (dic *DiseaseIdentifiedCreate) SetUploadedByID(id uuid.UUID) *DiseaseIdentifiedCreate {
	dic.mutation.SetUploadedByID(id)
	return dic
}

// SetUploadedBy sets the "uploaded_by" edge to the User entity.
func (dic *DiseaseIdentifiedCreate) SetUploadedBy(u *User) *DiseaseIdentifiedCreate {
	return dic.SetUploadedByID(u.ID)
}

// AddDiseaseIDs adds the "disease" edge to the Disease entity by IDs.
func (dic *DiseaseIdentifiedCreate) AddDiseaseIDs(ids ...uuid.UUID) *DiseaseIdentifiedCreate {
	dic.mutation.AddDiseaseIDs(ids...)
	return dic
}

// AddDisease adds the "disease" edges to the Disease entity.
func (dic *DiseaseIdentifiedCreate) AddDisease(d ...*Disease) *DiseaseIdentifiedCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dic.AddDiseaseIDs(ids...)
}

// Mutation returns the DiseaseIdentifiedMutation object of the builder.
func (dic *DiseaseIdentifiedCreate) Mutation() *DiseaseIdentifiedMutation {
	return dic.mutation
}

// Save creates the DiseaseIdentified in the database.
func (dic *DiseaseIdentifiedCreate) Save(ctx context.Context) (*DiseaseIdentified, error) {
	dic.defaults()
	return withHooks(ctx, dic.sqlSave, dic.mutation, dic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dic *DiseaseIdentifiedCreate) SaveX(ctx context.Context) *DiseaseIdentified {
	v, err := dic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dic *DiseaseIdentifiedCreate) Exec(ctx context.Context) error {
	_, err := dic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dic *DiseaseIdentifiedCreate) ExecX(ctx context.Context) {
	if err := dic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dic *DiseaseIdentifiedCreate) defaults() {
	if _, ok := dic.mutation.Severity(); !ok {
		v := diseaseidentified.DefaultSeverity
		dic.mutation.SetSeverity(v)
	}
	if _, ok := dic.mutation.CreatedAt(); !ok {
		v := diseaseidentified.DefaultCreatedAt()
		dic.mutation.SetCreatedAt(v)
	}
	if _, ok := dic.mutation.Status(); !ok {
		v := diseaseidentified.DefaultStatus
		dic.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dic *DiseaseIdentifiedCreate) check() error {
	if _, ok := dic.mutation.Location(); !ok {
		return &ValidationError{Name: "location", err: errors.New(`ent: missing required field "DiseaseIdentified.location"`)}
	}
	if _, ok := dic.mutation.Severity(); !ok {
		return &ValidationError{Name: "severity", err: errors.New(`ent: missing required field "DiseaseIdentified.severity"`)}
	}
	if v, ok := dic.mutation.Severity(); ok {
		if err := diseaseidentified.SeverityValidator(v); err != nil {
			return &ValidationError{Name: "severity", err: fmt.Errorf(`ent: validator failed for field "DiseaseIdentified.severity": %w`, err)}
		}
	}
	if _, ok := dic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "DiseaseIdentified.created_at"`)}
	}
	if _, ok := dic.mutation.Photos(); !ok {
		return &ValidationError{Name: "photos", err: errors.New(`ent: missing required field "DiseaseIdentified.photos"`)}
	}
	if _, ok := dic.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "DiseaseIdentified.status"`)}
	}
	if v, ok := dic.mutation.Status(); ok {
		if err := diseaseidentified.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "DiseaseIdentified.status": %w`, err)}
		}
	}
	if _, ok := dic.mutation.UploadedByID(); !ok {
		return &ValidationError{Name: "uploaded_by", err: errors.New(`ent: missing required edge "DiseaseIdentified.uploaded_by"`)}
	}
	return nil
}

func (dic *DiseaseIdentifiedCreate) sqlSave(ctx context.Context) (*DiseaseIdentified, error) {
	if err := dic.check(); err != nil {
		return nil, err
	}
	_node, _spec := dic.createSpec()
	if err := sqlgraph.CreateNode(ctx, dic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	dic.mutation.id = &_node.ID
	dic.mutation.done = true
	return _node, nil
}

func (dic *DiseaseIdentifiedCreate) createSpec() (*DiseaseIdentified, *sqlgraph.CreateSpec) {
	var (
		_node = &DiseaseIdentified{config: dic.config}
		_spec = sqlgraph.NewCreateSpec(diseaseidentified.Table, sqlgraph.NewFieldSpec(diseaseidentified.FieldID, field.TypeUUID))
	)
	if id, ok := dic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dic.mutation.Location(); ok {
		_spec.SetField(diseaseidentified.FieldLocation, field.TypeString, value)
		_node.Location = value
	}
	if value, ok := dic.mutation.Severity(); ok {
		_spec.SetField(diseaseidentified.FieldSeverity, field.TypeInt, value)
		_node.Severity = value
	}
	if value, ok := dic.mutation.CreatedAt(); ok {
		_spec.SetField(diseaseidentified.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dic.mutation.Photos(); ok {
		_spec.SetField(diseaseidentified.FieldPhotos, field.TypeJSON, value)
		_node.Photos = value
	}
	if value, ok := dic.mutation.Status(); ok {
		_spec.SetField(diseaseidentified.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := dic.mutation.UploadedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   diseaseidentified.UploadedByTable,
			Columns: []string{diseaseidentified.UploadedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.disease_identified_uploaded_by = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dic.mutation.DiseaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   diseaseidentified.DiseaseTable,
			Columns: diseaseidentified.DiseasePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(disease.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DiseaseIdentifiedCreateBulk is the builder for creating many DiseaseIdentified entities in bulk.
type DiseaseIdentifiedCreateBulk struct {
	config
	err      error
	builders []*DiseaseIdentifiedCreate
}

// Save creates the DiseaseIdentified entities in the database.
func (dicb *DiseaseIdentifiedCreateBulk) Save(ctx context.Context) ([]*DiseaseIdentified, error) {
	if dicb.err != nil {
		return nil, dicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dicb.builders))
	nodes := make([]*DiseaseIdentified, len(dicb.builders))
	mutators := make([]Mutator, len(dicb.builders))
	for i := range dicb.builders {
		func(i int, root context.Context) {
			builder := dicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DiseaseIdentifiedMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dicb *DiseaseIdentifiedCreateBulk) SaveX(ctx context.Context) []*DiseaseIdentified {
	v, err := dicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dicb *DiseaseIdentifiedCreateBulk) Exec(ctx context.Context) error {
	_, err := dicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dicb *DiseaseIdentifiedCreateBulk) ExecX(ctx context.Context) {
	if err := dicb.Exec(ctx); err != nil {
		panic(err)
	}
}
