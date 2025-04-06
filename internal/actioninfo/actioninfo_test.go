package actioninfo

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDataParser is a mock implementation of the DataParser interface
type MockDataParser struct {
	mock.Mock
}

func (m *MockDataParser) Parse(data string) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockDataParser) ActionInfo() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func TestInfo(t *testing.T) {
	// Set up log output capture
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

	tests := []struct {
		name           string
		dataset        []string
		setup          func(*MockDataParser)
		logMsgExpected bool
	}{
		{
			name:    "happy path - single item",
			dataset: []string{"test data"},
			setup: func(m *MockDataParser) {
				m.On("Parse", "test data").Return(nil)
				m.On("ActionInfo").Return("processed test data", nil)
			},
			logMsgExpected: false,
		},
		{
			name:    "parse error",
			dataset: []string{"invalid data"},
			setup: func(m *MockDataParser) {
				m.On("Parse", "invalid data").Return(assert.AnError)
			},
			logMsgExpected: true,
		},
		{
			name:           "empty dataset",
			dataset:        []string{},
			setup:          func(m *MockDataParser) {},
			logMsgExpected: false,
		},
		{
			name:    "action info error",
			dataset: []string{"data1"},
			setup: func(m *MockDataParser) {
				m.On("Parse", "data1").Return(nil)
				m.On("ActionInfo").Return("", assert.AnError)
			},
			logMsgExpected: true,
		},
		{
			name:    "multiple items",
			dataset: []string{"data1", "data2", "data3"},
			setup: func(m *MockDataParser) {
				m.On("Parse", "data1").Return(nil)
				m.On("Parse", "data2").Return(assert.AnError)
				m.On("Parse", "data3").Return(nil)
				m.On("ActionInfo").Return("processed data1", nil).Once()
				m.On("ActionInfo").Return("processed data3", nil).Once()
			},
			logMsgExpected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf.Reset()
			mockParser := new(MockDataParser)
			tt.setup(mockParser)

			Info(tt.dataset, mockParser)

			mockParser.AssertExpectations(t)
			if tt.logMsgExpected {
				assert.NotEmpty(t, buf.String(), "В случае ошибки сообщение о ней должно быть в логе")
				return
			}
			assert.Empty(t, buf.String(), "В случае успешного выполнения сообщение в логе не должно быть (получено сообщение: %q)", buf.String())
		})
	}
}
