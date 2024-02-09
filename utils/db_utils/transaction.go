package db_utils

import (
	"IM/internel/db"
	"IM/internel/db/ent"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

func WithTx(ctx context.Context, tx *ent.Tx, fn func(tx *ent.Tx) error) error {
	var err error
	if tx == nil {
		tx, err = db.DB.Tx(ctx)
		if err != nil {
			fmt.Printf("ctx's tx err: %v", err)
			return err
		}
		defer func() {
			// check if anything went wrong
			if v := recover(); v != nil {
				if err = tx.Rollback(); err != nil {
					logrus.Errorf("tx rollback: %v", err)
					return
				}
				panic(v)
			}
		}()
		err = fn(tx)
		// run the process
		if err != nil {
			// process go wrong
			if rErr := tx.Rollback(); rErr != nil {
				// rollback go wrong
				err = fmt.Errorf("%w: rolling back transaction: %v", err, rErr)
			}
			return err
		}
		if err = tx.Commit(); err != nil {
			// commit go wrong
			return fmt.Errorf("committing transaction: %w", err)
		}
		return nil
	} else {
		err = fn(tx)
		return err
	}
}
