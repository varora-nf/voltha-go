/*
 * Copyright 2018-present Open Networking Foundation

 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 * http://www.apache.org/licenses/LICENSE-2.0

 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package model

import (
	"context"
	"fmt"

	"github.com/opencord/voltha-lib-go/v3/pkg/log"
)

// Transaction -
type Transaction struct {
	proxy *Proxy
	txid  string
}

// NewTransaction -
func NewTransaction(proxy *Proxy, txid string) *Transaction {
	tx := &Transaction{
		proxy: proxy,
		txid:  txid,
	}
	return tx
}
func (t *Transaction) Get(ctx context.Context, path string, depth int, deep bool) (interface{}, error) {
	if t.txid == "" {
		log.Errorf("closed transaction")
		return nil, fmt.Errorf("closed transaction")
	}
	// TODO: need to review the return values at the different layers!!!!!
	return t.proxy.Get(ctx, path, depth, deep, t.txid)
}
func (t *Transaction) Update(ctx context.Context, path string, data interface{}, strict bool) (interface{}, error) {
	if t.txid == "" {
		log.Errorf("closed transaction")
		return nil, fmt.Errorf("closed transaction")
	}
	return t.proxy.Update(ctx, path, data, strict, t.txid)
}
func (t *Transaction) Add(ctx context.Context, path string, data interface{}) (interface{}, error) {
	if t.txid == "" {
		log.Errorf("closed transaction")
		return nil, fmt.Errorf("closed transaction")
	}
	return t.proxy.Add(ctx, path, data, t.txid)
}
func (t *Transaction) Remove(ctx context.Context, path string) (interface{}, error) {
	if t.txid == "" {
		log.Errorf("closed transaction")
		return nil, fmt.Errorf("closed transaction")
	}
	return t.proxy.Remove(ctx, path, t.txid)
}

// Cancel -
func (t *Transaction) Cancel() {
	t.proxy.cancelTransaction(t.txid)
	t.txid = ""
}

// Commit -
func (t *Transaction) Commit(ctx context.Context) {
	t.proxy.commitTransaction(ctx, t.txid)
	t.txid = ""
}
