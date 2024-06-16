package domain

import (
	"fmt"
	"slices"
	"strings"
)

const (
	queryTypeNamed = "NAMED"

	equalOperator        = "="
	lessOperator         = "<"
	greaterOperator      = ">"
	lessEqualOperator    = "<="
	greaterEqualOperator = ">="
)

var supportedOperators = []string{equalOperator, lessOperator, greaterOperator, lessEqualOperator, greaterEqualOperator}

type QueryBuilder struct {
	queryType    string
	counter      int
	conditions   []string
	args         []interface{}
	namedArgsMap map[string]interface{}
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		counter:      1,
		conditions:   make([]string, 0),
		args:         make([]interface{}, 0),
		namedArgsMap: make(map[string]interface{}),
	}
}

func NewNamedQueryBuilder() *QueryBuilder {
	qb := NewQueryBuilder()
	qb.queryType = queryTypeNamed
	return qb
}

func (qb *QueryBuilder) AddEqualCondition(key string, value interface{}) {
	qb.addCondition(key, value, equalOperator)
}

func (qb *QueryBuilder) AddLessCondition(key string, value interface{}) {
	qb.addCondition(key, value, lessOperator)
}

func (qb *QueryBuilder) AddGreaterCondition(key string, value interface{}) {
	qb.addCondition(key, value, greaterOperator)
}

func (qb *QueryBuilder) AddLessEqualCondition(key string, value interface{}) {
	qb.addCondition(key, value, lessEqualOperator)
}

func (qb *QueryBuilder) AddGreaterEqualCondition(key string, value interface{}) {
	qb.addCondition(key, value, greaterEqualOperator)
}

func (qb *QueryBuilder) Build() (string, []interface{}) {
	query := strings.Join(qb.conditions, " AND ")
	return query, qb.args
}

func (qb *QueryBuilder) BuildNamedConditions() (string, map[string]interface{}) {
	query := strings.Join(qb.conditions, ", ")
	return query, qb.namedArgsMap
}

func (qb *QueryBuilder) addCondition(key string, value interface{}, operator string) {
	if !slices.Contains(supportedOperators, operator) {
		return
	}

	if key == "" || value == nil || value == "" {
		return
	}

	condition := qb.getCondition(key, operator)
	qb.conditions = append(qb.conditions, condition)
	qb.args = append(qb.args, value)
	qb.namedArgsMap[key] = value
	qb.counter++
}

func (qb *QueryBuilder) getCondition(key, operator string) string {
	if qb.queryType == queryTypeNamed {
		return fmt.Sprintf("%s %s :%s", key, operator, key)
	}

	return fmt.Sprintf("%s %s $%d", key, operator, qb.counter)
}
