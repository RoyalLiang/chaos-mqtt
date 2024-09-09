package service

import (
	"bytes"
	"fmt"
	"reflect"
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

	var message string
	var err error
	if template == "" {
		message = reflect.ValueOf(param).String()
	} else {
		message, err = getTemplateMessage(template, param)
		if err != nil {
			return err
		}
	}

	c, err := NewMQTTClientWithConfig("assigned")
	if err != nil {
		return err
	}

	defer c.Close()
	if err = c.Publish(topic, message); err != nil {
		return err
	}

	return nil
}
