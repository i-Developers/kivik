package couchdb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/flimzy/kivik"
	"github.com/flimzy/kivik/driver"
	"github.com/flimzy/kivik/driver/couchdb/chttp"
	"github.com/flimzy/kivik/errors"
)

// deJSONify unmarshals a string, []byte, or json.RawMessage. All other types
// are returned as-is.
func deJSONify(i interface{}) (interface{}, error) {
	var data []byte
	switch t := i.(type) {
	case string:
		data = []byte(t)
	case []byte:
		data = t
	case json.RawMessage:
		data = []byte(t)
	default:
		return i, nil
	}
	var x interface{}
	err := json.Unmarshal(data, &x)
	return x, errors.WrapStatus(kivik.StatusBadRequest, err)
}

func (d *db) CreateIndexContext(ctx context.Context, ddoc, name string, index interface{}) error {
	if d.client.Compat == CompatCouch16 {
		return kivik.ErrNotImplemented
	}
	indexObj, err := deJSONify(index)
	if err != nil {
		return err
	}
	parameters := struct {
		Index interface{} `json:"index"`
		Ddoc  string      `json:"ddoc,omitempty"`
		Name  string      `json:"name,omitempty"`
	}{
		Index: indexObj,
		Ddoc:  ddoc,
		Name:  name,
	}
	body := &bytes.Buffer{}
	if err = json.NewEncoder(body).Encode(parameters); err != nil {
		return errors.WrapStatus(kivik.StatusBadRequest, err)
	}
	_, err = d.Client.DoError(ctx, kivik.MethodPost, d.path("_index", nil), &chttp.Options{Body: body})
	return err
}

func (d *db) DeleteIndexContext(ctx context.Context, ddoc, name string) error {
	path := fmt.Sprintf("_index/%s/json/%s", ddoc, name)
	_, err := d.Client.DoError(ctx, kivik.MethodDelete, d.path(path, nil), nil)
	return err
}

func (d *db) FindContext(ctx context.Context, query interface{}) (driver.Rows, error) {
	if d.client.Compat == CompatCouch16 {
		return nil, kivik.ErrNotImplemented
	}
	body, err := jsonify(query)
	if err != nil {
		return nil, err
	}
	resp, err := d.Client.DoReq(ctx, kivik.MethodPost, d.path("_find", nil), &chttp.Options{Body: body})
	if err != nil {
		return nil, err
	}
	if err = chttp.ResponseError(resp); err != nil {
		return nil, err
	}
	return newRows(resp.Body), nil
}
