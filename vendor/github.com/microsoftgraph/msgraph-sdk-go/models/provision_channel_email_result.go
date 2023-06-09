package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisionChannelEmailResult 
type ProvisionChannelEmailResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Represents the provisioned email address.
    email *string
    // The OdataType property
    odataType *string
}
// NewProvisionChannelEmailResult instantiates a new provisionChannelEmailResult and sets the default values.
func NewProvisionChannelEmailResult()(*ProvisionChannelEmailResult) {
    m := &ProvisionChannelEmailResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateProvisionChannelEmailResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisionChannelEmailResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisionChannelEmailResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProvisionChannelEmailResult) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEmail gets the email property value. Represents the provisioned email address.
func (m *ProvisionChannelEmailResult) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisionChannelEmailResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["email"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetEmail)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ProvisionChannelEmailResult) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ProvisionChannelEmailResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProvisionChannelEmailResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEmail sets the email property value. Represents the provisioned email address.
func (m *ProvisionChannelEmailResult) SetEmail(value *string)() {
    m.email = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ProvisionChannelEmailResult) SetOdataType(value *string)() {
    m.odataType = value
}
