// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/zibbp/ganymede/ent/multistreaminfo"
	"github.com/zibbp/ganymede/ent/playlist"
	"github.com/zibbp/ganymede/ent/vod"
)

// MultistreamInfoCreate is the builder for creating a MultistreamInfo entity.
type MultistreamInfoCreate struct {
	config
	mutation *MultistreamInfoMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDelayMs sets the "delay_ms" field.
func (mic *MultistreamInfoCreate) SetDelayMs(i int) *MultistreamInfoCreate {
	mic.mutation.SetDelayMs(i)
	return mic
}

// SetNillableDelayMs sets the "delay_ms" field if the given value is not nil.
func (mic *MultistreamInfoCreate) SetNillableDelayMs(i *int) *MultistreamInfoCreate {
	if i != nil {
		mic.SetDelayMs(*i)
	}
	return mic
}

// SetVodID sets the "vod" edge to the Vod entity by ID.
func (mic *MultistreamInfoCreate) SetVodID(id uuid.UUID) *MultistreamInfoCreate {
	mic.mutation.SetVodID(id)
	return mic
}

// SetVod sets the "vod" edge to the Vod entity.
func (mic *MultistreamInfoCreate) SetVod(v *Vod) *MultistreamInfoCreate {
	return mic.SetVodID(v.ID)
}

// SetPlaylistID sets the "playlist" edge to the Playlist entity by ID.
func (mic *MultistreamInfoCreate) SetPlaylistID(id uuid.UUID) *MultistreamInfoCreate {
	mic.mutation.SetPlaylistID(id)
	return mic
}

// SetPlaylist sets the "playlist" edge to the Playlist entity.
func (mic *MultistreamInfoCreate) SetPlaylist(p *Playlist) *MultistreamInfoCreate {
	return mic.SetPlaylistID(p.ID)
}

// Mutation returns the MultistreamInfoMutation object of the builder.
func (mic *MultistreamInfoCreate) Mutation() *MultistreamInfoMutation {
	return mic.mutation
}

// Save creates the MultistreamInfo in the database.
func (mic *MultistreamInfoCreate) Save(ctx context.Context) (*MultistreamInfo, error) {
	return withHooks(ctx, mic.sqlSave, mic.mutation, mic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mic *MultistreamInfoCreate) SaveX(ctx context.Context) *MultistreamInfo {
	v, err := mic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mic *MultistreamInfoCreate) Exec(ctx context.Context) error {
	_, err := mic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mic *MultistreamInfoCreate) ExecX(ctx context.Context) {
	if err := mic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mic *MultistreamInfoCreate) check() error {
	if _, ok := mic.mutation.VodID(); !ok {
		return &ValidationError{Name: "vod", err: errors.New(`ent: missing required edge "MultistreamInfo.vod"`)}
	}
	if _, ok := mic.mutation.PlaylistID(); !ok {
		return &ValidationError{Name: "playlist", err: errors.New(`ent: missing required edge "MultistreamInfo.playlist"`)}
	}
	return nil
}

func (mic *MultistreamInfoCreate) sqlSave(ctx context.Context) (*MultistreamInfo, error) {
	if err := mic.check(); err != nil {
		return nil, err
	}
	_node, _spec := mic.createSpec()
	if err := sqlgraph.CreateNode(ctx, mic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mic.mutation.id = &_node.ID
	mic.mutation.done = true
	return _node, nil
}

func (mic *MultistreamInfoCreate) createSpec() (*MultistreamInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &MultistreamInfo{config: mic.config}
		_spec = sqlgraph.NewCreateSpec(multistreaminfo.Table, sqlgraph.NewFieldSpec(multistreaminfo.FieldID, field.TypeInt))
	)
	_spec.OnConflict = mic.conflict
	if value, ok := mic.mutation.DelayMs(); ok {
		_spec.SetField(multistreaminfo.FieldDelayMs, field.TypeInt, value)
		_node.DelayMs = value
	}
	if nodes := mic.mutation.VodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   multistreaminfo.VodTable,
			Columns: []string{multistreaminfo.VodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vod.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.multistream_info_vod = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mic.mutation.PlaylistIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   multistreaminfo.PlaylistTable,
			Columns: []string{multistreaminfo.PlaylistColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playlist.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.playlist_multistream_info = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MultistreamInfo.Create().
//		SetDelayMs(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MultistreamInfoUpsert) {
//			SetDelayMs(v+v).
//		}).
//		Exec(ctx)
func (mic *MultistreamInfoCreate) OnConflict(opts ...sql.ConflictOption) *MultistreamInfoUpsertOne {
	mic.conflict = opts
	return &MultistreamInfoUpsertOne{
		create: mic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MultistreamInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (mic *MultistreamInfoCreate) OnConflictColumns(columns ...string) *MultistreamInfoUpsertOne {
	mic.conflict = append(mic.conflict, sql.ConflictColumns(columns...))
	return &MultistreamInfoUpsertOne{
		create: mic,
	}
}

type (
	// MultistreamInfoUpsertOne is the builder for "upsert"-ing
	//  one MultistreamInfo node.
	MultistreamInfoUpsertOne struct {
		create *MultistreamInfoCreate
	}

	// MultistreamInfoUpsert is the "OnConflict" setter.
	MultistreamInfoUpsert struct {
		*sql.UpdateSet
	}
)

// SetDelayMs sets the "delay_ms" field.
func (u *MultistreamInfoUpsert) SetDelayMs(v int) *MultistreamInfoUpsert {
	u.Set(multistreaminfo.FieldDelayMs, v)
	return u
}

// UpdateDelayMs sets the "delay_ms" field to the value that was provided on create.
func (u *MultistreamInfoUpsert) UpdateDelayMs() *MultistreamInfoUpsert {
	u.SetExcluded(multistreaminfo.FieldDelayMs)
	return u
}

// AddDelayMs adds v to the "delay_ms" field.
func (u *MultistreamInfoUpsert) AddDelayMs(v int) *MultistreamInfoUpsert {
	u.Add(multistreaminfo.FieldDelayMs, v)
	return u
}

// ClearDelayMs clears the value of the "delay_ms" field.
func (u *MultistreamInfoUpsert) ClearDelayMs() *MultistreamInfoUpsert {
	u.SetNull(multistreaminfo.FieldDelayMs)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.MultistreamInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *MultistreamInfoUpsertOne) UpdateNewValues() *MultistreamInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MultistreamInfo.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *MultistreamInfoUpsertOne) Ignore() *MultistreamInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MultistreamInfoUpsertOne) DoNothing() *MultistreamInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MultistreamInfoCreate.OnConflict
// documentation for more info.
func (u *MultistreamInfoUpsertOne) Update(set func(*MultistreamInfoUpsert)) *MultistreamInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MultistreamInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetDelayMs sets the "delay_ms" field.
func (u *MultistreamInfoUpsertOne) SetDelayMs(v int) *MultistreamInfoUpsertOne {
	return u.Update(func(s *MultistreamInfoUpsert) {
		s.SetDelayMs(v)
	})
}

// AddDelayMs adds v to the "delay_ms" field.
func (u *MultistreamInfoUpsertOne) AddDelayMs(v int) *MultistreamInfoUpsertOne {
	return u.Update(func(s *MultistreamInfoUpsert) {
		s.AddDelayMs(v)
	})
}

// UpdateDelayMs sets the "delay_ms" field to the value that was provided on create.
func (u *MultistreamInfoUpsertOne) UpdateDelayMs() *MultistreamInfoUpsertOne {
	return u.Update(func(s *MultistreamInfoUpsert) {
		s.UpdateDelayMs()
	})
}

// ClearDelayMs clears the value of the "delay_ms" field.
func (u *MultistreamInfoUpsertOne) ClearDelayMs() *MultistreamInfoUpsertOne {
	return u.Update(func(s *MultistreamInfoUpsert) {
		s.ClearDelayMs()
	})
}

// Exec executes the query.
func (u *MultistreamInfoUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MultistreamInfoCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MultistreamInfoUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *MultistreamInfoUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *MultistreamInfoUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// MultistreamInfoCreateBulk is the builder for creating many MultistreamInfo entities in bulk.
type MultistreamInfoCreateBulk struct {
	config
	err      error
	builders []*MultistreamInfoCreate
	conflict []sql.ConflictOption
}

// Save creates the MultistreamInfo entities in the database.
func (micb *MultistreamInfoCreateBulk) Save(ctx context.Context) ([]*MultistreamInfo, error) {
	if micb.err != nil {
		return nil, micb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(micb.builders))
	nodes := make([]*MultistreamInfo, len(micb.builders))
	mutators := make([]Mutator, len(micb.builders))
	for i := range micb.builders {
		func(i int, root context.Context) {
			builder := micb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MultistreamInfoMutation)
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
					_, err = mutators[i+1].Mutate(root, micb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = micb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, micb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, micb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (micb *MultistreamInfoCreateBulk) SaveX(ctx context.Context) []*MultistreamInfo {
	v, err := micb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (micb *MultistreamInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := micb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (micb *MultistreamInfoCreateBulk) ExecX(ctx context.Context) {
	if err := micb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.MultistreamInfo.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.MultistreamInfoUpsert) {
//			SetDelayMs(v+v).
//		}).
//		Exec(ctx)
func (micb *MultistreamInfoCreateBulk) OnConflict(opts ...sql.ConflictOption) *MultistreamInfoUpsertBulk {
	micb.conflict = opts
	return &MultistreamInfoUpsertBulk{
		create: micb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.MultistreamInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (micb *MultistreamInfoCreateBulk) OnConflictColumns(columns ...string) *MultistreamInfoUpsertBulk {
	micb.conflict = append(micb.conflict, sql.ConflictColumns(columns...))
	return &MultistreamInfoUpsertBulk{
		create: micb,
	}
}

// MultistreamInfoUpsertBulk is the builder for "upsert"-ing
// a bulk of MultistreamInfo nodes.
type MultistreamInfoUpsertBulk struct {
	create *MultistreamInfoCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.MultistreamInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *MultistreamInfoUpsertBulk) UpdateNewValues() *MultistreamInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.MultistreamInfo.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *MultistreamInfoUpsertBulk) Ignore() *MultistreamInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *MultistreamInfoUpsertBulk) DoNothing() *MultistreamInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the MultistreamInfoCreateBulk.OnConflict
// documentation for more info.
func (u *MultistreamInfoUpsertBulk) Update(set func(*MultistreamInfoUpsert)) *MultistreamInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&MultistreamInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetDelayMs sets the "delay_ms" field.
func (u *MultistreamInfoUpsertBulk) SetDelayMs(v int) *MultistreamInfoUpsertBulk {
	return u.Update(func(s *MultistreamInfoUpsert) {
		s.SetDelayMs(v)
	})
}

// AddDelayMs adds v to the "delay_ms" field.
func (u *MultistreamInfoUpsertBulk) AddDelayMs(v int) *MultistreamInfoUpsertBulk {
	return u.Update(func(s *MultistreamInfoUpsert) {
		s.AddDelayMs(v)
	})
}

// UpdateDelayMs sets the "delay_ms" field to the value that was provided on create.
func (u *MultistreamInfoUpsertBulk) UpdateDelayMs() *MultistreamInfoUpsertBulk {
	return u.Update(func(s *MultistreamInfoUpsert) {
		s.UpdateDelayMs()
	})
}

// ClearDelayMs clears the value of the "delay_ms" field.
func (u *MultistreamInfoUpsertBulk) ClearDelayMs() *MultistreamInfoUpsertBulk {
	return u.Update(func(s *MultistreamInfoUpsert) {
		s.ClearDelayMs()
	})
}

// Exec executes the query.
func (u *MultistreamInfoUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the MultistreamInfoCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for MultistreamInfoCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *MultistreamInfoUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}