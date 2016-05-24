// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sql

func (p *Parser) parseDefineFieldStatement(explain bool) (stmt *DefineFieldStatement, err error) {

	stmt = &DefineFieldStatement{}

	stmt.EX = explain

	stmt.KV = p.c.Get("KV").(string)
	stmt.NS = p.c.Get("NS").(string)
	stmt.DB = p.c.Get("DB").(string)

	if stmt.Name, err = p.parseIdent(); err != nil {
		return nil, err
	}

	if _, _, err = p.shouldBe(ON); err != nil {
		return nil, err
	}

	if stmt.What, err = p.parseTables(); err != nil {
		return nil, err
	}

	for {

		tok, _, exi := p.mightBe(MIN, MAX, TYPE, CODE, DEFAULT, NOTNULL, READONLY, MANDATORY)
		if !exi {
			break
		}

		if is(tok, MIN) {
			if stmt.Min, err = p.parseNumber(); err != nil {
				return nil, err
			}
		}

		if is(tok, MAX) {
			if stmt.Max, err = p.parseNumber(); err != nil {
				return nil, err
			}
		}

		if is(tok, TYPE) {
			if stmt.Type, err = p.parseType(); err != nil {
				return nil, err
			}
		}

		if is(tok, CODE) {
			if stmt.Code, err = p.parseCode(); err != nil {
				return nil, err
			}
		}

		if is(tok, DEFAULT) {
			if stmt.Default, err = p.parseDefault(); err != nil {
				return nil, err
			}
		}

		if is(tok, NOTNULL) {
			stmt.Notnull = true
			if tok, _, exi := p.mightBe(TRUE, FALSE); exi {
				if tok == FALSE {
					stmt.Notnull = false
				}
			}
		}

		if is(tok, READONLY) {
			stmt.Readonly = true
			if tok, _, exi := p.mightBe(TRUE, FALSE); exi {
				if tok == FALSE {
					stmt.Readonly = false
				}
			}
		}

		if is(tok, MANDATORY) {
			stmt.Mandatory = true
			if tok, _, exi := p.mightBe(TRUE, FALSE); exi {
				if tok == FALSE {
					stmt.Mandatory = false
				}
			}
		}

	}

	if stmt.Type == nil {
		return nil, &ParseError{Found: "", Expected: []string{"TYPE"}}
	}

	if _, _, err = p.shouldBe(EOF, SEMICOLON); err != nil {
		return nil, err
	}

	return

}

func (p *Parser) parseRemoveFieldStatement(explain bool) (stmt *RemoveFieldStatement, err error) {

	stmt = &RemoveFieldStatement{}

	stmt.EX = explain

	stmt.KV = p.c.Get("KV").(string)
	stmt.NS = p.c.Get("NS").(string)
	stmt.DB = p.c.Get("DB").(string)

	if stmt.Name, err = p.parseIdent(); err != nil {
		return nil, err
	}

	if _, _, err = p.shouldBe(ON); err != nil {
		return nil, err
	}

	if stmt.What, err = p.parseTables(); err != nil {
		return nil, err
	}

	if _, _, err = p.shouldBe(EOF, SEMICOLON); err != nil {
		return nil, err
	}

	return

}