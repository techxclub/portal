package domain

import (
	"fmt"
	"slices"
	"strings"
)

const (
	equalOperator        = "="
	lessOperator         = "<"
	greaterOperator      = ">"
	lessEqualOperator    = "<="
	greaterEqualOperator = ">="
)

var supportedOperators = []string{equalOperator, lessOperator, greaterOperator, lessEqualOperator, greaterEqualOperator}

type QueryBuilder struct {
	counter    int
	conditions []string
	args       []interface{}
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		counter:    1,
		conditions: make([]string, 0),
		args:       make([]interface{}, 0),
	}
}

func (qb *QueryBuilder) AddEqualParam(key string, value interface{}) {
	qb.AddParam(key, value, equalOperator)
}

func (qb *QueryBuilder) AddLessParam(key string, value interface{}) {
	qb.AddParam(key, value, lessOperator)
}

func (qb *QueryBuilder) AddGreaterParam(key string, value interface{}) {
	qb.AddParam(key, value, greaterOperator)
}

func (qb *QueryBuilder) AddLessEqualParam(key string, value interface{}) {
	qb.AddParam(key, value, lessEqualOperator)
}

func (qb *QueryBuilder) AddGreaterEqualParam(key string, value interface{}) {
	qb.AddParam(key, value, greaterEqualOperator)
}

func (qb *QueryBuilder) AddParam(key string, value interface{}, operator string) {
	if !slices.Contains(supportedOperators, operator) {
		return
	}

	if key == "" || value == nil || value == "" {
		return
	}

	condition := fmt.Sprintf("%s %s $%d", key, operator, qb.counter)
	qb.conditions = append(qb.conditions, condition)
	qb.args = append(qb.args, value)
	qb.counter++
}

func (qb *QueryBuilder) Build() (string, []interface{}) {
	query := strings.Join(qb.conditions, " AND ")
	return query, qb.args
}
