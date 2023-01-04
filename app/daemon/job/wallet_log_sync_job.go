package job

import (
	"context"
	"log"
	"time"

	"github.com/ixre/go2o/core/domain/interface/wallet"
	"github.com/ixre/go2o/core/infrastructure/locker"
	"github.com/ixre/go2o/core/repos"
	"github.com/ixre/go2o/core/repos/clickhouse"
)

func SyncWalletLogToClickHouse() {
	jobName := "SyncWalletLogToClickHouse"
	if !locker.Lock(jobName, 600) {
		return
	}
	defer locker.Unlock(jobName)
	repo := repos.Repo.GetWalletRepo()
	//log.Println("[ job]: start sync wallet log to clickhouse..")
	job := getJob(jobName)
	lastId := job.GetValue().LastExecIndex
	size := 1000

	for {
		list := repo.SelectWalletLog_("id > $1 ORDER BY id ASC LIMIT $2", lastId, size)
		l := len(list)
		if l > 0 {
			err := writeWalletLogToClickHouse(list)
			if err != nil {
				log.Println("[ Job]: handle wallet log write error", err.Error())
				time.Sleep(3 * time.Second)
				continue
			}
			lastId = list[len(list)-1].Id
			if err = job.UpdateExecCursor(int(lastId)); err == nil {
				err = job.Save()
			}
			if err != nil {
				log.Println("[ Job]: handle wallet log write error", err.Error())
				break
			}
		}
		if l < size {
			break
		}
	}
}

func writeWalletLogToClickHouse(list []*wallet.WalletLog) error {
	conn := clickhouse.GetClickhouseConn()
	batch, err := conn.PrepareBatch(context.TODO(),
		`INSERT INTO go2o_wal_wallet_log (
id,wallet_id,wallet_user,kind,title,outer_chan,
outer_no,change_value,balance,procedure_fee,
opr_uid,opr_name,account_no,
account_name,bank_name,audit_state,
audit_remark,audit_time,remark,create_time,
update_time)`)
	if err != nil {
		return err
	}
	for _, l := range list {
		if err = batch.Append(
			l.Id, l.WalletId, l.WalletUser,
			int32(l.Kind),
			l.Subject,
			l.OuterChan,
			l.OuterNo,
			l.ChangeValue,
			l.Balance,
			int32(l.ProcedureFee),
			int64(l.OperatorUid),
			l.OperatorName,
			l.AccountNo,
			l.AccountName,
			l.BankName,
			int32(l.AuditState),
			l.AuditRemark,
			l.AuditTime,
			l.Remark,
			l.CreateTime,
			l.UpdateTime,
		); err != nil {
			return err
		}
	}
	return batch.Send()
}
