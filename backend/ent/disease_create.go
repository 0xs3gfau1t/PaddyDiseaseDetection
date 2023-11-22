// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"segFault/PaddyDiseaseDetection/ent/disease"
	"segFault/PaddyDiseaseDetection/ent/diseaseidentified"
	"segFault/PaddyDiseaseDetection/ent/solution"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DiseaseCreate is the builder for creating a Disease entity.
type DiseaseCreate struct {
	config
	mutation *DiseaseMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (dc *DiseaseCreate) SetName(s string) *DiseaseCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetPhotos sets the "photos" field.
func (dc *DiseaseCreate) SetPhotos(s []string) *DiseaseCreate {
	dc.mutation.SetPhotos(s)
	return dc
}

// SetID sets the "id" field.
func (dc *DiseaseCreate) SetID(u uuid.UUID) *DiseaseCreate {
	dc.mutation.SetID(u)
	return dc
}

// AddSolutionIDs adds the "solutions" edge to the Solution entity by IDs.
func (dc *DiseaseCreate) AddSolutionIDs(ids ...uuid.UUID) *DiseaseCreate {
	dc.mutation.AddSolutionIDs(ids...)
	return dc
}

// AddSolutions adds the "solutions" edges to the Solution entity.
func (dc *DiseaseCreate) AddSolutions(s ...*Solution) *DiseaseCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return dc.AddSolutionIDs(ids...)
}

// AddDiseaseIdentifiedIDs adds the "disease_identified" edge to the DiseaseIdentified entity by IDs.
func (dc *DiseaseCreate) AddDiseaseIdentifiedIDs(ids ...uuid.UUID) *DiseaseCreate {
	dc.mutation.AddDiseaseIdentifiedIDs(ids...)
	return dc
}

// AddDiseaseIdentified adds the "disease_identified" edges to the DiseaseIdentified entity.
func (dc *DiseaseCreate) AddDiseaseIdentified(d ...*DiseaseIdentified) *DiseaseCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddDiseaseIdentifiedIDs(ids...)
}

// Mutation returns the DiseaseMutation object of the builder.
func (dc *DiseaseCreate) Mutation() *DiseaseMutation {
	return dc.mutation
}

// Save creates the Disease in the database.
func (dc *DiseaseCreate) Save(ctx context.Context) (*Disease, error) {
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DiseaseCreate) SaveX(ctx context.Context) *Disease {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DiseaseCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DiseaseCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DiseaseCreate) check() error {
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Disease.name"`)}
	}
	if _, ok := dc.mutation.Photos(); !ok {
		return &ValidationError{Name: "photos", err: errors.New(`ent: missing required field "Disease.photos"`)}
	}
	return nil
}

func (dc *DiseaseCreate) sqlSave(ctx context.Context) (*Disease, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
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
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DiseaseCreate) createSpec() (*Disease, *sqlgraph.CreateSpec) {
	var (
		_node = &Disease{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(disease.Table, sqlgraph.NewFieldSpec(disease.FieldID, field.TypeUUID))
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(disease.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dc.mutation.Photos(); ok {
		_spec.SetField(disease.FieldPhotos, field.TypeJSON, value)
		_node.Photos = value
	}
	if nodes := dc.mutation.SolutionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   disease.SolutionsTable,
			Columns: disease.SolutionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(solution.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.DiseaseIdentifiedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   disease.DiseaseIdentifiedTable,
			Columns: disease.DiseaseIdentifiedPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(diseaseidentified.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DiseaseCreateBulk is the builder for creating many Disease entities in bulk.
type DiseaseCreateBulk struct {
	config
	err      error
	builders []*DiseaseCreate
}

// Save creates the Disease entities in the database.
func (dcb *DiseaseCreateBulk) Save(ctx context.Context) ([]*Disease, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Disease, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DiseaseMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DiseaseCreateBulk) SaveX(ctx context.Context) []*Disease {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DiseaseCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DiseaseCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}