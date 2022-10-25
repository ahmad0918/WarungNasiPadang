package utils

const (
	// Menu
	INSERT_MENU = `insert into menu values($1,$2,$3)`
	SELECT_MENU_LIST = `select * from menu`
	DELETE_MENU = `delete from menu where id = $1`
	UPDATE_MENU = `update menu set name=$1, price=$2 where id =$3`

	// Transaksa
	INSERT_TRANSAKSI = `insert into transaksi values($1,$2,$3,$4)`
	SELECT_TRANSAKSI_LIST = `select * from transaksi`
	DELETE_TRANSAKSI = `delete from transaksi where id = $1`
	UPDATE_TRANSAKSI = `update transaksi set menu=$1, quantity=$2, date=$3 where id =$4`
)