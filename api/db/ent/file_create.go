// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/db/ent/file"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FileCreate is the builder for creating a File entity.
type FileCreate struct {
	config
	mutation *FileMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetPath sets the "path" field.
func (fc *FileCreate) SetPath(s string) *FileCreate {
	fc.mutation.SetPath(s)
	return fc
}

// SetTitle sets the "title" field.
func (fc *FileCreate) SetTitle(s string) *FileCreate {
	fc.mutation.SetTitle(s)
	return fc
}

// SetUpdated sets the "updated" field.
func (fc *FileCreate) SetUpdated(t time.Time) *FileCreate {
	fc.mutation.SetUpdated(t)
	return fc
}

// SetAuthor sets the "author" field.
func (fc *FileCreate) SetAuthor(s string) *FileCreate {
	fc.mutation.SetAuthor(s)
	return fc
}

// SetCommitHash sets the "commitHash" field.
func (fc *FileCreate) SetCommitHash(s string) *FileCreate {
	fc.mutation.SetCommitHash(s)
	return fc
}

// SetContent sets the "content" field.
func (fc *FileCreate) SetContent(s string) *FileCreate {
	fc.mutation.SetContent(s)
	return fc
}

// Mutation returns the FileMutation object of the builder.
func (fc *FileCreate) Mutation() *FileMutation {
	return fc.mutation
}

// Save creates the File in the database.
func (fc *FileCreate) Save(ctx context.Context) (*File, error) {
	return withHooks(ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FileCreate) SaveX(ctx context.Context) *File {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FileCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FileCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FileCreate) check() error {
	if _, ok := fc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "File.path"`)}
	}
	if v, ok := fc.mutation.Path(); ok {
		if err := file.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "File.path": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "File.title"`)}
	}
	if _, ok := fc.mutation.Updated(); !ok {
		return &ValidationError{Name: "updated", err: errors.New(`ent: missing required field "File.updated"`)}
	}
	if _, ok := fc.mutation.Author(); !ok {
		return &ValidationError{Name: "author", err: errors.New(`ent: missing required field "File.author"`)}
	}
	if _, ok := fc.mutation.CommitHash(); !ok {
		return &ValidationError{Name: "commitHash", err: errors.New(`ent: missing required field "File.commitHash"`)}
	}
	if _, ok := fc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "File.content"`)}
	}
	return nil
}

func (fc *FileCreate) sqlSave(ctx context.Context) (*File, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FileCreate) createSpec() (*File, *sqlgraph.CreateSpec) {
	var (
		_node = &File{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(file.Table, sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt))
	)
	_spec.OnConflict = fc.conflict
	if value, ok := fc.mutation.Path(); ok {
		_spec.SetField(file.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := fc.mutation.Title(); ok {
		_spec.SetField(file.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := fc.mutation.Updated(); ok {
		_spec.SetField(file.FieldUpdated, field.TypeTime, value)
		_node.Updated = value
	}
	if value, ok := fc.mutation.Author(); ok {
		_spec.SetField(file.FieldAuthor, field.TypeString, value)
		_node.Author = value
	}
	if value, ok := fc.mutation.CommitHash(); ok {
		_spec.SetField(file.FieldCommitHash, field.TypeString, value)
		_node.CommitHash = value
	}
	if value, ok := fc.mutation.Content(); ok {
		_spec.SetField(file.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.File.Create().
//		SetPath(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FileUpsert) {
//			SetPath(v+v).
//		}).
//		Exec(ctx)
func (fc *FileCreate) OnConflict(opts ...sql.ConflictOption) *FileUpsertOne {
	fc.conflict = opts
	return &FileUpsertOne{
		create: fc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fc *FileCreate) OnConflictColumns(columns ...string) *FileUpsertOne {
	fc.conflict = append(fc.conflict, sql.ConflictColumns(columns...))
	return &FileUpsertOne{
		create: fc,
	}
}

type (
	// FileUpsertOne is the builder for "upsert"-ing
	//  one File node.
	FileUpsertOne struct {
		create *FileCreate
	}

	// FileUpsert is the "OnConflict" setter.
	FileUpsert struct {
		*sql.UpdateSet
	}
)

// SetTitle sets the "title" field.
func (u *FileUpsert) SetTitle(v string) *FileUpsert {
	u.Set(file.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *FileUpsert) UpdateTitle() *FileUpsert {
	u.SetExcluded(file.FieldTitle)
	return u
}

// SetUpdated sets the "updated" field.
func (u *FileUpsert) SetUpdated(v time.Time) *FileUpsert {
	u.Set(file.FieldUpdated, v)
	return u
}

// UpdateUpdated sets the "updated" field to the value that was provided on create.
func (u *FileUpsert) UpdateUpdated() *FileUpsert {
	u.SetExcluded(file.FieldUpdated)
	return u
}

// SetAuthor sets the "author" field.
func (u *FileUpsert) SetAuthor(v string) *FileUpsert {
	u.Set(file.FieldAuthor, v)
	return u
}

// UpdateAuthor sets the "author" field to the value that was provided on create.
func (u *FileUpsert) UpdateAuthor() *FileUpsert {
	u.SetExcluded(file.FieldAuthor)
	return u
}

// SetCommitHash sets the "commitHash" field.
func (u *FileUpsert) SetCommitHash(v string) *FileUpsert {
	u.Set(file.FieldCommitHash, v)
	return u
}

// UpdateCommitHash sets the "commitHash" field to the value that was provided on create.
func (u *FileUpsert) UpdateCommitHash() *FileUpsert {
	u.SetExcluded(file.FieldCommitHash)
	return u
}

// SetContent sets the "content" field.
func (u *FileUpsert) SetContent(v string) *FileUpsert {
	u.Set(file.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *FileUpsert) UpdateContent() *FileUpsert {
	u.SetExcluded(file.FieldContent)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *FileUpsertOne) UpdateNewValues() *FileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.Path(); exists {
			s.SetIgnore(file.FieldPath)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.File.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *FileUpsertOne) Ignore() *FileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FileUpsertOne) DoNothing() *FileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FileCreate.OnConflict
// documentation for more info.
func (u *FileUpsertOne) Update(set func(*FileUpsert)) *FileUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FileUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *FileUpsertOne) SetTitle(v string) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateTitle() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateTitle()
	})
}

// SetUpdated sets the "updated" field.
func (u *FileUpsertOne) SetUpdated(v time.Time) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetUpdated(v)
	})
}

// UpdateUpdated sets the "updated" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateUpdated() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateUpdated()
	})
}

// SetAuthor sets the "author" field.
func (u *FileUpsertOne) SetAuthor(v string) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetAuthor(v)
	})
}

// UpdateAuthor sets the "author" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateAuthor() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateAuthor()
	})
}

// SetCommitHash sets the "commitHash" field.
func (u *FileUpsertOne) SetCommitHash(v string) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetCommitHash(v)
	})
}

// UpdateCommitHash sets the "commitHash" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateCommitHash() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateCommitHash()
	})
}

// SetContent sets the "content" field.
func (u *FileUpsertOne) SetContent(v string) *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *FileUpsertOne) UpdateContent() *FileUpsertOne {
	return u.Update(func(s *FileUpsert) {
		s.UpdateContent()
	})
}

// Exec executes the query.
func (u *FileUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FileCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FileUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *FileUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *FileUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// FileCreateBulk is the builder for creating many File entities in bulk.
type FileCreateBulk struct {
	config
	err      error
	builders []*FileCreate
	conflict []sql.ConflictOption
}

// Save creates the File entities in the database.
func (fcb *FileCreateBulk) Save(ctx context.Context) ([]*File, error) {
	if fcb.err != nil {
		return nil, fcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*File, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FileCreateBulk) SaveX(ctx context.Context) []*File {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FileCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FileCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.File.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.FileUpsert) {
//			SetPath(v+v).
//		}).
//		Exec(ctx)
func (fcb *FileCreateBulk) OnConflict(opts ...sql.ConflictOption) *FileUpsertBulk {
	fcb.conflict = opts
	return &FileUpsertBulk{
		create: fcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fcb *FileCreateBulk) OnConflictColumns(columns ...string) *FileUpsertBulk {
	fcb.conflict = append(fcb.conflict, sql.ConflictColumns(columns...))
	return &FileUpsertBulk{
		create: fcb,
	}
}

// FileUpsertBulk is the builder for "upsert"-ing
// a bulk of File nodes.
type FileUpsertBulk struct {
	create *FileCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *FileUpsertBulk) UpdateNewValues() *FileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.Path(); exists {
				s.SetIgnore(file.FieldPath)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.File.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *FileUpsertBulk) Ignore() *FileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *FileUpsertBulk) DoNothing() *FileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the FileCreateBulk.OnConflict
// documentation for more info.
func (u *FileUpsertBulk) Update(set func(*FileUpsert)) *FileUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&FileUpsert{UpdateSet: update})
	}))
	return u
}

// SetTitle sets the "title" field.
func (u *FileUpsertBulk) SetTitle(v string) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateTitle() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateTitle()
	})
}

// SetUpdated sets the "updated" field.
func (u *FileUpsertBulk) SetUpdated(v time.Time) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetUpdated(v)
	})
}

// UpdateUpdated sets the "updated" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateUpdated() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateUpdated()
	})
}

// SetAuthor sets the "author" field.
func (u *FileUpsertBulk) SetAuthor(v string) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetAuthor(v)
	})
}

// UpdateAuthor sets the "author" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateAuthor() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateAuthor()
	})
}

// SetCommitHash sets the "commitHash" field.
func (u *FileUpsertBulk) SetCommitHash(v string) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetCommitHash(v)
	})
}

// UpdateCommitHash sets the "commitHash" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateCommitHash() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateCommitHash()
	})
}

// SetContent sets the "content" field.
func (u *FileUpsertBulk) SetContent(v string) *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *FileUpsertBulk) UpdateContent() *FileUpsertBulk {
	return u.Update(func(s *FileUpsert) {
		s.UpdateContent()
	})
}

// Exec executes the query.
func (u *FileUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the FileCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for FileCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FileUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
