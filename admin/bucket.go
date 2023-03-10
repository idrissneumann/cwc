package admin

import (
	"net/mail"
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) AdminAddBucket(user_email string, name string, reg_type string) (*Bucket, error) {
	buf := bytes.Buffer{}
	bucket := Bucket{
		Name:  name,
		Type:  reg_type,
		Email: user_email,
	}

	err := json.NewEncoder(&buf).Encode(bucket)
	if err != nil {
		return nil, err
	}
	respBody, err := c.httpRequest(fmt.Sprintf("/admin/bucket/%s/%s/provision", c.provider, c.region), "POST", buf)
	if err != nil {
		return nil, err
	}
	created_bucket := &Bucket{}
	err = json.NewDecoder(respBody).Decode(created_bucket)
	if err != nil {
		return nil, err
	}
	return created_bucket, nil
}

func (c *Client) GetAllBuckets() (*[]Bucket, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/bucket/%s/%s/all", c.provider, c.region), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	buckets := []Bucket{}
	err = json.NewDecoder(body).Decode(&buckets)

	if err != nil {
		return nil, err
	}
	return &buckets, nil
}

func (c *Client) GetBucket(bucket_id string) (*Bucket, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/bucket/%s", bucket_id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	bucket := &Bucket{}
	err = json.NewDecoder(body).Decode(bucket)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

func valid(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil
}

func (c *Client) UpdateBucket(id string, email string) error {
	buf := bytes.Buffer{}
	var updateCreds bool = true

	renew := RenewCredentials{
		Email: email,
		UpdateCreds: updateCreds,
	}

	if len(email) > 0 && valid(email) {
		updateCreds = false
	}

	encode_err := json.NewEncoder(&buf).Encode(renew)
	if encode_err != nil {
		return encode_err
	}

	_, err := c.httpRequest(fmt.Sprintf("/admin/bucket/%s", id), "PATCH", buf)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteBucket(bucketId string) error {
	_, err := c.httpRequest(fmt.Sprintf("/admin/bucket/%s", bucketId), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
