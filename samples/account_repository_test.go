package account

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/tonybka/go-base-ddd/domain/event"
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

func randomString() string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 10)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func (ts *AccountRepositoryTestSuite) SetupSuite() {
	sqliteConn, err := tests.NewSqliteDBConnect()
	require.NoError(ts.T(), err)

	ts.sqliteConnect = sqliteConn
	ts.dbConn = sqliteConn.Connection()

	ts.dbConn.AutoMigrate(&AccountModel{})
	ts.accountRepo = NewAccountRepository(ts.dbConn)

	// Init global domain publisher
	event.InitDomainEventPublisher()
	publisher := event.GetDomainEventPublisher()

	// Register handlers of domain event
	accountCreatedSubscribers := []event.IDomainEvenHandler{&AccountCreatedEventHandler{}}
	publisher.RegisterSubscriber(&AccountCreatedEvent{}, accountCreatedSubscribers...)

	// Reset random seed to make sure the generated value is unique
	rand.Seed(time.Now().UnixNano())
}

func (ts *AccountRepositoryTestSuite) TestCreateAccount() {
	randId := rand.Intn(99999)

	account := AccountModel{
		BaseModel:   model.NewBaseModel(uint(randId)),
		AccountName: randomString(),
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

func (ts *AccountRepositoryTestSuite) TestAccountWithEvent() {
	randId := rand.Intn(99999)

	account := AccountModel{
		BaseModel:   model.NewBaseModel(uint(randId)),
		AccountName: randomString(),
	}

	account.AddEvent(&AccountCreatedEvent{})

	err := ts.accountRepo.Create(account)
	ts.NoError(err)
}

func (ts *AccountRepositoryTestSuite) TearDownSuite() {
	err := ts.sqliteConnect.CleanUp()
	ts.NoError(err)
}

func TestSuiteRunnerAccountRepository(t *testing.T) {
	ts := new(AccountRepositoryTestSuite)
	suite.Run(t, ts)
}
