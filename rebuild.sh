rm feestboek.db
sqlite3 feestboek.db < schema.sql

killall -q -9 feestboek-back
go install
feestboek-back -dev &
sleep 1
