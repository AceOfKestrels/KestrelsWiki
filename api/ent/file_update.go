// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/ent/file"
	"api/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FileUpdate is the builder for updating File entities.
type FileUpdate struct {
	config
	hooks    []Hook
	mutation *FileMutation
}

// Where appends a list predicates to the FileUpdate builder.
func (fu *FileUpdate) Where(ps ...predicate.File) *FileUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetTitle sets the "title" field.
func (fu *FileUpdate) SetTitle(s string) *FileUpdate {
	fu.mutation.SetTitle(s)
	return fu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (fu *FileUpdate) SetNillableTitle(s *string) *FileUpdate {
	if s != nil {
		fu.SetTitle(*s)
	}
	return fu
}

// SetUpdated sets the "updated" field.
func (fu *FileUpdate) SetUpdated(t time.Time) *FileUpdate {
	fu.mutation.SetUpdated(t)
	return fu
}

// SetNillableUpdated sets the "updated" field if the given value is not nil.
func (fu *FileUpdate) SetNillableUpdated(t *time.Time) *FileUpdate {
	if t != nil {
		fu.SetUpdated(*t)
	}
	return fu
}

// SetAuthor sets the "author" field.
func (fu *FileUpdate) SetAuthor(s string) *FileUpdate {
	fu.mutation.SetAuthor(s)
	return fu
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (fu *FileUpdate) SetNillableAuthor(s *string) *FileUpdate {
	if s != nil {
		fu.SetAuthor(*s)
	}
	return fu
}

// SetCommitHash sets the "commitHash" field.
func (fu *FileUpdate) SetCommitHash(s string) *FileUpdate {
	fu.mutation.SetCommitHash(s)
	return fu
}

// SetNillableCommitHash sets the "commitHash" field if the given value is not nil.
func (fu *FileUpdate) SetNillableCommitHash(s *string) *FileUpdate {
	if s != nil {
		fu.SetCommitHash(*s)
	}
	return fu
}

// SetContent sets the "content" field.
func (fu *FileUpdate) SetContent(s string) *FileUpdate {
	fu.mutation.SetContent(s)
	return fu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (fu *FileUpdate) SetNillableContent(s *string) *FileUpdate {
	if s != nil {
		fu.SetContent(*s)
	}
	return fu
}

// Mutation returns the FileMutation object of the builder.
func (fu *FileUpdate) Mutation() *FileMutation {
	return fu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FileUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FileUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FileUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FileUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(file.Table, file.Columns, sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Title(); ok {
		_spec.SetField(file.FieldTitle, field.TypeString, value)
	}
	if value, ok := fu.mutation.Updated(); ok {
		_spec.SetField(file.FieldUpdated, field.TypeTime, value)
	}
	if value, ok := fu.mutation.Author(); ok {
		_spec.SetField(file.FieldAuthor, field.TypeString, value)
	}
	if value, ok := fu.mutation.CommitHash(); ok {
		_spec.SetField(file.FieldCommitHash, field.TypeString, value)
	}
	if value, ok := fu.mutation.Content(); ok {
		_spec.SetField(file.FieldContent, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FileUpdateOne is the builder for updating a single File entity.
type FileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FileMutation
}

// SetTitle sets the "title" field.
func (fuo *FileUpdateOne) SetTitle(s string) *FileUpdateOne {
	fuo.mutation.SetTitle(s)
	return fuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableTitle(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetTitle(*s)
	}
	return fuo
}

// SetUpdated sets the "updated" field.
func (fuo *FileUpdateOne) SetUpdated(t time.Time) *FileUpdateOne {
	fuo.mutation.SetUpdated(t)
	return fuo
}

// SetNillableUpdated sets the "updated" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableUpdated(t *time.Time) *FileUpdateOne {
	if t != nil {
		fuo.SetUpdated(*t)
	}
	return fuo
}

// SetAuthor sets the "author" field.
func (fuo *FileUpdateOne) SetAuthor(s string) *FileUpdateOne {
	fuo.mutation.SetAuthor(s)
	return fuo
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableAuthor(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetAuthor(*s)
	}
	return fuo
}

// SetCommitHash sets the "commitHash" field.
func (fuo *FileUpdateOne) SetCommitHash(s string) *FileUpdateOne {
	fuo.mutation.SetCommitHash(s)
	return fuo
}

// SetNillableCommitHash sets the "commitHash" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableCommitHash(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetCommitHash(*s)
	}
	return fuo
}

// SetContent sets the "content" field.
func (fuo *FileUpdateOne) SetContent(s string) *FileUpdateOne {
	fuo.mutation.SetContent(s)
	return fuo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableContent(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetContent(*s)
	}
	return fuo
}

// Mutation returns the FileMutation object of the builder.
func (fuo *FileUpdateOne) Mutation() *FileMutation {
	return fuo.mutation
}

// Where appends a list predicates to the FileUpdate builder.
func (fuo *FileUpdateOne) Where(ps ...predicate.File) *FileUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FileUpdateOne) Select(field string, fields ...string) *FileUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated File entity.
func (fuo *FileUpdateOne) Save(ctx context.Context) (*File, error) {
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FileUpdateOne) SaveX(ctx context.Context) *File {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FileUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FileUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FileUpdateOne) sqlSave(ctx context.Context) (_node *File, err error) {
	_spec := sqlgraph.NewUpdateSpec(file.Table, file.Columns, sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "File.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, file.FieldID)
		for _, f := range fields {
			if !file.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != file.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.Title(); ok {
		_spec.SetField(file.FieldTitle, field.TypeString, value)
	}
	if value, ok := fuo.mutation.Updated(); ok {
		_spec.SetField(file.FieldUpdated, field.TypeTime, value)
	}
	if value, ok := fuo.mutation.Author(); ok {
		_spec.SetField(file.FieldAuthor, field.TypeString, value)
	}
	if value, ok := fuo.mutation.CommitHash(); ok {
		_spec.SetField(file.FieldCommitHash, field.TypeString, value)
	}
	if value, ok := fuo.mutation.Content(); ok {
		_spec.SetField(file.FieldContent, field.TypeString, value)
	}
	_node = &File{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}
