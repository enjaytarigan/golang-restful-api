package common

import "testing"

func TestNewSuccessResponseWithoutData(t *testing.T) {
	successResponse := NewSuccessResponse(nil)
	expectedStatus := true
	expectedData := new(interface{})

	if successResponse.Status != true {
		t.Errorf("expected status: %v but got %v", expectedStatus, successResponse.Status)
	}

	if successResponse.Data != nil {
		t.Errorf("expected daa: %v but got %v", expectedData, successResponse.Data)
	}
}

func TestNewSuccessResponseWithData(t *testing.T) {
	data := struct {
		userId int
	}{
		userId: 1,
	}
	successResponse := NewSuccessResponse(data)

	if successResponse.Status != true {
		t.Errorf("expected status: %v but got %v", true, successResponse.Status)

	}

	if successResponse.Data == nil {
		t.Errorf("expected data: %v but got: %v", data, nil)
	}
}
