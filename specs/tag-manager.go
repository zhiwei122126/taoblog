package main

import (
	"database/sql"
	"log"
	"strings"
)

type xTagObject struct {
	id    int64
	name  string
	alias int64
}

type xTagManager struct {
	db *sql.DB
}

func newTagManager(db *sql.DB) *xTagManager {
	return &xTagManager{
		db: db,
	}
}

func (tm *xTagManager) addTag(name string, alias uint) int64 {
	sql := `INSERT INTO tags (name,alias) values (?,?)`
	stmt, err := tm.db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	ret, err := stmt.Exec(name, alias)
	if err != nil {
		panic(err)
	}

	id, err := ret.LastInsertId()
	return id
}

func (tm *xTagManager) searchTag(tag string) (tags []xTagObject) {
	sql := `SELECT * FROM tags WHERE name LIKE ?`
	rows, err := tm.db.Query(sql, "%"+tag+"%")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var tag xTagObject
		if err = rows.Scan(&tag.id, &tag.name, &tag.alias); err != nil {
			panic(err)
		}
		tags = append(tags, tag)
	}

	return
}

func (tm *xTagManager) getTagID(name string) int64 {
	sql := `SELECT id FROM tags WHERE name=? LIMIT 1`
	stmt, err := tm.db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	row := stmt.QueryRow(name)

	var id int64
	if err = row.Scan(&id); err != nil {
		log.Printf("标签名不存在：%s\n", name)
		id = 0
	}

	return id
}

func (tm *xTagManager) hasTagName(name string) bool {
	return tm.getTagID(name) > 0
}

func (tm *xTagManager) getTagNames(pid int64) (names []string) {
	sql := `SELECT tags.name FROM post_tags,tags WHERE post_tags.post_id=? AND post_tags.tag_id=tags.id`
	stmt, err := tm.db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query(pid)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	names = make([]string, 0)

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			panic(err)
		}
		names = append(names, name)
	}

	return
}

func (tm *xTagManager) getTagIDs(pid int64, alias bool) (ids []int64) {
	sql := `SELECT tag_id FROM post_tags WHERE post_id=?`
	stmt, err := tm.db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query(pid)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int64
		err = rows.Scan(&id)
		if err != nil {
			panic(err)
		}
		ids = append(ids, id)
	}

	if alias {
		ids = tm.getAliasTagsAll(ids)
	}

	return
}

func (tm *xTagManager) getAliasTagsAll(ids []int64) []int64 {
	sids := joinInts(ids, ",")

	sql1 := `SELECT alias FROM tags WHERE id in (?)`
	sql2 := `SELECT id FROM tags WHERE alias in (?)`

	rows, err := tm.db.Query(sql1, sids)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var alias int64
		if err = rows.Scan(&alias); err != nil {
			panic(err)
		}

		if alias > 0 {
			ids = append(ids, alias)
		}
	}

	rows.Close()

	rows, err = tm.db.Query(sql2, sids)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int64
		if err = rows.Scan(&id); err != nil {
			panic(err)
		}

		ids = append(ids, id)
	}

	rows.Close()

	return ids
}

func (tm *xTagManager) addObjectTag(pid int64, tid int64) int64 {
	sql := `INSERT INTO post_tags (post_id,tag_id) VALUES (?,?)`
	ret, err := tm.db.Exec(sql, pid, tid)
	if err != nil {
		panic(err)
	}

	id, err := ret.LastInsertId()

	return id
}

func (tm *xTagManager) removeObjectTag(pid, tid int64) {
	sql := `DELETE FROM post_tags WHERE post_id=? AND tag_id=? LIMIT 1`
	ret, err := tm.db.Exec(sql, pid, tid)
	if err != nil {
		panic(err)
	}
	_ = ret
}

func (tm *xTagManager) updateObjectTags(pid int64, tagstr string) {
	newTags := strings.Split(tagstr, ",")
	oldTags := tm.getTagNames(pid)

	var (
		toBeDeled []string
		toBeAdded []string
	)

	for _, t := range oldTags {
		if !strInSlice(newTags, t) {
			toBeDeled = append(toBeDeled, t)
		}
	}

	for _, t := range newTags {
		t = strings.TrimSpace(t)
		if t != "" && !strInSlice(oldTags, t) {
			toBeAdded = append(toBeAdded, t)
		}
	}

	for _, t := range toBeDeled {
		tid := tm.getTagID(t)
		tm.removeObjectTag(pid, tid)
	}

	for _, t := range toBeAdded {
		var tid int64
		if !tm.hasTagName(t) {
			tid = tm.addTag(t, 0)
		} else {
			tid = tm.getTagID(t)
		}
		tm.addObjectTag(pid, tid)
	}
}
