package account

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/tonybka/go-base-persistence/model"
	"github.com/tonybka/go-base-persistence/tests"
	"gorm.io/gorm"
)

type AccountRepositoryTestSuite struct {
	suite.Suite

	sqliteConnect *tests.SqliteDBConnect
	dbConn        *gorm.DB

	accountRepo *AccountRepository
}

func (ts *AccountRepositoryTestSuite) SetupSuite() {
	sqliteConn, err := tests.NewSqliteDBConnect()
	require.NoError(ts.T(), err)

	ts.sqliteConnect = sqliteConn
	ts.dbConn = sqliteConn.Connection()

	ts.dbConn.AutoMigrate(&AccountModel{})

	ts.accountRepo = NewAccountRepository(ts.dbConn)

	model.InitDomainEventPublisher()
}

func (ts *AccountRepositoryTestSuite) TestCreateAccount() {
	randId := rand.Intn(99999)

	account := AccountModel{
		BaseModel:   model.NewBaseModel(uint(randId)),
		AccountName: "abc",
	}

	err := ts.accountRepo.Create(account)
	ts.NoError(err)

	all, err := ts.accountRepo.GetAll()
	ts.NoError(err)
	ts.Greater(len(all), 0)

	queriedAccount, err := ts.accountRepo.FindById(account.ID)
	ts.NoError(err)
	ts.Equal(account.AccountName, queriedAccount.AccountName)
	ts.Equal(account.ID, queriedAccount.ID)
}

func (ts *AccountRepositoryTestSuite) TearDownSuite() {
	err := ts.sqliteConnect.CleanUp()
	ts.NoError(err)
}

func TestSuiteRunnerAccountRepository(t *testing.T) {
	ts := new(AccountRepositoryTestSuite)
	suite.Run(t, ts)
}
