package query

const QAddUser = `INSERT INTO users VALUES($1,$2,$3,$4,$5,$6,$7)`
const QGetUser = `SELECT * FROM users WHERE email=$1 AND password=$2`
const QDeleteUser = `DELETE FROM users WHERE id=:id`
