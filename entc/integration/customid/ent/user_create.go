// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/customid/ent/group"
	"github.com/facebookincubator/ent/entc/integration/customid/ent/pet"
	"github.com/facebookincubator/ent/entc/integration/customid/ent/user"
	"github.com/facebookincubator/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetID sets the id field.
func (uc *UserCreate) SetID(i int) *UserCreate {
	uc.mutation.SetID(i)
	return uc
}

// AddGroupIDs adds the groups edge to Group by ids.
func (uc *UserCreate) AddGroupIDs(ids ...int) *UserCreate {
	uc.mutation.AddGroupIDs(ids...)
	return uc
}

// AddGroups adds the groups edges to Group.
func (uc *UserCreate) AddGroups(g ...*Group) *UserCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uc.AddGroupIDs(ids...)
}

// SetParentID sets the parent edge to User by id.
func (uc *UserCreate) SetParentID(id int) *UserCreate {
	uc.mutation.SetParentID(id)
	return uc
}

// SetNillableParentID sets the parent edge to User by id if the given value is not nil.
func (uc *UserCreate) SetNillableParentID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetParentID(*id)
	}
	return uc
}

// SetParent sets the parent edge to User.
func (uc *UserCreate) SetParent(u *User) *UserCreate {
	return uc.SetParentID(u.ID)
}

// AddChildIDs adds the children edge to User by ids.
func (uc *UserCreate) AddChildIDs(ids ...int) *UserCreate {
	uc.mutation.AddChildIDs(ids...)
	return uc
}

// AddChildren adds the children edges to User.
func (uc *UserCreate) AddChildren(u ...*User) *UserCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddChildIDs(ids...)
}

// AddPetIDs adds the pets edge to Pet by ids.
func (uc *UserCreate) AddPetIDs(ids ...string) *UserCreate {
	uc.mutation.AddPetIDs(ids...)
	return uc
}

// AddPets adds the pets edges to Pet.
func (uc *UserCreate) AddPets(p ...*Pet) *UserCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddPetIDs(ids...)
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uc.hooks) == 0 {
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uc.mutation = mutation
			node, err = uc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	var (
		u     = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if id, ok := uc.mutation.ID(); ok {
		u.ID = id
		_spec.ID.Value = id
	}
	if nodes := uc.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.ParentTable,
			Columns: []string{user.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ChildrenTable,
			Columns: []string{user.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: pet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if u.ID == 0 {
		id := _spec.ID.Value.(int64)
		u.ID = int(id)
	}
	return u, nil
}
