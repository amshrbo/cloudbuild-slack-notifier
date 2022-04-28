package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/slack-go/slack"
	cbpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"
)

func TestWriteMessage(t *testing.T) {
	n := new(slackNotifier)
	b := &cbpb.Build{
		ProjectId: "my-project-id",
		Id:        "some-build-id",
		Status:    cbpb.Build_SUCCESS,
		LogUrl:    "https://some.example.com/log/url?foo=bar",
	}

	got, err := n.writeMessage(b)
	if err != nil {
		t.Fatalf("writeMessage failed: %v", err)
	}

	want := &slack.WebhookMessage{
		Attachments: []slack.Attachment{{
			AuthorName: "ahmed.shrbiny",
			AuthorLink: "https://gitlab.com/ahmed.shrbiny",
			Text:  "Build Status: SUCCESS,\nRepository: ,\nBranch: ,\nCommit: ",
			Color: "good",
			Actions: []slack.AttachmentAction{
				slack.AttachmentAction { 	
					Text: "View Logs",
					Type: "button",
					URL: "https://some.example.com/log/url?foo=bar&utm_campaign=google-cloud-build-notifiers&utm_medium=chat&utm_source=google-cloud-build",
				}, 
				slack.AttachmentAction {
					Text: "View Logs +",
					Type: "button",
					URL:  "https://some.example.com/log/url?foo=bar&utm_campaign=google-cloud-build-notifiers&utm_medium=chat&utm_source=google-cloud-build",
				},
				slack.AttachmentAction {
					Text: "Additional btn",
					Type: "button",
					URL:  "https://some.example.com/log/url?foo=bar&utm_campaign=google-cloud-build-notifiers&utm_medium=chat&utm_source=google-cloud-build",
				},
			},
		}},
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("writeMessage got unexpected diff: %s", diff)
	}
}
