rm feestboek.db
sqlite3 feestboek.db < schema.sql
go install
feestboek-back -dev
