DB_FILE	:= sg.db
SQLITE	:= sqlite3

.PHONY: db
db: $(DB_FILE)

$(DB_FILE):
	$(SQLITE) $@ < configs/init.sql
	$(SQLITE) $@ < configs/test-data.sql

.PHONY: clean
clean:
	rm -f $(DB_FILE)

.PHONY: schema
schema:
	sqlboiler sqlite3
