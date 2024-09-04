package service

import (
	"bytes"
	"fmt"
	"text/template"
)

func getTemplateMessage(content string, param interface{}) (string, error) {
	t, _ := template.New("").Parse(content)

	var buf bytes.Buffer
	if err := t.Execute(&buf, &param); err != nil {
		fmt.Println("parse template error: ", err)
		return "", err
	}
	return buf.String(), nil
}

func PublishAssignedTopic(topic, template string, param interface{}) error {
	message, err := getTemplateMessage(template, param)
	if err != nil {
		return err
	}
	return MQTTClient.Publish(topic, message)
}

func PublishJobInstruction() error {
	return nil
}

func PublishDockPosition() error {
	return nil
}

func PublishCallInRequest() error {
	return nil
}

func PublishVesselBerth() error {
	return nil
}

func PublishMoveToQC() error {
	return nil
}

func PublishIngressToCallIn() error {
	return nil
}

func PublishCancelJob() error {
	return nil
}
