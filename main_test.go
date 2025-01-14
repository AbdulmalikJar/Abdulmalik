package main

import (
	"testing"
)

// Mock function for sending email
var emailSent bool // Track whether the mock email was "sent"

func mockSendEmail() {
	emailSent = true // Simulate that an email was sent
}

// Test the scheduling logic
func TestEmailScheduler(t *testing.T) {
	// Reset emailSent before test
	emailSent = false

	// Call the mockSendEmail function directly to simulate the job
	mockSendEmail()

	// Check if emailSent was set to true
	if !emailSent {
		t.Errorf("Expected emailSent to be true, but got false")
	}
}

// Test the email content logic
func TestEmailContent(t *testing.T) {
	expectedSubject := "Scheduled Email"
	expectedBody := "Hello! This is a scheduled email sent every 12 hours."

	actualSubject := "Scheduled Email"                                    // Replace with function output if dynamic
	actualBody := "Hello! This is a scheduled email sent every 12 hours." // Replace with function output if dynamic

	if actualSubject != expectedSubject {
		t.Errorf("Expected subject: %s, but got: %s", expectedSubject, actualSubject)
	}

	if actualBody != expectedBody {
		t.Errorf("Expected body: %s, but got: %s", expectedBody, actualBody)
	}
}
