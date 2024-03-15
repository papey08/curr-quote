package app

import (
	"context"
	apimocks "curr-quote/internal/app/exchange_mocks"
	repomocks "curr-quote/internal/app/repo_mocks"
	"curr-quote/internal/model"
	"curr-quote/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
	"time"
)

//go:generate mockery --dir=../adapters/exchange --output=./exchange_mocks --name=Exchange
//go:generate mockery --dir=../repo --output=./repo_mocks --name=Repo

type appTestSuite struct {
	suite.Suite
	repo    *repomocks.Repo
	api     *apimocks.Exchange
	logs    logger.Logger
	service App
}

func (s *appTestSuite) SetupSuite() {
	s.repo = new(repomocks.Repo)
	s.api = new(apimocks.Exchange)

	s.setApiMocks()
	s.setRepoMocks()

	s.service = New(context.Background(), s.api, s.repo, logger.New())
}

func (s *appTestSuite) TearDownSuite() {}

func (s *appTestSuite) setApiMocks() {
	s.api.On("GetLatestQuote", mock.Anything, model.EUR).Return(eurQuote, error(nil))
	s.api.On("GetLatestQuote", mock.Anything, model.USD).Return(usdQuote, error(nil))
	s.api.On("GetLatestQuote", mock.Anything, model.MXN).Return(mxnQuote, error(nil))
}

func (s *appTestSuite) setRepoMocks() {
	s.repo.On("SetQuote", mock.Anything, mock.Anything, model.EUR, mock.Anything).Return(error(nil))
	s.repo.On("SetQuote", mock.Anything, mock.Anything, model.USD, mock.Anything).Return(error(nil))
	s.repo.On("SetQuote", mock.Anything, mock.Anything, model.MXN, mock.Anything).Return(error(nil))
}

type getQuoteMock struct {
	id    string
	curr  model.Currency
	quote model.Quote
	err   error
}

type getQuoteByIdTest struct {
	description string
	id          string
	value       float64
	err         error
}

func (s *appTestSuite) TestGetQuoteById() {
	idEurMxn, _ := s.service.RefreshQuote(context.Background(), model.EUR, model.MXN)
	idMxnUsd, _ := s.service.RefreshQuote(context.Background(), model.MXN, model.USD)
	idUsdUsd, _ := s.service.RefreshQuote(context.Background(), model.USD, model.USD)

	var mocks = []getQuoteMock{
		{
			id:    strings.Split(idEurMxn, "-")[2],
			curr:  model.EUR,
			quote: eurQuote,
			err:   nil,
		},
		{
			id:    strings.Split(idMxnUsd, "-")[2],
			curr:  model.MXN,
			quote: mxnQuote,
			err:   nil,
		},
		{
			id:    strings.Split(idUsdUsd, "-")[2],
			curr:  model.USD,
			quote: usdQuote,
			err:   nil,
		},
		{
			id:    "aaaa",
			curr:  model.USD,
			quote: model.Quote{},
			err:   model.ErrQuoteNotFound,
		},
	}

	for _, m := range mocks {
		s.repo.On("GetQuote", mock.Anything, m.id, m.curr).Return(m.quote, m.err)
	}

	var tests = []getQuoteByIdTest{
		{
			description: "successful getting of EUR/MXN quote",
			id:          idEurMxn,
			value:       eurQuote.Data[model.MXN],
			err:         nil,
		},
		{
			description: "successful getting of MXN/USD quote",
			id:          idMxnUsd,
			value:       mxnQuote.Data[model.USD],
			err:         nil,
		},
		{
			description: "successful getting of USD/USD quote",
			id:          idUsdUsd,
			value:       usdQuote.Data[model.USD],
			err:         nil,
		},
		{
			description: "getting of non existing quote",
			id:          "USD-MXN-aaaa",
			value:       0,
			err:         model.ErrQuoteNotFound,
		},
		{
			description: "getting of quote with invalid id",
			id:          "aaaa",
			value:       0,
			err:         model.ErrInvalidId,
		},
	}

	for _, test := range tests {
		s.T().Run(test.description, func(t *testing.T) {
			quote, err := s.service.GetQuoteById(context.Background(), test.id)
			assert.Equal(s.T(), test.value, quote.Value)
			assert.ErrorIs(s.T(), err, test.err)
		})
	}
}

type getLastQuoteTest struct {
	description string
	baseCurr    model.Currency
	otherCurr   model.Currency
	value       float64
	err         error
}

func (s *appTestSuite) TestGetLastQuote() {
	var tests = []getLastQuoteTest{
		{
			description: "successful getting of EUR/USD quote",
			baseCurr:    model.EUR,
			otherCurr:   model.USD,
			value:       eurQuote.Data[model.USD],
			err:         nil,
		},
		{
			description: "successful getting of USD/EUR quote",
			baseCurr:    model.USD,
			otherCurr:   model.EUR,
			value:       usdQuote.Data[model.EUR],
			err:         nil,
		},
		{
			description: "getting of quote of non existing currency",
			baseCurr:    model.EUR,
			otherCurr:   "RUB",
			value:       0,
			err:         model.ErrInvalidCurr,
		},
	}

	// блокировка нужна, чтобы котировки в приложении успели обновиться
	time.Sleep(time.Second * 5)

	for _, test := range tests {
		s.T().Run(test.description, func(t *testing.T) {
			quote, err := s.service.GetLastQuote(context.Background(), test.baseCurr, test.otherCurr)
			assert.Equal(s.T(), test.value, quote.Value)
			assert.ErrorIs(s.T(), err, test.err)
		})
	}
}

var (
	eurQuote = model.Quote{
		Data: map[model.Currency]float64{
			model.EUR: 1,
			model.USD: 1.09,
			model.MXN: 18.23,
		},
		RefreshTime: time.Now().UTC(),
	}

	usdQuote = model.Quote{
		Data: map[model.Currency]float64{
			model.EUR: 0.91,
			model.USD: 1,
			model.MXN: 16.72,
		},
		RefreshTime: time.Now().UTC(),
	}

	mxnQuote = model.Quote{
		Data: map[model.Currency]float64{
			model.EUR: 0.05,
			model.USD: 0.06,
			model.MXN: 1,
		},
		RefreshTime: time.Now().UTC(),
	}
)

func TestAppTestSuite(t *testing.T) {
	suite.Run(t, new(appTestSuite))
}
