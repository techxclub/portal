package domain

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
)

const (
	conditionGet   = "GET"
	conditionSet   = "SET"
	conditionWHERE = "WHERE"

	separatorAnd   = " AND "
	separatorComma = ", "

	equalOperator        = "="
	lessOperator         = "<"
	greaterOperator      = ">"
	lessEqualOperator    = "<="
	greaterEqualOperator = ">="
)

var supportedOperators = []string{equalOperator, lessOperator, greaterOperator, lessEqualOperator, greaterEqualOperator}

type QueryBuilder struct {
	separator       string
	conditionType   string
	counter         int
	conditions      []string
	namedConditions []string
	args            []interface{}
	namedArgsMap    map[string]interface{}
}

func newQueryBuilder(conditionType, separator string) *QueryBuilder {
	return &QueryBuilder{
		counter:       1,
		conditionType: conditionType,
		separator:     separator,
		conditions:    make([]string, 0),
		args:          make([]interface{}, 0),
		namedArgsMap:  make(map[string]interface{}),
	}
}

func NewGetQueryBuilder() *QueryBuilder {
	return newQueryBuilder(conditionGet, separatorAnd)
}

func NewSetQueryBuilder() *QueryBuilder {
	return newQueryBuilder(conditionSet, separatorComma)
}

func NewWhereQueryBuilder() *QueryBuilder {
	return newQueryBuilder(conditionWHERE, separatorAnd)
}

func (qb *QueryBuilder) AddEqualCondition(key string, value interface{}) {
	qb.addCondition(key, value, equalOperator)
}

func (qb *QueryBuilder) AddEqualConditionForJSONB(field, key string, value interface{}) {
	qb.addJSONBCondition(key, field, equalOperator, value)
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
	query := strings.Join(qb.conditions, qb.separator)
	return query, qb.args
}

func (qb *QueryBuilder) BuildNamed() (string, map[string]interface{}) {
	query := strings.Join(qb.namedConditions, qb.separator)
	return query, qb.namedArgsMap
}

func (qb *QueryBuilder) addCondition(key string, value interface{}, operator string) {
	if !slices.Contains(supportedOperators, operator) {
		return
	}

	if key == "" || reflect.ValueOf(value).IsZero() {
		return
	}

	condition := qb.getCondition(key, operator)
	qb.conditions = append(qb.conditions, condition)
	qb.args = append(qb.args, value)

	namedCondition := qb.getNamedCondition(key, operator)
	qb.namedConditions = append(qb.namedConditions, namedCondition)
	namedKey := qb.getNamedConditionKey(key)
	qb.namedArgsMap[namedKey] = value

	qb.counter++
}

func (qb *QueryBuilder) addJSONBCondition(key, field, operator string, value interface{}) {
	if !slices.Contains(supportedOperators, operator) {
		return
	}

	if key == "" || reflect.ValueOf(value).IsZero() {
		return
	}

	condition := fmt.Sprintf("jsonb_extract_path_text(%s, '%s') %s $%d", key, field, operator, qb.counter)
	qb.conditions = append(qb.conditions, condition)
	qb.args = append(qb.args, value)

	namedCondition := fmt.Sprintf("jsonb_extract_path_text(%s, '%s') %s :%s", key, field, operator, key)
	qb.namedConditions = append(qb.namedConditions, namedCondition)
	namedKey := qb.getNamedConditionKey(key)
	qb.namedArgsMap[namedKey] = value

	qb.counter++
}

func (qb *QueryBuilder) getCondition(key, operator string) string {
	return fmt.Sprintf("%s %s $%d", key, operator, qb.counter)
}

func (qb *QueryBuilder) getNamedCondition(key, operator string) string {
	return fmt.Sprintf("%s %s :%s", key, operator, qb.getNamedConditionKey(key))
}

func (qb *QueryBuilder) getNamedConditionKey(key string) string {
	return fmt.Sprintf("%s_%s", qb.conditionType, key)
}
