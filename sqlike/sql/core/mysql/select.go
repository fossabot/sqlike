package mysql

import (
	"bitbucket.org/SianLoong/sqlike/sqlike/actions"
	sqlstmt "bitbucket.org/SianLoong/sqlike/sqlike/sql/stmt"
)

// Select :
func (ms *MySQL) Select(f *actions.FindActions) (stmt *sqlstmt.Statement, err error) {
	stmt = sqlstmt.NewStatement(ms)
	err = buildStatement(stmt, ms.parser, f)
	if err != nil {
		return
	}
	return
}

func buildStatement(stmt *sqlstmt.Statement, parser *sqlstmt.StatementParser, f interface{}) error {
	if err := parser.BuildStatement(stmt, f); err != nil {
		return err
	}
	stmt.WriteRune(';')
	return nil
}
