package console

import (
	"database/sql"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/impal-lms/lms-backend/helper"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	Long:  `This subcommand used to migrate database`,
	Run:   processMigration,
}

func init() {
	migrateCmd.PersistentFlags().Int("step", 0, "maximum migration steps")
	migrateCmd.PersistentFlags().String("direction", "up", "migration direction")
	RootCmd.AddCommand(migrateCmd)
}

func processMigration(cmd *cobra.Command, args []string) {

	direction := cmd.Flag("direction").Value.String()
	stepStr := cmd.Flag("step").Value.String()
	step, err := strconv.Atoi(stepStr)
	if err != nil {
		logrus.WithField("stepStr", stepStr).Fatal("Failed to parse step to int: ", err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "./migration",
	}

	migrate.SetTable("schema_migrations")
	dbURL := helper.GetEnv("DB_URL", "").(string)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		logrus.WithField("DatabaseDSN", dbURL).Fatal("Failed to connect database: ", err)
	}

	var n int
	if direction == "down" {
		n, err = migrate.ExecMax(db, "postgres", migrations, migrate.Down, step)
	} else {
		n, err = migrate.ExecMax(db, "postgres", migrations, migrate.Up, step)
	}
	if err != nil {
		logrus.Fatal("Failed to migrate database: ", err)
	}

	logrus.Infof("Applied %d migrations!\n", n)

}
